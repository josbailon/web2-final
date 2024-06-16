<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Tarifa extends Model
{
    protected $table = 'tarifas_estacionamiento';

    protected $fillable = [
        'tipo_vehiculo',
        'precio_por_hora',
    ];
}
