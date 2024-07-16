<?php
namespace App\Exports;

use App\Models\Vehicle;
use Maatwebsite\Excel\Concerns\FromCollection;
use Maatwebsite\Excel\Concerns\WithHeadings;

class VehiclesExport implements FromCollection, WithHeadings
{
    public function collection()
    {
        return Vehicle::all();
    }

    public function headings(): array
    {
        return [
            'ID',
            'Matrícula',
            'Modelo',
            'ID del Cliente',
            'Fecha de Creación',
            'Fecha de Actualización',
        ];
    }
}
