package database

import (
	"database/sql"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func (ur *UserRepository) GetUsers() ([]domain.User, error) {
	rows, err := ur.DB.Query("SELECT id, name, is_active, is_staff FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.IsActive, &user.IsStaff)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
