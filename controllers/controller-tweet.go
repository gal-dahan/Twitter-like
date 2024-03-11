package controllers

import (
	"errors"
	"fmt"
	"github.com/gal-dahan/Twitter-like/model"
)
func isFollowerAlready(user model.User, followerName string) bool {
	_, exists := user.Followers[followerName]
	return exists
}

func FollowUser(userName string, followerName string) error {
	user, err := GetUser(userName)
	if err != nil {
		return errors.New("userName not found")
	}
	follower, err := GetUser(followerName)
	if err != nil {
		return errors.New("followerName not found")
	}
	if user.Name == follower.Name {
		return errors.New("user can't follow himself")
	}
	if isFollowerAlready(*user, followerName) {
		return fmt.Errorf("%v is already following %v", followerName, userName)
	}

	user.Followers[followerName] = follower
	user.FollowersNum++
	fmt.Printf("User %v followed successfully by %v\n", userName, followerName)
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
	if user.Name == follower.Name {
		return errors.New("user can't unfollow himself")
	}
	if !isFollowerAlready(*user, followerName) {
		return fmt.Errorf("%v is not following %v", followerName, userName)
	}
	delete(user.Followers, followerName)
	user.FollowersNum--
	fmt.Printf("User %v unfollowed successfully by %v\n", userName, followerName)

	return nil
}

func PostTweet(userName string, tweetContent string) error {
	user, err := GetUser(userName)
	if err != nil {
		return errors.New("userName not found")
	}
	newTweet := model.Tweet{
		User: *user,
		Post: tweetContent,
	}
	user.Tweets = append(user.Tweets, newTweet)
	fmt.Println("Tweet posted successfully by", userName, ":", tweetContent)

	return nil
}

func GetUserFeed(userName string) ([]model.Tweet, error) {
	user, err := GetUser(userName)
	if err != nil {
		return nil, errors.New("The user name "+userName+" not found")
	}

	return user.Tweets, nil
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

	var mutualFollowers []*model.User
	for _, follower := range user2.Followers {
		if _, exists := user1.Followers[follower.Name]; exists {
			mutualFollowers = append(mutualFollowers, follower)
		}
	}

	return mutualFollowers, nil
}

func GetTopInfluencers(n int) ([]*model.User, error) {
	if n > len(model.AllUsers) {
		return nil, fmt.Errorf("%d is bigger than the number of users", n)
	}
	return model.AllUsers[:n], nil
}