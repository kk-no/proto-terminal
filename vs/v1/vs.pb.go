// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: vs/v1/vs.proto

package v1

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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vs_v1_vs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vs_v1_vs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_vs_v1_vs_proto_rawDescGZIP(), []int{0}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vs_v1_vs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vs_v1_vs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_vs_v1_vs_proto_rawDescGZIP(), []int{1}
}

type ListUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vs_v1_vs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vs_v1_vs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_vs_v1_vs_proto_rawDescGZIP(), []int{2}
}

type ListUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vs_v1_vs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vs_v1_vs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_vs_v1_vs_proto_rawDescGZIP(), []int{3}
}

func (x *ListUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserWithGameResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserWithGameResultsRequest) Reset() {
	*x = GetUserWithGameResultsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vs_v1_vs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserWithGameResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserWithGameResultsRequest) ProtoMessage() {}

func (x *GetUserWithGameResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vs_v1_vs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserWithGameResultsRequest.ProtoReflect.Descriptor instead.
func (*GetUserWithGameResultsRequest) Descriptor() ([]byte, []int) {
	return file_vs_v1_vs_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserWithGameResultsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserWithGameResultsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        *User         `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	GameResults []*GameResult `protobuf:"bytes,2,rep,name=game_results,json=gameResults,proto3" json:"game_results,omitempty"`
}

func (x *GetUserWithGameResultsResponse) Reset() {
	*x = GetUserWithGameResultsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vs_v1_vs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserWithGameResultsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserWithGameResultsResponse) ProtoMessage() {}

func (x *GetUserWithGameResultsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vs_v1_vs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserWithGameResultsResponse.ProtoReflect.Descriptor instead.
func (*GetUserWithGameResultsResponse) Descriptor() ([]byte, []int) {
	return file_vs_v1_vs_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserWithGameResultsResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetUserWithGameResultsResponse) GetGameResults() []*GameResult {
	if x != nil {
		return x.GameResults
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string  `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string  `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Win                  uint32  `protobuf:"varint,4,opt,name=win,proto3" json:"win,omitempty"`
	Lose                 uint32  `protobuf:"varint,5,opt,name=lose,proto3" json:"lose,omitempty"`
	WinRate              float32 `protobuf:"fixed32,6,opt,name=win_rate,json=winRate,proto3" json:"win_rate,omitempty"`
	Rating               uint32  `protobuf:"varint,7,opt,name=rating,proto3" json:"rating,omitempty"`
	MaxRating            uint32  `protobuf:"varint,8,opt,name=max_rating,json=maxRating,proto3" json:"max_rating,omitempty"`
	Privilege            string  `protobuf:"bytes,9,opt,name=privilege,proto3" json:"privilege,omitempty"`
	ContinuousAttendance string  `protobuf:"bytes,10,opt,name=continuous_attendance,json=continuousAttendance,proto3" json:"continuous_attendance,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vs_v1_vs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_vs_v1_vs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_vs_v1_vs_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetWin() uint32 {
	if x != nil {
		return x.Win
	}
	return 0
}

func (x *User) GetLose() uint32 {
	if x != nil {
		return x.Lose
	}
	return 0
}

func (x *User) GetWinRate() float32 {
	if x != nil {
		return x.WinRate
	}
	return 0
}

func (x *User) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *User) GetMaxRating() uint32 {
	if x != nil {
		return x.MaxRating
	}
	return 0
}

func (x *User) GetPrivilege() string {
	if x != nil {
		return x.Privilege
	}
	return ""
}

func (x *User) GetContinuousAttendance() string {
	if x != nil {
		return x.ContinuousAttendance
	}
	return ""
}

type GameResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Date              string             `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Player1           *User              `protobuf:"bytes,3,opt,name=player1,proto3" json:"player1,omitempty"`
	Player2           *User              `protobuf:"bytes,4,opt,name=player2,proto3" json:"player2,omitempty"`
	WinUserId         string             `protobuf:"bytes,5,opt,name=win_user_id,json=winUserId,proto3" json:"win_user_id,omitempty"`
	Note              string             `protobuf:"bytes,6,opt,name=note,proto3" json:"note,omitempty"`
	RatingFluctuation *RatingFluctuation `protobuf:"bytes,7,opt,name=rating_fluctuation,json=ratingFluctuation,proto3" json:"rating_fluctuation,omitempty"`
}

func (x *GameResult) Reset() {
	*x = GameResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vs_v1_vs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResult) ProtoMessage() {}

func (x *GameResult) ProtoReflect() protoreflect.Message {
	mi := &file_vs_v1_vs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameResult.ProtoReflect.Descriptor instead.
func (*GameResult) Descriptor() ([]byte, []int) {
	return file_vs_v1_vs_proto_rawDescGZIP(), []int{7}
}

func (x *GameResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GameResult) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *GameResult) GetPlayer1() *User {
	if x != nil {
		return x.Player1
	}
	return nil
}

func (x *GameResult) GetPlayer2() *User {
	if x != nil {
		return x.Player2
	}
	return nil
}

func (x *GameResult) GetWinUserId() string {
	if x != nil {
		return x.WinUserId
	}
	return ""
}

func (x *GameResult) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *GameResult) GetRatingFluctuation() *RatingFluctuation {
	if x != nil {
		return x.RatingFluctuation
	}
	return nil
}

type RatingFluctuation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Winner int32 `protobuf:"varint,1,opt,name=winner,proto3" json:"winner,omitempty"`
	Loser  int32 `protobuf:"varint,2,opt,name=loser,proto3" json:"loser,omitempty"`
}

func (x *RatingFluctuation) Reset() {
	*x = RatingFluctuation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vs_v1_vs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingFluctuation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingFluctuation) ProtoMessage() {}

func (x *RatingFluctuation) ProtoReflect() protoreflect.Message {
	mi := &file_vs_v1_vs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingFluctuation.ProtoReflect.Descriptor instead.
func (*RatingFluctuation) Descriptor() ([]byte, []int) {
	return file_vs_v1_vs_proto_rawDescGZIP(), []int{8}
}

func (x *RatingFluctuation) GetWinner() int32 {
	if x != nil {
		return x.Winner
	}
	return 0
}

func (x *RatingFluctuation) GetLoser() int32 {
	if x != nil {
		return x.Loser
	}
	return 0
}

var File_vs_v1_vs_proto protoreflect.FileDescriptor

var file_vs_v1_vs_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x76, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x76, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x76, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x38, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74,
	0x68, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x77, 0x0a, 0x1e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x34, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x9d, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x77, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x6f, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6c, 0x6f, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x77, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65,
	0x12, 0x33, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73, 0x5f, 0x61,
	0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73, 0x41, 0x74, 0x74, 0x65, 0x6e,
	0x64, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xfb, 0x01, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x12,
	0x25, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x76, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x12, 0x1e, 0x0a, 0x0b, 0x77, 0x69, 0x6e, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x69, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x47, 0x0a, 0x12, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x6c, 0x75, 0x63, 0x74, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x6c, 0x75, 0x63, 0x74, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x11, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x6c, 0x75, 0x63, 0x74, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x11, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x6c, 0x75,
	0x63, 0x74, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x32, 0xe9, 0x01, 0x0a, 0x09, 0x56, 0x53, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x76,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x76, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x76, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x76, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x76, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x6b, 0x2d, 0x6e, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_vs_v1_vs_proto_rawDescOnce sync.Once
	file_vs_v1_vs_proto_rawDescData = file_vs_v1_vs_proto_rawDesc
)

func file_vs_v1_vs_proto_rawDescGZIP() []byte {
	file_vs_v1_vs_proto_rawDescOnce.Do(func() {
		file_vs_v1_vs_proto_rawDescData = protoimpl.X.CompressGZIP(file_vs_v1_vs_proto_rawDescData)
	})
	return file_vs_v1_vs_proto_rawDescData
}

var file_vs_v1_vs_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_vs_v1_vs_proto_goTypes = []interface{}{
	(*PingRequest)(nil),                    // 0: vs.v1.PingRequest
	(*PingResponse)(nil),                   // 1: vs.v1.PingResponse
	(*ListUsersRequest)(nil),               // 2: vs.v1.ListUsersRequest
	(*ListUsersResponse)(nil),              // 3: vs.v1.ListUsersResponse
	(*GetUserWithGameResultsRequest)(nil),  // 4: vs.v1.GetUserWithGameResultsRequest
	(*GetUserWithGameResultsResponse)(nil), // 5: vs.v1.GetUserWithGameResultsResponse
	(*User)(nil),                           // 6: vs.v1.User
	(*GameResult)(nil),                     // 7: vs.v1.GameResult
	(*RatingFluctuation)(nil),              // 8: vs.v1.RatingFluctuation
}
var file_vs_v1_vs_proto_depIdxs = []int32{
	6, // 0: vs.v1.ListUsersResponse.users:type_name -> vs.v1.User
	6, // 1: vs.v1.GetUserWithGameResultsResponse.user:type_name -> vs.v1.User
	7, // 2: vs.v1.GetUserWithGameResultsResponse.game_results:type_name -> vs.v1.GameResult
	6, // 3: vs.v1.GameResult.player1:type_name -> vs.v1.User
	6, // 4: vs.v1.GameResult.player2:type_name -> vs.v1.User
	8, // 5: vs.v1.GameResult.rating_fluctuation:type_name -> vs.v1.RatingFluctuation
	0, // 6: vs.v1.VSService.Ping:input_type -> vs.v1.PingRequest
	2, // 7: vs.v1.VSService.ListUsers:input_type -> vs.v1.ListUsersRequest
	4, // 8: vs.v1.VSService.GetUserWithGameResults:input_type -> vs.v1.GetUserWithGameResultsRequest
	1, // 9: vs.v1.VSService.Ping:output_type -> vs.v1.PingResponse
	3, // 10: vs.v1.VSService.ListUsers:output_type -> vs.v1.ListUsersResponse
	5, // 11: vs.v1.VSService.GetUserWithGameResults:output_type -> vs.v1.GetUserWithGameResultsResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_vs_v1_vs_proto_init() }
func file_vs_v1_vs_proto_init() {
	if File_vs_v1_vs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vs_v1_vs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_vs_v1_vs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_vs_v1_vs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersRequest); i {
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
		file_vs_v1_vs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersResponse); i {
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
		file_vs_v1_vs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserWithGameResultsRequest); i {
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
		file_vs_v1_vs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserWithGameResultsResponse); i {
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
		file_vs_v1_vs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_vs_v1_vs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameResult); i {
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
		file_vs_v1_vs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingFluctuation); i {
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
			RawDescriptor: file_vs_v1_vs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vs_v1_vs_proto_goTypes,
		DependencyIndexes: file_vs_v1_vs_proto_depIdxs,
		MessageInfos:      file_vs_v1_vs_proto_msgTypes,
	}.Build()
	File_vs_v1_vs_proto = out.File
	file_vs_v1_vs_proto_rawDesc = nil
	file_vs_v1_vs_proto_goTypes = nil
	file_vs_v1_vs_proto_depIdxs = nil
}
