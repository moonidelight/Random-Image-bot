package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"lab4/events"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load("../configs/.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	token := os.Getenv("TELEGRAM_TOKEN")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	fmt.Println("starting something .....")
	opts := []bot.Option{
		bot.WithDebug(),
		bot.WithDefaultHandler(events.Handler),
	}
	b, err := bot.New(token, opts...)
	if err != nil {
		return
	}
	b.RegisterHandler(bot.HandlerTypeMessageText, "/image", bot.MatchTypeExact, events.RandomImageMutex)
	fmt.Println("hello world")
	// to terminate the program CTRL + C
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// block until a signal is received
		<-sigs
		// handle the signal
		fmt.Println("received interrupt signal, terminating...")
		os.Exit(0)
	}()

	b.Start(ctx)

}
