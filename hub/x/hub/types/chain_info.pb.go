// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hub/chain_info.proto

package types

import (
	fmt "fmt"
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

type ChainInfo struct {
	ChainKey              string `protobuf:"bytes,1,opt,name=chainKey,proto3" json:"chainKey,omitempty"`
	ChainAdministratorKey string `protobuf:"bytes,2,opt,name=chainAdministratorKey,proto3" json:"chainAdministratorKey,omitempty"`
	InitStatus            bool   `protobuf:"varint,3,opt,name=initStatus,proto3" json:"initStatus,omitempty"`
}

func (m *ChainInfo) Reset()         { *m = ChainInfo{} }
func (m *ChainInfo) String() string { return proto.CompactTextString(m) }
func (*ChainInfo) ProtoMessage()    {}
func (*ChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6a228684842ccbf, []int{0}
}
func (m *ChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainInfo.Merge(m, src)
}
func (m *ChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *ChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChainInfo proto.InternalMessageInfo

func (m *ChainInfo) GetChainKey() string {
	if m != nil {
		return m.ChainKey
	}
	return ""
}

func (m *ChainInfo) GetChainAdministratorKey() string {
	if m != nil {
		return m.ChainAdministratorKey
	}
	return ""
}

func (m *ChainInfo) GetInitStatus() bool {
	if m != nil {
		return m.InitStatus
	}
	return false
}

func init() {
	proto.RegisterType((*ChainInfo)(nil), "hub.hub.ChainInfo")
}

func init() { proto.RegisterFile("hub/chain_info.proto", fileDescriptor_f6a228684842ccbf) }

var fileDescriptor_f6a228684842ccbf = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x28, 0x4d, 0xd2,
	0x4f, 0xce, 0x48, 0xcc, 0xcc, 0x8b, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0xcf, 0x28, 0x4d, 0xd2, 0xcb, 0x28, 0x4d, 0x52, 0xaa, 0xe5, 0xe2, 0x74, 0x06, 0x49,
	0x7a, 0xe6, 0xa5, 0xe5, 0x0b, 0x49, 0x71, 0x71, 0x80, 0x55, 0x7a, 0xa7, 0x56, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x42, 0x26, 0x5c, 0xa2, 0x60, 0xb6, 0x63, 0x4a, 0x6e, 0x66,
	0x5e, 0x66, 0x71, 0x49, 0x51, 0x62, 0x49, 0x7e, 0x11, 0x48, 0x21, 0x13, 0x58, 0x21, 0x76, 0x49,
	0x21, 0x39, 0x2e, 0xae, 0xcc, 0xbc, 0xcc, 0x92, 0xe0, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0x09, 0x66,
	0x05, 0x46, 0x0d, 0x8e, 0x20, 0x24, 0x11, 0x27, 0xcd, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c,
	0x96, 0x63, 0x88, 0xe2, 0x07, 0xb9, 0xbb, 0x42, 0x1f, 0x44, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27,
	0xb1, 0x81, 0x5d, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x86, 0x17, 0x62, 0xd1, 0x00,
	0x00, 0x00,
}

func (m *ChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InitStatus {
		i--
		if m.InitStatus {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainAdministratorKey) > 0 {
		i -= len(m.ChainAdministratorKey)
		copy(dAtA[i:], m.ChainAdministratorKey)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.ChainAdministratorKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainKey) > 0 {
		i -= len(m.ChainKey)
		copy(dAtA[i:], m.ChainKey)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.ChainKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintChainInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovChainInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainKey)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	l = len(m.ChainAdministratorKey)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	if m.InitStatus {
		n += 2
	}
	return n
}

func sovChainInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChainInfo(x uint64) (n int) {
	return sovChainInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChainInfo
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
			return fmt.Errorf("proto: ChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainAdministratorKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainAdministratorKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitStatus", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
			m.InitStatus = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipChainInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChainInfo
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
func skipChainInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChainInfo
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
					return 0, ErrIntOverflowChainInfo
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
					return 0, ErrIntOverflowChainInfo
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
				return 0, ErrInvalidLengthChainInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChainInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChainInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChainInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChainInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChainInfo = fmt.Errorf("proto: unexpected end of group")
)