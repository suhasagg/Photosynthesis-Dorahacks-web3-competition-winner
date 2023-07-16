// DONTCOVER
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archway/rewards/v1beta1/events.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// ContractMetadataSetEvent is emitted when the contract metadata is created or
// updated.
type ContractMetadataSetEvent struct {
	// contract_address defines the contract address.
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// metadata defines the new contract metadata state.
	Metadata ContractMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata"`
}

func (m *ContractMetadataSetEvent) Reset()         { *m = ContractMetadataSetEvent{} }
func (m *ContractMetadataSetEvent) String() string { return proto.CompactTextString(m) }
func (*ContractMetadataSetEvent) ProtoMessage()    {}
func (*ContractMetadataSetEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad2689b4f7dc3cd8, []int{0}
}
func (m *ContractMetadataSetEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractMetadataSetEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractMetadataSetEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractMetadataSetEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractMetadataSetEvent.Merge(m, src)
}
func (m *ContractMetadataSetEvent) XXX_Size() int {
	return m.Size()
}
func (m *ContractMetadataSetEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractMetadataSetEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ContractMetadataSetEvent proto.InternalMessageInfo

func (m *ContractMetadataSetEvent) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractMetadataSetEvent) GetMetadata() ContractMetadata {
	if m != nil {
		return m.Metadata
	}
	return ContractMetadata{}
}

// ContractRewardCalculationEvent is emitted when the contract reward is
// calculated.
type ContractRewardCalculationEvent struct {
	// contract_address defines the contract address.
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// gas_consumed defines the total gas consumption by all WASM operations
	// within one transaction.
	GasConsumed uint64 `protobuf:"varint,2,opt,name=gas_consumed,json=gasConsumed,proto3" json:"gas_consumed,omitempty"`
	// inflation_rewards defines the inflation rewards portions of the rewards.
	InflationRewards types.Coin `protobuf:"bytes,3,opt,name=inflation_rewards,json=inflationRewards,proto3" json:"inflation_rewards"`
	// fee_rebate_rewards defines the fee rebate rewards portions of the rewards.
	FeeRebateRewards []types.Coin `protobuf:"bytes,4,rep,name=fee_rebate_rewards,json=feeRebateRewards,proto3" json:"fee_rebate_rewards"`
	// metadata defines the contract metadata (if set).
	Metadata *ContractMetadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *ContractRewardCalculationEvent) Reset()         { *m = ContractRewardCalculationEvent{} }
func (m *ContractRewardCalculationEvent) String() string { return proto.CompactTextString(m) }
func (*ContractRewardCalculationEvent) ProtoMessage()    {}
func (*ContractRewardCalculationEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad2689b4f7dc3cd8, []int{1}
}
func (m *ContractRewardCalculationEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractRewardCalculationEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractRewardCalculationEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractRewardCalculationEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractRewardCalculationEvent.Merge(m, src)
}
func (m *ContractRewardCalculationEvent) XXX_Size() int {
	return m.Size()
}
func (m *ContractRewardCalculationEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractRewardCalculationEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ContractRewardCalculationEvent proto.InternalMessageInfo

func (m *ContractRewardCalculationEvent) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractRewardCalculationEvent) GetGasConsumed() uint64 {
	if m != nil {
		return m.GasConsumed
	}
	return 0
}

func (m *ContractRewardCalculationEvent) GetInflationRewards() types.Coin {
	if m != nil {
		return m.InflationRewards
	}
	return types.Coin{}
}

func (m *ContractRewardCalculationEvent) GetFeeRebateRewards() []types.Coin {
	if m != nil {
		return m.FeeRebateRewards
	}
	return nil
}

func (m *ContractRewardCalculationEvent) GetMetadata() *ContractMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// RewardsWithdrawEvent is emitted when credited rewards for a specific
// rewards_address are distributed. Event could be triggered by a transaction
// (via CLI for example) or by a contract via WASM bindings.
type RewardsWithdrawEvent struct {
	// rewards_address defines the rewards address rewards are distributed to.
	RewardAddress string `protobuf:"bytes,1,opt,name=reward_address,json=rewardAddress,proto3" json:"reward_address,omitempty"`
	// rewards defines the total rewards being distributed.
	Rewards []types.Coin `protobuf:"bytes,2,rep,name=rewards,proto3" json:"rewards"`
}

func (m *RewardsWithdrawEvent) Reset()         { *m = RewardsWithdrawEvent{} }
func (m *RewardsWithdrawEvent) String() string { return proto.CompactTextString(m) }
func (*RewardsWithdrawEvent) ProtoMessage()    {}
func (*RewardsWithdrawEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad2689b4f7dc3cd8, []int{2}
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

func (m *RewardsWithdrawEvent) GetRewards() []types.Coin {
	if m != nil {
		return m.Rewards
	}
	return nil
}

// MinConsensusFeeSetEvent is emitted when the minimum consensus fee is updated.
type MinConsensusFeeSetEvent struct {
	// fee defines the updated minimum gas unit price.
	Fee types.DecCoin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`
}

func (m *MinConsensusFeeSetEvent) Reset()         { *m = MinConsensusFeeSetEvent{} }
func (m *MinConsensusFeeSetEvent) String() string { return proto.CompactTextString(m) }
func (*MinConsensusFeeSetEvent) ProtoMessage()    {}
func (*MinConsensusFeeSetEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad2689b4f7dc3cd8, []int{3}
}
func (m *MinConsensusFeeSetEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MinConsensusFeeSetEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MinConsensusFeeSetEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MinConsensusFeeSetEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinConsensusFeeSetEvent.Merge(m, src)
}
func (m *MinConsensusFeeSetEvent) XXX_Size() int {
	return m.Size()
}
func (m *MinConsensusFeeSetEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MinConsensusFeeSetEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MinConsensusFeeSetEvent proto.InternalMessageInfo

func (m *MinConsensusFeeSetEvent) GetFee() types.DecCoin {
	if m != nil {
		return m.Fee
	}
	return types.DecCoin{}
}

// ContractFlatFeeSetEvent is emitted when the contract flat fee is updated
type ContractFlatFeeSetEvent struct {
	// contract_address defines the bech32 address of the contract for which the
	// flat fee is set
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// flat_fee defines the amount that has been set as the minimum fee for the
	// contract
	FlatFee types.Coin `protobuf:"bytes,2,opt,name=flat_fee,json=flatFee,proto3" json:"flat_fee"`
}

func (m *ContractFlatFeeSetEvent) Reset()         { *m = ContractFlatFeeSetEvent{} }
func (m *ContractFlatFeeSetEvent) String() string { return proto.CompactTextString(m) }
func (*ContractFlatFeeSetEvent) ProtoMessage()    {}
func (*ContractFlatFeeSetEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad2689b4f7dc3cd8, []int{4}
}
func (m *ContractFlatFeeSetEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractFlatFeeSetEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractFlatFeeSetEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractFlatFeeSetEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractFlatFeeSetEvent.Merge(m, src)
}
func (m *ContractFlatFeeSetEvent) XXX_Size() int {
	return m.Size()
}
func (m *ContractFlatFeeSetEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractFlatFeeSetEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ContractFlatFeeSetEvent proto.InternalMessageInfo

func (m *ContractFlatFeeSetEvent) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractFlatFeeSetEvent) GetFlatFee() types.Coin {
	if m != nil {
		return m.FlatFee
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*ContractMetadataSetEvent)(nil), "archway.rewards.v1beta1.ContractMetadataSetEvent")
	proto.RegisterType((*ContractRewardCalculationEvent)(nil), "archway.rewards.v1beta1.ContractRewardCalculationEvent")
	proto.RegisterType((*RewardsWithdrawEvent)(nil), "archway.rewards.v1beta1.RewardsWithdrawEvent")
	proto.RegisterType((*MinConsensusFeeSetEvent)(nil), "archway.rewards.v1beta1.MinConsensusFeeSetEvent")
	proto.RegisterType((*ContractFlatFeeSetEvent)(nil), "archway.rewards.v1beta1.ContractFlatFeeSetEvent")
}

func init() {
	proto.RegisterFile("archway/rewards/v1beta1/events.proto", fileDescriptor_ad2689b4f7dc3cd8)
}

var fileDescriptor_ad2689b4f7dc3cd8 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x8d, 0x93, 0xfc, 0x7e, 0x2d, 0x1b, 0xfe, 0x14, 0xab, 0x52, 0x42, 0x85, 0x4c, 0x88, 0xa8,
	0xd4, 0x1e, 0xb0, 0xd5, 0xc0, 0x05, 0x6e, 0x34, 0xb4, 0x17, 0x1a, 0x21, 0x99, 0x03, 0x12, 0x17,
	0x6b, 0xbc, 0x1e, 0x3b, 0x16, 0xc9, 0x6e, 0xb5, 0xbb, 0x69, 0xda, 0x1b, 0x1f, 0x01, 0xf1, 0xa9,
	0x7a, 0xac, 0x38, 0x71, 0x42, 0x28, 0xf9, 0x22, 0xc8, 0xde, 0x5d, 0xcb, 0xaa, 0xa8, 0x94, 0xdc,
	0xec, 0xd9, 0x37, 0x6f, 0xde, 0x7b, 0xb3, 0x4b, 0x5e, 0x80, 0xa0, 0x93, 0x05, 0x5c, 0x05, 0x02,
	0x17, 0x20, 0x12, 0x19, 0x5c, 0x1c, 0xc5, 0xa8, 0xe0, 0x28, 0xc0, 0x0b, 0x64, 0x4a, 0xfa, 0xe7,
	0x82, 0x2b, 0xee, 0x76, 0x0d, 0xca, 0x37, 0x28, 0xdf, 0xa0, 0xf6, 0x76, 0x33, 0x9e, 0xf1, 0x12,
	0x13, 0x14, 0x5f, 0x1a, 0xbe, 0xe7, 0x51, 0x2e, 0x67, 0x5c, 0x06, 0x31, 0x48, 0xac, 0x08, 0x29,
	0xcf, 0x99, 0x39, 0xdf, 0xbf, 0x6b, 0xa8, 0xa5, 0x2f, 0x61, 0x83, 0x1f, 0x0e, 0xe9, 0x8d, 0x38,
	0x53, 0x02, 0xa8, 0x1a, 0xa3, 0x82, 0x04, 0x14, 0x7c, 0x42, 0x75, 0x52, 0x28, 0x73, 0x0f, 0xc9,
	0x0e, 0x35, 0x67, 0x11, 0x24, 0x89, 0x40, 0x29, 0x7b, 0x4e, 0xdf, 0x39, 0xb8, 0x17, 0x3e, 0xb2,
	0xf5, 0x77, 0xba, 0xec, 0x7e, 0x20, 0xdb, 0x33, 0xd3, 0xde, 0x6b, 0xf6, 0x9d, 0x83, 0xce, 0xf0,
	0xd0, 0xbf, 0xc3, 0x90, 0x7f, 0x7b, 0xde, 0x71, 0xfb, 0xfa, 0xf7, 0xb3, 0x46, 0x58, 0x11, 0x0c,
	0x7e, 0x36, 0x89, 0x67, 0x41, 0x61, 0xd9, 0x3c, 0x82, 0x29, 0x9d, 0x4f, 0x41, 0xe5, 0x9c, 0x6d,
	0x2c, 0xed, 0x39, 0xb9, 0x9f, 0x81, 0x8c, 0x28, 0x67, 0x72, 0x3e, 0xc3, 0xa4, 0x94, 0xd7, 0x0e,
	0x3b, 0x19, 0xc8, 0x91, 0x29, 0xb9, 0x67, 0xe4, 0x71, 0xce, 0x52, 0xcd, 0x1f, 0x19, 0xb9, 0xbd,
	0x56, 0x69, 0xe3, 0x89, 0xaf, 0x83, 0xf6, 0x8b, 0xa0, 0x6b, 0x16, 0x72, 0x66, 0x64, 0xef, 0x54,
	0x9d, 0x5a, 0xaa, 0x74, 0xc7, 0xc4, 0x4d, 0x11, 0x23, 0x81, 0x31, 0x28, 0xac, 0xe8, 0xda, 0xfd,
	0xd6, 0x5a, 0x74, 0x29, 0x62, 0x58, 0x76, 0x5a, 0xba, 0x93, 0x5a, 0xb4, 0xff, 0x6d, 0x18, 0x6d,
	0x2d, 0xd4, 0x4b, 0xb2, 0x6b, 0x18, 0x3f, 0xe7, 0x6a, 0x92, 0x08, 0x58, 0xe8, 0x24, 0xf7, 0xc9,
	0x43, 0xcd, 0x72, 0x2b, 0xc7, 0x07, 0xba, 0x6a, 0x53, 0x7c, 0x43, 0xb6, 0xac, 0x93, 0xe6, 0x7a,
	0x4e, 0x2c, 0x7e, 0xf0, 0x91, 0x74, 0xc7, 0x39, 0x2b, 0xc2, 0x46, 0x26, 0xe7, 0xf2, 0x14, 0xb1,
	0xba, 0x61, 0xaf, 0x49, 0x2b, 0x45, 0x2c, 0x27, 0x76, 0x86, 0x4f, 0xff, 0xc9, 0xf8, 0x1e, 0x69,
	0x8d, 0xb4, 0x80, 0x0f, 0xbe, 0x39, 0xa4, 0x6b, 0x9d, 0x9e, 0x4e, 0x41, 0xd5, 0x19, 0x37, 0xb8,
	0x18, 0x6f, 0xc9, 0x76, 0xb1, 0xb9, 0xa8, 0x50, 0xd0, 0x5c, 0x6f, 0xd9, 0x5b, 0xa9, 0x1e, 0x77,
	0x7c, 0x76, 0xbd, 0xf4, 0x9c, 0x9b, 0xa5, 0xe7, 0xfc, 0x59, 0x7a, 0xce, 0xf7, 0x95, 0xd7, 0xb8,
	0x59, 0x79, 0x8d, 0x5f, 0x2b, 0xaf, 0xf1, 0x65, 0x98, 0xe5, 0x6a, 0x32, 0x8f, 0x7d, 0xca, 0x67,
	0x81, 0x59, 0xd3, 0x4b, 0x86, 0x6a, 0xc1, 0xc5, 0x57, 0xfb, 0x1f, 0x5c, 0x56, 0xaf, 0x52, 0x5d,
	0x9d, 0xa3, 0x8c, 0xff, 0x2f, 0x1f, 0xe3, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xde, 0x3c,
	0x46, 0x8c, 0x2a, 0x04, 0x00, 0x00,
}

func (m *ContractMetadataSetEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractMetadataSetEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractMetadataSetEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContractRewardCalculationEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractRewardCalculationEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractRewardCalculationEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.FeeRebateRewards) > 0 {
		for iNdEx := len(m.FeeRebateRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeRebateRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.InflationRewards.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.GasConsumed != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.GasConsumed))
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

func (m *MinConsensusFeeSetEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MinConsensusFeeSetEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MinConsensusFeeSetEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ContractFlatFeeSetEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractFlatFeeSetEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractFlatFeeSetEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.FlatFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
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
func (m *ContractMetadataSetEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *ContractRewardCalculationEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.GasConsumed != 0 {
		n += 1 + sovEvents(uint64(m.GasConsumed))
	}
	l = m.InflationRewards.Size()
	n += 1 + l + sovEvents(uint64(l))
	if len(m.FeeRebateRewards) > 0 {
		for _, e := range m.FeeRebateRewards {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovEvents(uint64(l))
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

func (m *MinConsensusFeeSetEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Fee.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *ContractFlatFeeSetEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.FlatFee.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractMetadataSetEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContractMetadataSetEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractMetadataSetEvent: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
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
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ContractRewardCalculationEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContractRewardCalculationEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractRewardCalculationEvent: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field GasConsumed", wireType)
			}
			m.GasConsumed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasConsumed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRewards", wireType)
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
			if err := m.InflationRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRebateRewards", wireType)
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
			m.FeeRebateRewards = append(m.FeeRebateRewards, types.Coin{})
			if err := m.FeeRebateRewards[len(m.FeeRebateRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
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
			if m.Metadata == nil {
				m.Metadata = &ContractMetadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
			m.Rewards = append(m.Rewards, types.Coin{})
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
func (m *MinConsensusFeeSetEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MinConsensusFeeSetEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MinConsensusFeeSetEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
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
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ContractFlatFeeSetEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContractFlatFeeSetEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractFlatFeeSetEvent: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlatFee", wireType)
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
			if err := m.FlatFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
