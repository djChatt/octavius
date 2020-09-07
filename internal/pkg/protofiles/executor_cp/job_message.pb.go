// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: internal/pkg/protofiles/executor_cp/job_message.proto

package executor_cp

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasJob    bool              `protobuf:"varint,1,opt,name=hasJob,proto3" json:"hasJob,omitempty"`
	JobID     string            `protobuf:"bytes,2,opt,name=jobID,proto3" json:"jobID,omitempty"`
	ImageName string            `protobuf:"bytes,3,opt,name=imageName,proto3" json:"imageName,omitempty"`
	JobData   map[string]string `protobuf:"bytes,4,rep,name=job_data,json=jobData,proto3" json:"job_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_executor_cp_job_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_executor_cp_job_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetHasJob() bool {
	if x != nil {
		return x.HasJob
	}
	return false
}

func (x *Job) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

func (x *Job) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *Job) GetJobData() map[string]string {
	if x != nil {
		return x.JobData
	}
	return nil
}

type Start struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Start) Reset() {
	*x = Start{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_executor_cp_job_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Start) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Start) ProtoMessage() {}

func (x *Start) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_executor_cp_job_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Start.ProtoReflect.Descriptor instead.
func (*Start) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDescGZIP(), []int{1}
}

func (x *Start) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_internal_pkg_protofiles_executor_cp_job_message_proto protoreflect.FileDescriptor

var file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDesc = []byte{
	0x0a, 0x35, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x5f, 0x63, 0x70, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x4a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x68, 0x61, 0x73, 0x4a, 0x6f, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6a,
	0x6f, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x4a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x6a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x3a, 0x0a, 0x0c, 0x4a, 0x6f, 0x62,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x17, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x2e,
	0x5a, 0x2c, 0x6f, 0x63, 0x74, 0x61, 0x76, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDescOnce sync.Once
	file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDescData = file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDesc
)

func file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDescGZIP() []byte {
	file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDescOnce.Do(func() {
		file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDescData)
	})
	return file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDescData
}

var file_internal_pkg_protofiles_executor_cp_job_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_pkg_protofiles_executor_cp_job_message_proto_goTypes = []interface{}{
	(*Job)(nil),   // 0: Job
	(*Start)(nil), // 1: Start
	nil,           // 2: Job.JobDataEntry
}
var file_internal_pkg_protofiles_executor_cp_job_message_proto_depIdxs = []int32{
	2, // 0: Job.job_data:type_name -> Job.JobDataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_pkg_protofiles_executor_cp_job_message_proto_init() }
func file_internal_pkg_protofiles_executor_cp_job_message_proto_init() {
	if File_internal_pkg_protofiles_executor_cp_job_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pkg_protofiles_executor_cp_job_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_internal_pkg_protofiles_executor_cp_job_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Start); i {
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
			RawDescriptor: file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pkg_protofiles_executor_cp_job_message_proto_goTypes,
		DependencyIndexes: file_internal_pkg_protofiles_executor_cp_job_message_proto_depIdxs,
		MessageInfos:      file_internal_pkg_protofiles_executor_cp_job_message_proto_msgTypes,
	}.Build()
	File_internal_pkg_protofiles_executor_cp_job_message_proto = out.File
	file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDesc = nil
	file_internal_pkg_protofiles_executor_cp_job_message_proto_goTypes = nil
	file_internal_pkg_protofiles_executor_cp_job_message_proto_depIdxs = nil
}
