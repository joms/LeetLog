package bot

import (
	"time"
	"fmt"
	"strconv"
	"os"
)

func (b *Bot) leet(c *Cmd, t time.Time) {
	l := &Leet{User: c.User, Channel: c.Channel, Message: c.Message}
	l.Time = t.Format("2006/01/02-15:04:05.999")

	var h = t.Hour();
	var m = t.Minute();

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

	var text = l.Time +" "+ l.Channel +" "+ strconv.Itoa(l.Status) +" "+ l.User.Nick +" "+ l.Message +"\n"

	f, err := os.OpenFile("halla.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}

	fmt.Println(err)
}
