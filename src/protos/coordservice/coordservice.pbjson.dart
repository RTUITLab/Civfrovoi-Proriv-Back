///
//  Generated code. Do not modify.
//  source: coordservice.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use resuourceDescriptor instead')
const Resuource$json = const {
  '1': 'Resuource',
  '2': const [
    const {'1': 'TRUCK', '2': 0},
    const {'1': 'ECAVATOR', '2': 1},
    const {'1': 'ROTARY_LOADER', '2': 2},
    const {'1': 'PLOW_AND_BRUSH_MACHINE', '2': 3},
    const {'1': 'PLOW_AND_BRUSH_MACHINE_REG', '2': 4},
    const {'1': 'HUMAN', '2': 5},
    const {'1': 'DISPETCHER', '2': 6},
  ],
};

/// Descriptor for `Resuource`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List resuourceDescriptor = $convert.base64Decode('CglSZXN1b3VyY2USCQoFVFJVQ0sQABIMCghFQ0FWQVRPUhABEhEKDVJPVEFSWV9MT0FERVIQAhIaChZQTE9XX0FORF9CUlVTSF9NQUNISU5FEAMSHgoaUExPV19BTkRfQlJVU0hfTUFDSElORV9SRUcQBBIJCgVIVU1BThAFEg4KCkRJU1BFVENIRVIQBg==');
@$core.Deprecated('Use writeCoordsReqDescriptor instead')
const WriteCoordsReq$json = const {
  '1': 'WriteCoordsReq',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'lat', '3': 2, '4': 1, '5': 1, '10': 'lat'},
    const {'1': 'long', '3': 3, '4': 1, '5': 1, '10': 'long'},
  ],
};

/// Descriptor for `WriteCoordsReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List writeCoordsReqDescriptor = $convert.base64Decode('Cg5Xcml0ZUNvb3Jkc1JlcRIOCgJpZBgBIAEoCVICaWQSEAoDbGF0GAIgASgBUgNsYXQSEgoEbG9uZxgDIAEoAVIEbG9uZw==');
@$core.Deprecated('Use emptyDescriptor instead')
const Empty$json = const {
  '1': 'Empty',
};

/// Descriptor for `Empty`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List emptyDescriptor = $convert.base64Decode('CgVFbXB0eQ==');
@$core.Deprecated('Use writeCoordsRespDescriptor instead')
const WriteCoordsResp$json = const {
  '1': 'WriteCoordsResp',
};

/// Descriptor for `WriteCoordsResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List writeCoordsRespDescriptor = $convert.base64Decode('Cg9Xcml0ZUNvb3Jkc1Jlc3A=');
@$core.Deprecated('Use getCoordsReqDescriptor instead')
const GetCoordsReq$json = const {
  '1': 'GetCoordsReq',
  '2': const [
    const {'1': 'ids', '3': 1, '4': 3, '5': 9, '10': 'ids'},
  ],
};

/// Descriptor for `GetCoordsReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getCoordsReqDescriptor = $convert.base64Decode('CgxHZXRDb29yZHNSZXESEAoDaWRzGAEgAygJUgNpZHM=');
@$core.Deprecated('Use getCoordsRespDescriptor instead')
const GetCoordsResp$json = const {
  '1': 'GetCoordsResp',
  '2': const [
    const {'1': 'coords', '3': 1, '4': 3, '5': 11, '6': '.UnitWithCords', '10': 'coords'},
  ],
};

/// Descriptor for `GetCoordsResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getCoordsRespDescriptor = $convert.base64Decode('Cg1HZXRDb29yZHNSZXNwEiYKBmNvb3JkcxgBIAMoCzIOLlVuaXRXaXRoQ29yZHNSBmNvb3Jkcw==');
@$core.Deprecated('Use updateUnitsReqDescriptor instead')
const UpdateUnitsReq$json = const {
  '1': 'UpdateUnitsReq',
};

/// Descriptor for `UpdateUnitsReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateUnitsReqDescriptor = $convert.base64Decode('Cg5VcGRhdGVVbml0c1JlcQ==');
@$core.Deprecated('Use updateUnitsRespDescriptor instead')
const UpdateUnitsResp$json = const {
  '1': 'UpdateUnitsResp',
  '2': const [
    const {'1': 'coords', '3': 1, '4': 3, '5': 11, '6': '.Coordinates', '10': 'coords'},
  ],
};

/// Descriptor for `UpdateUnitsResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateUnitsRespDescriptor = $convert.base64Decode('Cg9VcGRhdGVVbml0c1Jlc3ASJAoGY29vcmRzGAEgAygLMgwuQ29vcmRpbmF0ZXNSBmNvb3Jkcw==');
@$core.Deprecated('Use coordinatesDescriptor instead')
const Coordinates$json = const {
  '1': 'Coordinates',
  '2': const [
    const {'1': 'dt', '3': 1, '4': 1, '5': 9, '10': 'dt'},
    const {'1': 'id', '3': 2, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'type', '3': 3, '4': 1, '5': 14, '6': '.Resuource', '10': 'type'},
    const {'1': 'lat', '3': 4, '4': 1, '5': 1, '10': 'lat'},
    const {'1': 'long', '3': 5, '4': 1, '5': 1, '10': 'long'},
  ],
};

/// Descriptor for `Coordinates`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List coordinatesDescriptor = $convert.base64Decode('CgtDb29yZGluYXRlcxIOCgJkdBgBIAEoCVICZHQSDgoCaWQYAiABKAlSAmlkEh4KBHR5cGUYAyABKA4yCi5SZXN1b3VyY2VSBHR5cGUSEAoDbGF0GAQgASgBUgNsYXQSEgoEbG9uZxgFIAEoAVIEbG9uZw==');
@$core.Deprecated('Use unitsDescriptor instead')
const Units$json = const {
  '1': 'Units',
  '2': const [
    const {'1': 'ids', '3': 1, '4': 3, '5': 9, '10': 'ids'},
  ],
};

/// Descriptor for `Units`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List unitsDescriptor = $convert.base64Decode('CgVVbml0cxIQCgNpZHMYASADKAlSA2lkcw==');
@$core.Deprecated('Use coordsDescriptor instead')
const Coords$json = const {
  '1': 'Coords',
  '2': const [
    const {'1': 'lat', '3': 1, '4': 1, '5': 1, '10': 'lat'},
    const {'1': 'long', '3': 2, '4': 1, '5': 1, '10': 'long'},
  ],
};

/// Descriptor for `Coords`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List coordsDescriptor = $convert.base64Decode('CgZDb29yZHMSEAoDbGF0GAEgASgBUgNsYXQSEgoEbG9uZxgCIAEoAVIEbG9uZw==');
@$core.Deprecated('Use opeataionOnDescriptor instead')
const OpeataionOn$json = const {
  '1': 'OpeataionOn',
  '2': const [
    const {'1': 'units', '3': 1, '4': 1, '5': 11, '6': '.Units', '10': 'units'},
    const {'1': 'coords', '3': 2, '4': 1, '5': 11, '6': '.Coords', '10': 'coords'},
  ],
};

/// Descriptor for `OpeataionOn`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List opeataionOnDescriptor = $convert.base64Decode('CgtPcGVhdGFpb25PbhIcCgV1bml0cxgBIAEoCzIGLlVuaXRzUgV1bml0cxIfCgZjb29yZHMYAiABKAsyBy5Db29yZHNSBmNvb3Jkcw==');
@$core.Deprecated('Use operationFromToDescriptor instead')
const OperationFromTo$json = const {
  '1': 'OperationFromTo',
  '2': const [
    const {'1': 'units', '3': 1, '4': 1, '5': 11, '6': '.Units', '10': 'units'},
    const {'1': 'from', '3': 2, '4': 1, '5': 11, '6': '.Coords', '10': 'from'},
    const {'1': 'to', '3': 3, '4': 1, '5': 11, '6': '.Coords', '10': 'to'},
  ],
};

/// Descriptor for `OperationFromTo`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List operationFromToDescriptor = $convert.base64Decode('Cg9PcGVyYXRpb25Gcm9tVG8SHAoFdW5pdHMYASABKAsyBi5Vbml0c1IFdW5pdHMSGwoEZnJvbRgCIAEoCzIHLkNvb3Jkc1IEZnJvbRIXCgJ0bxgDIAEoCzIHLkNvb3Jkc1ICdG8=');
@$core.Deprecated('Use unitDescriptor instead')
const Unit$json = const {
  '1': 'Unit',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

/// Descriptor for `Unit`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List unitDescriptor = $convert.base64Decode('CgRVbml0Eg4KAmlkGAEgASgJUgJpZA==');
@$core.Deprecated('Use unitWithCordsDescriptor instead')
const UnitWithCords$json = const {
  '1': 'UnitWithCords',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'coords', '3': 2, '4': 1, '5': 11, '6': '.Coords', '10': 'coords'},
  ],
};

/// Descriptor for `UnitWithCords`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List unitWithCordsDescriptor = $convert.base64Decode('Cg1Vbml0V2l0aENvcmRzEg4KAmlkGAEgASgJUgJpZBIfCgZjb29yZHMYAiABKAsyBy5Db29yZHNSBmNvb3Jkcw==');
@$core.Deprecated('Use operaionsDescriptor instead')
const Operaions$json = const {
  '1': 'Operaions',
  '2': const [
    const {'1': 'from', '3': 1, '4': 1, '5': 11, '6': '.Coords', '10': 'from'},
    const {'1': 'to', '3': 2, '4': 1, '5': 11, '6': '.Coords', '10': 'to'},
    const {'1': 'operation_name', '3': 3, '4': 1, '5': 9, '10': 'operationName'},
  ],
};

/// Descriptor for `Operaions`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List operaionsDescriptor = $convert.base64Decode('CglPcGVyYWlvbnMSGwoEZnJvbRgBIAEoCzIHLkNvb3Jkc1IEZnJvbRIXCgJ0bxgCIAEoCzIHLkNvb3Jkc1ICdG8SJQoOb3BlcmF0aW9uX25hbWUYAyABKAlSDW9wZXJhdGlvbk5hbWU=');
@$core.Deprecated('Use iDDescriptor instead')
const ID$json = const {
  '1': 'ID',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

/// Descriptor for `ID`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List iDDescriptor = $convert.base64Decode('CgJJRBIOCgJpZBgBIAEoCVICaWQ=');
@$core.Deprecated('Use initReqDescriptor instead')
const InitReq$json = const {
  '1': 'InitReq',
  '2': const [
    const {'1': 'type', '3': 1, '4': 1, '5': 14, '6': '.Resuource', '10': 'type'},
  ],
};

/// Descriptor for `InitReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List initReqDescriptor = $convert.base64Decode('CgdJbml0UmVxEh4KBHR5cGUYASABKA4yCi5SZXN1b3VyY2VSBHR5cGU=');
@$core.Deprecated('Use taskDescriptor instead')
const Task$json = const {
  '1': 'Task',
  '2': const [
    const {'1': 'on', '3': 1, '4': 1, '5': 11, '6': '.Coords', '10': 'on'},
    const {'1': 'operation_name', '3': 2, '4': 1, '5': 9, '10': 'operationName'},
  ],
};

/// Descriptor for `Task`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List taskDescriptor = $convert.base64Decode('CgRUYXNrEhcKAm9uGAEgASgLMgcuQ29vcmRzUgJvbhIlCg5vcGVyYXRpb25fbmFtZRgCIAEoCVINb3BlcmF0aW9uTmFtZQ==');
@$core.Deprecated('Use tasksDescriptor instead')
const Tasks$json = const {
  '1': 'Tasks',
  '2': const [
    const {'1': 'tasks', '3': 1, '4': 3, '5': 11, '6': '.Task', '10': 'tasks'},
  ],
};

/// Descriptor for `Tasks`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List tasksDescriptor = $convert.base64Decode('CgVUYXNrcxIbCgV0YXNrcxgBIAMoCzIFLlRhc2tSBXRhc2tz');
