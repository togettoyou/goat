package dao

import (
	"goat/internal/model"

	"gorm.io/gorm"
)

func newBook1(db *gorm.DB) *book1 {
	return &book1{db}
}

func newBook2() *book2 {
	return &book2{}
}

type book1 struct {
	db *gorm.DB
}

func (b *book1) List() ([]model.Book, error) {
	var books []model.Book
	if err := b.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

type book2 struct {
}

func (b *book2) List() ([]model.Book, error) {
	return []model.Book{
		{
			Name: "Go语言圣经（中文版）",
			Url:  "https://books.studygolang.com/gopl-zh/",
		},
		{
			Name: "Go语言高级编程(Advanced Go Programming)",
			Url:  "https://chai2010.cn/advanced-go-programming-book/",
		},
	}, nil
}
