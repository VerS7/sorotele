# Cоздание web-сайта для посредника организации ООО "Соро-телеком"

## Структура

Весь проект разделён на **фронтенд** и **бэкенд**. Сборка и запуск осуществляется через **docker** и **docker compose**

### Frontend

Визуальная часть написана на **Vue 3** с **Vuetify** с использованием **TypeScript**.

Так же использовано:

-   [Vuetify](https://vuetifyjs.com) для красивых компонентов
-   [Vue Router](https://router.vuejs.org/) роутер для SPA-функционала
-   [vue-recaptcha](https://dansnow.github.io/vue-recaptcha/) для Google-каптчи

### Backend

Серверная часть написана на **Golang**. В серверной части реализована авторизация и аутентификация, а так же оплата через **Юmoney**.

Вся информация по пользователям лежит в базе данных **Postgres**.

Так же использовано:

-   [Godotenv](https://github.com/joho/godotenv) для загрузки **.env** файлов
-   [Gorm](https://gorm.io) ORM для работы с базой данных

## Настройка

Для запуска приложения необходимо заполнить конфигурационный файл .env

```ini
# -- FRONTEND ----------------------------------------------
# > BUILD
RECAPTCHA_KEY="..."
API_PATH="api"
# -- BACKEND -----------------------------------------------
# > SERVER
FRONTEND_PATH="/app/dist"
BACKEND_PORT="8080"
# > DATABASE
DB_HOST="db"
DB_USERNAME="..."
DB_PASSWORD="..."
DB_NAME="Sorotele"
DB_SSL="disable" # disable / require
# > ADMIN
ADMIN_USERNAME="..."
ADMIN_PASSWORD="..."
# > EMAIL
EMAIL_SMTP_HOST="smtp.mail.ru"
EMAIL_SMTP_PORT="587"
EMAIL_FROM="..."
EMAIL_TO="..."
EMAIL_PASSWORD="..."
# > PAYMENT
YOOMONEY_CLIENT_TOKEN="..."
YOOMONEY_PAYMENT_NOTIFICATION_URL="<URL САЙТА>/api/payment/notification"
YOOMONEY_RECIEVER="..."
YOOMONEY_ENSURE_SECRET="..."
```

### Конфигурация .env

-   **RECAPTCHA_KEY**

    -   Необходимо для функционирования каптчи на сайте. Подробнее можно узнать **[тут](https://developers.google.com/recaptcha/docs/v3)**.

-   **ADMIN_USERNAME ADMIN_PASSWORD**

    -   Данные для генерации "нулевого" пользователя - **Администратора**.

-   **EMAIL_TO**

    -   **email**, куда будут прилетать оповещения о новых заявках на **подключение**.

-   **EMAIL_FROM EMAIL_PASSWORD**

    -   **smtp-аккаунт** для отправки оповещений. Пример с **mail.ru** можно посмотреть **[тут](https://help.mail.ru/mail/mailer/popsmtp/)**.

-   **DB_USERNAME DB_PASSWORD**

    -   Данные для доступа к БД.

-   **YOOMONEY_CLIENT_TOKEN**

    -   Токен для доступа к ЮMoney. Получение показано ниже.

-   **YOOMONEY_RECIEVER**

    -   Счёт получателя из ЛК ЮMoney.

-   **YOOMONEY_ENSURE_SECRET**
    -   Секретный код, выдаваемый при создании сервиса ЮMoney API.

### ЮMoney токен

Получить токен доступа можно разными способами. Для начала нужно зарегистрироваться **[тут](https://yoomoney.ru)**. Документацию по ЮMoney можно получить **[тут](https://yoomoney.ru/docs/wallet)**.

Для удобства был написан python-скрипт, упрощающий получение токена.

```python
...
# --- Пользовательские данные ---
client_id = "ID КЛИЕНТА"
redirect_uri = "REDIRECT URI ИЗ YOOMONEY SERVICES"
client_secret = "OAUTH SECRET ПРИ НАЛИЧИИ"
scope = [
    "account-info",
    "operation-history",
    "operation-details",
    "incoming-transfers",
    "payment-p2p",
    "payment-shop",
]
# --- Пользовательские данные ---
...
```

Запуск скрипта:

`python3 -m pip install requests`

`python3 yoomoney_client_token.py`

Код получения токена и некоторые идеи проекта были взяты из библиотеки **[yoomoney-api](https://github.com/AlekseyKorshuk/yoomoney-api/blob/master/yoomoney/authorize/authorize.py)** для **Python**. Рекоммендую поставить ⭐️ [автору](https://github.com/AlekseyKorshuk).

## Сборка

Для сборки проекта необходимо установить **[docker](https://docs.docker.com/get-started/get-docker/)**.

Собрать и запустить приложение можно командой `docker compose up -d --build` или через **Docker Desktop**.

## Тестовый функционал

По умолчанию зайти в личный кабинет как администратор можно указав лицевой счёт `0` и пароль из .env файла.

Добавить пользователя или тариф можно в админ-панели (`<URL САЙТА>/admin` или нажать на **Администратор** в личном кабинете).

Добавление пользователя доступно и через python-скрипт **create_user.py**. Все подробности внутри можно найти внутри файла.

```python
# --- Пользовательские данные ---
admin_token = "..."  # Токен доступа пользователя-админа
api_url = "<URL САЙТА>/api/user/create"  # URL для запросов на создание пользователя
headers = {"Content-Type": "application/json", "Authentication-Token": admin_token}
user = {
    "name": "...", # Имя пользователя
    "surname": "...", # Фамилия пользователя
    "password": "...", # Пароль пользователя
    "rate_id": "...", # ID тарифа. Можно посмотреть в БД
}
# --- Пользовательские данные ---
```

Тестовую оплату можно провести при помощи python-скрипта **new_payment.py**.

```python
# --- Пользовательские данные ---
api_url = "<URL САЙТА>/api/payment/notification"  # URL для запросов на создание платежа
secret = "..."  # Секретный ключ приложения для проверки подленности
account = "sr8125070062432387"  # Лицевой счёт пользователя
data = {
    "notification_type": "income",
    "operation_id": ...,
    "amount": "1337.07",
    "currency": "RUB",
    "datetime": ...,
    "sender": "TEST",
    "codepro": "false",
    "secret": secret,
    "label": account,
}  # YooMoney выдаёт другие значения, но для тестов не важно
# --- Пользовательские данные ---
```

Тестовое списание таким же образом можно провести при помощи python-скрипта **new_charge.py**.

```python
# --- Пользовательские данные ---
admin_token = "..."  # Токен доступа пользователя-админа
api_url = "<URL САЙТА>/api/payment/charge"  # URL для запросов на списание
data = {
    "account": "...",  # Лицевой счёт пользователя
    "amount": 1337.07,
}
# --- Пользовательские данные ---
```
