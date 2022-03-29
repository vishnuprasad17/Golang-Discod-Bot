package main

import (
	"fmt"

	"github.com/vishnu/discord-ping/bot"

	"github.com/vishnu/discord-ping/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()
	<-make(chan struct{})
	return
}
