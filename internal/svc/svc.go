package svc

import (
	"goat/internal/model"

	"go.uber.org/zap"
)

type Service struct {
	store *model.Store
	log   *zap.Logger
}

func New(store *model.Store, log *zap.Logger) *Service {
	return &Service{
		store: store,
		log:   log.Named("svc"),
	}
}
