<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Vehiculo extends Model
{
    protected $table = 'vehiculos';

    protected $fillable = [
        'cliente_id',
        'placa',
        'marca',
        'modelo',
    ];

    public function cliente()
    {
        return $this->belongsTo(Cliente::class);
    }
}
