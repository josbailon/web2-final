<?php

namespace App\Http\Controllers;

use App\Models\Vehiculo;
use Illuminate\Http\Request;

class VehiculoController extends Controller
{
    public function index()
    {
        return response()->json(Vehiculo::all());
    }

    public function store(Request $request)
    {
        $request->validate([
            'cliente_id' => 'required|exists:clientes,id',
            'placa' => 'required|string|max:255|unique:vehiculos',
            'marca' => 'required|string|max:255',
            'modelo' => 'required|string|max:255',
        ]);

        $vehiculo = Vehiculo::create($request->all());

        return response()->json($vehiculo, 201);
    }

    public function show($id)
    {
        $vehiculo = Vehiculo::find($id);

        if (is_null($vehiculo)) {
            return response()->json(['message' => 'Vehiculo no encontrado'], 404);
        }

        return response()->json($vehiculo);
    }

    public function update(Request $request, $id)
    {
        $vehiculo = Vehiculo::find($id);

        if (is_null($vehiculo)) {
            return response()->json(['message' => 'Vehiculo no encontrado'], 404);
        }

        $request->validate([
            'cliente_id' => 'sometimes|required|exists:clientes,id',
            'placa' => 'sometimes|required|string|max:255|unique:vehiculos,placa,' . $id,
            'marca' => 'sometimes|required|string|max:255',
            'modelo' => 'sometimes|required|string|max:255',
        ]);

        $vehiculo->update($request->all());

        return response()->json($vehiculo);
    }

    public function destroy($id)
    {
        $vehiculo = Vehiculo::find($id);

        if (is_null($vehiculo)) {
            return response()->json(['message' => 'Vehiculo no encontrado'], 404);
        }

        $vehiculo->delete();

        return response()->json(['message' => 'Vehiculo eliminado']);
    }
}
