// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/vesting/vesting.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
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

// BaseVestingAccount implements the VestingAccount interface. It contains all
// the necessary fields needed for any vesting account implementation.
type BaseVestingAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty"`
	OriginalVesting    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=original_vesting,json=originalVesting,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"original_vesting" yaml:"original_vesting"`
	DelegatedFree      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=delegated_free,json=delegatedFree,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"delegated_free" yaml:"delegated_free"`
	DelegatedVesting   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=delegated_vesting,json=delegatedVesting,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"delegated_vesting" yaml:"delegated_vesting"`
	EndTime            int64                                    `protobuf:"varint,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" yaml:"end_time"`
}

func (m *BaseVestingAccount) Reset()      { *m = BaseVestingAccount{} }
func (*BaseVestingAccount) ProtoMessage() {}
func (*BaseVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f0278a453c26b3, []int{0}
}
func (m *BaseVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseVestingAccount.Merge(m, src)
}
func (m *BaseVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *BaseVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_BaseVestingAccount proto.InternalMessageInfo

// Period defines a length of time and amount of coins that will vest.
type Period struct {
	StartTime  int64                                    `protobuf:"varint,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Length     int64                                    `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Amount     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	ActionType int32                                    `protobuf:"varint,4,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
}

func (m *Period) Reset()      { *m = Period{} }
func (*Period) ProtoMessage() {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f0278a453c26b3, []int{1}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Period.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(m, src)
}
func (m *Period) XXX_Size() int {
	return m.Size()
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Period) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Period) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Period) GetActionType() int32 {
	if m != nil {
		return m.ActionType
	}
	return 0
}

// StridePeriodicVestingAccount implements the VestingAccount interface. It
// periodically vests by unlocking coins during each specified period.
type StridePeriodicVestingAccount struct {
	*BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
	VestingPeriods      []Period `protobuf:"bytes,3,rep,name=vesting_periods,json=vestingPeriods,proto3" json:"vesting_periods" yaml:"vesting_periods"`
}

func (m *StridePeriodicVestingAccount) Reset()      { *m = StridePeriodicVestingAccount{} }
func (*StridePeriodicVestingAccount) ProtoMessage() {}
func (*StridePeriodicVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f0278a453c26b3, []int{2}
}
func (m *StridePeriodicVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StridePeriodicVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StridePeriodicVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StridePeriodicVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StridePeriodicVestingAccount.Merge(m, src)
}
func (m *StridePeriodicVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *StridePeriodicVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_StridePeriodicVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_StridePeriodicVestingAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseVestingAccount)(nil), "stride.vesting.BaseVestingAccount")
	proto.RegisterType((*Period)(nil), "stride.vesting.Period")
	proto.RegisterType((*StridePeriodicVestingAccount)(nil), "stride.vesting.StridePeriodicVestingAccount")
}

func init() { proto.RegisterFile("stride/vesting/vesting.proto", fileDescriptor_41f0278a453c26b3) }

var fileDescriptor_41f0278a453c26b3 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x41, 0x6f, 0xd3, 0x3e,
	0x1c, 0x8d, 0xd7, 0xae, 0xff, 0xfd, 0x5d, 0x58, 0x47, 0x18, 0x25, 0x4c, 0x23, 0xa9, 0x72, 0xea,
	0x65, 0x09, 0x1b, 0x12, 0x12, 0xbd, 0x11, 0x24, 0xa4, 0x89, 0x1d, 0xa6, 0x30, 0x71, 0xd8, 0x25,
	0x72, 0x12, 0x93, 0x5a, 0x34, 0x71, 0x15, 0xbb, 0x13, 0xfd, 0x06, 0x48, 0x5c, 0x40, 0xe2, 0xc0,
	0x71, 0x67, 0x3e, 0xc9, 0x24, 0x2e, 0x15, 0x27, 0x4e, 0x05, 0xb5, 0xe2, 0x0b, 0xec, 0x13, 0xa0,
	0xd8, 0x4e, 0xcb, 0xb2, 0x43, 0x35, 0x4e, 0x89, 0xfd, 0xf3, 0x7b, 0x7e, 0xbf, 0x67, 0x3f, 0xc3,
	0x5d, 0xc6, 0x73, 0x12, 0x63, 0xf7, 0x0c, 0x33, 0x4e, 0xb2, 0xa4, 0xfc, 0x3a, 0xc3, 0x9c, 0x72,
	0xaa, 0x6f, 0xca, 0xaa, 0xa3, 0x66, 0x77, 0xb6, 0x13, 0x9a, 0x50, 0x51, 0x72, 0x8b, 0x3f, 0xb9,
	0x6a, 0xc7, 0x8c, 0x28, 0x4b, 0x29, 0x73, 0x43, 0xc4, 0xb0, 0x7b, 0xb6, 0x1f, 0x62, 0x8e, 0xf6,
	0xdd, 0x88, 0x92, 0xac, 0x52, 0x47, 0x23, 0xde, 0x5f, 0xd4, 0x8b, 0x81, 0xac, 0xdb, 0xdf, 0xeb,
	0x50, 0xf7, 0x10, 0xc3, 0xaf, 0xe5, 0x2e, 0xcf, 0xa2, 0x88, 0x8e, 0x32, 0xae, 0x1f, 0xc2, 0x5b,
	0x05, 0x63, 0x80, 0xe4, 0xd8, 0x00, 0x1d, 0xd0, 0x6d, 0x1e, 0x74, 0x1c, 0xc9, 0xe6, 0x08, 0x02,
	0xc5, 0xe6, 0x14, 0x70, 0x85, 0xf3, 0xea, 0x93, 0xa9, 0x05, 0xfc, 0x66, 0xb8, 0x9c, 0xd2, 0x3f,
	0x01, 0xb8, 0x45, 0x73, 0x92, 0x90, 0x0c, 0x0d, 0x02, 0xd5, 0x8c, 0xb1, 0xd6, 0xa9, 0x75, 0x9b,
	0x07, 0x0f, 0x4a, 0xbe, 0x62, 0xfd, 0x82, 0xef, 0x39, 0x25, 0x99, 0xf7, 0xf2, 0x62, 0x6a, 0x69,
	0x97, 0x53, 0xeb, 0xfe, 0x18, 0xa5, 0x83, 0x9e, 0x5d, 0x25, 0xb0, 0xbf, 0xfe, 0xb4, 0xba, 0x09,
	0xe1, 0xfd, 0x51, 0xe8, 0x44, 0x34, 0x75, 0x55, 0x97, 0xf2, 0xb3, 0xc7, 0xe2, 0xb7, 0x2e, 0x1f,
	0x0f, 0x31, 0x13, 0x5c, 0xcc, 0x6f, 0x95, 0x70, 0xd5, 0xa5, 0xfe, 0x01, 0xc0, 0xcd, 0x18, 0x0f,
	0x70, 0x82, 0x38, 0x8e, 0x83, 0x37, 0x39, 0xc6, 0x46, 0x6d, 0x95, 0xa2, 0x43, 0xa5, 0xe8, 0x9e,
	0x54, 0x74, 0x15, 0x7e, 0x33, 0x3d, 0xb7, 0x17, 0xe0, 0x17, 0x39, 0xc6, 0xfa, 0x67, 0x00, 0xef,
	0x2c, 0xe9, 0x4a, 0x8b, 0xea, 0xab, 0x04, 0x1d, 0x29, 0x41, 0x46, 0x55, 0xd0, 0x3f, 0x79, 0xb4,
	0xb5, 0xc0, 0x97, 0x26, 0x39, 0x70, 0x03, 0x67, 0x71, 0xc0, 0x49, 0x8a, 0x8d, 0xf5, 0x0e, 0xe8,
	0xd6, 0xbc, 0xbb, 0x97, 0x53, 0xab, 0x25, 0x77, 0x2b, 0x2b, 0xb6, 0xff, 0x1f, 0xce, 0xe2, 0x13,
	0x92, 0xe2, 0xde, 0xc6, 0xfb, 0x73, 0x4b, 0xfb, 0x72, 0x6e, 0x69, 0xf6, 0x37, 0x00, 0x1b, 0xc7,
	0x38, 0x27, 0x34, 0xd6, 0x1f, 0x42, 0xc8, 0x38, 0xca, 0xb9, 0xa4, 0x29, 0xae, 0x51, 0xcd, 0xff,
	0x5f, 0xcc, 0x14, 0x18, 0xbd, 0x0d, 0x1b, 0x03, 0x9c, 0x25, 0xbc, 0x6f, 0xac, 0x89, 0x92, 0x1a,
	0xe9, 0x11, 0x6c, 0xa0, 0x54, 0xdc, 0xbc, 0x95, 0xe7, 0xf2, 0xa8, 0xb0, 0xe1, 0x46, 0xad, 0x2a,
	0x6a, 0xdd, 0x82, 0x4d, 0x14, 0x71, 0x42, 0xb3, 0xa0, 0xa8, 0x1a, 0xf5, 0x0e, 0xe8, 0xae, 0xfb,
	0x50, 0x4e, 0x9d, 0x8c, 0x87, 0xb8, 0x57, 0x17, 0xdd, 0xfc, 0x06, 0x70, 0xf7, 0x95, 0xc8, 0xa2,
	0xec, 0x89, 0x44, 0x95, 0xb0, 0x9c, 0xc2, 0x6d, 0x11, 0x16, 0xe5, 0x7b, 0x25, 0x34, 0xb6, 0x73,
	0x35, 0xc8, 0xce, 0xf5, 0xb8, 0xa9, 0xd8, 0xe8, 0xe1, 0xf5, 0x20, 0x06, 0xb0, 0x55, 0xd2, 0x0e,
	0xc5, 0xee, 0x4c, 0x39, 0xd2, 0xae, 0xd2, 0x4a, 0x71, 0x9e, 0xa9, 0x6e, 0x45, 0x5b, 0x9e, 0x53,
	0x05, 0x6c, 0xfb, 0x9b, 0x6a, 0x46, 0x2e, 0x67, 0xcb, 0x53, 0xf3, 0x8e, 0x2f, 0x66, 0x26, 0x98,
	0xcc, 0x4c, 0xf0, 0x6b, 0x66, 0x82, 0x8f, 0x73, 0x53, 0x9b, 0xcc, 0x4d, 0xed, 0xc7, 0xdc, 0xd4,
	0x4e, 0x9f, 0xfc, 0x65, 0xad, 0x74, 0x62, 0xef, 0x08, 0x85, 0xcc, 0x2d, 0xdf, 0xaf, 0xa7, 0xee,
	0x3b, 0x37, 0x1a, 0x20, 0x92, 0x2e, 0x9e, 0x32, 0x61, 0x77, 0xd8, 0x10, 0x6f, 0xcc, 0xe3, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x14, 0x07, 0xe2, 0xe9, 0x04, 0x00, 0x00,
}

func (m *BaseVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DelegatedVesting) > 0 {
		for iNdEx := len(m.DelegatedVesting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatedVesting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.DelegatedFree) > 0 {
		for iNdEx := len(m.DelegatedFree) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatedFree[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.OriginalVesting) > 0 {
		for iNdEx := len(m.OriginalVesting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OriginalVesting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Period) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Period) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Period) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ActionType != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.ActionType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Length != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.Length))
		i--
		dAtA[i] = 0x10
	}
	if m.StartTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StridePeriodicVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StridePeriodicVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StridePeriodicVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPeriods) > 0 {
		for iNdEx := len(m.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVesting(dAtA []byte, offset int, v uint64) int {
	offset -= sovVesting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	if len(m.OriginalVesting) > 0 {
		for _, e := range m.OriginalVesting {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if len(m.DelegatedFree) > 0 {
		for _, e := range m.DelegatedFree {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if len(m.DelegatedVesting) > 0 {
		for _, e := range m.DelegatedVesting {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if m.EndTime != 0 {
		n += 1 + sovVesting(uint64(m.EndTime))
	}
	return n
}

func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartTime != 0 {
		n += 1 + sovVesting(uint64(m.StartTime))
	}
	if m.Length != 0 {
		n += 1 + sovVesting(uint64(m.Length))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if m.ActionType != 0 {
		n += 1 + sovVesting(uint64(m.ActionType))
	}
	return n
}

func (m *StridePeriodicVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	if len(m.VestingPeriods) > 0 {
		for _, e := range m.VestingPeriods {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	return n
}

func sovVesting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVesting(x uint64) (n int) {
	return sovVesting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: BaseVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalVesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalVesting = append(m.OriginalVesting, types1.Coin{})
			if err := m.OriginalVesting[len(m.OriginalVesting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatedFree", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatedFree = append(m.DelegatedFree, types1.Coin{})
			if err := m.DelegatedFree[len(m.DelegatedFree)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatedVesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatedVesting = append(m.DelegatedVesting, types1.Coin{})
			if err := m.DelegatedVesting[len(m.DelegatedVesting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *Period) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: Period: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types1.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionType", wireType)
			}
			m.ActionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActionType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *StridePeriodicVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: StridePeriodicVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StridePeriodicVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriods = append(m.VestingPeriods, Period{})
			if err := m.VestingPeriods[len(m.VestingPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func skipVesting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
				return 0, ErrInvalidLengthVesting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVesting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVesting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVesting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVesting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVesting = fmt.Errorf("proto: unexpected end of group")
)
