// Code generated by protoc-gen-go.
// source: proto_files/RoomMessage.proto
// DO NOT EDIT!

/*
Package room_messages is a generated protocol buffer package.

It is generated from these files:
	proto_files/RoomMessage.proto

It has these top-level messages:
	Vec3
	Quat
	Color4
	Rpc
	CreateObj
	GetMissingCmd
	ReadyForGame
	Empty
*/
package room_messages

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

type Rpc struct {
	NetId            *int32  `protobuf:"varint,1,req,name=netId" json:"netId,omitempty"`
	Method           *string `protobuf:"bytes,2,req,name=method" json:"method,omitempty"`
	Argbuf           []byte  `protobuf:"bytes,3,req,name=argbuf" json:"argbuf,omitempty"`
	Frame            *int32  `protobuf:"varint,4,req,name=frame" json:"frame,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Rpc) Reset()                    { *m = Rpc{} }
func (m *Rpc) String() string            { return proto.CompactTextString(m) }
func (*Rpc) ProtoMessage()               {}
func (*Rpc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Rpc) GetNetId() int32 {
	if m != nil && m.NetId != nil {
		return *m.NetId
	}
	return 0
}

func (m *Rpc) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *Rpc) GetArgbuf() []byte {
	if m != nil {
		return m.Argbuf
	}
	return nil
}

func (m *Rpc) GetFrame() int32 {
	if m != nil && m.Frame != nil {
		return *m.Frame
	}
	return 0
}

type CreateObj struct {
	Path             *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	Pos              *Vec3   `protobuf:"bytes,2,req,name=pos" json:"pos,omitempty"`
	Rot              *Quat   `protobuf:"bytes,3,req,name=rot" json:"rot,omitempty"`
	Frame            *int32  `protobuf:"varint,4,req,name=frame" json:"frame,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateObj) Reset()                    { *m = CreateObj{} }
func (m *CreateObj) String() string            { return proto.CompactTextString(m) }
func (*CreateObj) ProtoMessage()               {}
func (*CreateObj) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateObj) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *CreateObj) GetPos() *Vec3 {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *CreateObj) GetRot() *Quat {
	if m != nil {
		return m.Rot
	}
	return nil
}

func (m *CreateObj) GetFrame() int32 {
	if m != nil && m.Frame != nil {
		return *m.Frame
	}
	return 0
}

type GetMissingCmd struct {
	PlayerIndex      *int32 `protobuf:"varint,1,req,name=playerIndex" json:"playerIndex,omitempty"`
	Frame            *int32 `protobuf:"varint,2,req,name=frame" json:"frame,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetMissingCmd) Reset()                    { *m = GetMissingCmd{} }
func (m *GetMissingCmd) String() string            { return proto.CompactTextString(m) }
func (*GetMissingCmd) ProtoMessage()               {}
func (*GetMissingCmd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetMissingCmd) GetPlayerIndex() int32 {
	if m != nil && m.PlayerIndex != nil {
		return *m.PlayerIndex
	}
	return 0
}

func (m *GetMissingCmd) GetFrame() int32 {
	if m != nil && m.Frame != nil {
		return *m.Frame
	}
	return 0
}

type ReadyForGame struct {
	Frame            *int32 `protobuf:"varint,1,req,name=frame" json:"frame,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReadyForGame) Reset()                    { *m = ReadyForGame{} }
func (m *ReadyForGame) String() string            { return proto.CompactTextString(m) }
func (*ReadyForGame) ProtoMessage()               {}
func (*ReadyForGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadyForGame) GetFrame() int32 {
	if m != nil && m.Frame != nil {
		return *m.Frame
	}
	return 0
}

type Empty struct {
	Frame            *int32 `protobuf:"varint,1,req,name=frame" json:"frame,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Empty) GetFrame() int32 {
	if m != nil && m.Frame != nil {
		return *m.Frame
	}
	return 0
}

func init() {
	proto.RegisterType((*Vec3)(nil), "room_messages.Vec3")
	proto.RegisterType((*Quat)(nil), "room_messages.Quat")
	proto.RegisterType((*Color4)(nil), "room_messages.Color4")
	proto.RegisterType((*Rpc)(nil), "room_messages.Rpc")
	proto.RegisterType((*CreateObj)(nil), "room_messages.CreateObj")
	proto.RegisterType((*GetMissingCmd)(nil), "room_messages.GetMissingCmd")
	proto.RegisterType((*ReadyForGame)(nil), "room_messages.ReadyForGame")
	proto.RegisterType((*Empty)(nil), "room_messages.Empty")
}

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x6b, 0xc2, 0x30,
	0x18, 0xc5, 0xb1, 0xb6, 0x42, 0x3f, 0x75, 0x87, 0x08, 0xa3, 0x17, 0x41, 0x7a, 0xd9, 0x4e, 0x0e,
	0xd4, 0xff, 0xa0, 0x6c, 0xe2, 0x41, 0xc6, 0x3c, 0xec, 0x2a, 0xa9, 0xfd, 0x5a, 0x3b, 0x9a, 0x26,
	0x24, 0x91, 0xd9, 0xfd, 0xf5, 0x6b, 0xbe, 0xb9, 0x95, 0xb9, 0xc3, 0x6e, 0x2f, 0x2f, 0xef, 0xf7,
	0x42, 0x1e, 0x4c, 0x95, 0x96, 0x56, 0xee, 0xf3, 0xb2, 0x42, 0xf3, 0xb0, 0x93, 0x52, 0x6c, 0xd1,
	0x18, 0x5e, 0xe0, 0x9c, 0x7c, 0x36, 0xd6, 0xad, 0xb5, 0x17, 0x5f, 0x9e, 0x89, 0xef, 0xc0, 0x7f,
	0xc5, 0xc3, 0x92, 0x85, 0xd0, 0x3b, 0x47, 0xbd, 0x99, 0x77, 0xef, 0x39, 0xd9, 0x44, 0xde, 0xb7,
	0xfc, 0x88, 0xfa, 0x4e, 0xc6, 0x0b, 0xf0, 0x5f, 0x4e, 0xdc, 0xfe, 0x1b, 0x74, 0xf2, 0x3d, 0xf2,
	0x89, 0x59, 0xc1, 0x20, 0x91, 0x95, 0xd4, 0x2b, 0x67, 0xea, 0x8e, 0x2a, 0x3a, 0x2a, 0xed, 0x28,
	0x7e, 0xa1, 0x12, 0xe8, 0xef, 0xd4, 0x81, 0x8d, 0x21, 0xa8, 0xd1, 0x6e, 0x32, 0xc2, 0x02, 0x76,
	0x03, 0x03, 0x81, 0xf6, 0x28, 0x33, 0x62, 0x43, 0x77, 0xe6, 0xba, 0x48, 0x4f, 0x39, 0x15, 0x8c,
	0x5c, 0x3c, 0xd7, 0x5c, 0x20, 0x95, 0x04, 0x71, 0x05, 0x61, 0xa2, 0x91, 0x5b, 0x7c, 0x4e, 0xdf,
	0xd8, 0x08, 0x7c, 0xc5, 0xed, 0x91, 0x9a, 0x42, 0x36, 0x83, 0xbe, 0x92, 0x86, 0x6a, 0x86, 0x8b,
	0xc9, 0xfc, 0xd7, 0x1e, 0x73, 0x1a, 0xa3, 0x4d, 0xb4, 0x5b, 0x51, 0xf1, 0xdf, 0x04, 0xad, 0x70,
	0xf5, 0xda, 0x12, 0xc6, 0x6b, 0xb4, 0xdb, 0xd2, 0x98, 0xb2, 0x2e, 0x12, 0x91, 0xb1, 0x09, 0x0c,
	0x55, 0xc5, 0x1b, 0xd4, 0x9b, 0x3a, 0xc3, 0xf3, 0xe5, 0x0b, 0x3f, 0x90, 0x47, 0xd0, 0x14, 0x46,
	0x3b, 0xe4, 0x59, 0xf3, 0x24, 0xf5, 0xba, 0x75, 0xbb, 0x6b, 0x4a, 0xc7, 0xb7, 0x10, 0x3c, 0x0a,
	0x65, 0x9b, 0x2b, 0xff, 0x33, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xbb, 0xca, 0x06, 0xe0, 0x01, 0x00,
	0x00,
}