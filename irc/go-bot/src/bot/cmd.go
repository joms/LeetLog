package bot

import (
	"log"
	"fmt"
)

// Input command result
type Cmd struct {
	Raw string
	Channel string
	User *User
	Message string
	Command string
	RawArgs string
	Args []string
}

// User structure
type User struct {
	Nick string
	Realname string
}

// Leet structure
type Leet struct {
	User *User
	Time string
	Status int
	Channel string
}

// Command structure
type Command struct {
	Cmd         string
	CmdFunc   activeCmdFunc
	Description string
	ExampleArgs string
	Admin bool
	Msg bool
}


type activeCmdFunc func(cmd *Cmd) (string, error)

var commands = make(map[string]*Command)

const errorExecutingCommand = "Error executing %s: %s"


// RegisterCommand adds a new command to the bot.
// The command(s) should be registered in the Init() func of your package
// command: String which the user will use to execute the command, example: reverse
// decription: Description of the command to use in !help, example: Reverses a string
// exampleArgs: Example args to be displayed in !help <command>, example: string to be reversed
// cmdFunc: Function which will be executed. It will received a parsed command as a Cmd value
func RegisterCommand(command, description, exampleArgs string, cmdFunc activeCmdFunc, admin bool, msg bool) {
	commands[command] = &Command{
		Cmd:         command,
		CmdFunc:   cmdFunc,
		Description: description,
		ExampleArgs: exampleArgs,
		Admin: admin,
		Msg: msg,
	}
}

// Run the command
func (b *Bot) handleCmd(c *Cmd) {
	cmd := commands[c.Command]

	if cmd.Admin == true && b.IsAdmin(c.User) == false {
		return;
	}

	if cmd.Msg == true && c.User.Nick != c.Channel {
		return;
	}

	if cmd == nil {
		log.Printf("Command not found %v", c.Command)
		return
	}

	message, err := cmd.CmdFunc(c)
	b.checkCmdError(err, c)
	if message != "" {
		b.handlers.Response(c.Channel, message, c.User)
	}
}

// Check for command errors
func (b *Bot) checkCmdError(err error, c *Cmd) {
	if err != nil {
		errorMsg := fmt.Sprintf(errorExecutingCommand, c.Command, err.Error())
		log.Printf(errorMsg)
		b.handlers.Response(c.Channel, errorMsg, c.User)
	}
}