package views

import (
    "context"
    appmodels "telegram-dice-bot/models"

    "github.com/go-telegram/bot"
)

func SendAuthPrompt(ctx context.Context, b *bot.Bot, msg *appmodels.Message) {
    msg.SendMessage(ctx, b, "Please login: /login <password>")
}

func SendAuthSuccess(ctx context.Context, b *bot.Bot, msg *appmodels.Message) {
    msg.SendMessage(ctx, b, "Logged in. Commands are available.")
}

func SendAuthFailed(ctx context.Context, b *bot.Bot, msg *appmodels.Message) {
    msg.SendMessage(ctx, b, "Invalid password. Try again: /login <password>")
}

func SendLogoutSuccess(ctx context.Context, b *bot.Bot, msg *appmodels.Message) {
    msg.SendMessage(ctx, b, "You have been logged out.")
}
