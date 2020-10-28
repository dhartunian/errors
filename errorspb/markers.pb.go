// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/cockroachdb/errors/errorspb/markers.proto

package errorspb

import (
	fmt "fmt"
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

// MarkPayload is the error payload for a forced marker.
// See errors/markers/markers.go and the RFC on
// error handling for details.
type MarkPayload struct {
	Msg                  string          `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Types                []ErrorTypeMark `protobuf:"bytes,2,rep,name=types,proto3" json:"types"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MarkPayload) Reset()         { *m = MarkPayload{} }
func (m *MarkPayload) String() string { return proto.CompactTextString(m) }
func (*MarkPayload) ProtoMessage()    {}
func (*MarkPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ac3bc232982be93, []int{0}
}
func (m *MarkPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarkPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MarkPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkPayload.Merge(m, src)
}
func (m *MarkPayload) XXX_Size() int {
	return m.Size()
}
func (m *MarkPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkPayload.DiscardUnknown(m)
}

var xxx_messageInfo_MarkPayload proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MarkPayload)(nil), "cockroach.errorspb.MarkPayload")
}

func init() {
	proto.RegisterFile("github.com/cockroachdb/errors/errorspb/markers.proto", fileDescriptor_5ac3bc232982be93)
}

var fileDescriptor_5ac3bc232982be93 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x4f, 0xce, 0x2e, 0xca, 0x4f, 0x4c, 0xce,
	0x48, 0x49, 0xd2, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0x2a, 0x86, 0x52, 0x05, 0x49, 0xfa, 0xb9, 0x89,
	0x45, 0xd9, 0xa9, 0x45, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x42, 0x70, 0xa5, 0x7a,
	0x30, 0x15, 0x52, 0xc6, 0x44, 0x9a, 0x04, 0x61, 0x40, 0x0c, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0x33, 0xf5, 0x41, 0x2c, 0x88, 0xa8, 0x52, 0x1c, 0x17, 0xb7, 0x6f, 0x62, 0x51, 0x76, 0x40,
	0x62, 0x65, 0x4e, 0x7e, 0x62, 0x8a, 0x90, 0x00, 0x17, 0x73, 0x6e, 0x71, 0xba, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x88, 0x29, 0x64, 0xcb, 0xc5, 0x5a, 0x52, 0x59, 0x90, 0x5a, 0x2c, 0xc1,
	0xa4, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa8, 0x87, 0xe9, 0x1e, 0x3d, 0x57, 0x10, 0x23, 0xa4, 0xb2,
	0x20, 0x15, 0x64, 0x94, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x10, 0x5d, 0x4e, 0x4a, 0x27,
	0x1e, 0xca, 0x31, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x8d, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x10, 0xc5, 0x01, 0xd3, 0x9e, 0xc4, 0x06, 0x76,
	0x8a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x84, 0xb7, 0x15, 0x21, 0x01, 0x00, 0x00,
}

func (m *MarkPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarkPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarkPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Types) > 0 {
		for iNdEx := len(m.Types) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Types[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMarkers(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintMarkers(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMarkers(dAtA []byte, offset int, v uint64) int {
	offset -= sovMarkers(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MarkPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovMarkers(uint64(l))
	}
	if len(m.Types) > 0 {
		for _, e := range m.Types {
			l = e.Size()
			n += 1 + l + sovMarkers(uint64(l))
		}
	}
	return n
}

func sovMarkers(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMarkers(x uint64) (n int) {
	return sovMarkers(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MarkPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarkers
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
			return fmt.Errorf("proto: MarkPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarkPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarkers
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
				return ErrInvalidLengthMarkers
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarkers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Types", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarkers
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
				return ErrInvalidLengthMarkers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarkers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Types = append(m.Types, ErrorTypeMark{})
			if err := m.Types[len(m.Types)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarkers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMarkers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMarkers
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
func skipMarkers(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarkers
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
					return 0, ErrIntOverflowMarkers
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
					return 0, ErrIntOverflowMarkers
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
				return 0, ErrInvalidLengthMarkers
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMarkers
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMarkers
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMarkers        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarkers          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMarkers = fmt.Errorf("proto: unexpected end of group")
)
