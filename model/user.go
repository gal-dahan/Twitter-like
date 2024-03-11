package model

type User struct {
	Name         string
	Followers    map[string]*User
	FollowersNum int
	Tweets       []Tweet
}

var DatabaseUser = make(map[string]*User)
var AllUsers []*User // AllUsers is a slice of pointers to User for func GetTopInfluencers
