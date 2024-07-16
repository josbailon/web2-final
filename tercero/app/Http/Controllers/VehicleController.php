<?php

namespace App\Http\Controllers;

use App\Models\Vehicle;
use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class VehicleController extends Controller
{
    public function index()
    {
        return Vehicle::all();
    }

    public function store(Request $request)
    {
        $validated = $request->validate([
            'license_plate' => 'required|string|max:10|unique:vehicles,license_plate',
            'model' => 'required|string|max:255',
            'client_id' => 'required|exists:clients,id',
        ]);

        return Vehicle::create($validated);
    }

    public function show(Vehicle $vehicle)
    {
        return $vehicle;
    }

    public function update(Request $request, Vehicle $vehicle)
    {
        $validated = $request->validate([
            'license_plate' => 'string|max:10|unique:vehicles,license_plate,' . $vehicle->id,
            'model' => 'string|max:255',
            'client_id' => 'exists:clients,id',
        ]);

        $vehicle->update($validated);
        return $vehicle;
    }

    public function destroy(Vehicle $vehicle)
    {
        $vehicle->delete();
        return response()->noContent();
    }
}