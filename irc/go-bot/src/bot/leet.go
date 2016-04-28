package bot

import (
	"time"
)

// We probably have a leet to deal with
func (b *Bot) Leet(channel string, sender *User, msg string, t time.Time) {
	l := &Leet{User: sender, Channel: channel, Message: msg}
	l.Time = t.Format("2006/01/02-15:04:05.999")

	var h = t.Hour();
	var m = t.Minute();

	// Is it valid?
	if l.Message == " " {
		if h < 13 {
			l.Status = 5
		} else if h > 13 {
			l.Status = 6
		} else if h == 13 {
			if m < 37 {
				l.Status = 5
			} else if m > 37 {
				l.Status = 6
			} else if h == 37 {
				l.Status = 0
			}
		}
	} else {
		if h < 13 {
			l.Status = 1
		} else if h > 13 {
			l.Status = 4
		} else if h == 13 {
			if m < 37 {
				l.Status = 1
			} else if m > 37 {
				l.Status = 4
			} else if h == 37 {
				l.Status = 2
			}
		}
	}
}
