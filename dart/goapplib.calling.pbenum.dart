//
//  Generated code. Do not modify.
//  source: goapplib.calling.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Code extends $pb.ProtobufEnum {
  static const Code OK = Code._(0, _omitEnumNames ? '' : 'OK');
  static const Code InvalidRequest = Code._(400, _omitEnumNames ? '' : 'InvalidRequest');
  static const Code MethodNotFound = Code._(404, _omitEnumNames ? '' : 'MethodNotFound');
  static const Code InternalError = Code._(500, _omitEnumNames ? '' : 'InternalError');
  static const Code MethodNullResponse = Code._(501, _omitEnumNames ? '' : 'MethodNullResponse');

  static const $core.List<Code> values = <Code> [
    OK,
    InvalidRequest,
    MethodNotFound,
    InternalError,
    MethodNullResponse,
  ];

  static final $core.Map<$core.int, Code> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Code? valueOf($core.int value) => _byValue[value];

  const Code._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
