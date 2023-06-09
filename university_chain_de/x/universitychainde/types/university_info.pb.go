// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: universitychainde/university_info.proto

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

type UniversityInfo struct {
	UniversityName  string `protobuf:"bytes,1,opt,name=universityName,proto3" json:"universityName,omitempty"`
	NextStudentId   uint32 `protobuf:"varint,2,opt,name=nextStudentId,proto3" json:"nextStudentId,omitempty"`
	SecretariatKey  string `protobuf:"bytes,3,opt,name=secretariatKey,proto3" json:"secretariatKey,omitempty"`
	UniversityKey   string `protobuf:"bytes,4,opt,name=universityKey,proto3" json:"universityKey,omitempty"`
	CaiKey          string `protobuf:"bytes,5,opt,name=caiKey,proto3" json:"caiKey,omitempty"`
	FifoHeadErasmus string `protobuf:"bytes,6,opt,name=fifoHeadErasmus,proto3" json:"fifoHeadErasmus,omitempty"`
	FifoTailErasmus string `protobuf:"bytes,7,opt,name=fifoTailErasmus,proto3" json:"fifoTailErasmus,omitempty"`
	DeadlineTaxes   string `protobuf:"bytes,8,opt,name=deadlineTaxes,proto3" json:"deadlineTaxes,omitempty"`
	DeadlineErasmus string `protobuf:"bytes,9,opt,name=deadlineErasmus,proto3" json:"deadlineErasmus,omitempty"`
	TaxesBrackets   string `protobuf:"bytes,10,opt,name=taxesBrackets,proto3" json:"taxesBrackets,omitempty"`
	MaxErasmusExams int32  `protobuf:"varint,11,opt,name=maxErasmusExams,proto3" json:"maxErasmusExams,omitempty"`
}

func (m *UniversityInfo) Reset()         { *m = UniversityInfo{} }
func (m *UniversityInfo) String() string { return proto.CompactTextString(m) }
func (*UniversityInfo) ProtoMessage()    {}
func (*UniversityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6f78a94827245c6, []int{0}
}
func (m *UniversityInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UniversityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UniversityInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UniversityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniversityInfo.Merge(m, src)
}
func (m *UniversityInfo) XXX_Size() int {
	return m.Size()
}
func (m *UniversityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UniversityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UniversityInfo proto.InternalMessageInfo

func (m *UniversityInfo) GetUniversityName() string {
	if m != nil {
		return m.UniversityName
	}
	return ""
}

func (m *UniversityInfo) GetNextStudentId() uint32 {
	if m != nil {
		return m.NextStudentId
	}
	return 0
}

func (m *UniversityInfo) GetSecretariatKey() string {
	if m != nil {
		return m.SecretariatKey
	}
	return ""
}

func (m *UniversityInfo) GetUniversityKey() string {
	if m != nil {
		return m.UniversityKey
	}
	return ""
}

func (m *UniversityInfo) GetCaiKey() string {
	if m != nil {
		return m.CaiKey
	}
	return ""
}

func (m *UniversityInfo) GetFifoHeadErasmus() string {
	if m != nil {
		return m.FifoHeadErasmus
	}
	return ""
}

func (m *UniversityInfo) GetFifoTailErasmus() string {
	if m != nil {
		return m.FifoTailErasmus
	}
	return ""
}

func (m *UniversityInfo) GetDeadlineTaxes() string {
	if m != nil {
		return m.DeadlineTaxes
	}
	return ""
}

func (m *UniversityInfo) GetDeadlineErasmus() string {
	if m != nil {
		return m.DeadlineErasmus
	}
	return ""
}

func (m *UniversityInfo) GetTaxesBrackets() string {
	if m != nil {
		return m.TaxesBrackets
	}
	return ""
}

func (m *UniversityInfo) GetMaxErasmusExams() int32 {
	if m != nil {
		return m.MaxErasmusExams
	}
	return 0
}

func init() {
	proto.RegisterType((*UniversityInfo)(nil), "university_chain_de.universitychainde.UniversityInfo")
}

func init() {
	proto.RegisterFile("universitychainde/university_info.proto", fileDescriptor_e6f78a94827245c6)
}

var fileDescriptor_e6f78a94827245c6 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0x40, 0x99, 0x8f, 0x0f, 0x94, 0x31, 0x60, 0xd2, 0x85, 0x99, 0xd5, 0x84, 0x18, 0x7f, 0xba,
	0x11, 0x16, 0xbe, 0x01, 0x09, 0x51, 0x62, 0xe2, 0x02, 0x71, 0xe3, 0x86, 0x5c, 0x3b, 0xb7, 0x71,
	0x22, 0x9d, 0x92, 0x99, 0xc1, 0x94, 0xb7, 0xf0, 0x71, 0x7c, 0x04, 0x97, 0x2c, 0x5d, 0x1a, 0x78,
	0x11, 0x33, 0x2d, 0xb5, 0xb4, 0xb8, 0xec, 0x99, 0x73, 0x4f, 0x6e, 0xd2, 0x4b, 0x2f, 0x17, 0x4a,
	0xbe, 0xa1, 0x36, 0xd2, 0x2e, 0x83, 0x17, 0x90, 0x4a, 0x60, 0xbf, 0x20, 0x53, 0xa9, 0xc2, 0xb8,
	0x37, 0xd7, 0xb1, 0x8d, 0xbd, 0xf3, 0x1d, 0x9c, 0x9a, 0x53, 0x81, 0xbd, 0xbd, 0xe1, 0xd3, 0x8f,
	0x3a, 0xed, 0x3c, 0xfe, 0xd2, 0x91, 0x0a, 0x63, 0xef, 0x82, 0x76, 0x0a, 0xef, 0x1e, 0x22, 0x64,
	0xa4, 0x4b, 0xfc, 0xd6, 0xb8, 0x42, 0xbd, 0x33, 0xda, 0x56, 0x98, 0xd8, 0x07, 0xbb, 0x10, 0xa8,
	0xec, 0x48, 0xb0, 0x7f, 0x5d, 0xe2, 0xb7, 0xc7, 0x65, 0xe8, 0x6a, 0x06, 0x03, 0x8d, 0x16, 0xb4,
	0x04, 0x7b, 0x87, 0x4b, 0x56, 0xcf, 0x6a, 0x65, 0xea, 0x6a, 0x45, 0xdf, 0x69, 0xff, 0x53, 0xad,
	0x0c, 0xbd, 0x13, 0xda, 0x0c, 0x40, 0xba, 0xe7, 0x46, 0xfa, 0xbc, 0xfd, 0xf2, 0x7c, 0x7a, 0x1c,
	0xca, 0x30, 0xbe, 0x45, 0x10, 0x43, 0x0d, 0x26, 0x5a, 0x18, 0xd6, 0x4c, 0x85, 0x2a, 0xce, 0xcd,
	0x09, 0xc8, 0x59, 0x6e, 0x1e, 0x14, 0xe6, 0x0e, 0x76, 0x1b, 0x09, 0x04, 0x31, 0x93, 0x0a, 0x27,
	0x90, 0xa0, 0x61, 0x87, 0xd9, 0x46, 0x25, 0xe8, 0x7a, 0x39, 0xc8, 0x7b, 0xad, 0xac, 0x57, 0xc1,
	0xae, 0x67, 0xdd, 0xc8, 0x40, 0x43, 0xf0, 0x8a, 0xd6, 0x30, 0x9a, 0xf5, 0x4a, 0xd0, 0xf5, 0x22,
	0x48, 0xb6, 0x33, 0xc3, 0x04, 0x22, 0xc3, 0x8e, 0xba, 0xc4, 0x6f, 0x8c, 0xab, 0x78, 0x70, 0xf3,
	0xb9, 0xe6, 0x64, 0xb5, 0xe6, 0xe4, 0x7b, 0xcd, 0xc9, 0xfb, 0x86, 0xd7, 0x56, 0x1b, 0x5e, 0xfb,
	0xda, 0xf0, 0xda, 0xd3, 0xd5, 0x1f, 0xff, 0xbe, 0x9f, 0xf4, 0xf7, 0x4f, 0xc7, 0x2e, 0xe7, 0x68,
	0x9e, 0x9b, 0xe9, 0xc5, 0x5c, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x01, 0xa2, 0x5c, 0x31, 0x5c,
	0x02, 0x00, 0x00,
}

func (m *UniversityInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UniversityInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UniversityInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxErasmusExams != 0 {
		i = encodeVarintUniversityInfo(dAtA, i, uint64(m.MaxErasmusExams))
		i--
		dAtA[i] = 0x58
	}
	if len(m.TaxesBrackets) > 0 {
		i -= len(m.TaxesBrackets)
		copy(dAtA[i:], m.TaxesBrackets)
		i = encodeVarintUniversityInfo(dAtA, i, uint64(len(m.TaxesBrackets)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.DeadlineErasmus) > 0 {
		i -= len(m.DeadlineErasmus)
		copy(dAtA[i:], m.DeadlineErasmus)
		i = encodeVarintUniversityInfo(dAtA, i, uint64(len(m.DeadlineErasmus)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.DeadlineTaxes) > 0 {
		i -= len(m.DeadlineTaxes)
		copy(dAtA[i:], m.DeadlineTaxes)
		i = encodeVarintUniversityInfo(dAtA, i, uint64(len(m.DeadlineTaxes)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.FifoTailErasmus) > 0 {
		i -= len(m.FifoTailErasmus)
		copy(dAtA[i:], m.FifoTailErasmus)
		i = encodeVarintUniversityInfo(dAtA, i, uint64(len(m.FifoTailErasmus)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.FifoHeadErasmus) > 0 {
		i -= len(m.FifoHeadErasmus)
		copy(dAtA[i:], m.FifoHeadErasmus)
		i = encodeVarintUniversityInfo(dAtA, i, uint64(len(m.FifoHeadErasmus)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.CaiKey) > 0 {
		i -= len(m.CaiKey)
		copy(dAtA[i:], m.CaiKey)
		i = encodeVarintUniversityInfo(dAtA, i, uint64(len(m.CaiKey)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.UniversityKey) > 0 {
		i -= len(m.UniversityKey)
		copy(dAtA[i:], m.UniversityKey)
		i = encodeVarintUniversityInfo(dAtA, i, uint64(len(m.UniversityKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SecretariatKey) > 0 {
		i -= len(m.SecretariatKey)
		copy(dAtA[i:], m.SecretariatKey)
		i = encodeVarintUniversityInfo(dAtA, i, uint64(len(m.SecretariatKey)))
		i--
		dAtA[i] = 0x1a
	}
	if m.NextStudentId != 0 {
		i = encodeVarintUniversityInfo(dAtA, i, uint64(m.NextStudentId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UniversityName) > 0 {
		i -= len(m.UniversityName)
		copy(dAtA[i:], m.UniversityName)
		i = encodeVarintUniversityInfo(dAtA, i, uint64(len(m.UniversityName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUniversityInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovUniversityInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UniversityInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UniversityName)
	if l > 0 {
		n += 1 + l + sovUniversityInfo(uint64(l))
	}
	if m.NextStudentId != 0 {
		n += 1 + sovUniversityInfo(uint64(m.NextStudentId))
	}
	l = len(m.SecretariatKey)
	if l > 0 {
		n += 1 + l + sovUniversityInfo(uint64(l))
	}
	l = len(m.UniversityKey)
	if l > 0 {
		n += 1 + l + sovUniversityInfo(uint64(l))
	}
	l = len(m.CaiKey)
	if l > 0 {
		n += 1 + l + sovUniversityInfo(uint64(l))
	}
	l = len(m.FifoHeadErasmus)
	if l > 0 {
		n += 1 + l + sovUniversityInfo(uint64(l))
	}
	l = len(m.FifoTailErasmus)
	if l > 0 {
		n += 1 + l + sovUniversityInfo(uint64(l))
	}
	l = len(m.DeadlineTaxes)
	if l > 0 {
		n += 1 + l + sovUniversityInfo(uint64(l))
	}
	l = len(m.DeadlineErasmus)
	if l > 0 {
		n += 1 + l + sovUniversityInfo(uint64(l))
	}
	l = len(m.TaxesBrackets)
	if l > 0 {
		n += 1 + l + sovUniversityInfo(uint64(l))
	}
	if m.MaxErasmusExams != 0 {
		n += 1 + sovUniversityInfo(uint64(m.MaxErasmusExams))
	}
	return n
}

func sovUniversityInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUniversityInfo(x uint64) (n int) {
	return sovUniversityInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UniversityInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUniversityInfo
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
			return fmt.Errorf("proto: UniversityInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UniversityInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniversityName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniversityInfo
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
				return ErrInvalidLengthUniversityInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUniversityInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UniversityName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextStudentId", wireType)
			}
			m.NextStudentId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniversityInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextStudentId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecretariatKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniversityInfo
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
				return ErrInvalidLengthUniversityInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUniversityInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecretariatKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniversityKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniversityInfo
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
				return ErrInvalidLengthUniversityInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUniversityInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UniversityKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CaiKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniversityInfo
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
				return ErrInvalidLengthUniversityInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUniversityInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CaiKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FifoHeadErasmus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniversityInfo
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
				return ErrInvalidLengthUniversityInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUniversityInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FifoHeadErasmus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FifoTailErasmus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniversityInfo
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
				return ErrInvalidLengthUniversityInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUniversityInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FifoTailErasmus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeadlineTaxes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniversityInfo
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
				return ErrInvalidLengthUniversityInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUniversityInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeadlineTaxes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeadlineErasmus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniversityInfo
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
				return ErrInvalidLengthUniversityInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUniversityInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeadlineErasmus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxesBrackets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniversityInfo
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
				return ErrInvalidLengthUniversityInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUniversityInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaxesBrackets = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxErasmusExams", wireType)
			}
			m.MaxErasmusExams = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniversityInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxErasmusExams |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUniversityInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUniversityInfo
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
func skipUniversityInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUniversityInfo
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
					return 0, ErrIntOverflowUniversityInfo
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
					return 0, ErrIntOverflowUniversityInfo
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
				return 0, ErrInvalidLengthUniversityInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUniversityInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUniversityInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUniversityInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUniversityInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUniversityInfo = fmt.Errorf("proto: unexpected end of group")
)
