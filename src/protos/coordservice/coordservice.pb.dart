///
//  Generated code. Do not modify.
//  source: coordservice.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'coordservice.pbenum.dart';

export 'coordservice.pbenum.dart';

class WriteCoordsReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'WriteCoordsReq', createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..a<$core.double>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'lat', $pb.PbFieldType.OD)
    ..a<$core.double>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'long', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  WriteCoordsReq._() : super();
  factory WriteCoordsReq({
    $core.String? id,
    $core.double? lat,
    $core.double? long,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (lat != null) {
      _result.lat = lat;
    }
    if (long != null) {
      _result.long = long;
    }
    return _result;
  }
  factory WriteCoordsReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory WriteCoordsReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  WriteCoordsReq clone() => WriteCoordsReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  WriteCoordsReq copyWith(void Function(WriteCoordsReq) updates) => super.copyWith((message) => updates(message as WriteCoordsReq)) as WriteCoordsReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static WriteCoordsReq create() => WriteCoordsReq._();
  WriteCoordsReq createEmptyInstance() => create();
  static $pb.PbList<WriteCoordsReq> createRepeated() => $pb.PbList<WriteCoordsReq>();
  @$core.pragma('dart2js:noInline')
  static WriteCoordsReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<WriteCoordsReq>(create);
  static WriteCoordsReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get lat => $_getN(1);
  @$pb.TagNumber(2)
  set lat($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLat() => $_has(1);
  @$pb.TagNumber(2)
  void clearLat() => clearField(2);

  @$pb.TagNumber(3)
  $core.double get long => $_getN(2);
  @$pb.TagNumber(3)
  set long($core.double v) { $_setDouble(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLong() => $_has(2);
  @$pb.TagNumber(3)
  void clearLong() => clearField(3);
}

class Empty extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Empty', createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  Empty._() : super();
  factory Empty() => create();
  factory Empty.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Empty.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Empty clone() => Empty()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Empty copyWith(void Function(Empty) updates) => super.copyWith((message) => updates(message as Empty)) as Empty; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Empty create() => Empty._();
  Empty createEmptyInstance() => create();
  static $pb.PbList<Empty> createRepeated() => $pb.PbList<Empty>();
  @$core.pragma('dart2js:noInline')
  static Empty getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Empty>(create);
  static Empty? _defaultInstance;
}

class WriteCoordsResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'WriteCoordsResp', createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  WriteCoordsResp._() : super();
  factory WriteCoordsResp() => create();
  factory WriteCoordsResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory WriteCoordsResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  WriteCoordsResp clone() => WriteCoordsResp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  WriteCoordsResp copyWith(void Function(WriteCoordsResp) updates) => super.copyWith((message) => updates(message as WriteCoordsResp)) as WriteCoordsResp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static WriteCoordsResp create() => WriteCoordsResp._();
  WriteCoordsResp createEmptyInstance() => create();
  static $pb.PbList<WriteCoordsResp> createRepeated() => $pb.PbList<WriteCoordsResp>();
  @$core.pragma('dart2js:noInline')
  static WriteCoordsResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<WriteCoordsResp>(create);
  static WriteCoordsResp? _defaultInstance;
}

class GetCoordsReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetCoordsReq', createEmptyInstance: create)
    ..pPS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ids')
    ..hasRequiredFields = false
  ;

  GetCoordsReq._() : super();
  factory GetCoordsReq({
    $core.Iterable<$core.String>? ids,
  }) {
    final _result = create();
    if (ids != null) {
      _result.ids.addAll(ids);
    }
    return _result;
  }
  factory GetCoordsReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetCoordsReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetCoordsReq clone() => GetCoordsReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetCoordsReq copyWith(void Function(GetCoordsReq) updates) => super.copyWith((message) => updates(message as GetCoordsReq)) as GetCoordsReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetCoordsReq create() => GetCoordsReq._();
  GetCoordsReq createEmptyInstance() => create();
  static $pb.PbList<GetCoordsReq> createRepeated() => $pb.PbList<GetCoordsReq>();
  @$core.pragma('dart2js:noInline')
  static GetCoordsReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetCoordsReq>(create);
  static GetCoordsReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.String> get ids => $_getList(0);
}

class GetCoordsResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetCoordsResp', createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..a<$core.double>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'lat', $pb.PbFieldType.OD)
    ..a<$core.double>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'long', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  GetCoordsResp._() : super();
  factory GetCoordsResp({
    $core.String? id,
    $core.double? lat,
    $core.double? long,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (lat != null) {
      _result.lat = lat;
    }
    if (long != null) {
      _result.long = long;
    }
    return _result;
  }
  factory GetCoordsResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetCoordsResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetCoordsResp clone() => GetCoordsResp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetCoordsResp copyWith(void Function(GetCoordsResp) updates) => super.copyWith((message) => updates(message as GetCoordsResp)) as GetCoordsResp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetCoordsResp create() => GetCoordsResp._();
  GetCoordsResp createEmptyInstance() => create();
  static $pb.PbList<GetCoordsResp> createRepeated() => $pb.PbList<GetCoordsResp>();
  @$core.pragma('dart2js:noInline')
  static GetCoordsResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetCoordsResp>(create);
  static GetCoordsResp? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get lat => $_getN(1);
  @$pb.TagNumber(2)
  set lat($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLat() => $_has(1);
  @$pb.TagNumber(2)
  void clearLat() => clearField(2);

  @$pb.TagNumber(3)
  $core.double get long => $_getN(2);
  @$pb.TagNumber(3)
  set long($core.double v) { $_setDouble(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLong() => $_has(2);
  @$pb.TagNumber(3)
  void clearLong() => clearField(3);
}

class UpdateUnitsReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateUnitsReq', createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  UpdateUnitsReq._() : super();
  factory UpdateUnitsReq() => create();
  factory UpdateUnitsReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateUnitsReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateUnitsReq clone() => UpdateUnitsReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateUnitsReq copyWith(void Function(UpdateUnitsReq) updates) => super.copyWith((message) => updates(message as UpdateUnitsReq)) as UpdateUnitsReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateUnitsReq create() => UpdateUnitsReq._();
  UpdateUnitsReq createEmptyInstance() => create();
  static $pb.PbList<UpdateUnitsReq> createRepeated() => $pb.PbList<UpdateUnitsReq>();
  @$core.pragma('dart2js:noInline')
  static UpdateUnitsReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateUnitsReq>(create);
  static UpdateUnitsReq? _defaultInstance;
}

class UpdateUnitsResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateUnitsResp', createEmptyInstance: create)
    ..pc<Coordinates>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'coords', $pb.PbFieldType.PM, subBuilder: Coordinates.create)
    ..hasRequiredFields = false
  ;

  UpdateUnitsResp._() : super();
  factory UpdateUnitsResp({
    $core.Iterable<Coordinates>? coords,
  }) {
    final _result = create();
    if (coords != null) {
      _result.coords.addAll(coords);
    }
    return _result;
  }
  factory UpdateUnitsResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateUnitsResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateUnitsResp clone() => UpdateUnitsResp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateUnitsResp copyWith(void Function(UpdateUnitsResp) updates) => super.copyWith((message) => updates(message as UpdateUnitsResp)) as UpdateUnitsResp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateUnitsResp create() => UpdateUnitsResp._();
  UpdateUnitsResp createEmptyInstance() => create();
  static $pb.PbList<UpdateUnitsResp> createRepeated() => $pb.PbList<UpdateUnitsResp>();
  @$core.pragma('dart2js:noInline')
  static UpdateUnitsResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateUnitsResp>(create);
  static UpdateUnitsResp? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Coordinates> get coords => $_getList(0);
}

class Coordinates extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Coordinates', createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'dt')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..e<Resuource>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: Resuource.TRUCK, valueOf: Resuource.valueOf, enumValues: Resuource.values)
    ..a<$core.double>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'lat', $pb.PbFieldType.OD)
    ..a<$core.double>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'long', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  Coordinates._() : super();
  factory Coordinates({
    $core.String? dt,
    $core.String? id,
    Resuource? type,
    $core.double? lat,
    $core.double? long,
  }) {
    final _result = create();
    if (dt != null) {
      _result.dt = dt;
    }
    if (id != null) {
      _result.id = id;
    }
    if (type != null) {
      _result.type = type;
    }
    if (lat != null) {
      _result.lat = lat;
    }
    if (long != null) {
      _result.long = long;
    }
    return _result;
  }
  factory Coordinates.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Coordinates.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Coordinates clone() => Coordinates()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Coordinates copyWith(void Function(Coordinates) updates) => super.copyWith((message) => updates(message as Coordinates)) as Coordinates; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Coordinates create() => Coordinates._();
  Coordinates createEmptyInstance() => create();
  static $pb.PbList<Coordinates> createRepeated() => $pb.PbList<Coordinates>();
  @$core.pragma('dart2js:noInline')
  static Coordinates getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Coordinates>(create);
  static Coordinates? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get dt => $_getSZ(0);
  @$pb.TagNumber(1)
  set dt($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDt() => $_has(0);
  @$pb.TagNumber(1)
  void clearDt() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get id => $_getSZ(1);
  @$pb.TagNumber(2)
  set id($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasId() => $_has(1);
  @$pb.TagNumber(2)
  void clearId() => clearField(2);

  @$pb.TagNumber(3)
  Resuource get type => $_getN(2);
  @$pb.TagNumber(3)
  set type(Resuource v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasType() => $_has(2);
  @$pb.TagNumber(3)
  void clearType() => clearField(3);

  @$pb.TagNumber(4)
  $core.double get lat => $_getN(3);
  @$pb.TagNumber(4)
  set lat($core.double v) { $_setDouble(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasLat() => $_has(3);
  @$pb.TagNumber(4)
  void clearLat() => clearField(4);

  @$pb.TagNumber(5)
  $core.double get long => $_getN(4);
  @$pb.TagNumber(5)
  set long($core.double v) { $_setDouble(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasLong() => $_has(4);
  @$pb.TagNumber(5)
  void clearLong() => clearField(5);
}

class Units extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Units', createEmptyInstance: create)
    ..pPS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ids')
    ..hasRequiredFields = false
  ;

  Units._() : super();
  factory Units({
    $core.Iterable<$core.String>? ids,
  }) {
    final _result = create();
    if (ids != null) {
      _result.ids.addAll(ids);
    }
    return _result;
  }
  factory Units.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Units.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Units clone() => Units()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Units copyWith(void Function(Units) updates) => super.copyWith((message) => updates(message as Units)) as Units; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Units create() => Units._();
  Units createEmptyInstance() => create();
  static $pb.PbList<Units> createRepeated() => $pb.PbList<Units>();
  @$core.pragma('dart2js:noInline')
  static Units getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Units>(create);
  static Units? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.String> get ids => $_getList(0);
}

class Coords extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Coords', createEmptyInstance: create)
    ..a<$core.double>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'lat', $pb.PbFieldType.OD)
    ..a<$core.double>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'long', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  Coords._() : super();
  factory Coords({
    $core.double? lat,
    $core.double? long,
  }) {
    final _result = create();
    if (lat != null) {
      _result.lat = lat;
    }
    if (long != null) {
      _result.long = long;
    }
    return _result;
  }
  factory Coords.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Coords.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Coords clone() => Coords()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Coords copyWith(void Function(Coords) updates) => super.copyWith((message) => updates(message as Coords)) as Coords; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Coords create() => Coords._();
  Coords createEmptyInstance() => create();
  static $pb.PbList<Coords> createRepeated() => $pb.PbList<Coords>();
  @$core.pragma('dart2js:noInline')
  static Coords getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Coords>(create);
  static Coords? _defaultInstance;

  @$pb.TagNumber(1)
  $core.double get lat => $_getN(0);
  @$pb.TagNumber(1)
  set lat($core.double v) { $_setDouble(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasLat() => $_has(0);
  @$pb.TagNumber(1)
  void clearLat() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get long => $_getN(1);
  @$pb.TagNumber(2)
  set long($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLong() => $_has(1);
  @$pb.TagNumber(2)
  void clearLong() => clearField(2);
}

class OpeataionOn extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'OpeataionOn', createEmptyInstance: create)
    ..aOM<Units>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'units', subBuilder: Units.create)
    ..aOM<Coords>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'coords', subBuilder: Coords.create)
    ..hasRequiredFields = false
  ;

  OpeataionOn._() : super();
  factory OpeataionOn({
    Units? units,
    Coords? coords,
  }) {
    final _result = create();
    if (units != null) {
      _result.units = units;
    }
    if (coords != null) {
      _result.coords = coords;
    }
    return _result;
  }
  factory OpeataionOn.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory OpeataionOn.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  OpeataionOn clone() => OpeataionOn()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  OpeataionOn copyWith(void Function(OpeataionOn) updates) => super.copyWith((message) => updates(message as OpeataionOn)) as OpeataionOn; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static OpeataionOn create() => OpeataionOn._();
  OpeataionOn createEmptyInstance() => create();
  static $pb.PbList<OpeataionOn> createRepeated() => $pb.PbList<OpeataionOn>();
  @$core.pragma('dart2js:noInline')
  static OpeataionOn getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<OpeataionOn>(create);
  static OpeataionOn? _defaultInstance;

  @$pb.TagNumber(1)
  Units get units => $_getN(0);
  @$pb.TagNumber(1)
  set units(Units v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasUnits() => $_has(0);
  @$pb.TagNumber(1)
  void clearUnits() => clearField(1);
  @$pb.TagNumber(1)
  Units ensureUnits() => $_ensure(0);

  @$pb.TagNumber(2)
  Coords get coords => $_getN(1);
  @$pb.TagNumber(2)
  set coords(Coords v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasCoords() => $_has(1);
  @$pb.TagNumber(2)
  void clearCoords() => clearField(2);
  @$pb.TagNumber(2)
  Coords ensureCoords() => $_ensure(1);
}

class OperationFromTo extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'OperationFromTo', createEmptyInstance: create)
    ..aOM<Units>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'units', subBuilder: Units.create)
    ..aOM<Coords>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'from', subBuilder: Coords.create)
    ..aOM<Coords>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'to', subBuilder: Coords.create)
    ..hasRequiredFields = false
  ;

  OperationFromTo._() : super();
  factory OperationFromTo({
    Units? units,
    Coords? from,
    Coords? to,
  }) {
    final _result = create();
    if (units != null) {
      _result.units = units;
    }
    if (from != null) {
      _result.from = from;
    }
    if (to != null) {
      _result.to = to;
    }
    return _result;
  }
  factory OperationFromTo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory OperationFromTo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  OperationFromTo clone() => OperationFromTo()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  OperationFromTo copyWith(void Function(OperationFromTo) updates) => super.copyWith((message) => updates(message as OperationFromTo)) as OperationFromTo; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static OperationFromTo create() => OperationFromTo._();
  OperationFromTo createEmptyInstance() => create();
  static $pb.PbList<OperationFromTo> createRepeated() => $pb.PbList<OperationFromTo>();
  @$core.pragma('dart2js:noInline')
  static OperationFromTo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<OperationFromTo>(create);
  static OperationFromTo? _defaultInstance;

  @$pb.TagNumber(1)
  Units get units => $_getN(0);
  @$pb.TagNumber(1)
  set units(Units v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasUnits() => $_has(0);
  @$pb.TagNumber(1)
  void clearUnits() => clearField(1);
  @$pb.TagNumber(1)
  Units ensureUnits() => $_ensure(0);

  @$pb.TagNumber(2)
  Coords get from => $_getN(1);
  @$pb.TagNumber(2)
  set from(Coords v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasFrom() => $_has(1);
  @$pb.TagNumber(2)
  void clearFrom() => clearField(2);
  @$pb.TagNumber(2)
  Coords ensureFrom() => $_ensure(1);

  @$pb.TagNumber(3)
  Coords get to => $_getN(2);
  @$pb.TagNumber(3)
  set to(Coords v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasTo() => $_has(2);
  @$pb.TagNumber(3)
  void clearTo() => clearField(3);
  @$pb.TagNumber(3)
  Coords ensureTo() => $_ensure(2);
}

class Unit extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Unit', createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  Unit._() : super();
  factory Unit({
    $core.String? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory Unit.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Unit.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Unit clone() => Unit()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Unit copyWith(void Function(Unit) updates) => super.copyWith((message) => updates(message as Unit)) as Unit; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Unit create() => Unit._();
  Unit createEmptyInstance() => create();
  static $pb.PbList<Unit> createRepeated() => $pb.PbList<Unit>();
  @$core.pragma('dart2js:noInline')
  static Unit getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Unit>(create);
  static Unit? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class Operaions extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Operaions', createEmptyInstance: create)
    ..aOM<Coords>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'from', subBuilder: Coords.create)
    ..aOM<Coords>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'to', subBuilder: Coords.create)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'operationName')
    ..hasRequiredFields = false
  ;

  Operaions._() : super();
  factory Operaions({
    Coords? from,
    Coords? to,
    $core.String? operationName,
  }) {
    final _result = create();
    if (from != null) {
      _result.from = from;
    }
    if (to != null) {
      _result.to = to;
    }
    if (operationName != null) {
      _result.operationName = operationName;
    }
    return _result;
  }
  factory Operaions.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Operaions.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Operaions clone() => Operaions()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Operaions copyWith(void Function(Operaions) updates) => super.copyWith((message) => updates(message as Operaions)) as Operaions; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Operaions create() => Operaions._();
  Operaions createEmptyInstance() => create();
  static $pb.PbList<Operaions> createRepeated() => $pb.PbList<Operaions>();
  @$core.pragma('dart2js:noInline')
  static Operaions getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Operaions>(create);
  static Operaions? _defaultInstance;

  @$pb.TagNumber(1)
  Coords get from => $_getN(0);
  @$pb.TagNumber(1)
  set from(Coords v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasFrom() => $_has(0);
  @$pb.TagNumber(1)
  void clearFrom() => clearField(1);
  @$pb.TagNumber(1)
  Coords ensureFrom() => $_ensure(0);

  @$pb.TagNumber(2)
  Coords get to => $_getN(1);
  @$pb.TagNumber(2)
  set to(Coords v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasTo() => $_has(1);
  @$pb.TagNumber(2)
  void clearTo() => clearField(2);
  @$pb.TagNumber(2)
  Coords ensureTo() => $_ensure(1);

  @$pb.TagNumber(3)
  $core.String get operationName => $_getSZ(2);
  @$pb.TagNumber(3)
  set operationName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasOperationName() => $_has(2);
  @$pb.TagNumber(3)
  void clearOperationName() => clearField(3);
}

class ID extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ID', createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  ID._() : super();
  factory ID({
    $core.String? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory ID.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ID.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ID clone() => ID()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ID copyWith(void Function(ID) updates) => super.copyWith((message) => updates(message as ID)) as ID; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ID create() => ID._();
  ID createEmptyInstance() => create();
  static $pb.PbList<ID> createRepeated() => $pb.PbList<ID>();
  @$core.pragma('dart2js:noInline')
  static ID getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ID>(create);
  static ID? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

