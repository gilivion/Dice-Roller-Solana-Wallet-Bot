package main

import (
    "context"
    "database/sql"
    "log"
    "os"
    "telegram-dice-bot/controllers"
    appmodels "telegram-dice-bot/models"

    _ "github.com/jackc/pgx/v5/stdlib"
    "github.com/go-telegram/bot"
)

func main() {
    botToken := os.Getenv("BOT_TOKEN")
    if botToken == "" {
        log.Fatal("BOT_TOKEN is not set")
    }

    if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
        conn, err := sql.Open("pgx", dsn)
        if err != nil {
            log.Fatalf("DB open error: %v", err)
        }
        if err := conn.Ping(); err != nil {
            log.Fatalf("DB ping error: %v", err)
        }
        appmodels.InitDB(conn)
        if err := appmodels.Migrate(); err != nil {
            log.Fatalf("DB migration error: %v", err)
        }
    }

    b, err := bot.New(botToken, bot.WithDefaultHandler(controllers.HandleMessage))
    if err != nil {
        log.Fatalf("Bot cretation error: %v", err)
    }

	ctx := context.Background()

	b.Start(ctx)

}
