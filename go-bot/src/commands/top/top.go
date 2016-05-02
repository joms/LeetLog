package top

import (
	"bot"
)

func top(command *bot.Cmd) (msg string, err error) {

	return "halla", nil
}

func init() {
	bot.RegisterCommand(
		"top",
		"Returns top 3 leets for a channel on any given date. Defaults to today",
		"",
		top,
		false,
		false)
}
