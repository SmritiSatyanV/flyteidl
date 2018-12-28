// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/task_execution.proto

package admin // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A message used to fetch a single task execution entity.
type TaskExecutionGetRequest struct {
	Id                   *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *TaskExecutionGetRequest) Reset()         { *m = TaskExecutionGetRequest{} }
func (m *TaskExecutionGetRequest) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionGetRequest) ProtoMessage()    {}
func (*TaskExecutionGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_execution_e30a76cbb86d71c8, []int{0}
}
func (m *TaskExecutionGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionGetRequest.Unmarshal(m, b)
}
func (m *TaskExecutionGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionGetRequest.Marshal(b, m, deterministic)
}
func (dst *TaskExecutionGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionGetRequest.Merge(dst, src)
}
func (m *TaskExecutionGetRequest) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionGetRequest.Size(m)
}
func (m *TaskExecutionGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionGetRequest proto.InternalMessageInfo

func (m *TaskExecutionGetRequest) GetId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

// Represents a request structure to retrieve a list of task execution entities.
type TaskExecutionListRequest struct {
	// Indicates the number of resources to be returned.
	// +optional
	Limit uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Indicates the offset to fetch the resource from.
	// It is the responsibility of the user to pass add value of limit to offset to fetch the next set of results.
	// +optional
	Offset uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Indicates a list of filters passed as string.
	// More info on constructing filters : <Link>
	// +optional
	Filters string `protobuf:"bytes,3,opt,name=filters,proto3" json:"filters,omitempty"`
	// Sort ordering for returned list.
	// +optional
	SortBy               *Sort    `protobuf:"bytes,4,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecutionListRequest) Reset()         { *m = TaskExecutionListRequest{} }
func (m *TaskExecutionListRequest) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionListRequest) ProtoMessage()    {}
func (*TaskExecutionListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_execution_e30a76cbb86d71c8, []int{1}
}
func (m *TaskExecutionListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionListRequest.Unmarshal(m, b)
}
func (m *TaskExecutionListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionListRequest.Marshal(b, m, deterministic)
}
func (dst *TaskExecutionListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionListRequest.Merge(dst, src)
}
func (m *TaskExecutionListRequest) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionListRequest.Size(m)
}
func (m *TaskExecutionListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionListRequest proto.InternalMessageInfo

func (m *TaskExecutionListRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TaskExecutionListRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *TaskExecutionListRequest) GetFilters() string {
	if m != nil {
		return m.Filters
	}
	return ""
}

func (m *TaskExecutionListRequest) GetSortBy() *Sort {
	if m != nil {
		return m.SortBy
	}
	return nil
}

// Encapsulates all details for a single task execution entity.
type TaskExecution struct {
	Id                   *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InputUri             string                        `protobuf:"bytes,2,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	Closure              *TaskExecutionClosure         `protobuf:"bytes,3,opt,name=closure,proto3" json:"closure,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *TaskExecution) Reset()         { *m = TaskExecution{} }
func (m *TaskExecution) String() string { return proto.CompactTextString(m) }
func (*TaskExecution) ProtoMessage()    {}
func (*TaskExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_execution_e30a76cbb86d71c8, []int{2}
}
func (m *TaskExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecution.Unmarshal(m, b)
}
func (m *TaskExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecution.Marshal(b, m, deterministic)
}
func (dst *TaskExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecution.Merge(dst, src)
}
func (m *TaskExecution) XXX_Size() int {
	return xxx_messageInfo_TaskExecution.Size(m)
}
func (m *TaskExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecution.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecution proto.InternalMessageInfo

func (m *TaskExecution) GetId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TaskExecution) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

func (m *TaskExecution) GetClosure() *TaskExecutionClosure {
	if m != nil {
		return m.Closure
	}
	return nil
}

// Request structure to retrieve a list of task execution entities.
type TaskExecutionList struct {
	TaskExecutions       []*TaskExecution `protobuf:"bytes,1,rep,name=task_executions,json=taskExecutions,proto3" json:"task_executions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TaskExecutionList) Reset()         { *m = TaskExecutionList{} }
func (m *TaskExecutionList) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionList) ProtoMessage()    {}
func (*TaskExecutionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_execution_e30a76cbb86d71c8, []int{3}
}
func (m *TaskExecutionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionList.Unmarshal(m, b)
}
func (m *TaskExecutionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionList.Marshal(b, m, deterministic)
}
func (dst *TaskExecutionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionList.Merge(dst, src)
}
func (m *TaskExecutionList) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionList.Size(m)
}
func (m *TaskExecutionList) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionList.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionList proto.InternalMessageInfo

func (m *TaskExecutionList) GetTaskExecutions() []*TaskExecution {
	if m != nil {
		return m.TaskExecutions
	}
	return nil
}

// Container for task execution details and results.
type TaskExecutionClosure struct {
	// Types that are valid to be assigned to OutputResult:
	//	*TaskExecutionClosure_OutputUri
	//	*TaskExecutionClosure_Error
	OutputResult isTaskExecutionClosure_OutputResult `protobuf_oneof:"output_result"`
	Phase        core.TaskExecutionPhase             `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.TaskExecutionPhase" json:"phase,omitempty"`
	Logs         []*core.TaskLog                     `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty"`
	// Time at which the task execution began running.
	StartedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// The amount of time the task execution spent running.
	Duration *duration.Duration `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	// Time at which the task execution was created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Time at which the task execution was last updated.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskExecutionClosure) Reset()         { *m = TaskExecutionClosure{} }
func (m *TaskExecutionClosure) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionClosure) ProtoMessage()    {}
func (*TaskExecutionClosure) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_execution_e30a76cbb86d71c8, []int{4}
}
func (m *TaskExecutionClosure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionClosure.Unmarshal(m, b)
}
func (m *TaskExecutionClosure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionClosure.Marshal(b, m, deterministic)
}
func (dst *TaskExecutionClosure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionClosure.Merge(dst, src)
}
func (m *TaskExecutionClosure) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionClosure.Size(m)
}
func (m *TaskExecutionClosure) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionClosure.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionClosure proto.InternalMessageInfo

type isTaskExecutionClosure_OutputResult interface {
	isTaskExecutionClosure_OutputResult()
}

type TaskExecutionClosure_OutputUri struct {
	OutputUri string `protobuf:"bytes,1,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type TaskExecutionClosure_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*TaskExecutionClosure_OutputUri) isTaskExecutionClosure_OutputResult() {}

func (*TaskExecutionClosure_Error) isTaskExecutionClosure_OutputResult() {}

func (m *TaskExecutionClosure) GetOutputResult() isTaskExecutionClosure_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *TaskExecutionClosure) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*TaskExecutionClosure_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *TaskExecutionClosure) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*TaskExecutionClosure_Error); ok {
		return x.Error
	}
	return nil
}

func (m *TaskExecutionClosure) GetPhase() core.TaskExecutionPhase {
	if m != nil {
		return m.Phase
	}
	return core.TaskExecutionPhase_TASK_PHASE_UNDEFINED
}

func (m *TaskExecutionClosure) GetLogs() []*core.TaskLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *TaskExecutionClosure) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *TaskExecutionClosure) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *TaskExecutionClosure) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TaskExecutionClosure) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TaskExecutionClosure) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TaskExecutionClosure_OneofMarshaler, _TaskExecutionClosure_OneofUnmarshaler, _TaskExecutionClosure_OneofSizer, []interface{}{
		(*TaskExecutionClosure_OutputUri)(nil),
		(*TaskExecutionClosure_Error)(nil),
	}
}

func _TaskExecutionClosure_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TaskExecutionClosure)
	// output_result
	switch x := m.OutputResult.(type) {
	case *TaskExecutionClosure_OutputUri:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OutputUri)
	case *TaskExecutionClosure_Error:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TaskExecutionClosure.OutputResult has unexpected type %T", x)
	}
	return nil
}

func _TaskExecutionClosure_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TaskExecutionClosure)
	switch tag {
	case 1: // output_result.output_uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OutputResult = &TaskExecutionClosure_OutputUri{x}
		return true, err
	case 2: // output_result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.ExecutionError)
		err := b.DecodeMessage(msg)
		m.OutputResult = &TaskExecutionClosure_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TaskExecutionClosure_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TaskExecutionClosure)
	// output_result
	switch x := m.OutputResult.(type) {
	case *TaskExecutionClosure_OutputUri:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.OutputUri)))
		n += len(x.OutputUri)
	case *TaskExecutionClosure_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*TaskExecutionGetRequest)(nil), "flyteidl.admin.TaskExecutionGetRequest")
	proto.RegisterType((*TaskExecutionListRequest)(nil), "flyteidl.admin.TaskExecutionListRequest")
	proto.RegisterType((*TaskExecution)(nil), "flyteidl.admin.TaskExecution")
	proto.RegisterType((*TaskExecutionList)(nil), "flyteidl.admin.TaskExecutionList")
	proto.RegisterType((*TaskExecutionClosure)(nil), "flyteidl.admin.TaskExecutionClosure")
}

func init() {
	proto.RegisterFile("flyteidl/admin/task_execution.proto", fileDescriptor_task_execution_e30a76cbb86d71c8)
}

var fileDescriptor_task_execution_e30a76cbb86d71c8 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xeb, 0x36, 0x5f, 0x9e, 0x28, 0xa9, 0x58, 0x45, 0xc5, 0xa4, 0x2a, 0x0d, 0x01, 0xa1,
	0x08, 0xa9, 0xb6, 0x48, 0x55, 0x10, 0x17, 0xa4, 0x04, 0x0a, 0x45, 0xea, 0x01, 0x96, 0x72, 0x81,
	0x43, 0xe4, 0xd8, 0x6b, 0x77, 0x55, 0x3b, 0x6b, 0x76, 0xc7, 0x12, 0x79, 0x0c, 0xde, 0x81, 0xc7,
	0xe1, 0xa1, 0x90, 0xd7, 0x1f, 0xe0, 0xb4, 0x6a, 0x25, 0x8e, 0xb3, 0xf3, 0xfb, 0xcf, 0xfc, 0x67,
	0xbc, 0x5e, 0x78, 0x1c, 0x44, 0x6b, 0x64, 0xdc, 0x8f, 0x1c, 0xd7, 0x8f, 0xf9, 0xca, 0x41, 0x57,
	0x5d, 0x2d, 0xd8, 0x0f, 0xe6, 0xa5, 0xc8, 0xc5, 0xca, 0x4e, 0xa4, 0x40, 0x41, 0xfa, 0x25, 0x64,
	0x6b, 0x68, 0xb8, 0xbf, 0x21, 0xf2, 0x44, 0x1c, 0x97, 0xf0, 0xf0, 0xa0, 0x4a, 0x7a, 0x42, 0x32,
	0x67, 0xa3, 0xd6, 0xf0, 0x61, 0x3d, 0xcd, 0x7d, 0xb6, 0x42, 0x1e, 0x70, 0x26, 0x8b, 0xfc, 0x61,
	0x28, 0x44, 0x18, 0x31, 0x47, 0x47, 0xcb, 0x34, 0x70, 0x90, 0xc7, 0x4c, 0xa1, 0x1b, 0x27, 0x65,
	0x81, 0x4d, 0xc0, 0x4f, 0xa5, 0xfb, 0xb7, 0xc1, 0xf8, 0x13, 0xdc, 0xbf, 0x70, 0xd5, 0xd5, 0x69,
	0xd9, 0xf7, 0x3d, 0x43, 0xca, 0xbe, 0xa7, 0x4c, 0x21, 0x79, 0x01, 0xdb, 0xdc, 0xb7, 0x8c, 0x91,
	0x31, 0xe9, 0x4e, 0x9f, 0xda, 0xd5, 0x50, 0x99, 0x11, 0xbb, 0xa6, 0xf9, 0x50, 0xb9, 0xa2, 0xdb,
	0xdc, 0x1f, 0xff, 0x34, 0xc0, 0xaa, 0xe5, 0xcf, 0xb9, 0xaa, 0x8a, 0x0e, 0xa0, 0x19, 0xf1, 0x98,
	0xa3, 0xae, 0xdb, 0xa3, 0x79, 0x40, 0xf6, 0xa0, 0x25, 0x82, 0x40, 0x31, 0xb4, 0xb6, 0xf5, 0x71,
	0x11, 0x11, 0x0b, 0xda, 0x01, 0x8f, 0x90, 0x49, 0x65, 0xed, 0x8c, 0x8c, 0x89, 0x49, 0xcb, 0x90,
	0x1c, 0x41, 0x5b, 0x09, 0x89, 0x8b, 0xe5, 0xda, 0x6a, 0x68, 0x87, 0x03, 0xbb, 0xbe, 0x76, 0xfb,
	0xb3, 0x90, 0x48, 0x5b, 0x19, 0x34, 0x5f, 0x8f, 0x7f, 0x19, 0xd0, 0xab, 0x79, 0xfa, 0xdf, 0xe9,
	0xc8, 0x3e, 0x98, 0x7c, 0x95, 0xa4, 0xb8, 0x48, 0x25, 0xd7, 0x6e, 0x4d, 0xda, 0xd1, 0x07, 0x5f,
	0x24, 0x27, 0xaf, 0xa1, 0xed, 0x45, 0x42, 0xa5, 0x92, 0x69, 0xbf, 0xdd, 0xe9, 0x93, 0x4d, 0x57,
	0xb5, 0xd2, 0x6f, 0x72, 0x96, 0x96, 0xa2, 0xf1, 0x37, 0xb8, 0x77, 0x6d, 0x73, 0xe4, 0x1d, 0xec,
	0xd6, 0xef, 0x99, 0xb2, 0x8c, 0xd1, 0xce, 0xa4, 0x3b, 0x3d, 0xb8, 0xb5, 0x38, 0xed, 0xe3, 0xbf,
	0xa1, 0x1a, 0xff, 0xde, 0x81, 0xc1, 0x4d, 0xed, 0xc9, 0x21, 0x80, 0x48, 0xb1, 0x9c, 0x29, 0x5b,
	0x89, 0x79, 0xb6, 0x45, 0xcd, 0xfc, 0x2c, 0x1b, 0xeb, 0x04, 0x9a, 0x4c, 0x4a, 0x21, 0xf5, 0xbc,
	0xb5, 0xbe, 0x7a, 0x5d, 0x55, 0xc1, 0xd3, 0x0c, 0x3a, 0xdb, 0xa2, 0x39, 0x4d, 0x5e, 0x42, 0x33,
	0xb9, 0x74, 0x55, 0xbe, 0x8b, 0xfe, 0xf4, 0xd1, 0x6d, 0x5b, 0xfe, 0x98, 0x81, 0x34, 0xe7, 0xc9,
	0x33, 0x68, 0x44, 0x22, 0x54, 0x56, 0x43, 0x8f, 0xb9, 0x77, 0x83, 0xee, 0x5c, 0x84, 0x54, 0x33,
	0xe4, 0x15, 0x80, 0x42, 0x57, 0x22, 0xf3, 0x17, 0x2e, 0x5a, 0x4d, 0x6d, 0x70, 0x68, 0xe7, 0xb7,
	0xde, 0x2e, 0x6f, 0xbd, 0x7d, 0x51, 0xfe, 0x16, 0xd4, 0x2c, 0xe8, 0x19, 0x92, 0x13, 0xe8, 0x94,
	0x7f, 0x83, 0xd5, 0xd2, 0xc2, 0x07, 0xd7, 0x84, 0x6f, 0x0b, 0x80, 0x56, 0x68, 0xd6, 0xd1, 0x93,
	0xcc, 0x2d, 0x3a, 0xb6, 0xef, 0xee, 0x58, 0xd0, 0x33, 0xcc, 0xa4, 0x69, 0xe2, 0x97, 0xd2, 0xce,
	0xdd, 0xd2, 0x82, 0x9e, 0xe1, 0x7c, 0x17, 0x7a, 0xc5, 0x47, 0x92, 0x4c, 0xa5, 0x11, 0xce, 0x8f,
	0xbf, 0x3e, 0x0f, 0x39, 0x5e, 0xa6, 0x4b, 0xdb, 0x13, 0xb1, 0x13, 0xad, 0x03, 0x74, 0xaa, 0xc7,
	0x22, 0x64, 0x2b, 0x27, 0x59, 0x1e, 0x85, 0xc2, 0xa9, 0xbf, 0x3d, 0xcb, 0x96, 0x6e, 0x72, 0xfc,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x95, 0x10, 0x76, 0xb5, 0xc9, 0x04, 0x00, 0x00,
}
