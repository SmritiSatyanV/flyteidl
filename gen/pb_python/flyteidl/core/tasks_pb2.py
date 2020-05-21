# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/core/tasks.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import identifier_pb2 as flyteidl_dot_core_dot_identifier__pb2
from flyteidl.core import interface_pb2 as flyteidl_dot_core_dot_interface__pb2
from flyteidl.core import literals_pb2 as flyteidl_dot_core_dot_literals__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/core/tasks.proto',
  package='flyteidl.core',
  syntax='proto3',
  serialized_options=_b('Z0github.com/lyft/flyteidl/gen/pb-go/flyteidl/core'),
  serialized_pb=_b('\n\x19\x66lyteidl/core/tasks.proto\x12\rflyteidl.core\x1a\x1e\x66lyteidl/core/identifier.proto\x1a\x1d\x66lyteidl/core/interface.proto\x1a\x1c\x66lyteidl/core/literals.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1cgoogle/protobuf/struct.proto\"\x9a\x02\n\tResources\x12\x38\n\x08requests\x18\x01 \x03(\x0b\x32&.flyteidl.core.Resources.ResourceEntry\x12\x36\n\x06limits\x18\x02 \x03(\x0b\x32&.flyteidl.core.Resources.ResourceEntry\x1aS\n\rResourceEntry\x12\x33\n\x04name\x18\x01 \x01(\x0e\x32%.flyteidl.core.Resources.ResourceName\x12\r\n\x05value\x18\x02 \x01(\t\"F\n\x0cResourceName\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x07\n\x03\x43PU\x10\x01\x12\x07\n\x03GPU\x10\x02\x12\n\n\x06MEMORY\x10\x03\x12\x0b\n\x07STORAGE\x10\x04\"\x95\x01\n\x0fRuntimeMetadata\x12\x38\n\x04type\x18\x01 \x01(\x0e\x32*.flyteidl.core.RuntimeMetadata.RuntimeType\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\x0e\n\x06\x66lavor\x18\x03 \x01(\t\"\'\n\x0bRuntimeType\x12\t\n\x05OTHER\x10\x00\x12\r\n\tFLYTE_SDK\x10\x01\"\x9d\x02\n\x0cTaskMetadata\x12\x14\n\x0c\x64iscoverable\x18\x01 \x01(\x08\x12/\n\x07runtime\x18\x02 \x01(\x0b\x32\x1e.flyteidl.core.RuntimeMetadata\x12*\n\x07timeout\x18\x04 \x01(\x0b\x32\x19.google.protobuf.Duration\x12-\n\x07retries\x18\x05 \x01(\x0b\x32\x1c.flyteidl.core.RetryStrategy\x12\x19\n\x11\x64iscovery_version\x18\x06 \x01(\t\x12 \n\x18\x64\x65precated_error_message\x18\x07 \x01(\t\x12\x17\n\rinterruptible\x18\x08 \x01(\x08H\x00\x42\x15\n\x13interruptible_value\"\x86\x02\n\x0cTaskTemplate\x12%\n\x02id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.Identifier\x12\x0c\n\x04type\x18\x02 \x01(\t\x12-\n\x08metadata\x18\x03 \x01(\x0b\x32\x1b.flyteidl.core.TaskMetadata\x12\x30\n\tinterface\x18\x04 \x01(\x0b\x32\x1d.flyteidl.core.TypedInterface\x12\'\n\x06\x63ustom\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12-\n\tcontainer\x18\x06 \x01(\x0b\x32\x18.flyteidl.core.ContainerH\x00\x42\x08\n\x06target\"\'\n\rContainerPort\x12\x16\n\x0e\x63ontainer_port\x18\x01 \x01(\r\"\xa1\x02\n\tContainer\x12\r\n\x05image\x18\x01 \x01(\t\x12\x0f\n\x07\x63ommand\x18\x02 \x03(\t\x12\x0c\n\x04\x61rgs\x18\x03 \x03(\t\x12+\n\tresources\x18\x04 \x01(\x0b\x32\x18.flyteidl.core.Resources\x12(\n\x03\x65nv\x18\x05 \x03(\x0b\x32\x1b.flyteidl.core.KeyValuePair\x12+\n\x06\x63onfig\x18\x06 \x03(\x0b\x32\x1b.flyteidl.core.KeyValuePair\x12+\n\x05ports\x18\x07 \x03(\x0b\x32\x1c.flyteidl.core.ContainerPort\x12\x35\n\x0b\x64\x61ta_config\x18\t \x01(\x0b\x32 .flyteidl.core.DataLoadingConfig\"\xd6\x03\n\x11\x44\x61taLoadingConfig\x12\x12\n\ninput_path\x18\x01 \x01(\t\x12\x13\n\x0boutput_path\x18\x02 \x01(\t\x12?\n\x06\x66ormat\x18\x03 \x01(\x0e\x32/.flyteidl.core.DataLoadingConfig.MetadataFormat\x12\x0f\n\x07\x65nabled\x18\x04 \x01(\x08\x12H\n\x11\x64ownload_strategy\x18\x05 \x01(\x0e\x32-.flyteidl.core.DataLoadingConfig.BlobDownload\x12\x44\n\x0fupload_strategy\x18\x06 \x01(\x0e\x32+.flyteidl.core.DataLoadingConfig.BlobUpload\"/\n\x0eMetadataFormat\x12\x08\n\x04JSON\x10\x00\x12\x08\n\x04YAML\x10\x01\x12\t\n\x05PROTO\x10\x02\"C\n\x0c\x42lobDownload\x12\x12\n\x0e\x42\x45\x46ORE_STARTUP\x10\x00\x12\n\n\x06STREAM\x10\x01\x12\x13\n\x0f\x44O_NOT_DOWNLOAD\x10\x02\"@\n\nBlobUpload\x12\x0b\n\x07ON_EXIT\x10\x00\x12\x12\n\x0eWHEN_AVAILABLE\x10\x01\x12\x11\n\rDO_NOT_UPLOAD\x10\x02\x42\x32Z0github.com/lyft/flyteidl/gen/pb-go/flyteidl/coreb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_identifier__pb2.DESCRIPTOR,flyteidl_dot_core_dot_interface__pb2.DESCRIPTOR,flyteidl_dot_core_dot_literals__pb2.DESCRIPTOR,google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,google_dot_protobuf_dot_struct__pb2.DESCRIPTOR,])



_RESOURCES_RESOURCENAME = _descriptor.EnumDescriptor(
  name='ResourceName',
  full_name='flyteidl.core.Resources.ResourceName',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CPU', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='GPU', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MEMORY', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STORAGE', index=4, number=4,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=412,
  serialized_end=482,
)
_sym_db.RegisterEnumDescriptor(_RESOURCES_RESOURCENAME)

_RUNTIMEMETADATA_RUNTIMETYPE = _descriptor.EnumDescriptor(
  name='RuntimeType',
  full_name='flyteidl.core.RuntimeMetadata.RuntimeType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='OTHER', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FLYTE_SDK', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=595,
  serialized_end=634,
)
_sym_db.RegisterEnumDescriptor(_RUNTIMEMETADATA_RUNTIMETYPE)

_DATALOADINGCONFIG_METADATAFORMAT = _descriptor.EnumDescriptor(
  name='MetadataFormat',
  full_name='flyteidl.core.DataLoadingConfig.MetadataFormat',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='JSON', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='YAML', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1811,
  serialized_end=1858,
)
_sym_db.RegisterEnumDescriptor(_DATALOADINGCONFIG_METADATAFORMAT)

_DATALOADINGCONFIG_BLOBDOWNLOAD = _descriptor.EnumDescriptor(
  name='BlobDownload',
  full_name='flyteidl.core.DataLoadingConfig.BlobDownload',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='BEFORE_STARTUP', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STREAM', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DO_NOT_DOWNLOAD', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1860,
  serialized_end=1927,
)
_sym_db.RegisterEnumDescriptor(_DATALOADINGCONFIG_BLOBDOWNLOAD)

_DATALOADINGCONFIG_BLOBUPLOAD = _descriptor.EnumDescriptor(
  name='BlobUpload',
  full_name='flyteidl.core.DataLoadingConfig.BlobUpload',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ON_EXIT', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='WHEN_AVAILABLE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DO_NOT_UPLOAD', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1929,
  serialized_end=1993,
)
_sym_db.RegisterEnumDescriptor(_DATALOADINGCONFIG_BLOBUPLOAD)


_RESOURCES_RESOURCEENTRY = _descriptor.Descriptor(
  name='ResourceEntry',
  full_name='flyteidl.core.Resources.ResourceEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='flyteidl.core.Resources.ResourceEntry.name', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='flyteidl.core.Resources.ResourceEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=327,
  serialized_end=410,
)

_RESOURCES = _descriptor.Descriptor(
  name='Resources',
  full_name='flyteidl.core.Resources',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='requests', full_name='flyteidl.core.Resources.requests', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limits', full_name='flyteidl.core.Resources.limits', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_RESOURCES_RESOURCEENTRY, ],
  enum_types=[
    _RESOURCES_RESOURCENAME,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=200,
  serialized_end=482,
)


_RUNTIMEMETADATA = _descriptor.Descriptor(
  name='RuntimeMetadata',
  full_name='flyteidl.core.RuntimeMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='flyteidl.core.RuntimeMetadata.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='version', full_name='flyteidl.core.RuntimeMetadata.version', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='flavor', full_name='flyteidl.core.RuntimeMetadata.flavor', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _RUNTIMEMETADATA_RUNTIMETYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=485,
  serialized_end=634,
)


_TASKMETADATA = _descriptor.Descriptor(
  name='TaskMetadata',
  full_name='flyteidl.core.TaskMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='discoverable', full_name='flyteidl.core.TaskMetadata.discoverable', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='runtime', full_name='flyteidl.core.TaskMetadata.runtime', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timeout', full_name='flyteidl.core.TaskMetadata.timeout', index=2,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='retries', full_name='flyteidl.core.TaskMetadata.retries', index=3,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='discovery_version', full_name='flyteidl.core.TaskMetadata.discovery_version', index=4,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='deprecated_error_message', full_name='flyteidl.core.TaskMetadata.deprecated_error_message', index=5,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='interruptible', full_name='flyteidl.core.TaskMetadata.interruptible', index=6,
      number=8, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='interruptible_value', full_name='flyteidl.core.TaskMetadata.interruptible_value',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=637,
  serialized_end=922,
)


_TASKTEMPLATE = _descriptor.Descriptor(
  name='TaskTemplate',
  full_name='flyteidl.core.TaskTemplate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.core.TaskTemplate.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='flyteidl.core.TaskTemplate.type', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='flyteidl.core.TaskTemplate.metadata', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='interface', full_name='flyteidl.core.TaskTemplate.interface', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='custom', full_name='flyteidl.core.TaskTemplate.custom', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='container', full_name='flyteidl.core.TaskTemplate.container', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='target', full_name='flyteidl.core.TaskTemplate.target',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=925,
  serialized_end=1187,
)


_CONTAINERPORT = _descriptor.Descriptor(
  name='ContainerPort',
  full_name='flyteidl.core.ContainerPort',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='container_port', full_name='flyteidl.core.ContainerPort.container_port', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1189,
  serialized_end=1228,
)


_CONTAINER = _descriptor.Descriptor(
  name='Container',
  full_name='flyteidl.core.Container',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='image', full_name='flyteidl.core.Container.image', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='command', full_name='flyteidl.core.Container.command', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='args', full_name='flyteidl.core.Container.args', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resources', full_name='flyteidl.core.Container.resources', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='env', full_name='flyteidl.core.Container.env', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='config', full_name='flyteidl.core.Container.config', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ports', full_name='flyteidl.core.Container.ports', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data_config', full_name='flyteidl.core.Container.data_config', index=7,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1231,
  serialized_end=1520,
)


_DATALOADINGCONFIG = _descriptor.Descriptor(
  name='DataLoadingConfig',
  full_name='flyteidl.core.DataLoadingConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='input_path', full_name='flyteidl.core.DataLoadingConfig.input_path', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='output_path', full_name='flyteidl.core.DataLoadingConfig.output_path', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='format', full_name='flyteidl.core.DataLoadingConfig.format', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='enabled', full_name='flyteidl.core.DataLoadingConfig.enabled', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='download_strategy', full_name='flyteidl.core.DataLoadingConfig.download_strategy', index=4,
      number=5, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='upload_strategy', full_name='flyteidl.core.DataLoadingConfig.upload_strategy', index=5,
      number=6, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _DATALOADINGCONFIG_METADATAFORMAT,
    _DATALOADINGCONFIG_BLOBDOWNLOAD,
    _DATALOADINGCONFIG_BLOBUPLOAD,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1523,
  serialized_end=1993,
)

_RESOURCES_RESOURCEENTRY.fields_by_name['name'].enum_type = _RESOURCES_RESOURCENAME
_RESOURCES_RESOURCEENTRY.containing_type = _RESOURCES
_RESOURCES.fields_by_name['requests'].message_type = _RESOURCES_RESOURCEENTRY
_RESOURCES.fields_by_name['limits'].message_type = _RESOURCES_RESOURCEENTRY
_RESOURCES_RESOURCENAME.containing_type = _RESOURCES
_RUNTIMEMETADATA.fields_by_name['type'].enum_type = _RUNTIMEMETADATA_RUNTIMETYPE
_RUNTIMEMETADATA_RUNTIMETYPE.containing_type = _RUNTIMEMETADATA
_TASKMETADATA.fields_by_name['runtime'].message_type = _RUNTIMEMETADATA
_TASKMETADATA.fields_by_name['timeout'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_TASKMETADATA.fields_by_name['retries'].message_type = flyteidl_dot_core_dot_literals__pb2._RETRYSTRATEGY
_TASKMETADATA.oneofs_by_name['interruptible_value'].fields.append(
  _TASKMETADATA.fields_by_name['interruptible'])
_TASKMETADATA.fields_by_name['interruptible'].containing_oneof = _TASKMETADATA.oneofs_by_name['interruptible_value']
_TASKTEMPLATE.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._IDENTIFIER
_TASKTEMPLATE.fields_by_name['metadata'].message_type = _TASKMETADATA
_TASKTEMPLATE.fields_by_name['interface'].message_type = flyteidl_dot_core_dot_interface__pb2._TYPEDINTERFACE
_TASKTEMPLATE.fields_by_name['custom'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_TASKTEMPLATE.fields_by_name['container'].message_type = _CONTAINER
_TASKTEMPLATE.oneofs_by_name['target'].fields.append(
  _TASKTEMPLATE.fields_by_name['container'])
_TASKTEMPLATE.fields_by_name['container'].containing_oneof = _TASKTEMPLATE.oneofs_by_name['target']
_CONTAINER.fields_by_name['resources'].message_type = _RESOURCES
_CONTAINER.fields_by_name['env'].message_type = flyteidl_dot_core_dot_literals__pb2._KEYVALUEPAIR
_CONTAINER.fields_by_name['config'].message_type = flyteidl_dot_core_dot_literals__pb2._KEYVALUEPAIR
_CONTAINER.fields_by_name['ports'].message_type = _CONTAINERPORT
_CONTAINER.fields_by_name['data_config'].message_type = _DATALOADINGCONFIG
_DATALOADINGCONFIG.fields_by_name['format'].enum_type = _DATALOADINGCONFIG_METADATAFORMAT
_DATALOADINGCONFIG.fields_by_name['download_strategy'].enum_type = _DATALOADINGCONFIG_BLOBDOWNLOAD
_DATALOADINGCONFIG.fields_by_name['upload_strategy'].enum_type = _DATALOADINGCONFIG_BLOBUPLOAD
_DATALOADINGCONFIG_METADATAFORMAT.containing_type = _DATALOADINGCONFIG
_DATALOADINGCONFIG_BLOBDOWNLOAD.containing_type = _DATALOADINGCONFIG
_DATALOADINGCONFIG_BLOBUPLOAD.containing_type = _DATALOADINGCONFIG
DESCRIPTOR.message_types_by_name['Resources'] = _RESOURCES
DESCRIPTOR.message_types_by_name['RuntimeMetadata'] = _RUNTIMEMETADATA
DESCRIPTOR.message_types_by_name['TaskMetadata'] = _TASKMETADATA
DESCRIPTOR.message_types_by_name['TaskTemplate'] = _TASKTEMPLATE
DESCRIPTOR.message_types_by_name['ContainerPort'] = _CONTAINERPORT
DESCRIPTOR.message_types_by_name['Container'] = _CONTAINER
DESCRIPTOR.message_types_by_name['DataLoadingConfig'] = _DATALOADINGCONFIG
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Resources = _reflection.GeneratedProtocolMessageType('Resources', (_message.Message,), dict(

  ResourceEntry = _reflection.GeneratedProtocolMessageType('ResourceEntry', (_message.Message,), dict(
    DESCRIPTOR = _RESOURCES_RESOURCEENTRY,
    __module__ = 'flyteidl.core.tasks_pb2'
    # @@protoc_insertion_point(class_scope:flyteidl.core.Resources.ResourceEntry)
    ))
  ,
  DESCRIPTOR = _RESOURCES,
  __module__ = 'flyteidl.core.tasks_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.Resources)
  ))
_sym_db.RegisterMessage(Resources)
_sym_db.RegisterMessage(Resources.ResourceEntry)

RuntimeMetadata = _reflection.GeneratedProtocolMessageType('RuntimeMetadata', (_message.Message,), dict(
  DESCRIPTOR = _RUNTIMEMETADATA,
  __module__ = 'flyteidl.core.tasks_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.RuntimeMetadata)
  ))
_sym_db.RegisterMessage(RuntimeMetadata)

TaskMetadata = _reflection.GeneratedProtocolMessageType('TaskMetadata', (_message.Message,), dict(
  DESCRIPTOR = _TASKMETADATA,
  __module__ = 'flyteidl.core.tasks_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.TaskMetadata)
  ))
_sym_db.RegisterMessage(TaskMetadata)

TaskTemplate = _reflection.GeneratedProtocolMessageType('TaskTemplate', (_message.Message,), dict(
  DESCRIPTOR = _TASKTEMPLATE,
  __module__ = 'flyteidl.core.tasks_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.TaskTemplate)
  ))
_sym_db.RegisterMessage(TaskTemplate)

ContainerPort = _reflection.GeneratedProtocolMessageType('ContainerPort', (_message.Message,), dict(
  DESCRIPTOR = _CONTAINERPORT,
  __module__ = 'flyteidl.core.tasks_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.ContainerPort)
  ))
_sym_db.RegisterMessage(ContainerPort)

Container = _reflection.GeneratedProtocolMessageType('Container', (_message.Message,), dict(
  DESCRIPTOR = _CONTAINER,
  __module__ = 'flyteidl.core.tasks_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.Container)
  ))
_sym_db.RegisterMessage(Container)

DataLoadingConfig = _reflection.GeneratedProtocolMessageType('DataLoadingConfig', (_message.Message,), dict(
  DESCRIPTOR = _DATALOADINGCONFIG,
  __module__ = 'flyteidl.core.tasks_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.DataLoadingConfig)
  ))
_sym_db.RegisterMessage(DataLoadingConfig)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
