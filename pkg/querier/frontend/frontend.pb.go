// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: frontend.proto

package frontend

import (
	context "context"
	fmt "fmt"
	queryrange "github.com/cortexproject/cortex/pkg/querier/queryrange"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	httpgrpc "github.com/weaveworks/common/httpgrpc"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ProcessRequest struct {
	HttpRequest       *httpgrpc.HTTPRequest `protobuf:"bytes,1,opt,name=httpRequest,proto3" json:"httpRequest,omitempty"`
	QueryRangeRequest *queryrange.Request   `protobuf:"bytes,2,opt,name=queryRangeRequest,proto3" json:"queryRangeRequest,omitempty"`
}

func (m *ProcessRequest) Reset()      { *m = ProcessRequest{} }
func (*ProcessRequest) ProtoMessage() {}
func (*ProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0}
}
func (m *ProcessRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessRequest.Merge(m, src)
}
func (m *ProcessRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessRequest proto.InternalMessageInfo

func (m *ProcessRequest) GetHttpRequest() *httpgrpc.HTTPRequest {
	if m != nil {
		return m.HttpRequest
	}
	return nil
}

func (m *ProcessRequest) GetQueryRangeRequest() *queryrange.Request {
	if m != nil {
		return m.QueryRangeRequest
	}
	return nil
}

type ProcessResponse struct {
	HttpResponse *httpgrpc.HTTPResponse  `protobuf:"bytes,1,opt,name=httpResponse,proto3" json:"httpResponse,omitempty"`
	ApiResponse  *queryrange.APIResponse `protobuf:"bytes,2,opt,name=apiResponse,proto3" json:"apiResponse,omitempty"`
}

func (m *ProcessResponse) Reset()      { *m = ProcessResponse{} }
func (*ProcessResponse) ProtoMessage() {}
func (*ProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{1}
}
func (m *ProcessResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessResponse.Merge(m, src)
}
func (m *ProcessResponse) XXX_Size() int {
	return m.Size()
}
func (m *ProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessResponse proto.InternalMessageInfo

func (m *ProcessResponse) GetHttpResponse() *httpgrpc.HTTPResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func (m *ProcessResponse) GetApiResponse() *queryrange.APIResponse {
	if m != nil {
		return m.ApiResponse
	}
	return nil
}

func init() {
	proto.RegisterType((*ProcessRequest)(nil), "frontend.ProcessRequest")
	proto.RegisterType((*ProcessResponse)(nil), "frontend.ProcessResponse")
}

func init() { proto.RegisterFile("frontend.proto", fileDescriptor_eca3873955a29cfe) }

var fileDescriptor_eca3873955a29cfe = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0xc6, 0xe3, 0x3b, 0xdc, 0x5b, 0xb9, 0x57, 0x45, 0x18, 0x01, 0xa5, 0x83, 0x85, 0x3a, 0x75,
	0x21, 0x41, 0x05, 0x09, 0xc1, 0x80, 0x54, 0x84, 0xf8, 0xb3, 0x45, 0x51, 0x27, 0xb6, 0x36, 0xb8,
	0x69, 0xa9, 0x1a, 0xbb, 0x8e, 0x43, 0x61, 0x63, 0x64, 0x60, 0xe0, 0x31, 0x78, 0x14, 0xc6, 0x8e,
	0x1d, 0xa9, 0xbb, 0x30, 0xf6, 0x11, 0x50, 0x6d, 0xd7, 0x32, 0x85, 0x29, 0xe7, 0xd3, 0xf9, 0xbe,
	0x73, 0x7e, 0xd1, 0x31, 0x2c, 0x75, 0x38, 0x4d, 0x05, 0x49, 0x6f, 0x7d, 0xc6, 0xa9, 0xa0, 0xa8,
	0xb0, 0xd4, 0x95, 0xbd, 0xa4, 0x27, 0xba, 0x79, 0xdb, 0x8f, 0xe9, 0x20, 0x48, 0x68, 0x42, 0x03,
	0x65, 0x68, 0xe7, 0x1d, 0xa5, 0x94, 0x50, 0x95, 0x0e, 0x56, 0x0e, 0x1d, 0xfb, 0x88, 0xb4, 0xee,
	0xc9, 0x88, 0xf2, 0x7e, 0x16, 0xc4, 0x74, 0x30, 0xa0, 0x69, 0xd0, 0x15, 0x82, 0x25, 0x9c, 0xc5,
	0xb6, 0x30, 0xa9, 0x4b, 0x27, 0x15, 0x53, 0x2e, 0xc8, 0x03, 0xe3, 0xf4, 0x8e, 0xc4, 0xc2, 0xa8,
	0x80, 0xf5, 0x93, 0x60, 0x98, 0x13, 0xde, 0x23, 0x5c, 0x7d, 0x1f, 0x79, 0x2b, 0x4d, 0x88, 0x53,
	0xea, 0x41, 0xd5, 0x17, 0x00, 0x4b, 0x21, 0xa7, 0x31, 0xc9, 0xb2, 0x88, 0x0c, 0x73, 0x92, 0x09,
	0x74, 0x04, 0x8b, 0x8b, 0x6d, 0x46, 0x96, 0xc1, 0x2e, 0xa8, 0x15, 0xeb, 0x9b, 0xbe, 0x25, 0xb8,
	0x6a, 0x36, 0x43, 0xd3, 0x8c, 0x5c, 0x27, 0x6a, 0xc0, 0x75, 0x35, 0x3f, 0x5a, 0xcc, 0x5f, 0xc6,
	0xff, 0xa8, 0xf8, 0x86, 0xef, 0x6c, 0x5e, 0x86, 0x7f, 0xba, 0xab, 0xcf, 0x00, 0xae, 0x59, 0x9c,
	0x8c, 0xd1, 0x34, 0x23, 0xe8, 0x04, 0xfe, 0xd7, 0x5b, 0xb4, 0x36, 0x40, 0x5b, 0xab, 0x40, 0xba,
	0x1b, 0x7d, 0xf3, 0xa2, 0x63, 0x58, 0x6c, 0xb1, 0x9e, 0x8d, 0x6a, 0x98, 0x6d, 0x17, 0xa6, 0x11,
	0x5e, 0xdb, 0xac, 0xeb, 0xad, 0x87, 0xb0, 0x70, 0x61, 0x6e, 0x8a, 0xce, 0xe1, 0x3f, 0x43, 0x85,
	0x76, 0x7c, 0x7b, 0xf9, 0x15, 0xd0, 0x4a, 0xf9, 0x97, 0x96, 0xfe, 0x2d, 0xaf, 0x06, 0xf6, 0xc1,
	0xd9, 0xe9, 0x78, 0x8a, 0xbd, 0xc9, 0x14, 0x7b, 0xf3, 0x29, 0x06, 0x4f, 0x12, 0x83, 0x37, 0x89,
	0xc1, 0xbb, 0xc4, 0x60, 0x2c, 0x31, 0xf8, 0x90, 0x18, 0x7c, 0x4a, 0xec, 0xcd, 0x25, 0x06, 0xaf,
	0x33, 0xec, 0x8d, 0x67, 0xd8, 0x9b, 0xcc, 0xb0, 0x77, 0x63, 0x5f, 0x56, 0xfb, 0xaf, 0x3a, 0xd9,
	0xc1, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x0d, 0xde, 0xba, 0x7c, 0x02, 0x00, 0x00,
}

func (this *ProcessRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProcessRequest)
	if !ok {
		that2, ok := that.(ProcessRequest)
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
	if !this.HttpRequest.Equal(that1.HttpRequest) {
		return false
	}
	if !this.QueryRangeRequest.Equal(that1.QueryRangeRequest) {
		return false
	}
	return true
}
func (this *ProcessResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProcessResponse)
	if !ok {
		that2, ok := that.(ProcessResponse)
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
	if !this.HttpResponse.Equal(that1.HttpResponse) {
		return false
	}
	if !this.ApiResponse.Equal(that1.ApiResponse) {
		return false
	}
	return true
}
func (this *ProcessRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&frontend.ProcessRequest{")
	if this.HttpRequest != nil {
		s = append(s, "HttpRequest: "+fmt.Sprintf("%#v", this.HttpRequest)+",\n")
	}
	if this.QueryRangeRequest != nil {
		s = append(s, "QueryRangeRequest: "+fmt.Sprintf("%#v", this.QueryRangeRequest)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ProcessResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&frontend.ProcessResponse{")
	if this.HttpResponse != nil {
		s = append(s, "HttpResponse: "+fmt.Sprintf("%#v", this.HttpResponse)+",\n")
	}
	if this.ApiResponse != nil {
		s = append(s, "ApiResponse: "+fmt.Sprintf("%#v", this.ApiResponse)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFrontend(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FrontendClient is the client API for Frontend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FrontendClient interface {
	Process(ctx context.Context, opts ...grpc.CallOption) (Frontend_ProcessClient, error)
}

type frontendClient struct {
	cc *grpc.ClientConn
}

func NewFrontendClient(cc *grpc.ClientConn) FrontendClient {
	return &frontendClient{cc}
}

func (c *frontendClient) Process(ctx context.Context, opts ...grpc.CallOption) (Frontend_ProcessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Frontend_serviceDesc.Streams[0], "/frontend.Frontend/Process", opts...)
	if err != nil {
		return nil, err
	}
	x := &frontendProcessClient{stream}
	return x, nil
}

type Frontend_ProcessClient interface {
	Send(*ProcessResponse) error
	Recv() (*ProcessRequest, error)
	grpc.ClientStream
}

type frontendProcessClient struct {
	grpc.ClientStream
}

func (x *frontendProcessClient) Send(m *ProcessResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *frontendProcessClient) Recv() (*ProcessRequest, error) {
	m := new(ProcessRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrontendServer is the server API for Frontend service.
type FrontendServer interface {
	Process(Frontend_ProcessServer) error
}

func RegisterFrontendServer(s *grpc.Server, srv FrontendServer) {
	s.RegisterService(&_Frontend_serviceDesc, srv)
}

func _Frontend_Process_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FrontendServer).Process(&frontendProcessServer{stream})
}

type Frontend_ProcessServer interface {
	Send(*ProcessRequest) error
	Recv() (*ProcessResponse, error)
	grpc.ServerStream
}

type frontendProcessServer struct {
	grpc.ServerStream
}

func (x *frontendProcessServer) Send(m *ProcessRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *frontendProcessServer) Recv() (*ProcessResponse, error) {
	m := new(ProcessResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Frontend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "frontend.Frontend",
	HandlerType: (*FrontendServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Process",
			Handler:       _Frontend_Process_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "frontend.proto",
}

func (m *ProcessRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.HttpRequest != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFrontend(dAtA, i, uint64(m.HttpRequest.Size()))
		n1, err := m.HttpRequest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.QueryRangeRequest != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFrontend(dAtA, i, uint64(m.QueryRangeRequest.Size()))
		n2, err := m.QueryRangeRequest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *ProcessResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.HttpResponse != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFrontend(dAtA, i, uint64(m.HttpResponse.Size()))
		n3, err := m.HttpResponse.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.ApiResponse != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFrontend(dAtA, i, uint64(m.ApiResponse.Size()))
		n4, err := m.ApiResponse.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeVarintFrontend(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ProcessRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpRequest != nil {
		l = m.HttpRequest.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	if m.QueryRangeRequest != nil {
		l = m.QueryRangeRequest.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	return n
}

func (m *ProcessResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpResponse != nil {
		l = m.HttpResponse.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	if m.ApiResponse != nil {
		l = m.ApiResponse.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	return n
}

func sovFrontend(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFrontend(x uint64) (n int) {
	return sovFrontend(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ProcessRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessRequest{`,
		`HttpRequest:` + strings.Replace(fmt.Sprintf("%v", this.HttpRequest), "HTTPRequest", "httpgrpc.HTTPRequest", 1) + `,`,
		`QueryRangeRequest:` + strings.Replace(fmt.Sprintf("%v", this.QueryRangeRequest), "Request", "queryrange.Request", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProcessResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessResponse{`,
		`HttpResponse:` + strings.Replace(fmt.Sprintf("%v", this.HttpResponse), "HTTPResponse", "httpgrpc.HTTPResponse", 1) + `,`,
		`ApiResponse:` + strings.Replace(fmt.Sprintf("%v", this.ApiResponse), "APIResponse", "queryrange.APIResponse", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFrontend(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ProcessRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
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
			return fmt.Errorf("proto: ProcessRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
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
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpRequest == nil {
				m.HttpRequest = &httpgrpc.HTTPRequest{}
			}
			if err := m.HttpRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryRangeRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
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
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.QueryRangeRequest == nil {
				m.QueryRangeRequest = &queryrange.Request{}
			}
			if err := m.QueryRangeRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFrontend
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
func (m *ProcessResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
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
			return fmt.Errorf("proto: ProcessResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
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
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpResponse == nil {
				m.HttpResponse = &httpgrpc.HTTPResponse{}
			}
			if err := m.HttpResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
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
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ApiResponse == nil {
				m.ApiResponse = &queryrange.APIResponse{}
			}
			if err := m.ApiResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFrontend
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
func skipFrontend(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFrontend
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
					return 0, ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFrontend
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
				return 0, ErrInvalidLengthFrontend
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFrontend
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFrontend
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFrontend(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFrontend
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFrontend = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFrontend   = fmt.Errorf("proto: integer overflow")
)