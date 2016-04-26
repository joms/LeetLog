package bot

import (
	"time"
	"github.com/thoj/go-ircevent"
	"os"
	"fmt"
	"go/types"
)

// Bot handles the bot instance
type Bot struct {
	handlers *Handlers
	admins []string
	ircCon *irc.Connection
}

const CmdPrefix = "&"
const LeetPrefix = []string{" ","^"}

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

	if (t.Hour() == 13 && t.Minute() >= 36 && t.Minute() <= 38) {
		logToFile(t.Format("2006/01/02-15:04:05.999") + " " + sender.Nick + ": " + text + "\n")
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
	}
}

// Check if a given *User is admin
func (b *Bot) IsAdmin(u *User) bool {
	return inArray(u.Nick, b.admins)
}

func logToFile(str string) {
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(str); err != nil {
		panic(err)
	}

	fmt.Println(err)
}

func inArray(needle string, haystack types.Slice) {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}