# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mysqlctl.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='mysqlctl.proto',
  package='mysqlctl',
  syntax='proto3',
  serialized_pb=_b('\n\x0emysqlctl.proto\x12\x08mysqlctl\"#\n\x0cStartRequest\x12\x13\n\x0bmysqld_args\x18\x01 \x03(\t\"\x0f\n\rStartResponse\"*\n\x0fShutdownRequest\x12\x17\n\x0fwait_for_mysqld\x18\x01 \x01(\x08\"\x12\n\x10ShutdownResponse\"\x18\n\x16RunMysqlUpgradeRequest\"\x19\n\x17RunMysqlUpgradeResponse\"\x15\n\x13ReinitConfigRequest\"\x16\n\x14ReinitConfigResponse2\xb6\x02\n\x08MysqlCtl\x12:\n\x05Start\x12\x16.mysqlctl.StartRequest\x1a\x17.mysqlctl.StartResponse\"\x00\x12\x43\n\x08Shutdown\x12\x19.mysqlctl.ShutdownRequest\x1a\x1a.mysqlctl.ShutdownResponse\"\x00\x12X\n\x0fRunMysqlUpgrade\x12 .mysqlctl.RunMysqlUpgradeRequest\x1a!.mysqlctl.RunMysqlUpgradeResponse\"\x00\x12O\n\x0cReinitConfig\x12\x1d.mysqlctl.ReinitConfigRequest\x1a\x1e.mysqlctl.ReinitConfigResponse\"\x00\x62\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_STARTREQUEST = _descriptor.Descriptor(
  name='StartRequest',
  full_name='mysqlctl.StartRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='mysqld_args', full_name='mysqlctl.StartRequest.mysqld_args', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=28,
  serialized_end=63,
)


_STARTRESPONSE = _descriptor.Descriptor(
  name='StartResponse',
  full_name='mysqlctl.StartResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=65,
  serialized_end=80,
)


_SHUTDOWNREQUEST = _descriptor.Descriptor(
  name='ShutdownRequest',
  full_name='mysqlctl.ShutdownRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='wait_for_mysqld', full_name='mysqlctl.ShutdownRequest.wait_for_mysqld', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=82,
  serialized_end=124,
)


_SHUTDOWNRESPONSE = _descriptor.Descriptor(
  name='ShutdownResponse',
  full_name='mysqlctl.ShutdownResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=126,
  serialized_end=144,
)


_RUNMYSQLUPGRADEREQUEST = _descriptor.Descriptor(
  name='RunMysqlUpgradeRequest',
  full_name='mysqlctl.RunMysqlUpgradeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=146,
  serialized_end=170,
)


_RUNMYSQLUPGRADERESPONSE = _descriptor.Descriptor(
  name='RunMysqlUpgradeResponse',
  full_name='mysqlctl.RunMysqlUpgradeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=172,
  serialized_end=197,
)


_REINITCONFIGREQUEST = _descriptor.Descriptor(
  name='ReinitConfigRequest',
  full_name='mysqlctl.ReinitConfigRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=199,
  serialized_end=220,
)


_REINITCONFIGRESPONSE = _descriptor.Descriptor(
  name='ReinitConfigResponse',
  full_name='mysqlctl.ReinitConfigResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=222,
  serialized_end=244,
)

DESCRIPTOR.message_types_by_name['StartRequest'] = _STARTREQUEST
DESCRIPTOR.message_types_by_name['StartResponse'] = _STARTRESPONSE
DESCRIPTOR.message_types_by_name['ShutdownRequest'] = _SHUTDOWNREQUEST
DESCRIPTOR.message_types_by_name['ShutdownResponse'] = _SHUTDOWNRESPONSE
DESCRIPTOR.message_types_by_name['RunMysqlUpgradeRequest'] = _RUNMYSQLUPGRADEREQUEST
DESCRIPTOR.message_types_by_name['RunMysqlUpgradeResponse'] = _RUNMYSQLUPGRADERESPONSE
DESCRIPTOR.message_types_by_name['ReinitConfigRequest'] = _REINITCONFIGREQUEST
DESCRIPTOR.message_types_by_name['ReinitConfigResponse'] = _REINITCONFIGRESPONSE

StartRequest = _reflection.GeneratedProtocolMessageType('StartRequest', (_message.Message,), dict(
  DESCRIPTOR = _STARTREQUEST,
  __module__ = 'mysqlctl_pb2'
  # @@protoc_insertion_point(class_scope:mysqlctl.StartRequest)
  ))
_sym_db.RegisterMessage(StartRequest)

StartResponse = _reflection.GeneratedProtocolMessageType('StartResponse', (_message.Message,), dict(
  DESCRIPTOR = _STARTRESPONSE,
  __module__ = 'mysqlctl_pb2'
  # @@protoc_insertion_point(class_scope:mysqlctl.StartResponse)
  ))
_sym_db.RegisterMessage(StartResponse)

ShutdownRequest = _reflection.GeneratedProtocolMessageType('ShutdownRequest', (_message.Message,), dict(
  DESCRIPTOR = _SHUTDOWNREQUEST,
  __module__ = 'mysqlctl_pb2'
  # @@protoc_insertion_point(class_scope:mysqlctl.ShutdownRequest)
  ))
_sym_db.RegisterMessage(ShutdownRequest)

ShutdownResponse = _reflection.GeneratedProtocolMessageType('ShutdownResponse', (_message.Message,), dict(
  DESCRIPTOR = _SHUTDOWNRESPONSE,
  __module__ = 'mysqlctl_pb2'
  # @@protoc_insertion_point(class_scope:mysqlctl.ShutdownResponse)
  ))
_sym_db.RegisterMessage(ShutdownResponse)

RunMysqlUpgradeRequest = _reflection.GeneratedProtocolMessageType('RunMysqlUpgradeRequest', (_message.Message,), dict(
  DESCRIPTOR = _RUNMYSQLUPGRADEREQUEST,
  __module__ = 'mysqlctl_pb2'
  # @@protoc_insertion_point(class_scope:mysqlctl.RunMysqlUpgradeRequest)
  ))
_sym_db.RegisterMessage(RunMysqlUpgradeRequest)

RunMysqlUpgradeResponse = _reflection.GeneratedProtocolMessageType('RunMysqlUpgradeResponse', (_message.Message,), dict(
  DESCRIPTOR = _RUNMYSQLUPGRADERESPONSE,
  __module__ = 'mysqlctl_pb2'
  # @@protoc_insertion_point(class_scope:mysqlctl.RunMysqlUpgradeResponse)
  ))
_sym_db.RegisterMessage(RunMysqlUpgradeResponse)

ReinitConfigRequest = _reflection.GeneratedProtocolMessageType('ReinitConfigRequest', (_message.Message,), dict(
  DESCRIPTOR = _REINITCONFIGREQUEST,
  __module__ = 'mysqlctl_pb2'
  # @@protoc_insertion_point(class_scope:mysqlctl.ReinitConfigRequest)
  ))
_sym_db.RegisterMessage(ReinitConfigRequest)

ReinitConfigResponse = _reflection.GeneratedProtocolMessageType('ReinitConfigResponse', (_message.Message,), dict(
  DESCRIPTOR = _REINITCONFIGRESPONSE,
  __module__ = 'mysqlctl_pb2'
  # @@protoc_insertion_point(class_scope:mysqlctl.ReinitConfigResponse)
  ))
_sym_db.RegisterMessage(ReinitConfigResponse)


import abc
from grpc.beta import implementations as beta_implementations
from grpc.framework.common import cardinality
from grpc.framework.interfaces.face import utilities as face_utilities

class BetaMysqlCtlServicer(object):
  """<fill me in later!>"""
  __metaclass__ = abc.ABCMeta
  @abc.abstractmethod
  def Start(self, request, context):
    raise NotImplementedError()
  @abc.abstractmethod
  def Shutdown(self, request, context):
    raise NotImplementedError()
  @abc.abstractmethod
  def RunMysqlUpgrade(self, request, context):
    raise NotImplementedError()
  @abc.abstractmethod
  def ReinitConfig(self, request, context):
    raise NotImplementedError()

class BetaMysqlCtlStub(object):
  """The interface to which stubs will conform."""
  __metaclass__ = abc.ABCMeta
  @abc.abstractmethod
  def Start(self, request, timeout):
    raise NotImplementedError()
  Start.future = None
  @abc.abstractmethod
  def Shutdown(self, request, timeout):
    raise NotImplementedError()
  Shutdown.future = None
  @abc.abstractmethod
  def RunMysqlUpgrade(self, request, timeout):
    raise NotImplementedError()
  RunMysqlUpgrade.future = None
  @abc.abstractmethod
  def ReinitConfig(self, request, timeout):
    raise NotImplementedError()
  ReinitConfig.future = None

def beta_create_MysqlCtl_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
  import mysqlctl_pb2
  import mysqlctl_pb2
  import mysqlctl_pb2
  import mysqlctl_pb2
  import mysqlctl_pb2
  import mysqlctl_pb2
  import mysqlctl_pb2
  import mysqlctl_pb2
  request_deserializers = {
    ('mysqlctl.MysqlCtl', 'ReinitConfig'): mysqlctl_pb2.ReinitConfigRequest.FromString,
    ('mysqlctl.MysqlCtl', 'RunMysqlUpgrade'): mysqlctl_pb2.RunMysqlUpgradeRequest.FromString,
    ('mysqlctl.MysqlCtl', 'Shutdown'): mysqlctl_pb2.ShutdownRequest.FromString,
    ('mysqlctl.MysqlCtl', 'Start'): mysqlctl_pb2.StartRequest.FromString,
  }
  response_serializers = {
    ('mysqlctl.MysqlCtl', 'ReinitConfig'): mysqlctl_pb2.ReinitConfigResponse.SerializeToString,
    ('mysqlctl.MysqlCtl', 'RunMysqlUpgrade'): mysqlctl_pb2.RunMysqlUpgradeResponse.SerializeToString,
    ('mysqlctl.MysqlCtl', 'Shutdown'): mysqlctl_pb2.ShutdownResponse.SerializeToString,
    ('mysqlctl.MysqlCtl', 'Start'): mysqlctl_pb2.StartResponse.SerializeToString,
  }
  method_implementations = {
    ('mysqlctl.MysqlCtl', 'ReinitConfig'): face_utilities.unary_unary_inline(servicer.ReinitConfig),
    ('mysqlctl.MysqlCtl', 'RunMysqlUpgrade'): face_utilities.unary_unary_inline(servicer.RunMysqlUpgrade),
    ('mysqlctl.MysqlCtl', 'Shutdown'): face_utilities.unary_unary_inline(servicer.Shutdown),
    ('mysqlctl.MysqlCtl', 'Start'): face_utilities.unary_unary_inline(servicer.Start),
  }
  server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
  return beta_implementations.server(method_implementations, options=server_options)

def beta_create_MysqlCtl_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
  import mysqlctl_pb2
  import mysqlctl_pb2
  import mysqlctl_pb2
  import mysqlctl_pb2
  import mysqlctl_pb2
  import mysqlctl_pb2
  import mysqlctl_pb2
  import mysqlctl_pb2
  request_serializers = {
    ('mysqlctl.MysqlCtl', 'ReinitConfig'): mysqlctl_pb2.ReinitConfigRequest.SerializeToString,
    ('mysqlctl.MysqlCtl', 'RunMysqlUpgrade'): mysqlctl_pb2.RunMysqlUpgradeRequest.SerializeToString,
    ('mysqlctl.MysqlCtl', 'Shutdown'): mysqlctl_pb2.ShutdownRequest.SerializeToString,
    ('mysqlctl.MysqlCtl', 'Start'): mysqlctl_pb2.StartRequest.SerializeToString,
  }
  response_deserializers = {
    ('mysqlctl.MysqlCtl', 'ReinitConfig'): mysqlctl_pb2.ReinitConfigResponse.FromString,
    ('mysqlctl.MysqlCtl', 'RunMysqlUpgrade'): mysqlctl_pb2.RunMysqlUpgradeResponse.FromString,
    ('mysqlctl.MysqlCtl', 'Shutdown'): mysqlctl_pb2.ShutdownResponse.FromString,
    ('mysqlctl.MysqlCtl', 'Start'): mysqlctl_pb2.StartResponse.FromString,
  }
  cardinalities = {
    'ReinitConfig': cardinality.Cardinality.UNARY_UNARY,
    'RunMysqlUpgrade': cardinality.Cardinality.UNARY_UNARY,
    'Shutdown': cardinality.Cardinality.UNARY_UNARY,
    'Start': cardinality.Cardinality.UNARY_UNARY,
  }
  stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
  return beta_implementations.dynamic_stub(channel, 'mysqlctl.MysqlCtl', cardinalities, options=stub_options)
# @@protoc_insertion_point(module_scope)
