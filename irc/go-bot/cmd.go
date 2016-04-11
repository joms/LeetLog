package main

import(
	"bot"
	"strings"
)

func main() {
	bot.Run(&bot.Config{
		Admins:		strings.Split("JoMs", ","),
		Server:		"localhost:6667",
		Channels: 	strings.Split("#sandkas.se", ","),
		User:		"LeetBot",
		Nick:		"LeetBot",
		Prefix: 	"&",
	})
}