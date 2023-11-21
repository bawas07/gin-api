package repository

import (
	"github.com/rs/zerolog"
)

type BaseRepository struct {
	log *zerolog.Logger
}

type RepoList struct {
	// Tambahkan repository lainnya sesuai kebutuhan
	UserRepo *UserRepository
}
