// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: universitychainit/personal_info.proto

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

type PersonalInfo struct {
	Gender             string `protobuf:"bytes,1,opt,name=gender,proto3" json:"gender,omitempty"`
	DateOfBirth        string `protobuf:"bytes,2,opt,name=dateOfBirth,proto3" json:"dateOfBirth,omitempty"`
	PrimaryNationality string `protobuf:"bytes,3,opt,name=primaryNationality,proto3" json:"primaryNationality,omitempty"`
	CountryOfBirth     string `protobuf:"bytes,4,opt,name=countryOfBirth,proto3" json:"countryOfBirth,omitempty"`
	ProvinceOfBirth    string `protobuf:"bytes,5,opt,name=provinceOfBirth,proto3" json:"provinceOfBirth,omitempty"`
	TownOfBirth        string `protobuf:"bytes,6,opt,name=townOfBirth,proto3" json:"townOfBirth,omitempty"`
	TaxCode            string `protobuf:"bytes,7,opt,name=taxCode,proto3" json:"taxCode,omitempty"`
}

func (m *PersonalInfo) Reset()         { *m = PersonalInfo{} }
func (m *PersonalInfo) String() string { return proto.CompactTextString(m) }
func (*PersonalInfo) ProtoMessage()    {}
func (*PersonalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c87983360b14e0, []int{0}
}
func (m *PersonalInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PersonalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PersonalInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PersonalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonalInfo.Merge(m, src)
}
func (m *PersonalInfo) XXX_Size() int {
	return m.Size()
}
func (m *PersonalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PersonalInfo proto.InternalMessageInfo

func (m *PersonalInfo) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *PersonalInfo) GetDateOfBirth() string {
	if m != nil {
		return m.DateOfBirth
	}
	return ""
}

func (m *PersonalInfo) GetPrimaryNationality() string {
	if m != nil {
		return m.PrimaryNationality
	}
	return ""
}

func (m *PersonalInfo) GetCountryOfBirth() string {
	if m != nil {
		return m.CountryOfBirth
	}
	return ""
}

func (m *PersonalInfo) GetProvinceOfBirth() string {
	if m != nil {
		return m.ProvinceOfBirth
	}
	return ""
}

func (m *PersonalInfo) GetTownOfBirth() string {
	if m != nil {
		return m.TownOfBirth
	}
	return ""
}

func (m *PersonalInfo) GetTaxCode() string {
	if m != nil {
		return m.TaxCode
	}
	return ""
}

func init() {
	proto.RegisterType((*PersonalInfo)(nil), "university_chain_it.universitychainit.PersonalInfo")
}

func init() {
	proto.RegisterFile("universitychainit/personal_info.proto", fileDescriptor_b2c87983360b14e0)
}

var fileDescriptor_b2c87983360b14e0 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe3, 0x02, 0xa9, 0x30, 0x08, 0x24, 0x0f, 0xc8, 0x93, 0x55, 0x21, 0x81, 0xba, 0x90,
	0x0c, 0xbc, 0x41, 0x19, 0x10, 0x0b, 0x20, 0x46, 0x96, 0xc8, 0x24, 0x0e, 0x3d, 0x09, 0x6c, 0xcb,
	0xb9, 0x96, 0xfa, 0x05, 0x98, 0x79, 0x2c, 0xc6, 0x8e, 0x8c, 0x28, 0x79, 0x11, 0x54, 0xd3, 0x54,
	0x51, 0x9b, 0xf1, 0xbe, 0xff, 0xbb, 0xf3, 0xe9, 0x4c, 0x2f, 0x66, 0x1a, 0xe6, 0xca, 0x55, 0x80,
	0x3e, 0x9f, 0x4a, 0xd0, 0x80, 0xa9, 0x55, 0xae, 0x32, 0x5a, 0xbe, 0x65, 0xa0, 0x4b, 0x93, 0x58,
	0x67, 0xd0, 0xb0, 0x8e, 0x96, 0x05, 0x2f, 0x03, 0x4c, 0x76, 0x5a, 0xcf, 0x3f, 0x07, 0xf4, 0xf8,
	0x71, 0xdd, 0x7e, 0xa7, 0x4b, 0xc3, 0xce, 0x68, 0xfc, 0xaa, 0x74, 0xa1, 0x1c, 0x27, 0x23, 0x32,
	0x3e, 0x7c, 0x5a, 0x57, 0x6c, 0x44, 0x8f, 0x0a, 0x89, 0xea, 0xa1, 0x9c, 0x80, 0xc3, 0x29, 0x1f,
	0x84, 0xb0, 0x8b, 0x58, 0x42, 0x99, 0x75, 0xf0, 0x2e, 0x9d, 0xbf, 0x97, 0x08, 0xab, 0x81, 0x80,
	0x9e, 0xef, 0x05, 0xb1, 0x27, 0x61, 0x97, 0xf4, 0x24, 0x37, 0x33, 0x8d, 0xce, 0xb7, 0x43, 0xf7,
	0x83, 0xbb, 0x45, 0xd9, 0x98, 0x9e, 0x5a, 0x67, 0xe6, 0xa0, 0xf3, 0xcd, 0xeb, 0x07, 0x41, 0xdc,
	0xc6, 0xab, 0x1d, 0xd1, 0x7c, 0xe8, 0xd6, 0x8a, 0xff, 0x77, 0xec, 0x20, 0xc6, 0xe9, 0x10, 0xe5,
	0xe2, 0xc6, 0x14, 0x8a, 0x0f, 0x43, 0xda, 0x96, 0x93, 0xdb, 0xef, 0x5a, 0x90, 0x65, 0x2d, 0xc8,
	0x6f, 0x2d, 0xc8, 0x57, 0x23, 0xa2, 0x65, 0x23, 0xa2, 0x9f, 0x46, 0x44, 0xcf, 0x57, 0x3d, 0x97,
	0x4c, 0x17, 0xe9, 0xee, 0x37, 0xa0, 0xb7, 0xaa, 0x7a, 0x89, 0xc3, 0xfd, 0xaf, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x29, 0x78, 0xf3, 0xd9, 0xa8, 0x01, 0x00, 0x00,
}

func (m *PersonalInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PersonalInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PersonalInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TaxCode) > 0 {
		i -= len(m.TaxCode)
		copy(dAtA[i:], m.TaxCode)
		i = encodeVarintPersonalInfo(dAtA, i, uint64(len(m.TaxCode)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TownOfBirth) > 0 {
		i -= len(m.TownOfBirth)
		copy(dAtA[i:], m.TownOfBirth)
		i = encodeVarintPersonalInfo(dAtA, i, uint64(len(m.TownOfBirth)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ProvinceOfBirth) > 0 {
		i -= len(m.ProvinceOfBirth)
		copy(dAtA[i:], m.ProvinceOfBirth)
		i = encodeVarintPersonalInfo(dAtA, i, uint64(len(m.ProvinceOfBirth)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CountryOfBirth) > 0 {
		i -= len(m.CountryOfBirth)
		copy(dAtA[i:], m.CountryOfBirth)
		i = encodeVarintPersonalInfo(dAtA, i, uint64(len(m.CountryOfBirth)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PrimaryNationality) > 0 {
		i -= len(m.PrimaryNationality)
		copy(dAtA[i:], m.PrimaryNationality)
		i = encodeVarintPersonalInfo(dAtA, i, uint64(len(m.PrimaryNationality)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DateOfBirth) > 0 {
		i -= len(m.DateOfBirth)
		copy(dAtA[i:], m.DateOfBirth)
		i = encodeVarintPersonalInfo(dAtA, i, uint64(len(m.DateOfBirth)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Gender) > 0 {
		i -= len(m.Gender)
		copy(dAtA[i:], m.Gender)
		i = encodeVarintPersonalInfo(dAtA, i, uint64(len(m.Gender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPersonalInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovPersonalInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PersonalInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Gender)
	if l > 0 {
		n += 1 + l + sovPersonalInfo(uint64(l))
	}
	l = len(m.DateOfBirth)
	if l > 0 {
		n += 1 + l + sovPersonalInfo(uint64(l))
	}
	l = len(m.PrimaryNationality)
	if l > 0 {
		n += 1 + l + sovPersonalInfo(uint64(l))
	}
	l = len(m.CountryOfBirth)
	if l > 0 {
		n += 1 + l + sovPersonalInfo(uint64(l))
	}
	l = len(m.ProvinceOfBirth)
	if l > 0 {
		n += 1 + l + sovPersonalInfo(uint64(l))
	}
	l = len(m.TownOfBirth)
	if l > 0 {
		n += 1 + l + sovPersonalInfo(uint64(l))
	}
	l = len(m.TaxCode)
	if l > 0 {
		n += 1 + l + sovPersonalInfo(uint64(l))
	}
	return n
}

func sovPersonalInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPersonalInfo(x uint64) (n int) {
	return sovPersonalInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PersonalInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPersonalInfo
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
			return fmt.Errorf("proto: PersonalInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PersonalInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonalInfo
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
				return ErrInvalidLengthPersonalInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPersonalInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DateOfBirth", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonalInfo
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
				return ErrInvalidLengthPersonalInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPersonalInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DateOfBirth = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimaryNationality", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonalInfo
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
				return ErrInvalidLengthPersonalInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPersonalInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrimaryNationality = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountryOfBirth", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonalInfo
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
				return ErrInvalidLengthPersonalInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPersonalInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CountryOfBirth = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvinceOfBirth", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonalInfo
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
				return ErrInvalidLengthPersonalInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPersonalInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProvinceOfBirth = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TownOfBirth", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonalInfo
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
				return ErrInvalidLengthPersonalInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPersonalInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TownOfBirth = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonalInfo
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
				return ErrInvalidLengthPersonalInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPersonalInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaxCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPersonalInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPersonalInfo
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
func skipPersonalInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPersonalInfo
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
					return 0, ErrIntOverflowPersonalInfo
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
					return 0, ErrIntOverflowPersonalInfo
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
				return 0, ErrInvalidLengthPersonalInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPersonalInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPersonalInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPersonalInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPersonalInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPersonalInfo = fmt.Errorf("proto: unexpected end of group")
)