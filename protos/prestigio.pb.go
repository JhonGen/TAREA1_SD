// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: prestigio.proto

package cliente

import (
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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nombre      string `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Valor       int32  `protobuf:"varint,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda      string `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino     string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
	TipoCliente string `protobuf:"bytes,6,opt,name=tipoCliente,proto3" json:"tipoCliente,omitempty"`
	Prioritario bool   `protobuf:"varint,7,opt,name=prioritario,proto3" json:"prioritario,omitempty"`
	Apruebo     bool   `protobuf:"varint,8,opt,name=apruebo,proto3" json:"apruebo,omitempty"`
	Intentos    int32  `protobuf:"varint,9,opt,name=Intentos,proto3" json:"Intentos,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prestigio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_prestigio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_prestigio_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Order) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *Order) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *Order) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *Order) GetTipoCliente() string {
	if x != nil {
		return x.TipoCliente
	}
	return ""
}

func (x *Order) GetPrioritario() bool {
	if x != nil {
		return x.Prioritario
	}
	return false
}

func (x *Order) GetApruebo() bool {
	if x != nil {
		return x.Apruebo
	}
	return false
}

func (x *Order) GetIntentos() int32 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

type Camion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tipo         string `protobuf:"bytes,1,opt,name=Tipo,proto3" json:"Tipo,omitempty"`
	Orden1       *Order `protobuf:"bytes,2,opt,name=orden1,proto3" json:"orden1,omitempty"`
	Orden2       *Order `protobuf:"bytes,3,opt,name=orden2,proto3" json:"orden2,omitempty"`
	Estado       string `protobuf:"bytes,4,opt,name=Estado,proto3" json:"Estado,omitempty"`
	TiempoEspera int32  `protobuf:"varint,5,opt,name=TiempoEspera,proto3" json:"TiempoEspera,omitempty"`
}

func (x *Camion) Reset() {
	*x = Camion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prestigio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Camion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Camion) ProtoMessage() {}

func (x *Camion) ProtoReflect() protoreflect.Message {
	mi := &file_prestigio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Camion.ProtoReflect.Descriptor instead.
func (*Camion) Descriptor() ([]byte, []int) {
	return file_prestigio_proto_rawDescGZIP(), []int{1}
}

func (x *Camion) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *Camion) GetOrden1() *Order {
	if x != nil {
		return x.Orden1
	}
	return nil
}

func (x *Camion) GetOrden2() *Order {
	if x != nil {
		return x.Orden2
	}
	return nil
}

func (x *Camion) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *Camion) GetTiempoEspera() int32 {
	if x != nil {
		return x.TiempoEspera
	}
	return 0
}

type CodigoSeguimiento struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codigo int32 `protobuf:"varint,1,opt,name=codigo,proto3" json:"codigo,omitempty"`
}

func (x *CodigoSeguimiento) Reset() {
	*x = CodigoSeguimiento{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prestigio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodigoSeguimiento) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodigoSeguimiento) ProtoMessage() {}

func (x *CodigoSeguimiento) ProtoReflect() protoreflect.Message {
	mi := &file_prestigio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodigoSeguimiento.ProtoReflect.Descriptor instead.
func (*CodigoSeguimiento) Descriptor() ([]byte, []int) {
	return file_prestigio_proto_rawDescGZIP(), []int{2}
}

func (x *CodigoSeguimiento) GetCodigo() int32 {
	if x != nil {
		return x.Codigo
	}
	return 0
}

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sample string `protobuf:"bytes,1,opt,name=sample,proto3" json:"sample,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prestigio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_prestigio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_prestigio_proto_rawDescGZIP(), []int{3}
}

func (x *Sample) GetSample() string {
	if x != nil {
		return x.Sample
	}
	return ""
}

type Confirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfirmationMessage string `protobuf:"bytes,1,opt,name=ConfirmationMessage,proto3" json:"ConfirmationMessage,omitempty"`
}

func (x *Confirmation) Reset() {
	*x = Confirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prestigio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmation) ProtoMessage() {}

func (x *Confirmation) ProtoReflect() protoreflect.Message {
	mi := &file_prestigio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmation.ProtoReflect.Descriptor instead.
func (*Confirmation) Descriptor() ([]byte, []int) {
	return file_prestigio_proto_rawDescGZIP(), []int{4}
}

func (x *Confirmation) GetConfirmationMessage() string {
	if x != nil {
		return x.ConfirmationMessage
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prestigio_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_prestigio_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_prestigio_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

var File_prestigio_proto protoreflect.FileDescriptor

var file_prestigio_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x65, 0x73, 0x74, 0x69, 0x67, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x22, 0xf1, 0x01, 0x0a, 0x05, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x70, 0x6f, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x69, 0x70, 0x6f, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x72, 0x75,
	0x65, 0x62, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x70, 0x72, 0x75, 0x65,
	0x62, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x22, 0xa8,
	0x01, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x70,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x69, 0x70, 0x6f, 0x12, 0x26, 0x0a,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x6e, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f,
	0x72, 0x64, 0x65, 0x6e, 0x31, 0x12, 0x26, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x6e, 0x32, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x6e, 0x32, 0x12, 0x16, 0x0a,
	0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x69, 0x65, 0x6d, 0x70, 0x6f, 0x45,
	0x73, 0x70, 0x65, 0x72, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x54, 0x69, 0x65,
	0x6d, 0x70, 0x6f, 0x45, 0x73, 0x70, 0x65, 0x72, 0x61, 0x22, 0x2b, 0x0a, 0x11, 0x43, 0x6f, 0x64,
	0x69, 0x67, 0x6f, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x40, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x32, 0xd1, 0x02, 0x0a, 0x09, 0x53,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x12, 0x2e, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x77,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x4d, 0x61, 0x6b, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x53, 0x65, 0x67, 0x75,
	0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x52, 0x65,
	0x74, 0x69, 0x72, 0x61, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x12, 0x0f, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x65, 0x2e, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x0d, 0x44, 0x65, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x12,
	0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e,
	0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x43, 0x61, 0x6d, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x65, 0x67, 0x61, 0x12, 0x0e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prestigio_proto_rawDescOnce sync.Once
	file_prestigio_proto_rawDescData = file_prestigio_proto_rawDesc
)

func file_prestigio_proto_rawDescGZIP() []byte {
	file_prestigio_proto_rawDescOnce.Do(func() {
		file_prestigio_proto_rawDescData = protoimpl.X.CompressGZIP(file_prestigio_proto_rawDescData)
	})
	return file_prestigio_proto_rawDescData
}

var file_prestigio_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_prestigio_proto_goTypes = []interface{}{
	(*Order)(nil),             // 0: cliente.Order
	(*Camion)(nil),            // 1: cliente.Camion
	(*CodigoSeguimiento)(nil), // 2: cliente.CodigoSeguimiento
	(*Sample)(nil),            // 3: cliente.Sample
	(*Confirmation)(nil),      // 4: cliente.Confirmation
	(*Status)(nil),            // 5: cliente.Status
}
var file_prestigio_proto_depIdxs = []int32{
	0, // 0: cliente.Camion.orden1:type_name -> cliente.Order
	0, // 1: cliente.Camion.orden2:type_name -> cliente.Order
	0, // 2: cliente.Solicitud.ShowOrder:input_type -> cliente.Order
	0, // 3: cliente.Solicitud.MakeOrder:input_type -> cliente.Order
	2, // 4: cliente.Solicitud.GetStatus:input_type -> cliente.CodigoSeguimiento
	1, // 5: cliente.Solicitud.RetirarOrden:input_type -> cliente.Camion
	1, // 6: cliente.Solicitud.DevolverOrden:input_type -> cliente.Camion
	0, // 7: cliente.Solicitud.ReporteEntrega:input_type -> cliente.Order
	3, // 8: cliente.Solicitud.ShowOrder:output_type -> cliente.Sample
	4, // 9: cliente.Solicitud.MakeOrder:output_type -> cliente.Confirmation
	5, // 10: cliente.Solicitud.GetStatus:output_type -> cliente.Status
	1, // 11: cliente.Solicitud.RetirarOrden:output_type -> cliente.Camion
	1, // 12: cliente.Solicitud.DevolverOrden:output_type -> cliente.Camion
	4, // 13: cliente.Solicitud.ReporteEntrega:output_type -> cliente.Confirmation
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_prestigio_proto_init() }
func file_prestigio_proto_init() {
	if File_prestigio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prestigio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_prestigio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Camion); i {
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
		file_prestigio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodigoSeguimiento); i {
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
		file_prestigio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
		file_prestigio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirmation); i {
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
		file_prestigio_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_prestigio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prestigio_proto_goTypes,
		DependencyIndexes: file_prestigio_proto_depIdxs,
		MessageInfos:      file_prestigio_proto_msgTypes,
	}.Build()
	File_prestigio_proto = out.File
	file_prestigio_proto_rawDesc = nil
	file_prestigio_proto_goTypes = nil
	file_prestigio_proto_depIdxs = nil
}
