package user

import (
	"../../models"
	"fmt"
)

type User struct {
	Name string
	Age  int
	State int
}

func (newUser *User) Add() error {
	fmt.Println("InService")
	user := map[string]interface{}{
		"name": newUser.Name,
		"age": newUser.Age,
		"state": newUser.State,
	}
	if err := models.AddUser(user); err != nil {
		return err
	}
	return nil
}