package model
type Tweet struct {
	User  User
	Post  string
}
var DatabaseTweet = make(map[string]*Tweet)
