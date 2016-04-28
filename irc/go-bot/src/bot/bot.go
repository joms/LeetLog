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
}

//const CmdPrefix = "&"
//const LeetPrefix = []string{" ","^"}

// ResponseHandler must be implemented by the protocol to handle the bot responses
type ResponseHandler func(target, message string, sender *User)

// Handlers that must be registered to receive callbacks from the bot
type Handlers struct {
	Response ResponseHandler
}

type JsonLeet struct {
	Time string
	Channel string
	Nick string
	Status int
}

// New configures a new bot instance
func New(h *Handlers, a []string, i *irc.Connection) *Bot {
	b := &Bot{
		handlers: h,
		admins: a,
		ircCon: i,
		CmdPrefix: "&",
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
		if (t.Hour() == 13 && t.Minute() >= 35 && t.Minute() <= 39) {
			jd := &JsonLeet{
				Time: "halla",
				Channel: "hei",
				Nick: "hei",
				Status: 0,
			}

			bd, err := json.Marshal(jd)
			if (err == nil) {
				fmt.Println(string(bd))
			}

			if err != nil {
				fmt.Println(err)
			}

			//b.postData("http://localhost:8000", json.Marshal(bd))

			b.Leet(channel, sender, text, t)
		}
	}
}

// Check if a given *User is admin
func (b *Bot) IsAdmin(u *User) bool {
	return inArray(u.Nick, b.admins)
}

func inArray(needle string, haystack []string) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}