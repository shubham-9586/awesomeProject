package appLogic

import (
	"awesomeProject/entity"
	"fmt"
)

var CurrentUser = []entity.User{}

func GetNewUser(ID string, name string, gender string, age int, state string, district string) entity.User {
	return entity.User{
		UserID:   ID,
		Name:     name,
		Gender:   gender,
		Age:      age,
		State:    state,
		District: district,
	}
}

func AddUserToDB(newUser entity.User) {
	if newUser.Age >= 18 {
		CurrentUser = append(CurrentUser, newUser)
		fmt.Println("User has been added successfully")
	} else {
		fmt.Println("User must be an adult")
	}
}
