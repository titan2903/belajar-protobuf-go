// Code generated by protoc-gen-go. DO NOT EDIT.
// source: garage.proto

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

type GarageCoordinate struct {
	Latitude             float32  `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GarageCoordinate) Reset()         { *m = GarageCoordinate{} }
func (m *GarageCoordinate) String() string { return proto.CompactTextString(m) }
func (*GarageCoordinate) ProtoMessage()    {}
func (*GarageCoordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{0}
}

func (m *GarageCoordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageCoordinate.Unmarshal(m, b)
}
func (m *GarageCoordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageCoordinate.Marshal(b, m, deterministic)
}
func (m *GarageCoordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageCoordinate.Merge(m, src)
}
func (m *GarageCoordinate) XXX_Size() int {
	return xxx_messageInfo_GarageCoordinate.Size(m)
}
func (m *GarageCoordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageCoordinate.DiscardUnknown(m)
}

var xxx_messageInfo_GarageCoordinate proto.InternalMessageInfo

func (m *GarageCoordinate) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *GarageCoordinate) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type Garage struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Coordinate           *GarageCoordinate `protobuf:"bytes,3,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Garage) Reset()         { *m = Garage{} }
func (m *Garage) String() string { return proto.CompactTextString(m) }
func (*Garage) ProtoMessage()    {}
func (*Garage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{1}
}

func (m *Garage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Garage.Unmarshal(m, b)
}
func (m *Garage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Garage.Marshal(b, m, deterministic)
}
func (m *Garage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Garage.Merge(m, src)
}
func (m *Garage) XXX_Size() int {
	return xxx_messageInfo_Garage.Size(m)
}
func (m *Garage) XXX_DiscardUnknown() {
	xxx_messageInfo_Garage.DiscardUnknown(m)
}

var xxx_messageInfo_Garage proto.InternalMessageInfo

func (m *Garage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Garage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Garage) GetCoordinate() *GarageCoordinate {
	if m != nil {
		return m.Coordinate
	}
	return nil
}

type GarageList struct {
	List                 []*Garage `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GarageList) Reset()         { *m = GarageList{} }
func (m *GarageList) String() string { return proto.CompactTextString(m) }
func (*GarageList) ProtoMessage()    {}
func (*GarageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{2}
}

func (m *GarageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageList.Unmarshal(m, b)
}
func (m *GarageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageList.Marshal(b, m, deterministic)
}
func (m *GarageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageList.Merge(m, src)
}
func (m *GarageList) XXX_Size() int {
	return xxx_messageInfo_GarageList.Size(m)
}
func (m *GarageList) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageList.DiscardUnknown(m)
}

var xxx_messageInfo_GarageList proto.InternalMessageInfo

func (m *GarageList) GetList() []*Garage {
	if m != nil {
		return m.List
	}
	return nil
}

type GarageListByUser struct {
	List                 map[string]*GarageList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GarageListByUser) Reset()         { *m = GarageListByUser{} }
func (m *GarageListByUser) String() string { return proto.CompactTextString(m) }
func (*GarageListByUser) ProtoMessage()    {}
func (*GarageListByUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{3}
}

func (m *GarageListByUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageListByUser.Unmarshal(m, b)
}
func (m *GarageListByUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageListByUser.Marshal(b, m, deterministic)
}
func (m *GarageListByUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageListByUser.Merge(m, src)
}
func (m *GarageListByUser) XXX_Size() int {
	return xxx_messageInfo_GarageListByUser.Size(m)
}
func (m *GarageListByUser) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageListByUser.DiscardUnknown(m)
}

var xxx_messageInfo_GarageListByUser proto.InternalMessageInfo

func (m *GarageListByUser) GetList() map[string]*GarageList {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*GarageCoordinate)(nil), "model.GarageCoordinate")
	proto.RegisterType((*Garage)(nil), "model.Garage")
	proto.RegisterType((*GarageList)(nil), "model.GarageList")
	proto.RegisterType((*GarageListByUser)(nil), "model.GarageListByUser")
	proto.RegisterMapType((map[string]*GarageList)(nil), "model.GarageListByUser.ListEntry")
}

func init() { proto.RegisterFile("garage.proto", fileDescriptor_d35248283e3db3c1) }

var fileDescriptor_d35248283e3db3c1 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0x26, 0x69, 0x77, 0xb1, 0x53, 0x95, 0x75, 0x2e, 0x96, 0xc5, 0x43, 0xb7, 0x17, 0x7b, 0xaa,
	0x50, 0x11, 0xc5, 0xa3, 0x22, 0x82, 0xec, 0x29, 0xe0, 0x03, 0x44, 0x13, 0x4a, 0x30, 0xdb, 0x48,
	0x36, 0x2b, 0xf4, 0x41, 0x7c, 0x5f, 0x49, 0xa2, 0xad, 0xab, 0xa7, 0xce, 0x7c, 0x7f, 0x9d, 0x8f,
	0xc0, 0x61, 0xc7, 0x2d, 0xef, 0x64, 0xf3, 0x6e, 0x8d, 0x33, 0x38, 0xdb, 0x18, 0x21, 0x75, 0xb5,
	0x86, 0xc5, 0x63, 0x80, 0xef, 0x8d, 0xb1, 0x42, 0xf5, 0xdc, 0x49, 0x5c, 0xc2, 0x81, 0xe6, 0x4e,
	0xb9, 0x9d, 0x90, 0x05, 0x29, 0x49, 0x4d, 0xd9, 0xb8, 0xe3, 0x19, 0x64, 0xda, 0xf4, 0x5d, 0x24,
	0x69, 0x20, 0x27, 0xa0, 0x92, 0x30, 0x8f, 0x69, 0x78, 0x0c, 0x54, 0x89, 0xe0, 0xce, 0x18, 0x55,
	0x02, 0x11, 0xd2, 0x9e, 0x6f, 0xa2, 0x25, 0x63, 0x61, 0xc6, 0x6b, 0x80, 0xd7, 0xf1, 0xaf, 0x45,
	0x52, 0x92, 0x3a, 0x6f, 0x4f, 0x9b, 0x70, 0x57, 0xf3, 0xf7, 0x28, 0xf6, 0x4b, 0x5a, 0x5d, 0x00,
	0x44, 0x7e, 0xad, 0xb6, 0x0e, 0x57, 0x90, 0xfa, 0x6f, 0x41, 0xca, 0xa4, 0xce, 0xdb, 0xa3, 0xbd,
	0x00, 0x16, 0xa8, 0xea, 0x93, 0xfc, 0xd4, 0xf4, 0xeb, 0xdd, 0xf0, 0xbc, 0x95, 0x16, 0xaf, 0x20,
	0xd5, 0x93, 0x6f, 0xb5, 0xe7, 0x9b, 0x64, 0x8d, 0x1f, 0x1f, 0x7a, 0x67, 0x07, 0x16, 0xe4, 0xcb,
	0x27, 0xc8, 0x46, 0x08, 0x17, 0x90, 0xbc, 0xc9, 0xe1, 0xbb, 0xa7, 0x1f, 0xf1, 0x1c, 0x66, 0x1f,
	0x5c, 0xef, 0x62, 0xd3, 0xbc, 0x3d, 0xf9, 0x17, 0xcb, 0x22, 0x7f, 0x4b, 0x6f, 0xc8, 0xcb, 0x3c,
	0xbc, 0xc5, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0xaa, 0xe1, 0x5f, 0x9b, 0x01, 0x00,
	0x00,
}