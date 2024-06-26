<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Pago extends Model
{
    protected $table = 'pagos';

    protected $fillable = [
        'cliente_id',
        'monto',
        'fecha_pago',
    ];

    protected $dates = [
        'fecha_pago',
    ];

    public function cliente()
    {
        return $this->belongsTo(Cliente::class);
    }
}
