package types

import "time"

type User struct {
	ID        int       `json:"id"`
	Nome      string    `json:"nome"`
	Sobrenome string    `json:"sobrenome"`
	Email     string    `json:"email"`
	Senha     string    `json:"-"`
	CriadoEm  time.Time `j́son:"criadoEm"`
}

type RegisterUSerPayload struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
}
