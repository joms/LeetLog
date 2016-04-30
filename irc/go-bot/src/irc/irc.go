package irc

import (
	"github.com/thoj/go-ircevent"
	"bot"
	"log"
	"time"
)

// Bot configuration
type Config struct {
	Admins		[]string
	Server		string
	Channels 	[]string
	User		string
	Nick		string
}

var (
	ircConn *irc.Connection
	config  *Config
	b		*bot.Bot
)

// Connection event for servers
func onWelcome(e *irc.Event) {
	// Let's join some channels
	for _, channel := range config.Channels {
		ircConn.Join(channel)
	}
}

// MSG event
func onPRIVMSG(e *irc.Event) {
	t := time.Now()
	if e.Arguments[0] == config.Nick {
		e.Arguments[0] = e.Nick
	}

	// Pass our message along
	b.MessageReceived(e.Arguments[0], e.Message(), &bot.User{Nick: e.Nick}, t)
}

// Send a response somewhere to someone
func responseHandler(target string, message string, sender *bot.User) {
	channel := target
	if ircConn.GetNick() == target {
		channel = sender.Nick
	}
	ircConn.Privmsg(channel, message)
}

// Run our bot
func Run(c *Config) {
	config = c

	// Set up our IRC Connection
	ircConn = irc.IRC(config.User, config.Nick)

	// Configure the bot
	b = bot.New(
		&bot.Handlers{
			Response: responseHandler,
		},
		c.Admins,
		ircConn,
	)

	// Add callbacks
	ircConn.AddCallback("001", onWelcome)

	ircConn.AddCallback("PRIVMSG", func(event *irc.Event) {
		go func(event *irc.Event) {
			onPRIVMSG(event)
		}(event)
	});

	// Something failed
	err := ircConn.Connect(config.Server)
	if err != nil {
		log.Fatal(err)
	}
	ircConn.Loop()
}