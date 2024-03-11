package main

import (
	"fmt"
	"github.com/gal-dahan/Twitter-like/controllers"
)

func main() {
	fmt.Println("Start manage a Twitter-like application!")
	userNames := []string{"Gal Dahan", "Bar Dvir", "Tomer"} //For example get from the API or from the Frontend
	userGal := userNames[0]
	userBar := userNames[1]
	userTomer := userNames[2]

	fmt.Println("-----------------")
	for _, userName := range userNames {
		err := controllers.CreateUser(userName)
		if err != nil {
			fmt.Printf("Error creating user %s: %v\n", userName, err)
		}
	}

	fmt.Println("-----------------")

	err := controllers.FollowUser(userGal, userTomer)
	if err != nil {
		fmt.Println("Error following user:", err)
	} 

	fmt.Println("-----------------")

	err = controllers.FollowUser(userBar, userTomer)
	if err != nil {
		fmt.Println("Error following user:", err)
	} 

	fmt.Println("-----------------")

	err = controllers.UpdateUser(userTomer, "Tomer Lavi")
	if err != nil {
		fmt.Println("Error updating user:", err)
	}
		
	fmt.Println("-----------------")

	err = controllers.FollowUser(userGal, userBar)
	if err != nil {
		fmt.Println("Error following user:", err)
	} 
		
	fmt.Println("-----------------")

	
	err = controllers.PostTweet(userGal, "Hello, Twitter!")
	if err != nil {
		fmt.Println("Error posting tweet:", err)
	}
		
	fmt.Println("-----------------")
	topOfFollowers:=3
	topInfluencers, err := controllers.GetTopInfluencers(topOfFollowers)
	if err != nil {
		fmt.Println("Error getting top influencers:", err)
	} else {
		fmt.Printf("Top %d Influencers:\n", topOfFollowers)
		for _, user := range topInfluencers {
			fmt.Println("Name:", user.Name, "| Followers:", user.FollowersNum)
		}
	}

	fmt.Println("-----------------")

	mutualFollowers, err := controllers.GetMutualFollowers(userBar, userGal)
	if err != nil {
		fmt.Println("Error getting mutual followers:", err)
	} else {
		fmt.Printf("Mutual Followers of %s and %s:\n", userBar, userGal)
		for _, follower := range mutualFollowers {
			fmt.Printf(" Name: %s\n",follower.Name)
		}
		}
	
	fmt.Println("-----------------")

	err = controllers.UnfollowUser(userGal, userBar)
	if err != nil {
		fmt.Println("Error unfollowing user:", err)
	} 
	
	fmt.Println("-----------------")

	err = controllers.PostTweet(userGal, "Hello, from Alma Security!")
	if err != nil {
		fmt.Println("Error posting tweet:", err)
	}
	
	fmt.Println("-----------------")

	feed, err := controllers.GetUserFeed(userGal)
	if err != nil {
		fmt.Println("Error getting user feed:\n", err)
	} else {
		for _, tweet := range feed {
			fmt.Printf("Feed of User %v: %v\n", tweet.User.Name, tweet.Post)
		}
		}
	
	fmt.Println("-----------------")

	err = controllers.DeleteUser(userGal)
	if err != nil {
		fmt.Println("Error deleting user:", err)
	}
	
	fmt.Println("-----------------")

	fmt.Println("End!")

}