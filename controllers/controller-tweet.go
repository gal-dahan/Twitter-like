package controllers

import (
	"errors"
	"fmt"

	"github.com/gal-dahan/Twitter-like/model"
)

func FollowUser(userName string, followerName string) error {
	fmt.Println("FollowUser")
	user, err := GetUser(userName)
	if err != nil {
		return errors.New("userName not found")
	}
	follower, err := GetUser(followerName)
	if err != nil {
		return errors.New("followerName not found")
	}
	user.Followers = append(user.Followers, follower)
	user.FollowersNum++
	return nil
}
func UnfollowUser(userName string, followerName string) error {
	user, err := GetUser(userName)
	if err != nil {
		return err
	}
	follower, err := GetUser(followerName)
	if err != nil {
		return err
	}
	for i, u := range user.Followers {
		if u.Name == follower.Name {
			user.Followers = append(user.Followers[:i], user.Followers[i+1:]...) //remove the follower from the followers list
			user.FollowersNum--
			return nil
		}
	}
	return nil
}

func PostTweet(userName string, tweetContent string) error {
	fmt.Println("PostTweet")
	user, err := GetUser(userName)
	if err != nil {
		return errors.New("userName not found")
	}
	user.Tweets = append(user.Tweets, model.Tweet{Post: tweetContent})
	return nil
}

func GetUserFeed(userName string) ([]model.Tweet, error) {
	user, err := GetUser(userName)
	if err != nil {
		return nil, err
	}

	return user.Tweets, nil
}
func populateFollowersMap(user *model.User) map[string]bool {
	followersMap := make(map[string]bool)
	for _, follower := range user.Followers {
		followersMap[follower.Name] = true
	}
	return followersMap
}

func GetMutualFollowers(userName1 string, userName2 string) ([]*model.User, error) {
	user1, err := GetUser(userName1)
	if err != nil {
		return nil, err
	}
	user2, err := GetUser(userName2)
	if err != nil {
		return nil, err
	}

	followersMap1 := populateFollowersMap(user1)
	var mutualFollowers []*model.User
	for _, follower := range user2.Followers {
		if followersMap1[follower.Name] {
			mutualFollowers = append(mutualFollowers, follower)
		}
	}

	return mutualFollowers, nil
}
func GetTopInfluencers(n int) ([]*model.User, error) {
	if n > len(model.AllUsers) {
		return nil, errors.New("n is bigger than the number of users")
	}
	return model.AllUsers[:n], nil
}