// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: microservice.proto

package pkg

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Microservice_UploadFile_FullMethodName        = "/merkle_storage.v1.Microservice/UploadFile"
	Microservice_DownloadFile_FullMethodName      = "/merkle_storage.v1.Microservice/DownloadFile"
	Microservice_VerifyMerkleProof_FullMethodName = "/merkle_storage.v1.Microservice/VerifyMerkleProof"
)

// MicroserviceClient is the client API for Microservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicroserviceClient interface {
	// Endpoint to upload a file to the server
	UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error)
	// Endpoint to download a file from the server and get its Merkle proof
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error)
	// Endpoint to verify the Merkle proof for a file against a client's saved Merkle root
	VerifyMerkleProof(ctx context.Context, in *VerifyMerkleProofRequest, opts ...grpc.CallOption) (*VerifyMerkleProofResponse, error)
}

type microserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroserviceClient(cc grpc.ClientConnInterface) MicroserviceClient {
	return &microserviceClient{cc}
}

func (c *microserviceClient) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error) {
	out := new(UploadFileResponse)
	err := c.cc.Invoke(ctx, Microservice_UploadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error) {
	out := new(DownloadFileResponse)
	err := c.cc.Invoke(ctx, Microservice_DownloadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceClient) VerifyMerkleProof(ctx context.Context, in *VerifyMerkleProofRequest, opts ...grpc.CallOption) (*VerifyMerkleProofResponse, error) {
	out := new(VerifyMerkleProofResponse)
	err := c.cc.Invoke(ctx, Microservice_VerifyMerkleProof_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroserviceServer is the server API for Microservice service.
// All implementations must embed UnimplementedMicroserviceServer
// for forward compatibility
type MicroserviceServer interface {
	// Endpoint to upload a file to the server
	UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error)
	// Endpoint to download a file from the server and get its Merkle proof
	DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error)
	// Endpoint to verify the Merkle proof for a file against a client's saved Merkle root
	VerifyMerkleProof(context.Context, *VerifyMerkleProofRequest) (*VerifyMerkleProofResponse, error)
	mustEmbedUnimplementedMicroserviceServer()
}

// UnimplementedMicroserviceServer must be embedded to have forward compatible implementations.
type UnimplementedMicroserviceServer struct {
}

func (UnimplementedMicroserviceServer) UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedMicroserviceServer) DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedMicroserviceServer) VerifyMerkleProof(context.Context, *VerifyMerkleProofRequest) (*VerifyMerkleProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyMerkleProof not implemented")
}
func (UnimplementedMicroserviceServer) mustEmbedUnimplementedMicroserviceServer() {}

// UnsafeMicroserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroserviceServer will
// result in compilation errors.
type UnsafeMicroserviceServer interface {
	mustEmbedUnimplementedMicroserviceServer()
}

func RegisterMicroserviceServer(s grpc.ServiceRegistrar, srv MicroserviceServer) {
	s.RegisterService(&Microservice_ServiceDesc, srv)
}

func _Microservice_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Microservice_UploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceServer).UploadFile(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microservice_DownloadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceServer).DownloadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Microservice_DownloadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceServer).DownloadFile(ctx, req.(*DownloadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microservice_VerifyMerkleProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyMerkleProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceServer).VerifyMerkleProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Microservice_VerifyMerkleProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceServer).VerifyMerkleProof(ctx, req.(*VerifyMerkleProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Microservice_ServiceDesc is the grpc.ServiceDesc for Microservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Microservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "merkle_storage.v1.Microservice",
	HandlerType: (*MicroserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadFile",
			Handler:    _Microservice_UploadFile_Handler,
		},
		{
			MethodName: "DownloadFile",
			Handler:    _Microservice_DownloadFile_Handler,
		},
		{
			MethodName: "VerifyMerkleProof",
			Handler:    _Microservice_VerifyMerkleProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "microservice.proto",
}
