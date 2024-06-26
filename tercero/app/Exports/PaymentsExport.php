<?php
namespace App\Exports;

use App\Models\Payment;
use Maatwebsite\Excel\Concerns\FromCollection;
use Maatwebsite\Excel\Concerns\WithHeadings;

class PaymentsExport implements FromCollection, WithHeadings
{
    public function collection()
    {
        return Payment::all();
    }

    public function headings(): array
    {
        return [
            'ID',
            'ID del Cliente',
            'Monto',
            'Fecha de Creación',
            'Fecha de Actualización',
        ];
    }
}
