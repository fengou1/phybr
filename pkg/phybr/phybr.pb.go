// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: phybr.proto

package phybr

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

	_ "github.com/gogo/protobuf/gogoproto"

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

type RegionMeta struct {
	RegionId             uint64   `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	AppliedIndex         uint64   `protobuf:"varint,2,opt,name=applied_index,json=appliedIndex,proto3" json:"applied_index,omitempty"`
	Term                 uint64   `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegionMeta) Reset()         { *m = RegionMeta{} }
func (m *RegionMeta) String() string { return proto.CompactTextString(m) }
func (*RegionMeta) ProtoMessage()    {}
func (*RegionMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_phybr_a3e79a03fe00b96a, []int{0}
}
func (m *RegionMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegionMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegionMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RegionMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionMeta.Merge(dst, src)
}
func (m *RegionMeta) XXX_Size() int {
	return m.Size()
}
func (m *RegionMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionMeta.DiscardUnknown(m)
}

var xxx_messageInfo_RegionMeta proto.InternalMessageInfo

func (m *RegionMeta) GetRegionId() uint64 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *RegionMeta) GetAppliedIndex() uint64 {
	if m != nil {
		return m.AppliedIndex
	}
	return 0
}

func (m *RegionMeta) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

type RegionRecover struct {
	RegionId             uint64   `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Term                 uint64   `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	Silence              bool     `protobuf:"varint,3,opt,name=silence,proto3" json:"silence,omitempty"`
	Tombstone            bool     `protobuf:"varint,4,opt,name=tombstone,proto3" json:"tombstone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegionRecover) Reset()         { *m = RegionRecover{} }
func (m *RegionRecover) String() string { return proto.CompactTextString(m) }
func (*RegionRecover) ProtoMessage()    {}
func (*RegionRecover) Descriptor() ([]byte, []int) {
	return fileDescriptor_phybr_a3e79a03fe00b96a, []int{1}
}
func (m *RegionRecover) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegionRecover) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegionRecover.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RegionRecover) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionRecover.Merge(dst, src)
}
func (m *RegionRecover) XXX_Size() int {
	return m.Size()
}
func (m *RegionRecover) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionRecover.DiscardUnknown(m)
}

var xxx_messageInfo_RegionRecover proto.InternalMessageInfo

func (m *RegionRecover) GetRegionId() uint64 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *RegionRecover) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RegionRecover) GetSilence() bool {
	if m != nil {
		return m.Silence
	}
	return false
}

func (m *RegionRecover) GetTombstone() bool {
	if m != nil {
		return m.Tombstone
	}
	return false
}

func init() {
	proto.RegisterType((*RegionMeta)(nil), "phybr.RegionMeta")
	proto.RegisterType((*RegionRecover)(nil), "phybr.RegionRecover")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Phybr service

type PhybrClient interface {
	RecoverRegions(ctx context.Context, opts ...grpc.CallOption) (Phybr_RecoverRegionsClient, error)
}

type phybrClient struct {
	cc *grpc.ClientConn
}

func NewPhybrClient(cc *grpc.ClientConn) PhybrClient {
	return &phybrClient{cc}
}

func (c *phybrClient) RecoverRegions(ctx context.Context, opts ...grpc.CallOption) (Phybr_RecoverRegionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Phybr_serviceDesc.Streams[0], "/phybr.Phybr/recover_regions", opts...)
	if err != nil {
		return nil, err
	}
	x := &phybrRecoverRegionsClient{stream}
	return x, nil
}

type Phybr_RecoverRegionsClient interface {
	Send(*RegionMeta) error
	Recv() (*RegionRecover, error)
	grpc.ClientStream
}

type phybrRecoverRegionsClient struct {
	grpc.ClientStream
}

func (x *phybrRecoverRegionsClient) Send(m *RegionMeta) error {
	return x.ClientStream.SendMsg(m)
}

func (x *phybrRecoverRegionsClient) Recv() (*RegionRecover, error) {
	m := new(RegionRecover)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Phybr service

type PhybrServer interface {
	RecoverRegions(Phybr_RecoverRegionsServer) error
}

func RegisterPhybrServer(s *grpc.Server, srv PhybrServer) {
	s.RegisterService(&_Phybr_serviceDesc, srv)
}

func _Phybr_RecoverRegions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PhybrServer).RecoverRegions(&phybrRecoverRegionsServer{stream})
}

type Phybr_RecoverRegionsServer interface {
	Send(*RegionRecover) error
	Recv() (*RegionMeta, error)
	grpc.ServerStream
}

type phybrRecoverRegionsServer struct {
	grpc.ServerStream
}

func (x *phybrRecoverRegionsServer) Send(m *RegionRecover) error {
	return x.ServerStream.SendMsg(m)
}

func (x *phybrRecoverRegionsServer) Recv() (*RegionMeta, error) {
	m := new(RegionMeta)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Phybr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "phybr.Phybr",
	HandlerType: (*PhybrServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "recover_regions",
			Handler:       _Phybr_RecoverRegions_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "phybr.proto",
}

func (m *RegionMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegionMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RegionId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPhybr(dAtA, i, uint64(m.RegionId))
	}
	if m.AppliedIndex != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPhybr(dAtA, i, uint64(m.AppliedIndex))
	}
	if m.Term != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPhybr(dAtA, i, uint64(m.Term))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RegionRecover) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegionRecover) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RegionId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPhybr(dAtA, i, uint64(m.RegionId))
	}
	if m.Term != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPhybr(dAtA, i, uint64(m.Term))
	}
	if m.Silence {
		dAtA[i] = 0x18
		i++
		if m.Silence {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Tombstone {
		dAtA[i] = 0x20
		i++
		if m.Tombstone {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPhybr(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RegionMeta) Size() (n int) {
	var l int
	_ = l
	if m.RegionId != 0 {
		n += 1 + sovPhybr(uint64(m.RegionId))
	}
	if m.AppliedIndex != 0 {
		n += 1 + sovPhybr(uint64(m.AppliedIndex))
	}
	if m.Term != 0 {
		n += 1 + sovPhybr(uint64(m.Term))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegionRecover) Size() (n int) {
	var l int
	_ = l
	if m.RegionId != 0 {
		n += 1 + sovPhybr(uint64(m.RegionId))
	}
	if m.Term != 0 {
		n += 1 + sovPhybr(uint64(m.Term))
	}
	if m.Silence {
		n += 2
	}
	if m.Tombstone {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPhybr(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPhybr(x uint64) (n int) {
	return sovPhybr(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegionMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPhybr
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegionMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegionMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionId", wireType)
			}
			m.RegionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPhybr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegionId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppliedIndex", wireType)
			}
			m.AppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPhybr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPhybr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPhybr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPhybr
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegionRecover) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPhybr
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegionRecover: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegionRecover: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionId", wireType)
			}
			m.RegionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPhybr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegionId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPhybr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Silence", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPhybr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Silence = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstone", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPhybr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstone = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPhybr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPhybr
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPhybr(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPhybr
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
					return 0, ErrIntOverflowPhybr
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
					return 0, ErrIntOverflowPhybr
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPhybr
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPhybr
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
				next, err := skipPhybr(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthPhybr = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPhybr   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("phybr.proto", fileDescriptor_phybr_a3e79a03fe00b96a) }

var fileDescriptor_phybr_a3e79a03fe00b96a = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xc8, 0xa8, 0x4c,
	0x2a, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x44, 0xd2, 0xf3, 0xd3,
	0xf3, 0xc1, 0x22, 0xfa, 0x20, 0x16, 0x44, 0x52, 0x8a, 0xbf, 0xa8, 0xb4, 0xb8, 0x04, 0xcc, 0x84,
	0x08, 0x28, 0x25, 0x71, 0x71, 0x05, 0xa5, 0xa6, 0x67, 0xe6, 0xe7, 0xf9, 0xa6, 0x96, 0x24, 0x0a,
	0x49, 0x73, 0x71, 0x16, 0x81, 0x79, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41,
	0x1c, 0x10, 0x01, 0xcf, 0x14, 0x21, 0x65, 0x2e, 0xde, 0xc4, 0x82, 0x82, 0x9c, 0xcc, 0xd4, 0x94,
	0xf8, 0xcc, 0xbc, 0x94, 0xd4, 0x0a, 0x09, 0x26, 0xb0, 0x02, 0x1e, 0xa8, 0xa0, 0x27, 0x48, 0x4c,
	0x48, 0x88, 0x8b, 0xa5, 0x24, 0xb5, 0x28, 0x57, 0x82, 0x19, 0x2c, 0x07, 0x66, 0x2b, 0x55, 0x70,
	0xf1, 0x42, 0xec, 0x08, 0x4a, 0x4d, 0xce, 0x2f, 0x4b, 0x2d, 0xc2, 0x6f, 0x0d, 0xcc, 0x04, 0x26,
	0x84, 0x09, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0x99, 0x39, 0xa9, 0x79, 0xc9, 0xa9, 0x60, 0x83, 0x39,
	0x82, 0x60, 0x5c, 0x21, 0x19, 0x2e, 0xce, 0x92, 0xfc, 0xdc, 0xa4, 0xe2, 0x92, 0xfc, 0xbc, 0x54,
	0x09, 0x16, 0xb0, 0x1c, 0x42, 0xc0, 0xc8, 0x93, 0x8b, 0x35, 0x00, 0x14, 0x1a, 0x42, 0x0e, 0x5c,
	0xfc, 0x45, 0x10, 0xcb, 0xe3, 0x21, 0x16, 0x15, 0x0b, 0x09, 0xea, 0x41, 0x42, 0x0d, 0xe1, 0x7d,
	0x29, 0x11, 0x14, 0x21, 0xa8, 0x6b, 0x95, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x9d, 0x44, 0x6e, 0xac,
	0xe0, 0x60, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67,
	0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x87, 0xa2, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x85, 0xce,
	0x6a, 0x19, 0x82, 0x01, 0x00, 0x00,
}