package modelos

import (
	"time"
)
type User struct {
	Id int
	Name string
	LastName string
	Email string
	Password string
	CreatedAt time.Time
	Status bool
}
func (usuario *User) AddUser(id int, name string, lastName string, email string, password string, createdAt time.Time, Status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.LastName = lastName
	usuario.Email = email
	usuario.Password = password
	usuario.CreatedAt = createdAt
	usuario.Status = Status

}


