// Code generated by protoc-gen-go. DO NOT EDIT.
// source: medium/import.proto

package _import

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserData struct {
	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age         int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	PhoneNumber []string `protobuf:"bytes,3,rep,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Birthday    *Date    `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	// we make the address repeated coz addresses will be multiple
	UserAddress          []*UserData_Address `protobuf:"bytes,5,rep,name=user_address,json=userAddress,proto3" json:"user_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserData) Reset()         { *m = UserData{} }
func (m *UserData) String() string { return proto.CompactTextString(m) }
func (*UserData) ProtoMessage()    {}
func (*UserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6657fd59da4a9f7, []int{0}
}

func (m *UserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserData.Unmarshal(m, b)
}
func (m *UserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserData.Marshal(b, m, deterministic)
}
func (m *UserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserData.Merge(m, src)
}
func (m *UserData) XXX_Size() int {
	return xxx_messageInfo_UserData.Size(m)
}
func (m *UserData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserData.DiscardUnknown(m)
}

var xxx_messageInfo_UserData proto.InternalMessageInfo

func (m *UserData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserData) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserData) GetPhoneNumber() []string {
	if m != nil {
		return m.PhoneNumber
	}
	return nil
}

func (m *UserData) GetBirthday() *Date {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *UserData) GetUserAddress() []*UserData_Address {
	if m != nil {
		return m.UserAddress
	}
	return nil
}

// nesting the messages
type UserData_Address struct {
	AddressLine1         string   `protobuf:"bytes,1,opt,name=address_line1,json=addressLine1,proto3" json:"address_line1,omitempty"`
	AddressLine2         string   `protobuf:"bytes,2,opt,name=address_line2,json=addressLine2,proto3" json:"address_line2,omitempty"`
	AddressLine3         string   `protobuf:"bytes,3,opt,name=address_line3,json=addressLine3,proto3" json:"address_line3,omitempty"`
	City                 string   `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserData_Address) Reset()         { *m = UserData_Address{} }
func (m *UserData_Address) String() string { return proto.CompactTextString(m) }
func (*UserData_Address) ProtoMessage()    {}
func (*UserData_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6657fd59da4a9f7, []int{0, 0}
}

func (m *UserData_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserData_Address.Unmarshal(m, b)
}
func (m *UserData_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserData_Address.Marshal(b, m, deterministic)
}
func (m *UserData_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserData_Address.Merge(m, src)
}
func (m *UserData_Address) XXX_Size() int {
	return xxx_messageInfo_UserData_Address.Size(m)
}
func (m *UserData_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_UserData_Address.DiscardUnknown(m)
}

var xxx_messageInfo_UserData_Address proto.InternalMessageInfo

func (m *UserData_Address) GetAddressLine1() string {
	if m != nil {
		return m.AddressLine1
	}
	return ""
}

func (m *UserData_Address) GetAddressLine2() string {
	if m != nil {
		return m.AddressLine2
	}
	return ""
}

func (m *UserData_Address) GetAddressLine3() string {
	if m != nil {
		return m.AddressLine3
	}
	return ""
}

func (m *UserData_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *UserData_Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func init() {
	proto.RegisterType((*UserData)(nil), "UserData")
	proto.RegisterType((*UserData_Address)(nil), "UserData.Address")
}

func init() { proto.RegisterFile("medium/import.proto", fileDescriptor_f6657fd59da4a9f7) }

var fileDescriptor_f6657fd59da4a9f7 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x49, 0xd3, 0xd0, 0xe6, 0x52, 0x10, 0x98, 0xc5, 0xea, 0xe4, 0x96, 0xc5, 0x4b, 0x83,
	0x48, 0x78, 0x01, 0xa4, 0x8e, 0x88, 0xc1, 0x12, 0x73, 0xe5, 0x26, 0x27, 0x6a, 0xa9, 0x4e, 0x22,
	0xc7, 0x41, 0xf4, 0x85, 0x78, 0x23, 0xde, 0x07, 0xc5, 0x71, 0x8b, 0x28, 0xdb, 0x7f, 0xdf, 0x7d,
	0x96, 0xee, 0xce, 0x70, 0xa7, 0xb1, 0x54, 0x9d, 0x7e, 0x50, 0xba, 0xa9, 0x8d, 0x4d, 0x1b, 0x53,
	0xdb, 0x7a, 0xce, 0x3c, 0x2c, 0xa5, 0xc5, 0xd5, 0xd0, 0x59, 0xe1, 0xa7, 0xd4, 0xcd, 0x1e, 0x07,
	0x63, 0xf9, 0x3d, 0x82, 0xe9, 0x5b, 0x8b, 0x66, 0x2d, 0xad, 0x24, 0x04, 0xc6, 0x95, 0xd4, 0x48,
	0x03, 0x16, 0xf0, 0x58, 0xb8, 0x4c, 0x6e, 0x20, 0x94, 0xef, 0x48, 0x47, 0x2c, 0xe0, 0x91, 0xe8,
	0x23, 0x59, 0xc0, 0xac, 0xd9, 0xd5, 0x15, 0x6e, 0xaa, 0x4e, 0x6f, 0xd1, 0xd0, 0x90, 0x85, 0x3c,
	0x16, 0x89, 0x63, 0xaf, 0x0e, 0x91, 0x05, 0x4c, 0xb7, 0xca, 0xd8, 0x5d, 0x29, 0x0f, 0x74, 0xcc,
	0x02, 0x9e, 0x64, 0x51, 0xba, 0x96, 0x16, 0xc5, 0x09, 0x93, 0x27, 0x98, 0x75, 0x2d, 0x9a, 0x8d,
	0x2c, 0x4b, 0x83, 0x6d, 0x4b, 0x23, 0x16, 0xf2, 0x24, 0xbb, 0x4d, 0x8f, 0xc3, 0xa4, 0xcf, 0x43,
	0x43, 0x24, 0xbd, 0xe6, 0x8b, 0xf9, 0x57, 0x00, 0x13, 0x9f, 0xc9, 0x3d, 0x5c, 0xf9, 0xc7, 0x9b,
	0xbd, 0xaa, 0xf0, 0xd1, 0x8f, 0x3d, 0xf3, 0xf0, 0xa5, 0x67, 0xe7, 0x52, 0xe6, 0x16, 0xf9, 0x2b,
	0x65, 0xe7, 0x52, 0x4e, 0xc3, 0x7f, 0x52, 0xde, 0x1f, 0xa7, 0x50, 0x76, 0xd8, 0x27, 0x16, 0x2e,
	0x13, 0x0a, 0x93, 0xa2, 0xee, 0x2a, 0x6b, 0x0e, 0x34, 0x72, 0xf8, 0x58, 0x66, 0x39, 0x4c, 0xfc,
	0xa1, 0x09, 0x87, 0x6b, 0x1f, 0x5b, 0x34, 0x1f, 0xaa, 0x40, 0x12, 0x9f, 0xb6, 0x9c, 0xff, 0xc6,
	0xe5, 0xc5, 0xf6, 0xd2, 0xfd, 0x49, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xca, 0x14, 0x9b, 0xfc,
	0xcc, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleClient interface {
	Exampleservice(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*UserData, error)
}

type exampleClient struct {
	cc *grpc.ClientConn
}

func NewExampleClient(cc *grpc.ClientConn) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Exampleservice(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*UserData, error) {
	out := new(UserData)
	err := c.cc.Invoke(ctx, "/example/exampleservice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServer is the server API for Example service.
type ExampleServer interface {
	Exampleservice(context.Context, *UserData) (*UserData, error)
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_Exampleservice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Exampleservice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example/Exampleservice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Exampleservice(ctx, req.(*UserData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "exampleservice",
			Handler:    _Example_Exampleservice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "medium/import.proto",
}
