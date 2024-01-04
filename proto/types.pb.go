// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: proto/types.proto

package proto

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

type HTTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body           []byte                   `protobuf:"bytes,1,opt,name=Body,proto3" json:"Body,omitempty"`
	Method         string                   `protobuf:"bytes,2,opt,name=Method,proto3" json:"Method,omitempty"`
	URL            string                   `protobuf:"bytes,3,opt,name=URL,proto3" json:"URL,omitempty"`
	EndpointID     string                   `protobuf:"bytes,4,opt,name=EndpointID,proto3" json:"EndpointID,omitempty"`
	ID             string                   `protobuf:"bytes,5,opt,name=ID,proto3" json:"ID,omitempty"`
	Header         map[string]*HeaderFields `protobuf:"bytes,6,rep,name=Header,proto3" json:"Header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Runtime        string                   `protobuf:"bytes,7,opt,name=runtime,proto3" json:"runtime,omitempty"`
	ActiveDeployID string                   `protobuf:"bytes,8,opt,name=activeDeployID,proto3" json:"activeDeployID,omitempty"`
	Env            map[string]string        `protobuf:"bytes,9,rep,name=Env,proto3" json:"Env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HTTPRequest) Reset() {
	*x = HTTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPRequest) ProtoMessage() {}

func (x *HTTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPRequest.ProtoReflect.Descriptor instead.
func (*HTTPRequest) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *HTTPRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HTTPRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *HTTPRequest) GetEndpointID() string {
	if x != nil {
		return x.EndpointID
	}
	return ""
}

func (x *HTTPRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *HTTPRequest) GetHeader() map[string]*HeaderFields {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *HTTPRequest) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *HTTPRequest) GetActiveDeployID() string {
	if x != nil {
		return x.ActiveDeployID
	}
	return ""
}

func (x *HTTPRequest) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

type HeaderFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *HeaderFields) Reset() {
	*x = HeaderFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderFields) ProtoMessage() {}

func (x *HeaderFields) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderFields.ProtoReflect.Descriptor instead.
func (*HeaderFields) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{1}
}

func (x *HeaderFields) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type HTTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response   []byte `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	StatusCode int32  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	RequestID  string `protobuf:"bytes,3,opt,name=RequestID,proto3" json:"RequestID,omitempty"`
}

func (x *HTTPResponse) Reset() {
	*x = HTTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPResponse) ProtoMessage() {}

func (x *HTTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPResponse.ProtoReflect.Descriptor instead.
func (*HTTPResponse) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{2}
}

func (x *HTTPResponse) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *HTTPResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *HTTPResponse) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

var File_proto_types_proto protoreflect.FileDescriptor

var file_proto_types_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x03, 0x0a, 0x0b, 0x48,
	0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f,
	0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x49, 0x44, 0x12, 0x2d, 0x0a, 0x03, 0x45, 0x6e, 0x76, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x45, 0x6e,
	0x76, 0x1a, 0x4e, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x26, 0x0a, 0x0c, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x22, 0x68, 0x0a, 0x0c, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x42, 0x1d, 0x5a, 0x1b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x74, 0x68, 0x64, 0x6d,
	0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_types_proto_rawDescOnce sync.Once
	file_proto_types_proto_rawDescData = file_proto_types_proto_rawDesc
)

func file_proto_types_proto_rawDescGZIP() []byte {
	file_proto_types_proto_rawDescOnce.Do(func() {
		file_proto_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_proto_rawDescData)
	})
	return file_proto_types_proto_rawDescData
}

var file_proto_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_types_proto_goTypes = []interface{}{
	(*HTTPRequest)(nil),  // 0: proto.HTTPRequest
	(*HeaderFields)(nil), // 1: proto.HeaderFields
	(*HTTPResponse)(nil), // 2: proto.HTTPResponse
	nil,                  // 3: proto.HTTPRequest.HeaderEntry
	nil,                  // 4: proto.HTTPRequest.EnvEntry
}
var file_proto_types_proto_depIdxs = []int32{
	3, // 0: proto.HTTPRequest.Header:type_name -> proto.HTTPRequest.HeaderEntry
	4, // 1: proto.HTTPRequest.Env:type_name -> proto.HTTPRequest.EnvEntry
	1, // 2: proto.HTTPRequest.HeaderEntry.value:type_name -> proto.HeaderFields
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_types_proto_init() }
func file_proto_types_proto_init() {
	if File_proto_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPRequest); i {
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
		file_proto_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderFields); i {
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
		file_proto_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPResponse); i {
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
			RawDescriptor: file_proto_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_proto_goTypes,
		DependencyIndexes: file_proto_types_proto_depIdxs,
		MessageInfos:      file_proto_types_proto_msgTypes,
	}.Build()
	File_proto_types_proto = out.File
	file_proto_types_proto_rawDesc = nil
	file_proto_types_proto_goTypes = nil
	file_proto_types_proto_depIdxs = nil
}
