package bot

import (
	"time"
	"fmt"
)

type Bot struct{}

const CmdPrefix = "&"
const LeetPrefix = " "



func (b *Bot) MessageReceived(channel string, text string, sender *User, t time.Time) {
	command := parse(text, channel, sender)
	fmt.Println(command)

	switch command.Command {
	case " ":
		b.leet(command, t)
	}
}