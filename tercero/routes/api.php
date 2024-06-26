<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\AuthController;
use App\Http\Controllers\ClientController;
use App\Http\Controllers\VehicleController;
use App\Http\Controllers\PaymentController;
use App\Http\Controllers\ParkingFeeController;
use App\Http\Controllers\ReportController;


Route::post('register', [AuthController::class, 'register']);
Route::post('login', [AuthController::class, 'login']);

Route::middleware('auth:sanctum')->group(function () {
    Route::post('logout', [AuthController::class, 'logout']);

    // Rutas API RESTful para Client
    Route::apiResource('clients', ClientController::class);

    // Rutas API RESTful para Vehicle
    Route::apiResource('vehicles', VehicleController::class);

    // Rutas API RESTful para Payment
    Route::apiResource('payments', PaymentController::class);

    // Rutas API RESTful para ParkingFee
    Route::apiResource('parkingfees', ParkingFeeController::class);
    Route::get('/reports/clients/pdf', [ReportController::class, 'pdfReport']);
Route::get('/reports/clients/excel', [ReportController::class, 'excelReport']);
});

