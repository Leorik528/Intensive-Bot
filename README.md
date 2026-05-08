# Intensive Bot MVP

MVP-бот для продажи доступов к Telegram-интенсивам.

## Что реализовано в каркасе
- Базовая структура проекта по слоям (app/bot/domain/repository/service).
- Команда `/start` с клавиатурой выбора интенсива.
- Заготовки для платежей (`sendInvoice`, `successful_payment`), инвайтов и напоминаний.
- PostgreSQL миграция для пользователей, интенсивов и регистраций.
- Docker Compose для локального Postgres.

## Следующие шаги
1. Реализовать callback-обработчики выбора интенсива.
2. Добавить отправку `sendInvoice` и обработку pre-checkout/successful payment.
3. Реализовать `createChatInviteLink`, повторную генерацию ссылок и `banChatMember` при возврате.
4. Добавить админ-панель (Gin/Fiber/Echo) для CRUD интенсивов и рассылок.
