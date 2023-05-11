package user

import "time"

type User struct {
	Id          int       `json:"id"`
	UserName    string    `json:"user_name"`
	startedTime time.Time `json:"started_time"`
}

type Data struct {
	Users []User `json:"users"`
}
