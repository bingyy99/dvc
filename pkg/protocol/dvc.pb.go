// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dvc.proto

package dvc_protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DVCMessage_DVCMessageType int32

const (
	DVCMessage_JoinConference DVCMessage_DVCMessageType = 0
)

var DVCMessage_DVCMessageType_name = map[int32]string{
	0: "JoinConference",
}
var DVCMessage_DVCMessageType_value = map[string]int32{
	"JoinConference": 0,
}

func (x DVCMessage_DVCMessageType) String() string {
	return proto.EnumName(DVCMessage_DVCMessageType_name, int32(x))
}
func (DVCMessage_DVCMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dvc_c0172c7b909cfabf, []int{1, 0}
}

type JoinConferenceMessage struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ConferenceId         string   `protobuf:"bytes,3,opt,name=conference_id,json=conferenceId,proto3" json:"conference_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinConferenceMessage) Reset()         { *m = JoinConferenceMessage{} }
func (m *JoinConferenceMessage) String() string { return proto.CompactTextString(m) }
func (*JoinConferenceMessage) ProtoMessage()    {}
func (*JoinConferenceMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_dvc_c0172c7b909cfabf, []int{0}
}
func (m *JoinConferenceMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinConferenceMessage.Unmarshal(m, b)
}
func (m *JoinConferenceMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinConferenceMessage.Marshal(b, m, deterministic)
}
func (dst *JoinConferenceMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinConferenceMessage.Merge(dst, src)
}
func (m *JoinConferenceMessage) XXX_Size() int {
	return xxx_messageInfo_JoinConferenceMessage.Size(m)
}
func (m *JoinConferenceMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinConferenceMessage.DiscardUnknown(m)
}

var xxx_messageInfo_JoinConferenceMessage proto.InternalMessageInfo

func (m *JoinConferenceMessage) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *JoinConferenceMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JoinConferenceMessage) GetConferenceId() string {
	if m != nil {
		return m.ConferenceId
	}
	return ""
}

type DVCMessage struct {
	Type                 DVCMessage_DVCMessageType `protobuf:"varint,1,opt,name=type,proto3,enum=dvc.protocol.DVCMessage_DVCMessageType" json:"type,omitempty"`
	JoinConfMsg          *JoinConferenceMessage    `protobuf:"bytes,2,opt,name=join_conf_msg,json=joinConfMsg,proto3" json:"join_conf_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DVCMessage) Reset()         { *m = DVCMessage{} }
func (m *DVCMessage) String() string { return proto.CompactTextString(m) }
func (*DVCMessage) ProtoMessage()    {}
func (*DVCMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_dvc_c0172c7b909cfabf, []int{1}
}
func (m *DVCMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DVCMessage.Unmarshal(m, b)
}
func (m *DVCMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DVCMessage.Marshal(b, m, deterministic)
}
func (dst *DVCMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DVCMessage.Merge(dst, src)
}
func (m *DVCMessage) XXX_Size() int {
	return xxx_messageInfo_DVCMessage.Size(m)
}
func (m *DVCMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DVCMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DVCMessage proto.InternalMessageInfo

func (m *DVCMessage) GetType() DVCMessage_DVCMessageType {
	if m != nil {
		return m.Type
	}
	return DVCMessage_JoinConference
}

func (m *DVCMessage) GetJoinConfMsg() *JoinConferenceMessage {
	if m != nil {
		return m.JoinConfMsg
	}
	return nil
}

func init() {
	proto.RegisterType((*JoinConferenceMessage)(nil), "dvc.protocol.JoinConferenceMessage")
	proto.RegisterType((*DVCMessage)(nil), "dvc.protocol.DVCMessage")
	proto.RegisterEnum("dvc.protocol.DVCMessage_DVCMessageType", DVCMessage_DVCMessageType_name, DVCMessage_DVCMessageType_value)
}

func init() { proto.RegisterFile("dvc.proto", fileDescriptor_dvc_c0172c7b909cfabf) }

var fileDescriptor_dvc_c0172c7b909cfabf = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x29, 0x4b, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x81, 0x33, 0x93, 0xf3, 0x73, 0x94, 0x32, 0xb9, 0x44,
	0xbd, 0xf2, 0x33, 0xf3, 0x9c, 0xf3, 0xf3, 0xd2, 0x52, 0x8b, 0x52, 0xf3, 0x92, 0x53, 0x7d, 0x53,
	0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0xc4, 0xb9, 0xd8, 0x4b, 0x8b, 0x53, 0x8b, 0xe2, 0x33, 0x53,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xd8, 0x40, 0x5c, 0xcf, 0x14, 0x21, 0x21, 0x2e, 0x96,
	0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0xb0, 0x28, 0x98, 0x2d, 0xa4, 0xcc, 0xc5, 0x9b, 0x0c, 0x37,
	0x01, 0xa4, 0x85, 0x19, 0x2c, 0xc9, 0x83, 0x10, 0xf4, 0x4c, 0x51, 0xda, 0xc1, 0xc8, 0xc5, 0xe5,
	0x12, 0xe6, 0x0c, 0xb3, 0xc0, 0x9a, 0x8b, 0xa5, 0xa4, 0xb2, 0x20, 0x15, 0x6c, 0x3a, 0x9f, 0x91,
	0xba, 0x1e, 0xb2, 0xb3, 0xf4, 0x10, 0xea, 0x90, 0x98, 0x21, 0x95, 0x05, 0xa9, 0x41, 0x60, 0x4d,
	0x42, 0xee, 0x5c, 0xbc, 0x59, 0xf9, 0x99, 0x79, 0xf1, 0x20, 0x0b, 0xe2, 0x73, 0x8b, 0xd3, 0xc1,
	0xae, 0xe1, 0x36, 0x52, 0x46, 0x35, 0x05, 0xab, 0xcf, 0x82, 0xb8, 0xb3, 0xa0, 0xc2, 0xbe, 0xc5,
	0xe9, 0x4a, 0x2a, 0x5c, 0x7c, 0xa8, 0x16, 0x08, 0x09, 0x71, 0xf1, 0xa1, 0xea, 0x13, 0x60, 0x48,
	0x62, 0x03, 0x1b, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x97, 0xb7, 0xd2, 0x3e, 0x47, 0x01,
	0x00, 0x00,
}