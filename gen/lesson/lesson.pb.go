// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: lesson.proto

package lesson

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

type TeacherLesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonName string `protobuf:"bytes,1,opt,name=lesson_name,json=lessonName,proto3" json:"lesson_name,omitempty"`
	ClassName  string `protobuf:"bytes,2,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
}

func (x *TeacherLesson) Reset() {
	*x = TeacherLesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherLesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherLesson) ProtoMessage() {}

func (x *TeacherLesson) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherLesson.ProtoReflect.Descriptor instead.
func (*TeacherLesson) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{0}
}

func (x *TeacherLesson) GetLessonName() string {
	if x != nil {
		return x.LessonName
	}
	return ""
}

func (x *TeacherLesson) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

type SetLessonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherId  int64  `protobuf:"varint,1,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
	LessonName string `protobuf:"bytes,2,opt,name=lesson_name,json=lessonName,proto3" json:"lesson_name,omitempty"`
	ClassName  string `protobuf:"bytes,3,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
}

func (x *SetLessonRequest) Reset() {
	*x = SetLessonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLessonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLessonRequest) ProtoMessage() {}

func (x *SetLessonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLessonRequest.ProtoReflect.Descriptor instead.
func (*SetLessonRequest) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{1}
}

func (x *SetLessonRequest) GetTeacherId() int64 {
	if x != nil {
		return x.TeacherId
	}
	return 0
}

func (x *SetLessonRequest) GetLessonName() string {
	if x != nil {
		return x.LessonName
	}
	return ""
}

func (x *SetLessonRequest) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

type SetLessonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetLessonResponse) Reset() {
	*x = SetLessonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLessonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLessonResponse) ProtoMessage() {}

func (x *SetLessonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLessonResponse.ProtoReflect.Descriptor instead.
func (*SetLessonResponse) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{2}
}

func (x *SetLessonResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetAllTeacherLessonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherId int64 `protobuf:"varint,1,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
}

func (x *GetAllTeacherLessonsRequest) Reset() {
	*x = GetAllTeacherLessonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTeacherLessonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTeacherLessonsRequest) ProtoMessage() {}

func (x *GetAllTeacherLessonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTeacherLessonsRequest.ProtoReflect.Descriptor instead.
func (*GetAllTeacherLessonsRequest) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllTeacherLessonsRequest) GetTeacherId() int64 {
	if x != nil {
		return x.TeacherId
	}
	return 0
}

type GetAllTeacherLessonsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lessons []*TeacherLesson `protobuf:"bytes,1,rep,name=lessons,proto3" json:"lessons,omitempty"`
}

func (x *GetAllTeacherLessonsResponse) Reset() {
	*x = GetAllTeacherLessonsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTeacherLessonsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTeacherLessonsResponse) ProtoMessage() {}

func (x *GetAllTeacherLessonsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTeacherLessonsResponse.ProtoReflect.Descriptor instead.
func (*GetAllTeacherLessonsResponse) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllTeacherLessonsResponse) GetLessons() []*TeacherLesson {
	if x != nil {
		return x.Lessons
	}
	return nil
}

type IsLessonExistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonName string `protobuf:"bytes,1,opt,name=lesson_name,json=lessonName,proto3" json:"lesson_name,omitempty"`
	ClassName  string `protobuf:"bytes,2,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
}

func (x *IsLessonExistsRequest) Reset() {
	*x = IsLessonExistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLessonExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLessonExistsRequest) ProtoMessage() {}

func (x *IsLessonExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLessonExistsRequest.ProtoReflect.Descriptor instead.
func (*IsLessonExistsRequest) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{5}
}

func (x *IsLessonExistsRequest) GetLessonName() string {
	if x != nil {
		return x.LessonName
	}
	return ""
}

func (x *IsLessonExistsRequest) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

type IsLessonExistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *IsLessonExistsResponse) Reset() {
	*x = IsLessonExistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLessonExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLessonExistsResponse) ProtoMessage() {}

func (x *IsLessonExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLessonExistsResponse.ProtoReflect.Descriptor instead.
func (*IsLessonExistsResponse) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{6}
}

func (x *IsLessonExistsResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_lesson_proto protoreflect.FileDescriptor

var file_lesson_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x0d, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x71, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x4c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x11, 0x53, 0x65,
	0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3c, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x6c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x57, 0x0a, 0x15, 0x49, 0x73, 0x4c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x30, 0x0a, 0x16, 0x49, 0x73, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x32, 0x84, 0x02, 0x0a, 0x06, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x09,
	0x53, 0x65, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74,
	0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x63, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x4c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e, 0x49, 0x73, 0x4c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x2e, 0x49, 0x73, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e,
	0x49, 0x73, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_lesson_proto_rawDescOnce sync.Once
	file_lesson_proto_rawDescData = file_lesson_proto_rawDesc
)

func file_lesson_proto_rawDescGZIP() []byte {
	file_lesson_proto_rawDescOnce.Do(func() {
		file_lesson_proto_rawDescData = protoimpl.X.CompressGZIP(file_lesson_proto_rawDescData)
	})
	return file_lesson_proto_rawDescData
}

var file_lesson_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_lesson_proto_goTypes = []interface{}{
	(*TeacherLesson)(nil),                // 0: lesson.TeacherLesson
	(*SetLessonRequest)(nil),             // 1: lesson.SetLessonRequest
	(*SetLessonResponse)(nil),            // 2: lesson.SetLessonResponse
	(*GetAllTeacherLessonsRequest)(nil),  // 3: lesson.GetAllTeacherLessonsRequest
	(*GetAllTeacherLessonsResponse)(nil), // 4: lesson.GetAllTeacherLessonsResponse
	(*IsLessonExistsRequest)(nil),        // 5: lesson.IsLessonExistsRequest
	(*IsLessonExistsResponse)(nil),       // 6: lesson.IsLessonExistsResponse
}
var file_lesson_proto_depIdxs = []int32{
	0, // 0: lesson.GetAllTeacherLessonsResponse.lessons:type_name -> lesson.TeacherLesson
	1, // 1: lesson.Lesson.SetLesson:input_type -> lesson.SetLessonRequest
	3, // 2: lesson.Lesson.GetAllTeacherLessons:input_type -> lesson.GetAllTeacherLessonsRequest
	5, // 3: lesson.Lesson.IsLessonExists:input_type -> lesson.IsLessonExistsRequest
	2, // 4: lesson.Lesson.SetLesson:output_type -> lesson.SetLessonResponse
	4, // 5: lesson.Lesson.GetAllTeacherLessons:output_type -> lesson.GetAllTeacherLessonsResponse
	6, // 6: lesson.Lesson.IsLessonExists:output_type -> lesson.IsLessonExistsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_lesson_proto_init() }
func file_lesson_proto_init() {
	if File_lesson_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lesson_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeacherLesson); i {
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
		file_lesson_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLessonRequest); i {
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
		file_lesson_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLessonResponse); i {
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
		file_lesson_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTeacherLessonsRequest); i {
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
		file_lesson_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTeacherLessonsResponse); i {
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
		file_lesson_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLessonExistsRequest); i {
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
		file_lesson_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLessonExistsResponse); i {
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
			RawDescriptor: file_lesson_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lesson_proto_goTypes,
		DependencyIndexes: file_lesson_proto_depIdxs,
		MessageInfos:      file_lesson_proto_msgTypes,
	}.Build()
	File_lesson_proto = out.File
	file_lesson_proto_rawDesc = nil
	file_lesson_proto_goTypes = nil
	file_lesson_proto_depIdxs = nil
}