package repositories

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/tinycloudtv/authn-service/app/internal/errors"
	"github.com/tinycloudtv/authn-service/app/internal/models"
	"time"
)

type UserTokensRepository struct {
	db DatabaseConnect
}

func (repo UserTokensRepository) Init() {
	repo.db = DatabaseConnect{}
}

func (repo UserTokensRepository) Get(token string) (models.UserToken, error) {
	repo.db.Open()
	defer repo.db.Close()

	results, _ := repo.db.DB.Query("SELECT * FROM user_tokens WHERE token = ?", token)

	for results.Next() {
		var token models.UserToken
		err := results.Scan(&token.Id, &token.UserId, &token.Token, &token.Expiration, &token.Expired)

		if err == nil {
			return token, nil
		}
	}

	return models.UserToken{}, &errors.AuthFailedError{}
}

func (repo UserTokensRepository) Create(user models.User, token string, expiration time.Time) models.UserToken {
	id := uuid.New()

	userToken := models.UserToken{
		Id:         id.String(),
		UserId:     user.Id,
		Token:      token,
		Expiration: expiration,
		Expired:    false,
	}

	repo.db.Open()
	defer repo.db.Close()

	insert, _ := repo.db.DB.Query("REPLACE INTO user_tokens VALUES(?, ?, ?, ?, ?)", userToken.Id, userToken.UserId, userToken.Token, userToken.Expiration, userToken.Expired)
	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {

		}
	}(insert)

	return userToken
}
