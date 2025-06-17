package modelos

import "time"

// estructuras
type User struct {
	Id       int 
	Name     string
	CreateAt time.Time//en que fecha fue creado...
	Status   bool
}
//el profe uso this.Id, etc..User es un puntero 
func (u *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	u.Id = id
	u.Name = name
	u.CreateAt = createdAt
	u.Status = status
}
