// Code generated by protoc-gen-go.
// source: mcpserver.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	mcpserver.proto
	processing.proto
	workflow.proto

It has these top-level messages:
	ApproveTransferRequest
	ApproveTransferResponse
	ApproveJobRequest
	ApproveJobResponse
	ChoiceListRequest
	ChoiceListResponse
	Job
	WorkflowGetRequest
	WorkflowGetResponse
	WorkflowData
	WatchedDirectory
	Chain
	Link
*/
package pb

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

type ChoiceListResponse_Job_UnitType int32

const (
	ChoiceListResponse_Job_UNKNOWN  ChoiceListResponse_Job_UnitType = 0
	ChoiceListResponse_Job_TRANSFER ChoiceListResponse_Job_UnitType = 1
	ChoiceListResponse_Job_INGEST   ChoiceListResponse_Job_UnitType = 2
	ChoiceListResponse_Job_DIP      ChoiceListResponse_Job_UnitType = 3
)

var ChoiceListResponse_Job_UnitType_name = map[int32]string{
	0: "UNKNOWN",
	1: "TRANSFER",
	2: "INGEST",
	3: "DIP",
}
var ChoiceListResponse_Job_UnitType_value = map[string]int32{
	"UNKNOWN":  0,
	"TRANSFER": 1,
	"INGEST":   2,
	"DIP":      3,
}

func (x ChoiceListResponse_Job_UnitType) String() string {
	return proto.EnumName(ChoiceListResponse_Job_UnitType_name, int32(x))
}
func (ChoiceListResponse_Job_UnitType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0, 0}
}

type ApproveTransferRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ApproveTransferRequest) Reset()                    { *m = ApproveTransferRequest{} }
func (m *ApproveTransferRequest) String() string            { return proto.CompactTextString(m) }
func (*ApproveTransferRequest) ProtoMessage()               {}
func (*ApproveTransferRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ApproveTransferRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ApproveTransferResponse struct {
	Approved bool `protobuf:"varint,1,opt,name=approved" json:"approved,omitempty"`
}

func (m *ApproveTransferResponse) Reset()                    { *m = ApproveTransferResponse{} }
func (m *ApproveTransferResponse) String() string            { return proto.CompactTextString(m) }
func (*ApproveTransferResponse) ProtoMessage()               {}
func (*ApproveTransferResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ApproveTransferResponse) GetApproved() bool {
	if m != nil {
		return m.Approved
	}
	return false
}

type ApproveJobRequest struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ChoiceId string `protobuf:"bytes,2,opt,name=choiceId" json:"choiceId,omitempty"`
}

func (m *ApproveJobRequest) Reset()                    { *m = ApproveJobRequest{} }
func (m *ApproveJobRequest) String() string            { return proto.CompactTextString(m) }
func (*ApproveJobRequest) ProtoMessage()               {}
func (*ApproveJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ApproveJobRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ApproveJobRequest) GetChoiceId() string {
	if m != nil {
		return m.ChoiceId
	}
	return ""
}

type ApproveJobResponse struct {
	Approved bool `protobuf:"varint,1,opt,name=approved" json:"approved,omitempty"`
}

func (m *ApproveJobResponse) Reset()                    { *m = ApproveJobResponse{} }
func (m *ApproveJobResponse) String() string            { return proto.CompactTextString(m) }
func (*ApproveJobResponse) ProtoMessage()               {}
func (*ApproveJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ApproveJobResponse) GetApproved() bool {
	if m != nil {
		return m.Approved
	}
	return false
}

type ChoiceListRequest struct {
}

func (m *ChoiceListRequest) Reset()                    { *m = ChoiceListRequest{} }
func (m *ChoiceListRequest) String() string            { return proto.CompactTextString(m) }
func (*ChoiceListRequest) ProtoMessage()               {}
func (*ChoiceListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ChoiceListResponse struct {
	Jobs          []*ChoiceListResponse_Job `protobuf:"bytes,1,rep,name=jobs" json:"jobs,omitempty"`
	TransferCount int32                     `protobuf:"varint,2,opt,name=transferCount" json:"transferCount,omitempty"`
	IngestCount   int32                     `protobuf:"varint,3,opt,name=ingestCount" json:"ingestCount,omitempty"`
}

func (m *ChoiceListResponse) Reset()                    { *m = ChoiceListResponse{} }
func (m *ChoiceListResponse) String() string            { return proto.CompactTextString(m) }
func (*ChoiceListResponse) ProtoMessage()               {}
func (*ChoiceListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ChoiceListResponse) GetJobs() []*ChoiceListResponse_Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *ChoiceListResponse) GetTransferCount() int32 {
	if m != nil {
		return m.TransferCount
	}
	return 0
}

func (m *ChoiceListResponse) GetIngestCount() int32 {
	if m != nil {
		return m.IngestCount
	}
	return 0
}

type ChoiceListResponse_Job struct {
	Id       string                           `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	UnitType ChoiceListResponse_Job_UnitType  `protobuf:"varint,2,opt,name=unitType,enum=pb.ChoiceListResponse_Job_UnitType" json:"unitType,omitempty"`
	Choices  []*ChoiceListResponse_Job_Choice `protobuf:"bytes,3,rep,name=choices" json:"choices,omitempty"`
}

func (m *ChoiceListResponse_Job) Reset()                    { *m = ChoiceListResponse_Job{} }
func (m *ChoiceListResponse_Job) String() string            { return proto.CompactTextString(m) }
func (*ChoiceListResponse_Job) ProtoMessage()               {}
func (*ChoiceListResponse_Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *ChoiceListResponse_Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ChoiceListResponse_Job) GetUnitType() ChoiceListResponse_Job_UnitType {
	if m != nil {
		return m.UnitType
	}
	return ChoiceListResponse_Job_UNKNOWN
}

func (m *ChoiceListResponse_Job) GetChoices() []*ChoiceListResponse_Job_Choice {
	if m != nil {
		return m.Choices
	}
	return nil
}

type ChoiceListResponse_Job_Choice struct {
	Value       string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *ChoiceListResponse_Job_Choice) Reset()         { *m = ChoiceListResponse_Job_Choice{} }
func (m *ChoiceListResponse_Job_Choice) String() string { return proto.CompactTextString(m) }
func (*ChoiceListResponse_Job_Choice) ProtoMessage()    {}
func (*ChoiceListResponse_Job_Choice) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0, 0}
}

func (m *ChoiceListResponse_Job_Choice) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ChoiceListResponse_Job_Choice) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*ApproveTransferRequest)(nil), "pb.ApproveTransferRequest")
	proto.RegisterType((*ApproveTransferResponse)(nil), "pb.ApproveTransferResponse")
	proto.RegisterType((*ApproveJobRequest)(nil), "pb.ApproveJobRequest")
	proto.RegisterType((*ApproveJobResponse)(nil), "pb.ApproveJobResponse")
	proto.RegisterType((*ChoiceListRequest)(nil), "pb.ChoiceListRequest")
	proto.RegisterType((*ChoiceListResponse)(nil), "pb.ChoiceListResponse")
	proto.RegisterType((*ChoiceListResponse_Job)(nil), "pb.ChoiceListResponse.Job")
	proto.RegisterType((*ChoiceListResponse_Job_Choice)(nil), "pb.ChoiceListResponse.Job.Choice")
	proto.RegisterEnum("pb.ChoiceListResponse_Job_UnitType", ChoiceListResponse_Job_UnitType_name, ChoiceListResponse_Job_UnitType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MCPServer service

type MCPServerClient interface {
	// ApproveJob proceeds processing with a choice given by the user
	ApproveJob(ctx context.Context, in *ApproveJobRequest, opts ...grpc.CallOption) (*ApproveJobResponse, error)
	// ApproveTransfer starts a transfer awaiting for approval
	ApproveTransfer(ctx context.Context, in *ApproveTransferRequest, opts ...grpc.CallOption) (*ApproveTransferResponse, error)
	// ChoiceList returns the list of job choices awaiting for decision
	ChoiceList(ctx context.Context, in *ChoiceListRequest, opts ...grpc.CallOption) (*ChoiceListResponse, error)
}

type mCPServerClient struct {
	cc *grpc.ClientConn
}

func NewMCPServerClient(cc *grpc.ClientConn) MCPServerClient {
	return &mCPServerClient{cc}
}

func (c *mCPServerClient) ApproveJob(ctx context.Context, in *ApproveJobRequest, opts ...grpc.CallOption) (*ApproveJobResponse, error) {
	out := new(ApproveJobResponse)
	err := grpc.Invoke(ctx, "/pb.MCPServer/ApproveJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mCPServerClient) ApproveTransfer(ctx context.Context, in *ApproveTransferRequest, opts ...grpc.CallOption) (*ApproveTransferResponse, error) {
	out := new(ApproveTransferResponse)
	err := grpc.Invoke(ctx, "/pb.MCPServer/ApproveTransfer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mCPServerClient) ChoiceList(ctx context.Context, in *ChoiceListRequest, opts ...grpc.CallOption) (*ChoiceListResponse, error) {
	out := new(ChoiceListResponse)
	err := grpc.Invoke(ctx, "/pb.MCPServer/ChoiceList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MCPServer service

type MCPServerServer interface {
	// ApproveJob proceeds processing with a choice given by the user
	ApproveJob(context.Context, *ApproveJobRequest) (*ApproveJobResponse, error)
	// ApproveTransfer starts a transfer awaiting for approval
	ApproveTransfer(context.Context, *ApproveTransferRequest) (*ApproveTransferResponse, error)
	// ChoiceList returns the list of job choices awaiting for decision
	ChoiceList(context.Context, *ChoiceListRequest) (*ChoiceListResponse, error)
}

func RegisterMCPServerServer(s *grpc.Server, srv MCPServerServer) {
	s.RegisterService(&_MCPServer_serviceDesc, srv)
}

func _MCPServer_ApproveJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MCPServerServer).ApproveJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MCPServer/ApproveJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MCPServerServer).ApproveJob(ctx, req.(*ApproveJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MCPServer_ApproveTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MCPServerServer).ApproveTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MCPServer/ApproveTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MCPServerServer).ApproveTransfer(ctx, req.(*ApproveTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MCPServer_ChoiceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChoiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MCPServerServer).ChoiceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MCPServer/ChoiceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MCPServerServer).ChoiceList(ctx, req.(*ChoiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MCPServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MCPServer",
	HandlerType: (*MCPServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApproveJob",
			Handler:    _MCPServer_ApproveJob_Handler,
		},
		{
			MethodName: "ApproveTransfer",
			Handler:    _MCPServer_ApproveTransfer_Handler,
		},
		{
			MethodName: "ChoiceList",
			Handler:    _MCPServer_ChoiceList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mcpserver.proto",
}

func init() { proto.RegisterFile("mcpserver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x18, 0xc4, 0x36, 0x4d, 0xdc, 0x2f, 0xd0, 0xa6, 0x1f, 0x50, 0x2c, 0x73, 0x09, 0x86, 0x43, 0x4e,
	0x16, 0x0a, 0xe2, 0x42, 0x0e, 0xa5, 0x0a, 0x05, 0x25, 0x80, 0xa9, 0x36, 0xae, 0x38, 0xc7, 0xf6,
	0x02, 0x8b, 0xc0, 0xbb, 0x78, 0xd7, 0x91, 0x78, 0x0c, 0xde, 0x80, 0x57, 0xe2, 0x8d, 0x90, 0x77,
	0xb7, 0xa9, 0x89, 0x09, 0xea, 0x71, 0xc7, 0x33, 0xf3, 0xcd, 0xf7, 0x63, 0x38, 0xfc, 0x96, 0x0b,
	0x49, 0xab, 0x35, 0xad, 0x62, 0x51, 0x71, 0xc5, 0xd1, 0x15, 0x59, 0x34, 0x86, 0xe3, 0x53, 0x21,
	0x2a, 0xbe, 0xa6, 0x69, 0xb5, 0x2a, 0xe5, 0x47, 0x5a, 0x11, 0xfa, 0xbd, 0xa6, 0x52, 0xe1, 0x01,
	0xb8, 0xac, 0x08, 0x9c, 0x91, 0x33, 0xde, 0x27, 0x2e, 0x2b, 0xa2, 0x67, 0x70, 0xbf, 0xc3, 0x94,
	0x82, 0x97, 0x92, 0x62, 0x08, 0xfe, 0xca, 0x7c, 0x32, 0x02, 0x9f, 0x6c, 0xde, 0xd1, 0x09, 0x1c,
	0x59, 0xd9, 0x82, 0x67, 0x3b, 0xbc, 0x1b, 0x83, 0xfc, 0x33, 0x67, 0x39, 0x9d, 0x17, 0x81, 0xab,
	0xd1, 0xcd, 0x3b, 0x7a, 0x02, 0xd8, 0x36, 0xb8, 0x46, 0xc9, 0x3b, 0x70, 0x34, 0xd3, 0xea, 0xb7,
	0x4c, 0x2a, 0x5b, 0x32, 0xfa, 0xe5, 0x01, 0xb6, 0x51, 0xeb, 0x13, 0xc3, 0xcd, 0x2f, 0x3c, 0x93,
	0x81, 0x33, 0xf2, 0xc6, 0x83, 0x49, 0x18, 0x8b, 0x2c, 0xee, 0xb2, 0xe2, 0xa6, 0xb2, 0xe6, 0xe1,
	0x63, 0xb8, 0xad, 0x6c, 0xfb, 0x33, 0x5e, 0x97, 0x4a, 0xc7, 0xdd, 0x23, 0x7f, 0x83, 0x38, 0x82,
	0x01, 0x2b, 0x3f, 0x51, 0xa9, 0x0c, 0xc7, 0xd3, 0x9c, 0x36, 0x14, 0xfe, 0x74, 0xc1, 0x5b, 0xf0,
	0xac, 0x33, 0x89, 0x13, 0xf0, 0xeb, 0x92, 0xa9, 0xf4, 0x87, 0xa0, 0xda, 0xfa, 0x60, 0xf2, 0x68,
	0x77, 0xa6, 0xf8, 0xc2, 0x52, 0xc9, 0x46, 0x84, 0x53, 0xe8, 0x9b, 0xd1, 0xc9, 0xc0, 0xd3, 0x3d,
	0x3d, 0xfc, 0x8f, 0xde, 0xc0, 0xe4, 0x52, 0x11, 0xbe, 0x80, 0x9e, 0x81, 0xf0, 0x2e, 0xec, 0xad,
	0x57, 0x5f, 0x6b, 0x6a, 0xa3, 0x99, 0x47, 0xd3, 0x57, 0x41, 0x65, 0x5e, 0x31, 0xa1, 0x18, 0x2f,
	0xed, 0xaa, 0xda, 0x50, 0xf4, 0x1c, 0xfc, 0xcb, 0x50, 0x38, 0x80, 0xfe, 0x45, 0xf2, 0x26, 0x79,
	0xff, 0x21, 0x19, 0xde, 0xc0, 0x5b, 0xe0, 0xa7, 0xe4, 0x34, 0x59, 0xbe, 0x3a, 0x23, 0x43, 0x07,
	0x01, 0x7a, 0xf3, 0xe4, 0xf5, 0xd9, 0x32, 0x1d, 0xba, 0xd8, 0x07, 0xef, 0xe5, 0xfc, 0x7c, 0xe8,
	0x4d, 0x7e, 0x3b, 0xb0, 0xff, 0x6e, 0x76, 0xbe, 0xd4, 0x37, 0x8a, 0x53, 0x80, 0xab, 0xbd, 0xe3,
	0xbd, 0xa6, 0x8b, 0xce, 0x21, 0x85, 0xc7, 0xdb, 0xb0, 0x5d, 0xeb, 0x02, 0x0e, 0xb7, 0x8e, 0x15,
	0xc3, 0x16, 0x75, 0xeb, 0xd6, 0xc3, 0x07, 0xff, 0xfc, 0x66, 0xbd, 0xa6, 0x00, 0x57, 0xe3, 0x33,
	0x41, 0x3a, 0xe7, 0x65, 0x82, 0x74, 0xa7, 0x9c, 0xf5, 0xf4, 0xaf, 0xf6, 0xf4, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa5, 0x14, 0xf7, 0xa5, 0x7d, 0x03, 0x00, 0x00,
}
