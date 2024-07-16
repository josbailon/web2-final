<?php

namespace App\Http\Controllers;

use App\Models\Tarifa;
use Illuminate\Http\Request;

class TarifaController extends Controller
{
    public function index()
    {
        return response()->json(Tarifa::all());
    }

    public function store(Request $request)
    {
        $request->validate([
            'tipo_vehiculo' => 'required|string|max:255',
            'precio_por_hora' => 'required|numeric',
        ]);

        $tarifa = Tarifa::create($request->all());

        return response()->json($tarifa, 201);
    }

    public function show($id)
    {
        $tarifa = Tarifa::find($id);

        if (is_null($tarifa)) {
            return response()->json(['message' => 'Tarifa no encontrada'], 404);
        }

        return response()->json($tarifa);
    }

    public function update(Request $request, $id)
    {
        $tarifa = Tarifa::find($id);

        if (is_null($tarifa)) {
            return response()->json(['message' => 'Tarifa no encontrada'], 404);
        }

        $request->validate([
            'tipo_vehiculo' => 'sometimes|required|string|max:255',
            'precio_por_hora' => 'sometimes|required|numeric',
        ]);

        $tarifa->update($request->all());

        return response()->json($tarifa);
    }

    public function destroy($id)
    {
        $tarifa = Tarifa::find($id);

        if (is_null($tarifa)) {
            return response()->json(['message' => 'Tarifa no encontrada'], 404);
        }

        $tarifa->delete();

        return response()->json(['message' => 'Tarifa eliminada']);
    }
}
