// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: reconciler/v1/reconciler.proto

package reconcilerpb

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

// An ingestion data source
type IngestionDataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ApiKey string `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
}

func (x *IngestionDataSource) Reset() {
	*x = IngestionDataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconciler_v1_reconciler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestionDataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestionDataSource) ProtoMessage() {}

func (x *IngestionDataSource) ProtoReflect() protoreflect.Message {
	mi := &file_reconciler_v1_reconciler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestionDataSource.ProtoReflect.Descriptor instead.
func (*IngestionDataSource) Descriptor() ([]byte, []int) {
	return file_reconciler_v1_reconciler_proto_rawDescGZIP(), []int{0}
}

func (x *IngestionDataSource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IngestionDataSource) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

// The request to retrieve ingestion data sources
type RetrieveIngestionDataSourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SdkName    string `protobuf:"bytes,2,opt,name=sdk_name,json=sdkName,proto3" json:"sdk_name,omitempty"`
	SdkVersion string `protobuf:"bytes,3,opt,name=sdk_version,json=sdkVersion,proto3" json:"sdk_version,omitempty"`
}

func (x *RetrieveIngestionDataSourcesRequest) Reset() {
	*x = RetrieveIngestionDataSourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconciler_v1_reconciler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveIngestionDataSourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveIngestionDataSourcesRequest) ProtoMessage() {}

func (x *RetrieveIngestionDataSourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reconciler_v1_reconciler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveIngestionDataSourcesRequest.ProtoReflect.Descriptor instead.
func (*RetrieveIngestionDataSourcesRequest) Descriptor() ([]byte, []int) {
	return file_reconciler_v1_reconciler_proto_rawDescGZIP(), []int{1}
}

func (x *RetrieveIngestionDataSourcesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RetrieveIngestionDataSourcesRequest) GetSdkName() string {
	if x != nil {
		return x.SdkName
	}
	return ""
}

func (x *RetrieveIngestionDataSourcesRequest) GetSdkVersion() string {
	if x != nil {
		return x.SdkVersion
	}
	return ""
}

// The response from the retrieve ingestion data sources request
type RetrieveIngestionDataSourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IngestionDataSources []*IngestionDataSource `protobuf:"bytes,1,rep,name=ingestion_data_sources,json=ingestionDataSources,proto3" json:"ingestion_data_sources,omitempty"`
}

func (x *RetrieveIngestionDataSourcesResponse) Reset() {
	*x = RetrieveIngestionDataSourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconciler_v1_reconciler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveIngestionDataSourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveIngestionDataSourcesResponse) ProtoMessage() {}

func (x *RetrieveIngestionDataSourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reconciler_v1_reconciler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveIngestionDataSourcesResponse.ProtoReflect.Descriptor instead.
func (*RetrieveIngestionDataSourcesResponse) Descriptor() ([]byte, []int) {
	return file_reconciler_v1_reconciler_proto_rawDescGZIP(), []int{2}
}

func (x *RetrieveIngestionDataSourcesResponse) GetIngestionDataSources() []*IngestionDataSource {
	if x != nil {
		return x.IngestionDataSources
	}
	return nil
}

var File_reconciler_v1_reconciler_proto protoreflect.FileDescriptor

var file_reconciler_v1_reconciler_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x13, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x28, 0x18, 0x28, 0x52, 0x06, 0x61, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x22, 0xab, 0x01, 0x0a, 0x23, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05,
	0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x73,
	0x64, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x07, 0x73, 0x64, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x73, 0x64, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xfa, 0x42, 0x19, 0x72, 0x17, 0x32, 0x15,
	0x5e, 0x28, 0x5c, 0x64, 0x29, 0x2b, 0x5c, 0x2e, 0x28, 0x5c, 0x64, 0x29, 0x2b, 0x5c, 0x2e, 0x28,
	0x5c, 0x64, 0x29, 0x2b, 0x24, 0x52, 0x0a, 0x73, 0x64, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x80, 0x01, 0x0a, 0x24, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x49, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x16, 0x69, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x14,
	0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x32, 0x9f, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69,
	0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x1c, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x32, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x33, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x74, 0x62, 0x6f, 0x78, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x64, 0x69, 0x6f, 0x64, 0x65, 0x2f, 0x64, 0x69, 0x6f, 0x64, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reconciler_v1_reconciler_proto_rawDescOnce sync.Once
	file_reconciler_v1_reconciler_proto_rawDescData = file_reconciler_v1_reconciler_proto_rawDesc
)

func file_reconciler_v1_reconciler_proto_rawDescGZIP() []byte {
	file_reconciler_v1_reconciler_proto_rawDescOnce.Do(func() {
		file_reconciler_v1_reconciler_proto_rawDescData = protoimpl.X.CompressGZIP(file_reconciler_v1_reconciler_proto_rawDescData)
	})
	return file_reconciler_v1_reconciler_proto_rawDescData
}

var file_reconciler_v1_reconciler_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_reconciler_v1_reconciler_proto_goTypes = []interface{}{
	(*IngestionDataSource)(nil),                  // 0: reconciler.v1.IngestionDataSource
	(*RetrieveIngestionDataSourcesRequest)(nil),  // 1: reconciler.v1.RetrieveIngestionDataSourcesRequest
	(*RetrieveIngestionDataSourcesResponse)(nil), // 2: reconciler.v1.RetrieveIngestionDataSourcesResponse
}
var file_reconciler_v1_reconciler_proto_depIdxs = []int32{
	0, // 0: reconciler.v1.RetrieveIngestionDataSourcesResponse.ingestion_data_sources:type_name -> reconciler.v1.IngestionDataSource
	1, // 1: reconciler.v1.ReconcilerService.RetrieveIngestionDataSources:input_type -> reconciler.v1.RetrieveIngestionDataSourcesRequest
	2, // 2: reconciler.v1.ReconcilerService.RetrieveIngestionDataSources:output_type -> reconciler.v1.RetrieveIngestionDataSourcesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reconciler_v1_reconciler_proto_init() }
func file_reconciler_v1_reconciler_proto_init() {
	if File_reconciler_v1_reconciler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reconciler_v1_reconciler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestionDataSource); i {
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
		file_reconciler_v1_reconciler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveIngestionDataSourcesRequest); i {
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
		file_reconciler_v1_reconciler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveIngestionDataSourcesResponse); i {
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
			RawDescriptor: file_reconciler_v1_reconciler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reconciler_v1_reconciler_proto_goTypes,
		DependencyIndexes: file_reconciler_v1_reconciler_proto_depIdxs,
		MessageInfos:      file_reconciler_v1_reconciler_proto_msgTypes,
	}.Build()
	File_reconciler_v1_reconciler_proto = out.File
	file_reconciler_v1_reconciler_proto_rawDesc = nil
	file_reconciler_v1_reconciler_proto_goTypes = nil
	file_reconciler_v1_reconciler_proto_depIdxs = nil
}
