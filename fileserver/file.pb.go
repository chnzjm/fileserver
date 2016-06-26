// Code generated by protoc-gen-go.
// source: file.proto
// DO NOT EDIT!

/*
Package fileserver is a generated protocol buffer package.

It is generated from these files:
	file.proto

It has these top-level messages:
	FileDescriptor
	FileContent
*/
package fileserver

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

type FileDescriptor struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
}

func (m *FileDescriptor) Reset()                    { *m = FileDescriptor{} }
func (m *FileDescriptor) String() string            { return proto.CompactTextString(m) }
func (*FileDescriptor) ProtoMessage()               {}
func (*FileDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FileContent struct {
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *FileContent) Reset()                    { *m = FileContent{} }
func (m *FileContent) String() string            { return proto.CompactTextString(m) }
func (*FileContent) ProtoMessage()               {}
func (*FileContent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*FileDescriptor)(nil), "fileserver.FileDescriptor")
	proto.RegisterType((*FileContent)(nil), "fileserver.FileContent")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for File service

type FileClient interface {
	GetFile(ctx context.Context, in *FileDescriptor, opts ...grpc.CallOption) (File_GetFileClient, error)
}

type fileClient struct {
	cc *grpc.ClientConn
}

func NewFileClient(cc *grpc.ClientConn) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) GetFile(ctx context.Context, in *FileDescriptor, opts ...grpc.CallOption) (File_GetFileClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_File_serviceDesc.Streams[0], c.cc, "/fileserver.File/GetFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileGetFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type File_GetFileClient interface {
	Recv() (*FileContent, error)
	grpc.ClientStream
}

type fileGetFileClient struct {
	grpc.ClientStream
}

func (x *fileGetFileClient) Recv() (*FileContent, error) {
	m := new(FileContent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for File service

type FileServer interface {
	GetFile(*FileDescriptor, File_GetFileServer) error
}

func RegisterFileServer(s *grpc.Server, srv FileServer) {
	s.RegisterService(&_File_serviceDesc, srv)
}

func _File_GetFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileDescriptor)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServer).GetFile(m, &fileGetFileServer{stream})
}

type File_GetFileServer interface {
	Send(*FileContent) error
	grpc.ServerStream
}

type fileGetFileServer struct {
	grpc.ServerStream
}

func (x *fileGetFileServer) Send(m *FileContent) error {
	return x.ServerStream.SendMsg(m)
}

var _File_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fileserver.File",
	HandlerType: (*FileServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFile",
			Handler:       _File_GetFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("file.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcb, 0xcc, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x02, 0xb3, 0x8b, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0x94,
	0x74, 0xb8, 0xf8, 0xdc, 0x80, 0x3c, 0x97, 0xd4, 0xe2, 0xe4, 0xa2, 0xcc, 0x82, 0x92, 0xfc, 0x22,
	0x21, 0x29, 0x2e, 0x0e, 0x90, 0x7c, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x9c, 0xaf, 0xa4, 0xce, 0xc5, 0x0d, 0x52, 0xed, 0x9c, 0x9f, 0x57, 0x92, 0x9a, 0x57, 0x22,
	0x24, 0xc1, 0xc5, 0x9e, 0x0c, 0x61, 0x82, 0x55, 0xf2, 0x04, 0xc1, 0xb8, 0x46, 0x5e, 0x5c, 0x2c,
	0x20, 0x85, 0x42, 0x4e, 0x5c, 0xec, 0xee, 0xa9, 0x25, 0x60, 0xa6, 0x94, 0x1e, 0xc2, 0x5a, 0x3d,
	0x54, 0x3b, 0xa5, 0xc4, 0xd1, 0xe5, 0xa0, 0x36, 0x28, 0x31, 0x18, 0x30, 0x26, 0xb1, 0x81, 0x5d,
	0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x33, 0x5e, 0x70, 0x3f, 0xc3, 0x00, 0x00, 0x00,
}