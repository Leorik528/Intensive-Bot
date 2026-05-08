# Intensive Bot (MVP)

Telegram-бот для продажи доступа к интенсивам: выбор потока, оплата через Telegram Payments, выдача инвайт-ссылки в закрытый чат и базовая админ-панель.

## Что реализовано
- `/start` показывает список открытых интенсивов и кнопки выбора.
- Callback `intensive:{id}` отправляет `sendInvoice`.
- Обработка `pre_checkout_query` и `successful_payment`.
- После успешной оплаты создается одноразовая invite-ссылка (`createChatInviteLink`).
- Повторная генерация ссылки через callback `regen:{id}`.
- Команда возврата (MVP): `/refund <chat_id> <user_id>` вызывает `banChatMember`.
- Админ-панель на Gin:
  - `GET /admin/intensives`
  - `POST /admin/intensives`
  - `PUT /admin/intensives/:id`
  - `POST /admin/intensives/:id/toggle?open=true|false`
  - `POST /admin/intensives/:id/broadcast`

## Переменные окружения
- `TELEGRAM_TOKEN` — токен бота.
- `PAYMENT_PROVIDER_TOKEN` — токен платежного провайдера Telegram.
- `DATABASE_URL` — строка подключения к PostgreSQL.
- `ADMIN_HTTP_ADDR` — адрес админки (по умолчанию `:8080`).

## Запуск
```bash
go run ./cmd/bot
```

## Важно
Для работы с закрытым чатом бот должен быть администратором с правами:
- Invite Users
- Ban Users (для возвратов)

## Следующие шаги
- Перенести in-memory хранилище интенсивов в PostgreSQL.
- Добавить полноценный учет оплат/участников и рассылки по реальным участникам.
- Включить cron-напоминания (за день, час, 10 минут).
