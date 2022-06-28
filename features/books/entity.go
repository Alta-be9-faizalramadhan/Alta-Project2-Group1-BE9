package books

import (
	"time"
)

type Core struct {
	ID          int
	Title       string
	Author      string
	Publisher   string
	ISBN        string
	Category    string
	Price       uint
	Stock       uint
	BookPage    string
	Sold        uint
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User
}

type User struct {
	ID       int
	UserName string
}

type Business interface {
	GetAllBook(limit, offset uint) (data []Core, err error)
	CreateNewBook(data Core) (row int, err error)
	GetBookById(id int) (data Core, err error)
	UpdatedBook(id int, data Core) (row int, err error)
	SoftDeleteBook(id int) (row int, err error)
	SelectBookByCategory(category string) (data []Core, err error)
	SelectBookByUserId(id int) (data []Core, err error)
}

type Data interface {
	SelectAllBook(limit, offset uint) (data []Core, err error)
	InsertNewBook(data Core) (row int, err error)
	SelectBookById(id int) (data Core, err error)
	UpdatedBook(id int, data Core) (row int, err error)
	SoftDeleteBook(id int) (row int, err error)
	SelectBookByCategory(category string) (data []Core, err error)
	SelectBookByUserId(id int) (data []Core, err error)
}
