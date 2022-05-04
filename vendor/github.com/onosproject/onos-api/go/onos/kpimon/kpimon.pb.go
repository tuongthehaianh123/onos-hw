package hw

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetRequest struct {
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dc9732ac3e9f77, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type GetResponse struct {
	Measurements map[string]*MeasurementItems `protobuf:"bytes,1,rep,name=measurements,proto3" json:"measurements,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dc9732ac3e9f77, []int{1}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetMeasurements() map[string]*MeasurementItems {
	if m != nil {
		return m.Measurements
	}
	return nil
}

type MeasurementItems struct {
	MeasurementItems []*MeasurementItem `protobuf:"bytes,1,rep,name=measurement_items,json=measurementItems,proto3" json:"measurement_items,omitempty"`
}

func (m *MeasurementItems) Reset()         { *m = MeasurementItems{} }
func (m *MeasurementItems) String() string { return proto.CompactTextString(m) }
func (*MeasurementItems) ProtoMessage()    {}
func (*MeasurementItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dc9732ac3e9f77, []int{2}
}
func (m *MeasurementItems) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MeasurementItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MeasurementItems.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MeasurementItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeasurementItems.Merge(m, src)
}
func (m *MeasurementItems) XXX_Size() int {
	return m.Size()
}
func (m *MeasurementItems) XXX_DiscardUnknown() {
	xxx_messageInfo_MeasurementItems.DiscardUnknown(m)
}

var xxx_messageInfo_MeasurementItems proto.InternalMessageInfo

func (m *MeasurementItems) GetMeasurementItems() []*MeasurementItem {
	if m != nil {
		return m.MeasurementItems
	}
	return nil
}

type MeasurementItem struct {
	MeasurementRecords []*MeasurementRecord `protobuf:"bytes,1,rep,name=measurement_records,json=measurementRecords,proto3" json:"measurement_records,omitempty"`
}

func (m *MeasurementItem) Reset()         { *m = MeasurementItem{} }
func (m *MeasurementItem) String() string { return proto.CompactTextString(m) }
func (*MeasurementItem) ProtoMessage()    {}
func (*MeasurementItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dc9732ac3e9f77, []int{3}
}
func (m *MeasurementItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MeasurementItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MeasurementItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MeasurementItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeasurementItem.Merge(m, src)
}
func (m *MeasurementItem) XXX_Size() int {
	return m.Size()
}
func (m *MeasurementItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MeasurementItem.DiscardUnknown(m)
}

var xxx_messageInfo_MeasurementItem proto.InternalMessageInfo

func (m *MeasurementItem) GetMeasurementRecords() []*MeasurementRecord {
	if m != nil {
		return m.MeasurementRecords
	}
	return nil
}

type IntegerValue struct {
	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *IntegerValue) Reset()         { *m = IntegerValue{} }
func (m *IntegerValue) String() string { return proto.CompactTextString(m) }
func (*IntegerValue) ProtoMessage()    {}
func (*IntegerValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dc9732ac3e9f77, []int{4}
}
func (m *IntegerValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IntegerValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IntegerValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IntegerValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegerValue.Merge(m, src)
}
func (m *IntegerValue) XXX_Size() int {
	return m.Size()
}
func (m *IntegerValue) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegerValue.DiscardUnknown(m)
}

var xxx_messageInfo_IntegerValue proto.InternalMessageInfo

func (m *IntegerValue) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type RealValue struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *RealValue) Reset()         { *m = RealValue{} }
func (m *RealValue) String() string { return proto.CompactTextString(m) }
func (*RealValue) ProtoMessage()    {}
func (*RealValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dc9732ac3e9f77, []int{5}
}
func (m *RealValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RealValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RealValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RealValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealValue.Merge(m, src)
}
func (m *RealValue) XXX_Size() int {
	return m.Size()
}
func (m *RealValue) XXX_DiscardUnknown() {
	xxx_messageInfo_RealValue.DiscardUnknown(m)
}

var xxx_messageInfo_RealValue proto.InternalMessageInfo

func (m *RealValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type NoValue struct {
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *NoValue) Reset()         { *m = NoValue{} }
func (m *NoValue) String() string { return proto.CompactTextString(m) }
func (*NoValue) ProtoMessage()    {}
func (*NoValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dc9732ac3e9f77, []int{6}
}
func (m *NoValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoValue.Merge(m, src)
}
func (m *NoValue) XXX_Size() int {
	return m.Size()
}
func (m *NoValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NoValue.DiscardUnknown(m)
}

var xxx_messageInfo_NoValue proto.InternalMessageInfo

func (m *NoValue) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MeasurementRecord struct {
	Timestamp        uint64     `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MeasurementName  string     `protobuf:"bytes,3,opt,name=measurement_name,json=measurementName,proto3" json:"measurement_name,omitempty"`
	MeasurementValue *types.Any `protobuf:"bytes,4,opt,name=measurement_value,json=measurementValue,proto3" json:"measurement_value,omitempty"`
}

func (m *MeasurementRecord) Reset()         { *m = MeasurementRecord{} }
func (m *MeasurementRecord) String() string { return proto.CompactTextString(m) }
func (*MeasurementRecord) ProtoMessage()    {}
func (*MeasurementRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dc9732ac3e9f77, []int{7}
}
func (m *MeasurementRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MeasurementRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MeasurementRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MeasurementRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeasurementRecord.Merge(m, src)
}
func (m *MeasurementRecord) XXX_Size() int {
	return m.Size()
}
func (m *MeasurementRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_MeasurementRecord.DiscardUnknown(m)
}

var xxx_messageInfo_MeasurementRecord proto.InternalMessageInfo

func (m *MeasurementRecord) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MeasurementRecord) GetMeasurementName() string {
	if m != nil {
		return m.MeasurementName
	}
	return ""
}

func (m *MeasurementRecord) GetMeasurementValue() *types.Any {
	if m != nil {
		return m.MeasurementValue
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "onos.hw.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "onos.hw.GetResponse")
	proto.RegisterMapType((map[string]*MeasurementItems)(nil), "onos.hw.GetResponse.MeasurementsEntry")
	proto.RegisterType((*MeasurementItems)(nil), "onos.hw.MeasurementItems")
	proto.RegisterType((*MeasurementItem)(nil), "onos.hw.MeasurementItem")
	proto.RegisterType((*IntegerValue)(nil), "onos.hw.IntegerValue")
	proto.RegisterType((*RealValue)(nil), "onos.hw.RealValue")
	proto.RegisterType((*NoValue)(nil), "onos.hw.NoValue")
	proto.RegisterType((*MeasurementRecord)(nil), "onos.hw.MeasurementRecord")
}

func init() { proto.RegisterFile("onos/hw/hw.proto", fileDescriptor_b4dc9732ac3e9f77) }

var fileDescriptor_b4dc9732ac3e9f77 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x4d, 0x5b, 0x94, 0x97, 0x48, 0x4d, 0x8e, 0x4a, 0x98, 0xa8, 0x98, 0x60, 0x31,
	0x04, 0x86, 0x0b, 0x4a, 0x17, 0xc4, 0x56, 0x24, 0x54, 0x22, 0x20, 0x48, 0x37, 0xc0, 0x04, 0xd5,
	0xa5, 0x3c, 0x42, 0xd4, 0xdc, 0x9d, 0xf1, 0x9d, 0x91, 0xfc, 0x2d, 0xd8, 0x19, 0xf8, 0x24, 0xec,
	0x8c, 0x1d, 0x19, 0x51, 0xf2, 0x45, 0x90, 0xcf, 0xb1, 0x38, 0xbb, 0x0d, 0x0b, 0x53, 0xe2, 0xf7,
	0x7e, 0xfe, 0xfb, 0xbd, 0xff, 0xfb, 0x43, 0xa0, 0x95, 0x36, 0xa3, 0x8b, 0x78, 0x21, 0xb5, 0xda,
	0xfc, 0xb0, 0x38, 0xd1, 0x56, 0xd3, 0x76, 0xde, 0x61, 0x45, 0xa9, 0x7f, 0x7b, 0xae, 0xf5, 0x7c,
	0x89, 0x23, 0xd7, 0x9a, 0xa5, 0x1f, 0x47, 0x42, 0x65, 0x05, 0x17, 0x75, 0x00, 0x4e, 0xd1, 0x72,
	0xfc, 0x9c, 0xa2, 0xb1, 0xd1, 0x0f, 0x02, 0x6d, 0xf7, 0x68, 0x62, 0xad, 0x0c, 0xd2, 0x29, 0x74,
	0x24, 0x0a, 0x93, 0x26, 0x28, 0x51, 0x59, 0x13, 0x90, 0x41, 0x73, 0xd8, 0x1e, 0x3f, 0x64, 0x9e,
	0x38, 0xf3, 0x78, 0xf6, 0xca, 0x83, 0x9f, 0x29, 0x9b, 0x64, 0xbc, 0xf2, 0x7e, 0xff, 0x3d, 0xf4,
	0xae, 0x20, 0xb4, 0x0b, 0xcd, 0x0b, 0xcc, 0x02, 0x32, 0x20, 0xc3, 0x16, 0xcf, 0xff, 0xd2, 0x63,
	0xd8, 0xfb, 0x22, 0x96, 0x29, 0x06, 0x3b, 0x03, 0x32, 0x6c, 0x8f, 0xef, 0x54, 0xbe, 0xe7, 0x09,
	0x4c, 0x2c, 0x4a, 0xc3, 0x0b, 0xf6, 0xc9, 0xce, 0x63, 0x12, 0xbd, 0x83, 0x6e, 0xbd, 0x4d, 0x27,
	0xd0, 0xf3, 0x66, 0x38, 0x5b, 0xe4, 0xc5, 0xcd, 0x22, 0x47, 0xff, 0x12, 0xe6, 0x5d, 0x59, 0x93,
	0x8a, 0x66, 0x70, 0x50, 0x83, 0xe8, 0x6b, 0xb8, 0xe9, 0xab, 0x27, 0x78, 0xae, 0x93, 0x0f, 0xa5,
	0x7e, 0xb8, 0x4d, 0x9f, 0x3b, 0x8c, 0x53, 0x59, 0x2f, 0x99, 0xe8, 0x3e, 0x74, 0x26, 0xca, 0xe2,
	0x1c, 0x93, 0x37, 0xf9, 0x5a, 0xf4, 0xb0, 0xf4, 0x22, 0xf7, 0xa7, 0xb9, 0x59, 0x36, 0xba, 0x07,
	0x2d, 0x8e, 0x62, 0x79, 0x0d, 0x42, 0x4a, 0xe4, 0x2e, 0xdc, 0x98, 0xea, 0x6b, 0x80, 0xbd, 0x12,
	0xf8, 0x4e, 0x2a, 0xd7, 0x28, 0x06, 0xa0, 0x47, 0xd0, 0xb2, 0x0b, 0x89, 0xc6, 0x0a, 0x19, 0x3b,
	0xff, 0x77, 0xf9, 0xdf, 0x02, 0x7d, 0x00, 0xbe, 0x2b, 0x67, 0x4a, 0x48, 0x0c, 0x9a, 0xee, 0x70,
	0x07, 0x5e, 0x7d, 0x2a, 0x24, 0xd2, 0x93, 0xaa, 0xef, 0xc5, 0x00, 0xbb, 0xee, 0xa0, 0x87, 0xac,
	0x08, 0x24, 0x2b, 0x03, 0xc9, 0x4e, 0x54, 0x56, 0xf1, 0xdb, 0xcd, 0x3d, 0xfe, 0x46, 0x60, 0xff,
	0x85, 0x33, 0x8f, 0x9e, 0x42, 0xf7, 0xe5, 0xc2, 0x58, 0x3f, 0x3d, 0xf4, 0xd6, 0xd5, 0x1c, 0xba,
	0x18, 0xf7, 0x83, 0x6d, 0x01, 0x8d, 0x1a, 0xf4, 0x39, 0xf4, 0xde, 0x0a, 0x7b, 0xfe, 0xe9, 0x3f,
	0x95, 0x1e, 0x91, 0xa7, 0xc1, 0xcf, 0x55, 0x48, 0x2e, 0x57, 0x21, 0xf9, 0xbd, 0x0a, 0xc9, 0xd7,
	0x75, 0xd8, 0xb8, 0x5c, 0x87, 0x8d, 0x5f, 0xeb, 0xb0, 0x31, 0xdb, 0x77, 0x7b, 0x1d, 0xff, 0x09,
	0x00, 0x00, 0xff, 0xff, 0xb4, 0xb2, 0x5c, 0xe1, 0x9f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HWClient is the client API for HW service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HWClient interface {
	ListMeasurements(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	WatchMeasurements(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (HW_WatchMeasurementsClient, error)
}

type hwClient struct {
	cc *grpc.ClientConn
}

func NewHWClient(cc *grpc.ClientConn) HWClient {
	return &hwClient{cc}
}

func (c *hwClient) ListMeasurements(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/onos.hw.hw/ListMeasurements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hwClient) WatchMeasurements(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (HW_WatchMeasurementsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HW_serviceDesc.Streams[0], "/onos.hw.hw/WatchMeasurements", opts...)
	if err != nil {
		return nil, err
	}
	x := &HWWatchMeasurementsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HW_WatchMeasurementsClient interface {
	Recv() (*GetResponse, error)
	grpc.ClientStream
}

type HWWatchMeasurementsClient struct {
	grpc.ClientStream
}

func (x *HWWatchMeasurementsClient) Recv() (*GetResponse, error) {
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HWServer is the server API for HW service.
type HWServer interface {
	ListMeasurements(context.Context, *GetRequest) (*GetResponse, error)
	WatchMeasurements(*GetRequest, HW_WatchMeasurementsServer) error
}

// UnimplementedHWServer can be embedded to have forward compatible implementations.
type UnimplementedHWServer struct {
}

func (*UnimplementedHWServer) ListMeasurements(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMeasurements not implemented")
}
func (*UnimplementedHWServer) WatchMeasurements(req *GetRequest, srv HW_WatchMeasurementsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchMeasurements not implemented")
}

func RegisterHwServer(s *grpc.Server, srv HWServer) {
	s.RegisterService(&_HW_serviceDesc, srv)
}

func _HW_ListMeasurements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HWServer).ListMeasurements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onos.hw.hw/ListMeasurements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HWServer).ListMeasurements(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HW_WatchMeasurements_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HWServer).WatchMeasurements(m, &HWWatchMeasurementsServer{stream})
}

type HW_WatchMeasurementsServer interface {
	Send(*GetResponse) error
	grpc.ServerStream
}

type HWWatchMeasurementsServer struct {
	grpc.ServerStream
}

func (x *HWWatchMeasurementsServer) Send(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _HW_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onos.hw.hw",
	HandlerType: (*HWServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMeasurements",
			Handler:    _HW_ListMeasurements_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchMeasurements",
			Handler:       _HW_WatchMeasurements_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "onos/hw/hw.proto",
}

func (m *GetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Measurements) > 0 {
		for k := range m.Measurements {
			v := m.Measurements[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintHW(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintHW(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintHW(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MeasurementItems) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MeasurementItems) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MeasurementItems) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MeasurementItems) > 0 {
		for iNdEx := len(m.MeasurementItems) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MeasurementItems[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHW(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MeasurementItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MeasurementItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MeasurementItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MeasurementRecords) > 0 {
		for iNdEx := len(m.MeasurementRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MeasurementRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHW(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *IntegerValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IntegerValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IntegerValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintHW(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RealValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RealValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RealValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *NoValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintHW(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MeasurementRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MeasurementRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MeasurementRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MeasurementValue != nil {
		{
			size, err := m.MeasurementValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHW(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.MeasurementName) > 0 {
		i -= len(m.MeasurementName)
		copy(dAtA[i:], m.MeasurementName)
		i = encodeVarintHW(dAtA, i, uint64(len(m.MeasurementName)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Timestamp != 0 {
		i = encodeVarintHW(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func encodeVarintHW(dAtA []byte, offset int, v uint64) int {
	offset -= sovHW(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Measurements) > 0 {
		for k, v := range m.Measurements {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovHW(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovHW(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovHW(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *MeasurementItems) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MeasurementItems) > 0 {
		for _, e := range m.MeasurementItems {
			l = e.Size()
			n += 1 + l + sovHW(uint64(l))
		}
	}
	return n
}

func (m *MeasurementItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MeasurementRecords) > 0 {
		for _, e := range m.MeasurementRecords {
			l = e.Size()
			n += 1 + l + sovHW(uint64(l))
		}
	}
	return n
}

func (m *IntegerValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovHW(uint64(m.Value))
	}
	return n
}

func (m *RealValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	return n
}

func (m *NoValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovHW(uint64(m.Value))
	}
	return n
}

func (m *MeasurementRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovHW(uint64(m.Timestamp))
	}
	l = len(m.MeasurementName)
	if l > 0 {
		n += 1 + l + sovHW(uint64(l))
	}
	if m.MeasurementValue != nil {
		l = m.MeasurementValue.Size()
		n += 1 + l + sovHW(uint64(l))
	}
	return n
}

func sovHW(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHW(x uint64) (n int) {
	return sovHW(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHW
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
			return fmt.Errorf("proto: GetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHW(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHW
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
func (m *GetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHW
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
			return fmt.Errorf("proto: GetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Measurements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHW
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHW
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHW
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Measurements == nil {
				m.Measurements = make(map[string]*MeasurementItems)
			}
			var mapkey string
			var mapvalue *MeasurementItems
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowHW
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHW
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthHW
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthHW
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHW
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthHW
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthHW
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &MeasurementItems{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipHW(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthHW
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Measurements[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHW(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHW
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
func (m *MeasurementItems) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHW
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
			return fmt.Errorf("proto: MeasurementItems: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MeasurementItems: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeasurementItems", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHW
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHW
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHW
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MeasurementItems = append(m.MeasurementItems, &MeasurementItem{})
			if err := m.MeasurementItems[len(m.MeasurementItems)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHW(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHW
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
func (m *MeasurementItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHW
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
			return fmt.Errorf("proto: MeasurementItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MeasurementItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeasurementRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHW
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHW
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHW
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MeasurementRecords = append(m.MeasurementRecords, &MeasurementRecord{})
			if err := m.MeasurementRecords[len(m.MeasurementRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHW(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHW
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
func (m *IntegerValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHW
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
			return fmt.Errorf("proto: IntegerValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IntegerValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHW
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHW(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHW
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
func (m *RealValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHW
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
			return fmt.Errorf("proto: RealValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RealValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipHW(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHW
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
func (m *NoValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHW
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
			return fmt.Errorf("proto: NoValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHW
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHW(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHW
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
func (m *MeasurementRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHW
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
			return fmt.Errorf("proto: MeasurementRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MeasurementRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHW
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeasurementName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHW
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHW
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHW
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MeasurementName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeasurementValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHW
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHW
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHW
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MeasurementValue == nil {
				m.MeasurementValue = &types.Any{}
			}
			if err := m.MeasurementValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHW(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHW
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
func skipHW(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHW
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
					return 0, ErrIntOverflowHW
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
					return 0, ErrIntOverflowHW
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
				return 0, ErrInvalidLengthHW
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHW
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHW
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHW        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHW          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHW = fmt.Errorf("proto: unexpected end of group")
)