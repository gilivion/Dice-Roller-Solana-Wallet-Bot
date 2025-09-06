package views

import (
    "context"
    "fmt"
    appmodels "telegram-dice-bot/models"

    "github.com/go-telegram/bot"
)

func SendWalletCreated(ctx context.Context, b *bot.Bot, msg *appmodels.Message, address string) {
    msg.SendMessage(ctx, b, fmt.Sprintf("Wallet created. Address: %s", address))
}

func SendWalletAddress(ctx context.Context, b *bot.Bot, msg *appmodels.Message, address string) {
    msg.SendMessage(ctx, b, fmt.Sprintf("Wallet address: %s", address))
}

func SendWalletPrivateKey(ctx context.Context, b *bot.Bot, msg *appmodels.Message, priv string) {
    msg.SendMessage(ctx, b, fmt.Sprintf("Private key (base58): %s", priv))
}

func SendWalletBalance(ctx context.Context, b *bot.Bot, msg *appmodels.Message, lamports uint64) {
    // 1 SOL = 1_000_000_000 lamports
    msg.SendMessage(ctx, b, fmt.Sprintf("Balance: %d lamports (%.9f SOL)", lamports, float64(lamports)/1_000_000_000.0))
}

func SendWalletError(ctx context.Context, b *bot.Bot, msg *appmodels.Message, err error) {
    msg.SendMessage(ctx, b, fmt.Sprintf("Wallet error: %v", err))
}
