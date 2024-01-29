// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: api/proto/sand.proto

package sand

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

type Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beginning uint64 `protobuf:"varint,1,opt,name=beginning,proto3" json:"beginning,omitempty"`
	Ending    uint64 `protobuf:"varint,2,opt,name=ending,proto3" json:"ending,omitempty"`
}

func (x *Range) Reset() {
	*x = Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_sand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Range) ProtoMessage() {}

func (x *Range) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Range.ProtoReflect.Descriptor instead.
func (*Range) Descriptor() ([]byte, []int) {
	return file_api_proto_sand_proto_rawDescGZIP(), []int{0}
}

func (x *Range) GetBeginning() uint64 {
	if x != nil {
		return x.Beginning
	}
	return 0
}

func (x *Range) GetEnding() uint64 {
	if x != nil {
		return x.Ending
	}
	return 0
}

var File_api_proto_sand_proto protoreflect.FileDescriptor

var file_api_proto_sand_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x61, 0x6e, 0x64, 0x22, 0x3d, 0x0a, 0x05,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x2d, 0x63, 0x61,
	0x70, 0x74, 0x2f, 0x53, 0x61, 0x4e, 0x44, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_sand_proto_rawDescOnce sync.Once
	file_api_proto_sand_proto_rawDescData = file_api_proto_sand_proto_rawDesc
)

func file_api_proto_sand_proto_rawDescGZIP() []byte {
	file_api_proto_sand_proto_rawDescOnce.Do(func() {
		file_api_proto_sand_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_sand_proto_rawDescData)
	})
	return file_api_proto_sand_proto_rawDescData
}

var file_api_proto_sand_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_proto_sand_proto_goTypes = []interface{}{
	(*Range)(nil), // 0: sand.Range
}
var file_api_proto_sand_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_sand_proto_init() }
func file_api_proto_sand_proto_init() {
	if File_api_proto_sand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_sand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Range); i {
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
			RawDescriptor: file_api_proto_sand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_sand_proto_goTypes,
		DependencyIndexes: file_api_proto_sand_proto_depIdxs,
		MessageInfos:      file_api_proto_sand_proto_msgTypes,
	}.Build()
	File_api_proto_sand_proto = out.File
	file_api_proto_sand_proto_rawDesc = nil
	file_api_proto_sand_proto_goTypes = nil
	file_api_proto_sand_proto_depIdxs = nil
}
