// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: diode/v1/site.proto

package diodepb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// A site
type Site struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Slug *string `protobuf:"bytes,2,opt,name=slug,proto3,oneof" json:"slug,omitempty"`
}

func (x *Site) Reset() {
	*x = Site{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diode_v1_site_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Site) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Site) ProtoMessage() {}

func (x *Site) ProtoReflect() protoreflect.Message {
	mi := &file_diode_v1_site_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Site.ProtoReflect.Descriptor instead.
func (*Site) Descriptor() ([]byte, []int) {
	return file_diode_v1_site_proto_rawDescGZIP(), []int{0}
}

func (x *Site) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Site) GetSlug() string {
	if x != nil && x.Slug != nil {
		return *x.Slug
	}
	return ""
}

var File_diode_v1_site_proto protoreflect.FileDescriptor

var file_diode_v1_site_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x69, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x69, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x04, 0x53, 0x69, 0x74, 0x65,
	0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1b, 0xfa, 0x42, 0x18, 0x72, 0x16, 0x10, 0x01, 0x18, 0x64, 0x32, 0x10, 0x5e,
	0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x2b, 0x24, 0x48,
	0x01, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x42, 0x44, 0x5a, 0x42,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x74, 0x62, 0x6f,
	0x78, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x64, 0x69, 0x6f, 0x64, 0x65, 0x2d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x69, 0x6f, 0x64, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67,
	0x6f, 0x2f, 0x64, 0x69, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x6f, 0x64, 0x65,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_diode_v1_site_proto_rawDescOnce sync.Once
	file_diode_v1_site_proto_rawDescData = file_diode_v1_site_proto_rawDesc
)

func file_diode_v1_site_proto_rawDescGZIP() []byte {
	file_diode_v1_site_proto_rawDescOnce.Do(func() {
		file_diode_v1_site_proto_rawDescData = protoimpl.X.CompressGZIP(file_diode_v1_site_proto_rawDescData)
	})
	return file_diode_v1_site_proto_rawDescData
}

var file_diode_v1_site_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_diode_v1_site_proto_goTypes = []interface{}{
	(*Site)(nil), // 0: diode.v1.Site
}
var file_diode_v1_site_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_diode_v1_site_proto_init() }
func file_diode_v1_site_proto_init() {
	if File_diode_v1_site_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_diode_v1_site_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Site); i {
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
	file_diode_v1_site_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_diode_v1_site_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_diode_v1_site_proto_goTypes,
		DependencyIndexes: file_diode_v1_site_proto_depIdxs,
		MessageInfos:      file_diode_v1_site_proto_msgTypes,
	}.Build()
	File_diode_v1_site_proto = out.File
	file_diode_v1_site_proto_rawDesc = nil
	file_diode_v1_site_proto_goTypes = nil
	file_diode_v1_site_proto_depIdxs = nil
}
