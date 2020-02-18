// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/adm_service/administrator.proto

package adm_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	adm_dto "github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common/adm_dto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("grpc/adm_service/administrator.proto", fileDescriptor_b95a594534a3daa7)
}

var fileDescriptor_b95a594534a3daa7 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x3d, 0xab, 0xc2, 0x30,
	0x18, 0x46, 0xef, 0x72, 0xef, 0xd0, 0xbb, 0x15, 0xa7, 0x4e, 0x22, 0xea, 0xd6, 0x14, 0xf4, 0x17,
	0x28, 0x1d, 0x9d, 0x74, 0x11, 0x41, 0x4a, 0x9a, 0x84, 0x36, 0xd0, 0xe4, 0x89, 0xc9, 0x5b, 0x7f,
	0xbf, 0x58, 0x3f, 0x88, 0x50, 0xdc, 0xc2, 0xc9, 0x39, 0xc3, 0xf3, 0x26, 0xf3, 0xc6, 0x3b, 0x51,
	0x70, 0x69, 0xaa, 0xa0, 0xfc, 0x55, 0x0b, 0x75, 0x7f, 0x6b, 0xab, 0x03, 0x79, 0x4e, 0xf0, 0xcc,
	0x79, 0x10, 0xd2, 0xff, 0x48, 0xc8, 0xa6, 0xef, 0x44, 0x12, 0xc6, 0xf4, 0x95, 0x4b, 0x26, 0x9b,
	0x18, 0x1f, 0x1e, 0x65, 0x7a, 0x4c, 0x7e, 0x77, 0x68, 0xb4, 0x4d, 0x17, 0xec, 0x99, 0xb3, 0x0f,
	0x6f, 0xf8, 0xdc, 0xab, 0x4b, 0xaf, 0x02, 0x95, 0x84, 0x6c, 0xf9, 0x55, 0x0b, 0x0e, 0x36, 0xa8,
	0x92, 0x30, 0xfb, 0xd9, 0x56, 0xa7, 0x73, 0xa3, 0xa9, 0xed, 0x6b, 0x26, 0x60, 0x8a, 0x4e, 0x19,
	0xd8, 0x5c, 0x74, 0xe8, 0x65, 0xfe, 0x9a, 0x15, 0x33, 0xc9, 0x43, 0x5b, 0x83, 0x7b, 0x39, 0x4e,
	0x73, 0x01, 0x63, 0x60, 0xe3, 0xab, 0xd4, 0x7f, 0xc3, 0xb2, 0xf5, 0x2d, 0x00, 0x00, 0xff, 0xff,
	0xd8, 0x25, 0x82, 0x23, 0x30, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdministratorServiceClient is the client API for AdministratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdministratorServiceClient interface {
	Login(ctx context.Context, in *adm_dto.AdministratorLoginRequestDto, opts ...grpc.CallOption) (*adm_dto.AdministratorLoginResponseDto, error)
}

type administratorServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdministratorServiceClient(cc *grpc.ClientConn) AdministratorServiceClient {
	return &administratorServiceClient{cc}
}

func (c *administratorServiceClient) Login(ctx context.Context, in *adm_dto.AdministratorLoginRequestDto, opts ...grpc.CallOption) (*adm_dto.AdministratorLoginResponseDto, error) {
	out := new(adm_dto.AdministratorLoginResponseDto)
	err := c.cc.Invoke(ctx, "/adm_service.AdministratorService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdministratorServiceServer is the server API for AdministratorService service.
type AdministratorServiceServer interface {
	Login(context.Context, *adm_dto.AdministratorLoginRequestDto) (*adm_dto.AdministratorLoginResponseDto, error)
}

// UnimplementedAdministratorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdministratorServiceServer struct {
}

func (*UnimplementedAdministratorServiceServer) Login(ctx context.Context, req *adm_dto.AdministratorLoginRequestDto) (*adm_dto.AdministratorLoginResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterAdministratorServiceServer(s *grpc.Server, srv AdministratorServiceServer) {
	s.RegisterService(&_AdministratorService_serviceDesc, srv)
}

func _AdministratorService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(adm_dto.AdministratorLoginRequestDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adm_service.AdministratorService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServiceServer).Login(ctx, req.(*adm_dto.AdministratorLoginRequestDto))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdministratorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adm_service.AdministratorService",
	HandlerType: (*AdministratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AdministratorService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/adm_service/administrator.proto",
}