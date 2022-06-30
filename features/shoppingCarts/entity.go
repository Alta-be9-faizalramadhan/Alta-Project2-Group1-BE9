package shoppingcarts

type Core struct {
	ID            int
	TotalQuantity uint
	TotalPrice    uint
	Status        string
	User          User
}

type User struct {
	ID       int
	UserName string
}

type Business interface {
	GetHistoryOrder(id int, limit int, offset int) (data []Core, err error)
	CreateCart(idUser int, idBook int, data Core) (idCart int, rowSC int, errSC error)
	UpdatedStatusCart(id int, status string) (row int, err error)
	UpdatedCart(idCart int, idUser int, idBook int, quantity int, price int) (rowSC int, errSC error)
	DeleteCart(idCart int, idUser int, idBook int) (rowSC int, errSC error)
}

type Data interface {
	SelectAllOrder(id int, limit int, offset int) (data []Core, err error)
	InsertNewCart(data Core) (idShoppingCart int, row int, err error)
	UpdatedStatusCart(id int, status string) (row int, err error)
	IsCartNotExist(id int) (cond bool, data Core)
	UpdatedCart(idUser int, dataInput Core) (dataUpdated Core, row int, err error)
	SelectOrder(idUser int) (data Core, err error)
}
