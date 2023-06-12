// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/claim/claim.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Action int32

const (
	ACTION_FREE           Action = 0
	ACTION_LIQUID_STAKE   Action = 1
	ACTION_DELEGATE_STAKE Action = 2
)

var Action_name = map[int32]string{
	0: "ACTION_FREE",
	1: "ACTION_LIQUID_STAKE",
	2: "ACTION_DELEGATE_STAKE",
}

var Action_value = map[string]int32{
	"ACTION_FREE":           0,
	"ACTION_LIQUID_STAKE":   1,
	"ACTION_DELEGATE_STAKE": 2,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4747d999b9dc0da, []int{0}
}

// A Claim Records is the metadata of claim data per address
type ClaimRecord struct {
	// airdrop identifier
	AirdropIdentifier string `protobuf:"bytes,1,opt,name=airdrop_identifier,json=airdropIdentifier,proto3" json:"airdrop_identifier,omitempty" yaml:"airdrop_identifier"`
	// address of claim user
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// weight that represent the portion from total allocation
	Weight github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight" yaml:"weight"`
	// true if action is completed
	// index of bool in array refers to action enum #
	ActionCompleted []bool `protobuf:"varint,4,rep,packed,name=action_completed,json=actionCompleted,proto3" json:"action_completed,omitempty" yaml:"action_completed"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4747d999b9dc0da, []int{0}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetAirdropIdentifier() string {
	if m != nil {
		return m.AirdropIdentifier
	}
	return ""
}

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetActionCompleted() []bool {
	if m != nil {
		return m.ActionCompleted
	}
	return nil
}

func init() {
	proto.RegisterEnum("stride.claim.Action", Action_name, Action_value)
	proto.RegisterType((*ClaimRecord)(nil), "stride.claim.ClaimRecord")
}

func init() { proto.RegisterFile("stride/claim/claim.proto", fileDescriptor_b4747d999b9dc0da) }

var fileDescriptor_b4747d999b9dc0da = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0xce, 0xd2, 0x40,
	0x00, 0xc5, 0x5b, 0x20, 0xa8, 0x83, 0x0a, 0x8e, 0x1a, 0x0a, 0xc6, 0x96, 0x74, 0x61, 0x88, 0x91,
	0x76, 0xe1, 0x4a, 0x37, 0xa6, 0x40, 0xd1, 0xc6, 0x46, 0x63, 0xc1, 0x98, 0xb8, 0x69, 0x4a, 0x67,
	0x2c, 0x13, 0x29, 0xd3, 0x74, 0xc6, 0x3f, 0xdc, 0xc0, 0xa5, 0x77, 0x70, 0xe7, 0x49, 0x58, 0xb2,
	0x34, 0x2e, 0x1a, 0x03, 0x37, 0xe8, 0x09, 0x0c, 0xd3, 0xc1, 0x98, 0xef, 0xdb, 0xb4, 0xcd, 0xef,
	0xf7, 0xf2, 0xd2, 0x99, 0x07, 0x34, 0xc6, 0x73, 0x82, 0xb0, 0x1d, 0xaf, 0x23, 0x92, 0x56, 0x4f,
	0x2b, 0xcb, 0x29, 0xa7, 0xf0, 0x7a, 0x65, 0x2c, 0xc1, 0xfa, 0x77, 0x12, 0x9a, 0x50, 0x21, 0xec,
	0xd3, 0x57, 0x95, 0x31, 0x7f, 0xd6, 0x40, 0x6b, 0x72, 0xf2, 0x01, 0x8e, 0x69, 0x8e, 0xa0, 0x0f,
	0x60, 0x44, 0x72, 0x94, 0xd3, 0x2c, 0x24, 0x08, 0x6f, 0x38, 0xf9, 0x40, 0x70, 0xae, 0xa9, 0x03,
	0x75, 0x78, 0x6d, 0x7c, 0xbf, 0x2c, 0x8c, 0xde, 0x36, 0x4a, 0xd7, 0x4f, 0xcd, 0xcb, 0x19, 0x33,
	0xb8, 0x25, 0xa1, 0xf7, 0x8f, 0xc1, 0x47, 0xe0, 0x4a, 0x84, 0x50, 0x8e, 0x19, 0xd3, 0x6a, 0xa2,
	0x02, 0x96, 0x85, 0x71, 0x53, 0x56, 0x54, 0xc2, 0x0c, 0xce, 0x11, 0xf8, 0x0e, 0x34, 0xbf, 0x60,
	0x92, 0xac, 0xb8, 0x56, 0x17, 0xe1, 0x67, 0xbb, 0xc2, 0x50, 0x7e, 0x17, 0xc6, 0x83, 0x84, 0xf0,
	0xd5, 0xa7, 0xa5, 0x15, 0xd3, 0xd4, 0x8e, 0x29, 0x4b, 0x29, 0x93, 0xaf, 0x11, 0x43, 0x1f, 0x6d,
	0xbe, 0xcd, 0x30, 0xb3, 0xa6, 0x38, 0x2e, 0x0b, 0xe3, 0x46, 0x55, 0x5d, 0xb5, 0x98, 0x81, 0xac,
	0x83, 0x33, 0xd0, 0x89, 0x62, 0x4e, 0xe8, 0x26, 0x8c, 0x69, 0x9a, 0xad, 0x31, 0xc7, 0x48, 0x6b,
	0x0c, 0xea, 0xc3, 0xab, 0xe3, 0x7b, 0x65, 0x61, 0x74, 0xe5, 0xff, 0x5c, 0x48, 0x98, 0x41, 0xbb,
	0x42, 0x93, 0x33, 0x79, 0x38, 0x07, 0x4d, 0x47, 0x20, 0xd8, 0x06, 0x2d, 0x67, 0xb2, 0xf0, 0x5e,
	0xbf, 0x0a, 0x67, 0x81, 0xeb, 0x76, 0x14, 0xd8, 0x05, 0xb7, 0x25, 0xf0, 0xbd, 0x37, 0x6f, 0xbd,
	0x69, 0x38, 0x5f, 0x38, 0x2f, 0xdd, 0x8e, 0x0a, 0x7b, 0xe0, 0xae, 0x14, 0x53, 0xd7, 0x77, 0x9f,
	0x3b, 0x0b, 0x57, 0xaa, 0x5a, 0xbf, 0xf1, 0xed, 0x87, 0xae, 0x8c, 0x5f, 0xec, 0x0e, 0xba, 0xba,
	0x3f, 0xe8, 0xea, 0x9f, 0x83, 0xae, 0x7e, 0x3f, 0xea, 0xca, 0xfe, 0xa8, 0x2b, 0xbf, 0x8e, 0xba,
	0xf2, 0xde, 0xfa, 0xef, 0xdc, 0x73, 0x31, 0xe5, 0xc8, 0x8f, 0x96, 0xcc, 0x96, 0x83, 0x7f, 0x7e,
	0x62, 0x7f, 0x95, 0xab, 0x8b, 0x3b, 0x58, 0x36, 0xc5, 0xa4, 0x8f, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xae, 0x58, 0x48, 0x17, 0x12, 0x02, 0x00, 0x00,
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionCompleted) > 0 {
		for iNdEx := len(m.ActionCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaim(dAtA, i, uint64(len(m.ActionCompleted)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaim(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AirdropIdentifier) > 0 {
		i -= len(m.AirdropIdentifier)
		copy(dAtA[i:], m.AirdropIdentifier)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.AirdropIdentifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AirdropIdentifier)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovClaim(uint64(l))
	if len(m.ActionCompleted) > 0 {
		n += 1 + sovClaim(uint64(len(m.ActionCompleted))) + len(m.ActionCompleted)*1
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AirdropIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
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
				m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionCompleted) == 0 {
					m.ActionCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
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
					m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)
