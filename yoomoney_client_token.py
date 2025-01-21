"""
Автоматическое получение токена YooMoney
"""

import requests

# --- Пользовательские данные ---
client_id = "<ID_КЛИЕНТА>"
redirect_uri = "<REDIRECT_URI_ИЗ_YOOMONEY_SERVICES>"
client_secret = "<OAUTH_SECRET_ПРИ_НАЛИЧИИ>"
scope = [
    "account-info",
    "operation-history",
    "operation-details",
    "incoming-transfers",
    "payment-p2p",
    "payment-shop",
]
# --- Пользовательские данные ---

url = (
    "https://yoomoney.ru/oauth/authorize?client_id={client_id}&response_type=code"
    "&redirect_uri={redirect_uri}&scope={scope}".format(
        client_id=client_id,
        redirect_uri=redirect_uri,
        scope="%20".join([str(elem) for elem in scope]),
    )
)
headers = {"Content-Type": "application/x-www-form-urlencoded"}

response = requests.request("POST", url, headers=headers)

if response.status_code == 200:
    print("Перейдите на сайт и подтвердите запрос:")
    print(response.url)

code = str(
    input(
        "Введите полученный URL (прим. https://yourredirect_uri?code=XXXXXXXXXXXXX): "
    )
)
try:
    code = code[code.index("code=") + 5 :].replace(" ", "")
except Exception:
    pass

url = (
    "https://yoomoney.ru/oauth/token?code={code}&client_id={client_id}&"
    "grant_type=authorization_code&redirect_uri={redirect_uri}&client_secret={client_secret}".format(
        code=str(code),
        client_id=client_id,
        redirect_uri=redirect_uri,
        client_secret=client_secret,
    )
)

response = requests.request("POST", url, headers=headers)
print("Ваш токен доступа:")
print(response.json()["access_token"])
