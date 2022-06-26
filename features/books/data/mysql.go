package data

import (
	"altaproject/features/books"

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
	// result := repo.db.Joins("inner join users on users.id = books.user_id").Find(&dataBooks)
	result := repo.db.Preload("User").Find(&dataBooks)
	if result.Error != nil {
		return []books.Core{}, result.Error
	}
	// fmt.Println("databook", dataBooks[0].User.ID)
	// fmt.Println("databook", dataBooks[0].User.Name)
	return toCoreList(dataBooks), nil
}
