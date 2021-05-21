///
//  Generated code. Do not modify.
//  source: coordservice.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'coordservice.pb.dart' as $0;
export 'coordservice.pb.dart';

class CoordsServiceClient extends $grpc.Client {
  static final _$writeCoords =
      $grpc.ClientMethod<$0.WriteCoordsReq, $0.WriteCoordsResp>(
          '/CoordsService/WriteCoords',
          ($0.WriteCoordsReq value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.WriteCoordsResp.fromBuffer(value));
  static final _$getCoords = $grpc.ClientMethod<$0.Units, $0.GetCoordsResp>(
      '/CoordsService/GetCoords',
      ($0.Units value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.GetCoordsResp.fromBuffer(value));
  static final _$updateUnits =
      $grpc.ClientMethod<$0.UpdateUnitsReq, $0.UpdateUnitsResp>(
          '/CoordsService/UpdateUnits',
          ($0.UpdateUnitsReq value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.UpdateUnitsResp.fromBuffer(value));
  static final _$sweepingUp = $grpc.ClientMethod<$0.OpeataionOn, $0.Empty>(
      '/CoordsService/SweepingUp',
      ($0.OpeataionOn value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$reagentTreatment =
      $grpc.ClientMethod<$0.OpeataionOn, $0.Empty>(
          '/CoordsService/ReagentTreatment',
          ($0.OpeataionOn value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$shaftFormation = $grpc.ClientMethod<$0.OpeataionOn, $0.Empty>(
      '/CoordsService/ShaftFormation',
      ($0.OpeataionOn value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$moveSnowToTemp =
      $grpc.ClientMethod<$0.OperationFromTo, $0.Empty>(
          '/CoordsService/MoveSnowToTemp',
          ($0.OperationFromTo value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$loadingSnowFromTemp =
      $grpc.ClientMethod<$0.OpeataionOn, $0.Empty>(
          '/CoordsService/LoadingSnowFromTemp',
          ($0.OpeataionOn value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$exportSnowFromTemp =
      $grpc.ClientMethod<$0.OperationFromTo, $0.Empty>(
          '/CoordsService/ExportSnowFromTemp',
          ($0.OperationFromTo value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$clearTemp = $grpc.ClientMethod<$0.OpeataionOn, $0.Empty>(
      '/CoordsService/ClearTemp',
      ($0.OpeataionOn value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$listenCommands = $grpc.ClientMethod<$0.Unit, $0.Operaions>(
      '/CoordsService/ListenCommands',
      ($0.Unit value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Operaions.fromBuffer(value));
  static final _$initApp = $grpc.ClientMethod<$0.Empty, $0.ID>(
      '/CoordsService/InitApp',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ID.fromBuffer(value));

  CoordsServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.WriteCoordsResp> writeCoords(
      $0.WriteCoordsReq request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$writeCoords, request, options: options);
  }

  $grpc.ResponseStream<$0.GetCoordsResp> getCoords($0.Units request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$getCoords, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$0.UpdateUnitsResp> updateUnits(
      $0.UpdateUnitsReq request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$updateUnits, request, options: options);
  }

  $grpc.ResponseFuture<$0.Empty> sweepingUp($0.OpeataionOn request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$sweepingUp, request, options: options);
  }

  $grpc.ResponseFuture<$0.Empty> reagentTreatment($0.OpeataionOn request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$reagentTreatment, request, options: options);
  }

  $grpc.ResponseFuture<$0.Empty> shaftFormation($0.OpeataionOn request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$shaftFormation, request, options: options);
  }

  $grpc.ResponseFuture<$0.Empty> moveSnowToTemp($0.OperationFromTo request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$moveSnowToTemp, request, options: options);
  }

  $grpc.ResponseFuture<$0.Empty> loadingSnowFromTemp($0.OpeataionOn request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$loadingSnowFromTemp, request, options: options);
  }

  $grpc.ResponseFuture<$0.Empty> exportSnowFromTemp($0.OperationFromTo request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$exportSnowFromTemp, request, options: options);
  }

  $grpc.ResponseFuture<$0.Empty> clearTemp($0.OpeataionOn request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$clearTemp, request, options: options);
  }

  $grpc.ResponseStream<$0.Operaions> listenCommands($0.Unit request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$listenCommands, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$0.ID> initApp($0.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$initApp, request, options: options);
  }
}

abstract class CoordsServiceBase extends $grpc.Service {
  $core.String get $name => 'CoordsService';

  CoordsServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.WriteCoordsReq, $0.WriteCoordsResp>(
        'WriteCoords',
        writeCoords_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.WriteCoordsReq.fromBuffer(value),
        ($0.WriteCoordsResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Units, $0.GetCoordsResp>(
        'GetCoords',
        getCoords_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $0.Units.fromBuffer(value),
        ($0.GetCoordsResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UpdateUnitsReq, $0.UpdateUnitsResp>(
        'UpdateUnits',
        updateUnits_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UpdateUnitsReq.fromBuffer(value),
        ($0.UpdateUnitsResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.OpeataionOn, $0.Empty>(
        'SweepingUp',
        sweepingUp_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.OpeataionOn.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.OpeataionOn, $0.Empty>(
        'ReagentTreatment',
        reagentTreatment_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.OpeataionOn.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.OpeataionOn, $0.Empty>(
        'ShaftFormation',
        shaftFormation_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.OpeataionOn.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.OperationFromTo, $0.Empty>(
        'MoveSnowToTemp',
        moveSnowToTemp_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.OperationFromTo.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.OpeataionOn, $0.Empty>(
        'LoadingSnowFromTemp',
        loadingSnowFromTemp_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.OpeataionOn.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.OperationFromTo, $0.Empty>(
        'ExportSnowFromTemp',
        exportSnowFromTemp_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.OperationFromTo.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.OpeataionOn, $0.Empty>(
        'ClearTemp',
        clearTemp_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.OpeataionOn.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Unit, $0.Operaions>(
        'ListenCommands',
        listenCommands_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $0.Unit.fromBuffer(value),
        ($0.Operaions value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Empty, $0.ID>(
        'InitApp',
        initApp_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($0.ID value) => value.writeToBuffer()));
  }

  $async.Future<$0.WriteCoordsResp> writeCoords_Pre(
      $grpc.ServiceCall call, $async.Future<$0.WriteCoordsReq> request) async {
    return writeCoords(call, await request);
  }

  $async.Stream<$0.GetCoordsResp> getCoords_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Units> request) async* {
    yield* getCoords(call, await request);
  }

  $async.Future<$0.UpdateUnitsResp> updateUnits_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UpdateUnitsReq> request) async {
    return updateUnits(call, await request);
  }

  $async.Future<$0.Empty> sweepingUp_Pre(
      $grpc.ServiceCall call, $async.Future<$0.OpeataionOn> request) async {
    return sweepingUp(call, await request);
  }

  $async.Future<$0.Empty> reagentTreatment_Pre(
      $grpc.ServiceCall call, $async.Future<$0.OpeataionOn> request) async {
    return reagentTreatment(call, await request);
  }

  $async.Future<$0.Empty> shaftFormation_Pre(
      $grpc.ServiceCall call, $async.Future<$0.OpeataionOn> request) async {
    return shaftFormation(call, await request);
  }

  $async.Future<$0.Empty> moveSnowToTemp_Pre(
      $grpc.ServiceCall call, $async.Future<$0.OperationFromTo> request) async {
    return moveSnowToTemp(call, await request);
  }

  $async.Future<$0.Empty> loadingSnowFromTemp_Pre(
      $grpc.ServiceCall call, $async.Future<$0.OpeataionOn> request) async {
    return loadingSnowFromTemp(call, await request);
  }

  $async.Future<$0.Empty> exportSnowFromTemp_Pre(
      $grpc.ServiceCall call, $async.Future<$0.OperationFromTo> request) async {
    return exportSnowFromTemp(call, await request);
  }

  $async.Future<$0.Empty> clearTemp_Pre(
      $grpc.ServiceCall call, $async.Future<$0.OpeataionOn> request) async {
    return clearTemp(call, await request);
  }

  $async.Stream<$0.Operaions> listenCommands_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Unit> request) async* {
    yield* listenCommands(call, await request);
  }

  $async.Future<$0.ID> initApp_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return initApp(call, await request);
  }

  $async.Future<$0.WriteCoordsResp> writeCoords(
      $grpc.ServiceCall call, $0.WriteCoordsReq request);
  $async.Stream<$0.GetCoordsResp> getCoords(
      $grpc.ServiceCall call, $0.Units request);
  $async.Future<$0.UpdateUnitsResp> updateUnits(
      $grpc.ServiceCall call, $0.UpdateUnitsReq request);
  $async.Future<$0.Empty> sweepingUp(
      $grpc.ServiceCall call, $0.OpeataionOn request);
  $async.Future<$0.Empty> reagentTreatment(
      $grpc.ServiceCall call, $0.OpeataionOn request);
  $async.Future<$0.Empty> shaftFormation(
      $grpc.ServiceCall call, $0.OpeataionOn request);
  $async.Future<$0.Empty> moveSnowToTemp(
      $grpc.ServiceCall call, $0.OperationFromTo request);
  $async.Future<$0.Empty> loadingSnowFromTemp(
      $grpc.ServiceCall call, $0.OpeataionOn request);
  $async.Future<$0.Empty> exportSnowFromTemp(
      $grpc.ServiceCall call, $0.OperationFromTo request);
  $async.Future<$0.Empty> clearTemp(
      $grpc.ServiceCall call, $0.OpeataionOn request);
  $async.Stream<$0.Operaions> listenCommands(
      $grpc.ServiceCall call, $0.Unit request);
  $async.Future<$0.ID> initApp($grpc.ServiceCall call, $0.Empty request);
}
