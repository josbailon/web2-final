from repositories.cliente_repository import ClienteRepository

class AuthService:

    @staticmethod
    def register(email, password):
        user = ClienteRepository.find_by_email(email)
        if user:
            return None, 'User already exists'
        new_user = ClienteRepository.create_cliente(email, password)
        return new_user, None

    @staticmethod
    def login(email, password):
        user = ClienteRepository.find_by_email(email)
        if user and user.check_password(password):
            access_token = create_access_token(identity={'email': user.email})
            return access_token, None
        return None, 'Invalid credentials'
