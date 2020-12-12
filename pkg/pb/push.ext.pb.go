// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push.ext.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PushCode int32

const (
	PushCode_PC_ADD_DEFAULT         PushCode = 0
	PushCode_PC_ADD_FRIEND          PushCode = 100
	PushCode_PC_AGREE_ADD_FRIEND    PushCode = 101
	PushCode_PC_UPDATE_GROUP        PushCode = 110
	PushCode_PC_ADD_GROUP_MEMBERS   PushCode = 120
	PushCode_PC_REMOVE_GROUP_MEMBER PushCode = 121
)

var PushCode_name = map[int32]string{
	0:   "PC_ADD_DEFAULT",
	100: "PC_ADD_FRIEND",
	101: "PC_AGREE_ADD_FRIEND",
	110: "PC_UPDATE_GROUP",
	120: "PC_ADD_GROUP_MEMBERS",
	121: "PC_REMOVE_GROUP_MEMBER",
}

var PushCode_value = map[string]int32{
	"PC_ADD_DEFAULT":         0,
	"PC_ADD_FRIEND":          100,
	"PC_AGREE_ADD_FRIEND":    101,
	"PC_UPDATE_GROUP":        110,
	"PC_ADD_GROUP_MEMBERS":   120,
	"PC_REMOVE_GROUP_MEMBER": 121,
}

func (x PushCode) String() string {
	return proto.EnumName(PushCode_name, int32(x))
}

func (PushCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b36e2192823d216, []int{0}
}

// 推送码 PC_ADD_FRIEND = 100
type AddFriendPush struct {
	FriendId             int64    `protobuf:"varint,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFriendPush) Reset()         { *m = AddFriendPush{} }
func (m *AddFriendPush) String() string { return proto.CompactTextString(m) }
func (*AddFriendPush) ProtoMessage()    {}
func (*AddFriendPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b36e2192823d216, []int{0}
}

func (m *AddFriendPush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFriendPush.Unmarshal(m, b)
}
func (m *AddFriendPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFriendPush.Marshal(b, m, deterministic)
}
func (m *AddFriendPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFriendPush.Merge(m, src)
}
func (m *AddFriendPush) XXX_Size() int {
	return xxx_messageInfo_AddFriendPush.Size(m)
}
func (m *AddFriendPush) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFriendPush.DiscardUnknown(m)
}

var xxx_messageInfo_AddFriendPush proto.InternalMessageInfo

func (m *AddFriendPush) GetFriendId() int64 {
	if m != nil {
		return m.FriendId
	}
	return 0
}

func (m *AddFriendPush) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AddFriendPush) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *AddFriendPush) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// 推送码 PC_AGREE_ADD_FRIEND = 101
type AgreeAddFriendPush struct {
	FriendId             int64    `protobuf:"varint,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgreeAddFriendPush) Reset()         { *m = AgreeAddFriendPush{} }
func (m *AgreeAddFriendPush) String() string { return proto.CompactTextString(m) }
func (*AgreeAddFriendPush) ProtoMessage()    {}
func (*AgreeAddFriendPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b36e2192823d216, []int{1}
}

func (m *AgreeAddFriendPush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgreeAddFriendPush.Unmarshal(m, b)
}
func (m *AgreeAddFriendPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgreeAddFriendPush.Marshal(b, m, deterministic)
}
func (m *AgreeAddFriendPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgreeAddFriendPush.Merge(m, src)
}
func (m *AgreeAddFriendPush) XXX_Size() int {
	return xxx_messageInfo_AgreeAddFriendPush.Size(m)
}
func (m *AgreeAddFriendPush) XXX_DiscardUnknown() {
	xxx_messageInfo_AgreeAddFriendPush.DiscardUnknown(m)
}

var xxx_messageInfo_AgreeAddFriendPush proto.InternalMessageInfo

func (m *AgreeAddFriendPush) GetFriendId() int64 {
	if m != nil {
		return m.FriendId
	}
	return 0
}

func (m *AgreeAddFriendPush) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AgreeAddFriendPush) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

// 更新群组 PC_UPDATE_GROUP = 110
type UpdateGroupPush struct {
	OptId                int64    `protobuf:"varint,1,opt,name=opt_id,json=optId,proto3" json:"opt_id,omitempty"`
	OptName              string   `protobuf:"bytes,2,opt,name=opt_name,json=optName,proto3" json:"opt_name,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Introduction         string   `protobuf:"bytes,5,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Extra                string   `protobuf:"bytes,6,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateGroupPush) Reset()         { *m = UpdateGroupPush{} }
func (m *UpdateGroupPush) String() string { return proto.CompactTextString(m) }
func (*UpdateGroupPush) ProtoMessage()    {}
func (*UpdateGroupPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b36e2192823d216, []int{2}
}

func (m *UpdateGroupPush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGroupPush.Unmarshal(m, b)
}
func (m *UpdateGroupPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGroupPush.Marshal(b, m, deterministic)
}
func (m *UpdateGroupPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGroupPush.Merge(m, src)
}
func (m *UpdateGroupPush) XXX_Size() int {
	return xxx_messageInfo_UpdateGroupPush.Size(m)
}
func (m *UpdateGroupPush) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGroupPush.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGroupPush proto.InternalMessageInfo

func (m *UpdateGroupPush) GetOptId() int64 {
	if m != nil {
		return m.OptId
	}
	return 0
}

func (m *UpdateGroupPush) GetOptName() string {
	if m != nil {
		return m.OptName
	}
	return ""
}

func (m *UpdateGroupPush) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateGroupPush) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *UpdateGroupPush) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *UpdateGroupPush) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

// 添加群组成员 PC_AGREE_ADD_GROUPS = 120
type AddGroupMembersPush struct {
	OptId                int64          `protobuf:"varint,1,opt,name=opt_id,json=optId,proto3" json:"opt_id,omitempty"`
	OptName              string         `protobuf:"bytes,2,opt,name=opt_name,json=optName,proto3" json:"opt_name,omitempty"`
	Members              []*GroupMember `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddGroupMembersPush) Reset()         { *m = AddGroupMembersPush{} }
func (m *AddGroupMembersPush) String() string { return proto.CompactTextString(m) }
func (*AddGroupMembersPush) ProtoMessage()    {}
func (*AddGroupMembersPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b36e2192823d216, []int{3}
}

func (m *AddGroupMembersPush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddGroupMembersPush.Unmarshal(m, b)
}
func (m *AddGroupMembersPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddGroupMembersPush.Marshal(b, m, deterministic)
}
func (m *AddGroupMembersPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddGroupMembersPush.Merge(m, src)
}
func (m *AddGroupMembersPush) XXX_Size() int {
	return xxx_messageInfo_AddGroupMembersPush.Size(m)
}
func (m *AddGroupMembersPush) XXX_DiscardUnknown() {
	xxx_messageInfo_AddGroupMembersPush.DiscardUnknown(m)
}

var xxx_messageInfo_AddGroupMembersPush proto.InternalMessageInfo

func (m *AddGroupMembersPush) GetOptId() int64 {
	if m != nil {
		return m.OptId
	}
	return 0
}

func (m *AddGroupMembersPush) GetOptName() string {
	if m != nil {
		return m.OptName
	}
	return ""
}

func (m *AddGroupMembersPush) GetMembers() []*GroupMember {
	if m != nil {
		return m.Members
	}
	return nil
}

// 删除群组成员 PC_REMOVE_GROUP_MEMBER = 121
type RemoveGroupMemberPush struct {
	OptId                int64    `protobuf:"varint,1,opt,name=opt_id,json=optId,proto3" json:"opt_id,omitempty"`
	OptName              string   `protobuf:"bytes,2,opt,name=opt_name,json=optName,proto3" json:"opt_name,omitempty"`
	DeletedUserId        int64    `protobuf:"varint,3,opt,name=deleted_user_id,json=deletedUserId,proto3" json:"deleted_user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveGroupMemberPush) Reset()         { *m = RemoveGroupMemberPush{} }
func (m *RemoveGroupMemberPush) String() string { return proto.CompactTextString(m) }
func (*RemoveGroupMemberPush) ProtoMessage()    {}
func (*RemoveGroupMemberPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b36e2192823d216, []int{4}
}

func (m *RemoveGroupMemberPush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveGroupMemberPush.Unmarshal(m, b)
}
func (m *RemoveGroupMemberPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveGroupMemberPush.Marshal(b, m, deterministic)
}
func (m *RemoveGroupMemberPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveGroupMemberPush.Merge(m, src)
}
func (m *RemoveGroupMemberPush) XXX_Size() int {
	return xxx_messageInfo_RemoveGroupMemberPush.Size(m)
}
func (m *RemoveGroupMemberPush) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveGroupMemberPush.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveGroupMemberPush proto.InternalMessageInfo

func (m *RemoveGroupMemberPush) GetOptId() int64 {
	if m != nil {
		return m.OptId
	}
	return 0
}

func (m *RemoveGroupMemberPush) GetOptName() string {
	if m != nil {
		return m.OptName
	}
	return ""
}

func (m *RemoveGroupMemberPush) GetDeletedUserId() int64 {
	if m != nil {
		return m.DeletedUserId
	}
	return 0
}

func init() {
	proto.RegisterEnum("pb.PushCode", PushCode_name, PushCode_value)
	proto.RegisterType((*AddFriendPush)(nil), "pb.AddFriendPush")
	proto.RegisterType((*AgreeAddFriendPush)(nil), "pb.AgreeAddFriendPush")
	proto.RegisterType((*UpdateGroupPush)(nil), "pb.UpdateGroupPush")
	proto.RegisterType((*AddGroupMembersPush)(nil), "pb.AddGroupMembersPush")
	proto.RegisterType((*RemoveGroupMemberPush)(nil), "pb.RemoveGroupMemberPush")
}

func init() { proto.RegisterFile("push.ext.proto", fileDescriptor_9b36e2192823d216) }

var fileDescriptor_9b36e2192823d216 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x41, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x71, 0x9d, 0xa4, 0xc9, 0x2b, 0xa9, 0xcd, 0xa4, 0x05, 0x13, 0x84, 0x14, 0x79, 0x81,
	0x02, 0x8b, 0x2c, 0xe0, 0x04, 0x26, 0x76, 0xa2, 0x48, 0xa4, 0xb5, 0x86, 0x9a, 0xed, 0xc8, 0xc9,
	0x3c, 0x5a, 0x0b, 0xc7, 0x33, 0x8c, 0xc7, 0x55, 0x38, 0x02, 0x07, 0xe0, 0x1e, 0x1c, 0x11, 0x79,
	0x1c, 0xa8, 0x8b, 0x58, 0x75, 0xc1, 0xf2, 0x7d, 0xbf, 0xe7, 0xff, 0x9e, 0x9e, 0x64, 0x38, 0x95,
	0x55, 0x79, 0x33, 0xc3, 0xbd, 0x9e, 0x49, 0x25, 0xb4, 0x20, 0x47, 0x72, 0x33, 0x76, 0x72, 0x71,
	0x9d, 0x6d, 0xef, 0xa0, 0xff, 0xdd, 0x82, 0x61, 0xc0, 0xf9, 0x42, 0x65, 0x58, 0xf0, 0xb8, 0x2a,
	0x6f, 0xc8, 0x0b, 0x18, 0x7c, 0x36, 0x13, 0xcb, 0xb8, 0x67, 0x4d, 0xac, 0xa9, 0x4d, 0xfb, 0x0d,
	0x58, 0x71, 0x32, 0x86, 0x7e, 0x91, 0x6d, 0xbf, 0x14, 0xe9, 0x0e, 0xbd, 0xa3, 0x89, 0x35, 0x1d,
	0xd0, 0x3f, 0x33, 0x79, 0x09, 0x90, 0xde, 0xa6, 0x3a, 0x55, 0xac, 0x52, 0xb9, 0x67, 0x9b, 0x74,
	0xd0, 0x90, 0x44, 0xe5, 0x64, 0x02, 0x27, 0x1c, 0xcb, 0xad, 0xca, 0xa4, 0xce, 0x44, 0xe1, 0x75,
	0x4c, 0xde, 0x46, 0x7e, 0x0e, 0x24, 0xb8, 0x56, 0x88, 0xff, 0x65, 0x1f, 0xff, 0xa7, 0x05, 0x4e,
	0x22, 0x79, 0xaa, 0x71, 0xa9, 0x44, 0x25, 0x8d, 0xeb, 0x1c, 0x7a, 0x42, 0xea, 0x3b, 0x51, 0x57,
	0x48, 0xbd, 0xe2, 0xe4, 0x39, 0xf4, 0x6b, 0xdc, 0xb2, 0x1c, 0x0b, 0xa9, 0x2f, 0x6a, 0x09, 0x81,
	0x8e, 0xc1, 0x4d, 0x7d, 0xe7, 0x1f, 0xe2, 0xce, 0xdf, 0x87, 0xf0, 0xe1, 0x71, 0x56, 0x68, 0x25,
	0x78, 0xb5, 0x35, 0x97, 0xe8, 0x9a, 0x0f, 0xee, 0x31, 0x72, 0x06, 0x5d, 0xdc, 0x6b, 0x95, 0x7a,
	0x3d, 0x13, 0x36, 0x83, 0xaf, 0x60, 0x14, 0x70, 0x6e, 0xd6, 0x5d, 0xe3, 0x6e, 0x83, 0xaa, 0x7c,
	0xe0, 0xd6, 0xaf, 0xe1, 0x78, 0xd7, 0x14, 0x78, 0xf6, 0xc4, 0x9e, 0x9e, 0xbc, 0x75, 0x66, 0x72,
	0x33, 0x6b, 0x15, 0xd3, 0xdf, 0xb9, 0xff, 0x15, 0xce, 0x29, 0xee, 0xc4, 0x2d, 0xb6, 0xd2, 0x07,
	0x5a, 0x5f, 0x81, 0xc3, 0x31, 0x47, 0x8d, 0x9c, 0x55, 0x25, 0xaa, 0xfa, 0xa9, 0x6d, 0x9e, 0x0e,
	0x0f, 0x38, 0x29, 0x51, 0xad, 0xf8, 0x9b, 0x1f, 0x16, 0xf4, 0x6b, 0xc5, 0x5c, 0xf0, 0xfa, 0xc0,
	0xa7, 0xf1, 0x9c, 0x05, 0x61, 0xc8, 0xc2, 0x68, 0x11, 0x24, 0x1f, 0xae, 0xdc, 0x47, 0xe4, 0x09,
	0x0c, 0x0f, 0x6c, 0x41, 0x57, 0xd1, 0x45, 0xe8, 0x72, 0xf2, 0x0c, 0x46, 0x35, 0x5a, 0xd2, 0x28,
	0x6a, 0x07, 0x48, 0x46, 0xe0, 0xc4, 0x73, 0x96, 0xc4, 0x61, 0x70, 0x15, 0xb1, 0x25, 0xbd, 0x4c,
	0x62, 0xb7, 0x20, 0x1e, 0x9c, 0x1d, 0x0a, 0x0c, 0x61, 0xeb, 0x68, 0xfd, 0x3e, 0xa2, 0x1f, 0xdd,
	0x3d, 0x19, 0xc3, 0xd3, 0x78, 0xce, 0x68, 0xb4, 0xbe, 0xfc, 0x14, 0xdd, 0x0b, 0xdd, 0x6f, 0x9b,
	0x9e, 0xf9, 0x65, 0xde, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x0f, 0x24, 0xe8, 0x59, 0x03,
	0x00, 0x00,
}
