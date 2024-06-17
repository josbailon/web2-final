from flask import request, jsonify
from app import app, db
from models import Cliente, Vehiculo, IngresoVehiculo

@app.route('/ingreso_vehiculo', methods=['POST'])
def registrar_ingreso():
    data = request.json
    placa = data['placa']
    espacio_asignado = data['espacio_asignado']
    
    vehiculo = Vehiculo.query.filter_by(placa=placa).first()
    if vehiculo is None:
        return jsonify({'error': 'Vehículo no encontrado'}), 404

    if not verificar_disponibilidad_espacio(espacio_asignado):
        return jsonify({'error': 'Espacio no disponible'}), 400

    nuevo_ingreso = IngresoVehiculo(
        vehiculo_id=vehiculo.id,
        espacio_asignado=espacio_asignado
    )
    db.session.add(nuevo_ingreso)
    db.session.commit()

    return jsonify({'message': 'Ingreso registrado exitosamente'})

def verificar_disponibilidad_espacio(espacio):
    ocupados = IngresoVehiculo.query.filter_by(espacio_asignado=espacio).count()
    return ocupados == 0
