package shoppingcartdetails

import "time"

type Core struct {
	ID              int
	Book            Book
	QuantityBuyBook uint
	TotalPriceBook  uint
	ShoppingCart    ShoppingCart
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Book struct {
	ID    int
	Title string
	Price uint
}

type ShoppingCart struct {
	ID            int
	UserID        uint
	TotalQuantity uint
	TotalPrice    uint
	Status        string
}

type Business interface {
	GetAllCartDetails(idCart, limit, offset int) (data []Core, err error)
	InsertCartDetails(data Core) (row int, err error)
	DeleteCartDetails(idCart int, idBook int) (row int, err error)
	UpdateCartDetails(idCart int, idBook int, data Core) (row int, err error)
}

type Data interface {
	SelectAllCartDetails(idCart, limit, offset int) (data []Core, err error)
	SelectCartDetail(idCart int, idBook int) (data Core, err error)
	InsertCartDetails(data Core) (row int, err error)
	DeleteCartDetails(idCart int, idBook int) (row int, err error)
	PutCartDetails(idCart int, idBook int, data Core) (row int, err error)
	IsBookNotInCartDetail(idBook int, idCart int) (cond bool, data Core)
	FindIDCart(idUser int) (idCart int, err error)
}
