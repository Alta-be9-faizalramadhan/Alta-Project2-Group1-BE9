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
	CreateCart(idUser int, idBook int, data Core) (rowSC int, errSC error)
	UpdatedStatusCart(id int, status string) (row int, err error)
}

type Data interface {
	SelectAllOrder(id int, limit int, offset int) (data []Core, err error)
	InsertNewCart(data Core) (idShoppingCart int, row int, err error)
	UpdatedStatusCart(id int, status string) (row int, err error)
	IsCartNotExist(id int) (cond bool, data Core)
	UpdatedCart(idUser int, dataInput Core) (dataUpdated Core, row int, err error)
}
