// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: payment.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreditCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique id of the credit card record
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The credit card number
	CreditCardNumber string `protobuf:"bytes,2,opt,name=credit_card_number,json=creditCardNumber,proto3" json:"credit_card_number,omitempty"`
	// The expiration date
	Expiration string `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// The 3-digit security code of the credit card
	Cvv int32 `protobuf:"varint,4,opt,name=cvv,proto3" json:"cvv,omitempty"`
}

func (x *CreditCard) Reset() {
	*x = CreditCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditCard) ProtoMessage() {}

func (x *CreditCard) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditCard.ProtoReflect.Descriptor instead.
func (*CreditCard) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

func (x *CreditCard) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreditCard) GetCreditCardNumber() string {
	if x != nil {
		return x.CreditCardNumber
	}
	return ""
}

func (x *CreditCard) GetExpiration() string {
	if x != nil {
		return x.Expiration
	}
	return ""
}

func (x *CreditCard) GetCvv() int32 {
	if x != nil {
		return x.Cvv
	}
	return 0
}

type PaymentError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The error message
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PaymentError) Reset() {
	*x = PaymentError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentError) ProtoMessage() {}

func (x *PaymentError) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentError.ProtoReflect.Descriptor instead.
func (*PaymentError) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AddFundsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique id of the customer
	CustomerId int32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The unique id of the credit card record
	CreditCardId int32 `protobuf:"varint,2,opt,name=credit_card_id,json=creditCardId,proto3" json:"credit_card_id,omitempty"`
	// The amount of funds to add
	Amount int32 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AddFundsRequest) Reset() {
	*x = AddFundsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFundsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFundsRequest) ProtoMessage() {}

func (x *AddFundsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFundsRequest.ProtoReflect.Descriptor instead.
func (*AddFundsRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{2}
}

func (x *AddFundsRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *AddFundsRequest) GetCreditCardId() int32 {
	if x != nil {
		return x.CreditCardId
	}
	return 0
}

func (x *AddFundsRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AddFundsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A message after successfully publishing message to Redis for wallet event
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddFundsResponse) Reset() {
	*x = AddFundsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFundsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFundsResponse) ProtoMessage() {}

func (x *AddFundsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFundsResponse.ProtoReflect.Descriptor instead.
func (*AddFundsResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{3}
}

func (x *AddFundsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AddFundsEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique id of the customer
	CustomerId int32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The amount of funds to add
	Amount int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AddFundsEvent) Reset() {
	*x = AddFundsEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFundsEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFundsEvent) ProtoMessage() {}

func (x *AddFundsEvent) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFundsEvent.ProtoReflect.Descriptor instead.
func (*AddFundsEvent) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{4}
}

func (x *AddFundsEvent) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *AddFundsEvent) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AddCreditCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique id of the customer
	CustomerId int32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The 16-digit credit card number
	CreditCardNumber string `protobuf:"bytes,2,opt,name=credit_card_number,json=creditCardNumber,proto3" json:"credit_card_number,omitempty"`
	// The expiration date
	Expiration string `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// The 3-digit security code of the credit card
	Cvv int32 `protobuf:"varint,4,opt,name=cvv,proto3" json:"cvv,omitempty"`
}

func (x *AddCreditCardRequest) Reset() {
	*x = AddCreditCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCreditCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCreditCardRequest) ProtoMessage() {}

func (x *AddCreditCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCreditCardRequest.ProtoReflect.Descriptor instead.
func (*AddCreditCardRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{5}
}

func (x *AddCreditCardRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *AddCreditCardRequest) GetCreditCardNumber() string {
	if x != nil {
		return x.CreditCardNumber
	}
	return ""
}

func (x *AddCreditCardRequest) GetExpiration() string {
	if x != nil {
		return x.Expiration
	}
	return ""
}

func (x *AddCreditCardRequest) GetCvv() int32 {
	if x != nil {
		return x.Cvv
	}
	return 0
}

type AddCreditCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique id of the credit card record
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddCreditCardResponse) Reset() {
	*x = AddCreditCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCreditCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCreditCardResponse) ProtoMessage() {}

func (x *AddCreditCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCreditCardResponse.ProtoReflect.Descriptor instead.
func (*AddCreditCardResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{6}
}

func (x *AddCreditCardResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCreditCardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique id of the customer
	CustomerId int32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetCreditCardsRequest) Reset() {
	*x = GetCreditCardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCreditCardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCreditCardsRequest) ProtoMessage() {}

func (x *GetCreditCardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCreditCardsRequest.ProtoReflect.Descriptor instead.
func (*GetCreditCardsRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{7}
}

func (x *GetCreditCardsRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type GetCreditCardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All credit cards that belong to the customer
	CreditCards []*GetCreditCardsResponse_CreditCard `protobuf:"bytes,1,rep,name=credit_cards,json=creditCards,proto3" json:"credit_cards,omitempty"`
}

func (x *GetCreditCardsResponse) Reset() {
	*x = GetCreditCardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCreditCardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCreditCardsResponse) ProtoMessage() {}

func (x *GetCreditCardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCreditCardsResponse.ProtoReflect.Descriptor instead.
func (*GetCreditCardsResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{8}
}

func (x *GetCreditCardsResponse) GetCreditCards() []*GetCreditCardsResponse_CreditCard {
	if x != nil {
		return x.CreditCards
	}
	return nil
}

type GetCreditCardsResponse_CreditCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique id of the credit card record
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The credit card number
	CreditCardNumber string `protobuf:"bytes,2,opt,name=credit_card_number,json=creditCardNumber,proto3" json:"credit_card_number,omitempty"`
}

func (x *GetCreditCardsResponse_CreditCard) Reset() {
	*x = GetCreditCardsResponse_CreditCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCreditCardsResponse_CreditCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCreditCardsResponse_CreditCard) ProtoMessage() {}

func (x *GetCreditCardsResponse_CreditCard) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCreditCardsResponse_CreditCard.ProtoReflect.Descriptor instead.
func (*GetCreditCardsResponse_CreditCard) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{8, 0}
}

func (x *GetCreditCardsResponse_CreditCard) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCreditCardsResponse_CreditCard) GetCreditCardNumber() string {
	if x != nil {
		return x.CreditCardNumber
	}
	return ""
}

var File_payment_proto protoreflect.FileDescriptor

var file_payment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x76, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x63, 0x76, 0x76, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x70,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x2c, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x48,
	0x0a, 0x0d, 0x41, 0x64, 0x64, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x14, 0x41, 0x64, 0x64,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72,
	0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x76, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63,
	0x76, 0x76, 0x22, 0x27, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0xb3, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x1a,
	0x4a, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a,
	0x12, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0xb4, 0x03, 0x0a, 0x0e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93,
	0x01, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x1d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x43, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d, 0x22, 0x38, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x64, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2d, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x93, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3a,
	0x22, 0x35, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2d, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2d, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x08, 0x41, 0x64,
	0x64, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x75,
	0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2f, 0x22, 0x2a, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x64, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x3a,
	0x01, 0x2a, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_payment_proto_rawDescOnce sync.Once
	file_payment_proto_rawDescData = file_payment_proto_rawDesc
)

func file_payment_proto_rawDescGZIP() []byte {
	file_payment_proto_rawDescOnce.Do(func() {
		file_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_proto_rawDescData)
	})
	return file_payment_proto_rawDescData
}

var file_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_payment_proto_goTypes = []interface{}{
	(*CreditCard)(nil),                        // 0: payment.CreditCard
	(*PaymentError)(nil),                      // 1: payment.PaymentError
	(*AddFundsRequest)(nil),                   // 2: payment.AddFundsRequest
	(*AddFundsResponse)(nil),                  // 3: payment.AddFundsResponse
	(*AddFundsEvent)(nil),                     // 4: payment.AddFundsEvent
	(*AddCreditCardRequest)(nil),              // 5: payment.AddCreditCardRequest
	(*AddCreditCardResponse)(nil),             // 6: payment.AddCreditCardResponse
	(*GetCreditCardsRequest)(nil),             // 7: payment.GetCreditCardsRequest
	(*GetCreditCardsResponse)(nil),            // 8: payment.GetCreditCardsResponse
	(*GetCreditCardsResponse_CreditCard)(nil), // 9: payment.GetCreditCardsResponse.CreditCard
}
var file_payment_proto_depIdxs = []int32{
	9, // 0: payment.GetCreditCardsResponse.credit_cards:type_name -> payment.GetCreditCardsResponse.CreditCard
	5, // 1: payment.PaymentService.AddCreditCard:input_type -> payment.AddCreditCardRequest
	7, // 2: payment.PaymentService.GetCreditCards:input_type -> payment.GetCreditCardsRequest
	2, // 3: payment.PaymentService.AddFunds:input_type -> payment.AddFundsRequest
	6, // 4: payment.PaymentService.AddCreditCard:output_type -> payment.AddCreditCardResponse
	8, // 5: payment.PaymentService.GetCreditCards:output_type -> payment.GetCreditCardsResponse
	3, // 6: payment.PaymentService.AddFunds:output_type -> payment.AddFundsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payment_proto_init() }
func file_payment_proto_init() {
	if File_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditCard); i {
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
		file_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentError); i {
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
		file_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFundsRequest); i {
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
		file_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFundsResponse); i {
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
		file_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFundsEvent); i {
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
		file_payment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCreditCardRequest); i {
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
		file_payment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCreditCardResponse); i {
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
		file_payment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCreditCardsRequest); i {
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
		file_payment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCreditCardsResponse); i {
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
		file_payment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCreditCardsResponse_CreditCard); i {
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
			RawDescriptor: file_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_proto_goTypes,
		DependencyIndexes: file_payment_proto_depIdxs,
		MessageInfos:      file_payment_proto_msgTypes,
	}.Build()
	File_payment_proto = out.File
	file_payment_proto_rawDesc = nil
	file_payment_proto_goTypes = nil
	file_payment_proto_depIdxs = nil
}
