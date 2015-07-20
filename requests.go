package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tSlackMux/env"
)

type SlacLResp struct {
	Status bool `json:"ok"`
	Resp   interface{}
}

/*SlackChannel
 *Unread_count is a full count of visible messages that the calling user has yet to read. unread_count_display is a count of messages that the calling user has yet to read that matter to them (this means it excludes things like join/leave messages)
 */
type SlackChannel struct {
	UnreadCount        int `json: "unread_count"`
	UndearCountDisplay int `json: "unread_count_display"`
}

func NewMessages(u *env.User) {
	query := "https://slack.com/api/channels.info?"
	parameters := fmt.Sprintf("token=%s&channel=%s", u.Token, "C06Q4G3K3")
	query += parameters
	fmt.Printf("%s\n", query)
	resp, err := http.Get(query)
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	slackChannel := SlacLResp{}
	slackChannel.Resp = new(SlackChannel)
	err = json.Unmarshal(body, &slackChannel)
	// TODO handle error
	fmt.Printf("%+v", slackChannel.Resp)
}
