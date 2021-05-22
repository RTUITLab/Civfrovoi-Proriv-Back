///
//  Generated code. Do not modify.
//  source: coordservice.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Resuource extends $pb.ProtobufEnum {
  static const Resuource TRUCK = Resuource._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'TRUCK');
  static const Resuource ECAVATOR = Resuource._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ECAVATOR');
  static const Resuource ROTARY_LOADER = Resuource._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ROTARY_LOADER');
  static const Resuource PLOW_AND_BRUSH_MACHINE = Resuource._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'PLOW_AND_BRUSH_MACHINE');
  static const Resuource PLOW_AND_BRUSH_MACHINE_REG = Resuource._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'PLOW_AND_BRUSH_MACHINE_REG');
  static const Resuource HUMAN = Resuource._(5, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'HUMAN');
  static const Resuource DISPETCHER = Resuource._(6, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'DISPETCHER');

  static const $core.List<Resuource> values = <Resuource> [
    TRUCK,
    ECAVATOR,
    ROTARY_LOADER,
    PLOW_AND_BRUSH_MACHINE,
    PLOW_AND_BRUSH_MACHINE_REG,
    HUMAN,
    DISPETCHER,
  ];

  static final $core.Map<$core.int, Resuource> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Resuource? valueOf($core.int value) => _byValue[value];

  const Resuource._($core.int v, $core.String n) : super(v, n);
}

