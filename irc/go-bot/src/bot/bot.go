package bot

import (
	"time"
)

// Bot handles the bot instance
type Bot struct {
	handlers *Handlers
	admins []string
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
func New(h *Handlers, a []string) *Bot {
	b := &Bot{
		handlers: h,
		admins: a,
	}
	return b
}

// We've received a message to massage
func (b *Bot) MessageReceived(channel string, text string, sender *User, t time.Time) {
	// If it was an msg, check for admin rights
	if sender.Nick == channel {
		if ! b.IsAdmin(sender) {
			return;
		}
	}

	// Parse input
	command := parse(text, channel, sender)

	// Do something with the result
	if command != nil {
		switch command.Command {
		case " ":
			b.leet(command, t)
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