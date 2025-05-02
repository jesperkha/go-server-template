package main

import (
	"log"
	"os"
	"syscall"

	"github.com/jesperkha/go-server-template/config"
	"github.com/jesperkha/go-server-template/server"
	"github.com/jesperkha/notifier"
)

func main() {
	notif := notifier.New()

	config := config.Load()
	server := server.New(config)

	go server.ListenAndServe(notif)

	notif.NotifyOnSignal(os.Interrupt, syscall.SIGTERM)
	log.Println("shutdown")
}
