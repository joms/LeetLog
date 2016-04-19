package bot

import (
	"time"
	"github.com/thoj/go-ircevent"
)

// Bot handles the bot instance
type Bot struct {
	handlers *Handlers
	admins []string
	ircCon *irc.Connection
}

const CmdPrefix = "&"
const LeetPrefix = " "

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
	command := parse(text, channel, sender, msg)

	// Do something with the result
	if command != nil {
		switch command.Command {
		case " ":
			b.leet(command, t)

		// Hardcoded commands
		case "join":
			if ! b.IsAdmin(command.User) {
				return;
			}

			if msg == false {
				return;
			}

			b.join(command)

		case "leave":
			if ! b.IsAdmin(command.User) {
				return;
			}

			if msg == false {
				return;
			}

			b.leave(command)

		default:
			b.handleCmd(command)
		}
	}
}

// Check if a given *User is admin
func (b *Bot) IsAdmin(u *User) bool {
	for _, a := range b.admins {
		if a == u.Nick {
			return true
		}
	}
	return false
}