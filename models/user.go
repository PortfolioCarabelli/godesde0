package models

import (
	"time"
)

//Modelos solo son estructuras o definciones

type User struct {
	Id       int
	Name     string
	Createat time.Time
	Status   bool
}

func (user *User) AddUser(id int, name string, createdAt time.Time, status bool) {

	user.Id = id
	user.Name = name
	user.Createat = createdAt
	user.Status = status

}
