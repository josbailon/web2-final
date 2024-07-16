<!DOCTYPE html>
<html>
<head>
    <title>Reporte de Vehículos</title>
    <style>
        table {
            width: 100%;
            border-collapse: collapse;
        }
        table, th, td {
            border: 1px solid black;
        }
        th, td {
            padding: 8px;
            text-align: left;
        }
    </style>
</head>
<body>
    <h1>Reporte de Vehículos</h1>
    <table>
        <thead>
            <tr>
                <th>ID</th>
                <th>Matrícula</th>
                <th>Modelo</th>
                <th>ID del Cliente</th>
            </tr>
        </thead>
        <tbody>
            @foreach ($vehicles as $vehicle)
                <tr>
                    <td>{{ $vehicle->id }}</td>
                    <td>{{ $vehicle->license_plate }}</td>
                    <td>{{ $vehicle->model }}</td>
                    <td>{{ $vehicle->client_id }}</td>
                </tr>
            @endforeach
        </tbody>
    </table>
</body>
</html>
