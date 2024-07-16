<?php

namespace App\Http\Controllers;

use App\Models\Client;
use App\Models\Vehicle;
use App\Models\Payment;
use App\Models\ParkingFee;
use Barryvdh\DomPDF\Facade\Pdf;
use Illuminate\Http\Request;
use Maatwebsite\Excel\Facades\Excel;
use App\Exports\ClientsExport;
use App\Exports\VehiclesExport;
use App\Exports\PaymentsExport;
use App\Exports\ParkingFeesExport;

class ReportController extends Controller
{
    public function pdfReportClients()
    {
        $clients = Client::all();
        $pdf = Pdf::loadView('reports.clients', compact('clients'));

        return $pdf->download('clients_report.pdf');
    }

    public function excelReportClients()
    {
        return Excel::download(new ClientsExport, 'clients.xlsx');
    }

    public function pdfReportVehicles()
    {
        $vehicles = Vehicle::all();
        $pdf = Pdf::loadView('reports.vehicles', compact('vehicles'));

        return $pdf->download('vehicles_report.pdf');
    }

    public function excelReportVehicles()
    {
        return Excel::download(new VehiclesExport, 'vehicles.xlsx');
    }

    public function pdfReportPayments()
    {
        $payments = Payment::all();
        $pdf = Pdf::loadView('reports.payments', compact('payments'));

        return $pdf->download('payments_report.pdf');
    }

    public function excelReportPayments()
    {
        return Excel::download(new PaymentsExport, 'payments.xlsx');
    }

    public function pdfReportParkingFees()
    {
        $parkingFees = ParkingFee::all();
        $pdf = Pdf::loadView('reports.parking_fees', compact('parkingFees'));

        return $pdf->download('parking_fees_report.pdf');
    }

    public function excelReportParkingFees()
    {
        return Excel::download(new ParkingFeesExport, 'parking_fees.xlsx');
    }
}
