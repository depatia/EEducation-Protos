// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: grades.proto

package grades

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

type SetGradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId  int64  `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	LessonName string `protobuf:"bytes,2,opt,name=lesson_name,json=lessonName,proto3" json:"lesson_name,omitempty"`
	Grade      int64  `protobuf:"varint,3,opt,name=grade,proto3" json:"grade,omitempty"`
	IsTerm     bool   `protobuf:"varint,4,opt,name=is_term,json=isTerm,proto3" json:"is_term,omitempty"`
}

func (x *SetGradeRequest) Reset() {
	*x = SetGradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetGradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGradeRequest) ProtoMessage() {}

func (x *SetGradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGradeRequest.ProtoReflect.Descriptor instead.
func (*SetGradeRequest) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{0}
}

func (x *SetGradeRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *SetGradeRequest) GetLessonName() string {
	if x != nil {
		return x.LessonName
	}
	return ""
}

func (x *SetGradeRequest) GetGrade() int64 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *SetGradeRequest) GetIsTerm() bool {
	if x != nil {
		return x.IsTerm
	}
	return false
}

type SetGradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetGradeResponse) Reset() {
	*x = SetGradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetGradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGradeResponse) ProtoMessage() {}

func (x *SetGradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGradeResponse.ProtoReflect.Descriptor instead.
func (*SetGradeResponse) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{1}
}

func (x *SetGradeResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Grade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId  int64  `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	LessonName string `protobuf:"bytes,2,opt,name=lesson_name,json=lessonName,proto3" json:"lesson_name,omitempty"`
	Date       string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Grade      int64  `protobuf:"varint,4,opt,name=grade,proto3" json:"grade,omitempty"`
}

func (x *Grade) Reset() {
	*x = Grade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grade) ProtoMessage() {}

func (x *Grade) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grade.ProtoReflect.Descriptor instead.
func (*Grade) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{2}
}

func (x *Grade) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *Grade) GetLessonName() string {
	if x != nil {
		return x.LessonName
	}
	return ""
}

func (x *Grade) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Grade) GetGrade() int64 {
	if x != nil {
		return x.Grade
	}
	return 0
}

type GetAllLessonsGradesByStudentIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId int64 `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
}

func (x *GetAllLessonsGradesByStudentIDRequest) Reset() {
	*x = GetAllLessonsGradesByStudentIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllLessonsGradesByStudentIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllLessonsGradesByStudentIDRequest) ProtoMessage() {}

func (x *GetAllLessonsGradesByStudentIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllLessonsGradesByStudentIDRequest.ProtoReflect.Descriptor instead.
func (*GetAllLessonsGradesByStudentIDRequest) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllLessonsGradesByStudentIDRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

type GetAllLessonsGradesByStudentIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grades []*Grade `protobuf:"bytes,2,rep,name=grades,proto3" json:"grades,omitempty"`
}

func (x *GetAllLessonsGradesByStudentIDResponse) Reset() {
	*x = GetAllLessonsGradesByStudentIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllLessonsGradesByStudentIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllLessonsGradesByStudentIDResponse) ProtoMessage() {}

func (x *GetAllLessonsGradesByStudentIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllLessonsGradesByStudentIDResponse.ProtoReflect.Descriptor instead.
func (*GetAllLessonsGradesByStudentIDResponse) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllLessonsGradesByStudentIDResponse) GetGrades() []*Grade {
	if x != nil {
		return x.Grades
	}
	return nil
}

type GetLessonGradesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId  int64  `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	LessonName string `protobuf:"bytes,2,opt,name=lesson_name,json=lessonName,proto3" json:"lesson_name,omitempty"`
}

func (x *GetLessonGradesRequest) Reset() {
	*x = GetLessonGradesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLessonGradesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLessonGradesRequest) ProtoMessage() {}

func (x *GetLessonGradesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLessonGradesRequest.ProtoReflect.Descriptor instead.
func (*GetLessonGradesRequest) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{5}
}

func (x *GetLessonGradesRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *GetLessonGradesRequest) GetLessonName() string {
	if x != nil {
		return x.LessonName
	}
	return ""
}

type GetLessonGradesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grades []*Grade `protobuf:"bytes,2,rep,name=grades,proto3" json:"grades,omitempty"`
}

func (x *GetLessonGradesResponse) Reset() {
	*x = GetLessonGradesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLessonGradesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLessonGradesResponse) ProtoMessage() {}

func (x *GetLessonGradesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLessonGradesResponse.ProtoReflect.Descriptor instead.
func (*GetLessonGradesResponse) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{6}
}

func (x *GetLessonGradesResponse) GetGrades() []*Grade {
	if x != nil {
		return x.Grades
	}
	return nil
}

type SetTermGradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId  int64  `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	LessonName string `protobuf:"bytes,2,opt,name=lesson_name,json=lessonName,proto3" json:"lesson_name,omitempty"`
}

func (x *SetTermGradeRequest) Reset() {
	*x = SetTermGradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTermGradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTermGradeRequest) ProtoMessage() {}

func (x *SetTermGradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTermGradeRequest.ProtoReflect.Descriptor instead.
func (*SetTermGradeRequest) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{7}
}

func (x *SetTermGradeRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *SetTermGradeRequest) GetLessonName() string {
	if x != nil {
		return x.LessonName
	}
	return ""
}

type SetTermGradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetTermGradeResponse) Reset() {
	*x = SetTermGradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTermGradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTermGradeResponse) ProtoMessage() {}

func (x *SetTermGradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTermGradeResponse.ProtoReflect.Descriptor instead.
func (*SetTermGradeResponse) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{8}
}

func (x *SetTermGradeResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DeleteGradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId  int64  `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	LessonName string `protobuf:"bytes,2,opt,name=lesson_name,json=lessonName,proto3" json:"lesson_name,omitempty"`
	Date       string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *DeleteGradeRequest) Reset() {
	*x = DeleteGradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGradeRequest) ProtoMessage() {}

func (x *DeleteGradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGradeRequest.ProtoReflect.Descriptor instead.
func (*DeleteGradeRequest) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteGradeRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *DeleteGradeRequest) GetLessonName() string {
	if x != nil {
		return x.LessonName
	}
	return ""
}

func (x *DeleteGradeRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type DeleteGradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteGradeResponse) Reset() {
	*x = DeleteGradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGradeResponse) ProtoMessage() {}

func (x *DeleteGradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGradeResponse.ProtoReflect.Descriptor instead.
func (*DeleteGradeResponse) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteGradeResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ChangeGradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId  int64  `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	LessonName string `protobuf:"bytes,2,opt,name=lesson_name,json=lessonName,proto3" json:"lesson_name,omitempty"`
	Date       string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Grade      int64  `protobuf:"varint,4,opt,name=grade,proto3" json:"grade,omitempty"`
}

func (x *ChangeGradeRequest) Reset() {
	*x = ChangeGradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeGradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeGradeRequest) ProtoMessage() {}

func (x *ChangeGradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeGradeRequest.ProtoReflect.Descriptor instead.
func (*ChangeGradeRequest) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{11}
}

func (x *ChangeGradeRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *ChangeGradeRequest) GetLessonName() string {
	if x != nil {
		return x.LessonName
	}
	return ""
}

func (x *ChangeGradeRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ChangeGradeRequest) GetGrade() int64 {
	if x != nil {
		return x.Grade
	}
	return 0
}

type ChangeGradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ChangeGradeResponse) Reset() {
	*x = ChangeGradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeGradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeGradeResponse) ProtoMessage() {}

func (x *ChangeGradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grades_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeGradeResponse.ProtoReflect.Descriptor instead.
func (*ChangeGradeResponse) Descriptor() ([]byte, []int) {
	return file_grades_proto_rawDescGZIP(), []int{12}
}

func (x *ChangeGradeResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_grades_proto protoreflect.FileDescriptor

var file_grades_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x2a, 0x0a, 0x10, 0x53, 0x65, 0x74,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x71, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0x46, 0x0a, 0x25, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x42,
	0x79, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x4f, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x73, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x42, 0x79, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x06, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x73, 0x22, 0x58, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x06, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x22, 0x55, 0x0a,
	0x13, 0x53, 0x65, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x68, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x2d,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x7e, 0x0a,
	0x12, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0x2d, 0x0a,
	0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x84, 0x04, 0x0a,
	0x06, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x47, 0x72, 0x61,
	0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x47, 0x72, 0x61,
	0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x81, 0x01,
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x42, 0x79, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x2d, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x42, 0x79, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x42, 0x79, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x17, 0x2e,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e,
	0x53, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x47, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x54,
	0x65, 0x72, 0x6d, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x65, 0x72, 0x6d,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x48, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1a,
	0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grades_proto_rawDescOnce sync.Once
	file_grades_proto_rawDescData = file_grades_proto_rawDesc
)

func file_grades_proto_rawDescGZIP() []byte {
	file_grades_proto_rawDescOnce.Do(func() {
		file_grades_proto_rawDescData = protoimpl.X.CompressGZIP(file_grades_proto_rawDescData)
	})
	return file_grades_proto_rawDescData
}

var file_grades_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_grades_proto_goTypes = []interface{}{
	(*SetGradeRequest)(nil),                        // 0: grades.SetGradeRequest
	(*SetGradeResponse)(nil),                       // 1: grades.SetGradeResponse
	(*Grade)(nil),                                  // 2: grades.Grade
	(*GetAllLessonsGradesByStudentIDRequest)(nil),  // 3: grades.GetAllLessonsGradesByStudentIDRequest
	(*GetAllLessonsGradesByStudentIDResponse)(nil), // 4: grades.GetAllLessonsGradesByStudentIDResponse
	(*GetLessonGradesRequest)(nil),                 // 5: grades.GetLessonGradesRequest
	(*GetLessonGradesResponse)(nil),                // 6: grades.GetLessonGradesResponse
	(*SetTermGradeRequest)(nil),                    // 7: grades.SetTermGradeRequest
	(*SetTermGradeResponse)(nil),                   // 8: grades.SetTermGradeResponse
	(*DeleteGradeRequest)(nil),                     // 9: grades.DeleteGradeRequest
	(*DeleteGradeResponse)(nil),                    // 10: grades.DeleteGradeResponse
	(*ChangeGradeRequest)(nil),                     // 11: grades.ChangeGradeRequest
	(*ChangeGradeResponse)(nil),                    // 12: grades.ChangeGradeResponse
}
var file_grades_proto_depIdxs = []int32{
	2,  // 0: grades.GetAllLessonsGradesByStudentIDResponse.grades:type_name -> grades.Grade
	2,  // 1: grades.GetLessonGradesResponse.grades:type_name -> grades.Grade
	5,  // 2: grades.Grades.GetLessonGrades:input_type -> grades.GetLessonGradesRequest
	3,  // 3: grades.Grades.GetAllLessonsGradesByStudentID:input_type -> grades.GetAllLessonsGradesByStudentIDRequest
	0,  // 4: grades.Grades.SetGrade:input_type -> grades.SetGradeRequest
	7,  // 5: grades.Grades.SetTermGrade:input_type -> grades.SetTermGradeRequest
	9,  // 6: grades.Grades.DeleteGrade:input_type -> grades.DeleteGradeRequest
	11, // 7: grades.Grades.ChangeGrade:input_type -> grades.ChangeGradeRequest
	6,  // 8: grades.Grades.GetLessonGrades:output_type -> grades.GetLessonGradesResponse
	4,  // 9: grades.Grades.GetAllLessonsGradesByStudentID:output_type -> grades.GetAllLessonsGradesByStudentIDResponse
	1,  // 10: grades.Grades.SetGrade:output_type -> grades.SetGradeResponse
	8,  // 11: grades.Grades.SetTermGrade:output_type -> grades.SetTermGradeResponse
	10, // 12: grades.Grades.DeleteGrade:output_type -> grades.DeleteGradeResponse
	12, // 13: grades.Grades.ChangeGrade:output_type -> grades.ChangeGradeResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_grades_proto_init() }
func file_grades_proto_init() {
	if File_grades_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grades_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetGradeRequest); i {
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
		file_grades_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetGradeResponse); i {
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
		file_grades_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grade); i {
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
		file_grades_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllLessonsGradesByStudentIDRequest); i {
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
		file_grades_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllLessonsGradesByStudentIDResponse); i {
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
		file_grades_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLessonGradesRequest); i {
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
		file_grades_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLessonGradesResponse); i {
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
		file_grades_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTermGradeRequest); i {
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
		file_grades_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTermGradeResponse); i {
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
		file_grades_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGradeRequest); i {
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
		file_grades_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGradeResponse); i {
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
		file_grades_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeGradeRequest); i {
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
		file_grades_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeGradeResponse); i {
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
			RawDescriptor: file_grades_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grades_proto_goTypes,
		DependencyIndexes: file_grades_proto_depIdxs,
		MessageInfos:      file_grades_proto_msgTypes,
	}.Build()
	File_grades_proto = out.File
	file_grades_proto_rawDesc = nil
	file_grades_proto_goTypes = nil
	file_grades_proto_depIdxs = nil
}
