module intensive-bot

go 1.23

require (
	github.com/joho/godotenv v1.5.1
	github.com/go-co-op/gocron v1.37.0
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/jackc/pgx/v5 v5.7.4
)

replace github.com/joho/godotenv => ./third_party/godotenv
