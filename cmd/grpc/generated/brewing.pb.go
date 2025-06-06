// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: api/proto/brewing.proto

package generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetCookingStatusResponse_Status int32

const (
	GetCookingStatusResponse_UNKNOWN     GetCookingStatusResponse_Status = 0
	GetCookingStatusResponse_IN_PROGRESS GetCookingStatusResponse_Status = 1
	GetCookingStatusResponse_COMPLETED   GetCookingStatusResponse_Status = 2
)

// Enum value maps for GetCookingStatusResponse_Status.
var (
	GetCookingStatusResponse_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "IN_PROGRESS",
		2: "COMPLETED",
	}
	GetCookingStatusResponse_Status_value = map[string]int32{
		"UNKNOWN":     0,
		"IN_PROGRESS": 1,
		"COMPLETED":   2,
	}
)

func (x GetCookingStatusResponse_Status) Enum() *GetCookingStatusResponse_Status {
	p := new(GetCookingStatusResponse_Status)
	*p = x
	return p
}

func (x GetCookingStatusResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetCookingStatusResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_brewing_proto_enumTypes[0].Descriptor()
}

func (GetCookingStatusResponse_Status) Type() protoreflect.EnumType {
	return &file_api_proto_brewing_proto_enumTypes[0]
}

func (x GetCookingStatusResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetCookingStatusResponse_Status.Descriptor instead.
func (GetCookingStatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_brewing_proto_rawDescGZIP(), []int{5, 0}
}

type Recipe struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Uuid            string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BrewTimeSeconds int32                  `protobuf:"varint,3,opt,name=brew_time_seconds,json=brewTimeSeconds,proto3" json:"brew_time_seconds,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Recipe) Reset() {
	*x = Recipe{}
	mi := &file_api_proto_brewing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Recipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipe) ProtoMessage() {}

func (x *Recipe) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_brewing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipe.ProtoReflect.Descriptor instead.
func (*Recipe) Descriptor() ([]byte, []int) {
	return file_api_proto_brewing_proto_rawDescGZIP(), []int{0}
}

func (x *Recipe) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Recipe) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Recipe) GetBrewTimeSeconds() int32 {
	if x != nil {
		return x.BrewTimeSeconds
	}
	return 0
}

type Witch struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Witch) Reset() {
	*x = Witch{}
	mi := &file_api_proto_brewing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Witch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Witch) ProtoMessage() {}

func (x *Witch) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_brewing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Witch.ProtoReflect.Descriptor instead.
func (*Witch) Descriptor() ([]byte, []int) {
	return file_api_proto_brewing_proto_rawDescGZIP(), []int{1}
}

func (x *Witch) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Witch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StartCookingRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	RecipeUuid      string                 `protobuf:"bytes,1,opt,name=recipe_uuid,json=recipeUuid,proto3" json:"recipe_uuid,omitempty"`
	WitchUuid       string                 `protobuf:"bytes,2,opt,name=witch_uuid,json=witchUuid,proto3" json:"witch_uuid,omitempty"`
	BrewTimeSeconds int32                  `protobuf:"varint,3,opt,name=brew_time_seconds,json=brewTimeSeconds,proto3" json:"brew_time_seconds,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *StartCookingRequest) Reset() {
	*x = StartCookingRequest{}
	mi := &file_api_proto_brewing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartCookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartCookingRequest) ProtoMessage() {}

func (x *StartCookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_brewing_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartCookingRequest.ProtoReflect.Descriptor instead.
func (*StartCookingRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_brewing_proto_rawDescGZIP(), []int{2}
}

func (x *StartCookingRequest) GetRecipeUuid() string {
	if x != nil {
		return x.RecipeUuid
	}
	return ""
}

func (x *StartCookingRequest) GetWitchUuid() string {
	if x != nil {
		return x.WitchUuid
	}
	return ""
}

func (x *StartCookingRequest) GetBrewTimeSeconds() int32 {
	if x != nil {
		return x.BrewTimeSeconds
	}
	return 0
}

type StartCookingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecipeUuid    string                 `protobuf:"bytes,1,opt,name=recipe_uuid,json=recipeUuid,proto3" json:"recipe_uuid,omitempty"`
	RecipeName    string                 `protobuf:"bytes,2,opt,name=recipe_name,json=recipeName,proto3" json:"recipe_name,omitempty"`
	WitchUuid     string                 `protobuf:"bytes,4,opt,name=witch_uuid,json=witchUuid,proto3" json:"witch_uuid,omitempty"`
	WitchName     string                 `protobuf:"bytes,5,opt,name=witch_name,json=witchName,proto3" json:"witch_name,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartCookingResponse) Reset() {
	*x = StartCookingResponse{}
	mi := &file_api_proto_brewing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartCookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartCookingResponse) ProtoMessage() {}

func (x *StartCookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_brewing_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartCookingResponse.ProtoReflect.Descriptor instead.
func (*StartCookingResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_brewing_proto_rawDescGZIP(), []int{3}
}

func (x *StartCookingResponse) GetRecipeUuid() string {
	if x != nil {
		return x.RecipeUuid
	}
	return ""
}

func (x *StartCookingResponse) GetRecipeName() string {
	if x != nil {
		return x.RecipeName
	}
	return ""
}

func (x *StartCookingResponse) GetWitchUuid() string {
	if x != nil {
		return x.WitchUuid
	}
	return ""
}

func (x *StartCookingResponse) GetWitchName() string {
	if x != nil {
		return x.WitchName
	}
	return ""
}

func (x *StartCookingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetCookingStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecipeUuid    string                 `protobuf:"bytes,1,opt,name=recipe_uuid,json=recipeUuid,proto3" json:"recipe_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCookingStatusRequest) Reset() {
	*x = GetCookingStatusRequest{}
	mi := &file_api_proto_brewing_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCookingStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCookingStatusRequest) ProtoMessage() {}

func (x *GetCookingStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_brewing_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCookingStatusRequest.ProtoReflect.Descriptor instead.
func (*GetCookingStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_brewing_proto_rawDescGZIP(), []int{4}
}

func (x *GetCookingStatusRequest) GetRecipeUuid() string {
	if x != nil {
		return x.RecipeUuid
	}
	return ""
}

type GetCookingStatusResponse struct {
	state         protoimpl.MessageState          `protogen:"open.v1"`
	Status        GetCookingStatusResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=brewing.GetCookingStatusResponse_Status" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCookingStatusResponse) Reset() {
	*x = GetCookingStatusResponse{}
	mi := &file_api_proto_brewing_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCookingStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCookingStatusResponse) ProtoMessage() {}

func (x *GetCookingStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_brewing_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCookingStatusResponse.ProtoReflect.Descriptor instead.
func (*GetCookingStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_brewing_proto_rawDescGZIP(), []int{5}
}

func (x *GetCookingStatusResponse) GetStatus() GetCookingStatusResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetCookingStatusResponse_UNKNOWN
}

var File_api_proto_brewing_proto protoreflect.FileDescriptor

var file_api_proto_brewing_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x72, 0x65, 0x77,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x72, 0x65, 0x77, 0x69,
	0x6e, 0x67, 0x22, 0x5c, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x72, 0x65, 0x77, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x62, 0x72, 0x65, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x22, 0x2f, 0x0a, 0x05, 0x57, 0x69, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x69,
	0x74, 0x63, 0x68, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x55, 0x75, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x72, 0x65,
	0x77, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x62, 0x72, 0x65, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x69, 0x74, 0x63, 0x68, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x55, 0x75,
	0x69, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x28, 0x2e, 0x62, 0x72, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x35, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0xb6, 0x01, 0x0a, 0x0e, 0x42, 0x72, 0x65,
	0x77, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x2e, 0x62, 0x72,
	0x65, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x72, 0x65, 0x77,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x2e, 0x62,
	0x72, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x62, 0x72, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_proto_brewing_proto_rawDescOnce sync.Once
	file_api_proto_brewing_proto_rawDescData []byte
)

func file_api_proto_brewing_proto_rawDescGZIP() []byte {
	file_api_proto_brewing_proto_rawDescOnce.Do(func() {
		file_api_proto_brewing_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_proto_brewing_proto_rawDesc), len(file_api_proto_brewing_proto_rawDesc)))
	})
	return file_api_proto_brewing_proto_rawDescData
}

var file_api_proto_brewing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_brewing_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_brewing_proto_goTypes = []any{
	(GetCookingStatusResponse_Status)(0), // 0: brewing.GetCookingStatusResponse.Status
	(*Recipe)(nil),                       // 1: brewing.Recipe
	(*Witch)(nil),                        // 2: brewing.Witch
	(*StartCookingRequest)(nil),          // 3: brewing.StartCookingRequest
	(*StartCookingResponse)(nil),         // 4: brewing.StartCookingResponse
	(*GetCookingStatusRequest)(nil),      // 5: brewing.GetCookingStatusRequest
	(*GetCookingStatusResponse)(nil),     // 6: brewing.GetCookingStatusResponse
}
var file_api_proto_brewing_proto_depIdxs = []int32{
	0, // 0: brewing.GetCookingStatusResponse.status:type_name -> brewing.GetCookingStatusResponse.Status
	3, // 1: brewing.BrewingService.StartCooking:input_type -> brewing.StartCookingRequest
	5, // 2: brewing.BrewingService.GetCookingStatus:input_type -> brewing.GetCookingStatusRequest
	4, // 3: brewing.BrewingService.StartCooking:output_type -> brewing.StartCookingResponse
	6, // 4: brewing.BrewingService.GetCookingStatus:output_type -> brewing.GetCookingStatusResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_brewing_proto_init() }
func file_api_proto_brewing_proto_init() {
	if File_api_proto_brewing_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_proto_brewing_proto_rawDesc), len(file_api_proto_brewing_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_brewing_proto_goTypes,
		DependencyIndexes: file_api_proto_brewing_proto_depIdxs,
		EnumInfos:         file_api_proto_brewing_proto_enumTypes,
		MessageInfos:      file_api_proto_brewing_proto_msgTypes,
	}.Build()
	File_api_proto_brewing_proto = out.File
	file_api_proto_brewing_proto_goTypes = nil
	file_api_proto_brewing_proto_depIdxs = nil
}
