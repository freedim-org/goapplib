//
//  Generated code. Do not modify.
//  source: goapplib.calling.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use codeDescriptor instead')
const Code$json = {
  '1': 'Code',
  '2': [
    {'1': 'OK', '2': 0},
    {'1': 'InvalidRequest', '2': 400},
    {'1': 'MethodNotFound', '2': 404},
    {'1': 'InternalError', '2': 500},
    {'1': 'MethodNullResponse', '2': 501},
  ],
};

/// Descriptor for `Code`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List codeDescriptor = $convert.base64Decode(
    'CgRDb2RlEgYKAk9LEAASEwoOSW52YWxpZFJlcXVlc3QQkAMSEwoOTWV0aG9kTm90Rm91bmQQlA'
    'MSEgoNSW50ZXJuYWxFcnJvchD0AxIXChJNZXRob2ROdWxsUmVzcG9uc2UQ9QM=');

@$core.Deprecated('Use goRequestDescriptor instead')
const GoRequest$json = {
  '1': 'GoRequest',
  '2': [
    {'1': 'traceId', '3': 1, '4': 1, '5': 9, '10': 'traceId'},
    {'1': 'method', '3': 2, '4': 1, '5': 9, '10': 'method'},
    {'1': 'data', '3': 3, '4': 1, '5': 12, '10': 'data'},
  ],
};

/// Descriptor for `GoRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List goRequestDescriptor = $convert.base64Decode(
    'CglHb1JlcXVlc3QSGAoHdHJhY2VJZBgBIAEoCVIHdHJhY2VJZBIWCgZtZXRob2QYAiABKAlSBm'
    '1ldGhvZBISCgRkYXRhGAMgASgMUgRkYXRh');

@$core.Deprecated('Use goResponseDescriptor instead')
const GoResponse$json = {
  '1': 'GoResponse',
  '2': [
    {'1': 'traceId', '3': 1, '4': 1, '5': 9, '10': 'traceId'},
    {'1': 'code', '3': 2, '4': 1, '5': 14, '6': '.goapplib.Code', '10': 'code'},
    {'1': 'data', '3': 3, '4': 1, '5': 12, '10': 'data'},
  ],
};

/// Descriptor for `GoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List goResponseDescriptor = $convert.base64Decode(
    'CgpHb1Jlc3BvbnNlEhgKB3RyYWNlSWQYASABKAlSB3RyYWNlSWQSIgoEY29kZRgCIAEoDjIOLm'
    'dvYXBwbGliLkNvZGVSBGNvZGUSEgoEZGF0YRgDIAEoDFIEZGF0YQ==');

