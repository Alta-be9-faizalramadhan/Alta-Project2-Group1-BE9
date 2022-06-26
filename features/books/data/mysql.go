package data

import (
	"altaproject/features/books"
	"fmt"

	"gorm.io/gorm"
)

type mysqlBookRepository struct {
	db *gorm.DB
}

func NewBookRepository(conn *gorm.DB) books.Data {
	return &mysqlBookRepository{
		db: conn,
	}
}

func (repo *mysqlBookRepository) SelectAllBook(limit, offset uint) (response []books.Core, err error) {
	var dataBooks []Book
	result := repo.db.Preload("User").Find(&dataBooks)
	if result.Error != nil {
		return []books.Core{}, result.Error
	}
	return toCoreList(dataBooks), nil
}

func (repo *mysqlBookRepository) InsertNewBook(input books.Core) (row int, err error) {
	book := fromCore(input)
	result := repo.db.Create(&book)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(result.RowsAffected), nil
}