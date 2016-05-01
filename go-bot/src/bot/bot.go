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
			if (len(command.Args) == 1) {
				if ! b.IsAdmin(command.User) {
					fmt.Printf("Insufficient permissions: %s tried to join %s\n", command.User.Nick, command.Args[0])
					return
				}

				if msg == false {
					fmt.Printf("Invalid command: %s tried to join %s\n", command.User.Nick, command.Args[0])
					return
				}

				fmt.Printf("Success: %s joined %s\n", command.User.Nick, command.Args[0])
				b.ircCon.Join(command.Args[0])
			} else {
				fmt.Printf("Missing parameters: %s tried to join nothing\n", command.User.Nick)
				return
			}

		case "leave":
			if (len(command.Args) == 1) {
				if ! b.IsAdmin(command.User) {
					fmt.Printf("Insufficient permissions: %s tried to leave %s\n", command.User.Nick, command.Args[0])
					return
				}

				if msg == false {
					fmt.Printf("Invalid command: %s tried to leave %s\n", command.User.Nick, command.Args[0])
					return
				}

				fmt.Printf("Success: %s left %s\n", command.User.Nick, command.Args[0])
				b.ircCon.Part(command.Args[0])
			} else {
				fmt.Printf("Missing parameters: %s tried to leave nothing\n", command.User.Nick)
				return
			}

		case "nick":
			if (len(command.Args) == 1) {
				if ! b.IsAdmin(command.User) {
					fmt.Printf("Insufficient permissions: %s tried to change nick to %s\n", command.User.Nick, command.Args[0])
					return
				}

				if msg == false {
					fmt.Printf("Invalid command: %s tried to change nick to %s\n", command.User.Nick, command.Args[0])
					return
				}

				fmt.Printf("Success: %s changed nick to %s\n", command.User.Nick, command.Args[0])
				b.ircCon.Nick(command.Args[0])
			} else {
				fmt.Printf("Missing parameters: %s tried to change nick to nothing\n", command.User.Nick)
				return
			}

		case "set":
			if (len(command.Args) == 2) {
				if ! b.IsAdmin(command.User) {
					fmt.Printf("Insufficient permissions: %s tried to set %s to %s\n", command.User.Nick, command.Args[0], command.Args[1])
					return
				}

				if msg == false {
					fmt.Printf("Invalid command: %s tried to set %s to %s\n", command.User.Nick, command.Args[0], command.Args[1])
					return
				}

				switch command.Args[0] {
				case "endpointkey":
					b.EndpointKey = command.Args[1]

				case "endpoint":
					b.Endpoint = command.Args[1]
				}
			} else {
				fmt.Printf("Missing parameters: %s tried to set something to nothing", command.User.Nick)
				return
			}

		case "get":
			if (len(command.Args) == 1) {
				if ! b.IsAdmin(command.User) {
					fmt.Printf("Insufficient permissions: %s tried to get %s\n", command.User.Nick, command.Args[0])
					return
				}

				if msg == false {
					fmt.Printf("Invalid command: %s tried to get%s\n", command.User.Nick, command.Args[0])
					return
				}

				switch command.Args[0] {
				case "endpointkey":
					b.handlers.Response(channel, b.EndpointKey, sender)

				case "endpoint":
					b.handlers.Response(channel, b.Endpoint, sender)
				}

				fmt.Printf("Success: %s got %s\n", command.User.Nick, command.Args[0])
			} else {
				fmt.Printf("Missing parameters: %s tried to get nothing from something", command.User.Nick)
				return
			}

		default:
			fmt.Printf("%s is not a special command. Running default handler", command.Command)
			b.handleCmd(command)
		}
	} else {
		// http://i.imgur.com/khRqBiC.gif
		//if (t.Hour() == 13 && t.Minute() >= 35 && t.Minute() <= 39) {
			fmt.Printf("Not a command, and leet is closing in")

			leet := b.Leet(channel, sender, text, t)

			leetData, err := json.Marshal(leet)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("Sending Leet to remote API at %s\n", b.Endpoint)
			fmt.Println(string(leetData))

			b.PostLeet(b.Endpoint, leetData)
		//}
	}
}

// Check if a given *User is admin
func (b *Bot) IsAdmin(u *User) bool {
	return inArray(u.Nick, b.admins)
}