package env

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type User struct {
	Usr   string
	Token string
}

func read(path string) (fileBytes []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileBytes, err = ioutil.ReadAll(file)
	return
}

func Load(path string) (usr *User, err error) {
	file, err := read(path)
	if err != nil {
		return nil, err
	}
	usr = new(User)
	err = json.Unmarshal(file, &usr)
	return
}
