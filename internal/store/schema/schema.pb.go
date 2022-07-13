// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: schema.proto

package schema

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

type NewsArticlePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Category    string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Source      string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Image       string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Title       string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Date        string `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	Link        string `protobuf:"bytes,7,opt,name=link,proto3" json:"link,omitempty"`
	Logo        string `protobuf:"bytes,8,opt,name=logo,proto3" json:"logo,omitempty"`
}

func (x *NewsArticlePost) Reset() {
	*x = NewsArticlePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsArticlePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsArticlePost) ProtoMessage() {}

func (x *NewsArticlePost) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsArticlePost.ProtoReflect.Descriptor instead.
func (*NewsArticlePost) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{0}
}

func (x *NewsArticlePost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewsArticlePost) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *NewsArticlePost) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *NewsArticlePost) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *NewsArticlePost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewsArticlePost) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *NewsArticlePost) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *NewsArticlePost) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

type NewsArticlePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Country  string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Source   string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Query    string `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
	Page     int64  `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *NewsArticlePostRequest) Reset() {
	*x = NewsArticlePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsArticlePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsArticlePostRequest) ProtoMessage() {}

func (x *NewsArticlePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsArticlePostRequest.ProtoReflect.Descriptor instead.
func (*NewsArticlePostRequest) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{1}
}

func (x *NewsArticlePostRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *NewsArticlePostRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *NewsArticlePostRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *NewsArticlePostRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *NewsArticlePostRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *NewsArticlePostRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type NewsArticlePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*NewsArticlePost `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *NewsArticlePostResponse) Reset() {
	*x = NewsArticlePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsArticlePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsArticlePostResponse) ProtoMessage() {}

func (x *NewsArticlePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsArticlePostResponse.ProtoReflect.Descriptor instead.
func (*NewsArticlePostResponse) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{2}
}

func (x *NewsArticlePostResponse) GetData() []*NewsArticlePost {
	if x != nil {
		return x.Data
	}
	return nil
}

type NewsArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Link   string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *NewsArticleRequest) Reset() {
	*x = NewsArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsArticleRequest) ProtoMessage() {}

func (x *NewsArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsArticleRequest.ProtoReflect.Descriptor instead.
func (*NewsArticleRequest) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{3}
}

func (x *NewsArticleRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *NewsArticleRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type NewsArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *NewsArticlePost `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *NewsArticleResponse) Reset() {
	*x = NewsArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsArticleResponse) ProtoMessage() {}

func (x *NewsArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsArticleResponse.ProtoReflect.Descriptor instead.
func (*NewsArticleResponse) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{4}
}

func (x *NewsArticleResponse) GetData() *NewsArticlePost {
	if x != nil {
		return x.Data
	}
	return nil
}

type NewsTvPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Category    string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Source      string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Gender      string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Video       string `protobuf:"bytes,5,opt,name=video,proto3" json:"video,omitempty"`
	Logo        string `protobuf:"bytes,6,opt,name=logo,proto3" json:"logo,omitempty"`
	Live        bool   `protobuf:"varint,7,opt,name=live,proto3" json:"live,omitempty"`
}

func (x *NewsTvPost) Reset() {
	*x = NewsTvPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsTvPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsTvPost) ProtoMessage() {}

func (x *NewsTvPost) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsTvPost.ProtoReflect.Descriptor instead.
func (*NewsTvPost) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{5}
}

func (x *NewsTvPost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewsTvPost) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *NewsTvPost) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *NewsTvPost) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *NewsTvPost) GetVideo() string {
	if x != nil {
		return x.Video
	}
	return ""
}

func (x *NewsTvPost) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *NewsTvPost) GetLive() bool {
	if x != nil {
		return x.Live
	}
	return false
}

type NewsTvPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Country  string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Source   string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Query    string `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
	Page     int64  `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *NewsTvPostRequest) Reset() {
	*x = NewsTvPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsTvPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsTvPostRequest) ProtoMessage() {}

func (x *NewsTvPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsTvPostRequest.ProtoReflect.Descriptor instead.
func (*NewsTvPostRequest) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{6}
}

func (x *NewsTvPostRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *NewsTvPostRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *NewsTvPostRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *NewsTvPostRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *NewsTvPostRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *NewsTvPostRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type NewsTvPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*NewsTvPost `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *NewsTvPostResponse) Reset() {
	*x = NewsTvPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsTvPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsTvPostResponse) ProtoMessage() {}

func (x *NewsTvPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsTvPostResponse.ProtoReflect.Descriptor instead.
func (*NewsTvPostResponse) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{7}
}

func (x *NewsTvPostResponse) GetData() []*NewsTvPost {
	if x != nil {
		return x.Data
	}
	return nil
}

type NewsCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *NewsCategoryRequest) Reset() {
	*x = NewsCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsCategoryRequest) ProtoMessage() {}

func (x *NewsCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsCategoryRequest.ProtoReflect.Descriptor instead.
func (*NewsCategoryRequest) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{8}
}

func (x *NewsCategoryRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type NewsCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleCategories []string `protobuf:"bytes,1,rep,name=article_categories,json=articleCategories,proto3" json:"article_categories,omitempty"`
	TvCategories      []string `protobuf:"bytes,2,rep,name=tv_categories,json=tvCategories,proto3" json:"tv_categories,omitempty"`
}

func (x *NewsCategoryResponse) Reset() {
	*x = NewsCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsCategoryResponse) ProtoMessage() {}

func (x *NewsCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsCategoryResponse.ProtoReflect.Descriptor instead.
func (*NewsCategoryResponse) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{9}
}

func (x *NewsCategoryResponse) GetArticleCategories() []string {
	if x != nil {
		return x.ArticleCategories
	}
	return nil
}

func (x *NewsCategoryResponse) GetTvCategories() []string {
	if x != nil {
		return x.TvCategories
	}
	return nil
}

var File_schema_proto protoreflect.FileDescriptor

var file_schema_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf,
	0x01, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f,
	0x22, 0xac, 0x01, 0x0a, 0x16, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22,
	0x3f, 0x0a, 0x17, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x40, 0x0a, 0x12, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x22, 0x3b, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0xb8, 0x01, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x73, 0x54, 0x76, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x69, 0x76, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x11, 0x4e,
	0x65, 0x77, 0x73, 0x54, 0x76, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x22, 0x35, 0x0a, 0x12, 0x4e, 0x65, 0x77, 0x73, 0x54, 0x76, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x54,
	0x76, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x13, 0x4e,
	0x65, 0x77, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x6a,
	0x0a, 0x14, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x11, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x76, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x76,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_proto_rawDescOnce sync.Once
	file_schema_proto_rawDescData = file_schema_proto_rawDesc
)

func file_schema_proto_rawDescGZIP() []byte {
	file_schema_proto_rawDescOnce.Do(func() {
		file_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_proto_rawDescData)
	})
	return file_schema_proto_rawDescData
}

var file_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_schema_proto_goTypes = []interface{}{
	(*NewsArticlePost)(nil),         // 0: NewsArticlePost
	(*NewsArticlePostRequest)(nil),  // 1: NewsArticlePostRequest
	(*NewsArticlePostResponse)(nil), // 2: NewsArticlePostResponse
	(*NewsArticleRequest)(nil),      // 3: NewsArticleRequest
	(*NewsArticleResponse)(nil),     // 4: NewsArticleResponse
	(*NewsTvPost)(nil),              // 5: NewsTvPost
	(*NewsTvPostRequest)(nil),       // 6: NewsTvPostRequest
	(*NewsTvPostResponse)(nil),      // 7: NewsTvPostResponse
	(*NewsCategoryRequest)(nil),     // 8: NewsCategoryRequest
	(*NewsCategoryResponse)(nil),    // 9: NewsCategoryResponse
}
var file_schema_proto_depIdxs = []int32{
	0, // 0: NewsArticlePostResponse.data:type_name -> NewsArticlePost
	0, // 1: NewsArticleResponse.data:type_name -> NewsArticlePost
	5, // 2: NewsTvPostResponse.data:type_name -> NewsTvPost
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_schema_proto_init() }
func file_schema_proto_init() {
	if File_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsArticlePost); i {
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
		file_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsArticlePostRequest); i {
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
		file_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsArticlePostResponse); i {
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
		file_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsArticleRequest); i {
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
		file_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsArticleResponse); i {
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
		file_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsTvPost); i {
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
		file_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsTvPostRequest); i {
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
		file_schema_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsTvPostResponse); i {
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
		file_schema_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsCategoryRequest); i {
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
		file_schema_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsCategoryResponse); i {
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
			RawDescriptor: file_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_proto_goTypes,
		DependencyIndexes: file_schema_proto_depIdxs,
		MessageInfos:      file_schema_proto_msgTypes,
	}.Build()
	File_schema_proto = out.File
	file_schema_proto_rawDesc = nil
	file_schema_proto_goTypes = nil
	file_schema_proto_depIdxs = nil
}
