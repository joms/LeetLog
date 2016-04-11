package main

import(
	"irc"
	"strings"
)

func main() {
	irc.Run(&irc.Config{
		Admins:		strings.Split("JoMs", ","),
		Server:		"localhost:6667",
		Channels: 	strings.Split("#sandkas.se", ","),
		User:		"LeetBot",
		Nick:		"LeetBot",
		Prefix: 	"&",
	})
}