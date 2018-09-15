package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/xiaq/tg"
)

func main() {
	buf, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatalln("cannot read token file:", err)
	}
	token := strings.TrimSpace(string(buf))

	bot := tg.NewCommandBot(token)
	bot.OnCommand("me", handleMe)
	//bot.OnCommand("slap", handlePia)
	//bot.OnCommand("prpr", handlePrpr)
	bot.Main()
}

func handleMe(bot *tg.CommandBot, text string, msg *tg.Message) {
	if msg.From == nil {
		return
	}
	reply := msg.From.DisplayName() + " " + text
	bot.Get("/sendMessage", tg.Query{
		"chat_id": msg.Chat.ID,
		"text":    reply,
	}, nil)
}
