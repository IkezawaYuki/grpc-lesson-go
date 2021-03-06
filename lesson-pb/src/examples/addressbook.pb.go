// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/addressbook.proto

package tutorial

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d4cca6f7b9e4178c, []int{0, 0}
}

// [START messages]
type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4cca6f7b9e4178c, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4cca6f7b9e4178c, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4cca6f7b9e4178c, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
}

func init() {
	proto.RegisterFile("examples/addressbook.proto", fileDescriptor_d4cca6f7b9e4178c)
}

var fileDescriptor_d4cca6f7b9e4178c = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x51, 0x4b, 0x83, 0x50,
	0x14, 0xc7, 0xd3, 0x39, 0xd9, 0x8e, 0x31, 0xec, 0x32, 0x42, 0x24, 0x48, 0x46, 0x0f, 0x42, 0x70,
	0x07, 0x2b, 0xe8, 0xa9, 0x87, 0x84, 0x51, 0x51, 0x6b, 0x22, 0x1b, 0x3d, 0xc6, 0x35, 0x6f, 0x4b,
	0xa6, 0xde, 0x8b, 0xf7, 0x0a, 0xed, 0x2b, 0xf5, 0x15, 0xfa, 0x72, 0xa1, 0x57, 0xc7, 0x88, 0xde,
	0xce, 0x3d, 0xe7, 0xc7, 0xf1, 0xef, 0xef, 0x80, 0x4b, 0xbf, 0x48, 0xce, 0x33, 0x2a, 0xa6, 0x24,
	0x49, 0x4a, 0x2a, 0x44, 0xcc, 0xd8, 0x16, 0xf3, 0x92, 0x49, 0x86, 0x06, 0xb2, 0x92, 0xac, 0x4c,
	0x49, 0xe6, 0x9e, 0x6f, 0x18, 0xdb, 0x64, 0x74, 0xda, 0xf4, 0xe3, 0xea, 0x63, 0x2a, 0xd3, 0x9c,
	0x0a, 0x49, 0x72, 0xae, 0xd0, 0xc9, 0x8f, 0x0e, 0x66, 0x48, 0x4b, 0xc1, 0x0a, 0x84, 0xc0, 0x28,
	0x48, 0x4e, 0x1d, 0xcd, 0xd3, 0xfc, 0x61, 0xd4, 0xd4, 0x68, 0x04, 0x7a, 0x9a, 0x38, 0xba, 0xa7,
	0xf9, 0xfd, 0x48, 0x4f, 0x13, 0x34, 0x86, 0x3e, 0xcd, 0x49, 0x9a, 0x39, 0xbd, 0x06, 0x52, 0x0f,
	0x74, 0x0d, 0x26, 0xff, 0x64, 0x05, 0x15, 0x8e, 0xe1, 0xf5, 0x7c, 0x6b, 0x76, 0x86, 0xbb, 0x00,
	0x58, 0xed, 0xc6, 0x61, 0x3d, 0x7e, 0xa9, 0xf2, 0x98, 0x96, 0x51, 0xcb, 0xa2, 0x5b, 0x38, 0xce,
	0x88, 0x90, 0x6f, 0x15, 0x4f, 0x88, 0xa4, 0x89, 0xd3, 0xf7, 0x34, 0xdf, 0x9a, 0xb9, 0x58, 0x45,
	0xc6, 0x5d, 0x64, 0xbc, 0xea, 0x22, 0x47, 0x56, 0xcd, 0xaf, 0x15, 0xee, 0xae, 0xc1, 0x3a, 0xd8,
	0x8a, 0x4e, 0xc1, 0x2c, 0x9a, 0xaa, 0xcd, 0xdf, 0xbe, 0x10, 0x06, 0x43, 0xee, 0x38, 0x6d, 0xfe,
	0x61, 0x34, 0x73, 0xff, 0x4f, 0xb6, 0xda, 0x71, 0x1a, 0x35, 0xdc, 0xe4, 0x12, 0x86, 0xfb, 0x16,
	0x02, 0x30, 0x17, 0xcb, 0xe0, 0xf1, 0x79, 0x6e, 0x1f, 0xa1, 0x01, 0x18, 0x0f, 0xcb, 0xc5, 0xdc,
	0xd6, 0xea, 0xea, 0x75, 0x19, 0x3d, 0xd9, 0xfa, 0xe4, 0x06, 0xac, 0x3b, 0x65, 0x3f, 0x60, 0x6c,
	0x8b, 0x7c, 0x30, 0x39, 0x65, 0x3c, 0xab, 0x1d, 0xd6, 0x1e, 0xec, 0xbf, 0x5f, 0x8b, 0xda, 0x79,
	0x10, 0xc2, 0xf8, 0x9d, 0xe5, 0xb8, 0xbd, 0xe1, 0x1e, 0x0b, 0x4e, 0x0e, 0xd6, 0x85, 0xb5, 0x00,
	0xf1, 0xad, 0x5f, 0xdc, 0x2b, 0x21, 0x61, 0x27, 0x64, 0xde, 0x5e, 0x1e, 0x1f, 0xc0, 0xb1, 0xd9,
	0xf8, 0xba, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x61, 0xc5, 0x79, 0x68, 0x18, 0x02, 0x00, 0x00,
}
