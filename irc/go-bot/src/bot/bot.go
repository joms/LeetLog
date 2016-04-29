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
	Endpoint string
	EndpointKey string
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
		Endpoint: "http://localhost:8000/api/leet",
		EndpointKey: "abc123", // This key MUST correspond on your server
	}
	return b
}

// We've received a message to massage
func (b *Bot) MessageReceived(channel string, text string, sender *User, t time.Time) {
	var msg = false

	// If it was an msg, check for admin rights
	if sender.Nick == channel {
		if ! b.IsAdmin(sender) {
			return
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
				return
			}

			if msg == false {
				return
			}

			if (len(command.Args) == 1) {
				b.ircCon.Join(command.Args[0])
			} else {
				return
			}

		case "leave":
			if ! b.IsAdmin(command.User) {
				return
			}

			if msg == false {
				return
			}

			if (len(command.Args) == 1) {
				b.ircCon.Part(command.Args[0])
			} else {
				return
			}

		case "nick":
			if ! b.IsAdmin(command.User) {
				return
			}

			if msg == false {
				return
			}

			if (len(command.Args) == 1) {
				b.ircCon.Nick(command.Args[0])
			} else {
				return
			}

		case "set":
			if ! b.IsAdmin(command.User) {
				return
			}

			if msg == false {
				return
			}

			if (len(command.Args) == 2) {
				switch command.Args[0] {
				case "endpointkey":
					b.EndpointKey = command.Args[1]

				case "endpoint":
					b.Endpoint = command.Args[1]
				}
			}

		case "get":
			if ! b.IsAdmin(command.User) {
				return
			}

			if msg == false {
				return
			}

			if (len(command.Args) == 1) {
				switch command.Args[0] {
				case "endpointkey":
					b.handlers.Response(channel, b.EndpointKey, sender)

				case "endpoint":
					b.handlers.Response(channel, b.Endpoint, sender)
				}
			} else {
				return
			}

			return

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

			b.postData(b.Endpoint, leetData)
		}
	}
}

// Check if a given *User is admin
func (b *Bot) IsAdmin(u *User) bool {
	return inArray(u.Nick, b.admins)
}