# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: diode/v1/device_type.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()

DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(
    b'\n\x1a\x64iode/v1/device_type.proto\x12\x08\x64iode.v1\x1a\x19google/protobuf/any.proto\x1a\x17validate/validate.proto\"\xa2\x01\n\nDeviceType\x12\x42\n\x0cmanufacturer\x18\x01 \x01(\x0b\x32\x14.google.protobuf.AnyB\x08\xfa\x42\x05\xa2\x01\x02\x08\x01R\x0cmanufacturer\x12\x1f\n\x05model\x18\x02 \x01(\tB\t\xfa\x42\x06r\x04\x10\x01\x18\x64R\x05model\x12/\n\x04slug\x18\x03 \x01(\tB\x1b\xfa\x42\x18r\x16\x10\x01\x18\x64\x32\x10^[-a-zA-Z0-9_]+$R\x04slugBDZBgithub.com/netboxlabs/diode-internal/diode-sdk-go/diode/v1/diodepbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'diode.v1.device_type_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'ZBgithub.com/netboxlabs/diode-internal/diode-sdk-go/diode/v1/diodepb'
    _globals['_DEVICETYPE'].fields_by_name['manufacturer']._options = None
    _globals['_DEVICETYPE'].fields_by_name['manufacturer']._serialized_options = b'\372B\005\242\001\002\010\001'
    _globals['_DEVICETYPE'].fields_by_name['model']._options = None
    _globals['_DEVICETYPE'].fields_by_name['model']._serialized_options = b'\372B\006r\004\020\001\030d'
    _globals['_DEVICETYPE'].fields_by_name['slug']._options = None
    _globals['_DEVICETYPE'].fields_by_name[
        'slug']._serialized_options = b'\372B\030r\026\020\001\030d2\020^[-a-zA-Z0-9_]+$'
    _globals['_DEVICETYPE']._serialized_start = 93
    _globals['_DEVICETYPE']._serialized_end = 255
# @@protoc_insertion_point(module_scope)
