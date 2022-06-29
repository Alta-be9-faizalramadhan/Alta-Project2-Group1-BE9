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
	CreateNewCart(data Core) (row int, err error)
	UpdatedStatusCart(id int, status string) (row int, err error)
}

type Data interface {
	SelectAllOrder(id int, limit int, offset int) (data []Core, err error)
	InsertNewCart(data Core) (row int, err error)
	UpdatedStatusCart(id int, status string) (row int, err error)
}
