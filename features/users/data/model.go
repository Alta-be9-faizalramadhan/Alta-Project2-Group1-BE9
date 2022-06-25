package data

import (
	"altaproject/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Alamat   string `json:"alamat"`
	NoTelp   string `json:"notelp"`
}

func (data *User) toCore() users.Core {
	return users.Core{
		ID:        int(data.ID),
		UserName:  data.UserName,
		Email:     data.Email,
		Password:  data.Password,
		Alamat:    data.Alamat,
		NoTelp:    data.NoTelp,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func toCoreList(data []User) []users.Core {
	result := []users.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func toCore(data User) users.Core {
	result := users.Core{
		ID:       int(data.ID),
		UserName: data.UserName,
		Email:    data.Email,
		Alamat:   data.Alamat,
		NoTelp:   data.NoTelp,
		// Password:  data.Password,
		CreatedAt: data.CreatedAt,
	}
	return result
}

func fromCore(core users.Core) User {
	return User{
		UserName: core.UserName,
		Email:    core.Email,
		Password: core.Password,
		Alamat:   core.Alamat,
		NoTelp:   core.NoTelp,
	}
}
