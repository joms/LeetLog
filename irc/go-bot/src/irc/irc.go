package irc

import (
	"github.com/thoj/go-ircevent"
	"log"
	"bot"
	"time"
)

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


func onWelcome(e *irc.Event) {
	for _, channel := range config.Channels {
		ircConn.Join(channel)
	}
}

func onPRIVMSG(e *irc.Event) {
	t := time.Now()
	b.MessageReceived(e.Arguments[0], e.Message(), &bot.User{Nick: e.Nick}, t)
}

func responseHandler(target string, message string, sender *bot.User) {
	channel := target
	if ircConn.GetNick() == target {
		channel = sender.Nick
	}
	ircConn.Privmsg(channel, message)
}

func Run(c *Config) {
	config = c

	ircConn = irc.IRC(config.User, config.Nick)

	ircConn.AddCallback("001", onWelcome)

	ircConn.AddCallback("PRIVMSG", func(event *irc.Event) {
		go func(event *irc.Event) {
			onPRIVMSG(event)
		}(event)
	});

	err := ircConn.Connect(config.Server)
	if err != nil {
		log.Fatal(err)
	}
	ircConn.Loop()
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}