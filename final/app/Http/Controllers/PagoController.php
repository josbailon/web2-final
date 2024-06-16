<?php

namespace App\Http\Controllers;

use App\Models\Pago;
use Illuminate\Http\Request;

class PagoController extends Controller
{
    public function index()
    {
        return response()->json(Pago::all());
    }

    public function store(Request $request)
    {
        $request->validate([
            'cliente_id' => 'required|exists:clientes,id',
            'monto' => 'required|numeric',
            'fecha_pago' => 'required|date',
        ]);

        $pago = Pago::create($request->all());

        return response()->json($pago, 201);
    }

    public function show($id)
    {
        $pago = Pago::find($id);

        if (is_null($pago)) {
            return response()->json(['message' => 'Pago no encontrado'], 404);
        }

        return response()->json($pago);
    }

    public function update(Request $request, $id)
    {
        $pago = Pago::find($id);

        if (is_null($pago)) {
            return response()->json(['message' => 'Pago no encontrado'], 404);
        }

        $request->validate([
            'cliente_id' => 'sometimes|required|exists:clientes,id',
            'monto' => 'sometimes|required|numeric',
            'fecha_pago' => 'sometimes|required|date',
        ]);

        $pago->update($request->all());

        return response()->json($pago);
    }

    public function destroy($id)
    {
        $pago = Pago::find($id);

        if (is_null($pago)) {
            return response()->json(['message' => 'Pago no encontrado'], 404);
        }

        $pago->delete();

        return response()->json(['message' => 'Pago eliminado']);
    }
}
