from flask import request, jsonify, Blueprint
from flask_jwt_extended import jwt_required
from services.ingreso_vehiculo_service import IngresoVehiculoService

ingreso_vehiculo_bp = Blueprint('ingreso_vehiculo_bp', __name__)

@ingreso_vehiculo_bp.route('/ingresos-vehiculos', methods=['POST'])
@jwt_required()
def registrar_ingreso_vehiculo():
    data = request.json
    ingreso_vehiculo = IngresoVehiculoService.registrar_ingreso_vehiculo(data['vehiculo_id'], data['fecha_ingreso'])
    return jsonify({'id': ingreso_vehiculo.id, 'vehiculo_id': ingreso_vehiculo.vehiculo_id, 'fecha_ingreso': ingreso_vehiculo.fecha_ingreso}), 201

@ingreso_vehiculo_bp.route('/ingresos-vehiculos', methods=['GET'])
@jwt_required()
def obtener_ingresos_vehiculos():
    ingresos_vehiculos = IngresoVehiculoService.get_all_ingresos_vehiculos()
    return jsonify([{'id': ingreso_vehiculo.id, 'vehiculo_id': ingreso_vehiculo.vehiculo_id, 'fecha_ingreso': ingreso_vehiculo.fecha_ingreso} for ingreso_vehiculo in ingresos_vehiculos])
