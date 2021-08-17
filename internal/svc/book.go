package svc

import (
	"goat/internal/model"
	"goat/pkg/e"
)

func (s *Service) GetBookList() ([]model.Book, error) {
	s.log.Info("业务处理")
	books, err := s.store.Book.List()
	if err != nil {
		return nil, e.New(e.DBError, err)
	}
	return books, nil
}
