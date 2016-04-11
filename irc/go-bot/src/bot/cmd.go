package bot

type Cmd struct {
	Raw string
	Channel string
	User *User
	Message string
	Command string
	RawArgs string
	Args []string
}

type User struct{
	Nick string
	Realname string
}
