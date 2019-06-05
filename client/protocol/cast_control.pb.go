// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cast_control.proto

package cast_control

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type DeviceID struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceID) Reset()         { *m = DeviceID{} }
func (m *DeviceID) String() string { return proto.CompactTextString(m) }
func (*DeviceID) ProtoMessage()    {}
func (*DeviceID) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd4c6cd65c0abc5, []int{0}
}

func (m *DeviceID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceID.Unmarshal(m, b)
}
func (m *DeviceID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceID.Marshal(b, m, deterministic)
}
func (m *DeviceID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceID.Merge(m, src)
}
func (m *DeviceID) XXX_Size() int {
	return xxx_messageInfo_DeviceID.Size(m)
}
func (m *DeviceID) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceID.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceID proto.InternalMessageInfo

func (m *DeviceID) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type Device struct {
	FriendlyName         string    `protobuf:"bytes,1,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	ModelName            string    `protobuf:"bytes,2,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	Manufacturer         string    `protobuf:"bytes,3,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	DeviceId             *DeviceID `protobuf:"bytes,4,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	CastType             string    `protobuf:"bytes,5,opt,name=cast_type,json=castType,proto3" json:"cast_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd4c6cd65c0abc5, []int{1}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *Device) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *Device) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *Device) GetDeviceId() *DeviceID {
	if m != nil {
		return m.DeviceId
	}
	return nil
}

func (m *Device) GetCastType() string {
	if m != nil {
		return m.CastType
	}
	return ""
}

type Status struct {
	Volume               float64  `protobuf:"fixed64,1,opt,name=volume,proto3" json:"volume,omitempty"`
	Muted                bool     `protobuf:"varint,2,opt,name=muted,proto3" json:"muted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd4c6cd65c0abc5, []int{2}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Status) GetMuted() bool {
	if m != nil {
		return m.Muted
	}
	return false
}

type Volume struct {
	Volume               float64  `protobuf:"fixed64,1,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Volume) Reset()         { *m = Volume{} }
func (m *Volume) String() string { return proto.CompactTextString(m) }
func (*Volume) ProtoMessage()    {}
func (*Volume) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd4c6cd65c0abc5, []int{3}
}

func (m *Volume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Volume.Unmarshal(m, b)
}
func (m *Volume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Volume.Marshal(b, m, deterministic)
}
func (m *Volume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Volume.Merge(m, src)
}
func (m *Volume) XXX_Size() int {
	return xxx_messageInfo_Volume.Size(m)
}
func (m *Volume) XXX_DiscardUnknown() {
	xxx_messageInfo_Volume.DiscardUnknown(m)
}

var xxx_messageInfo_Volume proto.InternalMessageInfo

func (m *Volume) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type ListDeviceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDeviceRequest) Reset()         { *m = ListDeviceRequest{} }
func (m *ListDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*ListDeviceRequest) ProtoMessage()    {}
func (*ListDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd4c6cd65c0abc5, []int{4}
}

func (m *ListDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeviceRequest.Unmarshal(m, b)
}
func (m *ListDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeviceRequest.Marshal(b, m, deterministic)
}
func (m *ListDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeviceRequest.Merge(m, src)
}
func (m *ListDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_ListDeviceRequest.Size(m)
}
func (m *ListDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeviceRequest proto.InternalMessageInfo

type AdjustVolumeRequest struct {
	DeviceId             *DeviceID `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	RelativeVolume       float64   `protobuf:"fixed64,2,opt,name=relative_volume,json=relativeVolume,proto3" json:"relative_volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AdjustVolumeRequest) Reset()         { *m = AdjustVolumeRequest{} }
func (m *AdjustVolumeRequest) String() string { return proto.CompactTextString(m) }
func (*AdjustVolumeRequest) ProtoMessage()    {}
func (*AdjustVolumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd4c6cd65c0abc5, []int{5}
}

func (m *AdjustVolumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdjustVolumeRequest.Unmarshal(m, b)
}
func (m *AdjustVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdjustVolumeRequest.Marshal(b, m, deterministic)
}
func (m *AdjustVolumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdjustVolumeRequest.Merge(m, src)
}
func (m *AdjustVolumeRequest) XXX_Size() int {
	return xxx_messageInfo_AdjustVolumeRequest.Size(m)
}
func (m *AdjustVolumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdjustVolumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdjustVolumeRequest proto.InternalMessageInfo

func (m *AdjustVolumeRequest) GetDeviceId() *DeviceID {
	if m != nil {
		return m.DeviceId
	}
	return nil
}

func (m *AdjustVolumeRequest) GetRelativeVolume() float64 {
	if m != nil {
		return m.RelativeVolume
	}
	return 0
}

type SetVolumeRequest struct {
	DeviceId             *DeviceID `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Volume               float64   `protobuf:"fixed64,2,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SetVolumeRequest) Reset()         { *m = SetVolumeRequest{} }
func (m *SetVolumeRequest) String() string { return proto.CompactTextString(m) }
func (*SetVolumeRequest) ProtoMessage()    {}
func (*SetVolumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd4c6cd65c0abc5, []int{6}
}

func (m *SetVolumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetVolumeRequest.Unmarshal(m, b)
}
func (m *SetVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetVolumeRequest.Marshal(b, m, deterministic)
}
func (m *SetVolumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetVolumeRequest.Merge(m, src)
}
func (m *SetVolumeRequest) XXX_Size() int {
	return xxx_messageInfo_SetVolumeRequest.Size(m)
}
func (m *SetVolumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetVolumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetVolumeRequest proto.InternalMessageInfo

func (m *SetVolumeRequest) GetDeviceId() *DeviceID {
	if m != nil {
		return m.DeviceId
	}
	return nil
}

func (m *SetVolumeRequest) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func init() {
	proto.RegisterType((*DeviceID)(nil), "cast_control.DeviceID")
	proto.RegisterType((*Device)(nil), "cast_control.Device")
	proto.RegisterType((*Status)(nil), "cast_control.Status")
	proto.RegisterType((*Volume)(nil), "cast_control.Volume")
	proto.RegisterType((*ListDeviceRequest)(nil), "cast_control.ListDeviceRequest")
	proto.RegisterType((*AdjustVolumeRequest)(nil), "cast_control.AdjustVolumeRequest")
	proto.RegisterType((*SetVolumeRequest)(nil), "cast_control.SetVolumeRequest")
}

func init() { proto.RegisterFile("cast_control.proto", fileDescriptor_edd4c6cd65c0abc5) }

var fileDescriptor_edd4c6cd65c0abc5 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0x6e, 0xe2, 0xbd, 0x21, 0x39, 0x8d, 0x5e, 0x3d, 0xf7, 0x72, 0x09, 0x8a, 0xb5, 0x8e, 0x0b,
	0xbb, 0x2a, 0xd2, 0x82, 0xfb, 0xd2, 0x82, 0x16, 0xc4, 0x45, 0x2a, 0x6e, 0xc3, 0x98, 0x39, 0x85,
	0x48, 0x7e, 0x6a, 0x66, 0xa6, 0xd0, 0x57, 0xf2, 0x31, 0x7c, 0x32, 0x71, 0x26, 0xb1, 0x49, 0xad,
	0xd0, 0xc5, 0xdd, 0x65, 0xbe, 0xf9, 0xce, 0xf7, 0x73, 0x86, 0x00, 0xa6, 0x5c, 0xaa, 0x24, 0xad,
	0x4a, 0x55, 0x57, 0xf9, 0x74, 0x57, 0x57, 0xaa, 0xc2, 0xb0, 0x8b, 0xb1, 0x11, 0xf8, 0x2b, 0xda,
	0x67, 0x29, 0xad, 0x57, 0x88, 0x70, 0xa5, 0x75, 0x26, 0x22, 0x67, 0xec, 0x4c, 0x82, 0xd8, 0x7c,
	0xb3, 0x5f, 0x0e, 0x78, 0x96, 0x80, 0x6f, 0xe0, 0xf1, 0xb6, 0xce, 0xa8, 0x14, 0xf9, 0x21, 0x29,
	0x79, 0x41, 0x0d, 0x2f, 0x6c, 0xc1, 0xcf, 0xbc, 0x20, 0x7c, 0x09, 0x50, 0x54, 0x82, 0x72, 0xcb,
	0x70, 0x0d, 0x23, 0x30, 0x88, 0xb9, 0x66, 0x10, 0x16, 0xbc, 0xd4, 0x5b, 0x9e, 0x2a, 0x5d, 0x53,
	0x1d, 0x3d, 0xb2, 0x12, 0x5d, 0x0c, 0xe7, 0x10, 0x08, 0xe3, 0x98, 0x64, 0x22, 0xba, 0x1a, 0x3b,
	0x93, 0xe1, 0xec, 0x7e, 0xda, 0x2b, 0xd2, 0x26, 0x8e, 0x7d, 0x4b, 0x5c, 0x0b, 0x7c, 0x01, 0x81,
	0xa1, 0xa8, 0xc3, 0x8e, 0xa2, 0x6b, 0xa3, 0xea, 0xff, 0x01, 0xbe, 0x1c, 0x76, 0xc4, 0xde, 0x83,
	0xb7, 0x51, 0x5c, 0x69, 0x89, 0xf7, 0xe0, 0xed, 0xab, 0x5c, 0x37, 0xe1, 0x9d, 0xb8, 0x39, 0xe1,
	0x1d, 0x5c, 0x17, 0x5a, 0x91, 0x30, 0x89, 0xfd, 0xd8, 0x1e, 0xd8, 0x18, 0xbc, 0xaf, 0xf6, 0xfe,
	0x3f, 0x73, 0xec, 0x16, 0x9e, 0x7d, 0xca, 0xa4, 0xb2, 0x81, 0x62, 0xfa, 0xa1, 0x49, 0x2a, 0x26,
	0xe1, 0x76, 0x21, 0xbe, 0x6b, 0xa9, 0xec, 0x70, 0x03, 0xf7, 0x7b, 0x39, 0x17, 0xf6, 0x7a, 0x0b,
	0x37, 0x35, 0xe5, 0x5c, 0x65, 0x7b, 0x4a, 0x9a, 0x04, 0xae, 0x49, 0xf0, 0xa4, 0x85, 0xad, 0x09,
	0x4b, 0xe0, 0xe9, 0x86, 0x1e, 0xc2, 0xf1, 0x58, 0xd5, 0xed, 0x56, 0x9d, 0xfd, 0x74, 0x61, 0xb8,
	0xe4, 0x52, 0x2d, 0xed, 0x28, 0x7e, 0x84, 0xe1, 0xb1, 0xba, 0xc4, 0x57, 0x7d, 0xe1, 0x7f, 0xb6,
	0xf2, 0xfc, 0xee, 0x9c, 0x33, 0x1b, 0xbc, 0x73, 0x70, 0x01, 0x37, 0x1f, 0xa8, 0x61, 0xb7, 0xef,
	0x74, 0x3e, 0xe6, 0xa9, 0x88, 0x65, 0xb3, 0x01, 0xae, 0x21, 0xec, 0xae, 0x1c, 0x5f, 0xf7, 0x79,
	0x67, 0x9e, 0xe3, 0x54, 0xaa, 0x59, 0xe3, 0x00, 0x97, 0x10, 0xfc, 0x5d, 0x24, 0x8e, 0x4e, 0xfc,
	0xe8, 0x32, 0x91, 0x6f, 0x9e, 0xf9, 0xd7, 0xe6, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x60, 0xb8,
	0x01, 0x24, 0x81, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CastControlClient is the client API for CastControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CastControlClient interface {
	// Method definitions go here
	ListDevices(ctx context.Context, in *ListDeviceRequest, opts ...grpc.CallOption) (CastControl_ListDevicesClient, error)
	GetDeviceStatus(ctx context.Context, in *DeviceID, opts ...grpc.CallOption) (*Status, error)
	AdjustVolume(ctx context.Context, in *AdjustVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
	SetVolume(ctx context.Context, in *SetVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
}

type castControlClient struct {
	cc *grpc.ClientConn
}

func NewCastControlClient(cc *grpc.ClientConn) CastControlClient {
	return &castControlClient{cc}
}

func (c *castControlClient) ListDevices(ctx context.Context, in *ListDeviceRequest, opts ...grpc.CallOption) (CastControl_ListDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CastControl_serviceDesc.Streams[0], "/cast_control.CastControl/ListDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &castControlListDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CastControl_ListDevicesClient interface {
	Recv() (*Device, error)
	grpc.ClientStream
}

type castControlListDevicesClient struct {
	grpc.ClientStream
}

func (x *castControlListDevicesClient) Recv() (*Device, error) {
	m := new(Device)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *castControlClient) GetDeviceStatus(ctx context.Context, in *DeviceID, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/cast_control.CastControl/GetDeviceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *castControlClient) AdjustVolume(ctx context.Context, in *AdjustVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := c.cc.Invoke(ctx, "/cast_control.CastControl/AdjustVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *castControlClient) SetVolume(ctx context.Context, in *SetVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := c.cc.Invoke(ctx, "/cast_control.CastControl/SetVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CastControlServer is the server API for CastControl service.
type CastControlServer interface {
	// Method definitions go here
	ListDevices(*ListDeviceRequest, CastControl_ListDevicesServer) error
	GetDeviceStatus(context.Context, *DeviceID) (*Status, error)
	AdjustVolume(context.Context, *AdjustVolumeRequest) (*Volume, error)
	SetVolume(context.Context, *SetVolumeRequest) (*Volume, error)
}

// UnimplementedCastControlServer can be embedded to have forward compatible implementations.
type UnimplementedCastControlServer struct {
}

func (*UnimplementedCastControlServer) ListDevices(req *ListDeviceRequest, srv CastControl_ListDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListDevices not implemented")
}
func (*UnimplementedCastControlServer) GetDeviceStatus(ctx context.Context, req *DeviceID) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceStatus not implemented")
}
func (*UnimplementedCastControlServer) AdjustVolume(ctx context.Context, req *AdjustVolumeRequest) (*Volume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdjustVolume not implemented")
}
func (*UnimplementedCastControlServer) SetVolume(ctx context.Context, req *SetVolumeRequest) (*Volume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVolume not implemented")
}

func RegisterCastControlServer(s *grpc.Server, srv CastControlServer) {
	s.RegisterService(&_CastControl_serviceDesc, srv)
}

func _CastControl_ListDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListDeviceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CastControlServer).ListDevices(m, &castControlListDevicesServer{stream})
}

type CastControl_ListDevicesServer interface {
	Send(*Device) error
	grpc.ServerStream
}

type castControlListDevicesServer struct {
	grpc.ServerStream
}

func (x *castControlListDevicesServer) Send(m *Device) error {
	return x.ServerStream.SendMsg(m)
}

func _CastControl_GetDeviceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CastControlServer).GetDeviceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cast_control.CastControl/GetDeviceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CastControlServer).GetDeviceStatus(ctx, req.(*DeviceID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CastControl_AdjustVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdjustVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CastControlServer).AdjustVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cast_control.CastControl/AdjustVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CastControlServer).AdjustVolume(ctx, req.(*AdjustVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CastControl_SetVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CastControlServer).SetVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cast_control.CastControl/SetVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CastControlServer).SetVolume(ctx, req.(*SetVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CastControl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cast_control.CastControl",
	HandlerType: (*CastControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeviceStatus",
			Handler:    _CastControl_GetDeviceStatus_Handler,
		},
		{
			MethodName: "AdjustVolume",
			Handler:    _CastControl_AdjustVolume_Handler,
		},
		{
			MethodName: "SetVolume",
			Handler:    _CastControl_SetVolume_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListDevices",
			Handler:       _CastControl_ListDevices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cast_control.proto",
}
