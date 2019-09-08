// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo-service.proto

/*
Package todo_service is a generated protocol buffer package.

It is generated from these files:
	todo-service.proto

It has these top-level messages:
	Task
	CreateRequest
	CreateResponse
	RetrieveRequest
	RetrieveResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	RetrieveAllRequest
	RetrieveAllResponse
*/
package todo_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Task struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Priority    int32  `protobuf:"varint,4,opt,name=priority" json:"priority,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Task) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

type CreateRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Task  *Task  `protobuf:"bytes,2,opt,name=task" json:"task,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CreateRequest) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type CreateResponse struct {
	Task *Task `protobuf:"bytes,1,opt,name=task" json:"task,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateResponse) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type RetrieveRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *RetrieveRequest) Reset()                    { *m = RetrieveRequest{} }
func (m *RetrieveRequest) String() string            { return proto.CompactTextString(m) }
func (*RetrieveRequest) ProtoMessage()               {}
func (*RetrieveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RetrieveRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RetrieveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RetrieveResponse struct {
	Task *Task `protobuf:"bytes,1,opt,name=task" json:"task,omitempty"`
}

func (m *RetrieveResponse) Reset()                    { *m = RetrieveResponse{} }
func (m *RetrieveResponse) String() string            { return proto.CompactTextString(m) }
func (*RetrieveResponse) ProtoMessage()               {}
func (*RetrieveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RetrieveResponse) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type UpdateRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Task  *Task  `protobuf:"bytes,2,opt,name=task" json:"task,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateRequest) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type UpdateResponse struct {
	Task *Task `protobuf:"bytes,1,opt,name=task" json:"task,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateResponse) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type DeleteRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	Success int64 `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DeleteResponse) GetSuccess() int64 {
	if m != nil {
		return m.Success
	}
	return 0
}

type RetrieveAllRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *RetrieveAllRequest) Reset()                    { *m = RetrieveAllRequest{} }
func (m *RetrieveAllRequest) String() string            { return proto.CompactTextString(m) }
func (*RetrieveAllRequest) ProtoMessage()               {}
func (*RetrieveAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RetrieveAllRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RetrieveAllResponse struct {
	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *RetrieveAllResponse) Reset()                    { *m = RetrieveAllResponse{} }
func (m *RetrieveAllResponse) String() string            { return proto.CompactTextString(m) }
func (*RetrieveAllResponse) ProtoMessage()               {}
func (*RetrieveAllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RetrieveAllResponse) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "Task")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*RetrieveRequest)(nil), "RetrieveRequest")
	proto.RegisterType((*RetrieveResponse)(nil), "RetrieveResponse")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
	proto.RegisterType((*RetrieveAllRequest)(nil), "RetrieveAllRequest")
	proto.RegisterType((*RetrieveAllResponse)(nil), "RetrieveAllResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TodoService service

type TodoServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*RetrieveResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	RetrieveAll(ctx context.Context, in *RetrieveAllRequest, opts ...grpc.CallOption) (*RetrieveAllResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/TodoService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*RetrieveResponse, error) {
	out := new(RetrieveResponse)
	err := grpc.Invoke(ctx, "/TodoService/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/TodoService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/TodoService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) RetrieveAll(ctx context.Context, in *RetrieveAllRequest, opts ...grpc.CallOption) (*RetrieveAllResponse, error) {
	out := new(RetrieveAllResponse)
	err := grpc.Invoke(ctx, "/TodoService/RetrieveAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TodoService service

type TodoServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Retrieve(context.Context, *RetrieveRequest) (*RetrieveResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	RetrieveAll(context.Context, *RetrieveAllRequest) (*RetrieveAllResponse, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_RetrieveAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).RetrieveAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/RetrieveAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).RetrieveAll(ctx, req.(*RetrieveAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TodoService_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _TodoService_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TodoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TodoService_Delete_Handler,
		},
		{
			MethodName: "RetrieveAll",
			Handler:    _TodoService_RetrieveAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo-service.proto",
}

func init() { proto.RegisterFile("todo-service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0x4d, 0x4b, 0xdb, 0x8f, 0x6f, 0x1a, 0x5a, 0x1c, 0x38, 0xd4, 0x7a, 0x21, 0x3d, 0x01, 0x86,
	0x35, 0xc1, 0x18, 0x3d, 0x62, 0xf4, 0x17, 0x54, 0xfc, 0x01, 0xb5, 0x9d, 0xc4, 0x0d, 0xd8, 0xad,
	0xdd, 0x85, 0xc4, 0x9f, 0xee, 0xcd, 0xd0, 0x52, 0xa0, 0xc5, 0x48, 0x8c, 0xb7, 0xce, 0x74, 0xe6,
	0xbd, 0x7d, 0xef, 0xed, 0x02, 0x2a, 0x91, 0x88, 0x89, 0xa4, 0x7c, 0xcd, 0x63, 0x62, 0x59, 0x2e,
	0x94, 0x08, 0x5e, 0xc1, 0x98, 0x47, 0x72, 0x81, 0x0e, 0xe8, 0x3c, 0xf1, 0xb4, 0x81, 0x36, 0xfc,
	0x1f, 0xea, 0x3c, 0x41, 0x04, 0x23, 0x8d, 0xde, 0xc8, 0xd3, 0x8b, 0x4e, 0xf1, 0x8d, 0x03, 0xb0,
	0x13, 0x92, 0x71, 0xce, 0x33, 0xc5, 0x45, 0xea, 0xb5, 0x8a, 0x5f, 0x87, 0x2d, 0xf4, 0xa1, 0x9d,
	0xe5, 0x5c, 0xe4, 0x5c, 0x7d, 0x78, 0xc6, 0x40, 0x1b, 0x9a, 0xe1, 0xae, 0x0e, 0x66, 0xd0, 0x79,
	0xc8, 0x29, 0x52, 0x14, 0xd2, 0xfb, 0x8a, 0xa4, 0xc2, 0x3e, 0x98, 0x4a, 0x2c, 0x28, 0xdd, 0xb2,
	0x96, 0x05, 0x9e, 0x83, 0xa1, 0x22, 0xb9, 0x28, 0x88, 0xed, 0xa9, 0xc9, 0x36, 0xa7, 0x0b, 0x8b,
	0x56, 0x70, 0x09, 0x4e, 0x85, 0x20, 0x33, 0x91, 0x4a, 0xda, 0x0d, 0x6b, 0xc7, 0xc3, 0xb7, 0xe0,
	0x86, 0xa4, 0x72, 0x4e, 0xeb, 0x13, 0x84, 0xa5, 0x72, 0xbd, 0x52, 0x1e, 0x4c, 0xa0, 0xbb, 0x5f,
	0x3c, 0xcd, 0x33, 0x83, 0xce, 0x73, 0x96, 0xfc, 0x51, 0x56, 0x85, 0x70, 0x9a, 0xee, 0x06, 0x3a,
	0x8f, 0xb4, 0x24, 0xf5, 0x4b, 0x51, 0x63, 0x70, 0xaa, 0xb5, 0x2d, 0x87, 0x07, 0xff, 0xe4, 0x2a,
	0x8e, 0x49, 0xca, 0x62, 0xb3, 0x15, 0x56, 0x65, 0x30, 0x06, 0xac, 0x0c, 0xb8, 0x5f, 0x2e, 0x7f,
	0xe4, 0x09, 0xa6, 0xd0, 0xab, 0xcd, 0x6e, 0xc1, 0x2f, 0xc0, 0xdc, 0x9c, 0x76, 0x03, 0xdd, 0xda,
	0x2b, 0x28, 0x7b, 0xd3, 0x4f, 0x0d, 0xec, 0xb9, 0x48, 0xc4, 0x53, 0x79, 0x11, 0x71, 0x04, 0x56,
	0x19, 0x2b, 0x3a, 0xac, 0x76, 0x43, 0x7c, 0x97, 0x35, 0xf2, 0xbe, 0x82, 0x76, 0x45, 0x87, 0x5d,
	0xd6, 0xc8, 0xd7, 0x3f, 0x63, 0x47, 0xc1, 0x8d, 0xc0, 0x2a, 0xbd, 0x45, 0x87, 0xd5, 0x62, 0xf2,
	0x5d, 0xd6, 0x30, 0x7d, 0x04, 0x56, 0x69, 0x11, 0x3a, 0xac, 0x66, 0xb1, 0xef, 0xb2, 0x86, 0x77,
	0x77, 0x60, 0x1f, 0xa8, 0xc6, 0x1e, 0x3b, 0xf6, 0xcb, 0xef, 0xb3, 0x6f, 0x8c, 0x79, 0xb1, 0x8a,
	0x57, 0x77, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x80, 0x27, 0x55, 0x8b, 0x03, 0x00, 0x00,
}
