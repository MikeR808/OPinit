// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: op_init/op_child/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the rollup module's genesis state.
type GenesisState struct {
	// params defines all the parameters of related to deposit.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// last_validator_powers is a special index that provides a historical list
	// of the last-block's bonded validators.
	LastValidatorPowers []LastValidatorPower `protobuf:"bytes,2,rep,name=last_validator_powers,json=lastValidatorPowers,proto3" json:"last_validator_powers"`
	// delegations defines the validator set at genesis.
	Validators []Validator `protobuf:"bytes,3,rep,name=validators,proto3" json:"validators"`
	// fee_pool defines the previous proposer at genesis.
	PreviousProposer string `protobuf:"bytes,4,opt,name=previous_proposer,json=previousProposer,proto3" json:"previous_proposer,omitempty"`
	Exported         bool   `protobuf:"varint,5,opt,name=exported,proto3" json:"exported,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f99f91fe3495431, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetLastValidatorPowers() []LastValidatorPower {
	if m != nil {
		return m.LastValidatorPowers
	}
	return nil
}

func (m *GenesisState) GetValidators() []Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *GenesisState) GetPreviousProposer() string {
	if m != nil {
		return m.PreviousProposer
	}
	return ""
}

func (m *GenesisState) GetExported() bool {
	if m != nil {
		return m.Exported
	}
	return false
}

// LastValidatorPower required for validator set update logic.
type LastValidatorPower struct {
	// address is the address of the validator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// power defines the power of the validator.
	Power int64 `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (m *LastValidatorPower) Reset()         { *m = LastValidatorPower{} }
func (m *LastValidatorPower) String() string { return proto.CompactTextString(m) }
func (*LastValidatorPower) ProtoMessage()    {}
func (*LastValidatorPower) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f99f91fe3495431, []int{1}
}
func (m *LastValidatorPower) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastValidatorPower) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastValidatorPower.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastValidatorPower) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastValidatorPower.Merge(m, src)
}
func (m *LastValidatorPower) XXX_Size() int {
	return m.Size()
}
func (m *LastValidatorPower) XXX_DiscardUnknown() {
	xxx_messageInfo_LastValidatorPower.DiscardUnknown(m)
}

var xxx_messageInfo_LastValidatorPower proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "op_init.op_child.v1.GenesisState")
	proto.RegisterType((*LastValidatorPower)(nil), "op_init.op_child.v1.LastValidatorPower")
}

func init() { proto.RegisterFile("op_init/op_child/v1/genesis.proto", fileDescriptor_6f99f91fe3495431) }

var fileDescriptor_6f99f91fe3495431 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0xcd, 0xb4, 0xee, 0xda, 0x9d, 0xf5, 0xe0, 0xce, 0x56, 0x88, 0x15, 0x92, 0xd8, 0x8b, 0x41,
	0x68, 0x42, 0xeb, 0xcd, 0x83, 0x60, 0x41, 0x54, 0x10, 0x0c, 0x29, 0x78, 0xf0, 0x12, 0xa6, 0xcd,
	0x98, 0x0e, 0x24, 0x99, 0x61, 0x66, 0x1a, 0xeb, 0x3f, 0xf0, 0xe8, 0x4f, 0xe8, 0xd1, 0xa3, 0x07,
	0x7f, 0x44, 0x8f, 0xc5, 0x93, 0x27, 0x91, 0x56, 0xd0, 0x9f, 0x21, 0x9d, 0x24, 0xb5, 0xd0, 0xe0,
	0x25, 0xe4, 0xcd, 0xbc, 0xf7, 0xbe, 0x6f, 0x1e, 0x0f, 0xde, 0x67, 0x3c, 0xa2, 0x39, 0x55, 0x3e,
	0xe3, 0xd1, 0x6c, 0x4e, 0xd3, 0xd8, 0x2f, 0x86, 0x7e, 0x42, 0x72, 0x22, 0xa9, 0xf4, 0xb8, 0x60,
	0x8a, 0xa1, 0xeb, 0x8a, 0xe2, 0xd5, 0x14, 0xaf, 0x18, 0xf6, 0xae, 0x70, 0x46, 0x73, 0xe6, 0xeb,
	0x6f, 0xc9, 0xeb, 0xdd, 0x9d, 0x31, 0x99, 0x31, 0x19, 0x69, 0xe4, 0x97, 0xa0, 0xba, 0xea, 0x26,
	0x2c, 0x61, 0xe5, 0xf9, 0xfe, 0xaf, 0x3a, 0xb5, 0x9b, 0x66, 0xab, 0x0f, 0x9c, 0x54, 0xb2, 0xfe,
	0xaf, 0x16, 0xbc, 0xf5, 0xbc, 0xdc, 0x65, 0xa2, 0xb0, 0x22, 0xe8, 0x09, 0x3c, 0xe7, 0x58, 0xe0,
	0x4c, 0x9a, 0xc0, 0x01, 0xee, 0xe5, 0xe8, 0x9e, 0xd7, 0xb0, 0x9b, 0x17, 0x68, 0xca, 0xf8, 0x62,
	0xfd, 0xc3, 0x36, 0x3e, 0xff, 0xfe, 0xf2, 0x10, 0x84, 0x95, 0x0a, 0xbd, 0x83, 0x77, 0x52, 0x2c,
	0x55, 0x54, 0xe0, 0x94, 0xc6, 0x58, 0x31, 0x11, 0x71, 0xf6, 0x9e, 0x08, 0x69, 0xb6, 0x9c, 0xb6,
	0x7b, 0x39, 0x7a, 0xd0, 0x68, 0xf7, 0x0a, 0x4b, 0xf5, 0xa6, 0x16, 0x04, 0x7b, 0xfe, 0xb1, 0xf5,
	0x75, 0x7a, 0x72, 0x2d, 0xd1, 0x4b, 0x08, 0x0f, 0x23, 0xa4, 0xd9, 0xd6, 0xe6, 0x56, 0xa3, 0xf9,
	0x41, 0x79, 0xec, 0x79, 0x24, 0x46, 0xcf, 0xe0, 0x15, 0x17, 0xa4, 0xa0, 0x6c, 0xa1, 0x93, 0xe5,
	0x4c, 0x12, 0x61, 0xde, 0x70, 0x80, 0x7b, 0x31, 0x36, 0xbf, 0x7d, 0x1d, 0x74, 0xab, 0x9c, 0x9f,
	0xc6, 0xb1, 0x20, 0x52, 0x4e, 0x94, 0xa0, 0x79, 0x12, 0xde, 0xae, 0x25, 0x41, 0xa5, 0x40, 0x3d,
	0xd8, 0x21, 0x4b, 0xce, 0x84, 0x22, 0xb1, 0x79, 0xe6, 0x00, 0xb7, 0x13, 0x1e, 0x70, 0x7f, 0x0e,
	0xd1, 0xe9, 0x1b, 0xd1, 0x08, 0xde, 0xc4, 0xa5, 0xa9, 0x0e, 0xfb, 0x7f, 0xe3, 0x6a, 0x22, 0xea,
	0xc2, 0x33, 0x1d, 0xa8, 0xd9, 0x72, 0x80, 0xdb, 0x0e, 0x4b, 0xf0, 0xb8, 0xf3, 0x71, 0x65, 0x1b,
	0x7f, 0x56, 0xb6, 0x31, 0x7e, 0xb1, 0xde, 0x5a, 0x60, 0xb3, 0xb5, 0xc0, 0xcf, 0xad, 0x05, 0x3e,
	0xed, 0x2c, 0x63, 0xb3, 0xb3, 0x8c, 0xef, 0x3b, 0xcb, 0x78, 0xeb, 0x25, 0x54, 0xcd, 0x17, 0x53,
	0x6f, 0xc6, 0x32, 0x7f, 0x1f, 0x12, 0xc5, 0x83, 0x14, 0x4f, 0xa5, 0xff, 0x3a, 0xd0, 0x0d, 0x59,
	0xfe, 0xeb, 0x88, 0x2e, 0xc8, 0xf4, 0x5c, 0x37, 0xe4, 0xd1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xa6, 0x3d, 0xbe, 0xe4, 0xc0, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Exported {
		i--
		if m.Exported {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.PreviousProposer) > 0 {
		i -= len(m.PreviousProposer)
		copy(dAtA[i:], m.PreviousProposer)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PreviousProposer)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.LastValidatorPowers) > 0 {
		for iNdEx := len(m.LastValidatorPowers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LastValidatorPowers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *LastValidatorPower) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastValidatorPower) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastValidatorPower) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Power != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.LastValidatorPowers) > 0 {
		for _, e := range m.LastValidatorPowers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.PreviousProposer)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Exported {
		n += 2
	}
	return n
}

func (m *LastValidatorPower) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovGenesis(uint64(m.Power))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastValidatorPowers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastValidatorPowers = append(m.LastValidatorPowers, LastValidatorPower{})
			if err := m.LastValidatorPowers[len(m.LastValidatorPowers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousProposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousProposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exported", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Exported = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LastValidatorPower) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LastValidatorPower: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastValidatorPower: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
