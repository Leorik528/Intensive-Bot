# Telegram Bot для продажи доступов к интенсивам

Бот для Telegram, который автоматизирует продажу доступов к интенсивам: выбор потока, оплата через Telegram Payments, выдача одноразовых invite-ссылок в закрытые чаты и базовое администрирование через HTTP API.

---

## Возможности

### Для пользователей

* Просмотр списка открытых интенсивов через `/start`
* Выбор потока через inline-кнопки
* Оплата через Telegram Payments
* Автоматическая выдача одноразовой invite-ссылки после оплаты
* Повторная генерация invite-ссылки
* Доступ только в закрытые Telegram-чаты

### Для администраторов

* CRUD для интенсивов
* Открытие / закрытие продаж
* Массовые рассылки участникам
* Базовый refund workflow
* HTTP admin API на Gin

---

## Архитектура

### Основные компоненты

* Telegram Bot API
* Telegram Payments
* PostgreSQL
* Gin HTTP Server
* Invite Links API
* Admin REST API

---

## Поток оплаты

```text
/start
   ↓
Пользователь выбирает интенсив
   ↓
Бот отправляет sendInvoice
   ↓
pre_checkout_query
   ↓
successful_payment
   ↓
createChatInviteLink
   ↓
Пользователь получает invite-ссылку
```

---

## Стек технологий

* Go
* Gin
* PostgreSQL
* Telegram Bot API
* Telegram Payments

---

# Быстрый старт

## 1. Клонирование репозитория

```bash
git clone https://github.com/yourname/intensive-bot.git
cd intensive-bot
```

---

## 2. Настройка переменных окружения

Создайте `.env` файл:

```env
TELEGRAM_TOKEN=your_bot_token
# BOT_TOKEN=your_bot_token # алиас для обратной совместимости
PAYMENT_PROVIDER_TOKEN=your_payment_provider_token
DATABASE_URL=postgres://user:password@localhost:5432/intensive_bot?sslmode=disable
ADMIN_HTTP_ADDR=:8080
YOOKASSA_SHOP_ID=your_shop_id
YOOKASSA_SECRET_KEY=your_secret_key
YOOKASSA_RETURN_URL=https://t.me/your_bot_username
```

---

## 3. Установка PostgreSQL

Создайте базу данных:

```sql
CREATE DATABASE intensive_bot;
```

---

## 4. Запуск

```bash
go run ./cmd/bot
```

---

# Переменные окружения

| Переменная               | Описание                             |
| ------------------------ | ------------------------------------ |
| `TELEGRAM_TOKEN`         | Токен Telegram-бота (рекомендуется)  |
| `BOT_TOKEN`              | Алиас `TELEGRAM_TOKEN` (legacy)      |
| `PAYMENT_PROVIDER_TOKEN` | Токен платежного провайдера Telegram |
| `DATABASE_URL`           | PostgreSQL connection string         |
| `ADMIN_HTTP_ADDR`        | Адрес admin HTTP сервера             |
| `YOOKASSA_SHOP_ID`       | Shop ID ЮKassa                       |
| `YOOKASSA_SECRET_KEY`    | Секретный ключ ЮKassa                |
| `YOOKASSA_RETURN_URL`    | URL возврата после оплаты            |

---

# Telegram Permissions

Бот должен быть администратором закрытого чата с правами:

* Invite Users
* Ban Users

Без этих прав:

* не будут работать invite-ссылки
* не будут работать возвраты

---

# API Админ-панели

## Получить список интенсивов

```http
GET /admin/intensives
```

---

## Создать интенсив

```http
POST /admin/intensives
Content-Type: application/json
```

Пример:

```json
{
  "title": "Go Intensive #5",
  "price": 9900,
  "chat_id": -1001234567890,
  "is_open": true
}
```

---

## Обновить интенсив

```http
PUT /admin/intensives/:id
```

---

## Открыть / закрыть продажи

```http
POST /admin/intensives/:id/toggle?open=true
```

или

```http
POST /admin/intensives/:id/toggle?open=false
```

---

## Рассылка участникам

```http
POST /admin/intensives/:id/broadcast
```

Пример body:

```json
{
  "message": "Завтра стартуем в 19:00 🚀"
}
```

---

# Telegram Handlers

## `/start`

Показывает список доступных интенсивов.

---

## `intensive:{id}`

Создает invoice через `sendInvoice`.

---

## `pre_checkout_query`

Подтверждает платеж перед оплатой.

---

## `successful_payment`

После успешной оплаты:

* фиксирует оплату
* создает одноразовую invite-ссылку
* отправляет ссылку пользователю

---

## `regen:{id}`

Повторно генерирует invite-ссылку.

---

## Refund (MVP)

```bash
/refund <chat_id> <user_id>
```

Текущая реализация:

* удаляет пользователя из закрытого чата через `banChatMember`

> В MVP возврат денежных средств через платежного провайдера не реализован.

---

# Структура проекта

```text
.
├── cmd/
│   └── bot/
├── internal/
│   ├── bot/
│   ├── payments/
│   ├── admin/
│   ├── db/
│   └── telegram/
├── migrations/
├── .env
├── go.mod
└── README.md
```

---


## Планируется

* JWT/Auth для админки
* Web UI
* Настоящие refunds
* Webhooks от платежных систем
* Поддержка нескольких потоков
* Аналитика продаж
* Автоматическое удаление после refund
* Ограничение срока invite-ссылок

---

# Пример сценария использования

```text
1. Пользователь нажимает /start
2. Выбирает интенсив
3. Оплачивает через Telegram
4. Получает invite-ссылку
5. Заходит в закрытый чат
```

# Автор

Разработано для автоматизации продаж Telegram-интенсивов 🚀
