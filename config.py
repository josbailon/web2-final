import os

class Config:
    SECRET_KEY = os.environ.get('SECRET_KEY') or 'a_really_long_random_string'
    SQLALCHEMY_DATABASE_URI = os.environ.get('DATABASE_URL') or 'sqlite:///project.db'
    SQLALCHEMY_TRACK_MODIFICATIONS = False
    JWT_SECRET_KEY = os.environ.get('JWT_SECRET_KEY') or 'another_really_long_random_string'
