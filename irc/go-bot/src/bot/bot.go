package bot

import (
	"time"
	"github.com/thoj/go-ircevent"
	"encoding/json"
	"fmt"
)

// Bot handles the bot instance
type Bot struct {
	handlers *Handlers
	admins []string
	ircCon *irc.Connection
	CmdPrefix string
	LeetPrefix []string
}

// ResponseHandler must be implemented by the protocol to handle the bot responses
type ResponseHandler func(target, message string, sender *User)

// Handlers that must be registered to receive callbacks from the bot
type Handlers struct {
	Response ResponseHandler
}

// New configures a new bot instance
func New(h *Handlers, a []string, i *irc.Connection) *Bot {
	b := &Bot{
		handlers: h,
		admins: a,
		ircCon: i,
		CmdPrefix: "&",
		LeetPrefix: []string{" ", "^"},
	}
	return b
}

// We've received a message to massage
func (b *Bot) MessageReceived(channel string, text string, sender *User, t time.Time) {
	var msg = false

	// If it was an msg, check for admin rights
	if sender.Nick == channel {
		if ! b.IsAdmin(sender) {
			return;
		} else {
			msg = true
		}
	}

	// Parse input
	command := b.Parse(text, channel, sender, msg)

	// Do something with the result
	if command != nil {
		switch command.Command {
		// Hardcoded commands
		case "join":
			if ! b.IsAdmin(command.User) {
				return;
			}

			if msg == false {
				return;
			}

			b.ircCon.Join(command.Args[0])

		case "leave":
			if ! b.IsAdmin(command.User) {
				return;
			}

			if msg == false {
				return;
			}

			b.ircCon.Part(command.Args[0])

		case "nick":
			if ! b.IsAdmin(command.User) {
				return;
			}

			if msg == false {
				return;
			}

			b.ircCon.Nick(command.Args[0])

		default:
			b.handleCmd(command)
		}
	} else {
		// http://i.imgur.com/khRqBiC.gif
		if (t.Hour() == 13 && t.Minute() >= 35 && t.Minute() <= 39) {
			leet := b.Leet(channel, sender, text, t)

			leetData, err := json.Marshal(leet)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(leetData))

			bot.postData("http://localhost:8000", leetData)
	}
}

// Check if a given *User is admin
func (b *Bot) IsAdmin(u *User) bool {
	return inArray(u.Nick, b.admins)
}
