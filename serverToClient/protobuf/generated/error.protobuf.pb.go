// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: error.protobuf

package s2cProtobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ErrorType int32

const (
	ErrorType_AccessDenied ErrorType = 0
)

var ErrorType_name = map[int32]string{
	0: "AccessDenied",
}
var ErrorType_value = map[string]int32{
	"AccessDenied": 0,
}

func (x ErrorType) String() string {
	return proto.EnumName(ErrorType_name, int32(x))
}
func (ErrorType) EnumDescriptor() ([]byte, []int) { return fileDescriptorErrorbuf, []int{0} }

type Error struct {
	MessageId   int64     `protobuf:"varint,1,opt,name=MessageId,json=messageId,proto3" json:"MessageId,omitempty"`
	Type        ErrorType `protobuf:"varint,2,opt,name=Type,json=type,proto3,enum=s2cProtobuf.ErrorType" json:"Type,omitempty"`
	Description string    `protobuf:"bytes,3,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptorErrorbuf, []int{0} }

func (m *Error) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *Error) GetType() ErrorType {
	if m != nil {
		return m.Type
	}
	return ErrorType_AccessDenied
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "s2cProtobuf.Error")
	proto.RegisterEnum("s2cProtobuf.ErrorType", ErrorType_name, ErrorType_value)
}
func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MessageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintErrorbuf(dAtA, i, uint64(m.MessageId))
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintErrorbuf(dAtA, i, uint64(m.Type))
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintErrorbuf(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	return i, nil
}

func encodeVarintErrorbuf(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Error) Size() (n int) {
	var l int
	_ = l
	if m.MessageId != 0 {
		n += 1 + sovErrorbuf(uint64(m.MessageId))
	}
	if m.Type != 0 {
		n += 1 + sovErrorbuf(uint64(m.Type))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovErrorbuf(uint64(l))
	}
	return n
}

func sovErrorbuf(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozErrorbuf(x uint64) (n int) {
	return sovErrorbuf(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrorbuf
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
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			m.MessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrorbuf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrorbuf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (ErrorType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrorbuf
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
				return ErrInvalidLengthErrorbuf
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrorbuf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrorbuf
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
func skipErrorbuf(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrorbuf
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
					return 0, ErrIntOverflowErrorbuf
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
					return 0, ErrIntOverflowErrorbuf
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
				return 0, ErrInvalidLengthErrorbuf
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowErrorbuf
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
				next, err := skipErrorbuf(dAtA[start:])
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
	ErrInvalidLengthErrorbuf = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrorbuf   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("error.protobuf", fileDescriptorErrorbuf) }

var fileDescriptorErrorbuf = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2d, 0x2a, 0xca,
	0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0x13, 0xe2, 0x2e, 0x36, 0x4a, 0x0e,
	0x80, 0x72, 0x94, 0x8a, 0xb9, 0x58, 0x5d, 0x41, 0xd2, 0x42, 0x32, 0x5c, 0x9c, 0xbe, 0xa9, 0xc5,
	0xc5, 0x89, 0xe9, 0xa9, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x9c, 0xb9, 0x30,
	0x01, 0x21, 0x2d, 0x2e, 0x96, 0x90, 0xca, 0x82, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x3e, 0x23,
	0x31, 0x3d, 0x24, 0x23, 0xf4, 0xc0, 0xfa, 0x41, 0xb2, 0x41, 0x2c, 0x25, 0x95, 0x05, 0xa9, 0x42,
	0x0a, 0x5c, 0xdc, 0x2e, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0x99, 0xf9, 0x79, 0x12, 0xcc,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xdc, 0x29, 0x08, 0x21, 0x2d, 0x59, 0x2e, 0x4e, 0xb8, 0x26, 0x21,
	0x01, 0x2e, 0x1e, 0xc7, 0xe4, 0xe4, 0xd4, 0xe2, 0x62, 0x97, 0xd4, 0xbc, 0xcc, 0xd4, 0x14, 0x01,
	0x06, 0x27, 0x81, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc6, 0x63, 0x39, 0x86, 0x24, 0x36, 0xb0, 0xe3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x45,
	0x4f, 0x39, 0xf9, 0xcb, 0x00, 0x00, 0x00,
}