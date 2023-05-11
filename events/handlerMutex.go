package events

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"lab4/user"
	"sync"
)

var (
	//counter = 0
	mutex = &sync.Mutex{}
)

func RandomImageMutex(ctx context.Context, b *bot.Bot, update *models.Update) {
	d := new(user.Data)
	repo.Open(d, "../storage/user.json")
	d.AddUser(update.Message.Chat.Username)
	repo.Save(d, "../storage/user.json")
	wg := sync.WaitGroup{}
	wg.Add(2)
	var err error
	var url string
	var sent bool
	go func() {
		defer wg.Done()
		mutex.Lock()
		if sent {
			mutex.Unlock()
			return
		}
		url, err = URL()
		fmt.Println("---------------first ---------------------------")
		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("Error fetching image URL: %s", err),
			})
			mutex.Unlock()
			return
		}
		params := &bot.SendPhotoParams{
			ChatID:  update.Message.Chat.ID,
			Photo:   &models.InputFileString{Data: url},
			Caption: "",
		}
		sent = true
		mutex.Unlock()
		b.SendPhoto(ctx, params)
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		if sent {
			mutex.Unlock()
			return
		}
		url, err = URL()
		fmt.Println("---------------second ---------------------------")
		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("Error fetching image URL: %s", err),
			})
			mutex.Unlock()
			return
		}
		params := &bot.SendPhotoParams{
			ChatID:  update.Message.Chat.ID,
			Photo:   &models.InputFileString{Data: url},
			Caption: "",
		}
		sent = true
		mutex.Unlock()
		b.SendPhoto(ctx, params)
	}()

	wg.Wait()
}
