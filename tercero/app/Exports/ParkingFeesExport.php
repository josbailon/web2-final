<?php
namespace App\Exports;

use App\Models\ParkingFee;
use Maatwebsite\Excel\Concerns\FromCollection;
use Maatwebsite\Excel\Concerns\WithHeadings;

class ParkingFeesExport implements FromCollection, WithHeadings
{
    public function collection()
    {
        return ParkingFee::all();
    }

    public function headings(): array
    {
        return [
            'ID',
            'Tarifa',
            'Fecha de Creación',
            'Fecha de Actualización',
        ];
    }
}
