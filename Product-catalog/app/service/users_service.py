from app.models.user_models import User

user_list = []

def get_users():
    return user_list


def get_user(id):
    for usr in user_list:
        if id ==  usr.id:
            return user_list[id]

def create_user(user):
    user_list.append(user)
    return True
    
    