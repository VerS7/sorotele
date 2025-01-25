"""
Тестовое создание нового пользователя
"""

import requests


admin_token = "ТОКЕН_АДМИНА_ТУТ"  # Токен доступа пользователя-админа
api_url = "<URL_ТУТ>/api/user/create"  # URL для запросов на создание пользователя
headers = {"Content-Type": "application/json", "Authentication-Token": admin_token}
user = {
    "name": "<ИМЯ_ПОЛЬЗОВАТЕЛЯ>",
    "surname": "<ФАМИЛИЯ_ПОЛЬЗОВАТЕЛЯ>",
    "password": "<ПАРОЛЬ_ПОЛЬЗОВАТЕЛЯ>",
    "rate_id": "<ID тарифа>",
}  # Данные пользователя

response = requests.post(api_url, headers=headers, json=user)

if response.status_code == 200:
    ur = response.json()
    print("Создан новый пользователь:")
    print(f"{ur['account']=}\n{ur['name']=}\n{ur['surname']=}\n{ur['token']=}\n")
else:
    print("Ошибка: ", response.status_code, response.text)
