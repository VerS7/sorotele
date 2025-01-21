"""
Тестовое создание платежа
"""

import requests


admin_token = "ТОКЕН_АДМИНА_ТУТ"  # Токен доступа пользователя-админа
api_url = "<URL_ТУТ>/api/payment/charge"  # URL для запросов на списание
headers = {
    "Content-Type": "application/json",
    "Authentication-Token": admin_token,
}

data = {
    "account": "<ЛИЦЕВОЙ_СЧЁТ>",  # Лицевой счёт пользователя
    "amount": 0,
}

response = requests.post(api_url, headers=headers, json=data)

if response.status_code == 200:
    print("Списание принято!")
    cr = response.json()
    print(
        f"{cr['status']=}\n{cr['account']=}\n{cr['balance']=}\n{cr['amount']=}\{cr['datetime']=}"
    )
else:
    print("Ошибка: ", response.status_code, response.text)
