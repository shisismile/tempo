// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tempodb/backend/v1/v1.proto

package backend

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type BlockMeta struct {
	Version          string           `protobuf:"bytes,1,opt,name=version,proto3" json:"format"`
	BlockID          UUID             `protobuf:"bytes,2,opt,name=block_id,json=blockId,proto3,customtype=UUID" json:"blockID"`
	TenantID         string           `protobuf:"bytes,5,opt,name=tenant_id,json=tenantId,proto3" json:"tenantID"`
	StartTime        time.Time        `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3,stdtime" json:"startTime"`
	EndTime          time.Time        `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3,stdtime" json:"endTime"`
	TotalObjects     int64            `protobuf:"varint,8,opt,name=total_objects,json=totalObjects,proto3" json:"totalObjects"`
	Size_            uint64           `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	CompactionLevel  uint32           `protobuf:"varint,10,opt,name=compaction_level,json=compactionLevel,proto3" json:"compactionLevel"`
	Encoding         Encoding         `protobuf:"bytes,11,opt,name=encoding,proto3,customtype=Encoding" json:"encoding"`
	IndexPageSize    uint32           `protobuf:"varint,12,opt,name=index_page_size,json=indexPageSize,proto3" json:"indexPageSize"`
	TotalRecords     uint32           `protobuf:"varint,13,opt,name=total_records,json=totalRecords,proto3" json:"totalRecords"`
	DataEncoding     string           `protobuf:"bytes,14,opt,name=data_encoding,json=dataEncoding,proto3" json:"dataEncoding"`
	BloomShardCount  uint32           `protobuf:"varint,15,opt,name=bloom_shard_count,json=bloomShardCount,proto3" json:"bloomShards"`
	FooterSize       uint32           `protobuf:"varint,16,opt,name=footer_size,json=footerSize,proto3" json:"footerSize"`
	DedicatedColumns DedicatedColumns `protobuf:"bytes,17,opt,name=dedicated_columns,json=dedicatedColumns,proto3,customtype=DedicatedColumns" json:"dedicatedColumns,omitempty"`
	// repeated bytes dedicated_columns = 17 [(gogoproto.customtype) = "DedicatedColumn", (gogoproto.jsontag) = "dedicatedColumns,omitempty", (gogoproto.nullable) = false];
	ReplicationFactor uint32 `protobuf:"varint,18,opt,name=replication_factor,json=replicationFactor,proto3" json:"replicationFactor,omitempty"`
}

func (m *BlockMeta) Reset()         { *m = BlockMeta{} }
func (m *BlockMeta) String() string { return proto.CompactTextString(m) }
func (*BlockMeta) ProtoMessage()    {}
func (*BlockMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bc10ae735c1a340, []int{0}
}
func (m *BlockMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMeta.Merge(m, src)
}
func (m *BlockMeta) XXX_Size() int {
	return m.Size()
}
func (m *BlockMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMeta.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMeta proto.InternalMessageInfo

func (m *BlockMeta) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *BlockMeta) GetTenantID() string {
	if m != nil {
		return m.TenantID
	}
	return ""
}

func (m *BlockMeta) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *BlockMeta) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *BlockMeta) GetTotalObjects() int64 {
	if m != nil {
		return m.TotalObjects
	}
	return 0
}

func (m *BlockMeta) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *BlockMeta) GetCompactionLevel() uint32 {
	if m != nil {
		return m.CompactionLevel
	}
	return 0
}

func (m *BlockMeta) GetIndexPageSize() uint32 {
	if m != nil {
		return m.IndexPageSize
	}
	return 0
}

func (m *BlockMeta) GetTotalRecords() uint32 {
	if m != nil {
		return m.TotalRecords
	}
	return 0
}

func (m *BlockMeta) GetDataEncoding() string {
	if m != nil {
		return m.DataEncoding
	}
	return ""
}

func (m *BlockMeta) GetBloomShardCount() uint32 {
	if m != nil {
		return m.BloomShardCount
	}
	return 0
}

func (m *BlockMeta) GetFooterSize() uint32 {
	if m != nil {
		return m.FooterSize
	}
	return 0
}

func (m *BlockMeta) GetReplicationFactor() uint32 {
	if m != nil {
		return m.ReplicationFactor
	}
	return 0
}

type CompactedBlockMeta struct {
	BlockMeta     `protobuf:"bytes,1,opt,name=block_meta,json=blockMeta,proto3,embedded=block_meta" json:""`
	CompactedTime time.Time `protobuf:"bytes,2,opt,name=compacted_time,json=compactedTime,proto3,stdtime" json:"compactedTime"`
}

func (m *CompactedBlockMeta) Reset()         { *m = CompactedBlockMeta{} }
func (m *CompactedBlockMeta) String() string { return proto.CompactTextString(m) }
func (*CompactedBlockMeta) ProtoMessage()    {}
func (*CompactedBlockMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bc10ae735c1a340, []int{1}
}
func (m *CompactedBlockMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompactedBlockMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompactedBlockMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompactedBlockMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactedBlockMeta.Merge(m, src)
}
func (m *CompactedBlockMeta) XXX_Size() int {
	return m.Size()
}
func (m *CompactedBlockMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactedBlockMeta.DiscardUnknown(m)
}

var xxx_messageInfo_CompactedBlockMeta proto.InternalMessageInfo

func (m *CompactedBlockMeta) GetCompactedTime() time.Time {
	if m != nil {
		return m.CompactedTime
	}
	return time.Time{}
}

type TenantIndex struct {
	CreatedAt     time.Time             `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3,stdtime" json:"createdAt"`
	Meta          []*BlockMeta          `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta"`
	CompactedMeta []*CompactedBlockMeta `protobuf:"bytes,3,rep,name=compacted_meta,json=compactedMeta,proto3" json:"compacted"`
}

func (m *TenantIndex) Reset()         { *m = TenantIndex{} }
func (m *TenantIndex) String() string { return proto.CompactTextString(m) }
func (*TenantIndex) ProtoMessage()    {}
func (*TenantIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bc10ae735c1a340, []int{2}
}
func (m *TenantIndex) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TenantIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TenantIndex.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TenantIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TenantIndex.Merge(m, src)
}
func (m *TenantIndex) XXX_Size() int {
	return m.Size()
}
func (m *TenantIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_TenantIndex.DiscardUnknown(m)
}

var xxx_messageInfo_TenantIndex proto.InternalMessageInfo

func (m *TenantIndex) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *TenantIndex) GetMeta() []*BlockMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *TenantIndex) GetCompactedMeta() []*CompactedBlockMeta {
	if m != nil {
		return m.CompactedMeta
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockMeta)(nil), "backend.v1.BlockMeta")
	proto.RegisterType((*CompactedBlockMeta)(nil), "backend.v1.CompactedBlockMeta")
	proto.RegisterType((*TenantIndex)(nil), "backend.v1.TenantIndex")
}

func init() { proto.RegisterFile("tempodb/backend/v1/v1.proto", fileDescriptor_6bc10ae735c1a340) }

var fileDescriptor_6bc10ae735c1a340 = []byte{
	// 787 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x8a, 0xdb, 0x46,
	0x14, 0xb6, 0x76, 0xb7, 0x6b, 0x7b, 0x6c, 0xaf, 0xed, 0x09, 0x01, 0xd5, 0x01, 0x8f, 0x31, 0xbd,
	0x70, 0x21, 0x95, 0xd9, 0x84, 0x14, 0x4a, 0x69, 0xa1, 0xda, 0x6d, 0x21, 0xa5, 0x3f, 0x41, 0xd9,
	0xdc, 0x94, 0x82, 0x18, 0x69, 0xc6, 0x8a, 0x1a, 0x49, 0x63, 0xa4, 0xb1, 0x69, 0xf3, 0x14, 0x79,
	0x9a, 0x3e, 0x43, 0xe8, 0xd5, 0x5e, 0x96, 0x5e, 0x4c, 0x8b, 0xf7, 0x4e, 0x4f, 0x51, 0xe6, 0xe8,
	0xcf, 0x76, 0x28, 0xe9, 0xcd, 0x72, 0xce, 0xf9, 0xce, 0x77, 0xe6, 0x7c, 0xa3, 0xf9, 0xd6, 0xe8,
	0x81, 0xe4, 0xf1, 0x5a, 0x30, 0x6f, 0xe9, 0x51, 0xff, 0x15, 0x4f, 0xd8, 0x72, 0x7b, 0xb9, 0xdc,
	0x5e, 0x5a, 0xeb, 0x54, 0x48, 0x81, 0x51, 0x59, 0xb4, 0xb6, 0x97, 0x13, 0x12, 0x08, 0x11, 0x44,
	0x7c, 0x09, 0x88, 0xb7, 0x59, 0x2d, 0x65, 0x18, 0xf3, 0x4c, 0xd2, 0x78, 0x5d, 0x34, 0x4f, 0x3e,
	0x09, 0x42, 0xf9, 0x72, 0xe3, 0x59, 0xbe, 0x88, 0x97, 0x81, 0x08, 0x44, 0xd3, 0xa9, 0x33, 0x48,
	0x20, 0x2a, 0xda, 0xe7, 0x7f, 0xb4, 0x51, 0xd7, 0x8e, 0x84, 0xff, 0xea, 0x7b, 0x2e, 0x29, 0xfe,
	0x08, 0xb5, 0xb7, 0x3c, 0xcd, 0x42, 0x91, 0x98, 0xc6, 0xcc, 0x58, 0x74, 0x6d, 0x94, 0x2b, 0x72,
	0xbe, 0x12, 0x69, 0x4c, 0xa5, 0x53, 0x41, 0xf8, 0x0b, 0xd4, 0xf1, 0x34, 0xc5, 0x0d, 0x99, 0x79,
	0x32, 0x33, 0x16, 0x7d, 0x7b, 0xfe, 0x56, 0x91, 0xd6, 0x5f, 0x8a, 0x9c, 0xbd, 0x78, 0xf1, 0xf4,
	0x7a, 0xa7, 0x48, 0x1b, 0x46, 0x3e, 0xbd, 0xce, 0x15, 0x69, 0x7b, 0x45, 0xe8, 0x94, 0x01, 0xc3,
	0x4f, 0x50, 0x57, 0xf2, 0x84, 0x26, 0x52, 0xf3, 0x3f, 0x80, 0x63, 0xcc, 0x9d, 0x22, 0x9d, 0x1b,
	0x28, 0x02, 0xa9, 0x23, 0xcb, 0xd8, 0xa9, 0x22, 0x86, 0x9f, 0x21, 0x94, 0x49, 0x9a, 0x4a, 0x57,
	0x2b, 0x36, 0xcf, 0x67, 0xc6, 0xa2, 0xf7, 0x68, 0x62, 0x15, 0xd7, 0x61, 0x55, 0x22, 0xad, 0x9b,
	0xea, 0x3a, 0xec, 0xfb, 0x7a, 0xa7, 0x5c, 0x91, 0x2e, 0xb0, 0x74, 0xfd, 0xcd, 0xdf, 0xc4, 0x70,
	0x9a, 0x14, 0x7f, 0x8b, 0x3a, 0x3c, 0x61, 0xc5, 0xbc, 0xf6, 0x7b, 0xe7, 0xdd, 0x2b, 0xe7, 0xb5,
	0x79, 0xc2, 0xea, 0x69, 0x55, 0x82, 0x9f, 0xa0, 0x81, 0x14, 0x92, 0x46, 0xae, 0xf0, 0x7e, 0xe1,
	0xbe, 0xcc, 0xcc, 0xce, 0xcc, 0x58, 0x9c, 0xda, 0xa3, 0x5c, 0x91, 0x3e, 0x00, 0x3f, 0x16, 0x75,
	0xe7, 0x20, 0xc3, 0x18, 0x9d, 0x65, 0xe1, 0x6b, 0x6e, 0x76, 0x67, 0xc6, 0xe2, 0xcc, 0x81, 0x18,
	0x7f, 0x89, 0x46, 0xbe, 0x88, 0xd7, 0xd4, 0x97, 0xa1, 0x48, 0xdc, 0x88, 0x6f, 0x79, 0x64, 0xa2,
	0x99, 0xb1, 0x18, 0xd8, 0xf7, 0x72, 0x45, 0x86, 0x0d, 0xf6, 0x9d, 0x86, 0x9c, 0xe3, 0x02, 0x7e,
	0xa8, 0x65, 0xf9, 0x82, 0x85, 0x49, 0x60, 0xf6, 0xe0, 0xf3, 0x8c, 0xca, 0xcf, 0xd3, 0xf9, 0xba,
	0xac, 0x3b, 0x75, 0x07, 0xfe, 0x0c, 0x0d, 0xc3, 0x84, 0xf1, 0x5f, 0xdd, 0x35, 0x0d, 0xb8, 0x0b,
	0xcb, 0xf4, 0xe1, 0xb0, 0x71, 0xae, 0xc8, 0x00, 0xa0, 0x67, 0x34, 0xe0, 0xcf, 0xc3, 0xd7, 0xdc,
	0x39, 0x4c, 0x1b, 0xcd, 0x29, 0xf7, 0x45, 0xca, 0x32, 0x73, 0x00, 0xc4, 0x46, 0xb3, 0x53, 0xd4,
	0x9d, 0x83, 0x4c, 0xd3, 0x18, 0x95, 0xd4, 0xad, 0x97, 0xbc, 0x80, 0x37, 0x00, 0x34, 0x0d, 0xd4,
	0x4b, 0x1e, 0x64, 0xf8, 0x73, 0x34, 0xf6, 0x22, 0x21, 0x62, 0x37, 0x7b, 0x49, 0x53, 0xe6, 0xfa,
	0x62, 0x93, 0x48, 0x73, 0x08, 0x27, 0x0e, 0x73, 0x45, 0x7a, 0x00, 0x3e, 0xd7, 0x58, 0xe6, 0x0c,
	0x9b, 0xe4, 0x4a, 0xf7, 0xe1, 0x25, 0xea, 0xad, 0x84, 0x90, 0x3c, 0x2d, 0x14, 0x8e, 0x80, 0x76,
	0x91, 0x2b, 0x82, 0x8a, 0x32, 0xc8, 0xdb, 0x8b, 0xb1, 0x8f, 0xc6, 0x8c, 0xb3, 0xd0, 0xa7, 0x92,
	0xeb, 0xb3, 0xa2, 0x4d, 0x9c, 0x64, 0xe6, 0x18, 0x6e, 0xf3, 0xd3, 0xf2, 0x36, 0x47, 0xd7, 0x55,
	0xc3, 0x55, 0x81, 0xe7, 0x8a, 0x4c, 0xd8, 0x51, 0xed, 0xa1, 0x88, 0x43, 0xed, 0x6d, 0xf9, 0x9b,
	0x33, 0x3a, 0xc6, 0xf0, 0x0f, 0x08, 0xa7, 0x7c, 0x1d, 0xe9, 0xa2, 0xfe, 0xd4, 0x2b, 0xea, 0x4b,
	0x91, 0x9a, 0x18, 0x96, 0x23, 0xb9, 0x22, 0x0f, 0xf6, 0xd0, 0x6f, 0x00, 0xdc, 0x1b, 0x37, 0x7e,
	0x07, 0x9c, 0xff, 0x6e, 0x20, 0x7c, 0x55, 0xbc, 0x06, 0xce, 0x1a, 0x57, 0xdb, 0x08, 0x15, 0x7e,
	0x8d, 0xb9, 0xa4, 0x60, 0xec, 0xde, 0xa3, 0xfb, 0x56, 0xf3, 0x4f, 0xc5, 0xaa, 0x5b, 0xed, 0xbe,
	0xd6, 0x76, 0xab, 0x88, 0x91, 0x2b, 0xd2, 0x72, 0xba, 0x5e, 0x3d, 0xe3, 0x67, 0x74, 0xe1, 0x57,
	0x93, 0x0b, 0xc7, 0x9c, 0xbc, 0xd7, 0x31, 0x1f, 0x96, 0x8e, 0x19, 0xd4, 0xcc, 0xda, 0x37, 0x87,
	0xa5, 0x79, 0x6e, 0xa0, 0x5e, 0x69, 0x7f, 0xfd, 0xc2, 0xb4, 0xd7, 0xfd, 0x94, 0xc3, 0xdd, 0x53,
	0x59, 0x6e, 0xfc, 0xbf, 0xbc, 0x5e, 0xb2, 0xbe, 0x92, 0x85, 0xd7, 0xeb, 0x14, 0x3f, 0x46, 0x67,
	0xa0, 0xfe, 0x64, 0x76, 0xfa, 0xdf, 0xea, 0x3b, 0xb9, 0x22, 0xd0, 0xe6, 0xc0, 0x5f, 0x7c, 0xb3,
	0x2f, 0x1a, 0xe8, 0xa7, 0x40, 0x9f, 0xee, 0xd3, 0xdf, 0xbd, 0x70, 0x7b, 0x00, 0xab, 0x54, 0xf5,
	0x3d, 0xb1, 0x80, 0x7e, 0xfc, 0x76, 0x37, 0x35, 0x6e, 0x77, 0x53, 0xe3, 0x9f, 0xdd, 0xd4, 0x78,
	0x73, 0x37, 0x6d, 0xdd, 0xde, 0x4d, 0x5b, 0x7f, 0xde, 0x4d, 0x5b, 0x3f, 0x0d, 0x8f, 0x7e, 0x05,
	0xbc, 0x73, 0xd0, 0xfa, 0xf8, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0x8d, 0x86, 0x48, 0x1f,
	0x06, 0x00, 0x00,
}

func (m *BlockMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReplicationFactor != 0 {
		i = encodeVarintV1(dAtA, i, uint64(m.ReplicationFactor))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	{
		size := m.DedicatedColumns.Size()
		i -= size
		if _, err := m.DedicatedColumns.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintV1(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x8a
	if m.FooterSize != 0 {
		i = encodeVarintV1(dAtA, i, uint64(m.FooterSize))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.BloomShardCount != 0 {
		i = encodeVarintV1(dAtA, i, uint64(m.BloomShardCount))
		i--
		dAtA[i] = 0x78
	}
	if len(m.DataEncoding) > 0 {
		i -= len(m.DataEncoding)
		copy(dAtA[i:], m.DataEncoding)
		i = encodeVarintV1(dAtA, i, uint64(len(m.DataEncoding)))
		i--
		dAtA[i] = 0x72
	}
	if m.TotalRecords != 0 {
		i = encodeVarintV1(dAtA, i, uint64(m.TotalRecords))
		i--
		dAtA[i] = 0x68
	}
	if m.IndexPageSize != 0 {
		i = encodeVarintV1(dAtA, i, uint64(m.IndexPageSize))
		i--
		dAtA[i] = 0x60
	}
	{
		size := m.Encoding.Size()
		i -= size
		if _, err := m.Encoding.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintV1(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if m.CompactionLevel != 0 {
		i = encodeVarintV1(dAtA, i, uint64(m.CompactionLevel))
		i--
		dAtA[i] = 0x50
	}
	if m.Size_ != 0 {
		i = encodeVarintV1(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x48
	}
	if m.TotalObjects != 0 {
		i = encodeVarintV1(dAtA, i, uint64(m.TotalObjects))
		i--
		dAtA[i] = 0x40
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintV1(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintV1(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x32
	if len(m.TenantID) > 0 {
		i -= len(m.TenantID)
		copy(dAtA[i:], m.TenantID)
		i = encodeVarintV1(dAtA, i, uint64(len(m.TenantID)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.BlockID.Size()
		i -= size
		if _, err := m.BlockID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintV1(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintV1(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CompactedBlockMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompactedBlockMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompactedBlockMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CompactedTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CompactedTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintV1(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.BlockMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintV1(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TenantIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TenantIndex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TenantIndex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CompactedMeta) > 0 {
		for iNdEx := len(m.CompactedMeta) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CompactedMeta[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintV1(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Meta) > 0 {
		for iNdEx := len(m.Meta) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Meta[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintV1(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	n5, err5 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintV1(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovV1(uint64(l))
	}
	l = m.BlockID.Size()
	n += 1 + l + sovV1(uint64(l))
	l = len(m.TenantID)
	if l > 0 {
		n += 1 + l + sovV1(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovV1(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovV1(uint64(l))
	if m.TotalObjects != 0 {
		n += 1 + sovV1(uint64(m.TotalObjects))
	}
	if m.Size_ != 0 {
		n += 1 + sovV1(uint64(m.Size_))
	}
	if m.CompactionLevel != 0 {
		n += 1 + sovV1(uint64(m.CompactionLevel))
	}
	l = m.Encoding.Size()
	n += 1 + l + sovV1(uint64(l))
	if m.IndexPageSize != 0 {
		n += 1 + sovV1(uint64(m.IndexPageSize))
	}
	if m.TotalRecords != 0 {
		n += 1 + sovV1(uint64(m.TotalRecords))
	}
	l = len(m.DataEncoding)
	if l > 0 {
		n += 1 + l + sovV1(uint64(l))
	}
	if m.BloomShardCount != 0 {
		n += 1 + sovV1(uint64(m.BloomShardCount))
	}
	if m.FooterSize != 0 {
		n += 2 + sovV1(uint64(m.FooterSize))
	}
	l = m.DedicatedColumns.Size()
	n += 2 + l + sovV1(uint64(l))
	if m.ReplicationFactor != 0 {
		n += 2 + sovV1(uint64(m.ReplicationFactor))
	}
	return n
}

func (m *CompactedBlockMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BlockMeta.Size()
	n += 1 + l + sovV1(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CompactedTime)
	n += 1 + l + sovV1(uint64(l))
	return n
}

func (m *TenantIndex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovV1(uint64(l))
	if len(m.Meta) > 0 {
		for _, e := range m.Meta {
			l = e.Size()
			n += 1 + l + sovV1(uint64(l))
		}
	}
	if len(m.CompactedMeta) > 0 {
		for _, e := range m.CompactedMeta {
			l = e.Size()
			n += 1 + l + sovV1(uint64(l))
		}
	}
	return n
}

func sovV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozV1(x uint64) (n int) {
	return sovV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV1
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
			return fmt.Errorf("proto: BlockMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TenantID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TenantID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalObjects", wireType)
			}
			m.TotalObjects = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalObjects |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompactionLevel", wireType)
			}
			m.CompactionLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompactionLevel |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encoding", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Encoding.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexPageSize", wireType)
			}
			m.IndexPageSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexPageSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalRecords", wireType)
			}
			m.TotalRecords = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalRecords |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataEncoding", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataEncoding = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BloomShardCount", wireType)
			}
			m.BloomShardCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BloomShardCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FooterSize", wireType)
			}
			m.FooterSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FooterSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DedicatedColumns", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DedicatedColumns.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicationFactor", wireType)
			}
			m.ReplicationFactor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplicationFactor |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthV1
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
func (m *CompactedBlockMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV1
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
			return fmt.Errorf("proto: CompactedBlockMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompactedBlockMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompactedTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CompactedTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthV1
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
func (m *TenantIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV1
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
			return fmt.Errorf("proto: TenantIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TenantIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = append(m.Meta, &BlockMeta{})
			if err := m.Meta[len(m.Meta)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompactedMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompactedMeta = append(m.CompactedMeta, &CompactedBlockMeta{})
			if err := m.CompactedMeta[len(m.CompactedMeta)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthV1
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
func skipV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowV1
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
					return 0, ErrIntOverflowV1
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
					return 0, ErrIntOverflowV1
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
				return 0, ErrInvalidLengthV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupV1 = fmt.Errorf("proto: unexpected end of group")
)