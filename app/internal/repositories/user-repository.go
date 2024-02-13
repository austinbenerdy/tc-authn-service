package repositories

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/tinycloudtv/authn-service/app/internal/errors"
	"github.com/tinycloudtv/authn-service/app/internal/models"
)

type UserRepository struct {
	db DatabaseConnect
}

func (repo UserRepository) Init() {
	repo.db = DatabaseConnect{}
}

func (repo UserRepository) GetUser(email string) (models.User, error) {
	repo.db.Open()
	defer repo.db.Close()

	results, _ := repo.db.DB.Query("SELECT * FROM users WHERE email = ?", email)

	for results.Next() {
		var user models.User
		err := results.Scan(&user.Id, &user.Email, &user.Password)
		if err == nil {
			return user, nil
		}
	}

	return models.User{}, &errors.AuthFailedError{}
}

func (repo UserRepository) CreateUser(email string, password string) models.User {
	id := uuid.New()

	user := models.User{
		Id:       id.String(),
		Email:    email,
		Password: password,
	}

	repo.db.Open()
	defer repo.db.Close()

	insert, _ := repo.db.DB.Query("REPLACE INTO users VALUES(?, ?, ?)", user.Id, user.Email, user.Password)
	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {

		}
	}(insert)

	return user
}
