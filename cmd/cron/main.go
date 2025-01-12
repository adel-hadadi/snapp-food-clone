package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"snapp-food/cmd/app"
	"snapp-food/data/database"
	"syscall"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	db := database.New()

	app := app.New(db)

	cron, err := gocron.NewScheduler()
	if err != nil {
		log.Panicf("error on create new cronjob %v", cron)
	}

	cron.NewJob(
		gocron.DurationJob(10*time.Second),
		gocron.NewTask(app.Services.Order.RemovePending),
	)

	cron.Start()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)

	<-s

	if err := cron.Shutdown(); err != nil {
		log.Println(fmt.Errorf("error on shuting down: %w", err))
	}
}
