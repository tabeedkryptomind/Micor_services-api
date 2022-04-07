from flask import Flask
from flask_restful import Api
from app.controller.user_controller import healthcheck
from app.controller.user_controller import Users
from app.controller.user_controller import Product
app = Flask(__name__)
api = Api(app)

api.add_resource(healthcheck,"/health" )
api.add_resource(Users, "/users")
api.add_resource(Product, "/product")
if __name__ == '__main__':
    app.run(host ='0.0.0.0',debug=True)