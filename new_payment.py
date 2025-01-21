"""
Тестовое создание платежа с хеш-проверкой подленности
"""

import random
import requests
import hashlib
from datetime import datetime


def sha1_hash(string):
    sha1_hash = hashlib.sha1()
    sha1_hash.update(string.encode("utf-8"))
    return sha1_hash.hexdigest()


api_url = "<URL_ТУТ>/api/payment/notification"  # URL для запросов на создание платежа
secret = "<YOOMONEY_SECRET>"  # Секретный ключ приложения для проверки подленности
headers = {"Content-Type": "application/x-www-form-urlencoded"}
account = "<ЛИЦЕВОЙ_СЧЁТ>"  # Лицевой счёт пользователя
data = {
    "notification_type": "income",
    "operation_id": random.randint(0, 100_000_000),
    "amount": "1111",
    "currency": "RUB",
    "datetime": datetime.now(),
    "sender": "TEST",
    "codepro": "false",
    "secret": secret,
    "label": account,
}  # YooMoney выдаёт другие значения, но для тестов не важно

data["sha1_hash"] = sha1_hash(
    "&".join([str(e[1]) for e in data.items()]).replace(" ", "%20")
)

response = requests.post(api_url, headers=headers, data=data)

if response.status_code == 200:
    print("Платёж принят!")
else:
    print("Ошибка: ", response.status_code, response.text)
