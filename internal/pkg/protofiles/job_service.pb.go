// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: internal/pkg/protofiles/job_service.proto

package protofiles

import (
	context "context"
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

var File_internal_pkg_protofiles_job_service_proto protoreflect.FileDescriptor

var file_internal_pkg_protofiles_job_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdb, 0x01, 0x0a, 0x0a, 0x4a, 0x6f, 0x62,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x12,
	0x11, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x1a, 0x04, 0x2e, 0x4c, 0x6f, 0x67, 0x12, 0x27, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x49, 0x44, 0x1a, 0x04, 0x2e, 0x4a, 0x6f, 0x62, 0x12, 0x38, 0x0a, 0x11, 0x50,
	0x6f, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x11, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x1a, 0x10, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x07, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x1a, 0x10, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x22, 0x5a, 0x20, 0x6f, 0x63, 0x74, 0x61, 0x76, 0x69,
	0x75, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_internal_pkg_protofiles_job_service_proto_goTypes = []interface{}{
	(*RequestToGetLogs)(nil), // 0: RequestToGetLogs
	(*RequestToExecute)(nil), // 1: RequestToExecute
	(*ExecutorID)(nil),       // 2: ExecutorID
	(*ExecutionContext)(nil), // 3: ExecutionContext
	(*Status)(nil),           // 4: Status
	(*Log)(nil),              // 5: Log
	(*Response)(nil),         // 6: Response
	(*Job)(nil),              // 7: Job
	(*Acknowledgement)(nil),  // 8: Acknowledgement
}
var file_internal_pkg_protofiles_job_service_proto_depIdxs = []int32{
	0, // 0: JobService.Logs:input_type -> RequestToGetLogs
	1, // 1: JobService.Execute:input_type -> RequestToExecute
	2, // 2: JobService.Get:input_type -> ExecutorID
	3, // 3: JobService.PostExecutionData:input_type -> ExecutionContext
	4, // 4: JobService.PostExecutorStatus:input_type -> Status
	5, // 5: JobService.Logs:output_type -> Log
	6, // 6: JobService.Execute:output_type -> Response
	7, // 7: JobService.Get:output_type -> Job
	8, // 8: JobService.PostExecutionData:output_type -> Acknowledgement
	8, // 9: JobService.PostExecutorStatus:output_type -> Acknowledgement
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pkg_protofiles_job_service_proto_init() }
func file_internal_pkg_protofiles_job_service_proto_init() {
	if File_internal_pkg_protofiles_job_service_proto != nil {
		return
	}
	file_internal_pkg_protofiles_job_messages_proto_init()
	file_internal_pkg_protofiles_execution_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_pkg_protofiles_job_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pkg_protofiles_job_service_proto_goTypes,
		DependencyIndexes: file_internal_pkg_protofiles_job_service_proto_depIdxs,
	}.Build()
	File_internal_pkg_protofiles_job_service_proto = out.File
	file_internal_pkg_protofiles_job_service_proto_rawDesc = nil
	file_internal_pkg_protofiles_job_service_proto_goTypes = nil
	file_internal_pkg_protofiles_job_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	Logs(ctx context.Context, in *RequestToGetLogs, opts ...grpc.CallOption) (*Log, error)
	Execute(ctx context.Context, in *RequestToExecute, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *ExecutorID, opts ...grpc.CallOption) (*Job, error)
	PostExecutionData(ctx context.Context, in *ExecutionContext, opts ...grpc.CallOption) (*Acknowledgement, error)
	PostExecutorStatus(ctx context.Context, in *Status, opts ...grpc.CallOption) (*Acknowledgement, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) Logs(ctx context.Context, in *RequestToGetLogs, opts ...grpc.CallOption) (*Log, error) {
	out := new(Log)
	err := c.cc.Invoke(ctx, "/JobService/Logs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) Execute(ctx context.Context, in *RequestToExecute, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/JobService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) Get(ctx context.Context, in *ExecutorID, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/JobService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) PostExecutionData(ctx context.Context, in *ExecutionContext, opts ...grpc.CallOption) (*Acknowledgement, error) {
	out := new(Acknowledgement)
	err := c.cc.Invoke(ctx, "/JobService/PostExecutionData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) PostExecutorStatus(ctx context.Context, in *Status, opts ...grpc.CallOption) (*Acknowledgement, error) {
	out := new(Acknowledgement)
	err := c.cc.Invoke(ctx, "/JobService/PostExecutorStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	Logs(context.Context, *RequestToGetLogs) (*Log, error)
	Execute(context.Context, *RequestToExecute) (*Response, error)
	Get(context.Context, *ExecutorID) (*Job, error)
	PostExecutionData(context.Context, *ExecutionContext) (*Acknowledgement, error)
	PostExecutorStatus(context.Context, *Status) (*Acknowledgement, error)
}

// UnimplementedJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (*UnimplementedJobServiceServer) Logs(context.Context, *RequestToGetLogs) (*Log, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logs not implemented")
}
func (*UnimplementedJobServiceServer) Execute(context.Context, *RequestToExecute) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (*UnimplementedJobServiceServer) Get(context.Context, *ExecutorID) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedJobServiceServer) PostExecutionData(context.Context, *ExecutionContext) (*Acknowledgement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostExecutionData not implemented")
}
func (*UnimplementedJobServiceServer) PostExecutorStatus(context.Context, *Status) (*Acknowledgement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostExecutorStatus not implemented")
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_Logs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestToGetLogs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).Logs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JobService/Logs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).Logs(ctx, req.(*RequestToGetLogs))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestToExecute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JobService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).Execute(ctx, req.(*RequestToExecute))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutorID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JobService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).Get(ctx, req.(*ExecutorID))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_PostExecutionData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutionContext)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).PostExecutionData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JobService/PostExecutionData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).PostExecutionData(ctx, req.(*ExecutionContext))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_PostExecutorStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Status)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).PostExecutorStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JobService/PostExecutorStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).PostExecutorStatus(ctx, req.(*Status))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Logs",
			Handler:    _JobService_Logs_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _JobService_Execute_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _JobService_Get_Handler,
		},
		{
			MethodName: "PostExecutionData",
			Handler:    _JobService_PostExecutionData_Handler,
		},
		{
			MethodName: "PostExecutorStatus",
			Handler:    _JobService_PostExecutorStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/pkg/protofiles/job_service.proto",
}
