package response

import (
	"altaproject/features/books"
	"time"
)

type Book struct {
	ID          int       `json:"id" form:"id"`
	Title       string    `json:"title" form:"title"`
	Author      string    `json:"author" form:"author"`
	Publisher   string    `json:"publisher" form:"publisher"`
	ISBN        string    `json:"isbn" form:"isbn"`
	Category    string    `json:"category" form:"category"`
	Price       uint      `json:"price" form:"price"`
	Stock       uint      `json:"stock" form:"stock"`
	BookPage    string    `json:"book_page" form:"book_page"`
	Sold        uint      `json:"sold" form:"sold"`
	Description string    `json:"description" form:"description"`
	ImageURL    string    `json:"image_url" form:"image_url"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	User        User      `json:"user" form:"user"`
}

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
}

func FromCore(data books.Core) Book {
	return Book{
		ID:          data.ID,
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
		ImageURL:    data.ImageURL,
		CreatedAt:   data.CreatedAt,
		User: User{
			ID:       data.User.ID,
			UserName: data.User.UserName,
		},
	}
}

func FromCoreList(data []books.Core) []Book {
	result := []Book{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
