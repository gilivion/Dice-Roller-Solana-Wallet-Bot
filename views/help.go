package views

import (
    "context"
    "strings"
    appmodels "telegram-dice-bot/models"

    "github.com/go-telegram/bot"
)

func SendHelpGroup(ctx context.Context, b *bot.Bot, msg *appmodels.Message) {
    lines := []string{
        "Available commands:",
        "/roll [NdX[+M]] - Roll dice (e.g. /roll 2d6+1)",
        "/help - Show this help",
    }
    msg.SendMessage(ctx, b, strings.Join(lines, "\n"))
}

func SendHelpPrivate(ctx context.Context, b *bot.Bot, msg *appmodels.Message) {
    lines := []string{
        "Available commands:",
        "/roll [NdX[+M]] - Roll dice (e.g. /roll 2d6+1)",
        "/wallet_address - Show your Solana wallet address",
        "/wallet_balance - Show your wallet balance",
        "/wallet_private_key - Show your private key (base58) — be careful!",
        "/help - Show this help",
    }
    msg.SendMessage(ctx, b, strings.Join(lines, "\n"))
}

