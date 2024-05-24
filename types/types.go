package types

import "time"

type UserStore interface {
	GetUSerByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int       `json:"id"`
	Nome      string    `json:"nome"`
	Sobrenome string    `json:"sobrenome"`
	Email     string    `json:"email"`
	Senha     string    `json:"-"`
	CriadoEm  time.Time `json:"criadoEm"`
}

type RegisterUSerPayload struct {
	Nome      string `json:"nome" validate:"required"`
	Sobrenome string `json:"sobrenome" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Senha     string `json:"senha" validate:"required,min=3,max=130"`
}
