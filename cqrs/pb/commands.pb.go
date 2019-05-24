// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commands.proto

package pb // import "github.com/go-ocf/kit/cqrs/pb"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AuthorizationContext struct {
	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceId    string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	AccessToken string `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (m *AuthorizationContext) Reset()         { *m = AuthorizationContext{} }
func (m *AuthorizationContext) String() string { return proto.CompactTextString(m) }
func (*AuthorizationContext) ProtoMessage()    {}
func (*AuthorizationContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_4971704f2dbf9596, []int{0}
}
func (m *AuthorizationContext) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthorizationContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthorizationContext.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AuthorizationContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationContext.Merge(dst, src)
}
func (m *AuthorizationContext) XXX_Size() int {
	return m.Size()
}
func (m *AuthorizationContext) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationContext.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationContext proto.InternalMessageInfo

func (m *AuthorizationContext) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AuthorizationContext) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *AuthorizationContext) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type CommandMetadata struct {
	ConnectionId string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	Sequence     uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *CommandMetadata) Reset()         { *m = CommandMetadata{} }
func (m *CommandMetadata) String() string { return proto.CompactTextString(m) }
func (*CommandMetadata) ProtoMessage()    {}
func (*CommandMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_4971704f2dbf9596, []int{1}
}
func (m *CommandMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommandMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommandMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CommandMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandMetadata.Merge(dst, src)
}
func (m *CommandMetadata) XXX_Size() int {
	return m.Size()
}
func (m *CommandMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CommandMetadata proto.InternalMessageInfo

func (m *CommandMetadata) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *CommandMetadata) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthorizationContext)(nil), "ocf.cloud.pb.AuthorizationContext")
	proto.RegisterType((*CommandMetadata)(nil), "ocf.cloud.pb.CommandMetadata")
}
func (m *AuthorizationContext) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthorizationContext) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCommands(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if len(m.DeviceId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCommands(dAtA, i, uint64(len(m.DeviceId)))
		i += copy(dAtA[i:], m.DeviceId)
	}
	if len(m.AccessToken) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCommands(dAtA, i, uint64(len(m.AccessToken)))
		i += copy(dAtA[i:], m.AccessToken)
	}
	return i, nil
}

func (m *CommandMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommandMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ConnectionId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCommands(dAtA, i, uint64(len(m.ConnectionId)))
		i += copy(dAtA[i:], m.ConnectionId)
	}
	if m.Sequence != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCommands(dAtA, i, uint64(m.Sequence))
	}
	return i, nil
}

func encodeVarintCommands(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AuthorizationContext) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovCommands(uint64(l))
	}
	l = len(m.DeviceId)
	if l > 0 {
		n += 1 + l + sovCommands(uint64(l))
	}
	l = len(m.AccessToken)
	if l > 0 {
		n += 1 + l + sovCommands(uint64(l))
	}
	return n
}

func (m *CommandMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovCommands(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovCommands(uint64(m.Sequence))
	}
	return n
}

func sovCommands(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCommands(x uint64) (n int) {
	return sovCommands(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuthorizationContext) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommands
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
			return fmt.Errorf("proto: AuthorizationContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthorizationContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommands
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
				return ErrInvalidLengthCommands
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommands
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
				return ErrInvalidLengthCommands
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommands
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
				return ErrInvalidLengthCommands
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommands(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommands
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
func (m *CommandMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommands
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
			return fmt.Errorf("proto: CommandMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommandMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommands
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
				return ErrInvalidLengthCommands
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommands
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommands(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommands
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
func skipCommands(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommands
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
					return 0, ErrIntOverflowCommands
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
					return 0, ErrIntOverflowCommands
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
				return 0, ErrInvalidLengthCommands
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCommands
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
				next, err := skipCommands(dAtA[start:])
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
	ErrInvalidLengthCommands = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommands   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("commands.proto", fileDescriptor_commands_4971704f2dbf9596) }

var fileDescriptor_commands_4971704f2dbf9596 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x13, 0x40, 0xa5, 0x35, 0x01, 0x24, 0x0b, 0x89, 0x0a, 0x24, 0xab, 0x94, 0x85, 0x85,
	0x64, 0x60, 0x83, 0x09, 0x3a, 0x65, 0x60, 0x89, 0x98, 0x58, 0x2a, 0xe7, 0xf9, 0xb5, 0xb5, 0x4a,
	0xfc, 0xd2, 0xd8, 0x46, 0x88, 0xaf, 0xe0, 0xb3, 0x18, 0x3b, 0x32, 0xa2, 0xe4, 0x47, 0x50, 0x12,
	0x89, 0x8e, 0xf7, 0x9c, 0xe1, 0x5e, 0x5d, 0x76, 0x02, 0x54, 0x14, 0xd2, 0x28, 0x1b, 0x97, 0x15,
	0x39, 0xe2, 0x11, 0xc1, 0x22, 0x86, 0x37, 0xf2, 0x2a, 0x2e, 0xf3, 0x29, 0xb1, 0xb3, 0x47, 0xef,
	0x56, 0x54, 0xe9, 0x4f, 0xe9, 0x34, 0x99, 0x19, 0x19, 0x87, 0x1f, 0x8e, 0x9f, 0xb3, 0x43, 0x6f,
	0xb1, 0x9a, 0x6b, 0x35, 0x0e, 0x27, 0xe1, 0xcd, 0x28, 0x1b, 0xb4, 0x31, 0x55, 0xfc, 0x92, 0x8d,
	0x14, 0xbe, 0x6b, 0xc0, 0x56, 0xed, 0x75, 0x6a, 0xd8, 0x83, 0x54, 0xf1, 0x2b, 0x16, 0x49, 0x00,
	0xb4, 0x76, 0xee, 0x68, 0x8d, 0x66, 0xbc, 0xdf, 0xf9, 0xa3, 0x9e, 0xbd, 0xb4, 0x68, 0x9a, 0xb1,
	0xd3, 0x59, 0x3f, 0xe8, 0x19, 0x9d, 0x54, 0xd2, 0x49, 0x7e, 0xcd, 0x8e, 0x81, 0x8c, 0x41, 0x68,
	0x07, 0xec, 0x1a, 0xa3, 0x1d, 0x4c, 0x15, 0xbf, 0x60, 0x43, 0x8b, 0x1b, 0x8f, 0x06, 0xb0, 0xab,
	0x3d, 0xc8, 0xfe, 0xf3, 0xd3, 0xfd, 0x77, 0x2d, 0xc2, 0x6d, 0x2d, 0xc2, 0xdf, 0x5a, 0x84, 0x5f,
	0x8d, 0x08, 0xb6, 0x8d, 0x08, 0x7e, 0x1a, 0x11, 0xbc, 0x4e, 0x96, 0xda, 0xad, 0x7c, 0x1e, 0x03,
	0x15, 0xc9, 0x92, 0x6e, 0x09, 0x16, 0xc9, 0x5a, 0xbb, 0x04, 0x36, 0x95, 0x4d, 0xca, 0xfc, 0xa1,
	0xcc, 0xf3, 0x41, 0xf7, 0xca, 0xdd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x1d, 0x36, 0xe6,
	0x27, 0x01, 0x00, 0x00,
}