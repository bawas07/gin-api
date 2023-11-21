package service

import (
	"github.com/rs/zerolog"
)

type BaseService struct {
	log *zerolog.Logger
}

type ServiceList struct {
	User *UserService
	// Tambahkan controller lainnya sesuai kebutuhan
}
