package kylpynallet

// User is a user
type User struct {
	ID             int
	Name, Location string
}

/*
// Greetings greet ppl
func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}
*/

// Player is a player of games
type Player struct {
	User
	GameID int
}
