// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/networkserver.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	go_thethings_network_lorawan_stack_v3_pkg_types "go.thethings.network/lorawan-stack/v3/pkg/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Response of GenerateDevAddr.
type GenerateDevAddrResponse struct {
	DevAddr              *go_thethings_network_lorawan_stack_v3_pkg_types.DevAddr `protobuf:"bytes,1,opt,name=dev_addr,json=devAddr,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.DevAddr" json:"dev_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                 `json:"-"`
	XXX_sizecache        int32                                                    `json:"-"`
}

func (m *GenerateDevAddrResponse) Reset()      { *m = GenerateDevAddrResponse{} }
func (*GenerateDevAddrResponse) ProtoMessage() {}
func (*GenerateDevAddrResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77e7504ad1081b8, []int{0}
}
func (m *GenerateDevAddrResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenerateDevAddrResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenerateDevAddrResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenerateDevAddrResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateDevAddrResponse.Merge(m, src)
}
func (m *GenerateDevAddrResponse) XXX_Size() int {
	return m.Size()
}
func (m *GenerateDevAddrResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateDevAddrResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateDevAddrResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenerateDevAddrResponse)(nil), "ttn.lorawan.v3.GenerateDevAddrResponse")
	golang_proto.RegisterType((*GenerateDevAddrResponse)(nil), "ttn.lorawan.v3.GenerateDevAddrResponse")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/networkserver.proto", fileDescriptor_c77e7504ad1081b8)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/networkserver.proto", fileDescriptor_c77e7504ad1081b8)
}

var fileDescriptor_c77e7504ad1081b8 = []byte{
	// 772 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0x41, 0x48, 0x2b, 0x47,
	0x18, 0xc7, 0x77, 0xa2, 0xd8, 0x32, 0x88, 0xa5, 0x53, 0xb1, 0x6d, 0xda, 0x4e, 0x25, 0xb6, 0xb4,
	0x58, 0xdc, 0x85, 0x78, 0x28, 0xd8, 0x53, 0x24, 0x69, 0x2c, 0x54, 0xa9, 0xd1, 0x16, 0xea, 0x25,
	0x4c, 0x76, 0x3f, 0x37, 0x4b, 0x36, 0x33, 0xeb, 0xce, 0x24, 0x69, 0x28, 0x82, 0xf4, 0x50, 0x3c,
	0x0a, 0xa5, 0xd0, 0x63, 0x2f, 0x05, 0x8f, 0xd2, 0x4b, 0x3d, 0x94, 0xe2, 0xd1, 0x53, 0x11, 0xde,
	0x45, 0xde, 0xe1, 0x61, 0x76, 0xdf, 0xc1, 0xa3, 0x47, 0x8f, 0x8f, 0x6c, 0x12, 0x63, 0xb2, 0xe6,
	0xa1, 0xef, 0x3d, 0xde, 0x6d, 0x26, 0xdf, 0x7f, 0xfe, 0xf3, 0x9b, 0xff, 0xf7, 0x2d, 0xc1, 0x9f,
	0xba, 0xc2, 0x67, 0x0d, 0xc6, 0x17, 0xa4, 0x62, 0x66, 0xc5, 0x60, 0x9e, 0x63, 0x70, 0x50, 0x0d,
	0xe1, 0x57, 0x24, 0xf8, 0x75, 0xf0, 0x75, 0xcf, 0x17, 0x4a, 0x90, 0x29, 0xa5, 0xb8, 0xde, 0x95,
	0xea, 0xf5, 0xc5, 0xe4, 0x82, 0xed, 0xa8, 0x72, 0xad, 0xa4, 0x9b, 0xa2, 0x6a, 0xd8, 0xc2, 0x16,
	0x46, 0x24, 0x2b, 0xd5, 0xb6, 0xa3, 0x5d, 0xb4, 0x89, 0x56, 0x9d, 0xe3, 0xc9, 0x0f, 0x6d, 0x21,
	0x6c, 0x17, 0x22, 0x7b, 0xc6, 0xb9, 0x50, 0x4c, 0x39, 0x82, 0xcb, 0x6e, 0xf5, 0x83, 0x6e, 0xf5,
	0xc6, 0x03, 0xaa, 0x9e, 0x6a, 0x76, 0x8b, 0xa9, 0x38, 0x20, 0x70, 0xab, 0x68, 0x41, 0xdd, 0x31,
	0xa1, 0xab, 0x99, 0x8b, 0x6b, 0x1c, 0x0b, 0xb8, 0x72, 0xb6, 0x1d, 0xf0, 0x7b, 0xb7, 0xcc, 0xc6,
	0x45, 0x55, 0x90, 0x92, 0xd9, 0xd0, 0x55, 0xa4, 0x76, 0xf0, 0xbb, 0x79, 0xe0, 0xe0, 0x33, 0x05,
	0x59, 0xa8, 0x67, 0x2c, 0xcb, 0x2f, 0x80, 0xf4, 0x04, 0x97, 0x40, 0x7e, 0xc0, 0x6f, 0x5a, 0x50,
	0x2f, 0x32, 0xcb, 0xf2, 0xdf, 0x43, 0xb3, 0xe8, 0xf3, 0xc9, 0xe5, 0xaf, 0x1e, 0x3f, 0xf9, 0xf8,
	0x4b, 0x5b, 0xe8, 0xaa, 0x0c, 0xaa, 0xec, 0x70, 0x5b, 0xea, 0xdd, 0xdc, 0x8c, 0xc1, 0x7b, 0xea,
	0x8b, 0x86, 0x57, 0xb1, 0x0d, 0xd5, 0xf4, 0x40, 0xea, 0x3d, 0xdb, 0x37, 0xac, 0xce, 0x22, 0xcd,
	0x71, 0x62, 0x4d, 0x92, 0x32, 0x7e, 0x6b, 0xe8, 0x62, 0x32, 0xa3, 0x77, 0x42, 0xd1, 0x7b, 0xa1,
	0xe8, 0xb9, 0x76, 0x28, 0xc9, 0xcf, 0xf4, 0xc1, 0x4e, 0xe8, 0x23, 0x88, 0x53, 0xd3, 0xbf, 0x3c,
	0x7a, 0xfa, 0x5b, 0x62, 0x8a, 0x4c, 0x1a, 0x5c, 0x1a, 0x3d, 0xf6, 0xf4, 0x41, 0x02, 0x8f, 0x67,
	0xe4, 0x9a, 0x24, 0x9b, 0x78, 0x3a, 0x2b, 0x1a, 0xdc, 0x75, 0x78, 0x65, 0xbd, 0x06, 0x35, 0x28,
	0x80, 0xe7, 0x32, 0x13, 0xc8, 0x27, 0xc3, 0xfe, 0x43, 0xaa, 0x9d, 0x1a, 0x48, 0x95, 0x1c, 0x41,
	0x47, 0xd6, 0xf1, 0xdb, 0x03, 0xfa, 0xef, 0x6a, 0xb2, 0xfc, 0x92, 0x96, 0xc5, 0x21, 0xcb, 0x6f,
	0x1d, 0xa9, 0xe2, 0x96, 0x39, 0x6e, 0x65, 0xa3, 0x89, 0xf8, 0xa6, 0xdf, 0xf7, 0x64, 0x4c, 0x95,
	0xf1, 0x3c, 0xd7, 0x31, 0xa3, 0xd9, 0xeb, 0x79, 0xca, 0xf4, 0x21, 0xc2, 0xe3, 0xf9, 0x76, 0x24,
	0x39, 0x3c, 0xb9, 0xc2, 0xb8, 0xe5, 0xc2, 0xf7, 0x5e, 0xbb, 0x42, 0x3e, 0x1a, 0x3e, 0xde, 0xf9,
	0x7d, 0xb5, 0x33, 0x34, 0x23, 0x81, 0x7f, 0xc4, 0x33, 0x05, 0xf0, 0x84, 0xaf, 0x36, 0x7f, 0xca,
	0x98, 0x15, 0x2e, 0x1a, 0x2e, 0x58, 0x76, 0x15, 0xb8, 0x22, 0xf1, 0xde, 0x31, 0x05, 0x0d, 0xd6,
	0x1c, 0x16, 0x8e, 0xb2, 0x4e, 0xff, 0x3b, 0x81, 0xdf, 0x59, 0x93, 0x37, 0x6f, 0x2d, 0x80, 0xed,
	0x48, 0xe5, 0x37, 0xc9, 0xdf, 0x08, 0x8f, 0xe5, 0x41, 0x91, 0xb9, 0xf8, 0x70, 0xa8, 0x5b, 0xea,
	0x4e, 0xd0, 0xef, 0x8f, 0xcc, 0x2e, 0x55, 0x89, 0x66, 0x06, 0x88, 0xd9, 0x9e, 0x19, 0xd6, 0x0f,
	0x4b, 0x1a, 0x3f, 0xf7, 0xbf, 0xb9, 0xa2, 0x63, 0x49, 0xfd, 0x56, 0xf1, 0x8e, 0xfd, 0xae, 0xd1,
	0x91, 0xc6, 0xcf, 0xdd, 0x2c, 0x77, 0xc9, 0xaf, 0x09, 0x3c, 0xb6, 0x71, 0x17, 0xf4, 0xc6, 0xc3,
	0xa0, 0xff, 0x43, 0x11, 0xf5, 0x3f, 0x28, 0xf9, 0x5c, 0x6c, 0xfd, 0x05, 0xb1, 0xf5, 0x41, 0xec,
	0x25, 0x34, 0xbf, 0xb5, 0x9a, 0x5a, 0x79, 0x55, 0x37, 0x2d, 0xa1, 0x79, 0xf2, 0x3f, 0xc2, 0xd3,
	0x05, 0x90, 0xa0, 0xbe, 0x66, 0xa6, 0x12, 0x7e, 0x33, 0x0b, 0xdb, 0xac, 0xe6, 0x2a, 0x49, 0xbe,
	0x18, 0x7e, 0x74, 0xa4, 0xca, 0x70, 0xeb, 0x81, 0x6d, 0xe5, 0x51, 0x40, 0xe5, 0xf4, 0xeb, 0x68,
	0x6b, 0xfb, 0x41, 0xbf, 0x23, 0x3c, 0x91, 0x05, 0x17, 0x14, 0xdc, 0xf3, 0x43, 0x1d, 0x31, 0xef,
	0xa9, 0xd5, 0x08, 0x3c, 0x3f, 0x9f, 0x8b, 0x83, 0xdf, 0x9b, 0xb4, 0x8f, 0xb6, 0xfc, 0x17, 0x3a,
	0x6d, 0x51, 0x74, 0xd6, 0xa2, 0xe8, 0xbc, 0x45, 0xb5, 0x8b, 0x16, 0xd5, 0x2e, 0x5b, 0x54, 0xbb,
	0x6a, 0x51, 0xed, 0xba, 0x45, 0xd1, 0x5e, 0x40, 0xd1, 0x7e, 0x40, 0xb5, 0xc3, 0x80, 0xa2, 0xa3,
	0x80, 0x6a, 0xc7, 0x01, 0xd5, 0x4e, 0x02, 0xaa, 0x9d, 0x06, 0x14, 0x9d, 0x05, 0x14, 0x9d, 0x07,
	0x54, 0xbb, 0x08, 0x28, 0xba, 0x0c, 0xa8, 0x76, 0x15, 0x50, 0x74, 0x1d, 0x50, 0x6d, 0x2f, 0xa4,
	0xda, 0x7e, 0x48, 0xd1, 0x41, 0x48, 0xb5, 0x3f, 0x42, 0x8a, 0xfe, 0x0c, 0xa9, 0x76, 0x18, 0x52,
	0xed, 0x28, 0xa4, 0xe8, 0x38, 0xa4, 0xe8, 0x24, 0xa4, 0x68, 0xcb, 0x78, 0xc0, 0xbf, 0x84, 0xe2,
	0x5e, 0xa9, 0x34, 0x11, 0xc5, 0xb0, 0xf8, 0x2c, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xe4, 0xa0, 0x43,
	0x9c, 0x07, 0x00, 0x00,
}

func (this *GenerateDevAddrResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenerateDevAddrResponse)
	if !ok {
		that2, ok := that.(GenerateDevAddrResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.DevAddr == nil {
		if this.DevAddr != nil {
			return false
		}
	} else if !this.DevAddr.Equal(*that1.DevAddr) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NsClient is the client API for Ns service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NsClient interface {
	// GenerateDevAddr requests a device address assignment from the Network Server.
	GenerateDevAddr(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*GenerateDevAddrResponse, error)
}

type nsClient struct {
	cc *grpc.ClientConn
}

func NewNsClient(cc *grpc.ClientConn) NsClient {
	return &nsClient{cc}
}

func (c *nsClient) GenerateDevAddr(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*GenerateDevAddrResponse, error) {
	out := new(GenerateDevAddrResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.Ns/GenerateDevAddr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NsServer is the server API for Ns service.
type NsServer interface {
	// GenerateDevAddr requests a device address assignment from the Network Server.
	GenerateDevAddr(context.Context, *types.Empty) (*GenerateDevAddrResponse, error)
}

// UnimplementedNsServer can be embedded to have forward compatible implementations.
type UnimplementedNsServer struct {
}

func (*UnimplementedNsServer) GenerateDevAddr(ctx context.Context, req *types.Empty) (*GenerateDevAddrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateDevAddr not implemented")
}

func RegisterNsServer(s *grpc.Server, srv NsServer) {
	s.RegisterService(&_Ns_serviceDesc, srv)
}

func _Ns_GenerateDevAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsServer).GenerateDevAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.Ns/GenerateDevAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsServer).GenerateDevAddr(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ns_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.Ns",
	HandlerType: (*NsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateDevAddr",
			Handler:    _Ns_GenerateDevAddr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/networkserver.proto",
}

// AsNsClient is the client API for AsNs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AsNsClient interface {
	// Replace the entire downlink queue with the specified messages.
	// This can also be used to empty the queue by specifying no messages.
	// Note that this will trigger an immediate downlink if a downlink slot is available.
	DownlinkQueueReplace(ctx context.Context, in *DownlinkQueueRequest, opts ...grpc.CallOption) (*types.Empty, error)
	// Push downlink messages to the end of the downlink queue.
	// Note that this will trigger an immediate downlink if a downlink slot is available.
	DownlinkQueuePush(ctx context.Context, in *DownlinkQueueRequest, opts ...grpc.CallOption) (*types.Empty, error)
	// List the items currently in the downlink queue.
	DownlinkQueueList(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*ApplicationDownlinks, error)
}

type asNsClient struct {
	cc *grpc.ClientConn
}

func NewAsNsClient(cc *grpc.ClientConn) AsNsClient {
	return &asNsClient{cc}
}

func (c *asNsClient) DownlinkQueueReplace(ctx context.Context, in *DownlinkQueueRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.AsNs/DownlinkQueueReplace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asNsClient) DownlinkQueuePush(ctx context.Context, in *DownlinkQueueRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.AsNs/DownlinkQueuePush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asNsClient) DownlinkQueueList(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*ApplicationDownlinks, error) {
	out := new(ApplicationDownlinks)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.AsNs/DownlinkQueueList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AsNsServer is the server API for AsNs service.
type AsNsServer interface {
	// Replace the entire downlink queue with the specified messages.
	// This can also be used to empty the queue by specifying no messages.
	// Note that this will trigger an immediate downlink if a downlink slot is available.
	DownlinkQueueReplace(context.Context, *DownlinkQueueRequest) (*types.Empty, error)
	// Push downlink messages to the end of the downlink queue.
	// Note that this will trigger an immediate downlink if a downlink slot is available.
	DownlinkQueuePush(context.Context, *DownlinkQueueRequest) (*types.Empty, error)
	// List the items currently in the downlink queue.
	DownlinkQueueList(context.Context, *EndDeviceIdentifiers) (*ApplicationDownlinks, error)
}

// UnimplementedAsNsServer can be embedded to have forward compatible implementations.
type UnimplementedAsNsServer struct {
}

func (*UnimplementedAsNsServer) DownlinkQueueReplace(ctx context.Context, req *DownlinkQueueRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownlinkQueueReplace not implemented")
}
func (*UnimplementedAsNsServer) DownlinkQueuePush(ctx context.Context, req *DownlinkQueueRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownlinkQueuePush not implemented")
}
func (*UnimplementedAsNsServer) DownlinkQueueList(ctx context.Context, req *EndDeviceIdentifiers) (*ApplicationDownlinks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownlinkQueueList not implemented")
}

func RegisterAsNsServer(s *grpc.Server, srv AsNsServer) {
	s.RegisterService(&_AsNs_serviceDesc, srv)
}

func _AsNs_DownlinkQueueReplace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownlinkQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsNsServer).DownlinkQueueReplace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.AsNs/DownlinkQueueReplace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsNsServer).DownlinkQueueReplace(ctx, req.(*DownlinkQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsNs_DownlinkQueuePush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownlinkQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsNsServer).DownlinkQueuePush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.AsNs/DownlinkQueuePush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsNsServer).DownlinkQueuePush(ctx, req.(*DownlinkQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsNs_DownlinkQueueList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDeviceIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsNsServer).DownlinkQueueList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.AsNs/DownlinkQueueList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsNsServer).DownlinkQueueList(ctx, req.(*EndDeviceIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _AsNs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.AsNs",
	HandlerType: (*AsNsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownlinkQueueReplace",
			Handler:    _AsNs_DownlinkQueueReplace_Handler,
		},
		{
			MethodName: "DownlinkQueuePush",
			Handler:    _AsNs_DownlinkQueuePush_Handler,
		},
		{
			MethodName: "DownlinkQueueList",
			Handler:    _AsNs_DownlinkQueueList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/networkserver.proto",
}

// GsNsClient is the client API for GsNs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GsNsClient interface {
	// Called by the Gateway Server when an uplink message arrives.
	HandleUplink(ctx context.Context, in *UplinkMessage, opts ...grpc.CallOption) (*types.Empty, error)
	// Called by the Gateway Server when a Tx acknowledgment arrives.
	ReportTxAcknowledgment(ctx context.Context, in *GatewayTxAcknowledgment, opts ...grpc.CallOption) (*types.Empty, error)
}

type gsNsClient struct {
	cc *grpc.ClientConn
}

func NewGsNsClient(cc *grpc.ClientConn) GsNsClient {
	return &gsNsClient{cc}
}

func (c *gsNsClient) HandleUplink(ctx context.Context, in *UplinkMessage, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GsNs/HandleUplink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gsNsClient) ReportTxAcknowledgment(ctx context.Context, in *GatewayTxAcknowledgment, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GsNs/ReportTxAcknowledgment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GsNsServer is the server API for GsNs service.
type GsNsServer interface {
	// Called by the Gateway Server when an uplink message arrives.
	HandleUplink(context.Context, *UplinkMessage) (*types.Empty, error)
	// Called by the Gateway Server when a Tx acknowledgment arrives.
	ReportTxAcknowledgment(context.Context, *GatewayTxAcknowledgment) (*types.Empty, error)
}

// UnimplementedGsNsServer can be embedded to have forward compatible implementations.
type UnimplementedGsNsServer struct {
}

func (*UnimplementedGsNsServer) HandleUplink(ctx context.Context, req *UplinkMessage) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleUplink not implemented")
}
func (*UnimplementedGsNsServer) ReportTxAcknowledgment(ctx context.Context, req *GatewayTxAcknowledgment) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportTxAcknowledgment not implemented")
}

func RegisterGsNsServer(s *grpc.Server, srv GsNsServer) {
	s.RegisterService(&_GsNs_serviceDesc, srv)
}

func _GsNs_HandleUplink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UplinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GsNsServer).HandleUplink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GsNs/HandleUplink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GsNsServer).HandleUplink(ctx, req.(*UplinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GsNs_ReportTxAcknowledgment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayTxAcknowledgment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GsNsServer).ReportTxAcknowledgment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GsNs/ReportTxAcknowledgment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GsNsServer).ReportTxAcknowledgment(ctx, req.(*GatewayTxAcknowledgment))
	}
	return interceptor(ctx, in, info, handler)
}

var _GsNs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GsNs",
	HandlerType: (*GsNsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleUplink",
			Handler:    _GsNs_HandleUplink_Handler,
		},
		{
			MethodName: "ReportTxAcknowledgment",
			Handler:    _GsNs_ReportTxAcknowledgment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/networkserver.proto",
}

// NsEndDeviceRegistryClient is the client API for NsEndDeviceRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NsEndDeviceRegistryClient interface {
	// Get returns the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// Set creates or updates the device.
	Set(ctx context.Context, in *SetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// ResetFactoryDefaults resets device state to factory defaults.
	ResetFactoryDefaults(ctx context.Context, in *ResetAndGetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// Delete deletes the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type nsEndDeviceRegistryClient struct {
	cc *grpc.ClientConn
}

func NewNsEndDeviceRegistryClient(cc *grpc.ClientConn) NsEndDeviceRegistryClient {
	return &nsEndDeviceRegistryClient{cc}
}

func (c *nsEndDeviceRegistryClient) Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsEndDeviceRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsEndDeviceRegistryClient) Set(ctx context.Context, in *SetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsEndDeviceRegistry/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsEndDeviceRegistryClient) ResetFactoryDefaults(ctx context.Context, in *ResetAndGetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsEndDeviceRegistry/ResetFactoryDefaults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsEndDeviceRegistryClient) Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsEndDeviceRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NsEndDeviceRegistryServer is the server API for NsEndDeviceRegistry service.
type NsEndDeviceRegistryServer interface {
	// Get returns the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Get(context.Context, *GetEndDeviceRequest) (*EndDevice, error)
	// Set creates or updates the device.
	Set(context.Context, *SetEndDeviceRequest) (*EndDevice, error)
	// ResetFactoryDefaults resets device state to factory defaults.
	ResetFactoryDefaults(context.Context, *ResetAndGetEndDeviceRequest) (*EndDevice, error)
	// Delete deletes the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Delete(context.Context, *EndDeviceIdentifiers) (*types.Empty, error)
}

// UnimplementedNsEndDeviceRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedNsEndDeviceRegistryServer struct {
}

func (*UnimplementedNsEndDeviceRegistryServer) Get(ctx context.Context, req *GetEndDeviceRequest) (*EndDevice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedNsEndDeviceRegistryServer) Set(ctx context.Context, req *SetEndDeviceRequest) (*EndDevice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedNsEndDeviceRegistryServer) ResetFactoryDefaults(ctx context.Context, req *ResetAndGetEndDeviceRequest) (*EndDevice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetFactoryDefaults not implemented")
}
func (*UnimplementedNsEndDeviceRegistryServer) Delete(ctx context.Context, req *EndDeviceIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterNsEndDeviceRegistryServer(s *grpc.Server, srv NsEndDeviceRegistryServer) {
	s.RegisterService(&_NsEndDeviceRegistry_serviceDesc, srv)
}

func _NsEndDeviceRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsEndDeviceRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsEndDeviceRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsEndDeviceRegistryServer).Get(ctx, req.(*GetEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsEndDeviceRegistry_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsEndDeviceRegistryServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsEndDeviceRegistry/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsEndDeviceRegistryServer).Set(ctx, req.(*SetEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsEndDeviceRegistry_ResetFactoryDefaults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetAndGetEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsEndDeviceRegistryServer).ResetFactoryDefaults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsEndDeviceRegistry/ResetFactoryDefaults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsEndDeviceRegistryServer).ResetFactoryDefaults(ctx, req.(*ResetAndGetEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsEndDeviceRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDeviceIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsEndDeviceRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsEndDeviceRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsEndDeviceRegistryServer).Delete(ctx, req.(*EndDeviceIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _NsEndDeviceRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.NsEndDeviceRegistry",
	HandlerType: (*NsEndDeviceRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _NsEndDeviceRegistry_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _NsEndDeviceRegistry_Set_Handler,
		},
		{
			MethodName: "ResetFactoryDefaults",
			Handler:    _NsEndDeviceRegistry_ResetFactoryDefaults_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NsEndDeviceRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/networkserver.proto",
}

func (m *GenerateDevAddrResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenerateDevAddrResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenerateDevAddrResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DevAddr != nil {
		{
			size := m.DevAddr.Size()
			i -= size
			if _, err := m.DevAddr.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintNetworkserver(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNetworkserver(dAtA []byte, offset int, v uint64) int {
	offset -= sovNetworkserver(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedGenerateDevAddrResponse(r randyNetworkserver, easy bool) *GenerateDevAddrResponse {
	this := &GenerateDevAddrResponse{}
	this.DevAddr = go_thethings_network_lorawan_stack_v3_pkg_types.NewPopulatedDevAddr(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyNetworkserver interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneNetworkserver(r randyNetworkserver) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringNetworkserver(r randyNetworkserver) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneNetworkserver(r)
	}
	return string(tmps)
}
func randUnrecognizedNetworkserver(r randyNetworkserver, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldNetworkserver(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldNetworkserver(dAtA []byte, r randyNetworkserver, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateNetworkserver(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateNetworkserver(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateNetworkserver(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateNetworkserver(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateNetworkserver(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateNetworkserver(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateNetworkserver(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *GenerateDevAddrResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DevAddr != nil {
		l = m.DevAddr.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	return n
}

func sovNetworkserver(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNetworkserver(x uint64) (n int) {
	return sovNetworkserver((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *GenerateDevAddrResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GenerateDevAddrResponse{`,
		`DevAddr:` + fmt.Sprintf("%v", this.DevAddr) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringNetworkserver(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GenerateDevAddrResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenerateDevAddrResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenerateDevAddrResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v go_thethings_network_lorawan_stack_v3_pkg_types.DevAddr
			m.DevAddr = &v
			if err := m.DevAddr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkserver
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkserver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNetworkserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetworkserver
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNetworkserver
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNetworkserver
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNetworkserver
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNetworkserver        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkserver          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNetworkserver = fmt.Errorf("proto: unexpected end of group")
)
