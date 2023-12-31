// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: clients.proto

package user_service

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

type CreateClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname        string  `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname         string  `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Phone            string  `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Photo            string  `protobuf:"bytes,4,opt,name=photo,proto3" json:"photo,omitempty"`
	BirthDate        string  `protobuf:"bytes,5,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	DiscountType     string  `protobuf:"bytes,6,opt,name=discount_type,json=discountType,proto3" json:"discount_type,omitempty"`
	DiscountAmount   float64 `protobuf:"fixed64,7,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount,omitempty"`
	TotalOrdersCount int64   `protobuf:"varint,8,opt,name=total_orders_count,json=totalOrdersCount,proto3" json:"total_orders_count,omitempty"`
	TotalOrdersSum   float64 `protobuf:"fixed64,9,opt,name=total_orders_sum,json=totalOrdersSum,proto3" json:"total_orders_sum,omitempty"`
}

func (x *CreateClientsRequest) Reset() {
	*x = CreateClientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clients_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientsRequest) ProtoMessage() {}

func (x *CreateClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clients_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientsRequest.ProtoReflect.Descriptor instead.
func (*CreateClientsRequest) Descriptor() ([]byte, []int) {
	return file_clients_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClientsRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *CreateClientsRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *CreateClientsRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateClientsRequest) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *CreateClientsRequest) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *CreateClientsRequest) GetDiscountType() string {
	if x != nil {
		return x.DiscountType
	}
	return ""
}

func (x *CreateClientsRequest) GetDiscountAmount() float64 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *CreateClientsRequest) GetTotalOrdersCount() int64 {
	if x != nil {
		return x.TotalOrdersCount
	}
	return 0
}

func (x *CreateClientsRequest) GetTotalOrdersSum() float64 {
	if x != nil {
		return x.TotalOrdersSum
	}
	return 0
}

type Clients struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname        string  `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname         string  `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Phone            string  `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Photo            string  `protobuf:"bytes,5,opt,name=photo,proto3" json:"photo,omitempty"`
	BirthDate        string  `protobuf:"bytes,6,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	LastOrderedDate  string  `protobuf:"bytes,7,opt,name=last_ordered_date,json=lastOrderedDate,proto3" json:"last_ordered_date,omitempty"`
	TotalOrdersCount int64   `protobuf:"varint,8,opt,name=total_orders_count,json=totalOrdersCount,proto3" json:"total_orders_count,omitempty"`
	TotalOrdersSum   float64 `protobuf:"fixed64,9,opt,name=total_orders_sum,json=totalOrdersSum,proto3" json:"total_orders_sum,omitempty"`
	DiscountType     string  `protobuf:"bytes,10,opt,name=discount_type,json=discountType,proto3" json:"discount_type,omitempty"`
	DiscountAmount   float64 `protobuf:"fixed64,11,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount,omitempty"`
	CreatedAt        string  `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        string  `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Clients) Reset() {
	*x = Clients{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clients_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clients) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clients) ProtoMessage() {}

func (x *Clients) ProtoReflect() protoreflect.Message {
	mi := &file_clients_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clients.ProtoReflect.Descriptor instead.
func (*Clients) Descriptor() ([]byte, []int) {
	return file_clients_proto_rawDescGZIP(), []int{1}
}

func (x *Clients) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Clients) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Clients) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Clients) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Clients) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *Clients) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *Clients) GetLastOrderedDate() string {
	if x != nil {
		return x.LastOrderedDate
	}
	return ""
}

func (x *Clients) GetTotalOrdersCount() int64 {
	if x != nil {
		return x.TotalOrdersCount
	}
	return 0
}

func (x *Clients) GetTotalOrdersSum() float64 {
	if x != nil {
		return x.TotalOrdersSum
	}
	return 0
}

func (x *Clients) GetDiscountType() string {
	if x != nil {
		return x.DiscountType
	}
	return ""
}

func (x *Clients) GetDiscountAmount() float64 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *Clients) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Clients) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname      string  `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname       string  `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Phone          string  `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Photo          string  `protobuf:"bytes,5,opt,name=photo,proto3" json:"photo,omitempty"`
	BirthDate      string  `protobuf:"bytes,6,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	DiscountType   string  `protobuf:"bytes,7,opt,name=discount_type,json=discountType,proto3" json:"discount_type,omitempty"`
	DiscountAmount float64 `protobuf:"fixed64,8,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount,omitempty"`
}

func (x *UpdateClientsRequest) Reset() {
	*x = UpdateClientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clients_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClientsRequest) ProtoMessage() {}

func (x *UpdateClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clients_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClientsRequest.ProtoReflect.Descriptor instead.
func (*UpdateClientsRequest) Descriptor() ([]byte, []int) {
	return file_clients_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateClientsRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateClientsRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UpdateClientsRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UpdateClientsRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateClientsRequest) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *UpdateClientsRequest) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *UpdateClientsRequest) GetDiscountType() string {
	if x != nil {
		return x.DiscountType
	}
	return ""
}

func (x *UpdateClientsRequest) GetDiscountAmount() float64 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

type UpdateClientsOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TotalOrdersCount int64   `protobuf:"varint,2,opt,name=total_orders_count,json=totalOrdersCount,proto3" json:"total_orders_count,omitempty"`
	TotalOrdersSum   float64 `protobuf:"fixed64,3,opt,name=total_orders_sum,json=totalOrdersSum,proto3" json:"total_orders_sum,omitempty"`
}

func (x *UpdateClientsOrderRequest) Reset() {
	*x = UpdateClientsOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clients_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClientsOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClientsOrderRequest) ProtoMessage() {}

func (x *UpdateClientsOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clients_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClientsOrderRequest.ProtoReflect.Descriptor instead.
func (*UpdateClientsOrderRequest) Descriptor() ([]byte, []int) {
	return file_clients_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateClientsOrderRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateClientsOrderRequest) GetTotalOrdersCount() int64 {
	if x != nil {
		return x.TotalOrdersCount
	}
	return 0
}

func (x *UpdateClientsOrderRequest) GetTotalOrdersSum() float64 {
	if x != nil {
		return x.TotalOrdersSum
	}
	return 0
}

type ListClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit                int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Firstname            string `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Phone                string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	CreatedAtFrom        string `protobuf:"bytes,6,opt,name=created_at_from,json=createdAtFrom,proto3" json:"created_at_from,omitempty"`
	CreatedAtTo          string `protobuf:"bytes,7,opt,name=created_at_to,json=createdAtTo,proto3" json:"created_at_to,omitempty"`
	LastOrderedDateFrom  string `protobuf:"bytes,8,opt,name=last_ordered_date_from,json=lastOrderedDateFrom,proto3" json:"last_ordered_date_from,omitempty"`
	LastOrderedDateTo    string `protobuf:"bytes,9,opt,name=last_ordered_date_to,json=lastOrderedDateTo,proto3" json:"last_ordered_date_to,omitempty"`
	TotalOrdersSumFrom   int64  `protobuf:"varint,10,opt,name=total_orders_sum_from,json=totalOrdersSumFrom,proto3" json:"total_orders_sum_from,omitempty"`
	TotalOrdersSumTo     int64  `protobuf:"varint,11,opt,name=total_orders_sum_to,json=totalOrdersSumTo,proto3" json:"total_orders_sum_to,omitempty"`
	TotalOrdersCountFrom int64  `protobuf:"varint,12,opt,name=total_orders_count_from,json=totalOrdersCountFrom,proto3" json:"total_orders_count_from,omitempty"`
	TotalOrdersCountTo   int64  `protobuf:"varint,13,opt,name=total_orders_count_to,json=totalOrdersCountTo,proto3" json:"total_orders_count_to,omitempty"`
	DiscountType         string `protobuf:"bytes,14,opt,name=discount_type,json=discountType,proto3" json:"discount_type,omitempty"`
	DiscountAmountFrom   string `protobuf:"bytes,15,opt,name=discount_amount_from,json=discountAmountFrom,proto3" json:"discount_amount_from,omitempty"`
	DiscountAmountTo     string `protobuf:"bytes,16,opt,name=discount_amount_to,json=discountAmountTo,proto3" json:"discount_amount_to,omitempty"`
}

func (x *ListClientsRequest) Reset() {
	*x = ListClientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clients_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClientsRequest) ProtoMessage() {}

func (x *ListClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clients_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClientsRequest.ProtoReflect.Descriptor instead.
func (*ListClientsRequest) Descriptor() ([]byte, []int) {
	return file_clients_proto_rawDescGZIP(), []int{4}
}

func (x *ListClientsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListClientsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListClientsRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *ListClientsRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *ListClientsRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ListClientsRequest) GetCreatedAtFrom() string {
	if x != nil {
		return x.CreatedAtFrom
	}
	return ""
}

func (x *ListClientsRequest) GetCreatedAtTo() string {
	if x != nil {
		return x.CreatedAtTo
	}
	return ""
}

func (x *ListClientsRequest) GetLastOrderedDateFrom() string {
	if x != nil {
		return x.LastOrderedDateFrom
	}
	return ""
}

func (x *ListClientsRequest) GetLastOrderedDateTo() string {
	if x != nil {
		return x.LastOrderedDateTo
	}
	return ""
}

func (x *ListClientsRequest) GetTotalOrdersSumFrom() int64 {
	if x != nil {
		return x.TotalOrdersSumFrom
	}
	return 0
}

func (x *ListClientsRequest) GetTotalOrdersSumTo() int64 {
	if x != nil {
		return x.TotalOrdersSumTo
	}
	return 0
}

func (x *ListClientsRequest) GetTotalOrdersCountFrom() int64 {
	if x != nil {
		return x.TotalOrdersCountFrom
	}
	return 0
}

func (x *ListClientsRequest) GetTotalOrdersCountTo() int64 {
	if x != nil {
		return x.TotalOrdersCountTo
	}
	return 0
}

func (x *ListClientsRequest) GetDiscountType() string {
	if x != nil {
		return x.DiscountType
	}
	return ""
}

func (x *ListClientsRequest) GetDiscountAmountFrom() string {
	if x != nil {
		return x.DiscountAmountFrom
	}
	return ""
}

func (x *ListClientsRequest) GetDiscountAmountTo() string {
	if x != nil {
		return x.DiscountAmountTo
	}
	return ""
}

type ListClientsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*Clients `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
	Count   int32      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListClientsResponse) Reset() {
	*x = ListClientsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clients_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClientsResponse) ProtoMessage() {}

func (x *ListClientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clients_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClientsResponse.ProtoReflect.Descriptor instead.
func (*ListClientsResponse) Descriptor() ([]byte, []int) {
	return file_clients_proto_rawDescGZIP(), []int{5}
}

func (x *ListClientsResponse) GetClients() []*Clients {
	if x != nil {
		return x.Clients
	}
	return nil
}

func (x *ListClientsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_clients_proto protoreflect.FileDescriptor

var file_clients_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x53, 0x75, 0x6d, 0x22,
	0xae, 0x03, 0x0a, 0x07, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x61, 0x73,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x53, 0x75, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xf9, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x83, 0x01, 0x0a,
	0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x53,
	0x75, 0x6d, 0x22, 0x91, 0x05, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x54, 0x6f, 0x12,
	0x33, 0x0a, 0x16, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x2f, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x31, 0x0a, 0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x75, 0x6d, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x53, 0x75, 0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x2d, 0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x53, 0x75, 0x6d, 0x54, 0x6f, 0x12, 0x35, 0x0a, 0x17, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x31,
	0x0a, 0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x22, 0x5c, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0xb6, 0x03, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a,
	0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clients_proto_rawDescOnce sync.Once
	file_clients_proto_rawDescData = file_clients_proto_rawDesc
)

func file_clients_proto_rawDescGZIP() []byte {
	file_clients_proto_rawDescOnce.Do(func() {
		file_clients_proto_rawDescData = protoimpl.X.CompressGZIP(file_clients_proto_rawDescData)
	})
	return file_clients_proto_rawDescData
}

var file_clients_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_clients_proto_goTypes = []interface{}{
	(*CreateClientsRequest)(nil),      // 0: user_service.CreateClientsRequest
	(*Clients)(nil),                   // 1: user_service.Clients
	(*UpdateClientsRequest)(nil),      // 2: user_service.UpdateClientsRequest
	(*UpdateClientsOrderRequest)(nil), // 3: user_service.UpdateClientsOrderRequest
	(*ListClientsRequest)(nil),        // 4: user_service.ListClientsRequest
	(*ListClientsResponse)(nil),       // 5: user_service.ListClientsResponse
	(*IdRequest)(nil),                 // 6: user_service.IdRequest
	(*Response)(nil),                  // 7: user_service.Response
}
var file_clients_proto_depIdxs = []int32{
	1, // 0: user_service.ListClientsResponse.clients:type_name -> user_service.Clients
	0, // 1: user_service.ClientService.Create:input_type -> user_service.CreateClientsRequest
	6, // 2: user_service.ClientService.Get:input_type -> user_service.IdRequest
	4, // 3: user_service.ClientService.List:input_type -> user_service.ListClientsRequest
	2, // 4: user_service.ClientService.Update:input_type -> user_service.UpdateClientsRequest
	6, // 5: user_service.ClientService.Delete:input_type -> user_service.IdRequest
	3, // 6: user_service.ClientService.UpdateOrder:input_type -> user_service.UpdateClientsOrderRequest
	7, // 7: user_service.ClientService.Create:output_type -> user_service.Response
	1, // 8: user_service.ClientService.Get:output_type -> user_service.Clients
	5, // 9: user_service.ClientService.List:output_type -> user_service.ListClientsResponse
	7, // 10: user_service.ClientService.Update:output_type -> user_service.Response
	7, // 11: user_service.ClientService.Delete:output_type -> user_service.Response
	7, // 12: user_service.ClientService.UpdateOrder:output_type -> user_service.Response
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_clients_proto_init() }
func file_clients_proto_init() {
	if File_clients_proto != nil {
		return
	}
	file_branch_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_clients_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientsRequest); i {
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
		file_clients_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clients); i {
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
		file_clients_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClientsRequest); i {
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
		file_clients_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClientsOrderRequest); i {
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
		file_clients_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClientsRequest); i {
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
		file_clients_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClientsResponse); i {
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
			RawDescriptor: file_clients_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clients_proto_goTypes,
		DependencyIndexes: file_clients_proto_depIdxs,
		MessageInfos:      file_clients_proto_msgTypes,
	}.Build()
	File_clients_proto = out.File
	file_clients_proto_rawDesc = nil
	file_clients_proto_goTypes = nil
	file_clients_proto_depIdxs = nil
}
