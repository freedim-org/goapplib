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

import 'goapplib.calling.pbenum.dart';

export 'goapplib.calling.pbenum.dart';

class GoRequest extends $pb.GeneratedMessage {
  factory GoRequest({
    $core.String? traceId,
    $core.String? method,
    $core.List<$core.int>? data,
  }) {
    final $result = create();
    if (traceId != null) {
      $result.traceId = traceId;
    }
    if (method != null) {
      $result.method = method;
    }
    if (data != null) {
      $result.data = data;
    }
    return $result;
  }
  GoRequest._() : super();
  factory GoRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GoRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GoRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'goapplib'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'traceId', protoName: 'traceId')
    ..aOS(2, _omitFieldNames ? '' : 'method')
    ..a<$core.List<$core.int>>(3, _omitFieldNames ? '' : 'data', $pb.PbFieldType.OY)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GoRequest clone() => GoRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GoRequest copyWith(void Function(GoRequest) updates) => super.copyWith((message) => updates(message as GoRequest)) as GoRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GoRequest create() => GoRequest._();
  GoRequest createEmptyInstance() => create();
  static $pb.PbList<GoRequest> createRepeated() => $pb.PbList<GoRequest>();
  @$core.pragma('dart2js:noInline')
  static GoRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GoRequest>(create);
  static GoRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get traceId => $_getSZ(0);
  @$pb.TagNumber(1)
  set traceId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTraceId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTraceId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get method => $_getSZ(1);
  @$pb.TagNumber(2)
  set method($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMethod() => $_has(1);
  @$pb.TagNumber(2)
  void clearMethod() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<$core.int> get data => $_getN(2);
  @$pb.TagNumber(3)
  set data($core.List<$core.int> v) { $_setBytes(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasData() => $_has(2);
  @$pb.TagNumber(3)
  void clearData() => clearField(3);
}

class GoResponse extends $pb.GeneratedMessage {
  factory GoResponse({
    $core.String? traceId,
    Code? code,
    $core.List<$core.int>? data,
  }) {
    final $result = create();
    if (traceId != null) {
      $result.traceId = traceId;
    }
    if (code != null) {
      $result.code = code;
    }
    if (data != null) {
      $result.data = data;
    }
    return $result;
  }
  GoResponse._() : super();
  factory GoResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GoResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GoResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'goapplib'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'traceId', protoName: 'traceId')
    ..e<Code>(2, _omitFieldNames ? '' : 'code', $pb.PbFieldType.OE, defaultOrMaker: Code.OK, valueOf: Code.valueOf, enumValues: Code.values)
    ..a<$core.List<$core.int>>(3, _omitFieldNames ? '' : 'data', $pb.PbFieldType.OY)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GoResponse clone() => GoResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GoResponse copyWith(void Function(GoResponse) updates) => super.copyWith((message) => updates(message as GoResponse)) as GoResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GoResponse create() => GoResponse._();
  GoResponse createEmptyInstance() => create();
  static $pb.PbList<GoResponse> createRepeated() => $pb.PbList<GoResponse>();
  @$core.pragma('dart2js:noInline')
  static GoResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GoResponse>(create);
  static GoResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get traceId => $_getSZ(0);
  @$pb.TagNumber(1)
  set traceId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTraceId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTraceId() => clearField(1);

  @$pb.TagNumber(2)
  Code get code => $_getN(1);
  @$pb.TagNumber(2)
  set code(Code v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<$core.int> get data => $_getN(2);
  @$pb.TagNumber(3)
  set data($core.List<$core.int> v) { $_setBytes(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasData() => $_has(2);
  @$pb.TagNumber(3)
  void clearData() => clearField(3);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
