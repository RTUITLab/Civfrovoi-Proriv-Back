// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: src/protos/coordservice/coordservice.proto

package coordservice

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Resuource int32

const (
	Resuource_TRUCK                      Resuource = 0
	Resuource_ECAVATOR                   Resuource = 1
	Resuource_ROTARY_LOADER              Resuource = 2
	Resuource_PLOW_AND_BRUSH_MACHINE     Resuource = 3
	Resuource_PLOW_AND_BRUSH_MACHINE_REG Resuource = 4
	Resuource_HUMAN                      Resuource = 5
)

// Enum value maps for Resuource.
var (
	Resuource_name = map[int32]string{
		0: "TRUCK",
		1: "ECAVATOR",
		2: "ROTARY_LOADER",
		3: "PLOW_AND_BRUSH_MACHINE",
		4: "PLOW_AND_BRUSH_MACHINE_REG",
		5: "HUMAN",
	}
	Resuource_value = map[string]int32{
		"TRUCK":                      0,
		"ECAVATOR":                   1,
		"ROTARY_LOADER":              2,
		"PLOW_AND_BRUSH_MACHINE":     3,
		"PLOW_AND_BRUSH_MACHINE_REG": 4,
		"HUMAN":                      5,
	}
)

func (x Resuource) Enum() *Resuource {
	p := new(Resuource)
	*p = x
	return p
}

func (x Resuource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Resuource) Descriptor() protoreflect.EnumDescriptor {
	return file_src_protos_coordservice_coordservice_proto_enumTypes[0].Descriptor()
}

func (Resuource) Type() protoreflect.EnumType {
	return &file_src_protos_coordservice_coordservice_proto_enumTypes[0]
}

func (x Resuource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Resuource.Descriptor instead.
func (Resuource) EnumDescriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{0}
}

type WriteCoordsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Lat  float64 `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Long float64 `protobuf:"fixed64,3,opt,name=long,proto3" json:"long,omitempty"`
}

func (x *WriteCoordsReq) Reset() {
	*x = WriteCoordsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteCoordsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteCoordsReq) ProtoMessage() {}

func (x *WriteCoordsReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteCoordsReq.ProtoReflect.Descriptor instead.
func (*WriteCoordsReq) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{0}
}

func (x *WriteCoordsReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WriteCoordsReq) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *WriteCoordsReq) GetLong() float64 {
	if x != nil {
		return x.Long
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{1}
}

type WriteCoordsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteCoordsResp) Reset() {
	*x = WriteCoordsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteCoordsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteCoordsResp) ProtoMessage() {}

func (x *WriteCoordsResp) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteCoordsResp.ProtoReflect.Descriptor instead.
func (*WriteCoordsResp) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{2}
}

type GetCoordsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetCoordsReq) Reset() {
	*x = GetCoordsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoordsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoordsReq) ProtoMessage() {}

func (x *GetCoordsReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoordsReq.ProtoReflect.Descriptor instead.
func (*GetCoordsReq) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{3}
}

func (x *GetCoordsReq) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetCoordsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Lat  float64 `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Long float64 `protobuf:"fixed64,3,opt,name=long,proto3" json:"long,omitempty"`
}

func (x *GetCoordsResp) Reset() {
	*x = GetCoordsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoordsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoordsResp) ProtoMessage() {}

func (x *GetCoordsResp) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoordsResp.ProtoReflect.Descriptor instead.
func (*GetCoordsResp) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{4}
}

func (x *GetCoordsResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetCoordsResp) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *GetCoordsResp) GetLong() float64 {
	if x != nil {
		return x.Long
	}
	return 0
}

type UpdateUnitsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUnitsReq) Reset() {
	*x = UpdateUnitsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUnitsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUnitsReq) ProtoMessage() {}

func (x *UpdateUnitsReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUnitsReq.ProtoReflect.Descriptor instead.
func (*UpdateUnitsReq) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{5}
}

type UpdateUnitsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coords []*Coordinates `protobuf:"bytes,1,rep,name=coords,proto3" json:"coords,omitempty"`
}

func (x *UpdateUnitsResp) Reset() {
	*x = UpdateUnitsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUnitsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUnitsResp) ProtoMessage() {}

func (x *UpdateUnitsResp) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUnitsResp.ProtoReflect.Descriptor instead.
func (*UpdateUnitsResp) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUnitsResp) GetCoords() []*Coordinates {
	if x != nil {
		return x.Coords
	}
	return nil
}

type Coordinates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dt   string    `protobuf:"bytes,1,opt,name=dt,proto3" json:"dt,omitempty"`
	Id   string    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type Resuource `protobuf:"varint,3,opt,name=type,proto3,enum=Resuource" json:"type,omitempty"`
	Lat  float64   `protobuf:"fixed64,4,opt,name=lat,proto3" json:"lat,omitempty"`
	Long float64   `protobuf:"fixed64,5,opt,name=long,proto3" json:"long,omitempty"`
}

func (x *Coordinates) Reset() {
	*x = Coordinates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinates) ProtoMessage() {}

func (x *Coordinates) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinates.ProtoReflect.Descriptor instead.
func (*Coordinates) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{7}
}

func (x *Coordinates) GetDt() string {
	if x != nil {
		return x.Dt
	}
	return ""
}

func (x *Coordinates) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Coordinates) GetType() Resuource {
	if x != nil {
		return x.Type
	}
	return Resuource_TRUCK
}

func (x *Coordinates) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Coordinates) GetLong() float64 {
	if x != nil {
		return x.Long
	}
	return 0
}

type Units struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *Units) Reset() {
	*x = Units{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Units) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Units) ProtoMessage() {}

func (x *Units) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Units.ProtoReflect.Descriptor instead.
func (*Units) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{8}
}

func (x *Units) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Coords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat  float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Long float64 `protobuf:"fixed64,2,opt,name=long,proto3" json:"long,omitempty"`
}

func (x *Coords) Reset() {
	*x = Coords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coords) ProtoMessage() {}

func (x *Coords) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coords.ProtoReflect.Descriptor instead.
func (*Coords) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{9}
}

func (x *Coords) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Coords) GetLong() float64 {
	if x != nil {
		return x.Long
	}
	return 0
}

type OpeataionOn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Units  *Units  `protobuf:"bytes,1,opt,name=units,proto3" json:"units,omitempty"`
	Coords *Coords `protobuf:"bytes,2,opt,name=coords,proto3" json:"coords,omitempty"`
}

func (x *OpeataionOn) Reset() {
	*x = OpeataionOn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpeataionOn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpeataionOn) ProtoMessage() {}

func (x *OpeataionOn) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpeataionOn.ProtoReflect.Descriptor instead.
func (*OpeataionOn) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{10}
}

func (x *OpeataionOn) GetUnits() *Units {
	if x != nil {
		return x.Units
	}
	return nil
}

func (x *OpeataionOn) GetCoords() *Coords {
	if x != nil {
		return x.Coords
	}
	return nil
}

type OperationFromTo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Units *Units  `protobuf:"bytes,1,opt,name=units,proto3" json:"units,omitempty"`
	From  *Coords `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To    *Coords `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *OperationFromTo) Reset() {
	*x = OperationFromTo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationFromTo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationFromTo) ProtoMessage() {}

func (x *OperationFromTo) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationFromTo.ProtoReflect.Descriptor instead.
func (*OperationFromTo) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{11}
}

func (x *OperationFromTo) GetUnits() *Units {
	if x != nil {
		return x.Units
	}
	return nil
}

func (x *OperationFromTo) GetFrom() *Coords {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *OperationFromTo) GetTo() *Coords {
	if x != nil {
		return x.To
	}
	return nil
}

type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit.ProtoReflect.Descriptor instead.
func (*Unit) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{12}
}

func (x *Unit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Operaions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From          *Coords `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            *Coords `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	OperationName string  `protobuf:"bytes,3,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
}

func (x *Operaions) Reset() {
	*x = Operaions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operaions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operaions) ProtoMessage() {}

func (x *Operaions) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_coordservice_coordservice_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operaions.ProtoReflect.Descriptor instead.
func (*Operaions) Descriptor() ([]byte, []int) {
	return file_src_protos_coordservice_coordservice_proto_rawDescGZIP(), []int{13}
}

func (x *Operaions) GetFrom() *Coords {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Operaions) GetTo() *Coords {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Operaions) GetOperationName() string {
	if x != nil {
		return x.OperationName
	}
	return ""
}

var File_src_protos_coordservice_coordservice_proto protoreflect.FileDescriptor

var file_src_protos_coordservice_coordservice_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x11, 0x0a,
	0x0f, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x20, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x22, 0x45, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x22, 0x37, 0x0a, 0x0f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24,
	0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x06, 0x63, 0x6f,
	0x6f, 0x72, 0x64, 0x73, 0x22, 0x73, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x64, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0a, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x19, 0x0a, 0x05, 0x55, 0x6e, 0x69,
	0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x2e, 0x0a, 0x06, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x4c, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x61, 0x74, 0x61, 0x69, 0x6f,
	0x6e, 0x4f, 0x6e, 0x12, 0x1c, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74,
	0x73, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x73, 0x22, 0x65, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x12, 0x1c, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x05, 0x75, 0x6e,
	0x69, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x17, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x04, 0x55, 0x6e, 0x69,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x68, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x7e, 0x0a, 0x09, 0x52,
	0x65, 0x73, 0x75, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x55, 0x43,
	0x4b, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x43, 0x41, 0x56, 0x41, 0x54, 0x4f, 0x52, 0x10,
	0x01, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x4f, 0x54, 0x41, 0x52, 0x59, 0x5f, 0x4c, 0x4f, 0x41, 0x44,
	0x45, 0x52, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4c, 0x4f, 0x57, 0x5f, 0x41, 0x4e, 0x44,
	0x5f, 0x42, 0x52, 0x55, 0x53, 0x48, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x10, 0x03,
	0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x4c, 0x4f, 0x57, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x42, 0x52, 0x55,
	0x53, 0x48, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x10, 0x04,
	0x12, 0x09, 0x0a, 0x05, 0x48, 0x55, 0x4d, 0x41, 0x4e, 0x10, 0x05, 0x32, 0xe3, 0x03, 0x0a, 0x0d,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a,
	0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0f, 0x2e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x25, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x06, 0x2e, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x1a, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x30, 0x01, 0x12, 0x30, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x6e,
	0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x0a, 0x53, 0x77, 0x65, 0x65,
	0x70, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x12, 0x0c, 0x2e, 0x4f, 0x70, 0x65, 0x61, 0x74, 0x61, 0x69,
	0x6f, 0x6e, 0x4f, 0x6e, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x10,
	0x52, 0x65, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0c, 0x2e, 0x4f, 0x70, 0x65, 0x61, 0x74, 0x61, 0x69, 0x6f, 0x6e, 0x4f, 0x6e, 0x1a, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x68, 0x61, 0x66, 0x74, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x4f, 0x70, 0x65, 0x61, 0x74,
	0x61, 0x69, 0x6f, 0x6e, 0x4f, 0x6e, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2a,
	0x0a, 0x0e, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x6e, 0x6f, 0x77, 0x54, 0x6f, 0x54, 0x65, 0x6d, 0x70,
	0x12, 0x10, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x6f, 0x6d,
	0x54, 0x6f, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x13, 0x4c, 0x6f,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x6e, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d,
	0x70, 0x12, 0x0c, 0x2e, 0x4f, 0x70, 0x65, 0x61, 0x74, 0x61, 0x69, 0x6f, 0x6e, 0x4f, 0x6e, 0x1a,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x53, 0x6e, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x10, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x1a,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x09, 0x43, 0x6c, 0x65, 0x61, 0x72,
	0x54, 0x65, 0x6d, 0x70, 0x12, 0x0c, 0x2e, 0x4f, 0x70, 0x65, 0x61, 0x74, 0x61, 0x69, 0x6f, 0x6e,
	0x4f, 0x6e, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x05, 0x2e, 0x55,
	0x6e, 0x69, 0x74, 0x1a, 0x0a, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x69, 0x6f, 0x6e, 0x73, 0x30,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_protos_coordservice_coordservice_proto_rawDescOnce sync.Once
	file_src_protos_coordservice_coordservice_proto_rawDescData = file_src_protos_coordservice_coordservice_proto_rawDesc
)

func file_src_protos_coordservice_coordservice_proto_rawDescGZIP() []byte {
	file_src_protos_coordservice_coordservice_proto_rawDescOnce.Do(func() {
		file_src_protos_coordservice_coordservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_protos_coordservice_coordservice_proto_rawDescData)
	})
	return file_src_protos_coordservice_coordservice_proto_rawDescData
}

var file_src_protos_coordservice_coordservice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_protos_coordservice_coordservice_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_src_protos_coordservice_coordservice_proto_goTypes = []interface{}{
	(Resuource)(0),          // 0: Resuource
	(*WriteCoordsReq)(nil),  // 1: WriteCoordsReq
	(*Empty)(nil),           // 2: Empty
	(*WriteCoordsResp)(nil), // 3: WriteCoordsResp
	(*GetCoordsReq)(nil),    // 4: GetCoordsReq
	(*GetCoordsResp)(nil),   // 5: GetCoordsResp
	(*UpdateUnitsReq)(nil),  // 6: UpdateUnitsReq
	(*UpdateUnitsResp)(nil), // 7: UpdateUnitsResp
	(*Coordinates)(nil),     // 8: Coordinates
	(*Units)(nil),           // 9: Units
	(*Coords)(nil),          // 10: Coords
	(*OpeataionOn)(nil),     // 11: OpeataionOn
	(*OperationFromTo)(nil), // 12: OperationFromTo
	(*Unit)(nil),            // 13: Unit
	(*Operaions)(nil),       // 14: Operaions
}
var file_src_protos_coordservice_coordservice_proto_depIdxs = []int32{
	8,  // 0: UpdateUnitsResp.coords:type_name -> Coordinates
	0,  // 1: Coordinates.type:type_name -> Resuource
	9,  // 2: OpeataionOn.units:type_name -> Units
	10, // 3: OpeataionOn.coords:type_name -> Coords
	9,  // 4: OperationFromTo.units:type_name -> Units
	10, // 5: OperationFromTo.from:type_name -> Coords
	10, // 6: OperationFromTo.to:type_name -> Coords
	10, // 7: Operaions.from:type_name -> Coords
	10, // 8: Operaions.to:type_name -> Coords
	1,  // 9: CoordsService.WriteCoords:input_type -> WriteCoordsReq
	9,  // 10: CoordsService.GetCoords:input_type -> Units
	6,  // 11: CoordsService.UpdateUnits:input_type -> UpdateUnitsReq
	11, // 12: CoordsService.SweepingUp:input_type -> OpeataionOn
	11, // 13: CoordsService.ReagentTreatment:input_type -> OpeataionOn
	11, // 14: CoordsService.ShaftFormation:input_type -> OpeataionOn
	12, // 15: CoordsService.MoveSnowToTemp:input_type -> OperationFromTo
	11, // 16: CoordsService.LoadingSnowFromTemp:input_type -> OpeataionOn
	12, // 17: CoordsService.ExportSnowFromTemp:input_type -> OperationFromTo
	11, // 18: CoordsService.ClearTemp:input_type -> OpeataionOn
	13, // 19: CoordsService.ListenCommands:input_type -> Unit
	3,  // 20: CoordsService.WriteCoords:output_type -> WriteCoordsResp
	5,  // 21: CoordsService.GetCoords:output_type -> GetCoordsResp
	7,  // 22: CoordsService.UpdateUnits:output_type -> UpdateUnitsResp
	2,  // 23: CoordsService.SweepingUp:output_type -> Empty
	2,  // 24: CoordsService.ReagentTreatment:output_type -> Empty
	2,  // 25: CoordsService.ShaftFormation:output_type -> Empty
	2,  // 26: CoordsService.MoveSnowToTemp:output_type -> Empty
	2,  // 27: CoordsService.LoadingSnowFromTemp:output_type -> Empty
	2,  // 28: CoordsService.ExportSnowFromTemp:output_type -> Empty
	2,  // 29: CoordsService.ClearTemp:output_type -> Empty
	14, // 30: CoordsService.ListenCommands:output_type -> Operaions
	20, // [20:31] is the sub-list for method output_type
	9,  // [9:20] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_src_protos_coordservice_coordservice_proto_init() }
func file_src_protos_coordservice_coordservice_proto_init() {
	if File_src_protos_coordservice_coordservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_protos_coordservice_coordservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteCoordsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteCoordsResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoordsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoordsResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUnitsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUnitsResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinates); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Units); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coords); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpeataionOn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationFromTo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protos_coordservice_coordservice_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operaions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_protos_coordservice_coordservice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_protos_coordservice_coordservice_proto_goTypes,
		DependencyIndexes: file_src_protos_coordservice_coordservice_proto_depIdxs,
		EnumInfos:         file_src_protos_coordservice_coordservice_proto_enumTypes,
		MessageInfos:      file_src_protos_coordservice_coordservice_proto_msgTypes,
	}.Build()
	File_src_protos_coordservice_coordservice_proto = out.File
	file_src_protos_coordservice_coordservice_proto_rawDesc = nil
	file_src_protos_coordservice_coordservice_proto_goTypes = nil
	file_src_protos_coordservice_coordservice_proto_depIdxs = nil
}
