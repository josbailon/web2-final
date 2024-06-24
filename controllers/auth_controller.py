from flask import request, jsonify, Blueprint
from services.auth_service import AuthService

auth_bp = Blueprint('auth_bp', __name__)

@auth_bp.route('/register', methods=['POST'])
def register():
    data = request.json
    user, error = AuthService.register(data['email'], data['password'])
    if error:
        return jsonify({'error': error}), 400
    return jsonify({'message': 'User registered successfully'}), 201

@auth_bp.route('/login', methods=['POST'])
def login():
    data = request.json
    access_token, error = AuthService.login(data['email'], data['password'])
    if error:
        return jsonify({'error': error}), 401
    return jsonify({'access_token': access_token}), 200
