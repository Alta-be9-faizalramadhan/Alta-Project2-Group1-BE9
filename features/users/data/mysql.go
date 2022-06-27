package data

import (
	"altaproject/encription"
	"altaproject/features/users"
	"altaproject/middlewares"
	"fmt"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) SelectData(data string) (response []users.Core, err error) {
	var dataUsers []User
	result := repo.db.Find(&dataUsers)
	if result.Error != nil {
		return []users.Core{}, result.Error
	}
	return toCoreList(dataUsers), nil
}

func (repo *mysqlUserRepository) InsertData(input users.Core) (row int, err error) {
	// user := fromCore(input)
	passwordHash := encription.GetMD5Hash(input.Password)
	user := User{
		UserName: input.UserName,
		Email:    input.Email,
		Password: string(passwordHash),
		Alamat:   input.Alamat,
		NoTelp:   input.NoTelp,
	}
	result := repo.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}

	return int(result.RowsAffected), nil
}

func (repo *mysqlUserRepository) UpdateData(id int, data users.Core) (int, error) {
	var update = fromCore(data)
	result := repo.db.Model(&User{}).Where("id = ?", id).Updates(&update)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("failed to update data")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlUserRepository) SelectDataUser(id int) (response users.Core, err error) {
	var dataUsers User

	result := repo.db.First(&dataUsers, id)
	if result.Error != nil {
		return users.Core{}, result.Error
	}
	return dataUsers.toCore(), nil
}

func (repo *mysqlUserRepository) DeleteDataUser(data int) (row int, err error) {
	var dataUsers User
	result := repo.db.Delete(&dataUsers, data)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to delete user")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlUserRepository) LoginUser(data users.Core) (token string, username string, err error) {
	userData := User{}

	result := repo.db.Where("email = ?", data.Email).First(&userData)
	result = repo.db.Select("password").First(&userData, "email = ?", data.Email)
	strPassword := encription.GetMD5Hash(data.Password)
	if strPassword != userData.toCore().Password {
		return "", "", fmt.Errorf("error")
	} else {
		if result.Error != nil {
			return "", "", result.Error
		}
		if result.RowsAffected != 1 {
			return "", "", fmt.Errorf("error")
		}
		token, errToken := middlewares.CreateToken(int(userData.ID))
		if errToken != nil {
			return "", "", errToken
		}
		username = userData.UserName
		return token, username, nil
	}
}
