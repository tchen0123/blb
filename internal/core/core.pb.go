// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/core/core.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	internal/core/core.proto

It has these top-level messages:
*/
package core

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// StorageClass specifies how blob data is stored. If a blob has a certain
// storage class, then all tracts in the blob are encoded with that class.
// (They may also be encoded with another class at the same time, to allow for
// transitions beteween classes).
type StorageClass int32

const (
	// Data is fully replicated n times.
	StorageClass_REPLICATED StorageClass = 0
	// Reed-Solomon N_M, with N data and M parity.
	StorageClass_RS_6_3  StorageClass = 1
	StorageClass_RS_8_3  StorageClass = 2
	StorageClass_RS_10_3 StorageClass = 3
	StorageClass_RS_12_5 StorageClass = 4
)

var StorageClass_name = map[int32]string{
	0: "REPLICATED",
	1: "RS_6_3",
	2: "RS_8_3",
	3: "RS_10_3",
	4: "RS_12_5",
}
var StorageClass_value = map[string]int32{
	"REPLICATED": 0,
	"RS_6_3":     1,
	"RS_8_3":     2,
	"RS_10_3":    3,
	"RS_12_5":    4,
}

func (x StorageClass) Enum() *StorageClass {
	p := new(StorageClass)
	*p = x
	return p
}
func (x StorageClass) String() string {
	return proto.EnumName(StorageClass_name, int32(x))
}
func (x *StorageClass) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StorageClass_value, data, "StorageClass")
	if err != nil {
		return err
	}
	*x = StorageClass(value)
	return nil
}
func (StorageClass) EnumDescriptor() ([]byte, []int) { return fileDescriptorCore, []int{0} }

// StorageHint is a hint specified by client code about how it intents to use
// this blob in the future. Blb will use the storage hint in conjunction with
// usage data that it collects to choose an appropriate storage class for the
// blob.
type StorageHint int32

const (
	StorageHint_DEFAULT StorageHint = 0
	StorageHint_HOT     StorageHint = 1
	StorageHint_WARM    StorageHint = 2
	StorageHint_COLD    StorageHint = 3
)

var StorageHint_name = map[int32]string{
	0: "DEFAULT",
	1: "HOT",
	2: "WARM",
	3: "COLD",
}
var StorageHint_value = map[string]int32{
	"DEFAULT": 0,
	"HOT":     1,
	"WARM":    2,
	"COLD":    3,
}

func (x StorageHint) Enum() *StorageHint {
	p := new(StorageHint)
	*p = x
	return p
}
func (x StorageHint) String() string {
	return proto.EnumName(StorageHint_name, int32(x))
}
func (x *StorageHint) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StorageHint_value, data, "StorageHint")
	if err != nil {
		return err
	}
	*x = StorageHint(value)
	return nil
}
func (StorageHint) EnumDescriptor() ([]byte, []int) { return fileDescriptorCore, []int{1} }

// Priority is a hint to the tractserver about how to order operations.
type Priority int32

const (
	Priority_TSDEFAULT Priority = 0
	Priority_LOW       Priority = 1
	Priority_MEDIUM    Priority = 2
	Priority_HIGH      Priority = 3
)

var Priority_name = map[int32]string{
	0: "TSDEFAULT",
	1: "LOW",
	2: "MEDIUM",
	3: "HIGH",
}
var Priority_value = map[string]int32{
	"TSDEFAULT": 0,
	"LOW":       1,
	"MEDIUM":    2,
	"HIGH":      3,
}

func (x Priority) Enum() *Priority {
	p := new(Priority)
	*p = x
	return p
}
func (x Priority) String() string {
	return proto.EnumName(Priority_name, int32(x))
}
func (x *Priority) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Priority_value, data, "Priority")
	if err != nil {
		return err
	}
	*x = Priority(value)
	return nil
}
func (Priority) EnumDescriptor() ([]byte, []int) { return fileDescriptorCore, []int{2} }

func init() {
	proto.RegisterEnum("core.StorageClass", StorageClass_name, StorageClass_value)
	proto.RegisterEnum("core.StorageHint", StorageHint_name, StorageHint_value)
	proto.RegisterEnum("core.Priority", Priority_name, Priority_value)
}

func init() { proto.RegisterFile("internal/core/core.proto", fileDescriptorCore) }

var fileDescriptorCore = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x4f, 0xce, 0x2f, 0x4a, 0x05, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x56, 0x00, 0x17, 0x4f, 0x70, 0x49, 0x7e, 0x51, 0x62, 0x7a, 0xaa,
	0x73, 0x4e, 0x62, 0x71, 0xb1, 0x10, 0x1f, 0x17, 0x57, 0x90, 0x6b, 0x80, 0x8f, 0xa7, 0xb3, 0x63,
	0x88, 0xab, 0x8b, 0x00, 0x83, 0x10, 0x17, 0x17, 0x5b, 0x50, 0x70, 0xbc, 0x59, 0xbc, 0xb1, 0x00,
	0x23, 0x94, 0x6d, 0x11, 0x6f, 0x2c, 0xc0, 0x24, 0xc4, 0xcd, 0xc5, 0x1e, 0x14, 0x1c, 0x6f, 0x68,
	0x10, 0x6f, 0x2c, 0xc0, 0x0c, 0xe3, 0x18, 0xc5, 0x9b, 0x0a, 0xb0, 0x68, 0x99, 0x73, 0x71, 0x43,
	0x4d, 0xf4, 0xc8, 0xcc, 0x2b, 0x01, 0xc9, 0xb9, 0xb8, 0xba, 0x39, 0x86, 0xfa, 0x84, 0x08, 0x30,
	0x08, 0xb1, 0x73, 0x31, 0x7b, 0xf8, 0x87, 0x08, 0x30, 0x0a, 0x71, 0x70, 0xb1, 0x84, 0x3b, 0x06,
	0xf9, 0x0a, 0x30, 0x81, 0x58, 0xce, 0xfe, 0x3e, 0x2e, 0x02, 0xcc, 0x5a, 0x16, 0x5c, 0x1c, 0x01,
	0x45, 0x99, 0xf9, 0x45, 0x99, 0x25, 0x95, 0x42, 0xbc, 0x5c, 0x9c, 0x21, 0xc1, 0x28, 0xfa, 0x7c,
	0xfc, 0xc3, 0x21, 0x4e, 0xf0, 0x75, 0x75, 0xf1, 0x0c, 0x85, 0xea, 0xf4, 0xf0, 0x74, 0xf7, 0x10,
	0x60, 0x76, 0xe2, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x81, 0x22, 0x27, 0x92, 0xf3, 0x00, 0x00, 0x00,
}
