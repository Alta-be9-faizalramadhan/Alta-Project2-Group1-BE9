package data

import (
	"altaproject/features/users"
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
	user := fromCore(input)

	result := repo.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
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
