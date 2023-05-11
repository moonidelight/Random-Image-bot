package events

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"lab4/user"
	"sync"
)

func RandomImageWg(ctx context.Context, b *bot.Bot, update *models.Update) {
	d := new(user.Data)
	repo.Open(d, "../storage/user.json")
	d.AddUser(update.Message.Chat.Username)
	repo.Save(d, "../storage/user.json")
	var url string
	wg := sync.WaitGroup{}
	wg.Add(2)
	errChan := make(chan error)
	urlChan := make(chan string, 2)
	go func() {
		defer wg.Done()
		u, err := URL()
		if err != nil {
			errChan <- err
		} else {
			urlChan <- u
			fmt.Println("--------------first-----------------------")
		}
	}()
	go func() {
		defer wg.Done()
		u, err := URL()
		if err != nil {
			errChan <- err
		} else {
			urlChan <- u
			fmt.Println("--------------second---------------------")
		}
	}()
	select {
	case url = <-urlChan:
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

	wg.Wait()
}
