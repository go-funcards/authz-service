// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: v1/definition.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SaveDefsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Defs []*SaveDefsRequest_Def `protobuf:"bytes,1,rep,name=defs,proto3" json:"defs,omitempty"`
}

func (x *SaveDefsRequest) Reset() {
	*x = SaveDefsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_definition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveDefsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveDefsRequest) ProtoMessage() {}

func (x *SaveDefsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_definition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveDefsRequest.ProtoReflect.Descriptor instead.
func (*SaveDefsRequest) Descriptor() ([]byte, []int) {
	return file_v1_definition_proto_rawDescGZIP(), []int{0}
}

func (x *SaveDefsRequest) GetDefs() []*SaveDefsRequest_Def {
	if x != nil {
		return x.Defs
	}
	return nil
}

type DeleteDefsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefIds []string `protobuf:"bytes,1,rep,name=def_ids,json=defIds,proto3" json:"def_ids,omitempty"`
}

func (x *DeleteDefsRequest) Reset() {
	*x = DeleteDefsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_definition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDefsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDefsRequest) ProtoMessage() {}

func (x *DeleteDefsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_definition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDefsRequest.ProtoReflect.Descriptor instead.
func (*DeleteDefsRequest) Descriptor() ([]byte, []int) {
	return file_v1_definition_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteDefsRequest) GetDefIds() []string {
	if x != nil {
		return x.DefIds
	}
	return nil
}

type DefsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Defs []*DefsResponse_Def `protobuf:"bytes,1,rep,name=defs,proto3" json:"defs,omitempty"`
}

func (x *DefsResponse) Reset() {
	*x = DefsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_definition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefsResponse) ProtoMessage() {}

func (x *DefsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_definition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefsResponse.ProtoReflect.Descriptor instead.
func (*DefsResponse) Descriptor() ([]byte, []int) {
	return file_v1_definition_proto_rawDescGZIP(), []int{2}
}

func (x *DefsResponse) GetDefs() []*DefsResponse_Def {
	if x != nil {
		return x.Defs
	}
	return nil
}

type SaveDefsRequest_Def struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefId string `protobuf:"bytes,1,opt,name=def_id,json=defId,proto3" json:"def_id,omitempty"`
	Sec   string `protobuf:"bytes,2,opt,name=sec,proto3" json:"sec,omitempty"`
	Key   string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SaveDefsRequest_Def) Reset() {
	*x = SaveDefsRequest_Def{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_definition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveDefsRequest_Def) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveDefsRequest_Def) ProtoMessage() {}

func (x *SaveDefsRequest_Def) ProtoReflect() protoreflect.Message {
	mi := &file_v1_definition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveDefsRequest_Def.ProtoReflect.Descriptor instead.
func (*SaveDefsRequest_Def) Descriptor() ([]byte, []int) {
	return file_v1_definition_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SaveDefsRequest_Def) GetDefId() string {
	if x != nil {
		return x.DefId
	}
	return ""
}

func (x *SaveDefsRequest_Def) GetSec() string {
	if x != nil {
		return x.Sec
	}
	return ""
}

func (x *SaveDefsRequest_Def) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SaveDefsRequest_Def) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DefsResponse_Def struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefId string `protobuf:"bytes,1,opt,name=def_id,json=defId,proto3" json:"def_id,omitempty"`
	Sec   string `protobuf:"bytes,2,opt,name=sec,proto3" json:"sec,omitempty"`
	Key   string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DefsResponse_Def) Reset() {
	*x = DefsResponse_Def{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_definition_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefsResponse_Def) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefsResponse_Def) ProtoMessage() {}

func (x *DefsResponse_Def) ProtoReflect() protoreflect.Message {
	mi := &file_v1_definition_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefsResponse_Def.ProtoReflect.Descriptor instead.
func (*DefsResponse_Def) Descriptor() ([]byte, []int) {
	return file_v1_definition_proto_rawDescGZIP(), []int{2, 0}
}

func (x *DefsResponse_Def) GetDefId() string {
	if x != nil {
		return x.DefId
	}
	return ""
}

func (x *DefsResponse_Def) GetSec() string {
	if x != nil {
		return x.Sec
	}
	return ""
}

func (x *DefsResponse_Def) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DefsResponse_Def) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_v1_definition_proto protoreflect.FileDescriptor

var file_v1_definition_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a,
	0x0f, 0x53, 0x61, 0x76, 0x65, 0x44, 0x65, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x31, 0x0a, 0x04, 0x64, 0x65, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x44, 0x65,
	0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x66, 0x52, 0x04, 0x64,
	0x65, 0x66, 0x73, 0x1a, 0x56, 0x0a, 0x03, 0x44, 0x65, 0x66, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x65,
	0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x66, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x65, 0x66, 0x49, 0x64, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x0c, 0x44, 0x65,
	0x66, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x65,
	0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x66, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x44, 0x65, 0x66, 0x52, 0x04, 0x64, 0x65, 0x66, 0x73, 0x1a, 0x56, 0x0a, 0x03, 0x44, 0x65,
	0x66, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x65, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x64, 0x65, 0x66, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x32, 0xc9, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x44, 0x65, 0x66, 0x73, 0x12, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x44, 0x65, 0x66,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x41, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x66, 0x73, 0x12, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x65, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x66, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4d,
	0x0a, 0x1b, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0xaa, 0x02, 0x13, 0x46, 0x75, 0x6e, 0x43, 0x61, 0x72,
	0x64, 0x73, 0x4f, 0x72, 0x67, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_definition_proto_rawDescOnce sync.Once
	file_v1_definition_proto_rawDescData = file_v1_definition_proto_rawDesc
)

func file_v1_definition_proto_rawDescGZIP() []byte {
	file_v1_definition_proto_rawDescOnce.Do(func() {
		file_v1_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_definition_proto_rawDescData)
	})
	return file_v1_definition_proto_rawDescData
}

var file_v1_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_definition_proto_goTypes = []interface{}{
	(*SaveDefsRequest)(nil),     // 0: proto.v1.SaveDefsRequest
	(*DeleteDefsRequest)(nil),   // 1: proto.v1.DeleteDefsRequest
	(*DefsResponse)(nil),        // 2: proto.v1.DefsResponse
	(*SaveDefsRequest_Def)(nil), // 3: proto.v1.SaveDefsRequest.Def
	(*DefsResponse_Def)(nil),    // 4: proto.v1.DefsResponse.Def
	(*emptypb.Empty)(nil),       // 5: google.protobuf.Empty
}
var file_v1_definition_proto_depIdxs = []int32{
	3, // 0: proto.v1.SaveDefsRequest.defs:type_name -> proto.v1.SaveDefsRequest.Def
	4, // 1: proto.v1.DefsResponse.defs:type_name -> proto.v1.DefsResponse.Def
	0, // 2: proto.v1.Definition.SaveDefs:input_type -> proto.v1.SaveDefsRequest
	1, // 3: proto.v1.Definition.DeleteDefs:input_type -> proto.v1.DeleteDefsRequest
	5, // 4: proto.v1.Definition.GetDefs:input_type -> google.protobuf.Empty
	5, // 5: proto.v1.Definition.SaveDefs:output_type -> google.protobuf.Empty
	5, // 6: proto.v1.Definition.DeleteDefs:output_type -> google.protobuf.Empty
	2, // 7: proto.v1.Definition.GetDefs:output_type -> proto.v1.DefsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_definition_proto_init() }
func file_v1_definition_proto_init() {
	if File_v1_definition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_definition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveDefsRequest); i {
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
		file_v1_definition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDefsRequest); i {
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
		file_v1_definition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefsResponse); i {
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
		file_v1_definition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveDefsRequest_Def); i {
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
		file_v1_definition_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefsResponse_Def); i {
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
			RawDescriptor: file_v1_definition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_definition_proto_goTypes,
		DependencyIndexes: file_v1_definition_proto_depIdxs,
		MessageInfos:      file_v1_definition_proto_msgTypes,
	}.Build()
	File_v1_definition_proto = out.File
	file_v1_definition_proto_rawDesc = nil
	file_v1_definition_proto_goTypes = nil
	file_v1_definition_proto_depIdxs = nil
}
