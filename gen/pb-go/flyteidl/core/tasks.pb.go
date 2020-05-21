// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/tasks.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Known resource names.
type Resources_ResourceName int32

const (
	Resources_UNKNOWN Resources_ResourceName = 0
	Resources_CPU     Resources_ResourceName = 1
	Resources_GPU     Resources_ResourceName = 2
	Resources_MEMORY  Resources_ResourceName = 3
	Resources_STORAGE Resources_ResourceName = 4
)

var Resources_ResourceName_name = map[int32]string{
	0: "UNKNOWN",
	1: "CPU",
	2: "GPU",
	3: "MEMORY",
	4: "STORAGE",
}

var Resources_ResourceName_value = map[string]int32{
	"UNKNOWN": 0,
	"CPU":     1,
	"GPU":     2,
	"MEMORY":  3,
	"STORAGE": 4,
}

func (x Resources_ResourceName) String() string {
	return proto.EnumName(Resources_ResourceName_name, int32(x))
}

func (Resources_ResourceName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{0, 0}
}

type RuntimeMetadata_RuntimeType int32

const (
	RuntimeMetadata_OTHER     RuntimeMetadata_RuntimeType = 0
	RuntimeMetadata_FLYTE_SDK RuntimeMetadata_RuntimeType = 1
)

var RuntimeMetadata_RuntimeType_name = map[int32]string{
	0: "OTHER",
	1: "FLYTE_SDK",
}

var RuntimeMetadata_RuntimeType_value = map[string]int32{
	"OTHER":     0,
	"FLYTE_SDK": 1,
}

func (x RuntimeMetadata_RuntimeType) String() string {
	return proto.EnumName(RuntimeMetadata_RuntimeType_name, int32(x))
}

func (RuntimeMetadata_RuntimeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{1, 0}
}

// MetadataFormat decides the encoding format in which the input metadata should be made available to the containers.
// If the user has access to the protocol buffer definitions, it is recommended to use the PROTO format.
// JSON and YAML do not need any protobuf definitions to read it
// All remote references in core.LiteralMap are replaced with local filesystem references (the data is downloaded to local filesystem)
type DataLoadingConfig_MetadataFormat int32

const (
	// JSON / YAML for the metadata (which contains inlined primitive values). The representation is inline with the standard json specification as specified - https://www.json.org/json-en.html
	DataLoadingConfig_JSON DataLoadingConfig_MetadataFormat = 0
	DataLoadingConfig_YAML DataLoadingConfig_MetadataFormat = 1
	// Proto is a serialized binary of `core.LiteralMap` defined in flyteidl/core
	DataLoadingConfig_PROTO DataLoadingConfig_MetadataFormat = 2
)

var DataLoadingConfig_MetadataFormat_name = map[int32]string{
	0: "JSON",
	1: "YAML",
	2: "PROTO",
}

var DataLoadingConfig_MetadataFormat_value = map[string]int32{
	"JSON":  0,
	"YAML":  1,
	"PROTO": 2,
}

func (x DataLoadingConfig_MetadataFormat) String() string {
	return proto.EnumName(DataLoadingConfig_MetadataFormat_name, int32(x))
}

func (DataLoadingConfig_MetadataFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{6, 0}
}

// Strategy to use when dealing with Blob, Schema, or multipart blob data (large datasets)
type DataLoadingConfig_BlobDownload int32

const (
	// All data will be downloaded before the main container is executed
	DataLoadingConfig_BEFORE_STARTUP DataLoadingConfig_BlobDownload = 0
	// Data will be downloaded as a stream and an End-Of-Stream marker will be written to indicate all data has been downloaded. Refer to protocol for details
	DataLoadingConfig_STREAM DataLoadingConfig_BlobDownload = 1
	// Large objects (offloaded) will not be downloaded
	DataLoadingConfig_DO_NOT_DOWNLOAD DataLoadingConfig_BlobDownload = 2
)

var DataLoadingConfig_BlobDownload_name = map[int32]string{
	0: "BEFORE_STARTUP",
	1: "STREAM",
	2: "DO_NOT_DOWNLOAD",
}

var DataLoadingConfig_BlobDownload_value = map[string]int32{
	"BEFORE_STARTUP":  0,
	"STREAM":          1,
	"DO_NOT_DOWNLOAD": 2,
}

func (x DataLoadingConfig_BlobDownload) String() string {
	return proto.EnumName(DataLoadingConfig_BlobDownload_name, int32(x))
}

func (DataLoadingConfig_BlobDownload) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{6, 1}
}

type DataLoadingConfig_BlobUpload int32

const (
	// All data will be uploaded after the main container exits
	DataLoadingConfig_ON_EXIT DataLoadingConfig_BlobUpload = 0
	// Data will be uploaded as it appears. Refer to protocol specification for details
	DataLoadingConfig_WHEN_AVAILABLE DataLoadingConfig_BlobUpload = 1
	// Data will not be uploaded, only references will be written
	DataLoadingConfig_DO_NOT_UPLOAD DataLoadingConfig_BlobUpload = 2
)

var DataLoadingConfig_BlobUpload_name = map[int32]string{
	0: "ON_EXIT",
	1: "WHEN_AVAILABLE",
	2: "DO_NOT_UPLOAD",
}

var DataLoadingConfig_BlobUpload_value = map[string]int32{
	"ON_EXIT":        0,
	"WHEN_AVAILABLE": 1,
	"DO_NOT_UPLOAD":  2,
}

func (x DataLoadingConfig_BlobUpload) String() string {
	return proto.EnumName(DataLoadingConfig_BlobUpload_name, int32(x))
}

func (DataLoadingConfig_BlobUpload) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{6, 2}
}

// A customizable interface to convey resources requested for a container. This can be interpretted differently for different
// container engines.
type Resources struct {
	// The desired set of resources requested. ResourceNames must be unique within the list.
	Requests []*Resources_ResourceEntry `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	// Defines a set of bounds (e.g. min/max) within which the task can reliably run. ResourceNames must be unique
	// within the list.
	Limits               []*Resources_ResourceEntry `protobuf:"bytes,2,rep,name=limits,proto3" json:"limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{0}
}

func (m *Resources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources.Unmarshal(m, b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
}
func (m *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(m, src)
}
func (m *Resources) XXX_Size() int {
	return xxx_messageInfo_Resources.Size(m)
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func (m *Resources) GetRequests() []*Resources_ResourceEntry {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *Resources) GetLimits() []*Resources_ResourceEntry {
	if m != nil {
		return m.Limits
	}
	return nil
}

// Encapsulates a resource name and value.
type Resources_ResourceEntry struct {
	// Resource name.
	Name Resources_ResourceName `protobuf:"varint,1,opt,name=name,proto3,enum=flyteidl.core.Resources_ResourceName" json:"name,omitempty"`
	// Value must be a valid k8s quantity. See
	// https://github.com/kubernetes/apimachinery/blob/master/pkg/api/resource/quantity.go#L30-L80
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resources_ResourceEntry) Reset()         { *m = Resources_ResourceEntry{} }
func (m *Resources_ResourceEntry) String() string { return proto.CompactTextString(m) }
func (*Resources_ResourceEntry) ProtoMessage()    {}
func (*Resources_ResourceEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{0, 0}
}

func (m *Resources_ResourceEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources_ResourceEntry.Unmarshal(m, b)
}
func (m *Resources_ResourceEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources_ResourceEntry.Marshal(b, m, deterministic)
}
func (m *Resources_ResourceEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources_ResourceEntry.Merge(m, src)
}
func (m *Resources_ResourceEntry) XXX_Size() int {
	return xxx_messageInfo_Resources_ResourceEntry.Size(m)
}
func (m *Resources_ResourceEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources_ResourceEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Resources_ResourceEntry proto.InternalMessageInfo

func (m *Resources_ResourceEntry) GetName() Resources_ResourceName {
	if m != nil {
		return m.Name
	}
	return Resources_UNKNOWN
}

func (m *Resources_ResourceEntry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Runtime information. This is losely defined to allow for extensibility.
type RuntimeMetadata struct {
	// Type of runtime.
	Type RuntimeMetadata_RuntimeType `protobuf:"varint,1,opt,name=type,proto3,enum=flyteidl.core.RuntimeMetadata_RuntimeType" json:"type,omitempty"`
	// Version of the runtime. All versions should be backward compatible. However, certain cases call for version
	// checks to ensure tighter validation or setting expectations.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	//+optional It can be used to provide extra information about the runtime (e.g. python, golang... etc.).
	Flavor               string   `protobuf:"bytes,3,opt,name=flavor,proto3" json:"flavor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuntimeMetadata) Reset()         { *m = RuntimeMetadata{} }
func (m *RuntimeMetadata) String() string { return proto.CompactTextString(m) }
func (*RuntimeMetadata) ProtoMessage()    {}
func (*RuntimeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{1}
}

func (m *RuntimeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeMetadata.Unmarshal(m, b)
}
func (m *RuntimeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeMetadata.Marshal(b, m, deterministic)
}
func (m *RuntimeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeMetadata.Merge(m, src)
}
func (m *RuntimeMetadata) XXX_Size() int {
	return xxx_messageInfo_RuntimeMetadata.Size(m)
}
func (m *RuntimeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeMetadata proto.InternalMessageInfo

func (m *RuntimeMetadata) GetType() RuntimeMetadata_RuntimeType {
	if m != nil {
		return m.Type
	}
	return RuntimeMetadata_OTHER
}

func (m *RuntimeMetadata) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RuntimeMetadata) GetFlavor() string {
	if m != nil {
		return m.Flavor
	}
	return ""
}

// Task Metadata
type TaskMetadata struct {
	// Indicates whether the system should attempt to lookup this task's output to avoid duplication of work.
	Discoverable bool `protobuf:"varint,1,opt,name=discoverable,proto3" json:"discoverable,omitempty"`
	// Runtime information about the task.
	Runtime *RuntimeMetadata `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// The overall timeout of a task including user-triggered retries.
	Timeout *duration.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Number of retries per task.
	Retries *RetryStrategy `protobuf:"bytes,5,opt,name=retries,proto3" json:"retries,omitempty"`
	// Indicates a logical version to apply to this task for the purpose of discovery.
	DiscoveryVersion string `protobuf:"bytes,6,opt,name=discovery_version,json=discoveryVersion,proto3" json:"discovery_version,omitempty"`
	// If set, this indicates that this task is deprecated.  This will enable owners of tasks to notify consumers
	// of the ending of support for a given task.
	DeprecatedErrorMessage string `protobuf:"bytes,7,opt,name=deprecated_error_message,json=deprecatedErrorMessage,proto3" json:"deprecated_error_message,omitempty"`
	// Identify whether task is interruptible
	//
	// Types that are valid to be assigned to InterruptibleValue:
	//	*TaskMetadata_Interruptible
	InterruptibleValue   isTaskMetadata_InterruptibleValue `protobuf_oneof:"interruptible_value"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TaskMetadata) Reset()         { *m = TaskMetadata{} }
func (m *TaskMetadata) String() string { return proto.CompactTextString(m) }
func (*TaskMetadata) ProtoMessage()    {}
func (*TaskMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{2}
}

func (m *TaskMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskMetadata.Unmarshal(m, b)
}
func (m *TaskMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskMetadata.Marshal(b, m, deterministic)
}
func (m *TaskMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskMetadata.Merge(m, src)
}
func (m *TaskMetadata) XXX_Size() int {
	return xxx_messageInfo_TaskMetadata.Size(m)
}
func (m *TaskMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaskMetadata proto.InternalMessageInfo

func (m *TaskMetadata) GetDiscoverable() bool {
	if m != nil {
		return m.Discoverable
	}
	return false
}

func (m *TaskMetadata) GetRuntime() *RuntimeMetadata {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *TaskMetadata) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *TaskMetadata) GetRetries() *RetryStrategy {
	if m != nil {
		return m.Retries
	}
	return nil
}

func (m *TaskMetadata) GetDiscoveryVersion() string {
	if m != nil {
		return m.DiscoveryVersion
	}
	return ""
}

func (m *TaskMetadata) GetDeprecatedErrorMessage() string {
	if m != nil {
		return m.DeprecatedErrorMessage
	}
	return ""
}

type isTaskMetadata_InterruptibleValue interface {
	isTaskMetadata_InterruptibleValue()
}

type TaskMetadata_Interruptible struct {
	Interruptible bool `protobuf:"varint,8,opt,name=interruptible,proto3,oneof"`
}

func (*TaskMetadata_Interruptible) isTaskMetadata_InterruptibleValue() {}

func (m *TaskMetadata) GetInterruptibleValue() isTaskMetadata_InterruptibleValue {
	if m != nil {
		return m.InterruptibleValue
	}
	return nil
}

func (m *TaskMetadata) GetInterruptible() bool {
	if x, ok := m.GetInterruptibleValue().(*TaskMetadata_Interruptible); ok {
		return x.Interruptible
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TaskMetadata) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TaskMetadata_Interruptible)(nil),
	}
}

// A Task structure that uniquely identifies a task in the system
// Tasks are registered as a first step in the system.
type TaskTemplate struct {
	// Auto generated taskId by the system. Task Id uniquely identifies this task globally.
	Id *Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A predefined yet extensible Task type identifier. This can be used to customize any of the components. If no
	// extensions are provided in the system, Flyte will resolve the this task to its TaskCategory and default the
	// implementation registered for the TaskCategory.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Extra metadata about the task.
	Metadata *TaskMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// A strongly typed interface for the task. This enables others to use this task within a workflow and gauarantees
	// compile-time validation of the workflow to avoid costly runtime failures.
	Interface *TypedInterface `protobuf:"bytes,4,opt,name=interface,proto3" json:"interface,omitempty"`
	// Custom data about the task. This is extensible to allow various plugins in the system.
	Custom *_struct.Struct `protobuf:"bytes,5,opt,name=custom,proto3" json:"custom,omitempty"`
	// Known target types that the system will guarantee plugins for. Custom SDK plugins are allowed to set these if needed.
	// If no corresponding execution-layer plugins are found, the system will default to handling these using built-in
	// handlers.
	//
	// Types that are valid to be assigned to Target:
	//	*TaskTemplate_Container
	Target               isTaskTemplate_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TaskTemplate) Reset()         { *m = TaskTemplate{} }
func (m *TaskTemplate) String() string { return proto.CompactTextString(m) }
func (*TaskTemplate) ProtoMessage()    {}
func (*TaskTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{3}
}

func (m *TaskTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskTemplate.Unmarshal(m, b)
}
func (m *TaskTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskTemplate.Marshal(b, m, deterministic)
}
func (m *TaskTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskTemplate.Merge(m, src)
}
func (m *TaskTemplate) XXX_Size() int {
	return xxx_messageInfo_TaskTemplate.Size(m)
}
func (m *TaskTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_TaskTemplate proto.InternalMessageInfo

func (m *TaskTemplate) GetId() *Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TaskTemplate) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TaskTemplate) GetMetadata() *TaskMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TaskTemplate) GetInterface() *TypedInterface {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *TaskTemplate) GetCustom() *_struct.Struct {
	if m != nil {
		return m.Custom
	}
	return nil
}

type isTaskTemplate_Target interface {
	isTaskTemplate_Target()
}

type TaskTemplate_Container struct {
	Container *Container `protobuf:"bytes,6,opt,name=container,proto3,oneof"`
}

func (*TaskTemplate_Container) isTaskTemplate_Target() {}

func (m *TaskTemplate) GetTarget() isTaskTemplate_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *TaskTemplate) GetContainer() *Container {
	if x, ok := m.GetTarget().(*TaskTemplate_Container); ok {
		return x.Container
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TaskTemplate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TaskTemplate_Container)(nil),
	}
}

// Defines port properties for a container.
type ContainerPort struct {
	// Number of port to expose on the pod's IP address.
	// This must be a valid port number, 0 < x < 65536.
	ContainerPort        uint32   `protobuf:"varint,1,opt,name=container_port,json=containerPort,proto3" json:"container_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerPort) Reset()         { *m = ContainerPort{} }
func (m *ContainerPort) String() string { return proto.CompactTextString(m) }
func (*ContainerPort) ProtoMessage()    {}
func (*ContainerPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{4}
}

func (m *ContainerPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerPort.Unmarshal(m, b)
}
func (m *ContainerPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerPort.Marshal(b, m, deterministic)
}
func (m *ContainerPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerPort.Merge(m, src)
}
func (m *ContainerPort) XXX_Size() int {
	return xxx_messageInfo_ContainerPort.Size(m)
}
func (m *ContainerPort) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerPort.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerPort proto.InternalMessageInfo

func (m *ContainerPort) GetContainerPort() uint32 {
	if m != nil {
		return m.ContainerPort
	}
	return 0
}

type Container struct {
	// Container image url. Eg: docker/redis:latest
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// Command to be executed, if not provided, the default entrypoint in the container image will be used.
	Command []string `protobuf:"bytes,2,rep,name=command,proto3" json:"command,omitempty"`
	// These will default to Flyte given paths. If provided, the system will not append known paths. If the task still
	// needs flyte's inputs and outputs path, add $(FLYTE_INPUT_FILE), $(FLYTE_OUTPUT_FILE) wherever makes sense and the
	// system will populate these before executing the container.
	Args []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	// Container resources requirement as specified by the container engine.
	Resources *Resources `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
	// Environment variables will be set as the container is starting up.
	Env []*KeyValuePair `protobuf:"bytes,5,rep,name=env,proto3" json:"env,omitempty"`
	// Allows extra configs to be available for the container.
	// TODO: elaborate on how configs will become available.
	Config []*KeyValuePair `protobuf:"bytes,6,rep,name=config,proto3" json:"config,omitempty"`
	// Ports to open in the container. This feature is not supported by all execution engines. (e.g. supported on K8s but
	// not supported on AWS Batch)
	// Only K8s
	Ports []*ContainerPort `protobuf:"bytes,7,rep,name=ports,proto3" json:"ports,omitempty"`
	// BETA: Optional configuration for DataLoading. If not specified, then default values are used.
	// This makes it possible to to run a completely portable container, that uses inputs and outputs
	// only from the local file-system and without having any reference to flyteidl. This is supported only on K8s at the moment.
	// If data loading is enabled, then data will be mounted in accompanying directories specified in the DataLoadingConfig. If the directories
	// are not specified, inputs will be mounted onto and outputs will be uploaded from a pre-determined file-system path. Refer to the documentation
	// to understand the default paths.
	// Only K8s
	DataConfig           *DataLoadingConfig `protobuf:"bytes,9,opt,name=data_config,json=dataConfig,proto3" json:"data_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{5}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Container) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Container) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Container) GetEnv() []*KeyValuePair {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Container) GetConfig() []*KeyValuePair {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Container) GetPorts() []*ContainerPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Container) GetDataConfig() *DataLoadingConfig {
	if m != nil {
		return m.DataConfig
	}
	return nil
}

// This configuration allows executing raw containers in Flyte using the Flyte CoPilot system.
// Flyte CoPilot, eliminates the needs of flytekit or sdk inside the container. Any inputs required by the users container are side-loaded in the input_path
// Any outputs generated by the user container - within output_path are automatically uploaded.
type DataLoadingConfig struct {
	// File system path (start at root). This folder will contain all the inputs exploded to a separate file.
	// Example, if the input interface needs (x: int, y: blob, z: multipart_blob) and the input path is "/var/flyte/inputs", then the file system will look like
	// /var/flyte/inputs/inputs.<metadata format dependent -> .pb .json .yaml> -> Format as defined previously. The Blob and Multipart blob will reference local filesystem instead of remote locations
	// /var/flyte/inputs/x -> X is a file that contains the value of x (integer) in string format
	// /var/flyte/inputs/y -> Y is a file in Binary format
	// /var/flyte/inputs/z/... -> Note Z itself is a directory
	// More information about the protocol - refer to docs #TODO reference docs here
	InputPath string `protobuf:"bytes,1,opt,name=input_path,json=inputPath,proto3" json:"input_path,omitempty"`
	// File system path (start at root). This folder should contain all the outputs for the task as individual files and/or an error text file
	OutputPath string `protobuf:"bytes,2,opt,name=output_path,json=outputPath,proto3" json:"output_path,omitempty"`
	// In the inputs folder, there will be an additional summary/metadata file that contains references to all files or inlined primitive values.
	// This format decides the actual encoding for the data. Refer to the encoding to understand the specifics of the contents and the encoding
	Format DataLoadingConfig_MetadataFormat `protobuf:"varint,3,opt,name=format,proto3,enum=flyteidl.core.DataLoadingConfig_MetadataFormat" json:"format,omitempty"`
	// Flag enables DataLoading Config. If this is not set, data loading will not be used!
	Enabled              bool                           `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DownloadStrategy     DataLoadingConfig_BlobDownload `protobuf:"varint,5,opt,name=download_strategy,json=downloadStrategy,proto3,enum=flyteidl.core.DataLoadingConfig_BlobDownload" json:"download_strategy,omitempty"`
	UploadStrategy       DataLoadingConfig_BlobUpload   `protobuf:"varint,6,opt,name=upload_strategy,json=uploadStrategy,proto3,enum=flyteidl.core.DataLoadingConfig_BlobUpload" json:"upload_strategy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *DataLoadingConfig) Reset()         { *m = DataLoadingConfig{} }
func (m *DataLoadingConfig) String() string { return proto.CompactTextString(m) }
func (*DataLoadingConfig) ProtoMessage()    {}
func (*DataLoadingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{6}
}

func (m *DataLoadingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataLoadingConfig.Unmarshal(m, b)
}
func (m *DataLoadingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataLoadingConfig.Marshal(b, m, deterministic)
}
func (m *DataLoadingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataLoadingConfig.Merge(m, src)
}
func (m *DataLoadingConfig) XXX_Size() int {
	return xxx_messageInfo_DataLoadingConfig.Size(m)
}
func (m *DataLoadingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DataLoadingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DataLoadingConfig proto.InternalMessageInfo

func (m *DataLoadingConfig) GetInputPath() string {
	if m != nil {
		return m.InputPath
	}
	return ""
}

func (m *DataLoadingConfig) GetOutputPath() string {
	if m != nil {
		return m.OutputPath
	}
	return ""
}

func (m *DataLoadingConfig) GetFormat() DataLoadingConfig_MetadataFormat {
	if m != nil {
		return m.Format
	}
	return DataLoadingConfig_JSON
}

func (m *DataLoadingConfig) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *DataLoadingConfig) GetDownloadStrategy() DataLoadingConfig_BlobDownload {
	if m != nil {
		return m.DownloadStrategy
	}
	return DataLoadingConfig_BEFORE_STARTUP
}

func (m *DataLoadingConfig) GetUploadStrategy() DataLoadingConfig_BlobUpload {
	if m != nil {
		return m.UploadStrategy
	}
	return DataLoadingConfig_ON_EXIT
}

func init() {
	proto.RegisterEnum("flyteidl.core.Resources_ResourceName", Resources_ResourceName_name, Resources_ResourceName_value)
	proto.RegisterEnum("flyteidl.core.RuntimeMetadata_RuntimeType", RuntimeMetadata_RuntimeType_name, RuntimeMetadata_RuntimeType_value)
	proto.RegisterEnum("flyteidl.core.DataLoadingConfig_MetadataFormat", DataLoadingConfig_MetadataFormat_name, DataLoadingConfig_MetadataFormat_value)
	proto.RegisterEnum("flyteidl.core.DataLoadingConfig_BlobDownload", DataLoadingConfig_BlobDownload_name, DataLoadingConfig_BlobDownload_value)
	proto.RegisterEnum("flyteidl.core.DataLoadingConfig_BlobUpload", DataLoadingConfig_BlobUpload_name, DataLoadingConfig_BlobUpload_value)
	proto.RegisterType((*Resources)(nil), "flyteidl.core.Resources")
	proto.RegisterType((*Resources_ResourceEntry)(nil), "flyteidl.core.Resources.ResourceEntry")
	proto.RegisterType((*RuntimeMetadata)(nil), "flyteidl.core.RuntimeMetadata")
	proto.RegisterType((*TaskMetadata)(nil), "flyteidl.core.TaskMetadata")
	proto.RegisterType((*TaskTemplate)(nil), "flyteidl.core.TaskTemplate")
	proto.RegisterType((*ContainerPort)(nil), "flyteidl.core.ContainerPort")
	proto.RegisterType((*Container)(nil), "flyteidl.core.Container")
	proto.RegisterType((*DataLoadingConfig)(nil), "flyteidl.core.DataLoadingConfig")
}

func init() { proto.RegisterFile("flyteidl/core/tasks.proto", fileDescriptor_bd8423ba58d6ed80) }

var fileDescriptor_bd8423ba58d6ed80 = []byte{
	// 1134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xff, 0x6e, 0xdb, 0x36,
	0x10, 0x8e, 0x1d, 0xc7, 0x3f, 0xce, 0xb1, 0xab, 0xb0, 0x5b, 0xa7, 0x66, 0x6d, 0x17, 0x18, 0x68,
	0xd7, 0xad, 0xa8, 0x3d, 0xb8, 0x40, 0xd7, 0x61, 0x40, 0x31, 0x3b, 0x56, 0x9a, 0xac, 0x89, 0x65,
	0xd0, 0x4a, 0xbb, 0xf6, 0x1f, 0x8d, 0x96, 0x68, 0x57, 0xa8, 0x24, 0x6a, 0x14, 0x95, 0xc1, 0x6f,
	0xb3, 0x07, 0xd8, 0x03, 0x6c, 0xaf, 0xb3, 0x17, 0xd9, 0x40, 0xea, 0x47, 0x22, 0x17, 0x5d, 0xb6,
	0xbf, 0xcc, 0xe3, 0x7d, 0xdf, 0xf1, 0xf8, 0xdd, 0xf9, 0x28, 0xb8, 0xbd, 0xf4, 0xd7, 0x82, 0x7a,
	0xae, 0x3f, 0x70, 0x18, 0xa7, 0x03, 0x41, 0xe2, 0xf7, 0x71, 0x3f, 0xe2, 0x4c, 0x30, 0xd4, 0xc9,
	0x5d, 0x7d, 0xe9, 0xda, 0xbf, 0x57, 0x46, 0x7a, 0x2e, 0x0d, 0x85, 0xb7, 0xf4, 0x28, 0x4f, 0xe1,
	0xfb, 0x77, 0x37, 0xfc, 0xa1, 0xa0, 0x7c, 0x49, 0x1c, 0x9a, 0xb9, 0xef, 0x94, 0xdd, 0xbe, 0x27,
	0x28, 0x27, 0x7e, 0x76, 0xd6, 0xfe, 0xbd, 0x15, 0x63, 0x2b, 0x9f, 0x0e, 0x94, 0xb5, 0x48, 0x96,
	0x03, 0x37, 0xe1, 0x44, 0x78, 0x2c, 0xcc, 0xd9, 0x9b, 0xfe, 0x58, 0xf0, 0xc4, 0x11, 0xa9, 0xb7,
	0xf7, 0x67, 0x15, 0x5a, 0x98, 0xc6, 0x2c, 0xe1, 0x0e, 0x8d, 0xd1, 0x18, 0x9a, 0x9c, 0xfe, 0x92,
	0xd0, 0x58, 0xc4, 0x7a, 0xe5, 0x60, 0xfb, 0x61, 0x7b, 0xf8, 0xa0, 0x5f, 0xba, 0x4a, 0xbf, 0xc0,
	0x16, 0x2b, 0x23, 0x14, 0x7c, 0x8d, 0x0b, 0x1e, 0x7a, 0x0e, 0x75, 0xdf, 0x0b, 0x3c, 0x11, 0xeb,
	0xd5, 0xff, 0x15, 0x21, 0x63, 0xed, 0xff, 0x0c, 0x9d, 0x92, 0x03, 0x7d, 0x07, 0xb5, 0x90, 0x04,
	0x54, 0xaf, 0x1c, 0x54, 0x1e, 0x76, 0x87, 0xf7, 0xaf, 0x0d, 0x37, 0x25, 0x01, 0xc5, 0x8a, 0x82,
	0x3e, 0x81, 0x9d, 0x0b, 0xe2, 0x27, 0x54, 0xaf, 0x1e, 0x54, 0x1e, 0xb6, 0x70, 0x6a, 0xf4, 0x8e,
	0x60, 0xf7, 0x2a, 0x16, 0xb5, 0xa1, 0x71, 0x3e, 0x7d, 0x39, 0x35, 0x5f, 0x4f, 0xb5, 0x2d, 0xd4,
	0x80, 0xed, 0xc3, 0xd9, 0xb9, 0x56, 0x91, 0x8b, 0x17, 0xb3, 0x73, 0xad, 0x8a, 0x00, 0xea, 0x67,
	0xc6, 0x99, 0x89, 0xdf, 0x68, 0xdb, 0x12, 0x3a, 0xb7, 0x4c, 0x3c, 0x7a, 0x61, 0x68, 0xb5, 0xde,
	0xef, 0x15, 0xb8, 0x81, 0x93, 0x50, 0x78, 0x01, 0x3d, 0xa3, 0x82, 0xb8, 0x44, 0x10, 0xf4, 0x1c,
	0x6a, 0x62, 0x1d, 0xe5, 0xc9, 0x7e, 0xbd, 0x99, 0x6c, 0x19, 0x9d, 0xdb, 0xd6, 0x3a, 0xa2, 0x58,
	0xf1, 0x90, 0x0e, 0x8d, 0x0b, 0xca, 0x63, 0x8f, 0x85, 0x59, 0xce, 0xb9, 0x89, 0x6e, 0x41, 0x7d,
	0xe9, 0x93, 0x0b, 0xc6, 0xf5, 0x6d, 0xe5, 0xc8, 0xac, 0xde, 0x97, 0xd0, 0xbe, 0x12, 0x06, 0xb5,
	0x60, 0xc7, 0xb4, 0x8e, 0x0d, 0xac, 0x6d, 0xa1, 0x0e, 0xb4, 0x8e, 0x4e, 0xdf, 0x58, 0x86, 0x3d,
	0x9f, 0xbc, 0xd4, 0x2a, 0xbd, 0xbf, 0xab, 0xb0, 0x6b, 0x91, 0xf8, 0x7d, 0x91, 0x6b, 0x0f, 0x76,
	0x5d, 0x2f, 0x76, 0xd8, 0x05, 0xe5, 0x64, 0xe1, 0xa7, 0x39, 0x37, 0x71, 0x69, 0x0f, 0x3d, 0x83,
	0x06, 0x4f, 0xa3, 0xab, 0x7c, 0xda, 0xc3, 0x7b, 0xff, 0x7e, 0x25, 0x9c, 0xc3, 0xd1, 0x13, 0x68,
	0xc8, 0x5f, 0x96, 0x08, 0xbd, 0xa6, 0x98, 0xb7, 0xfb, 0x69, 0x27, 0xf6, 0xf3, 0x4e, 0xec, 0x4f,
	0xb2, 0x4e, 0xc5, 0x39, 0x12, 0x3d, 0x85, 0x06, 0xa7, 0x82, 0x7b, 0x34, 0xd6, 0x77, 0x14, 0xe9,
	0xce, 0x07, 0xe5, 0x16, 0x7c, 0x3d, 0x17, 0x9c, 0x08, 0xba, 0x5a, 0xe3, 0x1c, 0x8c, 0x1e, 0xc1,
	0x5e, 0x9e, 0xf6, 0xda, 0xce, 0x05, 0xac, 0x2b, 0x9d, 0xb4, 0xc2, 0xf1, 0x2a, 0x53, 0xf2, 0x19,
	0xe8, 0x2e, 0x8d, 0x38, 0x75, 0x88, 0xa0, 0xae, 0x4d, 0x39, 0x67, 0xdc, 0x0e, 0x68, 0x1c, 0x93,
	0x15, 0xd5, 0x1b, 0x8a, 0x73, 0xeb, 0xd2, 0x6f, 0x48, 0xf7, 0x59, 0xea, 0x45, 0x0f, 0xa0, 0xa3,
	0xfe, 0x9c, 0x3c, 0x89, 0x84, 0x27, 0x25, 0x6b, 0x4a, 0xc9, 0x8e, 0xb7, 0x70, 0x79, 0x7b, 0xfc,
	0x29, 0xdc, 0x2c, 0x6d, 0xd8, 0x69, 0xe3, 0xfd, 0x91, 0x55, 0xc0, 0xa2, 0x41, 0xe4, 0x13, 0x41,
	0xd1, 0x57, 0x50, 0xf5, 0x5c, 0xa5, 0xbb, 0x94, 0xa7, 0x7c, 0xd3, 0x93, 0x62, 0x4a, 0xe0, 0xaa,
	0xe7, 0x22, 0x94, 0x35, 0x56, 0xda, 0x15, 0x69, 0xb3, 0x7c, 0x0b, 0xcd, 0x20, 0xd3, 0x5d, 0x35,
	0x45, 0x7b, 0xf8, 0xf9, 0x46, 0x90, 0xab, 0xf5, 0xc6, 0x05, 0x18, 0x7d, 0x0f, 0xad, 0x62, 0xc8,
	0x64, 0xd5, 0xb9, 0xbb, 0xc9, 0x5c, 0x47, 0xd4, 0x3d, 0xc9, 0x41, 0xf8, 0x12, 0x8f, 0x06, 0x50,
	0x77, 0x92, 0x58, 0xb0, 0x20, 0x2b, 0xd1, 0x67, 0x1f, 0xd4, 0x75, 0xae, 0x26, 0x0c, 0xce, 0x60,
	0xe8, 0x19, 0xb4, 0x1c, 0x16, 0x0a, 0xe2, 0x85, 0x94, 0xab, 0xa2, 0xb4, 0x87, 0xfa, 0xc6, 0x69,
	0x87, 0xb9, 0xff, 0x78, 0x0b, 0x5f, 0x82, 0xc7, 0x4d, 0xa8, 0x0b, 0xc2, 0x57, 0x54, 0xf4, 0x9e,
	0x42, 0xa7, 0xc0, 0xcc, 0x18, 0x17, 0xe8, 0x3e, 0x74, 0x0b, 0x9c, 0x1d, 0x31, 0x2e, 0x94, 0x8c,
	0x1d, 0xdc, 0x71, 0xae, 0xc2, 0x7a, 0x7f, 0x55, 0xa1, 0x55, 0x10, 0xe5, 0x3c, 0xf0, 0x02, 0x59,
	0xe6, 0x4a, 0x3a, 0x0f, 0x94, 0x21, 0xff, 0x73, 0x0e, 0x0b, 0x02, 0x12, 0xba, 0x6a, 0x64, 0xb5,
	0x70, 0x6e, 0x4a, 0xd1, 0x09, 0x5f, 0xc5, 0xfa, 0xb6, 0xda, 0x56, 0x6b, 0xf4, 0x14, 0x5a, 0x3c,
	0x9f, 0x39, 0x99, 0x76, 0xfa, 0xc7, 0x66, 0x12, 0xbe, 0x84, 0xa2, 0xc7, 0xb0, 0x4d, 0xc3, 0x0b,
	0x7d, 0x47, 0x0d, 0xc5, 0xcd, 0x3a, 0xbd, 0xa4, 0xeb, 0x57, 0xb2, 0x45, 0x66, 0xc4, 0xe3, 0x58,
	0xe2, 0xd0, 0x13, 0xa8, 0x3b, 0x2c, 0x5c, 0x7a, 0x2b, 0xbd, 0x7e, 0x3d, 0x23, 0x83, 0xa2, 0x21,
	0xec, 0x48, 0x29, 0x62, 0xbd, 0xa1, 0x38, 0x77, 0x3e, 0xa6, 0xb2, 0x94, 0x06, 0xa7, 0x50, 0x34,
	0x82, 0xb6, 0xec, 0x09, 0x3b, 0x3b, 0xad, 0xa5, 0x6e, 0x74, 0xb0, 0xc1, 0x9c, 0x10, 0x41, 0x4e,
	0x19, 0x71, 0xbd, 0x70, 0x75, 0xa8, 0x70, 0x18, 0x24, 0x29, 0x5d, 0xf7, 0x7e, 0xab, 0xc1, 0xde,
	0x07, 0x08, 0x74, 0x17, 0xc0, 0x0b, 0xa3, 0x44, 0xd8, 0x11, 0x11, 0xef, 0x32, 0xc5, 0x5b, 0x6a,
	0x67, 0x46, 0xc4, 0x3b, 0xf4, 0x05, 0xb4, 0x59, 0x22, 0x0a, 0x7f, 0xda, 0xd7, 0x90, 0x6e, 0x29,
	0xc0, 0x0b, 0xa8, 0x2f, 0x19, 0x0f, 0x88, 0x50, 0xbd, 0xdd, 0x1d, 0x0e, 0xae, 0xcb, 0xa9, 0x9f,
	0x77, 0xfa, 0x91, 0xa2, 0xe1, 0x8c, 0x2e, 0xeb, 0x4b, 0x43, 0x39, 0xcd, 0x5c, 0x55, 0xaf, 0x26,
	0xce, 0x4d, 0xf4, 0x16, 0xf6, 0x5c, 0xf6, 0x6b, 0xe8, 0x33, 0xe2, 0xda, 0x71, 0x36, 0x54, 0x54,
	0x57, 0x77, 0x87, 0x8f, 0xaf, 0x3d, 0x6d, 0xec, 0xb3, 0xc5, 0x24, 0x63, 0x63, 0x2d, 0x8f, 0x93,
	0xcf, 0x26, 0x64, 0xc1, 0x8d, 0x24, 0x2a, 0x47, 0xae, 0xab, 0xc8, 0x8f, 0xfe, 0x53, 0xe4, 0x73,
	0xc5, 0xc5, 0xdd, 0x34, 0x46, 0x1e, 0xb5, 0x37, 0x80, 0x6e, 0xf9, 0x96, 0xa8, 0x09, 0xb5, 0x1f,
	0xe7, 0xa6, 0x7c, 0xba, 0x9a, 0x50, 0x7b, 0x33, 0x3a, 0x3b, 0xd5, 0x2a, 0xf2, 0x11, 0x98, 0x61,
	0xd3, 0x32, 0xb5, 0x6a, 0xef, 0x10, 0x76, 0xaf, 0x26, 0x8a, 0x10, 0x74, 0xc7, 0xc6, 0x91, 0x89,
	0x0d, 0x7b, 0x6e, 0x8d, 0xb0, 0x75, 0x3e, 0xd3, 0xb6, 0xe4, 0x0b, 0x37, 0xb7, 0xb0, 0x31, 0x3a,
	0xd3, 0x2a, 0xe8, 0x26, 0xdc, 0x98, 0x98, 0xf6, 0xd4, 0xb4, 0xec, 0x89, 0xf9, 0x7a, 0x7a, 0x6a,
	0x8e, 0x26, 0x5a, 0xb5, 0xf7, 0x03, 0xc0, 0x65, 0x4e, 0xf2, 0x11, 0x34, 0xa7, 0xb6, 0xf1, 0xd3,
	0x89, 0xa5, 0x6d, 0xc9, 0x78, 0xaf, 0x8f, 0x8d, 0xa9, 0x3d, 0x7a, 0x35, 0x3a, 0x39, 0x1d, 0x8d,
	0x4f, 0x0d, 0xad, 0x82, 0xf6, 0xa0, 0x93, 0xc5, 0x38, 0x9f, 0xa5, 0x11, 0xc6, 0xc3, 0xb7, 0xdf,
	0xac, 0x3c, 0xf1, 0x2e, 0x59, 0xf4, 0x1d, 0x16, 0x0c, 0xfc, 0xf5, 0x52, 0x0c, 0x8a, 0xaf, 0x9a,
	0x15, 0x0d, 0x07, 0xd1, 0xe2, 0xf1, 0x8a, 0x0d, 0x4a, 0x1f, 0x3a, 0x8b, 0xba, 0x1a, 0x28, 0x4f,
	0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x30, 0x04, 0x59, 0x69, 0x09, 0x00, 0x00,
}
