package bot

import (
	"regexp"
	"strings"
)

var (
	re = regexp.MustCompile("\\s+") // Matches one or more spaces
)

// Parse incoming message into useful data
func (b *Bot) Parse(s string, channel string, user *User, msg bool) *Cmd {
	c := &Cmd{Raw: s}

	if !strings.HasPrefix(s, b.CmdPrefix) && msg == false {
		return nil
	}

	c.Channel = strings.TrimSpace(channel)
	c.User = user


	// Trim the prefix and extra spaces
	c.Message = strings.TrimPrefix(s, b.CmdPrefix)
	c.Message = strings.TrimSpace(c.Message)

	// check if we have the command and not only the prefix
	if c.Message == "" {
		return nil
	}

	// get the command
	pieces := strings.SplitN(c.Message, " ", 2)
	c.Command = pieces[0]

	if len(pieces) > 1 {
		// get the arguments and remove extra spaces
		c.RawArgs = pieces[1]
		c.Args = strings.Split(removeExtraSpaces(c.RawArgs), " ")
	}

	return c
}

// Remove spaces from string
func removeExtraSpaces(args string) string {
	return re.ReplaceAllString(strings.TrimSpace(args), " ")
}

// Check if a string is in an array
func inArray(needle string, haystack []string) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}