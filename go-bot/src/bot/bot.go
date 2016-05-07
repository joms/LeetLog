package bot

import (
	"time"
	"github.com/thoj/go-ircevent"
	"encoding/json"
	"fmt"
)

// Bot handles the bot instance
type Bot struct {
	handlers       *Handlers
	admins         []string
	ircCon         *irc.Connection
	CmdPrefix      string
	LeetPrefix     []string
	APIEndpoint    string
	APIEndpointKey string
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
		APIEndpoint: "http://localhost:8000/api/leet",
		APIEndpointKey: "abc123", // This key MUST correspond on your server
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

				b.ircCon.Join(command.Args[0])

				b.ircCon.Privmsgf(command.User.Nick, "Joined %s", command.Args[0])
				fmt.Printf("Success: %s joined %s\n", command.User.Nick, command.Args[0])
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

				b.ircCon.Part(command.Args[0])

				b.ircCon.Privmsgf(command.User.Nick, "Left %s", command.Args[0])
				fmt.Printf("Success: %s left %s\n", command.User.Nick, command.Args[0])
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

				b.ircCon.Nick(command.Args[0])
				b.ircCon.Privmsgf(command.User.Nick, "Changed nick to %s", command.Args[0])
				fmt.Printf("Success: %s changed nick to %s\n", command.User.Nick, command.Args[0])
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
					b.APIEndpointKey = command.Args[1]

				case "endpoint":
					b.APIEndpoint = command.Args[1]
				}

				b.ircCon.Privmsgf(command.User.Nick, "Successfully set %s to %s\n", command.Args[0], command.Args[1])
				fmt.Printf("Success: %s set %s to %s\n", command.User.Nick, command.Args[0], command.Args[1])
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
					b.handlers.Response(channel, b.APIEndpointKey, sender)

				case "endpoint":
					b.handlers.Response(channel, b.APIEndpoint, sender)
				}

				fmt.Printf("Success: %s got %s\n", command.User.Nick, command.Args[0])
			} else {
				fmt.Printf("Missing parameters: %s tried to get nothing from something", command.User.Nick)
				return
			}

		case "admin":
			if (len(command.Args) == 2) {
				switch command.Args[0] {
				case "remove":
					if ! b.IsAdmin(command.User) {
						fmt.Printf("Insufficient permissions: %s tried to remove %s from admins\n", command.User.Nick, command.Args[1])
						return
					}

					if msg == false {
						fmt.Printf("Invalid command: %s tried to remove %s from admins\n", command.User.Nick, command.Args[1])
						return
					}

					for i, admin := range b.admins {
						if admin == command.Args[1] {

							b.admins = b.admins[:i+copy(b.admins[i:], b.admins[i+1:])]

							b.ircCon.Privmsgf(command.User.Nick, "Removed %s from admins", command.Args[1])
							b.ircCon.Privmsgf(command.Args[1], "You were removed from the admin list by %s", command.User.Nick)
							fmt.Printf("Success: %s removed %s from admins\n", command.User.Nick, command.Args[1])

							return
						}
					}

				case "add":
					if ! b.IsAdmin(command.User) {
						fmt.Printf("Insufficient permissions: %s tried to add %s to admins\n", command.User.Nick, command.Args[1])
						return
					}

					if msg == false {
						fmt.Printf("Invalid command: %s tried to add %s to admins\n", command.User.Nick, command.Args[1])
						return
					}

					for _, admin := range b.admins {
						if admin == command.Args[1] {
							return
						}
					}

					b.admins = append(b.admins, command.Args[1])
					b.ircCon.Privmsgf(command.User.Nick, "Added %s to admins", command.Args[1])
					b.ircCon.Privmsgf(command.Args[1], "You're added to the admin list by %s", command.User.Nick)
					fmt.Printf("Success: %s added %s to admins\n", command.User.Nick, command.Args[1])
				}
			} else {
				fmt.Printf("Missing parameters: %s tried to run an admin command\n", command.User.Nick)
			}

		default:
			fmt.Printf("%s is not a special command. Running default handler\n", command.Command)
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

			fmt.Printf("Sending Leet to remote API at %s\n", b.APIEndpoint)
			fmt.Println(string(leetData))

			b.PostLeet(b.APIEndpoint, leetData)
		//}
	}
}

// Check if a given *User is admin
func (b *Bot) IsAdmin(u *User) bool {
	return inArray(u.Nick, b.admins)
}