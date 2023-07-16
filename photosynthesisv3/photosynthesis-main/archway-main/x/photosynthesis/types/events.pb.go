// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archway/photosynthesis/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type LiquidStakeDepositRecordCreatedEvent struct {
	RecordId      string `protobuf:"bytes,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	RewardsAmount int64  `protobuf:"varint,2,opt,name=rewards_amount,json=rewardsAmount,proto3" json:"rewards_amount,omitempty"`
}

func (m *LiquidStakeDepositRecordCreatedEvent) Reset()         { *m = LiquidStakeDepositRecordCreatedEvent{} }
func (m *LiquidStakeDepositRecordCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*LiquidStakeDepositRecordCreatedEvent) ProtoMessage()    {}
func (*LiquidStakeDepositRecordCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bdf0e4f89738036, []int{0}
}
func (m *LiquidStakeDepositRecordCreatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidStakeDepositRecordCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidStakeDepositRecordCreatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidStakeDepositRecordCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidStakeDepositRecordCreatedEvent.Merge(m, src)
}
func (m *LiquidStakeDepositRecordCreatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *LiquidStakeDepositRecordCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidStakeDepositRecordCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidStakeDepositRecordCreatedEvent proto.InternalMessageInfo

func (m *LiquidStakeDepositRecordCreatedEvent) GetRecordId() string {
	if m != nil {
		return m.RecordId
	}
	return ""
}

func (m *LiquidStakeDepositRecordCreatedEvent) GetRewardsAmount() int64 {
	if m != nil {
		return m.RewardsAmount
	}
	return 0
}

type RedemptionRateUpdatedEvent struct {
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	NewThreshold    int64  `protobuf:"varint,2,opt,name=new_threshold,json=newThreshold,proto3" json:"new_threshold,omitempty"`
}

func (m *RedemptionRateUpdatedEvent) Reset()         { *m = RedemptionRateUpdatedEvent{} }
func (m *RedemptionRateUpdatedEvent) String() string { return proto.CompactTextString(m) }
func (*RedemptionRateUpdatedEvent) ProtoMessage()    {}
func (*RedemptionRateUpdatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bdf0e4f89738036, []int{1}
}
func (m *RedemptionRateUpdatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RedemptionRateUpdatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RedemptionRateUpdatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RedemptionRateUpdatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedemptionRateUpdatedEvent.Merge(m, src)
}
func (m *RedemptionRateUpdatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *RedemptionRateUpdatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RedemptionRateUpdatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RedemptionRateUpdatedEvent proto.InternalMessageInfo

func (m *RedemptionRateUpdatedEvent) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *RedemptionRateUpdatedEvent) GetNewThreshold() int64 {
	if m != nil {
		return m.NewThreshold
	}
	return 0
}

type RewardsDistributedEvent struct {
	RewardAddress string `protobuf:"bytes,1,opt,name=reward_address,json=rewardAddress,proto3" json:"reward_address,omitempty"`
	RewardsAmount int64  `protobuf:"varint,2,opt,name=rewards_amount,json=rewardsAmount,proto3" json:"rewards_amount,omitempty"`
	NumContracts  int32  `protobuf:"varint,3,opt,name=num_contracts,json=numContracts,proto3" json:"num_contracts,omitempty"`
}

func (m *RewardsDistributedEvent) Reset()         { *m = RewardsDistributedEvent{} }
func (m *RewardsDistributedEvent) String() string { return proto.CompactTextString(m) }
func (*RewardsDistributedEvent) ProtoMessage()    {}
func (*RewardsDistributedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bdf0e4f89738036, []int{2}
}
func (m *RewardsDistributedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardsDistributedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardsDistributedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardsDistributedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardsDistributedEvent.Merge(m, src)
}
func (m *RewardsDistributedEvent) XXX_Size() int {
	return m.Size()
}
func (m *RewardsDistributedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardsDistributedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RewardsDistributedEvent proto.InternalMessageInfo

func (m *RewardsDistributedEvent) GetRewardAddress() string {
	if m != nil {
		return m.RewardAddress
	}
	return ""
}

func (m *RewardsDistributedEvent) GetRewardsAmount() int64 {
	if m != nil {
		return m.RewardsAmount
	}
	return 0
}

func (m *RewardsDistributedEvent) GetNumContracts() int32 {
	if m != nil {
		return m.NumContracts
	}
	return 0
}

type RewardsWithdrawEvent struct {
	RewardAddress string  `protobuf:"bytes,1,opt,name=reward_address,json=rewardAddress,proto3" json:"reward_address,omitempty"`
	Rewards       []*Coin `protobuf:"bytes,2,rep,name=rewards,proto3" json:"rewards,omitempty"`
}

func (m *RewardsWithdrawEvent) Reset()         { *m = RewardsWithdrawEvent{} }
func (m *RewardsWithdrawEvent) String() string { return proto.CompactTextString(m) }
func (*RewardsWithdrawEvent) ProtoMessage()    {}
func (*RewardsWithdrawEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bdf0e4f89738036, []int{3}
}
func (m *RewardsWithdrawEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardsWithdrawEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardsWithdrawEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardsWithdrawEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardsWithdrawEvent.Merge(m, src)
}
func (m *RewardsWithdrawEvent) XXX_Size() int {
	return m.Size()
}
func (m *RewardsWithdrawEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardsWithdrawEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RewardsWithdrawEvent proto.InternalMessageInfo

func (m *RewardsWithdrawEvent) GetRewardAddress() string {
	if m != nil {
		return m.RewardAddress
	}
	return ""
}

func (m *RewardsWithdrawEvent) GetRewards() []*Coin {
	if m != nil {
		return m.Rewards
	}
	return nil
}

type Event struct {
	Type       string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Time       string `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Attributes string `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bdf0e4f89738036, []int{4}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Event) GetAttributes() string {
	if m != nil {
		return m.Attributes
	}
	return ""
}

type Events struct {
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *Events) Reset()         { *m = Events{} }
func (m *Events) String() string { return proto.CompactTextString(m) }
func (*Events) ProtoMessage()    {}
func (*Events) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bdf0e4f89738036, []int{5}
}
func (m *Events) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Events) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Events.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Events) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Events.Merge(m, src)
}
func (m *Events) XXX_Size() int {
	return m.Size()
}
func (m *Events) XXX_DiscardUnknown() {
	xxx_messageInfo_Events.DiscardUnknown(m)
}

var xxx_messageInfo_Events proto.InternalMessageInfo

func (m *Events) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*LiquidStakeDepositRecordCreatedEvent)(nil), "LiquidStakeDepositRecordCreatedEvent")
	proto.RegisterType((*RedemptionRateUpdatedEvent)(nil), "RedemptionRateUpdatedEvent")
	proto.RegisterType((*RewardsDistributedEvent)(nil), "RewardsDistributedEvent")
	proto.RegisterType((*RewardsWithdrawEvent)(nil), "RewardsWithdrawEvent")
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*Events)(nil), "Events")
}

func init() {
	proto.RegisterFile("archway/photosynthesis/events.proto", fileDescriptor_3bdf0e4f89738036)
}

var fileDescriptor_3bdf0e4f89738036 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0x80, 0x9b, 0xae, 0xbb, 0xba, 0x63, 0xab, 0x12, 0x04, 0x43, 0x85, 0x74, 0x49, 0x15, 0x22,
	0xe2, 0x86, 0xea, 0xc9, 0x63, 0xdd, 0x7a, 0x10, 0x04, 0x61, 0xaa, 0x08, 0x1e, 0x0c, 0x93, 0xcc,
	0x73, 0x33, 0x76, 0x67, 0x26, 0xce, 0xbc, 0x34, 0xee, 0x7f, 0xf0, 0xe0, 0xcf, 0xf2, 0xd8, 0xa3,
	0x47, 0xd9, 0xfd, 0x23, 0xb2, 0x93, 0x49, 0xd5, 0x82, 0xa0, 0xb7, 0x9d, 0x6f, 0x86, 0xf7, 0x7d,
	0x1b, 0x1e, 0x39, 0x60, 0xa6, 0xac, 0x5a, 0xb6, 0xcc, 0xea, 0x4a, 0xa3, 0xb6, 0x4b, 0x85, 0x15,
	0x58, 0x61, 0x33, 0x38, 0x03, 0x85, 0x76, 0x5a, 0x1b, 0x8d, 0x7a, 0x6f, 0x7f, 0xae, 0xf5, 0x7c,
	0x01, 0x99, 0x3b, 0x15, 0xcd, 0x87, 0x0c, 0x85, 0x04, 0x8b, 0x4c, 0xd6, 0xfe, 0x41, 0x5c, 0x6a,
	0x2b, 0xb5, 0xcd, 0x0a, 0x66, 0x21, 0x3b, 0x3b, 0x2c, 0x00, 0xd9, 0x61, 0x56, 0x6a, 0xa1, 0xfc,
	0xfd, 0xc3, 0xbf, 0x58, 0xfe, 0x3c, 0x76, 0x8f, 0x93, 0x8f, 0xe4, 0xde, 0x4b, 0xf1, 0xa9, 0x11,
	0xfc, 0x04, 0xd9, 0x29, 0x1c, 0x43, 0xad, 0xad, 0x40, 0x0a, 0xa5, 0x36, 0x7c, 0x66, 0x80, 0x21,
	0xf0, 0xe7, 0x9b, 0xb8, 0xf0, 0x2e, 0x19, 0x1b, 0x47, 0x73, 0xc1, 0xa3, 0x60, 0x12, 0xa4, 0x63,
	0x7a, 0xad, 0x03, 0x2f, 0x78, 0x78, 0x9f, 0xdc, 0x30, 0xd0, 0x32, 0xc3, 0x6d, 0xce, 0xa4, 0x6e,
	0x14, 0x46, 0xdb, 0x93, 0x20, 0x1d, 0xd0, 0x5d, 0x4f, 0x8f, 0x1c, 0x4c, 0x16, 0x64, 0x8f, 0x02,
	0x07, 0x59, 0xa3, 0xd0, 0x8a, 0x32, 0x84, 0x37, 0x35, 0xff, 0x65, 0x78, 0x40, 0x6e, 0x95, 0x5a,
	0xa1, 0x61, 0x25, 0xe6, 0x8c, 0x73, 0x03, 0xd6, 0x7a, 0xd1, 0xcd, 0x9e, 0x1f, 0x75, 0x38, 0x3c,
	0x20, 0xbb, 0x0a, 0xda, 0x1c, 0x2b, 0x03, 0xb6, 0xd2, 0x0b, 0xee, 0x75, 0x3b, 0x0a, 0xda, 0xd7,
	0x3d, 0x4b, 0xbe, 0x04, 0xe4, 0x0e, 0xed, 0xfc, 0xc7, 0xc2, 0xa2, 0x11, 0x45, 0x73, 0xe1, 0xba,
	0x08, 0xbe, 0x64, 0xf2, 0xc1, 0xbd, 0xe7, 0xdf, 0xfe, 0x97, 0xcb, 0x69, 0x64, 0xde, 0x57, 0xda,
	0x68, 0x30, 0x09, 0xd2, 0x21, 0xdd, 0x51, 0x8d, 0x9c, 0xf5, 0x2c, 0x79, 0x4f, 0x6e, 0xfb, 0x9a,
	0xb7, 0x02, 0x2b, 0x6e, 0x58, 0xfb, 0x5f, 0x29, 0xfb, 0xe4, 0xaa, 0x97, 0x46, 0xdb, 0x93, 0x41,
	0x7a, 0xfd, 0xf1, 0x70, 0x3a, 0xd3, 0x42, 0xd1, 0x9e, 0x26, 0xaf, 0xc8, 0xb0, 0x1b, 0x18, 0x92,
	0x2b, 0xb8, 0xac, 0xc1, 0x8f, 0x71, 0xbf, 0x1d, 0x13, 0x12, 0x5c, 0xfe, 0x86, 0x09, 0x09, 0x61,
	0x4c, 0x08, 0x43, 0xff, 0x59, 0xba, 0xe4, 0x31, 0xfd, 0x8d, 0x24, 0x29, 0x19, 0xb9, 0x81, 0x36,
	0x8c, 0xc9, 0xa8, 0xdb, 0xd0, 0x28, 0x70, 0xea, 0xd1, 0xd4, 0x5d, 0x50, 0x4f, 0x9f, 0x9d, 0x7c,
	0x5b, 0xc5, 0xc1, 0xf9, 0x2a, 0x0e, 0x7e, 0xac, 0xe2, 0xe0, 0xeb, 0x3a, 0xde, 0x3a, 0x5f, 0xc7,
	0x5b, 0xdf, 0xd7, 0xf1, 0xd6, 0xbb, 0xa7, 0x73, 0x81, 0x55, 0x53, 0x4c, 0x4b, 0x2d, 0x33, 0xbf,
	0x95, 0x8f, 0x14, 0x60, 0xab, 0xcd, 0x69, 0x7f, 0xce, 0x3e, 0x5f, 0xde, 0xd3, 0x4d, 0xb1, 0x2d,
	0x46, 0x6e, 0x3f, 0x9f, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xc1, 0xc7, 0x05, 0x34, 0x03,
	0x00, 0x00,
}

func (m *LiquidStakeDepositRecordCreatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidStakeDepositRecordCreatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidStakeDepositRecordCreatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RewardsAmount != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.RewardsAmount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RecordId) > 0 {
		i -= len(m.RecordId)
		copy(dAtA[i:], m.RecordId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RecordId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RedemptionRateUpdatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RedemptionRateUpdatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RedemptionRateUpdatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewThreshold != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.NewThreshold))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RewardsDistributedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardsDistributedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardsDistributedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumContracts != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.NumContracts))
		i--
		dAtA[i] = 0x18
	}
	if m.RewardsAmount != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.RewardsAmount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RewardAddress) > 0 {
		i -= len(m.RewardAddress)
		copy(dAtA[i:], m.RewardAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RewardAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RewardsWithdrawEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardsWithdrawEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardsWithdrawEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RewardAddress) > 0 {
		i -= len(m.RewardAddress)
		copy(dAtA[i:], m.RewardAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RewardAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		i -= len(m.Attributes)
		copy(dAtA[i:], m.Attributes)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Attributes)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Time) > 0 {
		i -= len(m.Time)
		copy(dAtA[i:], m.Time)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Time)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Events) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Events) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Events) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LiquidStakeDepositRecordCreatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RecordId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.RewardsAmount != 0 {
		n += 1 + sovEvents(uint64(m.RewardsAmount))
	}
	return n
}

func (m *RedemptionRateUpdatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.NewThreshold != 0 {
		n += 1 + sovEvents(uint64(m.NewThreshold))
	}
	return n
}

func (m *RewardsDistributedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RewardAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.RewardsAmount != 0 {
		n += 1 + sovEvents(uint64(m.RewardsAmount))
	}
	if m.NumContracts != 0 {
		n += 1 + sovEvents(uint64(m.NumContracts))
	}
	return n
}

func (m *RewardsWithdrawEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RewardAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Attributes)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *Events) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LiquidStakeDepositRecordCreatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: LiquidStakeDepositRecordCreatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidStakeDepositRecordCreatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecordId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsAmount", wireType)
			}
			m.RewardsAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardsAmount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *RedemptionRateUpdatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: RedemptionRateUpdatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RedemptionRateUpdatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewThreshold", wireType)
			}
			m.NewThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewThreshold |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *RewardsDistributedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: RewardsDistributedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardsDistributedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsAmount", wireType)
			}
			m.RewardsAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardsAmount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumContracts", wireType)
			}
			m.NumContracts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumContracts |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *RewardsWithdrawEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: RewardsWithdrawEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardsWithdrawEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, &Coin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *Events) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: Events: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Events: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
