# Cоздание web-сайта для посредника организации ООО "Соро-телеком"

По **ТЗ** необходимо было создать похожий на этот **[сайт](https://www.sorotele.com/)**, с кастомной реализацией оплаты (использовал ЮMoney) и личного кабинета.

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
FRONTEND_PATH="./dist"
BACKEND_PORT="8080"
# > DATABASE
DB_HOST="..."
DB_USERNAME="..."
DB_PASSWORD="..."
DB_NAME="..."
SSL="require"
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
YOOMONEY_RECIEVER="..."
SUCCESS_URL="..."
```

### Конфигурация .env

-   **RECAPTCHA_KEY**

    -   Необходимо для функционирования каптчи на сайте. Подробнее можно узнать **[тут](https://developers.google.com/recaptcha/docs/v3)**.

-   **ADMIN_USERNAME ADMIN_PASSWORD**

    -   Данные для генерации "нулевого" пользователя - **Администратора**. Пароль можно будет поменять в личном кабинете.

-   **EMAIL_TO**

    -   **email**, куда будут прилетать оповещения о новых заявках на **подключение**.

-   **EMAIL_FROM EMAIL_PASSWORD**
    -   **smtp-аккаунт** для отправки оповещений. Пример с **mail.ru** можно посмотреть **[тут](https://help.mail.ru/mail/mailer/popsmtp/)**.

### ЮMoney токен

Получить токен доступа можно разными способами. Для начала нужно зарегистрироваться **[тут](https://yoomoney.ru)**. Документацию по ЮMoney можно получить **[тут](https://yoomoney.ru/docs/wallet)**.

Для удобства был написан **python-скрипт**, упрощающий получение токена.

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

Код получения токена и некоторые идеи были взяты из библиотеки **[yoomoney-api](https://github.com/AlekseyKorshuk/yoomoney-api/blob/master/yoomoney/authorize/authorize.py)** для **Python**. Рекоммендую поставить ⭐️ [автору](https://github.com/AlekseyKorshuk).

## Сборка

Для сборки проекта необходимо установить **[docker](https://docs.docker.com/get-started/get-docker/)**.

Собрать и запустить приложение можно командой `docker compose up -d` или через **Docker Desktop**.
