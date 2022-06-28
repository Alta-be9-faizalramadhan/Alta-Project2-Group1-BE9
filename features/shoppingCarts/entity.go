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
}

type Data interface {
	SelectAllOrder(id int, limit int, offset int) (data []Core, err error)
}
