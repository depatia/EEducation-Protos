// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: schedule.proto

package schedule

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

type ScheduleDay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonName string `protobuf:"bytes,1,opt,name=lesson_name,json=lessonName,proto3" json:"lesson_name,omitempty"`
	Classroom  int64  `protobuf:"varint,2,opt,name=classroom,proto3" json:"classroom,omitempty"`
	Date       string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ScheduleDay) Reset() {
	*x = ScheduleDay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleDay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleDay) ProtoMessage() {}

func (x *ScheduleDay) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleDay.ProtoReflect.Descriptor instead.
func (*ScheduleDay) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleDay) GetLessonName() string {
	if x != nil {
		return x.LessonName
	}
	return ""
}

func (x *ScheduleDay) GetClassroom() int64 {
	if x != nil {
		return x.Classroom
	}
	return 0
}

func (x *ScheduleDay) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ScheduleDayForStudents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleDay *ScheduleDay `protobuf:"bytes,1,opt,name=schedule_day,json=scheduleDay,proto3" json:"schedule_day,omitempty"`
	Homework    string       `protobuf:"bytes,2,opt,name=homework,proto3" json:"homework,omitempty"`
	Grade       int64        `protobuf:"varint,3,opt,name=grade,proto3" json:"grade,omitempty"`
}

func (x *ScheduleDayForStudents) Reset() {
	*x = ScheduleDayForStudents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleDayForStudents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleDayForStudents) ProtoMessage() {}

func (x *ScheduleDayForStudents) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleDayForStudents.ProtoReflect.Descriptor instead.
func (*ScheduleDayForStudents) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{1}
}

func (x *ScheduleDayForStudents) GetScheduleDay() *ScheduleDay {
	if x != nil {
		return x.ScheduleDay
	}
	return nil
}

func (x *ScheduleDayForStudents) GetHomework() string {
	if x != nil {
		return x.Homework
	}
	return ""
}

func (x *ScheduleDayForStudents) GetGrade() int64 {
	if x != nil {
		return x.Grade
	}
	return 0
}

type ScheduleWeek struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleWeek []*ScheduleDayForStudents `protobuf:"bytes,1,rep,name=schedule_week,json=scheduleWeek,proto3" json:"schedule_week,omitempty"`
}

func (x *ScheduleWeek) Reset() {
	*x = ScheduleWeek{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleWeek) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleWeek) ProtoMessage() {}

func (x *ScheduleWeek) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleWeek.ProtoReflect.Descriptor instead.
func (*ScheduleWeek) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{2}
}

func (x *ScheduleWeek) GetScheduleWeek() []*ScheduleDayForStudents {
	if x != nil {
		return x.ScheduleWeek
	}
	return nil
}

type SetScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassName   string       `protobuf:"bytes,1,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	ScheduleDay *ScheduleDay `protobuf:"bytes,2,opt,name=schedule_day,json=scheduleDay,proto3" json:"schedule_day,omitempty"`
}

func (x *SetScheduleRequest) Reset() {
	*x = SetScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScheduleRequest) ProtoMessage() {}

func (x *SetScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScheduleRequest.ProtoReflect.Descriptor instead.
func (*SetScheduleRequest) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{3}
}

func (x *SetScheduleRequest) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *SetScheduleRequest) GetScheduleDay() *ScheduleDay {
	if x != nil {
		return x.ScheduleDay
	}
	return nil
}

type SetScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetScheduleResponse) Reset() {
	*x = SetScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScheduleResponse) ProtoMessage() {}

func (x *SetScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScheduleResponse.ProtoReflect.Descriptor instead.
func (*SetScheduleResponse) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{4}
}

func (x *SetScheduleResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetWeekScheduleByClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassName string `protobuf:"bytes,1,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
}

func (x *GetWeekScheduleByClassRequest) Reset() {
	*x = GetWeekScheduleByClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeekScheduleByClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeekScheduleByClassRequest) ProtoMessage() {}

func (x *GetWeekScheduleByClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeekScheduleByClassRequest.ProtoReflect.Descriptor instead.
func (*GetWeekScheduleByClassRequest) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{5}
}

func (x *GetWeekScheduleByClassRequest) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

type GetWeekScheduleByClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schedule *ScheduleWeek `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
}

func (x *GetWeekScheduleByClassResponse) Reset() {
	*x = GetWeekScheduleByClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeekScheduleByClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeekScheduleByClassResponse) ProtoMessage() {}

func (x *GetWeekScheduleByClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeekScheduleByClassResponse.ProtoReflect.Descriptor instead.
func (*GetWeekScheduleByClassResponse) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{6}
}

func (x *GetWeekScheduleByClassResponse) GetSchedule() *ScheduleWeek {
	if x != nil {
		return x.Schedule
	}
	return nil
}

type SetHomeworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassName  string `protobuf:"bytes,1,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	Date       string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	LessonName string `protobuf:"bytes,3,opt,name=lesson_name,json=lessonName,proto3" json:"lesson_name,omitempty"`
	Homework   string `protobuf:"bytes,4,opt,name=homework,proto3" json:"homework,omitempty"`
}

func (x *SetHomeworkRequest) Reset() {
	*x = SetHomeworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHomeworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHomeworkRequest) ProtoMessage() {}

func (x *SetHomeworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHomeworkRequest.ProtoReflect.Descriptor instead.
func (*SetHomeworkRequest) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{7}
}

func (x *SetHomeworkRequest) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *SetHomeworkRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *SetHomeworkRequest) GetLessonName() string {
	if x != nil {
		return x.LessonName
	}
	return ""
}

func (x *SetHomeworkRequest) GetHomework() string {
	if x != nil {
		return x.Homework
	}
	return ""
}

type SetHomeworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetHomeworkResponse) Reset() {
	*x = SetHomeworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHomeworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHomeworkResponse) ProtoMessage() {}

func (x *SetHomeworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHomeworkResponse.ProtoReflect.Descriptor instead.
func (*SetHomeworkResponse) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{8}
}

func (x *SetHomeworkResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_schedule_proto protoreflect.FileDescriptor

var file_schedule_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x60, 0x0a, 0x0b, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x61, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x84, 0x01, 0x0a,
	0x16, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x61, 0x79, 0x46, 0x6f, 0x72, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x44, 0x61, 0x79, 0x52, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x61,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x22, 0x55, 0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x57,
	0x65, 0x65, 0x6b, 0x12, 0x45, 0x0a, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x77, 0x65, 0x65, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x61,
	0x79, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0c, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x22, 0x6d, 0x0a, 0x12, 0x53, 0x65,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x38, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x61, 0x79, 0x52, 0x0b, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x61, 0x79, 0x22, 0x2d, 0x0a, 0x13, 0x53, 0x65, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3e, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x57,
	0x65, 0x65, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x42, 0x79, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x57,
	0x65, 0x65, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x42, 0x79, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x57, 0x65, 0x65, 0x6b, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x84,
	0x01, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x2d, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x32, 0x95, 0x02, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x6d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x27, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x65, 0x65, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x42, 0x79, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x0b, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1c, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e,
	0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schedule_proto_rawDescOnce sync.Once
	file_schedule_proto_rawDescData = file_schedule_proto_rawDesc
)

func file_schedule_proto_rawDescGZIP() []byte {
	file_schedule_proto_rawDescOnce.Do(func() {
		file_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_schedule_proto_rawDescData)
	})
	return file_schedule_proto_rawDescData
}

var file_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_schedule_proto_goTypes = []interface{}{
	(*ScheduleDay)(nil),                    // 0: schedule.ScheduleDay
	(*ScheduleDayForStudents)(nil),         // 1: schedule.ScheduleDayForStudents
	(*ScheduleWeek)(nil),                   // 2: schedule.ScheduleWeek
	(*SetScheduleRequest)(nil),             // 3: schedule.SetScheduleRequest
	(*SetScheduleResponse)(nil),            // 4: schedule.SetScheduleResponse
	(*GetWeekScheduleByClassRequest)(nil),  // 5: schedule.GetWeekScheduleByClassRequest
	(*GetWeekScheduleByClassResponse)(nil), // 6: schedule.GetWeekScheduleByClassResponse
	(*SetHomeworkRequest)(nil),             // 7: schedule.SetHomeworkRequest
	(*SetHomeworkResponse)(nil),            // 8: schedule.SetHomeworkResponse
}
var file_schedule_proto_depIdxs = []int32{
	0, // 0: schedule.ScheduleDayForStudents.schedule_day:type_name -> schedule.ScheduleDay
	1, // 1: schedule.ScheduleWeek.schedule_week:type_name -> schedule.ScheduleDayForStudents
	0, // 2: schedule.SetScheduleRequest.schedule_day:type_name -> schedule.ScheduleDay
	2, // 3: schedule.GetWeekScheduleByClassResponse.schedule:type_name -> schedule.ScheduleWeek
	3, // 4: schedule.Schedule.SetSchedule:input_type -> schedule.SetScheduleRequest
	5, // 5: schedule.Schedule.GetWeekScheduleByClass:input_type -> schedule.GetWeekScheduleByClassRequest
	7, // 6: schedule.Schedule.SetHomework:input_type -> schedule.SetHomeworkRequest
	4, // 7: schedule.Schedule.SetSchedule:output_type -> schedule.SetScheduleResponse
	6, // 8: schedule.Schedule.GetWeekScheduleByClass:output_type -> schedule.GetWeekScheduleByClassResponse
	8, // 9: schedule.Schedule.SetHomework:output_type -> schedule.SetHomeworkResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_schedule_proto_init() }
func file_schedule_proto_init() {
	if File_schedule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schedule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleDay); i {
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
		file_schedule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleDayForStudents); i {
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
		file_schedule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleWeek); i {
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
		file_schedule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScheduleRequest); i {
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
		file_schedule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScheduleResponse); i {
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
		file_schedule_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeekScheduleByClassRequest); i {
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
		file_schedule_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeekScheduleByClassResponse); i {
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
		file_schedule_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHomeworkRequest); i {
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
		file_schedule_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHomeworkResponse); i {
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
			RawDescriptor: file_schedule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schedule_proto_goTypes,
		DependencyIndexes: file_schedule_proto_depIdxs,
		MessageInfos:      file_schedule_proto_msgTypes,
	}.Build()
	File_schedule_proto = out.File
	file_schedule_proto_rawDesc = nil
	file_schedule_proto_goTypes = nil
	file_schedule_proto_depIdxs = nil
}