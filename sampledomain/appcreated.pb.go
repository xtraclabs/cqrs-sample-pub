// Code generated by protoc-gen-go.
// source: appcreated.proto
// DO NOT EDIT!

/*
Package domain is a generated protocol buffer package.

It is generated from these files:
	appcreated.proto

It has these top-level messages:
	ApplicationRegistrationCreated
*/
package sampledomain

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

type ApplicationRegistrationCreated struct {
	AggregateId     string `protobuf:"bytes,1,opt,name=aggregateId" json:"aggregateId,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description     string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	CreateTimestamp int64  `protobuf:"varint,4,opt,name=createTimestamp" json:"createTimestamp,omitempty"`
}

func (m *ApplicationRegistrationCreated) Reset()                    { *m = ApplicationRegistrationCreated{} }
func (m *ApplicationRegistrationCreated) String() string            { return proto.CompactTextString(m) }
func (*ApplicationRegistrationCreated) ProtoMessage()               {}
func (*ApplicationRegistrationCreated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*ApplicationRegistrationCreated)(nil), "domain.ApplicationRegistrationCreated")
}

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0x28, 0x48,
	0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xc9,
	0xcf, 0x4d, 0xcc, 0xcc, 0x53, 0x5a, 0xc4, 0xc8, 0x25, 0xe7, 0x58, 0x50, 0x90, 0x93, 0x99, 0x9c,
	0x58, 0x92, 0x99, 0x9f, 0x17, 0x94, 0x9a, 0x9e, 0x59, 0x5c, 0x52, 0x04, 0x66, 0x3b, 0x43, 0x34,
	0x08, 0x29, 0x70, 0x71, 0x27, 0xa6, 0xa7, 0x17, 0xa5, 0xa6, 0x03, 0x79, 0x9e, 0x29, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xc8, 0x42, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12,
	0x4c, 0x60, 0x29, 0x30, 0x1b, 0xa4, 0x2b, 0x25, 0xb5, 0x38, 0xb9, 0x28, 0xb3, 0x00, 0x64, 0x96,
	0x04, 0x33, 0x44, 0x17, 0x92, 0x90, 0x90, 0x06, 0x17, 0x3f, 0xc4, 0x4d, 0x21, 0x99, 0xb9, 0xa9,
	0xc5, 0x25, 0x89, 0xb9, 0x05, 0x12, 0x2c, 0x40, 0x55, 0xcc, 0x41, 0xe8, 0xc2, 0x49, 0x6c, 0x60,
	0x37, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x4d, 0x52, 0x87, 0xc7, 0x00, 0x00, 0x00,
}