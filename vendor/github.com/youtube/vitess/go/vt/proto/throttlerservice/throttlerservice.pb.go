// Code generated by protoc-gen-go.
// source: throttlerservice.proto
// DO NOT EDIT!

/*
Package throttlerservice is a generated protocol buffer package.

It is generated from these files:
	throttlerservice.proto

It has these top-level messages:
*/
package throttlerservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import throttlerdata "github.com/youtube/vitess/go/vt/proto/throttlerdata"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Throttler service

type ThrottlerClient interface {
	// MaxRates returns the current max rate for each throttler of the process.
	MaxRates(ctx context.Context, in *throttlerdata.MaxRatesRequest, opts ...grpc.CallOption) (*throttlerdata.MaxRatesResponse, error)
	// SetMaxRate allows to change the current max rate for all throttlers
	// of the process.
	SetMaxRate(ctx context.Context, in *throttlerdata.SetMaxRateRequest, opts ...grpc.CallOption) (*throttlerdata.SetMaxRateResponse, error)
}

type throttlerClient struct {
	cc *grpc.ClientConn
}

func NewThrottlerClient(cc *grpc.ClientConn) ThrottlerClient {
	return &throttlerClient{cc}
}

func (c *throttlerClient) MaxRates(ctx context.Context, in *throttlerdata.MaxRatesRequest, opts ...grpc.CallOption) (*throttlerdata.MaxRatesResponse, error) {
	out := new(throttlerdata.MaxRatesResponse)
	err := grpc.Invoke(ctx, "/throttlerservice.Throttler/MaxRates", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *throttlerClient) SetMaxRate(ctx context.Context, in *throttlerdata.SetMaxRateRequest, opts ...grpc.CallOption) (*throttlerdata.SetMaxRateResponse, error) {
	out := new(throttlerdata.SetMaxRateResponse)
	err := grpc.Invoke(ctx, "/throttlerservice.Throttler/SetMaxRate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Throttler service

type ThrottlerServer interface {
	// MaxRates returns the current max rate for each throttler of the process.
	MaxRates(context.Context, *throttlerdata.MaxRatesRequest) (*throttlerdata.MaxRatesResponse, error)
	// SetMaxRate allows to change the current max rate for all throttlers
	// of the process.
	SetMaxRate(context.Context, *throttlerdata.SetMaxRateRequest) (*throttlerdata.SetMaxRateResponse, error)
}

func RegisterThrottlerServer(s *grpc.Server, srv ThrottlerServer) {
	s.RegisterService(&_Throttler_serviceDesc, srv)
}

func _Throttler_MaxRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(throttlerdata.MaxRatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThrottlerServer).MaxRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/throttlerservice.Throttler/MaxRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThrottlerServer).MaxRates(ctx, req.(*throttlerdata.MaxRatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Throttler_SetMaxRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(throttlerdata.SetMaxRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThrottlerServer).SetMaxRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/throttlerservice.Throttler/SetMaxRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThrottlerServer).SetMaxRate(ctx, req.(*throttlerdata.SetMaxRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Throttler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "throttlerservice.Throttler",
	HandlerType: (*ThrottlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MaxRates",
			Handler:    _Throttler_MaxRates_Handler,
		},
		{
			MethodName: "SetMaxRate",
			Handler:    _Throttler_SetMaxRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("throttlerservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xc9, 0x28, 0xca,
	0x2f, 0x29, 0xc9, 0x49, 0x2d, 0x2a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x40, 0x17, 0x97, 0x12, 0x86, 0x8b, 0xa4, 0x24, 0x96, 0x24, 0x42, 0x94,
	0x19, 0xad, 0x67, 0xe4, 0xe2, 0x0c, 0x81, 0x89, 0x0b, 0xf9, 0x72, 0x71, 0xf8, 0x26, 0x56, 0x04,
	0x25, 0x96, 0xa4, 0x16, 0x0b, 0xc9, 0xe9, 0xa1, 0xaa, 0x87, 0x49, 0x04, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x48, 0xc9, 0xe3, 0x94, 0x2f, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x55, 0x62, 0x10, 0x0a,
	0xe6, 0xe2, 0x0a, 0x4e, 0x2d, 0x81, 0x4a, 0x08, 0x29, 0xa0, 0x69, 0x40, 0x48, 0xc1, 0x8c, 0x54,
	0xc4, 0xa3, 0x02, 0x66, 0x68, 0x12, 0x1b, 0xd8, 0xe1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf0, 0x88, 0x13, 0xa8, 0xf9, 0x00, 0x00, 0x00,
}
