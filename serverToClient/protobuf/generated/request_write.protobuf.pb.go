// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: request_write.protobuf

package s2cProtobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RequestWrite struct {
	Path string `protobuf:"bytes,1,opt,name=Path,json=path,proto3" json:"Path,omitempty"`
}

func (m *RequestWrite) Reset()                    { *m = RequestWrite{} }
func (m *RequestWrite) String() string            { return proto.CompactTextString(m) }
func (*RequestWrite) ProtoMessage()               {}
func (*RequestWrite) Descriptor() ([]byte, []int) { return fileDescriptorRequestWritebuf, []int{0} }

func (m *RequestWrite) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestWrite)(nil), "s2cProtobuf.RequestWrite")
}
func (m *RequestWrite) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestWrite) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Path) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRequestWritebuf(dAtA, i, uint64(len(m.Path)))
		i += copy(dAtA[i:], m.Path)
	}
	return i, nil
}

func encodeVarintRequestWritebuf(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RequestWrite) Size() (n int) {
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovRequestWritebuf(uint64(l))
	}
	return n
}

func sovRequestWritebuf(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRequestWritebuf(x uint64) (n int) {
	return sovRequestWritebuf(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestWrite) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequestWritebuf
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RequestWrite: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestWrite: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestWritebuf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestWritebuf
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequestWritebuf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRequestWritebuf
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
func skipRequestWritebuf(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequestWritebuf
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
					return 0, ErrIntOverflowRequestWritebuf
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRequestWritebuf
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRequestWritebuf
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRequestWritebuf
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRequestWritebuf(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRequestWritebuf = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequestWritebuf   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("request_write.protobuf", fileDescriptorRequestWritebuf) }

var fileDescriptorRequestWritebuf = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x89, 0x2f, 0x2f, 0xca, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x4f,
	0x2a, 0x4d, 0x13, 0xe2, 0x2e, 0x36, 0x4a, 0x0e, 0x80, 0x72, 0x94, 0x94, 0xb8, 0x78, 0x82, 0x20,
	0xca, 0xc2, 0x41, 0xaa, 0x84, 0x84, 0xb8, 0x58, 0x02, 0x12, 0x4b, 0x32, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x58, 0x0a, 0x12, 0x4b, 0x32, 0x9c, 0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x86, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x34, 0xed, 0x26, 0x63, 0x00, 0x00, 0x00,
}
