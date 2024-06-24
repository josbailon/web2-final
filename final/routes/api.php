<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\ClienteController;
use App\Http\Controllers\VehiculoController;
use App\Http\Controllers\PagoController;
use App\Http\Controllers\TarifaController;



Route::get('clientes', [ClienteController::class, 'index']);
Route::post('clientes', [ClienteController::class, 'store']);
Route::get('clientes/{id}', [ClienteController::class, 'show']);
Route::put('clientes/{id}', [ClienteController::class, 'update']);
Route::patch('clientes/{id}', [ClienteController::class, 'update']);
Route::delete('clientes/{id}', [ClienteController::class, 'destroy']);
//rutas de vehiculos

Route::get('vehiculos', [VehiculoController::class, 'index']);
Route::post('vehiculos', [VehiculoController::class, 'store']);
Route::get('vehiculos/{id}', [VehiculoController::class, 'show']);
Route::put('vehiculos/{id}', [VehiculoController::class, 'update']);
Route::patch('vehiculos/{id}', [VehiculoController::class, 'update']);
Route::delete('vehiculos/{id}', [VehiculoController::class, 'destroy']);
//rutas de pagos

Route::get('pagos', [PagoController::class, 'index']);
Route::post('pagos', [PagoController::class, 'store']);
Route::get('pagos/{id}', [PagoController::class, 'show']);
Route::put('pagos/{id}', [PagoController::class, 'update']);
Route::patch('pagos/{id}', [PagoController::class, 'update']);
Route::delete('pagos/{id}', [PagoController::class, 'destroy']);

//rutas de tarifas

Route::get('tarifas-estacionamiento', [TarifaController::class, 'index']);
Route::post('tarifas-estacionamiento', [TarifaController::class, 'store']);
Route::get('tarifas-estacionamiento/{id}', [TarifaController::class, 'show']);
Route::put('tarifas-estacionamiento/{id}', [TarifaController::class, 'update']);
Route::patch('tarifas-estacionamiento/{id}', [TarifaController::class, 'update']);
Route::delete('tarifas-estacionamiento/{id}', [TarifaController::class, 'destroy']);
