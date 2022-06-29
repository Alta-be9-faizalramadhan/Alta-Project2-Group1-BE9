package users

import (
	"time"
)

type Core struct {
	ID        int
	UserName  string
	Email     string
	Password  string
	Alamat    string
	NoTelp    string
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	GetAllData(limit int, offset int) (data []Core, err error)
	InsertData(data Core) (row int, err error)
	UpdateData(id int, data Core) (row int, er error)
	GetDataUser(id int) (data Core, err error)
	DeleteData(id int) (row int, err error)
	Login(email string, password string) (token string, nama string, id int, err error)
}

type Data interface {
	SelectData(param string) (data []Core, err error)
	InsertData(data Core) (row int, err error)
	UpdateData(id int, data Core) (row int, er error)
	SelectDataUser(id int) (data Core, err error)
	DeleteDataUser(param int) (row int, err error)
	LoginUser(email string, password string) (token string, nama string, id int, err error)
}
