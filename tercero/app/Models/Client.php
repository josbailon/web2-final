<?php
namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Client extends Model
{
    use HasFactory;

    protected $fillable = ['name', 'email'];

    public function vehicles()
    {
        return $this->hasMany(Vehicle::class);
    }

    public function payments()
    {
        return $this->hasMany(Payment::class);
    }
}
