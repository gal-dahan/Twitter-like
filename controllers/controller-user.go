package controllers

import (
	"errors"
	"fmt"
	"github.com/gal-dahan/Twitter-like/model"
)

func CreateUser(name string) error {
	if _, exists := model.DatabaseUser[name]; exists {
		return errors.New("The user with "+name+ " already exists")
	}

	user := &model.User{
		Name:         name,
		Followers:    make(map[string]*model.User),
		FollowersNum: 0,
		Tweets:       []model.Tweet{},
	}

	model.DatabaseUser[name] = user
	model.AllUsers = append(model.AllUsers, user) 
	fmt.Printf("User %v created successfully!\n", name)

	return nil
}

func GetUser(name string) (*model.User, error) {
	user, exists := model.DatabaseUser[name]
	if !exists {
		return nil, errors.New("The user name "+name+" not found")
	}

	return user, nil
}

func UpdateUser(name string, newName string) error {
	user, err := GetUser(name)
	if err != nil {
		return err
	}

	user.Name = newName
	fmt.Printf("User updated name %v to %v successfully!\n", name , newName)
	return nil
}

func DeleteUser(name string) error {
	_, err  := GetUser(name)
	if err != nil {
		return errors.New("user not found")
	}
	for i, user := range model.AllUsers {
		if user.Name == name {
			model.AllUsers[i] = model.AllUsers[len(model.AllUsers)-1]
			model.AllUsers = model.AllUsers[:len(model.AllUsers)-1]
			break
		}
	}

	delete(model.DatabaseUser, name) //To delete a key from a map(DatabaseUser), Go built-in delete() function.
	fmt.Printf("User %v deleted successfully!\n", name)
	return nil
}


