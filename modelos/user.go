package modelos

import "time"

// estructuras
type User struct {
	Id       int
	Name     string
	CreateAt time.Time
	Status   bool
}

func (this User) AddUser(id int, name string, createdAt time.Time, status bool) {
	this.Id = id
	this.Name = name
	this.CreateAt = createdAt
	this.Status = status
}
