// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: send.proto

package pb

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

//in case that is needed to specify the date
type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//examples : Siguente lunes, Lunes 25 feb etc
	Day      string `protobuf:"bytes,1,opt,name=day,proto3" json:"day,omitempty"`
	Hour     string `protobuf:"bytes,2,opt,name=hour,proto3" json:"hour,omitempty"`
	ToHour   string `protobuf:"bytes,3,opt,name=toHour,proto3" json:"toHour,omitempty"`
	FromHour string `protobuf:"bytes,4,opt,name=fromHour,proto3" json:"fromHour,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_send_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_send_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_send_proto_rawDescGZIP(), []int{0}
}

func (x *Date) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *Date) GetHour() string {
	if x != nil {
		return x.Hour
	}
	return ""
}

func (x *Date) GetToHour() string {
	if x != nil {
		return x.ToHour
	}
	return ""
}

func (x *Date) GetFromHour() string {
	if x != nil {
		return x.FromHour
	}
	return ""
}

//The contact of whom the message is going to be sent
type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//Just the name of the contact
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//if the contact has a degree title
	DegreeTitle string `protobuf:"bytes,2,opt,name=degreeTitle,proto3" json:"degreeTitle,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_send_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_send_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_send_proto_rawDescGZIP(), []int{1}
}

func (x *Contact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Contact) GetDegreeTitle() string {
	if x != nil {
		return x.DegreeTitle
	}
	return ""
}

//The person of CEMIAC
type Ambassador struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *Ambassador) Reset() {
	*x = Ambassador{}
	if protoimpl.UnsafeEnabled {
		mi := &file_send_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ambassador) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ambassador) ProtoMessage() {}

func (x *Ambassador) ProtoReflect() protoreflect.Message {
	mi := &file_send_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ambassador.ProtoReflect.Descriptor instead.
func (*Ambassador) Descriptor() ([]byte, []int) {
	return file_send_proto_rawDescGZIP(), []int{2}
}

func (x *Ambassador) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Ambassador) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

//The data that is going to be parsed to the template
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contact    *Contact    `protobuf:"bytes,1,opt,name=contact,proto3" json:"contact,omitempty"`
	Date       *Date       `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Ambassador *Ambassador `protobuf:"bytes,3,opt,name=ambassador,proto3" json:"ambassador,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_send_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_send_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_send_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *Data) GetDate() *Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Data) GetAmbassador() *Ambassador {
	if x != nil {
		return x.Ambassador
	}
	return nil
}

type SendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mail string `protobuf:"bytes,1,opt,name=mail,proto3" json:"mail,omitempty"`
	//Could be school principal, student
	TypeOfContact string `protobuf:"bytes,2,opt,name=typeOfContact,proto3" json:"typeOfContact,omitempty"`
	//Could be first contact student, first contact school principal, informative session,  etc
	TypeOfTemplate string `protobuf:"bytes,3,opt,name=typeOfTemplate,proto3" json:"typeOfTemplate,omitempty"`
	//The data that is going to be the information parsed to the template parsed to the template
	Data *Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SendReq) Reset() {
	*x = SendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_send_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReq) ProtoMessage() {}

func (x *SendReq) ProtoReflect() protoreflect.Message {
	mi := &file_send_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReq.ProtoReflect.Descriptor instead.
func (*SendReq) Descriptor() ([]byte, []int) {
	return file_send_proto_rawDescGZIP(), []int{4}
}

func (x *SendReq) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *SendReq) GetTypeOfContact() string {
	if x != nil {
		return x.TypeOfContact
	}
	return ""
}

func (x *SendReq) GetTypeOfTemplate() string {
	if x != nil {
		return x.TypeOfTemplate
	}
	return ""
}

func (x *SendReq) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//Returns the id of the sended mail to be processed
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SendResp) Reset() {
	*x = SendResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_send_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResp) ProtoMessage() {}

func (x *SendResp) ProtoReflect() protoreflect.Message {
	mi := &file_send_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResp.ProtoReflect.Descriptor instead.
func (*SendResp) Descriptor() ([]byte, []int) {
	return file_send_proto_rawDescGZIP(), []int{5}
}

func (x *SendResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_send_proto protoreflect.FileDescriptor

var file_send_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x60, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x75, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x6f, 0x48, 0x6f, 0x75, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x6f, 0x48, 0x6f, 0x75, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x48, 0x6f,
	0x75, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x48, 0x6f,
	0x75, 0x72, 0x22, 0x3f, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x22, 0x36, 0x0a, 0x0a, 0x41, 0x6d, 0x62, 0x61, 0x73, 0x73, 0x61, 0x64, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x7b, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x61, 0x6d, 0x62, 0x61,
	0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x6d, 0x62, 0x61, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x52, 0x0a, 0x61, 0x6d,
	0x62, 0x61, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x22, 0x89, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x79, 0x70, 0x65,
	0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x74, 0x79, 0x70, 0x65, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x26,
	0x0a, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x4f, 0x66, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x4f, 0x66, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x32, 0x60, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x25, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x61, 0x75, 0x2d, 0x4d, 0x52, 0x2f, 0x63, 0x65, 0x6d, 0x69, 0x61, 0x63, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_send_proto_rawDescOnce sync.Once
	file_send_proto_rawDescData = file_send_proto_rawDesc
)

func file_send_proto_rawDescGZIP() []byte {
	file_send_proto_rawDescOnce.Do(func() {
		file_send_proto_rawDescData = protoimpl.X.CompressGZIP(file_send_proto_rawDescData)
	})
	return file_send_proto_rawDescData
}

var file_send_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_send_proto_goTypes = []interface{}{
	(*Date)(nil),       // 0: pb.Date
	(*Contact)(nil),    // 1: pb.Contact
	(*Ambassador)(nil), // 2: pb.Ambassador
	(*Data)(nil),       // 3: pb.Data
	(*SendReq)(nil),    // 4: pb.SendReq
	(*SendResp)(nil),   // 5: pb.SendResp
}
var file_send_proto_depIdxs = []int32{
	1, // 0: pb.Data.contact:type_name -> pb.Contact
	0, // 1: pb.Data.date:type_name -> pb.Date
	2, // 2: pb.Data.ambassador:type_name -> pb.Ambassador
	3, // 3: pb.SendReq.data:type_name -> pb.Data
	4, // 4: pb.SendService.SendMail:input_type -> pb.SendReq
	4, // 5: pb.SendService.SendMails:input_type -> pb.SendReq
	5, // 6: pb.SendService.SendMail:output_type -> pb.SendResp
	5, // 7: pb.SendService.SendMails:output_type -> pb.SendResp
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_send_proto_init() }
func file_send_proto_init() {
	if File_send_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_send_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Date); i {
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
		file_send_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_send_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ambassador); i {
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
		file_send_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_send_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReq); i {
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
		file_send_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResp); i {
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
			RawDescriptor: file_send_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_send_proto_goTypes,
		DependencyIndexes: file_send_proto_depIdxs,
		MessageInfos:      file_send_proto_msgTypes,
	}.Build()
	File_send_proto = out.File
	file_send_proto_rawDesc = nil
	file_send_proto_goTypes = nil
	file_send_proto_depIdxs = nil
}