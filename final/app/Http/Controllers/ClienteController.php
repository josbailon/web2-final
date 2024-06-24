<?php

namespace App\Http\Controllers;

use App\Models\Cliente;
use Illuminate\Http\Request;

class ClienteController extends Controller
{
    public function index()
    {
        return response()->json(Cliente::all());
    }

    public function store(Request $request)
    {
        $request->validate([
            'nombre' => 'required|string|max:255',
            'documento_identidad' => 'required|string|max:255|unique:clientes',
            'correo' => 'required|string|email|max:255|unique:clientes',
            'telefono' => 'required|string|max:20',
        ]);

        $cliente = Cliente::create($request->all());

        return response()->json($cliente, 201);
    }

    public function show($id)
    {
        $cliente = Cliente::find($id);

        if (is_null($cliente)) {
            return response()->json(['message' => 'Cliente no encontrado'], 404);
        }

        return response()->json($cliente);
    }

    public function update(Request $request, $id)
    {
        $cliente = Cliente::find($id);

        if (is_null($cliente)) {
            return response()->json(['message' => 'Cliente no encontrado'], 404);
        }

        $request->validate([
            'nombre' => 'sometimes|required|string|max:255',
            'documento_identidad' => 'sometimes|required|string|max:255|unique:clientes,documento_identidad,' . $id,
            'correo' => 'sometimes|required|string|email|max:255|unique:clientes,correo,' . $id,
            'telefono' => 'sometimes|required|string|max:20',
        ]);

        $cliente->update($request->all());

        return response()->json($cliente);
    }

    public function destroy($id)
    {
        $cliente = Cliente::find($id);

        if (is_null($cliente)) {
            return response()->json(['message' => 'Cliente no encontrado'], 404);
        }

        $cliente->delete();

        return response()->json(['message' => 'Cliente eliminado']);
    }
}
