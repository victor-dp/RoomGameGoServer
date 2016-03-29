// Code generated by protoc-gen-go.
// source: proto_files/WorldMessage.proto
// DO NOT EDIT!

/*
Package world_messages is a generated protocol buffer package.

It is generated from these files:
	proto_files/WorldMessage.proto

It has these top-level messages:
	MsgPlayerEnterRoom
	MsgEnterRoomReply
	WorldMessage
	Room
	ReplyMsg
	MsgGetRoomListReply
	MsgCreateRoomReply
*/
package world_messages

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
	MessageType_CreateRoom       MessageType = 1
	MessageType_EnterRoom        MessageType = 2
	MessageType_QuitRoom         MessageType = 3
	MessageType_GetRoomList      MessageType = 4
	MessageType_CreateRoomReply  MessageType = 5
	MessageType_EnterRoomReply   MessageType = 6
	MessageType_QuitRoomReply    MessageType = 7
	MessageType_GetRoomListReply MessageType = 8
	MessageType_PlayerEnterRoom  MessageType = 9
)

var MessageType_name = map[int32]string{
	1: "CreateRoom",
	2: "EnterRoom",
	3: "QuitRoom",
	4: "GetRoomList",
	5: "CreateRoomReply",
	6: "EnterRoomReply",
	7: "QuitRoomReply",
	8: "GetRoomListReply",
	9: "PlayerEnterRoom",
}
var MessageType_value = map[string]int32{
	"CreateRoom":       1,
	"EnterRoom":        2,
	"QuitRoom":         3,
	"GetRoomList":      4,
	"CreateRoomReply":  5,
	"EnterRoomReply":   6,
	"QuitRoomReply":    7,
	"GetRoomListReply": 8,
	"PlayerEnterRoom":  9,
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

type EnterRoomResult int32

const (
	EnterRoomResult_Ok            EnterRoomResult = 1
	EnterRoomResult_AlreadyIn     EnterRoomResult = 2
	EnterRoomResult_OutOfCapacity EnterRoomResult = 3
	EnterRoomResult_RoomNotExists EnterRoomResult = 4
)

var EnterRoomResult_name = map[int32]string{
	1: "Ok",
	2: "AlreadyIn",
	3: "OutOfCapacity",
	4: "RoomNotExists",
}
var EnterRoomResult_value = map[string]int32{
	"Ok":            1,
	"AlreadyIn":     2,
	"OutOfCapacity": 3,
	"RoomNotExists": 4,
}

func (x EnterRoomResult) Enum() *EnterRoomResult {
	p := new(EnterRoomResult)
	*p = x
	return p
}
func (x EnterRoomResult) String() string {
	return proto.EnumName(EnterRoomResult_name, int32(x))
}
func (x *EnterRoomResult) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnterRoomResult_value, data, "EnterRoomResult")
	if err != nil {
		return err
	}
	*x = EnterRoomResult(value)
	return nil
}
func (EnterRoomResult) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type MsgPlayerEnterRoom struct {
	RoomId           *string `protobuf:"bytes,1,req,name=roomId" json:"roomId,omitempty"`
	PlayerId         *string `protobuf:"bytes,2,req,name=playerId" json:"playerId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgPlayerEnterRoom) Reset()                    { *m = MsgPlayerEnterRoom{} }
func (m *MsgPlayerEnterRoom) String() string            { return proto.CompactTextString(m) }
func (*MsgPlayerEnterRoom) ProtoMessage()               {}
func (*MsgPlayerEnterRoom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MsgPlayerEnterRoom) GetRoomId() string {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return ""
}

func (m *MsgPlayerEnterRoom) GetPlayerId() string {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return ""
}

type MsgEnterRoomReply struct {
	Players          []string         `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
	Result           *EnterRoomResult `protobuf:"varint,2,req,name=result,enum=world_messages.EnterRoomResult" json:"result,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MsgEnterRoomReply) Reset()                    { *m = MsgEnterRoomReply{} }
func (m *MsgEnterRoomReply) String() string            { return proto.CompactTextString(m) }
func (*MsgEnterRoomReply) ProtoMessage()               {}
func (*MsgEnterRoomReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MsgEnterRoomReply) GetPlayers() []string {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *MsgEnterRoomReply) GetResult() EnterRoomResult {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return EnterRoomResult_Ok
}

type WorldMessage struct {
	Type             *MessageType `protobuf:"varint,1,req,name=type,enum=world_messages.MessageType" json:"type,omitempty"`
	PlayerId         *string      `protobuf:"bytes,2,req,name=playerId" json:"playerId,omitempty"`
	Buff             []byte       `protobuf:"bytes,3,opt,name=buff" json:"buff,omitempty"`
	MsgId            *int32       `protobuf:"varint,4,req,name=msgId" json:"msgId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *WorldMessage) Reset()                    { *m = WorldMessage{} }
func (m *WorldMessage) String() string            { return proto.CompactTextString(m) }
func (*WorldMessage) ProtoMessage()               {}
func (*WorldMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WorldMessage) GetType() MessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MessageType_CreateRoom
}

func (m *WorldMessage) GetPlayerId() string {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return ""
}

func (m *WorldMessage) GetBuff() []byte {
	if m != nil {
		return m.Buff
	}
	return nil
}

func (m *WorldMessage) GetMsgId() int32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

type Room struct {
	Id               *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	PlayerCount      *int32  `protobuf:"varint,2,req,name=playerCount" json:"playerCount,omitempty"`
	Capacity         *int32  `protobuf:"varint,3,req,name=capacity" json:"capacity,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Room) Reset()                    { *m = Room{} }
func (m *Room) String() string            { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()               {}
func (*Room) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Room) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Room) GetPlayerCount() int32 {
	if m != nil && m.PlayerCount != nil {
		return *m.PlayerCount
	}
	return 0
}

func (m *Room) GetCapacity() int32 {
	if m != nil && m.Capacity != nil {
		return *m.Capacity
	}
	return 0
}

type ReplyMsg struct {
	Type             *MessageType `protobuf:"varint,1,req,name=type,enum=world_messages.MessageType" json:"type,omitempty"`
	MsgId            *int32       `protobuf:"varint,2,req,name=msgId" json:"msgId,omitempty"`
	Buff             []byte       `protobuf:"bytes,3,req,name=buff" json:"buff,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ReplyMsg) Reset()                    { *m = ReplyMsg{} }
func (m *ReplyMsg) String() string            { return proto.CompactTextString(m) }
func (*ReplyMsg) ProtoMessage()               {}
func (*ReplyMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReplyMsg) GetType() MessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MessageType_CreateRoom
}

func (m *ReplyMsg) GetMsgId() int32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *ReplyMsg) GetBuff() []byte {
	if m != nil {
		return m.Buff
	}
	return nil
}

type MsgGetRoomListReply struct {
	Rooms            []*Room `protobuf:"bytes,1,rep,name=rooms" json:"rooms,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgGetRoomListReply) Reset()                    { *m = MsgGetRoomListReply{} }
func (m *MsgGetRoomListReply) String() string            { return proto.CompactTextString(m) }
func (*MsgGetRoomListReply) ProtoMessage()               {}
func (*MsgGetRoomListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MsgGetRoomListReply) GetRooms() []*Room {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type MsgCreateRoomReply struct {
	ErrorCode        *int32  `protobuf:"varint,1,req,name=errorCode" json:"errorCode,omitempty"`
	Id               *string `protobuf:"bytes,2,req,name=id" json:"id,omitempty"`
	Capacity         *int32  `protobuf:"varint,3,req,name=capacity" json:"capacity,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgCreateRoomReply) Reset()                    { *m = MsgCreateRoomReply{} }
func (m *MsgCreateRoomReply) String() string            { return proto.CompactTextString(m) }
func (*MsgCreateRoomReply) ProtoMessage()               {}
func (*MsgCreateRoomReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MsgCreateRoomReply) GetErrorCode() int32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *MsgCreateRoomReply) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *MsgCreateRoomReply) GetCapacity() int32 {
	if m != nil && m.Capacity != nil {
		return *m.Capacity
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgPlayerEnterRoom)(nil), "world_messages.MsgPlayerEnterRoom")
	proto.RegisterType((*MsgEnterRoomReply)(nil), "world_messages.MsgEnterRoomReply")
	proto.RegisterType((*WorldMessage)(nil), "world_messages.WorldMessage")
	proto.RegisterType((*Room)(nil), "world_messages.Room")
	proto.RegisterType((*ReplyMsg)(nil), "world_messages.ReplyMsg")
	proto.RegisterType((*MsgGetRoomListReply)(nil), "world_messages.MsgGetRoomListReply")
	proto.RegisterType((*MsgCreateRoomReply)(nil), "world_messages.MsgCreateRoomReply")
	proto.RegisterEnum("world_messages.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("world_messages.EnterRoomResult", EnterRoomResult_name, EnterRoomResult_value)
}

var fileDescriptor0 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x45, 0xb2, 0xe4, 0xd8, 0x63, 0x47, 0xde, 0x4c, 0x42, 0x31, 0x6d, 0x69, 0x8d, 0x4e, 0xc1,
	0x07, 0x07, 0x7c, 0xe9, 0x39, 0x98, 0xd0, 0x06, 0x9a, 0xb8, 0x5d, 0x02, 0xbd, 0x35, 0xa8, 0xf1,
	0xca, 0x88, 0xca, 0x96, 0xd0, 0xae, 0x68, 0xf5, 0x01, 0xfd, 0xa1, 0x7e, 0x61, 0x67, 0x77, 0x65,
	0x59, 0xd1, 0x21, 0xd0, 0x93, 0x35, 0xb3, 0x6f, 0xdf, 0xbc, 0xf7, 0x66, 0x0d, 0xef, 0xf2, 0x22,
	0x53, 0xd9, 0x63, 0x9c, 0xa4, 0x42, 0x5e, 0x7d, 0xcb, 0x8a, 0x74, 0x73, 0x27, 0xa4, 0x8c, 0xb6,
	0x62, 0x61, 0x0e, 0x30, 0xf8, 0xa5, 0x7b, 0x8f, 0x3b, 0xdb, 0x94, 0xe1, 0x27, 0xc0, 0x3b, 0xb9,
	0xfd, 0x92, 0x46, 0x95, 0x28, 0x6e, 0xf6, 0x4a, 0x14, 0x3c, 0xcb, 0x76, 0xf8, 0x0a, 0xfa, 0x05,
	0xfd, 0xde, 0x6e, 0xa6, 0xce, 0xcc, 0xbd, 0x1c, 0xf2, 0xba, 0xc2, 0xd7, 0x30, 0xc8, 0x0d, 0x94,
	0x4e, 0x5c, 0x73, 0xd2, 0xd4, 0x61, 0x0c, 0x67, 0xc4, 0xd4, 0x70, 0x70, 0x91, 0xa7, 0x15, 0x4e,
	0xe1, 0xc4, 0x02, 0x24, 0x31, 0xf5, 0x08, 0x7f, 0x28, 0xf1, 0x03, 0x8d, 0x10, 0xb2, 0x4c, 0x95,
	0x21, 0x0a, 0x96, 0xef, 0x17, 0xcf, 0x95, 0x2d, 0x5a, 0x4c, 0x1a, 0xc6, 0x6b, 0x78, 0xf8, 0xc7,
	0x81, 0x71, 0xdb, 0x18, 0x5e, 0x81, 0xa7, 0xaa, 0x5c, 0x18, 0xa9, 0xc1, 0xf2, 0x4d, 0x97, 0xa7,
	0x86, 0x3d, 0x10, 0x84, 0x1b, 0xe0, 0x4b, 0x2e, 0x10, 0xc1, 0xfb, 0x51, 0xc6, 0xf1, 0xb4, 0x37,
	0x73, 0x2e, 0xc7, 0xdc, 0x7c, 0xe3, 0x05, 0xf8, 0x3b, 0xb9, 0x25, 0xb0, 0x47, 0x60, 0x9f, 0xdb,
	0x22, 0x7c, 0x00, 0xcf, 0x64, 0x15, 0x80, 0x9b, 0x1c, 0x72, 0xa2, 0x2f, 0x9c, 0xc1, 0xc8, 0xb2,
	0xad, 0xb2, 0x72, 0x6f, 0xdd, 0xf9, 0xbc, 0xdd, 0xd2, 0xf3, 0x9f, 0xa2, 0x3c, 0x7a, 0x4a, 0x54,
	0x45, 0x73, 0xf4, 0x71, 0x53, 0x87, 0x02, 0x06, 0x26, 0x39, 0x8a, 0xf2, 0xff, 0x8d, 0x35, 0x42,
	0xdd, 0x96, 0xd0, 0x96, 0x25, 0xf7, 0x60, 0x29, 0xbc, 0x86, 0x73, 0x9a, 0xf0, 0x51, 0x28, 0x6d,
	0xe1, 0x73, 0x22, 0x95, 0x5d, 0xd7, 0x1c, 0x7c, 0xbd, 0x69, 0xbb, 0xac, 0xd1, 0xf2, 0xa2, 0x3b,
	0xd2, 0xac, 0xc3, 0x42, 0xc2, 0xef, 0xe6, 0xe5, 0xac, 0x0a, 0x11, 0x29, 0x71, 0x5c, 0xf8, 0x5b,
	0x18, 0x8a, 0xa2, 0xc8, 0xc8, 0xe9, 0xc6, 0x0a, 0xf7, 0xf9, 0xb1, 0x51, 0x67, 0xe5, 0x36, 0x59,
	0xbd, 0x90, 0xc4, 0xfc, 0xaf, 0x03, 0xa3, 0x96, 0x45, 0xba, 0x0b, 0xc7, 0x61, 0xcc, 0xc1, 0x53,
	0x18, 0x36, 0x4f, 0x84, 0xb9, 0x38, 0x86, 0xc1, 0xd7, 0x32, 0x31, 0x7e, 0x58, 0x0f, 0x27, 0x30,
	0x6a, 0x99, 0x63, 0x1e, 0x9e, 0xc3, 0xa4, 0x23, 0x95, 0xf9, 0x94, 0x4c, 0xf0, 0xfc, 0xbd, 0xb2,
	0x3e, 0x9e, 0xc1, 0xe9, 0x81, 0xc7, 0xb6, 0x4e, 0x28, 0x56, 0xd6, 0x4d, 0x8a, 0x0d, 0x34, 0x63,
	0xe7, 0x6f, 0xc3, 0x86, 0xf3, 0x7b, 0x98, 0x74, 0xde, 0x2d, 0xf6, 0xc1, 0x5d, 0xff, 0xb4, 0x7a,
	0xaf, 0x53, 0x92, 0xb0, 0xa9, 0x6e, 0xf7, 0xa4, 0x97, 0xe6, 0xac, 0x4b, 0xb5, 0x8e, 0x57, 0xb5,
	0x5f, 0x12, 0x4d, 0x2d, 0x7d, 0xef, 0x3e, 0x53, 0x37, 0xbf, 0x69, 0x90, 0x64, 0xde, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x23, 0xb8, 0x5d, 0x93, 0xcf, 0x03, 0x00, 0x00,
}