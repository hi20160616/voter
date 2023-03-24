// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.3
// source: api/voter/v1/voter-votes.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListVotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent resource name
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListVotesRequest) Reset() {
	*x = ListVotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_votes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVotesRequest) ProtoMessage() {}

func (x *ListVotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_votes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVotesRequest.ProtoReflect.Descriptor instead.
func (*ListVotesRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_votes_proto_rawDescGZIP(), []int{0}
}

func (x *ListVotesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListVotesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListVotesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListVotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Votes         []*Vote `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes,omitempty"`
	NextPageToken string  `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListVotesResponse) Reset() {
	*x = ListVotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_votes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVotesResponse) ProtoMessage() {}

func (x *ListVotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_votes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVotesResponse.ProtoReflect.Descriptor instead.
func (*ListVotesResponse) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_votes_proto_rawDescGZIP(), []int{1}
}

func (x *ListVotesResponse) GetVotes() []*Vote {
	if x != nil {
		return x.Votes
	}
	return nil
}

func (x *ListVotesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetVoteRequest) Reset() {
	*x = GetVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_votes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoteRequest) ProtoMessage() {}

func (x *GetVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_votes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoteRequest.ProtoReflect.Descriptor instead.
func (*GetVoteRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_votes_proto_rawDescGZIP(), []int{2}
}

func (x *GetVoteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SearchVotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *SearchVotesRequest) Reset() {
	*x = SearchVotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_votes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchVotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchVotesRequest) ProtoMessage() {}

func (x *SearchVotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_votes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchVotesRequest.ProtoReflect.Descriptor instead.
func (*SearchVotesRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_votes_proto_rawDescGZIP(), []int{3}
}

func (x *SearchVotesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchVotesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchVotesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type SearchVotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Votes         []*Vote `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes,omitempty"`
	NextPageToken string  `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *SearchVotesResponse) Reset() {
	*x = SearchVotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_votes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchVotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchVotesResponse) ProtoMessage() {}

func (x *SearchVotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_votes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchVotesResponse.ProtoReflect.Descriptor instead.
func (*SearchVotesResponse) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_votes_proto_rawDescGZIP(), []int{4}
}

func (x *SearchVotesResponse) GetVotes() []*Vote {
	if x != nil {
		return x.Votes
	}
	return nil
}

func (x *SearchVotesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Vote   *Vote  `protobuf:"bytes,2,opt,name=vote,proto3" json:"vote,omitempty"`
}

func (x *CreateVoteRequest) Reset() {
	*x = CreateVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_votes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVoteRequest) ProtoMessage() {}

func (x *CreateVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_votes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVoteRequest.ProtoReflect.Descriptor instead.
func (*CreateVoteRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_votes_proto_rawDescGZIP(), []int{5}
}

func (x *CreateVoteRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateVoteRequest) GetVote() *Vote {
	if x != nil {
		return x.Vote
	}
	return nil
}

type UpdateVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vote *Vote `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
	// The update mask applies to the resource. For the `FieldMask` definition,
	// see https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateVoteRequest) Reset() {
	*x = UpdateVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_votes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVoteRequest) ProtoMessage() {}

func (x *UpdateVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_votes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateVoteRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_votes_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateVoteRequest) GetVote() *Vote {
	if x != nil {
		return x.Vote
	}
	return nil
}

func (x *UpdateVoteRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteVoteRequest) Reset() {
	*x = DeleteVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_votes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVoteRequest) ProtoMessage() {}

func (x *DeleteVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_votes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVoteRequest.ProtoReflect.Descriptor instead.
func (*DeleteVoteRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_votes_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteVoteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VoteId      int32                  `protobuf:"varint,2,opt,name=vote_id,json=voteId,proto3" json:"vote_id,omitempty"`
	Title       string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	IsRadio     int32                  `protobuf:"varint,4,opt,name=is_radio,json=isRadio,proto3" json:"is_radio,omitempty"`
	A           string                 `protobuf:"bytes,5,opt,name=a,proto3" json:"a,omitempty"`
	B           string                 `protobuf:"bytes,6,opt,name=b,proto3" json:"b,omitempty"`
	C           string                 `protobuf:"bytes,7,opt,name=c,proto3" json:"c,omitempty"`
	D           string                 `protobuf:"bytes,8,opt,name=d,proto3" json:"d,omitempty"`
	E           string                 `protobuf:"bytes,9,opt,name=e,proto3" json:"e,omitempty"`
	F           string                 `protobuf:"bytes,10,opt,name=f,proto3" json:"f,omitempty"`
	G           string                 `protobuf:"bytes,11,opt,name=g,proto3" json:"g,omitempty"`
	H           string                 `protobuf:"bytes,12,opt,name=h,proto3" json:"h,omitempty"`
	HasTxtField int32                  `protobuf:"varint,13,opt,name=HasTxtField,proto3" json:"HasTxtField,omitempty"`
	CreateTime  *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime  *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_votes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_votes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_votes_proto_rawDescGZIP(), []int{8}
}

func (x *Vote) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Vote) GetVoteId() int32 {
	if x != nil {
		return x.VoteId
	}
	return 0
}

func (x *Vote) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Vote) GetIsRadio() int32 {
	if x != nil {
		return x.IsRadio
	}
	return 0
}

func (x *Vote) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *Vote) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

func (x *Vote) GetC() string {
	if x != nil {
		return x.C
	}
	return ""
}

func (x *Vote) GetD() string {
	if x != nil {
		return x.D
	}
	return ""
}

func (x *Vote) GetE() string {
	if x != nil {
		return x.E
	}
	return ""
}

func (x *Vote) GetF() string {
	if x != nil {
		return x.F
	}
	return ""
}

func (x *Vote) GetG() string {
	if x != nil {
		return x.G
	}
	return ""
}

func (x *Vote) GetH() string {
	if x != nil {
		return x.H
	}
	return ""
}

func (x *Vote) GetHasTxtField() int32 {
	if x != nil {
		return x.HasTxtField
	}
	return 0
}

func (x *Vote) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Vote) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_api_voter_v1_voter_votes_proto protoreflect.FileDescriptor

var file_api_voter_v1_voter_votes_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x76,
	0x6f, 0x74, 0x65, 0x72, 0x2d, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x67, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76,
	0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x64, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x56,
	0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x69, 0x0a, 0x13, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x55, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x22, 0x7a, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0xf0, 0x02, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x76, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x73, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x69, 0x73, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x63, 0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x64, 0x12, 0x0c, 0x0a, 0x01, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x65, 0x12,
	0x0c, 0x0a, 0x01, 0x66, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x66, 0x12, 0x0c, 0x0a,
	0x01, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x67, 0x12, 0x0c, 0x0a, 0x01, 0x68,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x61, 0x73,
	0x54, 0x78, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x48, 0x61, 0x73, 0x54, 0x78, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x8e, 0x06, 0x0a, 0x08, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x41,
	0x50, 0x49, 0x12, 0xd7, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73,
	0x12, 0x20, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x84, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x7e, 0x12, 0x19,
	0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x72, 0x6f, 0x6c, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x5a, 0x21, 0x12, 0x1f, 0x2f, 0x76, 0x31,
	0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x76, 0x6f, 0x74, 0x65, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x5a, 0x22, 0x12, 0x20,
	0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x5a, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d,
	0x70, 0x69, 0x64, 0x2f, 0x2a, 0x7d, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x79, 0x0a, 0x0b, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76,
	0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x7d, 0x12, 0x6f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f,
	0x74, 0x65, 0x12, 0x21, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x6f,
	0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x22, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x3d, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x7d, 0x3a,
	0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x73, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x6f, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x6f, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x22, 0x2c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x26, 0x32, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x76, 0x6f, 0x74, 0x65, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x7d, 0x3a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x6a, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x2a, 0x19, 0x2f, 0x76, 0x31,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x7d, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x32, 0x30, 0x31, 0x36, 0x30, 0x36, 0x31, 0x36, 0x2f,
	0x76, 0x6f, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_voter_v1_voter_votes_proto_rawDescOnce sync.Once
	file_api_voter_v1_voter_votes_proto_rawDescData = file_api_voter_v1_voter_votes_proto_rawDesc
)

func file_api_voter_v1_voter_votes_proto_rawDescGZIP() []byte {
	file_api_voter_v1_voter_votes_proto_rawDescOnce.Do(func() {
		file_api_voter_v1_voter_votes_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_voter_v1_voter_votes_proto_rawDescData)
	})
	return file_api_voter_v1_voter_votes_proto_rawDescData
}

var file_api_voter_v1_voter_votes_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_voter_v1_voter_votes_proto_goTypes = []interface{}{
	(*ListVotesRequest)(nil),      // 0: voter.votes.v1.ListVotesRequest
	(*ListVotesResponse)(nil),     // 1: voter.votes.v1.ListVotesResponse
	(*GetVoteRequest)(nil),        // 2: voter.votes.v1.GetVoteRequest
	(*SearchVotesRequest)(nil),    // 3: voter.votes.v1.SearchVotesRequest
	(*SearchVotesResponse)(nil),   // 4: voter.votes.v1.SearchVotesResponse
	(*CreateVoteRequest)(nil),     // 5: voter.votes.v1.CreateVoteRequest
	(*UpdateVoteRequest)(nil),     // 6: voter.votes.v1.UpdateVoteRequest
	(*DeleteVoteRequest)(nil),     // 7: voter.votes.v1.DeleteVoteRequest
	(*Vote)(nil),                  // 8: voter.votes.v1.Vote
	(*fieldmaskpb.FieldMask)(nil), // 9: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_api_voter_v1_voter_votes_proto_depIdxs = []int32{
	8,  // 0: voter.votes.v1.ListVotesResponse.votes:type_name -> voter.votes.v1.Vote
	8,  // 1: voter.votes.v1.SearchVotesResponse.votes:type_name -> voter.votes.v1.Vote
	8,  // 2: voter.votes.v1.CreateVoteRequest.vote:type_name -> voter.votes.v1.Vote
	8,  // 3: voter.votes.v1.UpdateVoteRequest.vote:type_name -> voter.votes.v1.Vote
	9,  // 4: voter.votes.v1.UpdateVoteRequest.update_mask:type_name -> google.protobuf.FieldMask
	10, // 5: voter.votes.v1.Vote.create_time:type_name -> google.protobuf.Timestamp
	10, // 6: voter.votes.v1.Vote.update_time:type_name -> google.protobuf.Timestamp
	0,  // 7: voter.votes.v1.VotesAPI.ListVotes:input_type -> voter.votes.v1.ListVotesRequest
	2,  // 8: voter.votes.v1.VotesAPI.GetVote:input_type -> voter.votes.v1.GetVoteRequest
	3,  // 9: voter.votes.v1.VotesAPI.SearchVotes:input_type -> voter.votes.v1.SearchVotesRequest
	5,  // 10: voter.votes.v1.VotesAPI.CreateVote:input_type -> voter.votes.v1.CreateVoteRequest
	6,  // 11: voter.votes.v1.VotesAPI.UpdateVote:input_type -> voter.votes.v1.UpdateVoteRequest
	7,  // 12: voter.votes.v1.VotesAPI.DeleteVote:input_type -> voter.votes.v1.DeleteVoteRequest
	1,  // 13: voter.votes.v1.VotesAPI.ListVotes:output_type -> voter.votes.v1.ListVotesResponse
	8,  // 14: voter.votes.v1.VotesAPI.GetVote:output_type -> voter.votes.v1.Vote
	4,  // 15: voter.votes.v1.VotesAPI.SearchVotes:output_type -> voter.votes.v1.SearchVotesResponse
	8,  // 16: voter.votes.v1.VotesAPI.CreateVote:output_type -> voter.votes.v1.Vote
	8,  // 17: voter.votes.v1.VotesAPI.UpdateVote:output_type -> voter.votes.v1.Vote
	11, // 18: voter.votes.v1.VotesAPI.DeleteVote:output_type -> google.protobuf.Empty
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_voter_v1_voter_votes_proto_init() }
func file_api_voter_v1_voter_votes_proto_init() {
	if File_api_voter_v1_voter_votes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_voter_v1_voter_votes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVotesRequest); i {
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
		file_api_voter_v1_voter_votes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVotesResponse); i {
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
		file_api_voter_v1_voter_votes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVoteRequest); i {
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
		file_api_voter_v1_voter_votes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchVotesRequest); i {
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
		file_api_voter_v1_voter_votes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchVotesResponse); i {
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
		file_api_voter_v1_voter_votes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVoteRequest); i {
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
		file_api_voter_v1_voter_votes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVoteRequest); i {
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
		file_api_voter_v1_voter_votes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVoteRequest); i {
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
		file_api_voter_v1_voter_votes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
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
			RawDescriptor: file_api_voter_v1_voter_votes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_voter_v1_voter_votes_proto_goTypes,
		DependencyIndexes: file_api_voter_v1_voter_votes_proto_depIdxs,
		MessageInfos:      file_api_voter_v1_voter_votes_proto_msgTypes,
	}.Build()
	File_api_voter_v1_voter_votes_proto = out.File
	file_api_voter_v1_voter_votes_proto_rawDesc = nil
	file_api_voter_v1_voter_votes_proto_goTypes = nil
	file_api_voter_v1_voter_votes_proto_depIdxs = nil
}
