// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package model

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

type UserGender int32

const (
	UserGender_UNDEFINED UserGender = 0
	UserGender_MALE      UserGender = 1
	UserGender_FEMALE    UserGender = 2
)

var UserGender_name = map[int32]string{
	0: "UNDEFINED",
	1: "MALE",
	2: "FEMALE",
}

var UserGender_value = map[string]int32{
	"UNDEFINED": 0,
	"MALE":      1,
	"FEMALE":    2,
}

func (x UserGender) String() string {
	return proto.EnumName(UserGender_name, int32(x))
}

func (UserGender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

type User struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string     `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Gender               UserGender `protobuf:"varint,4,opt,name=gender,proto3,enum=model.UserGender" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetGender() UserGender {
	if m != nil {
		return m.Gender
	}
	return UserGender_UNDEFINED
}

type UserList struct {
	List                 []*User  `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetList() []*User {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterEnum("model.UserGender", UserGender_name, UserGender_value)
	proto.RegisterType((*User)(nil), "model.User")
	proto.RegisterType((*UserList)(nil), "model.UserList")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xd1, 0x4a, 0x85, 0x40,
	0x10, 0x86, 0xdb, 0x3d, 0x9b, 0x78, 0xe6, 0x90, 0xd8, 0x5c, 0x2d, 0xdd, 0x24, 0x5e, 0x59, 0x81,
	0x90, 0x3d, 0x41, 0xe0, 0x1a, 0x81, 0x79, 0x21, 0xf8, 0x00, 0xc6, 0x2e, 0xb1, 0xa0, 0xae, 0xed,
	0x1a, 0xbd, 0x7e, 0x38, 0x44, 0x9d, 0xbb, 0x7f, 0xbe, 0x7f, 0xe6, 0x83, 0x01, 0xf8, 0x0a, 0xc6,
	0x97, 0xab, 0x77, 0x9b, 0xc3, 0xcb, 0xd9, 0x69, 0x33, 0xe5, 0x9f, 0x20, 0x86, 0x60, 0x3c, 0x26,
	0xc0, 0xad, 0x96, 0x2c, 0x63, 0xc5, 0xb1, 0xe7, 0x56, 0x23, 0x82, 0x58, 0xc6, 0xd9, 0x48, 0x4e,
	0x84, 0x32, 0xde, 0x40, 0xbc, 0x8e, 0x21, 0x7c, 0x3b, 0xaf, 0xe5, 0x81, 0xf8, 0xdf, 0x8c, 0x77,
	0x10, 0x7d, 0x98, 0x45, 0x1b, 0x2f, 0x45, 0xc6, 0x8a, 0xa4, 0xba, 0x2e, 0xc9, 0x5f, 0xee, 0xf2,
	0x17, 0x2a, 0xfa, 0xdf, 0x85, 0xfc, 0x01, 0xe2, 0x9d, 0xb6, 0x36, 0x6c, 0x78, 0x0b, 0x62, 0xb2,
	0x61, 0x93, 0x2c, 0x3b, 0x14, 0xa7, 0xea, 0x74, 0x76, 0xd4, 0x53, 0x71, 0xff, 0x08, 0xf0, 0xaf,
	0xc0, 0x2b, 0x38, 0x0e, 0x5d, 0xad, 0x9a, 0xd7, 0x4e, 0xd5, 0xe9, 0x05, 0xc6, 0x20, 0xde, 0x9e,
	0x5b, 0x95, 0x32, 0x04, 0x88, 0x1a, 0x45, 0x99, 0xbf, 0x47, 0xf4, 0xe0, 0xd3, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf6, 0x3b, 0x16, 0x35, 0xee, 0x00, 0x00, 0x00,
}
