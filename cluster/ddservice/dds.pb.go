// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dds.proto

/*
Package ddservice is a generated protocol buffer package.

It is generated from these files:
	dds.proto

It has these top-level messages:
	DDSRequest
	DDSResponse
*/
package ddservice

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

// The request message
type DDSRequest struct {
	Ip      string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Time    int64  `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Cid     int32  `protobuf:"varint,3,opt,name=cid" json:"cid,omitempty"`
	Payload string `protobuf:"bytes,4,opt,name=payload" json:"payload,omitempty"`
}

func (m *DDSRequest) Reset()                    { *m = DDSRequest{} }
func (m *DDSRequest) String() string            { return proto.CompactTextString(m) }
func (*DDSRequest) ProtoMessage()               {}
func (*DDSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DDSRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *DDSRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *DDSRequest) GetCid() int32 {
	if m != nil {
		return m.Cid
	}
	return 0
}

func (m *DDSRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

// The response message
type DDSResponse struct {
	Payload string `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *DDSResponse) Reset()                    { *m = DDSResponse{} }
func (m *DDSResponse) String() string            { return proto.CompactTextString(m) }
func (*DDSResponse) ProtoMessage()               {}
func (*DDSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DDSResponse) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterType((*DDSRequest)(nil), "ddservice.DDSRequest")
	proto.RegisterType((*DDSResponse)(nil), "ddservice.DDSResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DDService service

type DDServiceClient interface {
	// Sends a greeting
	Call(ctx context.Context, in *DDSRequest, opts ...grpc.CallOption) (*DDSResponse, error)
}

type dDServiceClient struct {
	cc *grpc.ClientConn
}

func NewDDServiceClient(cc *grpc.ClientConn) DDServiceClient {
	return &dDServiceClient{cc}
}

func (c *dDServiceClient) Call(ctx context.Context, in *DDSRequest, opts ...grpc.CallOption) (*DDSResponse, error) {
	out := new(DDSResponse)
	err := grpc.Invoke(ctx, "/ddservice.DDService/Call", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DDService service

type DDServiceServer interface {
	// Sends a greeting
	Call(context.Context, *DDSRequest) (*DDSResponse, error)
}

func RegisterDDServiceServer(s *grpc.Server, srv DDServiceServer) {
	s.RegisterService(&_DDService_serviceDesc, srv)
}

func _DDService_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DDSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDServiceServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ddservice.DDService/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDServiceServer).Call(ctx, req.(*DDSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DDService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ddservice.DDService",
	HandlerType: (*DDServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _DDService_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dds.proto",
}

func init() { proto.RegisterFile("dds.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbd, 0x4b, 0xc5, 0x30,
	0x14, 0xc5, 0x4d, 0x5b, 0x95, 0x5e, 0x41, 0xe5, 0x82, 0x12, 0x74, 0x29, 0x5d, 0xec, 0x14, 0x44,
	0x07, 0xf7, 0x67, 0x06, 0xc7, 0x47, 0xde, 0xe0, 0xe2, 0x12, 0x9b, 0x8b, 0x06, 0x52, 0x13, 0x93,
	0xfa, 0xf5, 0xdf, 0x4b, 0x52, 0xfc, 0xdc, 0x7e, 0x09, 0xe7, 0x1e, 0x7e, 0x1c, 0x68, 0x8d, 0x49,
	0x22, 0x44, 0x3f, 0x7b, 0xcc, 0x48, 0xf1, 0xd5, 0x8e, 0xd4, 0xdf, 0x01, 0x48, 0xb9, 0x51, 0xf4,
	0xfc, 0x42, 0x69, 0xc6, 0x7d, 0xa8, 0x6c, 0xe0, 0xac, 0x63, 0x43, 0xab, 0x2a, 0x1b, 0x10, 0xa1,
	0x99, 0xed, 0x44, 0xbc, 0xea, 0xd8, 0x50, 0xab, 0xc2, 0x78, 0x08, 0xf5, 0x68, 0x0d, 0xaf, 0x3b,
	0x36, 0x6c, 0xab, 0x8c, 0xc8, 0x61, 0x37, 0xe8, 0x0f, 0xe7, 0xb5, 0xe1, 0x4d, 0x39, 0xfd, 0x7a,
	0xf6, 0x67, 0xb0, 0x57, 0xda, 0x53, 0xf0, 0x4f, 0x89, 0x7e, 0x07, 0xd9, 0x9f, 0xe0, 0x85, 0x84,
	0x56, 0xca, 0xcd, 0xe2, 0x84, 0x57, 0xd0, 0x5c, 0x6b, 0xe7, 0xf0, 0x48, 0x7c, 0x7b, 0x8a, 0x1f,
	0xc9, 0x93, 0xe3, 0xff, 0xdf, 0x4b, 0x7b, 0xbf, 0xb5, 0x3a, 0x87, 0x53, 0xeb, 0xc5, 0x43, 0x0c,
	0xa3, 0xa0, 0x77, 0x3d, 0x05, 0x47, 0x49, 0x3c, 0x92, 0x73, 0xfe, 0xcd, 0x47, 0x67, 0x56, 0x07,
	0x37, 0x99, 0x6f, 0x33, 0xaf, 0xf3, 0x0e, 0x6b, 0x76, 0xbf, 0x53, 0x06, 0xb9, 0xfc, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0xe1, 0x17, 0xd1, 0x88, 0x1d, 0x01, 0x00, 0x00,
}