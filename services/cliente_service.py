from repositories.cliente_repository import ClienteRepository

class ClienteService:

    @staticmethod
    def create_cliente(nombre, email):
        return ClienteRepository.create_cliente(nombre, email)

    @staticmethod
    def get_all_clientes():
        return ClienteRepository.get_all_client
