package bot

import (
	"github.com/thoj/go-ircevent"
	"log"
)

type Config struct {
	Admins		[]string
	Server		string
	Channels 	[]string
	User		string
	Nick		string
	Prefix 		string
}

var (
	ircConn *irc.Connection
	config  *Config
)


func onWelcome(e *irc.Event) {
	for _, channel := range config.Channels {
		ircConn.Join(channel)
	}
}

func onPRIVMSG(e *irc.Event) {
	if (string(e.Message()[0]) == " ") {
		ircConn.Privmsg(e.Arguments[0], e.User +": You can't even leet");
	} else if (string(e.Message()[0]) == config.Prefix && stringInSlice(e.User, config.Admins) == true) {
		ircConn.Privmsg(e.Arguments[0], e.User +": You can't even admin");
	} else {
		ircConn.Privmsg(e.Arguments[0], e.User +": Booooring");
	}
}

func Run(c *Config) {
	config = c

	ircConn = irc.IRC(config.User, config.Nick)

	ircConn.AddCallback("001", onWelcome)
	ircConn.AddCallback("PRIVMSG", onPRIVMSG)

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