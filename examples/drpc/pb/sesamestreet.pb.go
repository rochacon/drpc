// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: sesamestreet.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Cookie_Type int32

const (
	Cookie_Sugar      Cookie_Type = 0
	Cookie_Oatmeal    Cookie_Type = 1
	Cookie_Chocolate  Cookie_Type = 2
	Cookie_Nutella    Cookie_Type = 3
	Cookie_Strawberry Cookie_Type = 4
	Cookie_Vanilla    Cookie_Type = 5
)

// Enum value maps for Cookie_Type.
var (
	Cookie_Type_name = map[int32]string{
		0: "Sugar",
		1: "Oatmeal",
		2: "Chocolate",
		3: "Nutella",
		4: "Strawberry",
		5: "Vanilla",
	}
	Cookie_Type_value = map[string]int32{
		"Sugar":      0,
		"Oatmeal":    1,
		"Chocolate":  2,
		"Nutella":    3,
		"Strawberry": 4,
		"Vanilla":    5,
	}
)

func (x Cookie_Type) Enum() *Cookie_Type {
	p := new(Cookie_Type)
	*p = x
	return p
}

func (x Cookie_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cookie_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_sesamestreet_proto_enumTypes[0].Descriptor()
}

func (Cookie_Type) Type() protoreflect.EnumType {
	return &file_sesamestreet_proto_enumTypes[0]
}

func (x Cookie_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cookie_Type.Descriptor instead.
func (Cookie_Type) EnumDescriptor() ([]byte, []int) {
	return file_sesamestreet_proto_rawDescGZIP(), []int{1, 0}
}

type CookieJar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cookie    *Cookie `protobuf:"bytes,1,opt,name=cookie,proto3" json:"cookie,omitempty"`
	Remaining int64   `protobuf:"varint,2,opt,name=remaining,proto3" json:"remaining,omitempty"`
}

func (x *CookieJar) Reset() {
	*x = CookieJar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sesamestreet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CookieJar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CookieJar) ProtoMessage() {}

func (x *CookieJar) ProtoReflect() protoreflect.Message {
	mi := &file_sesamestreet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CookieJar.ProtoReflect.Descriptor instead.
func (*CookieJar) Descriptor() ([]byte, []int) {
	return file_sesamestreet_proto_rawDescGZIP(), []int{0}
}

func (x *CookieJar) GetCookie() *Cookie {
	if x != nil {
		return x.Cookie
	}
	return nil
}

func (x *CookieJar) GetRemaining() int64 {
	if x != nil {
		return x.Remaining
	}
	return 0
}

type Cookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Cookie_Type `protobuf:"varint,1,opt,name=type,proto3,enum=sesamestreet.Cookie_Type" json:"type,omitempty"`
}

func (x *Cookie) Reset() {
	*x = Cookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sesamestreet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cookie) ProtoMessage() {}

func (x *Cookie) ProtoReflect() protoreflect.Message {
	mi := &file_sesamestreet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cookie.ProtoReflect.Descriptor instead.
func (*Cookie) Descriptor() ([]byte, []int) {
	return file_sesamestreet_proto_rawDescGZIP(), []int{1}
}

func (x *Cookie) GetType() Cookie_Type {
	if x != nil {
		return x.Type
	}
	return Cookie_Sugar
}

var File_sesamestreet_proto protoreflect.FileDescriptor

var file_sesamestreet_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x65, 0x73, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x65, 0x73, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x22, 0x57, 0x0a, 0x09, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x4a, 0x61, 0x72, 0x12,
	0x2c, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x65, 0x73, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x43,
	0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x90, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x73, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x57, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x53, 0x75, 0x67, 0x61, 0x72, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x61, 0x74, 0x6d,
	0x65, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x68, 0x6f, 0x63, 0x6f, 0x6c, 0x61,
	0x74, 0x65, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x75, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x10,
	0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x61, 0x77, 0x62, 0x65, 0x72, 0x72, 0x79, 0x10,
	0x04, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61, 0x10, 0x05, 0x32, 0x8a,
	0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x3b, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4a, 0x61, 0x72, 0x12, 0x14, 0x2e, 0x73,
	0x65, 0x73, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x73, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x4a, 0x61, 0x72, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x09, 0x45, 0x61, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x73,
	0x61, 0x6d, 0x65, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x73, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x4a, 0x61, 0x72, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x73,
	0x74, 0x6f, 0x72, 0x6a, 0x2e, 0x69, 0x6f, 0x2f, 0x64, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x64, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sesamestreet_proto_rawDescOnce sync.Once
	file_sesamestreet_proto_rawDescData = file_sesamestreet_proto_rawDesc
)

func file_sesamestreet_proto_rawDescGZIP() []byte {
	file_sesamestreet_proto_rawDescOnce.Do(func() {
		file_sesamestreet_proto_rawDescData = protoimpl.X.CompressGZIP(file_sesamestreet_proto_rawDescData)
	})
	return file_sesamestreet_proto_rawDescData
}

var file_sesamestreet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sesamestreet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sesamestreet_proto_goTypes = []interface{}{
	(Cookie_Type)(0),  // 0: sesamestreet.Cookie.Type
	(*CookieJar)(nil), // 1: sesamestreet.CookieJar
	(*Cookie)(nil),    // 2: sesamestreet.Cookie
}
var file_sesamestreet_proto_depIdxs = []int32{
	2, // 0: sesamestreet.CookieJar.cookie:type_name -> sesamestreet.Cookie
	0, // 1: sesamestreet.Cookie.type:type_name -> sesamestreet.Cookie.Type
	2, // 2: sesamestreet.CookieMonster.CheckJar:input_type -> sesamestreet.Cookie
	2, // 3: sesamestreet.CookieMonster.EatCookie:input_type -> sesamestreet.Cookie
	1, // 4: sesamestreet.CookieMonster.CheckJar:output_type -> sesamestreet.CookieJar
	1, // 5: sesamestreet.CookieMonster.EatCookie:output_type -> sesamestreet.CookieJar
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sesamestreet_proto_init() }
func file_sesamestreet_proto_init() {
	if File_sesamestreet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sesamestreet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CookieJar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sesamestreet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cookie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sesamestreet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sesamestreet_proto_goTypes,
		DependencyIndexes: file_sesamestreet_proto_depIdxs,
		EnumInfos:         file_sesamestreet_proto_enumTypes,
		MessageInfos:      file_sesamestreet_proto_msgTypes,
	}.Build()
	File_sesamestreet_proto = out.File
	file_sesamestreet_proto_rawDesc = nil
	file_sesamestreet_proto_goTypes = nil
	file_sesamestreet_proto_depIdxs = nil
}
