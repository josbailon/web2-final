from flask import request, jsonify, Blueprint
from flask_jwt_extended import jwt_required
from services.vehiculo_service import VehiculoService

vehiculo_bp = Blueprint('vehiculo_bp', __name__)

@vehiculo_bp.route('/vehiculos', methods=['POST'])
@jwt_required()
def crear_vehiculo():
    data = request.json
    vehiculo = VehiculoService.create_vehiculo(data['placa'], data['marca'], data['modelo'])
    return jsonify({'id': vehiculo.id, 'placa': vehiculo.placa, 'marca': vehiculo.marca, 'modelo': vehiculo.modelo}), 201

@vehiculo_bp.route('/vehiculos', methods=['GET'])
@jwt_required()
def obtener_vehiculos():
    vehiculos = VehiculoService.get_all_vehiculos()
    return jsonify([{'id': vehiculo.id, 'placa': vehiculo.placa, 'marca': vehiculo.marca, 'modelo': vehiculo.modelo} for vehiculo in vehiculos])
