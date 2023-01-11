package entity

import (
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Age       int       `json:"age" db:"age"`
	Sex       string    `json:"sex" db:"sex"`
	Interests string    `json:"interests" db:"interests"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
type SecureUser struct {
	User
	Password string `db:"password"`
}
