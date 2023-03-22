// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.3
// source: api/voter/v1/voter-post_votes.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListPostVotesRequest struct {
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

func (x *ListPostVotesRequest) Reset() {
	*x = ListPostVotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostVotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostVotesRequest) ProtoMessage() {}

func (x *ListPostVotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostVotesRequest.ProtoReflect.Descriptor instead.
func (*ListPostVotesRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_post_votes_proto_rawDescGZIP(), []int{0}
}

func (x *ListPostVotesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListPostVotesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPostVotesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListPostVotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostVotes     []*PostVote `protobuf:"bytes,1,rep,name=postVotes,proto3" json:"postVotes,omitempty"`
	NextPageToken string      `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPostVotesResponse) Reset() {
	*x = ListPostVotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostVotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostVotesResponse) ProtoMessage() {}

func (x *ListPostVotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostVotesResponse.ProtoReflect.Descriptor instead.
func (*ListPostVotesResponse) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_post_votes_proto_rawDescGZIP(), []int{1}
}

func (x *ListPostVotesResponse) GetPostVotes() []*PostVote {
	if x != nil {
		return x.PostVotes
	}
	return nil
}

func (x *ListPostVotesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetPostVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPostVoteRequest) Reset() {
	*x = GetPostVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostVoteRequest) ProtoMessage() {}

func (x *GetPostVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostVoteRequest.ProtoReflect.Descriptor instead.
func (*GetPostVoteRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_post_votes_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostVoteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetByPidVidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetByPidVidRequest) Reset() {
	*x = GetByPidVidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByPidVidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByPidVidRequest) ProtoMessage() {}

func (x *GetByPidVidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByPidVidRequest.ProtoReflect.Descriptor instead.
func (*GetByPidVidRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_post_votes_proto_rawDescGZIP(), []int{3}
}

func (x *GetByPidVidRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListVidsByPidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListVidsByPidRequest) Reset() {
	*x = ListVidsByPidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVidsByPidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVidsByPidRequest) ProtoMessage() {}

func (x *ListVidsByPidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVidsByPidRequest.ProtoReflect.Descriptor instead.
func (*ListVidsByPidRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_post_votes_proto_rawDescGZIP(), []int{4}
}

func (x *ListVidsByPidRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListVidsByPidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid  int32   `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Vids []int32 `protobuf:"varint,2,rep,packed,name=Vids,proto3" json:"Vids,omitempty"`
}

func (x *ListVidsByPidResponse) Reset() {
	*x = ListVidsByPidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVidsByPidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVidsByPidResponse) ProtoMessage() {}

func (x *ListVidsByPidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVidsByPidResponse.ProtoReflect.Descriptor instead.
func (*ListVidsByPidResponse) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_post_votes_proto_rawDescGZIP(), []int{5}
}

func (x *ListVidsByPidResponse) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ListVidsByPidResponse) GetVids() []int32 {
	if x != nil {
		return x.Vids
	}
	return nil
}

type CreatePostVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent     string    `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	PostVoteId int32     `protobuf:"varint,2,opt,name=post_vote_id,json=postVoteId,proto3" json:"post_vote_id,omitempty"`
	PostVote   *PostVote `protobuf:"bytes,3,opt,name=post_vote,json=postVote,proto3" json:"post_vote,omitempty"`
}

func (x *CreatePostVoteRequest) Reset() {
	*x = CreatePostVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostVoteRequest) ProtoMessage() {}

func (x *CreatePostVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostVoteRequest.ProtoReflect.Descriptor instead.
func (*CreatePostVoteRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_post_votes_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePostVoteRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreatePostVoteRequest) GetPostVoteId() int32 {
	if x != nil {
		return x.PostVoteId
	}
	return 0
}

func (x *CreatePostVoteRequest) GetPostVote() *PostVote {
	if x != nil {
		return x.PostVote
	}
	return nil
}

type UpdatePostVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostVote *PostVote `protobuf:"bytes,1,opt,name=post_vote,json=postVote,proto3" json:"post_vote,omitempty"`
	// The update mask applies to the resource. For the `FieldMask` definition,
	// see https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdatePostVoteRequest) Reset() {
	*x = UpdatePostVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostVoteRequest) ProtoMessage() {}

func (x *UpdatePostVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostVoteRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostVoteRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_post_votes_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePostVoteRequest) GetPostVote() *PostVote {
	if x != nil {
		return x.PostVote
	}
	return nil
}

func (x *UpdatePostVoteRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeletePostVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeletePostVoteRequest) Reset() {
	*x = DeletePostVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostVoteRequest) ProtoMessage() {}

func (x *DeletePostVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostVoteRequest.ProtoReflect.Descriptor instead.
func (*DeletePostVoteRequest) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_post_votes_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePostVoteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PostVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PostVoteId int32  `protobuf:"varint,2,opt,name=post_vote_id,json=postVoteId,proto3" json:"post_vote_id,omitempty"`
	PostId     int32  `protobuf:"varint,3,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	VoteId     int32  `protobuf:"varint,4,opt,name=vote_id,json=voteId,proto3" json:"vote_id,omitempty"`
}

func (x *PostVote) Reset() {
	*x = PostVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostVote) ProtoMessage() {}

func (x *PostVote) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostVote.ProtoReflect.Descriptor instead.
func (*PostVote) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_post_votes_proto_rawDescGZIP(), []int{9}
}

func (x *PostVote) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PostVote) GetPostVoteId() int32 {
	if x != nil {
		return x.PostVoteId
	}
	return 0
}

func (x *PostVote) GetPostId() int32 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *PostVote) GetVoteId() int32 {
	if x != nil {
		return x.VoteId
	}
	return 0
}

type PostVotes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostVotes []*PostVote `protobuf:"bytes,1,rep,name=post_votes,json=postVotes,proto3" json:"post_votes,omitempty"`
}

func (x *PostVotes) Reset() {
	*x = PostVotes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostVotes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostVotes) ProtoMessage() {}

func (x *PostVotes) ProtoReflect() protoreflect.Message {
	mi := &file_api_voter_v1_voter_post_votes_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostVotes.ProtoReflect.Descriptor instead.
func (*PostVotes) Descriptor() ([]byte, []int) {
	return file_api_voter_v1_voter_post_votes_proto_rawDescGZIP(), []int{10}
}

func (x *PostVotes) GetPostVotes() []*PostVote {
	if x != nil {
		return x.PostVotes
	}
	return nil
}

var File_api_voter_v1_voter_post_votes_proto protoreflect.FileDescriptor

var file_api_voter_v1_voter_post_votes_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x76,
	0x6f, 0x74, 0x65, 0x72, 0x2d, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x7c, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x56,
	0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09,
	0x70, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x09,
	0x70, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x50, 0x69, 0x64, 0x56, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64,
	0x73, 0x42, 0x79, 0x50, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x3d, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x73, 0x42, 0x79, 0x50,
	0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x56, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x56, 0x69, 0x64, 0x73,
	0x22, 0x8d, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x22, 0x90, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x73, 0x6b, 0x22, 0x2b, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x72, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x76,
	0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x76, 0x6f,
	0x74, 0x65, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x32,
	0xc9, 0x07, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x41, 0x50, 0x49,
	0x12, 0x88, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74,
	0x65, 0x73, 0x12, 0x29, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x7d, 0x12, 0x7b, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x50, 0x69, 0x64, 0x56, 0x69, 0x64, 0x12, 0x27, 0x2e, 0x76, 0x6f, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x69, 0x64, 0x56, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x2f, 0x2a, 0x2f, 0x2a, 0x2f, 0x69, 0x64, 0x7d, 0x12, 0x76, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f,
	0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x2a, 0x7d,
	0x12, 0x8c, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x73, 0x42, 0x79, 0x50,
	0x69, 0x64, 0x12, 0x29, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64,
	0x73, 0x42, 0x79, 0x50, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x73, 0x42, 0x79, 0x50, 0x69,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x7d, 0x12,
	0x90, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x12, 0x2a, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x22, 0x33, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x2a, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x7d, 0x3a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f,
	0x74, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76,
	0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x32, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x7d, 0x3a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x7c, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12,
	0x2a, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x76, 0x31,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65,
	0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x7d, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x32, 0x30, 0x31, 0x36,
	0x30, 0x36, 0x31, 0x36, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x6f, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_voter_v1_voter_post_votes_proto_rawDescOnce sync.Once
	file_api_voter_v1_voter_post_votes_proto_rawDescData = file_api_voter_v1_voter_post_votes_proto_rawDesc
)

func file_api_voter_v1_voter_post_votes_proto_rawDescGZIP() []byte {
	file_api_voter_v1_voter_post_votes_proto_rawDescOnce.Do(func() {
		file_api_voter_v1_voter_post_votes_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_voter_v1_voter_post_votes_proto_rawDescData)
	})
	return file_api_voter_v1_voter_post_votes_proto_rawDescData
}

var file_api_voter_v1_voter_post_votes_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_voter_v1_voter_post_votes_proto_goTypes = []interface{}{
	(*ListPostVotesRequest)(nil),  // 0: voter.post_votes.v1.ListPostVotesRequest
	(*ListPostVotesResponse)(nil), // 1: voter.post_votes.v1.ListPostVotesResponse
	(*GetPostVoteRequest)(nil),    // 2: voter.post_votes.v1.GetPostVoteRequest
	(*GetByPidVidRequest)(nil),    // 3: voter.post_votes.v1.GetByPidVidRequest
	(*ListVidsByPidRequest)(nil),  // 4: voter.post_votes.v1.ListVidsByPidRequest
	(*ListVidsByPidResponse)(nil), // 5: voter.post_votes.v1.ListVidsByPidResponse
	(*CreatePostVoteRequest)(nil), // 6: voter.post_votes.v1.CreatePostVoteRequest
	(*UpdatePostVoteRequest)(nil), // 7: voter.post_votes.v1.UpdatePostVoteRequest
	(*DeletePostVoteRequest)(nil), // 8: voter.post_votes.v1.DeletePostVoteRequest
	(*PostVote)(nil),              // 9: voter.post_votes.v1.PostVote
	(*PostVotes)(nil),             // 10: voter.post_votes.v1.PostVotes
	(*fieldmaskpb.FieldMask)(nil), // 11: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_api_voter_v1_voter_post_votes_proto_depIdxs = []int32{
	9,  // 0: voter.post_votes.v1.ListPostVotesResponse.postVotes:type_name -> voter.post_votes.v1.PostVote
	9,  // 1: voter.post_votes.v1.CreatePostVoteRequest.post_vote:type_name -> voter.post_votes.v1.PostVote
	9,  // 2: voter.post_votes.v1.UpdatePostVoteRequest.post_vote:type_name -> voter.post_votes.v1.PostVote
	11, // 3: voter.post_votes.v1.UpdatePostVoteRequest.update_mask:type_name -> google.protobuf.FieldMask
	9,  // 4: voter.post_votes.v1.PostVotes.post_votes:type_name -> voter.post_votes.v1.PostVote
	0,  // 5: voter.post_votes.v1.PostVotesAPI.ListPostVotes:input_type -> voter.post_votes.v1.ListPostVotesRequest
	3,  // 6: voter.post_votes.v1.PostVotesAPI.GetByPidVid:input_type -> voter.post_votes.v1.GetByPidVidRequest
	2,  // 7: voter.post_votes.v1.PostVotesAPI.GetPostVote:input_type -> voter.post_votes.v1.GetPostVoteRequest
	4,  // 8: voter.post_votes.v1.PostVotesAPI.ListVidsByPid:input_type -> voter.post_votes.v1.ListVidsByPidRequest
	6,  // 9: voter.post_votes.v1.PostVotesAPI.CreatePostVote:input_type -> voter.post_votes.v1.CreatePostVoteRequest
	7,  // 10: voter.post_votes.v1.PostVotesAPI.UpdatePostVote:input_type -> voter.post_votes.v1.UpdatePostVoteRequest
	8,  // 11: voter.post_votes.v1.PostVotesAPI.DeletePostVote:input_type -> voter.post_votes.v1.DeletePostVoteRequest
	1,  // 12: voter.post_votes.v1.PostVotesAPI.ListPostVotes:output_type -> voter.post_votes.v1.ListPostVotesResponse
	9,  // 13: voter.post_votes.v1.PostVotesAPI.GetByPidVid:output_type -> voter.post_votes.v1.PostVote
	9,  // 14: voter.post_votes.v1.PostVotesAPI.GetPostVote:output_type -> voter.post_votes.v1.PostVote
	5,  // 15: voter.post_votes.v1.PostVotesAPI.ListVidsByPid:output_type -> voter.post_votes.v1.ListVidsByPidResponse
	9,  // 16: voter.post_votes.v1.PostVotesAPI.CreatePostVote:output_type -> voter.post_votes.v1.PostVote
	9,  // 17: voter.post_votes.v1.PostVotesAPI.UpdatePostVote:output_type -> voter.post_votes.v1.PostVote
	12, // 18: voter.post_votes.v1.PostVotesAPI.DeletePostVote:output_type -> google.protobuf.Empty
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_voter_v1_voter_post_votes_proto_init() }
func file_api_voter_v1_voter_post_votes_proto_init() {
	if File_api_voter_v1_voter_post_votes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_voter_v1_voter_post_votes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostVotesRequest); i {
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
		file_api_voter_v1_voter_post_votes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostVotesResponse); i {
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
		file_api_voter_v1_voter_post_votes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostVoteRequest); i {
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
		file_api_voter_v1_voter_post_votes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByPidVidRequest); i {
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
		file_api_voter_v1_voter_post_votes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVidsByPidRequest); i {
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
		file_api_voter_v1_voter_post_votes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVidsByPidResponse); i {
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
		file_api_voter_v1_voter_post_votes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostVoteRequest); i {
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
		file_api_voter_v1_voter_post_votes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostVoteRequest); i {
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
		file_api_voter_v1_voter_post_votes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostVoteRequest); i {
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
		file_api_voter_v1_voter_post_votes_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostVote); i {
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
		file_api_voter_v1_voter_post_votes_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostVotes); i {
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
			RawDescriptor: file_api_voter_v1_voter_post_votes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_voter_v1_voter_post_votes_proto_goTypes,
		DependencyIndexes: file_api_voter_v1_voter_post_votes_proto_depIdxs,
		MessageInfos:      file_api_voter_v1_voter_post_votes_proto_msgTypes,
	}.Build()
	File_api_voter_v1_voter_post_votes_proto = out.File
	file_api_voter_v1_voter_post_votes_proto_rawDesc = nil
	file_api_voter_v1_voter_post_votes_proto_goTypes = nil
	file_api_voter_v1_voter_post_votes_proto_depIdxs = nil
}
