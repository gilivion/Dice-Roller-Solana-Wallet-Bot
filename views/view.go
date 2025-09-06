package views

import (
	"context"
	"fmt"
	"strings"
	"telegram-dice-bot/models"

	"github.com/go-telegram/bot"
)

func SendRollResultMessage(ctx context.Context, b *bot.Bot, msg *models.Message, diceType string, total int, rolls []int, modifier int) {
	rollsStr := make([]string, len(rolls))
	for i, roll := range rolls {
		rollsStr[i] = fmt.Sprintf("%d", roll)
	}
	rollsResult := strings.Join(rollsStr, ", ")

    response := fmt.Sprintf("Result %d%s: %s (sum: %d)", len(rolls), diceType, rollsResult, total)
    if modifier != 0 {
        response += fmt.Sprintf(" + %d (sum: %d)", modifier, total)
    }

	msg.SendMessage(ctx, b, response)
}

func SendInvalidDiceMessage(ctx context.Context, b *bot.Bot, msg *models.Message) {
    response := "Invalid dice! Valid dice: d4, d6, d8, d10, d12, d20, d100."
    msg.SendMessage(ctx, b, response)
}
