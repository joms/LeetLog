package bot

import (
	"fmt"
	"time"
)

type Bot struct{}

const CmdPrefix = "&"
const LeetPrefix = " "

func (b *Bot) MessageReceived(channel string, text string, sender *User) {
	//command := parse(text, channel, sender)
	//fmt.Println(command)
	t := time.Now()
	fmt.Println(t.Format("13:37:01:123"))
}