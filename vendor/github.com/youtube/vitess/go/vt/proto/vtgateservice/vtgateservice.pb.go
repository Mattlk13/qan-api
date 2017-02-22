// Code generated by protoc-gen-go.
// source: vtgateservice.proto
// DO NOT EDIT!

/*
Package vtgateservice is a generated protocol buffer package.

It is generated from these files:
	vtgateservice.proto

It has these top-level messages:
*/
package vtgateservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import vtgate "github.com/youtube/vitess/go/vt/proto/vtgate"

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

// Client API for Vitess service

type VitessClient interface {
	// Execute tries to route the query to the right shard.
	// It depends on the query and bind variables to provide enough
	// information in conjonction with the vindexes to route the query.
	// API group: v3 API (alpha)
	Execute(ctx context.Context, in *vtgate.ExecuteRequest, opts ...grpc.CallOption) (*vtgate.ExecuteResponse, error)
	// ExecuteShards executes the query on the specified shards.
	// API group: Custom Sharding
	ExecuteShards(ctx context.Context, in *vtgate.ExecuteShardsRequest, opts ...grpc.CallOption) (*vtgate.ExecuteShardsResponse, error)
	// ExecuteKeyspaceIds executes the query based on the specified keyspace ids.
	// API group: Range-based Sharding
	ExecuteKeyspaceIds(ctx context.Context, in *vtgate.ExecuteKeyspaceIdsRequest, opts ...grpc.CallOption) (*vtgate.ExecuteKeyspaceIdsResponse, error)
	// ExecuteKeyRanges executes the query based on the specified key ranges.
	// API group: Range-based Sharding
	ExecuteKeyRanges(ctx context.Context, in *vtgate.ExecuteKeyRangesRequest, opts ...grpc.CallOption) (*vtgate.ExecuteKeyRangesResponse, error)
	// ExecuteEntityIds executes the query based on the specified external id to keyspace id map.
	// API group: Range-based Sharding
	ExecuteEntityIds(ctx context.Context, in *vtgate.ExecuteEntityIdsRequest, opts ...grpc.CallOption) (*vtgate.ExecuteEntityIdsResponse, error)
	// ExecuteBatchShards executes the list of queries on the specified shards.
	// API group: Custom Sharding
	ExecuteBatchShards(ctx context.Context, in *vtgate.ExecuteBatchShardsRequest, opts ...grpc.CallOption) (*vtgate.ExecuteBatchShardsResponse, error)
	// ExecuteBatchKeyspaceIds executes the list of queries based on the specified keyspace ids.
	// API group: Range-based Sharding
	ExecuteBatchKeyspaceIds(ctx context.Context, in *vtgate.ExecuteBatchKeyspaceIdsRequest, opts ...grpc.CallOption) (*vtgate.ExecuteBatchKeyspaceIdsResponse, error)
	// StreamExecute executes a streaming query based on shards.
	// It depends on the query and bind variables to provide enough
	// information in conjonction with the vindexes to route the query.
	// Use this method if the query returns a large number of rows.
	// API group: v3 API (alpha)
	StreamExecute(ctx context.Context, in *vtgate.StreamExecuteRequest, opts ...grpc.CallOption) (Vitess_StreamExecuteClient, error)
	// StreamExecuteShards executes a streaming query based on shards.
	// Use this method if the query returns a large number of rows.
	// API group: Custom Sharding
	StreamExecuteShards(ctx context.Context, in *vtgate.StreamExecuteShardsRequest, opts ...grpc.CallOption) (Vitess_StreamExecuteShardsClient, error)
	// StreamExecuteKeyspaceIds executes a streaming query based on keyspace ids.
	// Use this method if the query returns a large number of rows.
	// API group: Range-based Sharding
	StreamExecuteKeyspaceIds(ctx context.Context, in *vtgate.StreamExecuteKeyspaceIdsRequest, opts ...grpc.CallOption) (Vitess_StreamExecuteKeyspaceIdsClient, error)
	// StreamExecuteKeyRanges executes a streaming query based on key ranges.
	// Use this method if the query returns a large number of rows.
	// API group: Range-based Sharding
	StreamExecuteKeyRanges(ctx context.Context, in *vtgate.StreamExecuteKeyRangesRequest, opts ...grpc.CallOption) (Vitess_StreamExecuteKeyRangesClient, error)
	// Begin a transaction.
	// API group: Transactions
	Begin(ctx context.Context, in *vtgate.BeginRequest, opts ...grpc.CallOption) (*vtgate.BeginResponse, error)
	// Commit a transaction.
	// API group: Transactions
	Commit(ctx context.Context, in *vtgate.CommitRequest, opts ...grpc.CallOption) (*vtgate.CommitResponse, error)
	// Rollback a transaction.
	// API group: Transactions
	Rollback(ctx context.Context, in *vtgate.RollbackRequest, opts ...grpc.CallOption) (*vtgate.RollbackResponse, error)
	// Split a query into non-overlapping sub queries
	// API group: Map Reduce
	SplitQuery(ctx context.Context, in *vtgate.SplitQueryRequest, opts ...grpc.CallOption) (*vtgate.SplitQueryResponse, error)
	// GetSrvKeyspace returns a SrvKeyspace object (as seen by this vtgate).
	// This method is provided as a convenient way for clients to take a
	// look at the sharding configuration for a Keyspace. Looking at the
	// sharding information should not be used for routing queries (as the
	// information may change, use the Execute calls for that).
	// It is convenient for monitoring applications for instance, or if
	// using custom sharding.
	// API group: Topology
	GetSrvKeyspace(ctx context.Context, in *vtgate.GetSrvKeyspaceRequest, opts ...grpc.CallOption) (*vtgate.GetSrvKeyspaceResponse, error)
}

type vitessClient struct {
	cc *grpc.ClientConn
}

func NewVitessClient(cc *grpc.ClientConn) VitessClient {
	return &vitessClient{cc}
}

func (c *vitessClient) Execute(ctx context.Context, in *vtgate.ExecuteRequest, opts ...grpc.CallOption) (*vtgate.ExecuteResponse, error) {
	out := new(vtgate.ExecuteResponse)
	err := grpc.Invoke(ctx, "/vtgateservice.Vitess/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vitessClient) ExecuteShards(ctx context.Context, in *vtgate.ExecuteShardsRequest, opts ...grpc.CallOption) (*vtgate.ExecuteShardsResponse, error) {
	out := new(vtgate.ExecuteShardsResponse)
	err := grpc.Invoke(ctx, "/vtgateservice.Vitess/ExecuteShards", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vitessClient) ExecuteKeyspaceIds(ctx context.Context, in *vtgate.ExecuteKeyspaceIdsRequest, opts ...grpc.CallOption) (*vtgate.ExecuteKeyspaceIdsResponse, error) {
	out := new(vtgate.ExecuteKeyspaceIdsResponse)
	err := grpc.Invoke(ctx, "/vtgateservice.Vitess/ExecuteKeyspaceIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vitessClient) ExecuteKeyRanges(ctx context.Context, in *vtgate.ExecuteKeyRangesRequest, opts ...grpc.CallOption) (*vtgate.ExecuteKeyRangesResponse, error) {
	out := new(vtgate.ExecuteKeyRangesResponse)
	err := grpc.Invoke(ctx, "/vtgateservice.Vitess/ExecuteKeyRanges", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vitessClient) ExecuteEntityIds(ctx context.Context, in *vtgate.ExecuteEntityIdsRequest, opts ...grpc.CallOption) (*vtgate.ExecuteEntityIdsResponse, error) {
	out := new(vtgate.ExecuteEntityIdsResponse)
	err := grpc.Invoke(ctx, "/vtgateservice.Vitess/ExecuteEntityIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vitessClient) ExecuteBatchShards(ctx context.Context, in *vtgate.ExecuteBatchShardsRequest, opts ...grpc.CallOption) (*vtgate.ExecuteBatchShardsResponse, error) {
	out := new(vtgate.ExecuteBatchShardsResponse)
	err := grpc.Invoke(ctx, "/vtgateservice.Vitess/ExecuteBatchShards", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vitessClient) ExecuteBatchKeyspaceIds(ctx context.Context, in *vtgate.ExecuteBatchKeyspaceIdsRequest, opts ...grpc.CallOption) (*vtgate.ExecuteBatchKeyspaceIdsResponse, error) {
	out := new(vtgate.ExecuteBatchKeyspaceIdsResponse)
	err := grpc.Invoke(ctx, "/vtgateservice.Vitess/ExecuteBatchKeyspaceIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vitessClient) StreamExecute(ctx context.Context, in *vtgate.StreamExecuteRequest, opts ...grpc.CallOption) (Vitess_StreamExecuteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Vitess_serviceDesc.Streams[0], c.cc, "/vtgateservice.Vitess/StreamExecute", opts...)
	if err != nil {
		return nil, err
	}
	x := &vitessStreamExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Vitess_StreamExecuteClient interface {
	Recv() (*vtgate.StreamExecuteResponse, error)
	grpc.ClientStream
}

type vitessStreamExecuteClient struct {
	grpc.ClientStream
}

func (x *vitessStreamExecuteClient) Recv() (*vtgate.StreamExecuteResponse, error) {
	m := new(vtgate.StreamExecuteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vitessClient) StreamExecuteShards(ctx context.Context, in *vtgate.StreamExecuteShardsRequest, opts ...grpc.CallOption) (Vitess_StreamExecuteShardsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Vitess_serviceDesc.Streams[1], c.cc, "/vtgateservice.Vitess/StreamExecuteShards", opts...)
	if err != nil {
		return nil, err
	}
	x := &vitessStreamExecuteShardsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Vitess_StreamExecuteShardsClient interface {
	Recv() (*vtgate.StreamExecuteShardsResponse, error)
	grpc.ClientStream
}

type vitessStreamExecuteShardsClient struct {
	grpc.ClientStream
}

func (x *vitessStreamExecuteShardsClient) Recv() (*vtgate.StreamExecuteShardsResponse, error) {
	m := new(vtgate.StreamExecuteShardsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vitessClient) StreamExecuteKeyspaceIds(ctx context.Context, in *vtgate.StreamExecuteKeyspaceIdsRequest, opts ...grpc.CallOption) (Vitess_StreamExecuteKeyspaceIdsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Vitess_serviceDesc.Streams[2], c.cc, "/vtgateservice.Vitess/StreamExecuteKeyspaceIds", opts...)
	if err != nil {
		return nil, err
	}
	x := &vitessStreamExecuteKeyspaceIdsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Vitess_StreamExecuteKeyspaceIdsClient interface {
	Recv() (*vtgate.StreamExecuteKeyspaceIdsResponse, error)
	grpc.ClientStream
}

type vitessStreamExecuteKeyspaceIdsClient struct {
	grpc.ClientStream
}

func (x *vitessStreamExecuteKeyspaceIdsClient) Recv() (*vtgate.StreamExecuteKeyspaceIdsResponse, error) {
	m := new(vtgate.StreamExecuteKeyspaceIdsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vitessClient) StreamExecuteKeyRanges(ctx context.Context, in *vtgate.StreamExecuteKeyRangesRequest, opts ...grpc.CallOption) (Vitess_StreamExecuteKeyRangesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Vitess_serviceDesc.Streams[3], c.cc, "/vtgateservice.Vitess/StreamExecuteKeyRanges", opts...)
	if err != nil {
		return nil, err
	}
	x := &vitessStreamExecuteKeyRangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Vitess_StreamExecuteKeyRangesClient interface {
	Recv() (*vtgate.StreamExecuteKeyRangesResponse, error)
	grpc.ClientStream
}

type vitessStreamExecuteKeyRangesClient struct {
	grpc.ClientStream
}

func (x *vitessStreamExecuteKeyRangesClient) Recv() (*vtgate.StreamExecuteKeyRangesResponse, error) {
	m := new(vtgate.StreamExecuteKeyRangesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vitessClient) Begin(ctx context.Context, in *vtgate.BeginRequest, opts ...grpc.CallOption) (*vtgate.BeginResponse, error) {
	out := new(vtgate.BeginResponse)
	err := grpc.Invoke(ctx, "/vtgateservice.Vitess/Begin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vitessClient) Commit(ctx context.Context, in *vtgate.CommitRequest, opts ...grpc.CallOption) (*vtgate.CommitResponse, error) {
	out := new(vtgate.CommitResponse)
	err := grpc.Invoke(ctx, "/vtgateservice.Vitess/Commit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vitessClient) Rollback(ctx context.Context, in *vtgate.RollbackRequest, opts ...grpc.CallOption) (*vtgate.RollbackResponse, error) {
	out := new(vtgate.RollbackResponse)
	err := grpc.Invoke(ctx, "/vtgateservice.Vitess/Rollback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vitessClient) SplitQuery(ctx context.Context, in *vtgate.SplitQueryRequest, opts ...grpc.CallOption) (*vtgate.SplitQueryResponse, error) {
	out := new(vtgate.SplitQueryResponse)
	err := grpc.Invoke(ctx, "/vtgateservice.Vitess/SplitQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vitessClient) GetSrvKeyspace(ctx context.Context, in *vtgate.GetSrvKeyspaceRequest, opts ...grpc.CallOption) (*vtgate.GetSrvKeyspaceResponse, error) {
	out := new(vtgate.GetSrvKeyspaceResponse)
	err := grpc.Invoke(ctx, "/vtgateservice.Vitess/GetSrvKeyspace", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Vitess service

type VitessServer interface {
	// Execute tries to route the query to the right shard.
	// It depends on the query and bind variables to provide enough
	// information in conjonction with the vindexes to route the query.
	// API group: v3 API (alpha)
	Execute(context.Context, *vtgate.ExecuteRequest) (*vtgate.ExecuteResponse, error)
	// ExecuteShards executes the query on the specified shards.
	// API group: Custom Sharding
	ExecuteShards(context.Context, *vtgate.ExecuteShardsRequest) (*vtgate.ExecuteShardsResponse, error)
	// ExecuteKeyspaceIds executes the query based on the specified keyspace ids.
	// API group: Range-based Sharding
	ExecuteKeyspaceIds(context.Context, *vtgate.ExecuteKeyspaceIdsRequest) (*vtgate.ExecuteKeyspaceIdsResponse, error)
	// ExecuteKeyRanges executes the query based on the specified key ranges.
	// API group: Range-based Sharding
	ExecuteKeyRanges(context.Context, *vtgate.ExecuteKeyRangesRequest) (*vtgate.ExecuteKeyRangesResponse, error)
	// ExecuteEntityIds executes the query based on the specified external id to keyspace id map.
	// API group: Range-based Sharding
	ExecuteEntityIds(context.Context, *vtgate.ExecuteEntityIdsRequest) (*vtgate.ExecuteEntityIdsResponse, error)
	// ExecuteBatchShards executes the list of queries on the specified shards.
	// API group: Custom Sharding
	ExecuteBatchShards(context.Context, *vtgate.ExecuteBatchShardsRequest) (*vtgate.ExecuteBatchShardsResponse, error)
	// ExecuteBatchKeyspaceIds executes the list of queries based on the specified keyspace ids.
	// API group: Range-based Sharding
	ExecuteBatchKeyspaceIds(context.Context, *vtgate.ExecuteBatchKeyspaceIdsRequest) (*vtgate.ExecuteBatchKeyspaceIdsResponse, error)
	// StreamExecute executes a streaming query based on shards.
	// It depends on the query and bind variables to provide enough
	// information in conjonction with the vindexes to route the query.
	// Use this method if the query returns a large number of rows.
	// API group: v3 API (alpha)
	StreamExecute(*vtgate.StreamExecuteRequest, Vitess_StreamExecuteServer) error
	// StreamExecuteShards executes a streaming query based on shards.
	// Use this method if the query returns a large number of rows.
	// API group: Custom Sharding
	StreamExecuteShards(*vtgate.StreamExecuteShardsRequest, Vitess_StreamExecuteShardsServer) error
	// StreamExecuteKeyspaceIds executes a streaming query based on keyspace ids.
	// Use this method if the query returns a large number of rows.
	// API group: Range-based Sharding
	StreamExecuteKeyspaceIds(*vtgate.StreamExecuteKeyspaceIdsRequest, Vitess_StreamExecuteKeyspaceIdsServer) error
	// StreamExecuteKeyRanges executes a streaming query based on key ranges.
	// Use this method if the query returns a large number of rows.
	// API group: Range-based Sharding
	StreamExecuteKeyRanges(*vtgate.StreamExecuteKeyRangesRequest, Vitess_StreamExecuteKeyRangesServer) error
	// Begin a transaction.
	// API group: Transactions
	Begin(context.Context, *vtgate.BeginRequest) (*vtgate.BeginResponse, error)
	// Commit a transaction.
	// API group: Transactions
	Commit(context.Context, *vtgate.CommitRequest) (*vtgate.CommitResponse, error)
	// Rollback a transaction.
	// API group: Transactions
	Rollback(context.Context, *vtgate.RollbackRequest) (*vtgate.RollbackResponse, error)
	// Split a query into non-overlapping sub queries
	// API group: Map Reduce
	SplitQuery(context.Context, *vtgate.SplitQueryRequest) (*vtgate.SplitQueryResponse, error)
	// GetSrvKeyspace returns a SrvKeyspace object (as seen by this vtgate).
	// This method is provided as a convenient way for clients to take a
	// look at the sharding configuration for a Keyspace. Looking at the
	// sharding information should not be used for routing queries (as the
	// information may change, use the Execute calls for that).
	// It is convenient for monitoring applications for instance, or if
	// using custom sharding.
	// API group: Topology
	GetSrvKeyspace(context.Context, *vtgate.GetSrvKeyspaceRequest) (*vtgate.GetSrvKeyspaceResponse, error)
}

func RegisterVitessServer(s *grpc.Server, srv VitessServer) {
	s.RegisterService(&_Vitess_serviceDesc, srv)
}

func _Vitess_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(vtgate.ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitessServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vtgateservice.Vitess/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitessServer).Execute(ctx, req.(*vtgate.ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vitess_ExecuteShards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(vtgate.ExecuteShardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitessServer).ExecuteShards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vtgateservice.Vitess/ExecuteShards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitessServer).ExecuteShards(ctx, req.(*vtgate.ExecuteShardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vitess_ExecuteKeyspaceIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(vtgate.ExecuteKeyspaceIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitessServer).ExecuteKeyspaceIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vtgateservice.Vitess/ExecuteKeyspaceIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitessServer).ExecuteKeyspaceIds(ctx, req.(*vtgate.ExecuteKeyspaceIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vitess_ExecuteKeyRanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(vtgate.ExecuteKeyRangesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitessServer).ExecuteKeyRanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vtgateservice.Vitess/ExecuteKeyRanges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitessServer).ExecuteKeyRanges(ctx, req.(*vtgate.ExecuteKeyRangesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vitess_ExecuteEntityIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(vtgate.ExecuteEntityIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitessServer).ExecuteEntityIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vtgateservice.Vitess/ExecuteEntityIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitessServer).ExecuteEntityIds(ctx, req.(*vtgate.ExecuteEntityIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vitess_ExecuteBatchShards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(vtgate.ExecuteBatchShardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitessServer).ExecuteBatchShards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vtgateservice.Vitess/ExecuteBatchShards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitessServer).ExecuteBatchShards(ctx, req.(*vtgate.ExecuteBatchShardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vitess_ExecuteBatchKeyspaceIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(vtgate.ExecuteBatchKeyspaceIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitessServer).ExecuteBatchKeyspaceIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vtgateservice.Vitess/ExecuteBatchKeyspaceIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitessServer).ExecuteBatchKeyspaceIds(ctx, req.(*vtgate.ExecuteBatchKeyspaceIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vitess_StreamExecute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(vtgate.StreamExecuteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VitessServer).StreamExecute(m, &vitessStreamExecuteServer{stream})
}

type Vitess_StreamExecuteServer interface {
	Send(*vtgate.StreamExecuteResponse) error
	grpc.ServerStream
}

type vitessStreamExecuteServer struct {
	grpc.ServerStream
}

func (x *vitessStreamExecuteServer) Send(m *vtgate.StreamExecuteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Vitess_StreamExecuteShards_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(vtgate.StreamExecuteShardsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VitessServer).StreamExecuteShards(m, &vitessStreamExecuteShardsServer{stream})
}

type Vitess_StreamExecuteShardsServer interface {
	Send(*vtgate.StreamExecuteShardsResponse) error
	grpc.ServerStream
}

type vitessStreamExecuteShardsServer struct {
	grpc.ServerStream
}

func (x *vitessStreamExecuteShardsServer) Send(m *vtgate.StreamExecuteShardsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Vitess_StreamExecuteKeyspaceIds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(vtgate.StreamExecuteKeyspaceIdsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VitessServer).StreamExecuteKeyspaceIds(m, &vitessStreamExecuteKeyspaceIdsServer{stream})
}

type Vitess_StreamExecuteKeyspaceIdsServer interface {
	Send(*vtgate.StreamExecuteKeyspaceIdsResponse) error
	grpc.ServerStream
}

type vitessStreamExecuteKeyspaceIdsServer struct {
	grpc.ServerStream
}

func (x *vitessStreamExecuteKeyspaceIdsServer) Send(m *vtgate.StreamExecuteKeyspaceIdsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Vitess_StreamExecuteKeyRanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(vtgate.StreamExecuteKeyRangesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VitessServer).StreamExecuteKeyRanges(m, &vitessStreamExecuteKeyRangesServer{stream})
}

type Vitess_StreamExecuteKeyRangesServer interface {
	Send(*vtgate.StreamExecuteKeyRangesResponse) error
	grpc.ServerStream
}

type vitessStreamExecuteKeyRangesServer struct {
	grpc.ServerStream
}

func (x *vitessStreamExecuteKeyRangesServer) Send(m *vtgate.StreamExecuteKeyRangesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Vitess_Begin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(vtgate.BeginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitessServer).Begin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vtgateservice.Vitess/Begin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitessServer).Begin(ctx, req.(*vtgate.BeginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vitess_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(vtgate.CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitessServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vtgateservice.Vitess/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitessServer).Commit(ctx, req.(*vtgate.CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vitess_Rollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(vtgate.RollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitessServer).Rollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vtgateservice.Vitess/Rollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitessServer).Rollback(ctx, req.(*vtgate.RollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vitess_SplitQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(vtgate.SplitQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitessServer).SplitQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vtgateservice.Vitess/SplitQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitessServer).SplitQuery(ctx, req.(*vtgate.SplitQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vitess_GetSrvKeyspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(vtgate.GetSrvKeyspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitessServer).GetSrvKeyspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vtgateservice.Vitess/GetSrvKeyspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitessServer).GetSrvKeyspace(ctx, req.(*vtgate.GetSrvKeyspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Vitess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vtgateservice.Vitess",
	HandlerType: (*VitessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Vitess_Execute_Handler,
		},
		{
			MethodName: "ExecuteShards",
			Handler:    _Vitess_ExecuteShards_Handler,
		},
		{
			MethodName: "ExecuteKeyspaceIds",
			Handler:    _Vitess_ExecuteKeyspaceIds_Handler,
		},
		{
			MethodName: "ExecuteKeyRanges",
			Handler:    _Vitess_ExecuteKeyRanges_Handler,
		},
		{
			MethodName: "ExecuteEntityIds",
			Handler:    _Vitess_ExecuteEntityIds_Handler,
		},
		{
			MethodName: "ExecuteBatchShards",
			Handler:    _Vitess_ExecuteBatchShards_Handler,
		},
		{
			MethodName: "ExecuteBatchKeyspaceIds",
			Handler:    _Vitess_ExecuteBatchKeyspaceIds_Handler,
		},
		{
			MethodName: "Begin",
			Handler:    _Vitess_Begin_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _Vitess_Commit_Handler,
		},
		{
			MethodName: "Rollback",
			Handler:    _Vitess_Rollback_Handler,
		},
		{
			MethodName: "SplitQuery",
			Handler:    _Vitess_SplitQuery_Handler,
		},
		{
			MethodName: "GetSrvKeyspace",
			Handler:    _Vitess_GetSrvKeyspace_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamExecute",
			Handler:       _Vitess_StreamExecute_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamExecuteShards",
			Handler:       _Vitess_StreamExecuteShards_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamExecuteKeyspaceIds",
			Handler:       _Vitess_StreamExecuteKeyspaceIds_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamExecuteKeyRanges",
			Handler:       _Vitess_StreamExecuteKeyRanges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("vtgateservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x94, 0x5f, 0x6b, 0xdb, 0x30,
	0x14, 0xc5, 0xb7, 0x87, 0x65, 0xe3, 0xb2, 0x8c, 0xa1, 0x6c, 0x49, 0x16, 0x96, 0x65, 0xcb, 0x58,
	0xb2, 0x27, 0x33, 0x36, 0x18, 0x0c, 0x06, 0x83, 0x8c, 0x50, 0x4a, 0xa1, 0x34, 0x31, 0xb4, 0x4f,
	0x7d, 0xb0, 0x9d, 0x8b, 0x63, 0xe2, 0x7f, 0xb1, 0x64, 0x53, 0x7f, 0xd3, 0x7e, 0x9c, 0xd2, 0xd8,
	0x52, 0x25, 0x59, 0x4e, 0xde, 0xaa, 0x73, 0xce, 0xfd, 0xa9, 0x39, 0xbe, 0x08, 0x7a, 0x05, 0xf3,
	0x1d, 0x86, 0x14, 0xb3, 0x22, 0xf0, 0xd0, 0x4a, 0xb3, 0x84, 0x25, 0xa4, 0xab, 0x88, 0xa3, 0xd7,
	0xd5, 0xb1, 0x32, 0x7f, 0xde, 0x03, 0x74, 0xae, 0x03, 0x86, 0x94, 0x92, 0xbf, 0xf0, 0x72, 0x79,
	0x87, 0x5e, 0xce, 0x90, 0xf4, 0xad, 0x3a, 0x54, 0x0b, 0x6b, 0xdc, 0xe7, 0x48, 0xd9, 0x68, 0xd0,
	0xd0, 0x69, 0x9a, 0xc4, 0x14, 0xa7, 0xcf, 0xc8, 0x25, 0x74, 0x6b, 0xd1, 0xde, 0x3a, 0xd9, 0x86,
	0x92, 0x8f, 0x5a, 0xb6, 0x92, 0x39, 0x69, 0xdc, 0xe2, 0x0a, 0xde, 0x2d, 0x90, 0xda, 0xba, 0xc0,
	0x92, 0xa6, 0x8e, 0x87, 0xe7, 0x1b, 0x4a, 0xbe, 0x68, 0x63, 0x92, 0xc7, 0xc9, 0xd3, 0x63, 0x11,
	0x81, 0xbf, 0x81, 0xb7, 0x4f, 0xfe, 0xda, 0x89, 0x7d, 0xa4, 0x64, 0xd2, 0x9c, 0xac, 0x1c, 0x8e,
	0xfe, 0xdc, 0x1e, 0x30, 0x80, 0x97, 0x31, 0x0b, 0x58, 0xf9, 0xf8, 0x5f, 0xeb, 0x60, 0xe1, 0xb4,
	0x81, 0xa5, 0x80, 0xa1, 0x90, 0x85, 0xc3, 0xbc, 0x6d, 0xdd, 0xb2, 0x5e, 0x88, 0xe4, 0xb5, 0x15,
	0xa2, 0x44, 0x04, 0x3e, 0x84, 0x81, 0xec, 0xcb, 0xa5, 0xcf, 0x4c, 0x00, 0x43, 0xf3, 0xf3, 0x93,
	0x39, 0x71, 0xdb, 0x15, 0x74, 0x6d, 0x96, 0xa1, 0x13, 0xf1, 0x8d, 0x13, 0xdb, 0xa2, 0xc8, 0x8d,
	0x6d, 0xd1, 0x5c, 0xce, 0xfb, 0xf1, 0x9c, 0xb8, 0xd0, 0x53, 0xcc, 0xba, 0x9f, 0xa9, 0x71, 0x52,
	0x2d, 0xe8, 0xeb, 0xd1, 0x8c, 0x74, 0xc7, 0x1e, 0x86, 0x4a, 0x44, 0x2e, 0x69, 0x6e, 0x84, 0x18,
	0x5a, 0xfa, 0x7e, 0x3a, 0x28, 0x5d, 0xb9, 0x83, 0xbe, 0x9e, 0xab, 0xb7, 0xf5, 0x5b, 0x1b, 0x47,
	0xdd, 0xd9, 0xd9, 0xa9, 0x98, 0x74, 0xd9, 0x6f, 0x78, 0xb1, 0x40, 0x3f, 0x88, 0xc9, 0x3b, 0x3e,
	0x74, 0x38, 0x72, 0xd4, 0x7b, 0x4d, 0x15, 0x5f, 0xf3, 0x0f, 0x74, 0xfe, 0x27, 0x51, 0x14, 0x30,
	0x22, 0x22, 0xd5, 0x99, 0x4f, 0xf6, 0x75, 0x59, 0x8c, 0xfe, 0x83, 0x57, 0xeb, 0x24, 0x0c, 0x5d,
	0xc7, 0xdb, 0x11, 0xf1, 0xba, 0x70, 0x85, 0x8f, 0x0f, 0x9b, 0x86, 0x00, 0x2c, 0x01, 0xec, 0x34,
	0x0c, 0xd8, 0x2a, 0xc7, 0xac, 0x24, 0x1f, 0xc4, 0xaf, 0x15, 0x1a, 0x87, 0x8c, 0x4c, 0x96, 0xc0,
	0xac, 0xe0, 0xcd, 0x19, 0x32, 0x3b, 0x2b, 0xf8, 0x87, 0x20, 0x62, 0xe7, 0x54, 0x9d, 0xe3, 0x3e,
	0xb5, 0xd9, 0x1c, 0xb9, 0x98, 0xc0, 0xd8, 0x4b, 0x22, 0xab, 0x4c, 0x72, 0x96, 0xbb, 0x68, 0x15,
	0x87, 0x57, 0xb6, 0x7a, 0x76, 0x2d, 0x3f, 0x4b, 0x3d, 0xb7, 0x73, 0xf8, 0xfb, 0xd7, 0x43, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x36, 0x12, 0x3e, 0x98, 0xb6, 0x05, 0x00, 0x00,
}
