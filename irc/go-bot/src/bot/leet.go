package bot

import (
	"time"
)

// We probably have a leet to deal with
func (b *Bot) Leet(channel string, sender *User, msg string, t time.Time) *Leet {
	leet := &Leet{User: sender, Channel: channel}
	leet.Time = t.Format("2006/01/02-15:04:05.999")

	// Save time for prettier if-statements
	var h = t.Hour();
	var m = t.Minute();

	// Find the status of the leet
	if inArray(msg, b.LeetPrefix) {
		if h < 13 {
			// Empty string before 13:00
			leet.Status = 5
		} else if h > 13 {
			// Empty string after 14:00
			leet.Status = 6
		} else if h == 13 {
			if m < 37 {
				// Empty string before 13:37
				leet.Status = 5
			} else if m > 37 {
				// Empty string after 13:37
				leet.Status = 6
			} else if h == 37 {
				// Empty string on 13:37
				leet.Status = 0
			}
		}
	} else {
		if h < 13 {
			// Text before 13:00
			leet.Status = 1
		} else if h > 13 {
			// Text after 14:00
			leet.Status = 4
		} else if h == 13 {
			if m < 37 {
				// Text before 13:37
				leet.Status = 1
			} else if m > 37 {
				// Text after 13:37
				leet.Status = 4
			} else if h == 37 {
				// http://i.imgur.com/SWzmsBF.gif
				leet.Status = 2
			}
		}
	}

	return leet
}