package views

import (
    "bytes"
    "context"
    appmodels "telegram-dice-bot/models"

    "github.com/go-telegram/bot"
    tgmodels "github.com/go-telegram/bot/models"
)

func SendQRImage(ctx context.Context, b *bot.Bot, msg *appmodels.Message, png []byte, caption string) {
    b.SendPhoto(ctx, &bot.SendPhotoParams{
        ChatID:          msg.ChatID,
        Caption:         caption,
        MessageThreadID: msg.ThreadID,
        Photo: &tgmodels.InputFileUpload{
            Filename: "solana_qr.png",
            Data:     bytes.NewReader(png),
        },
        ReplyParameters: &tgmodels.ReplyParameters{MessageID: msg.MessageID},
    })
}

func SendQRError(ctx context.Context, b *bot.Bot, msg *appmodels.Message, err error) {
    msg.SendMessage(ctx, b, "QR error: "+err.Error())
}

