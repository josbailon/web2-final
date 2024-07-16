<?php

namespace App\Http\Controllers;

use App\Models\ParkingFee;
use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class ParkingFeeController extends Controller
{
    public function index()
    {
        return ParkingFee::all();
    }

    public function store(Request $request)
    {
        $validated = $request->validate([
            'fee' => 'required|numeric|min:0',
        ]);

        return ParkingFee::create($validated);
    }

    public function show(ParkingFee $parkingFee)
    {
        return $parkingFee;
    }

    public function update(Request $request, ParkingFee $parkingFee)
    {
        $validated = $request->validate([
            'fee' => 'numeric|min:0',
        ]);

        $parkingFee->update($validated);
        return $parkingFee;
    }

    public function destroy(ParkingFee $parkingFee)
    {
        $parkingFee->delete();
        return response()->noContent();
    }
}