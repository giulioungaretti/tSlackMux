package main

import "tSlackMux/env"

type channel struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	IsChannel  bool   `json:"is_channel" `
	IsArchived bool   `json :"is_archived"`
}

func GetChannels(u *env.User) []string {
	// get all the channel
	// if channel name not in env.Ingore
	// if Not archvied
	// return list of ids
	return []string{"id1", "id2"}
}
