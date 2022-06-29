package response

import (
	"altaproject/features/users"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Alamat    string    `json:"alamat"`
	NoTelp    string    `json:"notelp"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func FromCore(data users.Core) User {
	return User{
		ID:        data.ID,
		UserName:  data.UserName,
		Email:     data.Email,
		Alamat:    data.Alamat,
		NoTelp:    data.NoTelp,
		ImageURL:  data.ImageURL,
		CreatedAt: data.CreatedAt,
	}
}

func FromCoreList(data []users.Core) []User {
	result := []User{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
