package svc

import (
	"goat-layout/internal/model"
	"goat-layout/pkg/e"
)

// Book 书籍业务
type Book struct {
	Service
}

func (b *Book) GetList() ([]model.Book, error) {
	b.log.Info("业务处理")
	// 使用store调用dao层
	books, err := b.store.Book.List()
	if err != nil {
		// 返回包装错误，包含调用栈信息
		return nil, e.New(e.DBError, err)
	}
	return books, nil
}

func (b *Book) Add(name, url string) error {
	b.log.Info("业务处理")
	// 使用store调用dao层
	err := b.store.Book.Create(name, url)
	if err != nil {
		// 返回包装错误，包含调用栈信息
		return e.New(e.DBError, err)
	}
	return nil
}
