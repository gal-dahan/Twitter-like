package main

import (
	"fmt"

	"github.com/gal-dahan/Twitter-like/controllers"
)

func main() {
	fmt.Println("Start!")
	userGal := "Gal Dahan"
	userBar := "Bar Dvir"
	userTomer := "Tomer"
		err := controllers.CreateUser(userGal)
	if err != nil {
		fmt.Println("Error creating user:", err)
	} else {
		fmt.Printf("User %v created successfully\n", userGal)
	}
	err = controllers.CreateUser(userBar)
	if err != nil {
		fmt.Println("Error creating user:", err)
	} else {
		fmt.Printf("User %v created successfully\n", userBar)
	}
	err = controllers.CreateUser(userTomer)
	if err != nil {
		fmt.Println("Error creating user:", err)
	} else {
		fmt.Printf("User %v created successfully\n", userTomer)
	}

	err = controllers.FollowUser(userGal, userBar)
	if err != nil {
		fmt.Println("Error following user:", err)
	} else {
		fmt.Printf("User %v followed successfully by %v\n", userBar, userGal)
	}
	
	err = controllers.PostTweet(userGal, "Hello, Twitter!")
	if err != nil {
		fmt.Println("Error posting tweet:", err)
	} else {
		fmt.Println("Tweet posted successfully")
	}
	
	topInfluencers, err := controllers.GetTopInfluencers(2)
	if err != nil {
		fmt.Println("Error getting top influencers:", err)
	} else {
		fmt.Println("Top Influencers:")
		for _, user := range topInfluencers {
			fmt.Println("Name:", user.Name, "| Followers:", user.FollowersNum)
		}
	}

	err = controllers.UnfollowUser(userGal, userBar)
	if err != nil {
		fmt.Println("Error unfollowing user:", err)
	} else {
		fmt.Printf("User %v unfollowed successfully by %v\n", userBar, userGal)
	}

	err = controllers.PostTweet(userGal, "Hello, from Alma Security!")
	if err != nil {
		fmt.Println("Error posting tweet:", err)
	} else {
		fmt.Println("Tweet posted successfully")
	}
	feed, err := controllers.GetUserFeed(userGal)
	if err != nil {
		fmt.Println("Error getting user feed:", err)
	} else {
		fmt.Printf("Feed of User %v:%v\n", userGal, feed)
	}
	users, err := controllers.GetMutualFollowers("Gal Dahan", "Bar Dvir")
	if err != nil {
		fmt.Println("Error getting mutual followers:", err)
	} else {
		fmt.Println("Mutual followers:", users)
	}

	fmt.Println("end!")

}