package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New(os.Getenv("BOT_TOKEN"))
	if nil != err {
		panic(err)
	}

	b.RegisterHandlerMatchFunc(matchFunc, helloHandler)

	b.Start(ctx)
}

func matchFunc(update *models.Update) bool {
	if len(update.Message.NewChatMembers) > 0 {
		return true
	}
	return false
}

func helloHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	for _, user := range update.Message.NewChatMembers {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      bot.EscapeMarkdown(fmt.Sprintf("Добро пожаловать в чат, %s! \nПожалуйста напиши рассказ #осебе и скинь нюдсы", user.FirstName)),
			ParseMode: models.ParseModeMarkdown,
			ReplyParameters: &models.ReplyParameters{
				ChatID:    update.Message.Chat.ID,
				MessageID: update.Message.ID,
			},
		})
		if err != nil {
			panic(err)
		}
	}
}
