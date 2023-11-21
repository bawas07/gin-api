package repository

import (
	"gin-api/cmd/app/model"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type UserRepository struct {
	BaseRepository
	db *sqlx.DB
}

func NewUserRepository(
	log *zerolog.Logger,
	db *sqlx.DB,
) *UserRepository {
	return &UserRepository{
		BaseRepository: BaseRepository{log: log},
		db:             db,
	}
}

func (r *UserRepository) GetAllUser() []model.User {
	users := []model.User{}
	query := "select * from users"
	err := r.db.Select(&users, query)
	if err != nil {
		panic(err)
	}

	return users

}
