package data

import (
	"altaproject/features/books"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `json:"title" form:"title"`
	Author      string `json:"author" form:"author"`
	Publisher   string `json:"publisher" form:"publisher"`
	ISBN        string `json:"isbn" form:"isbn"`
	Category    string `json:"category" form:"category"`
	Price       uint   `json:"price" form:"price"`
	Stock       uint   `json:"stock" form:"stock"`
	BookPage    string `json:"book_page" form:"book_page"`
	Sold        uint   `json:"sold" form:"sold"`
	Description string `json:"description" form:"description"`
	UserID      uint   `json:"user_id" form:"user_id"`
	User        User
}

type User struct {
	gorm.Model
	UserName string `json:"user_name" form:"user_name"`
	Books    []Book
}

func (data *Book) toCore() books.Core {
	return books.Core{
		ID:          int(data.ID),
		Title:       data.Title,
		Author:      data.Author,
		Publisher:   data.Publisher,
		ISBN:        data.ISBN,
		Category:    data.Category,
		Price:       data.Price,
		Stock:       data.Stock,
		BookPage:    data.BookPage,
		Sold:        data.Sold,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		User: books.User{
			ID:       int(data.User.ID),
			UserName: data.User.UserName,
		},
	}
}

func toCoreList(data []Book) []books.Core {
	result := []books.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core books.Core) Book {
	return Book{
		Title:       core.Title,
		Author:      core.Author,
		Publisher:   core.Publisher,
		ISBN:        core.ISBN,
		Price:       core.Price,
		Stock:       core.Stock,
		Category:    core.Category,
		BookPage:    core.BookPage,
		Sold:        core.Sold,
		Description: core.Description,
		UserID:      uint(core.User.ID),
	}
}
