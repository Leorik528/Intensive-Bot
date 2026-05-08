# Intensive Bot (MVP)

Telegram-бот для продажи доступа к интенсивам: выбор интенсива, оплата через Telegram Payments, выдача ссылки в закрытый Telegram-чат и базовые admin-функции.

## Что умеет MVP

### Для ученика
- Команда `/start` показывает список открытых интенсивов.
- При выборе интенсива бот отправляет счёт через `sendInvoice`.
- После `pre_checkout_query` и `successful_payment` бот:
  - фиксирует факт успешной оплаты (MVP-обработка события),
  - создаёт одноразовую invite-ссылку через `createChatInviteLink`,
  - отправляет ссылку ученику.
- Можно повторно запросить ссылку (callback `regen:{id}`).

### Для администратора
- Команда `/refund <chat_id> <user_id>` вызывает `banChatMember` для удаления участника при возврате.
- HTTP admin API на Gin:
  - `GET /admin/intensives`
  - `POST /admin/intensives`
  - `PUT /admin/intensives/:id`
  - `POST /admin/intensives/:id/toggle?open=true|false`
  - `POST /admin/intensives/:id/broadcast`

## Где находится «окружение» (`TELEGRAM_TOKEN`, `DATABASE_URL` и т.д.)

В Go `os.Getenv(...)` читает переменные окружения **из процесса**, а не из отдельного встроенного файла.

Это значит:
- переменные нужно передать перед запуском команды,
- либо загрузить их из `.env` в shell,
- либо прокинуть через Docker/CI.

### Вариант 1: локально через `.env`

1. Скопируйте шаблон:

```bash
cp .env.example .env
```

2. Заполните `.env` своими значениями.

3. Загрузите переменные в текущую сессию shell и запустите бота:

```bash
set -a
source .env
set +a
go run ./cmd/bot
```

### Вариант 2: разово перед запуском

```bash
TELEGRAM_TOKEN=... \
PAYMENT_PROVIDER_TOKEN=... \
DATABASE_URL='postgres://intensive_bot:intensive_bot@localhost:5432/intensive_bot?sslmode=disable' \
ADMIN_HTTP_ADDR=':8080' \
go run ./cmd/bot
```

## Переменные окружения
- `TELEGRAM_TOKEN` — токен Telegram-бота от `@BotFather`.
- `PAYMENT_PROVIDER_TOKEN` — токен провайдера Telegram Payments.
- `DATABASE_URL` — DSN подключения к PostgreSQL.
- `ADMIN_HTTP_ADDR` — адрес HTTP admin API (по умолчанию `:8080`).

## Быстрый старт

### 1) Поднять PostgreSQL
```bash
docker compose up -d db
```

### 2) Подготовить окружение
```bash
cp .env.example .env
# заполнить .env
set -a && source .env && set +a
```

### 3) Запустить бота
```bash
go run ./cmd/bot
```

## Важно для доступа в закрытый чат
Бот должен быть админом целевого чата и иметь права:
- приглашать пользователей (invite links),
- удалять/банить пользователей (для возвратов).

## Ограничения текущего MVP
- CRUD интенсивов сейчас in-memory (без постоянного хранения в БД).
- Нужна доработка полноценного учёта оплат/участников в PostgreSQL.
- Рассылки и напоминания пока в виде заготовки/базовой логики.
