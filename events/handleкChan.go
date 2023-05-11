package events

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"lab4/user"
)

func RandomImageChan(ctx context.Context, b *bot.Bot, update *models.Update) {
	d := new(user.Data)
	repo.Open(d, "../storage/user.json")
	d.AddUser(update.Message.Chat.Username)
	repo.Save(d, "../storage/user.json")

	image := make(chan string)
	errChan := make(chan error)

	go func() {
		url, err := URL()
		if err != nil {
			errChan <- err
		} else {
			image <- url
			fmt.Println("---------------------first goroutines-----------------------")
		}
	}()
	go func() {
		url, err := URL()
		if err != nil {
			errChan <- err
		} else {
			image <- url
			fmt.Println("---------------------second goroutines-----------------------")
		}
	}()
	var url string
	select {
	case url = <-image:
	case err := <-errChan:
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Error fetching image URL: %s", err),
		})
		return
	}
	params := &bot.SendPhotoParams{
		ChatID:  update.Message.Chat.ID,
		Photo:   &models.InputFileString{Data: url},
		Caption: "",
	}
	b.SendPhoto(ctx, params)
}
func Handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	d := new(user.Data)
	repo.Open(d, "../storage/user.json")
	fmt.Println(update.Message.Chat.Username)
	d.AddUser(update.Message.Chat.Username)
	repo.Save(d, "../storage/user.json")
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "wrong command",
	})
}
