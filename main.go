package main

import (
	"fmt"
	"tSlackMux/env"
)

func main() {
	user, err := env.Load("./config.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	NewMessages(user)
}
