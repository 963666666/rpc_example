// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/search.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type SearchRequest struct {
	Request              string   `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e70f09e802673d, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

type SearchResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e70f09e802673d, []int{1}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type UserRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e70f09e802673d, []int{2}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserResponse struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Admin                string   `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty"`
	Loginci              int64    `protobuf:"varint,3,opt,name=loginci,proto3" json:"loginci,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e70f09e802673d, []int{3}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserResponse) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func (m *UserResponse) GetLoginci() int64 {
	if m != nil {
		return m.Loginci
	}
	return 0
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "proto.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "proto.SearchResponse")
	proto.RegisterType((*UserRequest)(nil), "proto.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "proto.UserResponse")
}

func init() { proto.RegisterFile("proto/search.proto", fileDescriptor_47e70f09e802673d) }

var fileDescriptor_47e70f09e802673d = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x41, 0x4b, 0xc5, 0x30,
	0x10, 0x84, 0x6d, 0x9f, 0xef, 0xa9, 0xfb, 0xb4, 0x87, 0xb5, 0x42, 0x28, 0x08, 0x92, 0x93, 0x82,
	0x54, 0xd0, 0x83, 0x9e, 0xbd, 0x88, 0xd7, 0x14, 0x7f, 0x40, 0x6d, 0x16, 0x0d, 0x68, 0x53, 0x93,
	0x56, 0xf0, 0xea, 0x2f, 0x97, 0x6e, 0x12, 0xb1, 0x9e, 0xba, 0xf3, 0x75, 0x32, 0xcc, 0x00, 0x0e,
	0xce, 0x8e, 0xf6, 0xca, 0x53, 0xeb, 0xba, 0xd7, 0x9a, 0x05, 0xae, 0xf9, 0x23, 0x2f, 0xe0, 0xa8,
	0x61, 0xac, 0xe8, 0x63, 0x22, 0x3f, 0xa2, 0x80, 0x3d, 0x17, 0x4e, 0x91, 0x9d, 0x65, 0xe7, 0x07,
	0x2a, 0x49, 0x79, 0x09, 0x45, 0xb2, 0xfa, 0xc1, 0xf6, 0x9e, 0xb0, 0x82, 0x7d, 0x17, 0xef, 0x68,
	0xfe, 0xd5, 0xf2, 0x14, 0xb6, 0x4f, 0x9e, 0x5c, 0x8a, 0x2d, 0x20, 0x37, 0x9a, 0x4d, 0x2b, 0x95,
	0x1b, 0x2d, 0x15, 0x1c, 0x86, 0xdf, 0x31, 0x0a, 0x61, 0x77, 0xf2, 0xe4, 0x62, 0x0c, 0xdf, 0x58,
	0xc2, 0xba, 0xd5, 0xef, 0xa6, 0x17, 0x39, 0xc3, 0x20, 0xe6, 0x82, 0x6f, 0xf6, 0xc5, 0xf4, 0x9d,
	0x11, 0x2b, 0x8e, 0x4b, 0xf2, 0xfa, 0x3b, 0x4b, 0x63, 0x1a, 0x72, 0x9f, 0xa6, 0x23, 0xbc, 0x85,
	0x4d, 0x00, 0x58, 0x86, 0xd9, 0xf5, 0x62, 0x6c, 0x75, 0xf2, 0x8f, 0xc6, 0xee, 0x3b, 0x78, 0x07,
	0xdb, 0x07, 0x1a, 0xe7, 0x86, 0xf7, 0x5f, 0x8f, 0x1a, 0x31, 0xfa, 0xfe, 0x2c, 0xaa, 0x8e, 0x17,
	0x2c, 0xbd, 0x7c, 0xde, 0x30, 0xbd, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0x41, 0xdf, 0x27, 0xd8,
	0x74, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	GetUserById(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type searchServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchServiceClient(cc *grpc.ClientConn) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetUserById(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	GetUserById(context.Context, *UserRequest) (*UserResponse, error)
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetUserById(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _SearchService_GetUserById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/search.proto",
}
