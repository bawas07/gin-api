package configuration

import (
	"gin-api/cmd/app/repository"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

func LoadRepository(log *zerolog.Logger, db *sqlx.DB) repository.RepoList {
	userRepo := repository.NewUserRepository(log, db)
	return repository.RepoList{
		UserRepo: userRepo,
	}
}
