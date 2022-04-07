from flask_restful import Resource
from app.service.users_service import *
from urllib.request import urlopen
import json

class healthcheck(Resource):
    def get(self):
        return {"message": "is alive"}
    
class Users(Resource):
    def get(self):
        users = get_users()
        return json.dumps(users)

class Product(Resource):
    def get(self):
        response = urlopen("http://localhost:8080/api/v1/products/")
        data_json = json.loads(response.read())
        users =  get_users()
        users_prod = {
            "user": users,
            "products": data_json
        }
        return users_prod