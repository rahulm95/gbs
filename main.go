package main

import (
	"log"

	_ "github.com/voonik/goFramework/pkg/event"
	_ "github.com/voonik/goFramework/pkg/logger"
	"github.com/voonik/goFramework/pkg/server"
)

func main() {
	log.Println("Service starting up")
	server.Init()
	log.Println("Service started")
	defer server.Finish()
}
