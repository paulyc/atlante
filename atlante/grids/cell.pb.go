// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cell.proto

package grids

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type HEMIType int32

const (
	HEMIType_NORTH HEMIType = 0
	HEMIType_SOUTH HEMIType = 1
)

var HEMIType_name = map[int32]string{
	0: "NORTH",
	1: "SOUTH",
}

var HEMIType_value = map[string]int32{
	"NORTH": 0,
	"SOUTH": 1,
}

func (x HEMIType) String() string {
	return proto.EnumName(HEMIType_name, int32(x))
}

func (HEMIType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d2698cf62ebc0420, []int{0}
}

type MDGID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Part                 uint32   `protobuf:"varint,2,opt,name=part,proto3" json:"part,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MDGID) Reset()         { *m = MDGID{} }
func (m *MDGID) String() string { return proto.CompactTextString(m) }
func (*MDGID) ProtoMessage()    {}
func (*MDGID) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2698cf62ebc0420, []int{0}
}

func (m *MDGID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MDGID.Unmarshal(m, b)
}
func (m *MDGID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MDGID.Marshal(b, m, deterministic)
}
func (m *MDGID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MDGID.Merge(m, src)
}
func (m *MDGID) XXX_Size() int {
	return xxx_messageInfo_MDGID.Size(m)
}
func (m *MDGID) XXX_DiscardUnknown() {
	xxx_messageInfo_MDGID.DiscardUnknown(m)
}

var xxx_messageInfo_MDGID proto.InternalMessageInfo

func (m *MDGID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MDGID) GetPart() uint32 {
	if m != nil {
		return m.Part
	}
	return 0
}

type EditInfo struct {
	By                   string               `protobuf:"bytes,1,opt,name=by,proto3" json:"by,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EditInfo) Reset()         { *m = EditInfo{} }
func (m *EditInfo) String() string { return proto.CompactTextString(m) }
func (*EditInfo) ProtoMessage()    {}
func (*EditInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2698cf62ebc0420, []int{1}
}

func (m *EditInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditInfo.Unmarshal(m, b)
}
func (m *EditInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditInfo.Marshal(b, m, deterministic)
}
func (m *EditInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditInfo.Merge(m, src)
}
func (m *EditInfo) XXX_Size() int {
	return xxx_messageInfo_EditInfo.Size(m)
}
func (m *EditInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EditInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EditInfo proto.InternalMessageInfo

func (m *EditInfo) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

func (m *EditInfo) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type UTMInfo struct {
	Zone uint32 `protobuf:"varint,1,opt,name=zone,proto3" json:"zone,omitempty"`
	// should be formated at 01; 01-60
	Hemi                 HEMIType `protobuf:"varint,2,opt,name=hemi,proto3,enum=atlante.grids.HEMIType" json:"hemi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UTMInfo) Reset()         { *m = UTMInfo{} }
func (m *UTMInfo) String() string { return proto.CompactTextString(m) }
func (*UTMInfo) ProtoMessage()    {}
func (*UTMInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2698cf62ebc0420, []int{2}
}

func (m *UTMInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UTMInfo.Unmarshal(m, b)
}
func (m *UTMInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UTMInfo.Marshal(b, m, deterministic)
}
func (m *UTMInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UTMInfo.Merge(m, src)
}
func (m *UTMInfo) XXX_Size() int {
	return xxx_messageInfo_UTMInfo.Size(m)
}
func (m *UTMInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UTMInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UTMInfo proto.InternalMessageInfo

func (m *UTMInfo) GetZone() uint32 {
	if m != nil {
		return m.Zone
	}
	return 0
}

func (m *UTMInfo) GetHemi() HEMIType {
	if m != nil {
		return m.Hemi
	}
	return HEMIType_NORTH
}

type Cell struct {
	Mdgid                *MDGID               `protobuf:"bytes,1,opt,name=mdgid,proto3" json:"mdgid,omitempty"`
	Sw                   *Cell_LatLng         `protobuf:"bytes,2,opt,name=sw,proto3" json:"sw,omitempty"`
	Ne                   *Cell_LatLng         `protobuf:"bytes,3,opt,name=ne,proto3" json:"ne,omitempty"`
	Len                  *Cell_LatLng         `protobuf:"bytes,4,opt,name=len,proto3" json:"len,omitempty"`
	Country              string               `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	City                 string               `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	Utm                  *UTMInfo             `protobuf:"bytes,7,opt,name=utm,proto3" json:"utm,omitempty"`
	Edited               *EditInfo            `protobuf:"bytes,8,opt,name=edited,proto3" json:"edited,omitempty"`
	PublishedAt          *timestamp.Timestamp `protobuf:"bytes,10,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
	MetaData             map[string]string    `protobuf:"bytes,11,rep,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Nrn                  string               `protobuf:"bytes,12,opt,name=nrn,proto3" json:"nrn,omitempty"`
	Sheet                string               `protobuf:"bytes,13,opt,name=sheet,proto3" json:"sheet,omitempty"`
	Series               string               `protobuf:"bytes,14,opt,name=series,proto3" json:"series,omitempty"`
	SwDms                *Cell_LatLngDMS      `protobuf:"bytes,15,opt,name=sw_dms,json=swDms,proto3" json:"sw_dms,omitempty"`
	NeDms                *Cell_LatLngDMS      `protobuf:"bytes,16,opt,name=ne_dms,json=neDms,proto3" json:"ne_dms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Cell) Reset()         { *m = Cell{} }
func (m *Cell) String() string { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()    {}
func (*Cell) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2698cf62ebc0420, []int{3}
}

func (m *Cell) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cell.Unmarshal(m, b)
}
func (m *Cell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cell.Marshal(b, m, deterministic)
}
func (m *Cell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cell.Merge(m, src)
}
func (m *Cell) XXX_Size() int {
	return xxx_messageInfo_Cell.Size(m)
}
func (m *Cell) XXX_DiscardUnknown() {
	xxx_messageInfo_Cell.DiscardUnknown(m)
}

var xxx_messageInfo_Cell proto.InternalMessageInfo

func (m *Cell) GetMdgid() *MDGID {
	if m != nil {
		return m.Mdgid
	}
	return nil
}

func (m *Cell) GetSw() *Cell_LatLng {
	if m != nil {
		return m.Sw
	}
	return nil
}

func (m *Cell) GetNe() *Cell_LatLng {
	if m != nil {
		return m.Ne
	}
	return nil
}

func (m *Cell) GetLen() *Cell_LatLng {
	if m != nil {
		return m.Len
	}
	return nil
}

func (m *Cell) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Cell) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Cell) GetUtm() *UTMInfo {
	if m != nil {
		return m.Utm
	}
	return nil
}

func (m *Cell) GetEdited() *EditInfo {
	if m != nil {
		return m.Edited
	}
	return nil
}

func (m *Cell) GetPublishedAt() *timestamp.Timestamp {
	if m != nil {
		return m.PublishedAt
	}
	return nil
}

func (m *Cell) GetMetaData() map[string]string {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func (m *Cell) GetNrn() string {
	if m != nil {
		return m.Nrn
	}
	return ""
}

func (m *Cell) GetSheet() string {
	if m != nil {
		return m.Sheet
	}
	return ""
}

func (m *Cell) GetSeries() string {
	if m != nil {
		return m.Series
	}
	return ""
}

func (m *Cell) GetSwDms() *Cell_LatLngDMS {
	if m != nil {
		return m.SwDms
	}
	return nil
}

func (m *Cell) GetNeDms() *Cell_LatLngDMS {
	if m != nil {
		return m.NeDms
	}
	return nil
}

type Cell_LatLng struct {
	Lat                  float32  `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  float32  `protobuf:"fixed32,2,opt,name=lng,proto3" json:"lng,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cell_LatLng) Reset()         { *m = Cell_LatLng{} }
func (m *Cell_LatLng) String() string { return proto.CompactTextString(m) }
func (*Cell_LatLng) ProtoMessage()    {}
func (*Cell_LatLng) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2698cf62ebc0420, []int{3, 0}
}

func (m *Cell_LatLng) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cell_LatLng.Unmarshal(m, b)
}
func (m *Cell_LatLng) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cell_LatLng.Marshal(b, m, deterministic)
}
func (m *Cell_LatLng) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cell_LatLng.Merge(m, src)
}
func (m *Cell_LatLng) XXX_Size() int {
	return xxx_messageInfo_Cell_LatLng.Size(m)
}
func (m *Cell_LatLng) XXX_DiscardUnknown() {
	xxx_messageInfo_Cell_LatLng.DiscardUnknown(m)
}

var xxx_messageInfo_Cell_LatLng proto.InternalMessageInfo

func (m *Cell_LatLng) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Cell_LatLng) GetLng() float32 {
	if m != nil {
		return m.Lng
	}
	return 0
}

type Cell_LatLngDMS struct {
	Lat                  string   `protobuf:"bytes,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  string   `protobuf:"bytes,2,opt,name=lng,proto3" json:"lng,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cell_LatLngDMS) Reset()         { *m = Cell_LatLngDMS{} }
func (m *Cell_LatLngDMS) String() string { return proto.CompactTextString(m) }
func (*Cell_LatLngDMS) ProtoMessage()    {}
func (*Cell_LatLngDMS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2698cf62ebc0420, []int{3, 1}
}

func (m *Cell_LatLngDMS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cell_LatLngDMS.Unmarshal(m, b)
}
func (m *Cell_LatLngDMS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cell_LatLngDMS.Marshal(b, m, deterministic)
}
func (m *Cell_LatLngDMS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cell_LatLngDMS.Merge(m, src)
}
func (m *Cell_LatLngDMS) XXX_Size() int {
	return xxx_messageInfo_Cell_LatLngDMS.Size(m)
}
func (m *Cell_LatLngDMS) XXX_DiscardUnknown() {
	xxx_messageInfo_Cell_LatLngDMS.DiscardUnknown(m)
}

var xxx_messageInfo_Cell_LatLngDMS proto.InternalMessageInfo

func (m *Cell_LatLngDMS) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *Cell_LatLngDMS) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func init() {
	proto.RegisterEnum("atlante.grids.HEMIType", HEMIType_name, HEMIType_value)
	proto.RegisterType((*MDGID)(nil), "atlante.grids.MDGID")
	proto.RegisterType((*EditInfo)(nil), "atlante.grids.EditInfo")
	proto.RegisterType((*UTMInfo)(nil), "atlante.grids.UTMInfo")
	proto.RegisterType((*Cell)(nil), "atlante.grids.Cell")
	proto.RegisterMapType((map[string]string)(nil), "atlante.grids.Cell.MetaDataEntry")
	proto.RegisterType((*Cell_LatLng)(nil), "atlante.grids.Cell.LatLng")
	proto.RegisterType((*Cell_LatLngDMS)(nil), "atlante.grids.Cell.LatLngDMS")
}

func init() { proto.RegisterFile("cell.proto", fileDescriptor_d2698cf62ebc0420) }

var fileDescriptor_d2698cf62ebc0420 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0xe9, 0xaf, 0xac, 0xbd, 0xae, 0xa3, 0xb2, 0xa6, 0x61, 0x55, 0x42, 0x94, 0x3e, 0x55,
	0xdb, 0x48, 0x50, 0xe1, 0x01, 0x81, 0x40, 0x02, 0x3a, 0xb1, 0x4d, 0x2b, 0x93, 0xbc, 0xee, 0x85,
	0x97, 0xc9, 0x6d, 0xbc, 0xd4, 0x5a, 0xe2, 0x44, 0xf5, 0x95, 0xaa, 0xfc, 0x05, 0xfc, 0xd9, 0xc8,
	0x97, 0x64, 0x62, 0x03, 0x34, 0xde, 0xee, 0x2e, 0x9f, 0xef, 0xd9, 0xf7, 0xbd, 0x18, 0x60, 0xae,
	0xe2, 0xd8, 0xcf, 0x96, 0x29, 0xa6, 0xac, 0x23, 0x31, 0x96, 0x06, 0x95, 0x1f, 0x2d, 0x75, 0x68,
	0x7b, 0xcf, 0xa2, 0x34, 0x8d, 0x62, 0x15, 0xd0, 0xc7, 0xd9, 0xea, 0x3a, 0x40, 0x9d, 0x28, 0x8b,
	0x32, 0xc9, 0x72, 0x7e, 0x70, 0x00, 0x8d, 0xc9, 0xf8, 0xcb, 0xc9, 0x98, 0xed, 0x40, 0x55, 0x87,
	0xbc, 0xd2, 0xaf, 0x0c, 0x5b, 0xa2, 0xaa, 0x43, 0xc6, 0xa0, 0x9e, 0xc9, 0x25, 0xf2, 0x6a, 0xbf,
	0x32, 0xec, 0x08, 0x8a, 0x07, 0xa7, 0xd0, 0x3c, 0x0a, 0x35, 0x9e, 0x98, 0xeb, 0xd4, 0xf1, 0xb3,
	0x4d, 0xc9, 0xcf, 0x36, 0xcc, 0x87, 0x7a, 0x28, 0x51, 0x11, 0xdf, 0x1e, 0xf5, 0xfc, 0xfc, 0x60,
	0xbf, 0x3c, 0xd8, 0x9f, 0x96, 0x07, 0x0b, 0xe2, 0x06, 0xa7, 0xb0, 0x75, 0x39, 0x9d, 0x50, 0x2b,
	0x06, 0xf5, 0x1f, 0xa9, 0x51, 0xd4, 0xac, 0x23, 0x28, 0x66, 0x07, 0x50, 0x5f, 0xa8, 0x44, 0x53,
	0xbb, 0x9d, 0xd1, 0x13, 0xff, 0xce, 0x58, 0xfe, 0xf1, 0xd1, 0xe4, 0x64, 0xba, 0xc9, 0x94, 0x20,
	0x68, 0xf0, 0xd3, 0x83, 0xfa, 0x67, 0x15, 0xc7, 0x6c, 0x1f, 0x1a, 0x49, 0x18, 0x15, 0x73, 0xb4,
	0x47, 0xbb, 0xf7, 0x64, 0x34, 0xa9, 0xc8, 0x11, 0xb6, 0x0f, 0x55, 0xbb, 0xbe, 0xbd, 0xee, 0x5d,
	0xd0, 0x35, 0xf3, 0xcf, 0x24, 0x9e, 0x99, 0x48, 0x54, 0xed, 0xda, 0xb1, 0x46, 0xf1, 0xda, 0xc3,
	0xac, 0x51, 0xec, 0x10, 0x6a, 0xb1, 0x32, 0xbc, 0xfe, 0x20, 0xec, 0x30, 0xc6, 0x61, 0x6b, 0x9e,
	0xae, 0x0c, 0x2e, 0x37, 0xbc, 0x41, 0x5e, 0x96, 0xa9, 0x73, 0x65, 0xae, 0x71, 0xc3, 0x3d, 0x2a,
	0x53, 0xcc, 0x86, 0x50, 0x5b, 0x61, 0xc2, 0xb7, 0xa8, 0xf7, 0xde, 0xbd, 0xde, 0x85, 0x9d, 0xc2,
	0x21, 0x2c, 0x00, 0x4f, 0x85, 0x1a, 0x55, 0xc8, 0x9b, 0x04, 0xdf, 0x77, 0xb0, 0xdc, 0xa3, 0x28,
	0x30, 0xf6, 0x1e, 0xb6, 0xb3, 0xd5, 0x2c, 0xd6, 0x76, 0xa1, 0xc2, 0x2b, 0x89, 0x1c, 0x1e, 0xdc,
	0x63, 0xfb, 0x96, 0xff, 0x88, 0xec, 0x03, 0xb4, 0x12, 0x85, 0xf2, 0x2a, 0x94, 0x28, 0x79, 0xbb,
	0x5f, 0x1b, 0xb6, 0x47, 0xcf, 0xff, 0x36, 0xfb, 0x44, 0xa1, 0x1c, 0x4b, 0x94, 0x47, 0x6e, 0x46,
	0xd1, 0x4c, 0x8a, 0x94, 0x75, 0xa1, 0x66, 0x96, 0x86, 0x6f, 0xd3, 0xb0, 0x2e, 0x64, 0xbb, 0xd0,
	0xb0, 0x0b, 0xa5, 0x90, 0x77, 0xa8, 0x96, 0x27, 0x6c, 0x0f, 0x3c, 0xab, 0x96, 0x5a, 0x59, 0xbe,
	0x43, 0xe5, 0x22, 0x63, 0xaf, 0xc1, 0xb3, 0xeb, 0xab, 0x30, 0xb1, 0xfc, 0x31, 0x5d, 0xfc, 0xe9,
	0xbf, 0x8d, 0x1f, 0x4f, 0x2e, 0x44, 0xc3, 0xae, 0xc7, 0x09, 0xa9, 0x8c, 0x22, 0x55, 0xf7, 0xbf,
	0x54, 0x46, 0x8d, 0x13, 0xdb, 0x3b, 0x04, 0x2f, 0xaf, 0xb9, 0x5b, 0xc7, 0x12, 0xe9, 0x6f, 0xab,
	0x0a, 0x17, 0x52, 0xc5, 0x44, 0xf4, 0x5b, 0xb9, 0x8a, 0x89, 0x7a, 0x01, 0xb4, 0x6e, 0x3b, 0xfc,
	0x2e, 0x68, 0xfd, 0x21, 0x68, 0xe5, 0x82, 0x77, 0xd0, 0xb9, 0xe3, 0x92, 0x43, 0x6e, 0x54, 0xf9,
	0xd6, 0x5c, 0xe8, 0xbc, 0xf9, 0x2e, 0xe3, 0x95, 0x2a, 0x64, 0x79, 0xf2, 0xb6, 0xfa, 0xa6, 0xb2,
	0xdf, 0x87, 0x66, 0xf9, 0x38, 0x58, 0x0b, 0x1a, 0x5f, 0xcf, 0xc5, 0xf4, 0xb8, 0xfb, 0xc8, 0x85,
	0x17, 0xe7, 0x97, 0xd3, 0xe3, 0x6e, 0xe5, 0xd3, 0xcb, 0x6f, 0x7e, 0xa4, 0x71, 0xb1, 0x9a, 0xf9,
	0xf3, 0x34, 0x09, 0xa2, 0xf4, 0x85, 0xcd, 0x24, 0x6a, 0x19, 0x07, 0x89, 0xcc, 0x30, 0x4d, 0xe3,
	0x1b, 0x8d, 0x41, 0xe1, 0x42, 0x40, 0x2e, 0xcc, 0x3c, 0x5a, 0xfe, 0xab, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x6d, 0x24, 0x53, 0x48, 0x68, 0x04, 0x00, 0x00,
}
