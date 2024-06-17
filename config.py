import os

class Config:
    SECRET_KEY = os.environ.get('SECRET_KEY') or 'a_really_long_random_string'
    SQLALCHEMY_DATABASE_URI = os.environ.get('DATABASE_URL') or 'sqlite:///parking_system.db'
    SQLALCHEMY_TRACK_MODIFICATIONS = False
