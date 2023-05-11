package user

import (
	"fmt"
	"time"
)

func (d *Data) AddUser(name string) {
	fmt.Println("add users")
	for _, user := range d.Users {
		if user.UserName == name {
			return
		}
	}
	id := len(d.Users) + 1
	d.Users = append(d.Users, User{
		Id:          id,
		UserName:    name,
		startedTime: time.Now(),
	})
}
