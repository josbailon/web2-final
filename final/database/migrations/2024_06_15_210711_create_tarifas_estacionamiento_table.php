
<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateTarifasEstacionamientoTable extends Migration
{
    public function up()
    {
        Schema::create('tarifas_estacionamiento', function (Blueprint $table) {
            $table->id();
            $table->string('tipo_vehiculo');
            $table->decimal('precio_por_hora');
            $table->timestamps();
        });
    }

    public function down()
    {
        Schema::dropIfExists('tarifas_estacionamiento');
    }
}
