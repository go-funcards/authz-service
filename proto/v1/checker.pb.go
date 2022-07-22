// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: v1/checker.proto

package v1

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

type IsGrantedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params []string `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *IsGrantedRequest) Reset() {
	*x = IsGrantedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_checker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsGrantedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsGrantedRequest) ProtoMessage() {}

func (x *IsGrantedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_checker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsGrantedRequest.ProtoReflect.Descriptor instead.
func (*IsGrantedRequest) Descriptor() ([]byte, []int) {
	return file_v1_checker_proto_rawDescGZIP(), []int{0}
}

func (x *IsGrantedRequest) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

type Granted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yes bool `protobuf:"varint,1,opt,name=yes,proto3" json:"yes,omitempty"`
}

func (x *Granted) Reset() {
	*x = Granted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_checker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Granted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Granted) ProtoMessage() {}

func (x *Granted) ProtoReflect() protoreflect.Message {
	mi := &file_v1_checker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Granted.ProtoReflect.Descriptor instead.
func (*Granted) Descriptor() ([]byte, []int) {
	return file_v1_checker_proto_rawDescGZIP(), []int{1}
}

func (x *Granted) GetYes() bool {
	if x != nil {
		return x.Yes
	}
	return false
}

var File_v1_checker_proto protoreflect.FileDescriptor

var file_v1_checker_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x2a, 0x0a, 0x10,
	0x49, 0x73, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x1b, 0x0a, 0x07, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x79, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x79, 0x65, 0x73, 0x32, 0x52, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x3a, 0x0a,
	0x09, 0x49, 0x73, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x42, 0x57, 0x0a, 0x1b, 0x6f, 0x72, 0x67,
	0x2e, 0x66, 0x75, 0x6e, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x19, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0xaa, 0x02, 0x13, 0x46,
	0x75, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x73, 0x4f, 0x72, 0x67, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_checker_proto_rawDescOnce sync.Once
	file_v1_checker_proto_rawDescData = file_v1_checker_proto_rawDesc
)

func file_v1_checker_proto_rawDescGZIP() []byte {
	file_v1_checker_proto_rawDescOnce.Do(func() {
		file_v1_checker_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_checker_proto_rawDescData)
	})
	return file_v1_checker_proto_rawDescData
}

var file_v1_checker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_checker_proto_goTypes = []interface{}{
	(*IsGrantedRequest)(nil), // 0: proto.v1.IsGrantedRequest
	(*Granted)(nil),          // 1: proto.v1.Granted
}
var file_v1_checker_proto_depIdxs = []int32{
	0, // 0: proto.v1.AuthorizationChecker.IsGranted:input_type -> proto.v1.IsGrantedRequest
	1, // 1: proto.v1.AuthorizationChecker.IsGranted:output_type -> proto.v1.Granted
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_checker_proto_init() }
func file_v1_checker_proto_init() {
	if File_v1_checker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_checker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsGrantedRequest); i {
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
		file_v1_checker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Granted); i {
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
			RawDescriptor: file_v1_checker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_checker_proto_goTypes,
		DependencyIndexes: file_v1_checker_proto_depIdxs,
		MessageInfos:      file_v1_checker_proto_msgTypes,
	}.Build()
	File_v1_checker_proto = out.File
	file_v1_checker_proto_rawDesc = nil
	file_v1_checker_proto_goTypes = nil
	file_v1_checker_proto_depIdxs = nil
}
