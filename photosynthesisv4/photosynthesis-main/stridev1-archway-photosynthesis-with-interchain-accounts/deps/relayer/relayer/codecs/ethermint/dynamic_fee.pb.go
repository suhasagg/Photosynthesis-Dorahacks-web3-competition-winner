// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/types/v1/dynamic_fee.proto

package ethermint

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
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

// ExtensionOptionDynamicFeeTx is an extension option that specifies the
// maxPrioPrice for cosmos tx
type ExtensionOptionDynamicFeeTx struct {
	// max_priority_price is the same as `max_priority_fee_per_gas` in eip-1559
	// spec
	MaxPriorityPrice github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=max_priority_price,json=maxPriorityPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_priority_price"`
}

func (m *ExtensionOptionDynamicFeeTx) Reset()         { *m = ExtensionOptionDynamicFeeTx{} }
func (m *ExtensionOptionDynamicFeeTx) String() string { return proto.CompactTextString(m) }
func (*ExtensionOptionDynamicFeeTx) ProtoMessage()    {}
func (*ExtensionOptionDynamicFeeTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d7cf05c9992c480, []int{0}
}
func (m *ExtensionOptionDynamicFeeTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtensionOptionDynamicFeeTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtensionOptionDynamicFeeTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtensionOptionDynamicFeeTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionOptionDynamicFeeTx.Merge(m, src)
}
func (m *ExtensionOptionDynamicFeeTx) XXX_Size() int {
	return m.Size()
}
func (m *ExtensionOptionDynamicFeeTx) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionOptionDynamicFeeTx.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionOptionDynamicFeeTx proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExtensionOptionDynamicFeeTx)(nil), "ethermint.types.v1.ExtensionOptionDynamicFeeTx")
}

func init() {
	proto.RegisterFile("ethermint/types/v1/dynamic_fee.proto", fileDescriptor_9d7cf05c9992c480)
}

var fileDescriptor_9d7cf05c9992c480 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0x93, 0xcb, 0x07, 0x5f, 0x4e, 0x12, 0x3c, 0x88, 0xc2, 0x56, 0x44, 0xc4, 0x4b, 0x77,
	0x29, 0x5e, 0x3d, 0x15, 0x15, 0x3c, 0x59, 0x44, 0x3c, 0x88, 0x50, 0xd2, 0xcd, 0x98, 0x2e, 0x66,
	0x67, 0xc2, 0xee, 0x34, 0x24, 0xf8, 0x27, 0xfc, 0x59, 0x3d, 0xf6, 0x28, 0x1e, 0x8a, 0x24, 0x7f,
	0x44, 0x92, 0x94, 0xe2, 0x69, 0x5e, 0x98, 0x87, 0x87, 0xf7, 0x8d, 0xce, 0x81, 0x97, 0xe0, 0xac,
	0x41, 0x56, 0x5c, 0x17, 0xe0, 0x55, 0x39, 0x51, 0x69, 0x8d, 0x89, 0x35, 0x7a, 0xfe, 0x06, 0x20,
	0x0b, 0x47, 0x4c, 0x71, 0xbc, 0xa7, 0x64, 0x4f, 0xc9, 0x72, 0x72, 0x7c, 0x98, 0x51, 0x46, 0xfd,
	0x5b, 0x75, 0x69, 0x20, 0xcf, 0x3e, 0xa2, 0x93, 0xdb, 0x8a, 0x01, 0xbd, 0x21, 0x7c, 0x28, 0xd8,
	0x10, 0xde, 0x0c, 0xb6, 0x3b, 0x80, 0xa7, 0x2a, 0x7e, 0x8d, 0x62, 0x9b, 0x54, 0xf3, 0xc2, 0x19,
	0x72, 0x86, 0xeb, 0x2e, 0x68, 0x38, 0x0a, 0x4f, 0xc3, 0xcb, 0xff, 0x53, 0xb9, 0xde, 0x8e, 0x82,
	0xef, 0xed, 0xe8, 0x22, 0x33, 0xbc, 0x5c, 0x2d, 0xa4, 0x26, 0xab, 0x34, 0x79, 0x4b, 0x7e, 0x77,
	0xc6, 0x3e, 0x7d, 0x1f, 0x5a, 0xca, 0x7b, 0xe4, 0xc7, 0x03, 0x9b, 0x54, 0xb3, 0x9d, 0x68, 0xd6,
	0x79, 0xa6, 0xcf, 0xeb, 0x46, 0x84, 0x9b, 0x46, 0x84, 0x3f, 0x8d, 0x08, 0x3f, 0x5b, 0x11, 0x6c,
	0x5a, 0x11, 0x7c, 0xb5, 0x22, 0x78, 0xb9, 0xfe, 0xe3, 0xf4, 0xec, 0x12, 0xcc, 0x20, 0xa7, 0x12,
	0xc6, 0x25, 0x20, 0xaf, 0x1c, 0x78, 0x95, 0x03, 0x7a, 0xa5, 0x73, 0x03, 0xc8, 0x4a, 0x53, 0x0a,
	0xda, 0xab, 0xfd, 0xe6, 0xc5, 0xbf, 0x7e, 0xdb, 0xd5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e,
	0x31, 0x52, 0x1f, 0x2d, 0x01, 0x00, 0x00,
}

func (m *ExtensionOptionDynamicFeeTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtensionOptionDynamicFeeTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtensionOptionDynamicFeeTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxPriorityPrice.Size()
		i -= size
		if _, err := m.MaxPriorityPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDynamicFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDynamicFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovDynamicFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtensionOptionDynamicFeeTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MaxPriorityPrice.Size()
	n += 1 + l + sovDynamicFee(uint64(l))
	return n
}

func sovDynamicFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDynamicFee(x uint64) (n int) {
	return sovDynamicFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtensionOptionDynamicFeeTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicFee
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
			return fmt.Errorf("proto: ExtensionOptionDynamicFeeTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtensionOptionDynamicFeeTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPriorityPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicFee
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
				return ErrInvalidLengthDynamicFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxPriorityPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDynamicFee
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
func skipDynamicFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDynamicFee
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
					return 0, ErrIntOverflowDynamicFee
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
					return 0, ErrIntOverflowDynamicFee
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
				return 0, ErrInvalidLengthDynamicFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDynamicFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDynamicFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDynamicFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDynamicFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDynamicFee = fmt.Errorf("proto: unexpected end of group")
)
