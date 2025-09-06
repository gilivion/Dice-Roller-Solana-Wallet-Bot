package models

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)
type Message struct {
	ChatID    int64 
	Text      string
	ThreadID  int
	MessageID int
}

func (m *Message) SendMessage(ctx context.Context, b *bot.Bot, response string) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:          m.ChatID,
		Text:            response,
		MessageThreadID: m.ThreadID,
		ReplyParameters: &models.ReplyParameters{
			MessageID: m.MessageID,
		},
	})
}
