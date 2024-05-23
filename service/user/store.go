package user

import (
	"database/sql"
	"fmt"

	"github.com/Luan-Pereira66/api-ecommerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUSerByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM usuarios WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {

			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("Usuario não encontrado")
	}

	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Nome,
		&user.Sobrenome,
		&user.Email,
		&user.Senha,
		&user.CriadoEm,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil

}

func (s *Store) CreateUser(user types.User) error {
	return nil
}
