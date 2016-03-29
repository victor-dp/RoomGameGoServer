// Code generated by protoc-gen-go.
// source: proto_files/Messages.proto
// DO NOT EDIT!

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	proto_files/Messages.proto

It has these top-level messages:
	Vec3
	Quat
	Color4
	MsgRpc
	MsgCreateObj
	MsgGetMissingCmd
	GenMessage
*/
package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type MessageType int32

const (
	MessageType_CreateObj     MessageType = 0
	MessageType_Rpc           MessageType = 1
	MessageType_Auth          MessageType = 2
	MessageType_Empty         MessageType = 3
	MessageType_AddPlayer     MessageType = 4
	MessageType_GetMissingCmd MessageType = 5
	MessageType_ReadyForGame  MessageType = 6
)

var MessageType_name = map[int32]string{
	0: "CreateObj",
	1: "Rpc",
	2: "Auth",
	3: "Empty",
	4: "AddPlayer",
	5: "GetMissingCmd",
	6: "ReadyForGame",
}
var MessageType_value = map[string]int32{
	"CreateObj":     0,
	"Rpc":           1,
	"Auth":          2,
	"Empty":         3,
	"AddPlayer":     4,
	"GetMissingCmd": 5,
	"ReadyForGame":  6,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Vec3 struct {
	X                *float32 `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	Y                *float32 `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	Z                *float32 `protobuf:"fixed32,3,req,name=z" json:"z,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Vec3) Reset()                    { *m = Vec3{} }
func (m *Vec3) String() string            { return proto.CompactTextString(m) }
func (*Vec3) ProtoMessage()               {}
func (*Vec3) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Vec3) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Vec3) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *Vec3) GetZ() float32 {
	if m != nil && m.Z != nil {
		return *m.Z
	}
	return 0
}

type Quat struct {
	X                *float32 `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	Y                *float32 `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	Z                *float32 `protobuf:"fixed32,3,req,name=z" json:"z,omitempty"`
	W                *float32 `protobuf:"fixed32,4,req,name=w" json:"w,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Quat) Reset()                    { *m = Quat{} }
func (m *Quat) String() string            { return proto.CompactTextString(m) }
func (*Quat) ProtoMessage()               {}
func (*Quat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Quat) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Quat) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *Quat) GetZ() float32 {
	if m != nil && m.Z != nil {
		return *m.Z
	}
	return 0
}

func (m *Quat) GetW() float32 {
	if m != nil && m.W != nil {
		return *m.W
	}
	return 0
}

type Color4 struct {
	R                *float32 `protobuf:"fixed32,1,req,name=r" json:"r,omitempty"`
	G                *float32 `protobuf:"fixed32,2,req,name=g" json:"g,omitempty"`
	B                *float32 `protobuf:"fixed32,3,req,name=b" json:"b,omitempty"`
	A                *float32 `protobuf:"fixed32,4,req,name=a" json:"a,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Color4) Reset()                    { *m = Color4{} }
func (m *Color4) String() string            { return proto.CompactTextString(m) }
func (*Color4) ProtoMessage()               {}
func (*Color4) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Color4) GetR() float32 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *Color4) GetG() float32 {
	if m != nil && m.G != nil {
		return *m.G
	}
	return 0
}

func (m *Color4) GetB() float32 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

func (m *Color4) GetA() float32 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

type MsgRpc struct {
	NetId            *int32  `protobuf:"varint,1,req,name=netId" json:"netId,omitempty"`
	Method           *string `protobuf:"bytes,2,req,name=method" json:"method,omitempty"`
	Argbuf           []byte  `protobuf:"bytes,3,req,name=argbuf" json:"argbuf,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgRpc) Reset()                    { *m = MsgRpc{} }
func (m *MsgRpc) String() string            { return proto.CompactTextString(m) }
func (*MsgRpc) ProtoMessage()               {}
func (*MsgRpc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MsgRpc) GetNetId() int32 {
	if m != nil && m.NetId != nil {
		return *m.NetId
	}
	return 0
}

func (m *MsgRpc) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *MsgRpc) GetArgbuf() []byte {
	if m != nil {
		return m.Argbuf
	}
	return nil
}

type MsgCreateObj struct {
	Path             *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	Pos              *Vec3   `protobuf:"bytes,2,req,name=pos" json:"pos,omitempty"`
	Rot              *Quat   `protobuf:"bytes,3,req,name=rot" json:"rot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgCreateObj) Reset()                    { *m = MsgCreateObj{} }
func (m *MsgCreateObj) String() string            { return proto.CompactTextString(m) }
func (*MsgCreateObj) ProtoMessage()               {}
func (*MsgCreateObj) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MsgCreateObj) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *MsgCreateObj) GetPos() *Vec3 {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *MsgCreateObj) GetRot() *Quat {
	if m != nil {
		return m.Rot
	}
	return nil
}

type MsgGetMissingCmd struct {
	Player           *string `protobuf:"bytes,1,req,name=player" json:"player,omitempty"`
	Frame            *int32  `protobuf:"varint,2,req,name=frame" json:"frame,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgGetMissingCmd) Reset()                    { *m = MsgGetMissingCmd{} }
func (m *MsgGetMissingCmd) String() string            { return proto.CompactTextString(m) }
func (*MsgGetMissingCmd) ProtoMessage()               {}
func (*MsgGetMissingCmd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MsgGetMissingCmd) GetPlayer() string {
	if m != nil && m.Player != nil {
		return *m.Player
	}
	return ""
}

func (m *MsgGetMissingCmd) GetFrame() int32 {
	if m != nil && m.Frame != nil {
		return *m.Frame
	}
	return 0
}

type GenMessage struct {
	MsgType          *MessageType `protobuf:"varint,1,req,name=msgType,enum=messages.MessageType" json:"msgType,omitempty"`
	PIdx             *int32       `protobuf:"varint,2,req,name=pIdx" json:"pIdx,omitempty"`
	MsgId            *int32       `protobuf:"varint,3,req,name=msgId" json:"msgId,omitempty"`
	Buf              []byte       `protobuf:"bytes,4,opt,name=buf" json:"buf,omitempty"`
	Frame            *int32       `protobuf:"varint,5,req,name=frame" json:"frame,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GenMessage) Reset()                    { *m = GenMessage{} }
func (m *GenMessage) String() string            { return proto.CompactTextString(m) }
func (*GenMessage) ProtoMessage()               {}
func (*GenMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GenMessage) GetMsgType() MessageType {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return MessageType_CreateObj
}

func (m *GenMessage) GetPIdx() int32 {
	if m != nil && m.PIdx != nil {
		return *m.PIdx
	}
	return 0
}

func (m *GenMessage) GetMsgId() int32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *GenMessage) GetBuf() []byte {
	if m != nil {
		return m.Buf
	}
	return nil
}

func (m *GenMessage) GetFrame() int32 {
	if m != nil && m.Frame != nil {
		return *m.Frame
	}
	return 0
}

func init() {
	proto.RegisterType((*Vec3)(nil), "messages.Vec3")
	proto.RegisterType((*Quat)(nil), "messages.Quat")
	proto.RegisterType((*Color4)(nil), "messages.Color4")
	proto.RegisterType((*MsgRpc)(nil), "messages.MsgRpc")
	proto.RegisterType((*MsgCreateObj)(nil), "messages.MsgCreateObj")
	proto.RegisterType((*MsgGetMissingCmd)(nil), "messages.MsgGetMissingCmd")
	proto.RegisterType((*GenMessage)(nil), "messages.GenMessage")
	proto.RegisterEnum("messages.MessageType", MessageType_name, MessageType_value)
}

var fileDescriptor0 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0x6c, 0xa7, 0xcd, 0xd4, 0xad, 0x96, 0x15, 0x20, 0x8b, 0x53, 0xe4, 0x13, 0xe2,
	0x90, 0xa2, 0xc2, 0x19, 0x51, 0x45, 0x10, 0xf5, 0x10, 0xfe, 0xac, 0x10, 0x57, 0xb4, 0xae, 0xd7,
	0x4e, 0x50, 0xfc, 0x47, 0xbb, 0x1b, 0x35, 0xee, 0x97, 0xe0, 0x2b, 0xb3, 0x33, 0x5e, 0xd3, 0xe6,
	0xd6, 0xdb, 0xfe, 0x76, 0x66, 0xdf, 0x9b, 0x79, 0x0b, 0xaf, 0x3b, 0xdd, 0xda, 0xf6, 0x77, 0xb9,
	0xdd, 0x29, 0x73, 0xb9, 0x56, 0xc6, 0xc8, 0x4a, 0x99, 0x05, 0x5d, 0xf2, 0xd3, 0xda, 0x73, 0xf6,
	0x0e, 0xa2, 0x5f, 0xea, 0xf6, 0x3d, 0x4f, 0x20, 0x38, 0xa4, 0xc1, 0x7c, 0xf2, 0x66, 0x22, 0x82,
	0x03, 0x52, 0x9f, 0x4e, 0x06, 0xea, 0x91, 0xee, 0xd3, 0x70, 0xa0, 0xfb, 0xec, 0x23, 0x44, 0x3f,
	0xf6, 0xd2, 0x3e, 0xfd, 0x05, 0xd2, 0x5d, 0x1a, 0x0d, 0x74, 0x97, 0x7d, 0x82, 0xe9, 0xb2, 0xdd,
	0xb5, 0xfa, 0x03, 0xde, 0xeb, 0x51, 0x41, 0x23, 0x55, 0xa3, 0x42, 0x85, 0x94, 0x8f, 0x0a, 0x39,
	0x92, 0x1c, 0x15, 0x64, 0xf6, 0x15, 0xa6, 0x6b, 0x53, 0x89, 0xee, 0x96, 0xbf, 0x80, 0xb8, 0x51,
	0xf6, 0xa6, 0x20, 0x95, 0x58, 0x0c, 0xc0, 0x5f, 0xc1, 0xb4, 0x56, 0x76, 0xd3, 0x16, 0x24, 0x37,
	0x13, 0x9e, 0xf0, 0x5e, 0xea, 0x2a, 0xdf, 0x97, 0x24, 0x9c, 0x08, 0x4f, 0x59, 0x09, 0x89, 0xd3,
	0x5b, 0x6a, 0x25, 0xad, 0xfa, 0x96, 0xff, 0xe1, 0x1c, 0xa2, 0x4e, 0xda, 0x0d, 0x89, 0xce, 0x04,
	0x9d, 0xf9, 0x1c, 0xc2, 0xae, 0x35, 0x24, 0x78, 0x76, 0x75, 0xb1, 0x18, 0xf3, 0x5b, 0x60, 0x78,
	0x02, 0x4b, 0xd8, 0xe1, 0xb2, 0x25, 0xe9, 0xa3, 0x0e, 0x0c, 0x4b, 0x60, 0xc9, 0x6d, 0xce, 0x9c,
	0xcf, 0x4a, 0xd9, 0xf5, 0xd6, 0x98, 0x6d, 0x53, 0x2d, 0x6b, 0x9a, 0xa9, 0xdb, 0xc9, 0x5e, 0x69,
	0xef, 0xe6, 0x09, 0x37, 0x2b, 0xb5, 0xac, 0x15, 0x39, 0xba, 0xcd, 0x08, 0xb2, 0xbf, 0x01, 0xc0,
	0x4a, 0x35, 0xfe, 0x37, 0xf9, 0x25, 0x9c, 0xd4, 0xa6, 0xfa, 0xd9, 0x77, 0x8a, 0x5e, 0x5f, 0x5c,
	0xbd, 0x7c, 0xb0, 0xf5, 0x3d, 0x58, 0x14, 0x63, 0x17, 0x6d, 0x76, 0x53, 0x1c, 0xbc, 0x28, 0x9d,
	0xd1, 0xc9, 0x95, 0x5d, 0x86, 0xe1, 0xe0, 0x44, 0xc0, 0x19, 0x84, 0x18, 0x54, 0x34, 0x0f, 0x5c,
	0x50, 0x78, 0x7c, 0x98, 0x28, 0x7e, 0x34, 0xd1, 0xdb, 0x06, 0xce, 0x1e, 0x39, 0xf1, 0x73, 0x98,
	0xfd, 0xcf, 0x91, 0x3d, 0xe3, 0x27, 0x10, 0xba, 0x6f, 0x62, 0x01, 0x3f, 0x85, 0xe8, 0x7a, 0x6f,
	0x37, 0x6c, 0xc2, 0x67, 0x10, 0x7f, 0xae, 0x3b, 0xdb, 0xb3, 0x10, 0x9b, 0xaf, 0x8b, 0xe2, 0x3b,
	0x2d, 0xcc, 0x22, 0xfe, 0x1c, 0xce, 0x8f, 0xb2, 0x61, 0xb1, 0x9b, 0x22, 0x11, 0x4a, 0x16, 0xfd,
	0x97, 0x56, 0xaf, 0x9c, 0x1b, 0x9b, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x80, 0x7d, 0x1d,
	0xd6, 0x02, 0x00, 0x00,
}