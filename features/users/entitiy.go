package users

import "time"

type Core struct {
	ID        int
	UserName  string
	Email     string
	Password  string
	Alamat    string
	NoTelp    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	GetAllData(limit int, offset int) (data []Core, err error)
	InsertData(data Core) (row int, err error)
	GetDataUser(id int) (data Core, err error)
}

type Data interface {
	SelectData(param string) (data []Core, err error)
	InsertData(data Core) (row int, err error)
	SelectDataUser(id int) (data Core, err error)
}
