# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: diode/v1/site.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()

from validate import validate_pb2 as validate_dot_validate__pb2

DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(
    b'\n\x13\x64iode/v1/site.proto\x12\x08\x64iode.v1\x1a\x17validate/validate.proto\"r\n\x04Site\x12\"\n\x04name\x18\x01 \x01(\tB\t\xfa\x42\x06r\x04\x10\x01\x18\x64H\x00R\x04name\x88\x01\x01\x12\x34\n\x04slug\x18\x02 \x01(\tB\x1b\xfa\x42\x18r\x16\x10\x01\x18\x64\x32\x10^[-a-zA-Z0-9_]+$H\x01R\x04slug\x88\x01\x01\x42\x07\n\x05_nameB\x07\n\x05_slugBDZBgithub.com/netboxlabs/diode-internal/diode-sdk-go/diode/v1/diodepbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'diode.v1.site_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'ZBgithub.com/netboxlabs/diode-internal/diode-sdk-go/diode/v1/diodepb'
    _globals['_SITE'].fields_by_name['name']._options = None
    _globals['_SITE'].fields_by_name['name']._serialized_options = b'\372B\006r\004\020\001\030d'
    _globals['_SITE'].fields_by_name['slug']._options = None
    _globals['_SITE'].fields_by_name['slug']._serialized_options = b'\372B\030r\026\020\001\030d2\020^[-a-zA-Z0-9_]+$'
    _globals['_SITE']._serialized_start = 58
    _globals['_SITE']._serialized_end = 172
# @@protoc_insertion_point(module_scope)
