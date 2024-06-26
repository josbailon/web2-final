<!DOCTYPE html>
<html>
<head>
    <title>Reporte de Tarifas de Estacionamiento</title>
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
    <h1>Reporte de Tarifas de Estacionamiento</h1>
    <table>
        <thead>
            <tr>
                <th>ID</th>
                <th>Tarifa</th>
                <th>Fecha de Creación</th>
                <th>Fecha de Actualización</th>
            </tr>
        </thead>
        <tbody>
            @foreach ($parkingFees as $fee)
                <tr>
                    <td>{{ $fee->id }}</td>
                    <td>{{ $fee->fee }}</td>
                    <td>{{ $fee->created_at }}</td>
                    <td>{{ $fee->updated_at }}</td>
                </tr>
            @endforeach
        </tbody>
    </table>
</body>
</html>
