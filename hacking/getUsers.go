package hacking

import (
	"bufio"
	"io"
	"os"
	"os/user"
	"strings"
)

type User struct {
	Username    string
	HomeDir     string
	GroupID     string
	DisplayName string
}

func GetUsers() ([]User, error) {
	var passWdUsers []string
	var users []User

	file, err := os.Open("/etc/passwd")

	if err != nil {
		return nil, err
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if equal := strings.Index(line, "#"); equal < 0 {
			lineSlice := strings.FieldsFunc(line, func(divide rune) bool {
				return divide == ':'
			})

			if len(lineSlice) > 0 {
				passWdUsers = append(passWdUsers, lineSlice[0])
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

	}

	for _, name := range passWdUsers {
		usr, err := user.Lookup(name)
		if err != nil {
			panic(err)
		}
		users = append(users, User{usr.Username, usr.HomeDir, usr.Gid, usr.Name})
	}
	return users, nil
}
