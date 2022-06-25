package request

import (
	"altaproject/features/users"
	"time"
)

type User struct {
	UserName  string    `json:"username" form:"username"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Alamat    string    `json:"alamat" form:"alamat"`
	NoTelp    string    `json:"notelp" form:"notelp"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func ToCore(req User) users.Core {
	return users.Core{
		UserName: req.UserName,
		Email:    req.Email,
		Password: req.Password,
		Alamat:   req.Alamat,
		NoTelp:   req.NoTelp,
	}
}
