package controllers

import (
    "context"
    "strings"
    dice "telegram-dice-bot/models"
    view "telegram-dice-bot/views"

    "github.com/go-telegram/bot"
    tgmodels "github.com/go-telegram/bot/models"
)

func HandleMessage(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
    msg := update.Message
    if msg == nil {
        return
    }

    messageModel := &dice.Message{
        ChatID:    msg.Chat.ID,
        Text:      msg.Text,
        ThreadID:  int(msg.MessageThreadID),
        MessageID: msg.ID,
    }

    isPrivate := msg.Chat.Type == tgmodels.ChatTypePrivate

    if strings.HasPrefix(messageModel.Text, "/help") {
        if isPrivate {
            view.SendHelpPrivate(ctx, b, messageModel)
        } else {
            view.SendHelpGroup(ctx, b, messageModel)
        }
        return
    }

    if isPrivate {
        if strings.HasPrefix(messageModel.Text, "/qr") {
            parts := strings.Fields(messageModel.Text)
            var amount string
            if len(parts) > 1 {
                amount = parts[1]
            }
            w, ok := dice.GetWallet(msg.From.ID)
            if !ok {
                w = dice.EnsureWallet(msg.From.ID)
            }
            uri := dice.BuildSolanaPayURI(w.AddressBase58(), amount)
            png, err := generateQRPNG(uri)
            if err != nil {
                view.SendQRError(ctx, b, messageModel, err)
                return
            }
            caption := "Solana Pay QR"
            if amount != "" {
                caption += " (amount: " + amount + ")"
            }
            view.SendQRImage(ctx, b, messageModel, png, caption)
            return
        }
        if strings.HasPrefix(messageModel.Text, "/wallet_address") {
            if w, ok := dice.GetWallet(msg.From.ID); ok {
                view.SendWalletAddress(ctx, b, messageModel, w.AddressBase58())
            } else {
                w := dice.EnsureWallet(msg.From.ID)
                view.SendWalletAddress(ctx, b, messageModel, w.AddressBase58())
            }
            return
        }

        if strings.HasPrefix(messageModel.Text, "/wallet_private_key") {
            if w, ok := dice.GetWallet(msg.From.ID); ok {
                view.SendWalletPrivateKey(ctx, b, messageModel, w.PrivateKeyBase58())
            } else {
                w := dice.EnsureWallet(msg.From.ID)
                view.SendWalletPrivateKey(ctx, b, messageModel, w.PrivateKeyBase58())
            }
            return
        }

        if strings.HasPrefix(messageModel.Text, "/wallet_balance") {
            if w, ok := dice.GetWallet(msg.From.ID); ok {
                if bal, err := w.GetBalance(); err != nil {
                    view.SendWalletError(ctx, b, messageModel, err)
                } else {
                    view.SendWalletBalance(ctx, b, messageModel, bal)
                }
            } else {
                w := dice.EnsureWallet(msg.From.ID)
                if bal, err := w.GetBalance(); err != nil {
                    view.SendWalletError(ctx, b, messageModel, err)
                } else {
                    view.SendWalletBalance(ctx, b, messageModel, bal)
                }
            }
            return
        }
    }

    if strings.HasPrefix(messageModel.Text, "/roll") {
        command := strings.TrimSpace(strings.TrimPrefix(messageModel.Text, "/roll"))

        if command == "" {
            total, rolls := dice.RollMultipleDice("d20", 1)
            view.SendRollResultMessage(ctx, b, messageModel, "d20", total, rolls, 0)
            return
        }

		numOfDice, diceType, modifier, err := dice.ParseRollCommand(command)
		if err != nil || !dice.IsValidDice(diceType) {
			view.SendInvalidDiceMessage(ctx, b, messageModel)
			return
		}

		total, rolls := dice.RollMultipleDice(diceType, numOfDice)

		totalWithModifier := total + modifier

        view.SendRollResultMessage(ctx, b, messageModel, diceType, totalWithModifier, rolls, modifier)
    }
}
