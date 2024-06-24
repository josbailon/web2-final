from flask import request, jsonify, Blueprint
from flask_jwt_extended import jwt_required, get_jwt_identity
from services.cliente_service import ClienteService

cliente_bp = Blueprint('cliente_bp', __name__)

@cliente_bp.route('/clientes', methods=['POST'])
def crear_cliente():
    data = request.json
    cliente = ClienteService.create_cliente(data['nombre'], data['email'])
    return jsonify({'id': cliente.id, 'nombre': cliente.nombre, 'email': cliente.email}), 201

@cliente_bp.route('/clientes', methods=['GET'])
@jwt_required()
def obtener_clientes():
    clientes = ClienteService.get_all_clientes()
    return jsonify([{'id': cliente.id, 'nombre': cliente.nombre, 'email': cliente.email} for cliente in clientes])
