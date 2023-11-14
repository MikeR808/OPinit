import { getDB } from './db';
import { DataSource } from 'typeorm';
import { batchLogger as logger } from 'lib/logger';

import { BlockBulk, getBlockBulk } from 'lib/rpc';
import { compressor } from 'lib/compressor';
import { getLatestBlockHeight } from 'lib/tx';
import { RecordEntity } from 'orm';
import { Wallet, MnemonicKey, MsgExecute, BCS } from '@initia/initia.js';
import { fetchBridgeConfig } from 'lib/lcd';
import { delay } from 'bluebird';
import { INTERVAL_BATCH } from 'config';
import { getConfig } from 'config';
import { sendTx } from 'lib/tx';

const config = getConfig();
const bcs = BCS.getInstance();

export class BatchSubmitter {
  private batchIndex = 0;
  private batchL2StartHeight: number;
  private latestBlockHeight: number;
  private dataSource: DataSource;
  private submitter: Wallet;
  private submissionInterval: number;
  private isRunning = false;

  async init() {
    [this.dataSource] = getDB();
    this.latestBlockHeight = await getLatestBlockHeight(config.l2lcd);
    this.submitter = new Wallet(
      config.l1lcd,
      new MnemonicKey({ mnemonic: config.BATCH_SUBMITTER_MNEMONIC })
    );
    const bridgeCfg = await fetchBridgeConfig();
    this.batchL2StartHeight = parseInt(bridgeCfg.starting_block_number);
    this.submissionInterval = parseInt(bridgeCfg.submission_interval);
    this.isRunning = true;
  }

  public stop() {
    this.isRunning = false;
  }

  public async run() {
    await this.init();

    while (this.isRunning) {
      try {
        const latestBatch = await this.getStoredBatch(this.dataSource);
        if (latestBatch) {
          this.batchIndex = latestBatch.batchIndex + 1;
        }

        // e.g [start_height + 0, start_height + 99], [start_height + 100, start_height + 199], ...
        const startHeight =
          this.batchL2StartHeight + this.batchIndex * this.submissionInterval;
        const endHeight =
          this.batchL2StartHeight +
          (this.batchIndex + 1) * this.submissionInterval -
          1;

        this.latestBlockHeight = await getLatestBlockHeight(config.l2lcd);
        if (endHeight > this.latestBlockHeight) {
          await delay(INTERVAL_BATCH);
          continue;
        }

        const batch = await this.getBatch(startHeight, endHeight);
        await this.publishBatchToL1(batch);
        await this.saveBatchToDB(this.dataSource, batch, this.batchIndex);
        logger.info(`${this.batchIndex}th batch is successfully saved`);
      } catch (err) {
        throw new Error(`Error in BatchSubmitter: ${err}`);
      }
    }
  }

  // Get [start, end] batch from L2
  async getBatch(start: number, end: number): Promise<Buffer> {
    const bulk: BlockBulk | null = await getBlockBulk(
      start.toString(),
      end.toString()
    );
    if (!bulk) {
      throw new Error(`Error getting block bulk from L2`);
    }

    return compressor(bulk.blocks);
  }

  async getStoredBatch(db: DataSource): Promise<RecordEntity | null> {
    const storedRecord = await db
      .getRepository(RecordEntity)
      .find({
        order: {
          batchIndex: 'DESC'
        },
        take: 1
      })
      .catch((err) => {
        logger.error(`Error getting stored batch: ${err}`);
        return null;
      });

    return storedRecord ? storedRecord[0] : null;
  }

  // Publish a batch to L1
  async publishBatchToL1(batch: Buffer) {
    try {
      const executeMsg = new MsgExecute(
        this.submitter.key.accAddress,
        '0x1',
        'op_batch_inbox',
        'record_batch',
        [config.L2ID],
        [bcs.serialize('vector<u8>', batch, this.submissionInterval * 1000)]
      );

      return await sendTx(this.submitter, [executeMsg]);
    } catch (err) {
      throw new Error(`Error publishing batch to L1: ${err}`);
    }
  }

  // Save batch record to database
  async saveBatchToDB(
    db: DataSource,
    batch: Buffer,
    batchIndex: number
  ): Promise<RecordEntity> {
    const record = new RecordEntity();

    record.l2Id = config.L2ID;
    record.batchIndex = batchIndex;
    record.batch = batch;

    await db
      .getRepository(RecordEntity)
      .save(record)
      .catch((error) => {
        throw new Error(
          `Error saving record ${record.l2Id} batch ${batchIndex} to database: ${error}`
        );
      });

    return record;
  }
}