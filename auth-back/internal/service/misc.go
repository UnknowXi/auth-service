package service

import (
	"auth/auth-back/internal/domain/auth"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadUsers() (auth.UsersJson, error) {
	usersJson := auth.UsersJson{}

	file, err := ioutil.ReadFile("user.json")
	if err != nil {
		return usersJson, err
	}

	err = json.Unmarshal(file, &usersJson)
	if err != nil {
		fmt.Printf("Error %s \n", err.Error())
		return usersJson, err
	}

	return usersJson, nil
}
