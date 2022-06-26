package request

import "altaproject/features/books"

type Book struct {
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
	UserId      int    `json:"user_id" form:"user_id"`
}

func ToCore(req Book) books.Core {
	return books.Core{
		Title:       req.Title,
		Author:      req.Author,
		Publisher:   req.Publisher,
		ISBN:        req.ISBN,
		Category:    req.Category,
		Price:       req.Price,
		Stock:       req.Stock,
		BookPage:    req.BookPage,
		Sold:        req.Sold,
		Description: req.Description,
		User: books.User{
			ID: req.UserId,
		},
	}
}
