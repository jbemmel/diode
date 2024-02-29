# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: reconciler/v1/reconciler.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from netbox_diode_plugin.diode_reconciler_sdk.reconciler.v1 import device_pb2 as reconciler_dot_v1_dot_device__pb2
from netbox_diode_plugin.diode_reconciler_sdk.reconciler.v1 import device_role_pb2 as reconciler_dot_v1_dot_device__role__pb2
from netbox_diode_plugin.diode_reconciler_sdk.reconciler.v1 import device_type_pb2 as reconciler_dot_v1_dot_device__type__pb2
from netbox_diode_plugin.diode_reconciler_sdk.reconciler.v1 import manufacturer_pb2 as reconciler_dot_v1_dot_manufacturer__pb2
from netbox_diode_plugin.diode_reconciler_sdk.reconciler.v1 import platform_pb2 as reconciler_dot_v1_dot_platform__pb2
from netbox_diode_plugin.diode_reconciler_sdk.reconciler.v1 import site_pb2 as reconciler_dot_v1_dot_site__pb2
from netbox_diode_plugin.diode_reconciler_sdk.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ereconciler/v1/reconciler.proto\x12\rreconciler.v1\x1a\x1areconciler/v1/device.proto\x1a\x1freconciler/v1/device_role.proto\x1a\x1freconciler/v1/device_type.proto\x1a reconciler/v1/manufacturer.proto\x1a\x1creconciler/v1/platform.proto\x1a\x18reconciler/v1/site.proto\x1a\x17validate/validate.proto\"Y\n\x13IngestionDataSource\x12\x1e\n\x04name\x18\x01 \x01(\tB\n\xfa\x42\x07r\x05\x10\x01\x18\xff\x01R\x04name\x12\"\n\x07\x61pi_key\x18\x02 \x01(\tB\t\xfa\x42\x06r\x04\x10(\x18(R\x06\x61piKey\"\xab\x01\n#RetrieveIngestionDataSourcesRequest\x12\x1e\n\x04name\x18\x01 \x01(\tB\n\xfa\x42\x07r\x05\x10\x01\x18\xff\x01R\x04name\x12%\n\x08sdk_name\x18\x02 \x01(\tB\n\xfa\x42\x07r\x05\x10\x01\x18\xff\x01R\x07sdkName\x12=\n\x0bsdk_version\x18\x03 \x01(\tB\x1c\xfa\x42\x19r\x17\x32\x15^(\\d)+\\.(\\d)+\\.(\\d)+$R\nsdkVersion\"\x80\x01\n$RetrieveIngestionDataSourcesResponse\x12X\n\x16ingestion_data_sources\x18\x01 \x03(\x0b\x32\".reconciler.v1.IngestionDataSourceR\x14ingestionDataSources\"\xe4\x02\n\x06Object\x12/\n\x06\x64\x65vice\x18\x01 \x01(\x0b\x32\x15.reconciler.v1.DeviceH\x00R\x06\x64\x65vice\x12<\n\x0b\x64\x65vice_role\x18\x02 \x01(\x0b\x32\x19.reconciler.v1.DeviceRoleH\x00R\ndeviceRole\x12<\n\x0b\x64\x65vice_type\x18\x03 \x01(\x0b\x32\x19.reconciler.v1.DeviceTypeH\x00R\ndeviceType\x12\x41\n\x0cmanufacturer\x18\x04 \x01(\x0b\x32\x1b.reconciler.v1.ManufacturerH\x00R\x0cmanufacturer\x12\x35\n\x08platform\x18\x05 \x01(\x0b\x32\x17.reconciler.v1.PlatformH\x00R\x08platform\x12)\n\x04site\x18\x06 \x01(\x0b\x32\x13.reconciler.v1.SiteH\x00R\x04siteB\x08\n\x06object\"\xb2\x02\n\x15\x41\x64\x64ObjectStateRequest\x12$\n\tobject_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02(\x01R\x08objectId\x12\x31\n\x10object_change_id\x18\x02 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02(\x01R\x0eobjectChangeId\x12+\n\x0bobject_type\x18\x03 \x01(\tB\n\xfa\x42\x07r\x05\x10\x01\x18\xff\x01R\nobjectType\x12-\n\x06object\x18\x04 \x01(\x0b\x32\x15.reconciler.v1.ObjectR\x06object\x12%\n\x08sdk_name\x18\x05 \x01(\tB\n\xfa\x42\x07r\x05\x10\x01\x18\xff\x01R\x07sdkName\x12=\n\x0bsdk_version\x18\x06 \x01(\tB\x1c\xfa\x42\x19r\x17\x32\x15^(\\d)+\\.(\\d)+\\.(\\d)+$R\nsdkVersion\"0\n\x16\x41\x64\x64ObjectStateResponse\x12\x16\n\x06\x65rrors\x18\x01 \x03(\tR\x06\x65rrors2\x80\x02\n\x11ReconcilerService\x12\x89\x01\n\x1cRetrieveIngestionDataSources\x12\x32.reconciler.v1.RetrieveIngestionDataSourcesRequest\x1a\x33.reconciler.v1.RetrieveIngestionDataSourcesResponse\"\x00\x12_\n\x0e\x41\x64\x64ObjectState\x12$.reconciler.v1.AddObjectStateRequest\x1a%.reconciler.v1.AddObjectStateResponse\"\x00\x42\x45ZCgithub.com/netboxlabs/diode/diode-server/reconciler/v1/reconcilerpbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'reconciler.v1.reconciler_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZCgithub.com/netboxlabs/diode/diode-server/reconciler/v1/reconcilerpb'
  _globals['_INGESTIONDATASOURCE'].fields_by_name['name']._options = None
  _globals['_INGESTIONDATASOURCE'].fields_by_name['name']._serialized_options = b'\372B\007r\005\020\001\030\377\001'
  _globals['_INGESTIONDATASOURCE'].fields_by_name['api_key']._options = None
  _globals['_INGESTIONDATASOURCE'].fields_by_name['api_key']._serialized_options = b'\372B\006r\004\020(\030('
  _globals['_RETRIEVEINGESTIONDATASOURCESREQUEST'].fields_by_name['name']._options = None
  _globals['_RETRIEVEINGESTIONDATASOURCESREQUEST'].fields_by_name['name']._serialized_options = b'\372B\007r\005\020\001\030\377\001'
  _globals['_RETRIEVEINGESTIONDATASOURCESREQUEST'].fields_by_name['sdk_name']._options = None
  _globals['_RETRIEVEINGESTIONDATASOURCESREQUEST'].fields_by_name['sdk_name']._serialized_options = b'\372B\007r\005\020\001\030\377\001'
  _globals['_RETRIEVEINGESTIONDATASOURCESREQUEST'].fields_by_name['sdk_version']._options = None
  _globals['_RETRIEVEINGESTIONDATASOURCESREQUEST'].fields_by_name['sdk_version']._serialized_options = b'\372B\031r\0272\025^(\\d)+\\.(\\d)+\\.(\\d)+$'
  _globals['_ADDOBJECTSTATEREQUEST'].fields_by_name['object_id']._options = None
  _globals['_ADDOBJECTSTATEREQUEST'].fields_by_name['object_id']._serialized_options = b'\372B\0042\002(\001'
  _globals['_ADDOBJECTSTATEREQUEST'].fields_by_name['object_change_id']._options = None
  _globals['_ADDOBJECTSTATEREQUEST'].fields_by_name['object_change_id']._serialized_options = b'\372B\0042\002(\001'
  _globals['_ADDOBJECTSTATEREQUEST'].fields_by_name['object_type']._options = None
  _globals['_ADDOBJECTSTATEREQUEST'].fields_by_name['object_type']._serialized_options = b'\372B\007r\005\020\001\030\377\001'
  _globals['_ADDOBJECTSTATEREQUEST'].fields_by_name['sdk_name']._options = None
  _globals['_ADDOBJECTSTATEREQUEST'].fields_by_name['sdk_name']._serialized_options = b'\372B\007r\005\020\001\030\377\001'
  _globals['_ADDOBJECTSTATEREQUEST'].fields_by_name['sdk_version']._options = None
  _globals['_ADDOBJECTSTATEREQUEST'].fields_by_name['sdk_version']._serialized_options = b'\372B\031r\0272\025^(\\d)+\\.(\\d)+\\.(\\d)+$'
  _globals['_INGESTIONDATASOURCE']._serialized_start=258
  _globals['_INGESTIONDATASOURCE']._serialized_end=347
  _globals['_RETRIEVEINGESTIONDATASOURCESREQUEST']._serialized_start=350
  _globals['_RETRIEVEINGESTIONDATASOURCESREQUEST']._serialized_end=521
  _globals['_RETRIEVEINGESTIONDATASOURCESRESPONSE']._serialized_start=524
  _globals['_RETRIEVEINGESTIONDATASOURCESRESPONSE']._serialized_end=652
  _globals['_OBJECT']._serialized_start=655
  _globals['_OBJECT']._serialized_end=1011
  _globals['_ADDOBJECTSTATEREQUEST']._serialized_start=1014
  _globals['_ADDOBJECTSTATEREQUEST']._serialized_end=1320
  _globals['_ADDOBJECTSTATERESPONSE']._serialized_start=1322
  _globals['_ADDOBJECTSTATERESPONSE']._serialized_end=1370
  _globals['_RECONCILERSERVICE']._serialized_start=1373
  _globals['_RECONCILERSERVICE']._serialized_end=1629
# @@protoc_insertion_point(module_scope)
