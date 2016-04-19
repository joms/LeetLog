package bot

func (b *Bot) join(cmd *Cmd) {
	b.ircCon.Join(cmd.Args[0])
}

func (b *Bot) leave(cmd *Cmd) {
	b.ircCon.Part(cmd.Args[0])
}