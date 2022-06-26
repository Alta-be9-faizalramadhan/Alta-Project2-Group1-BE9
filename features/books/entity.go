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
	BookPage    uint
	Sold        uint
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User
}

type User struct {
	ID   int
	Name string
}

type Business interface {
	GetAllBook(limit, offset uint) (data []Core, err error)
}

type Data interface {
	SelectAllBook(limit, offset uint) (data []Core, err error)
}
