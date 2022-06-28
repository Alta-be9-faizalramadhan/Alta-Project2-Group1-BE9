package business

import (
	"altaproject/features/users"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockUserData struct{}

func (mock mockUserData) SelectData(param string) (data []users.Core, err error) {
	return []users.Core{
		{ID: 1, UserName: "alta", Email: "alta@gmail.com", Password: "qwerty", Alamat: "Jakarta", NoTelp: "081234256"},
	}, nil
}

func (mock mockUserData) SelectDataUser(id int) (data users.Core, err error) {
	return users.Core{
		ID: 1, UserName: "alta", Email: "alta@gmail.com", Password: "qwerty", Alamat: "Jakarta", NoTelp: "081234256",
	}, nil
}

func (mock mockUserData) InsertData(data users.Core) (row int, err error) {
	return 1, nil
}

func (mock mockUserData) DeleteDataUser(id int) (row int, err error) {
	return 1, nil
}

func (mock mockUserData) UpdateData(id int, data users.Core) (row int, err error) {
	return 1, nil
}

func (mock mockUserData) LoginUser(email string, password string) (token string, nama string, id int, err error) {
	return "abncdgru465599f8c", "alta", 1, nil
}

type mockUserDataFailed struct{}

func (mock mockUserDataFailed) SelectData(param string) (data []users.Core, err error) {
	return nil, fmt.Errorf("Failed to select data")
}

func (mock mockUserDataFailed) SelectDataUser(id int) (data users.Core, err error) {
	return users.Core{}, fmt.Errorf("failed to select data user")
}

func (mock mockUserDataFailed) InsertData(data users.Core) (row int, err error) {
	return 0, fmt.Errorf("failed to insert data")
}

func (mock mockUserDataFailed) DeleteDataUser(id int) (row int, err error) {
	return 0, fmt.Errorf("failed to delete data user")
}

func (mock mockUserDataFailed) UpdateData(id int, data users.Core) (row int, err error) {
	return 0, fmt.Errorf("failed to edit data user")
}

func (mock mockUserDataFailed) LoginUser(email string, password string) (token string, nama string, id int, err error) {
	return "", "", 0, fmt.Errorf("failed to edit data user")
}

func TestGetAllData(t *testing.T) {
	t.Run("Test Get All Data Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.GetAllData(0, 0)
		assert.Nil(t, err)
		assert.Equal(t, "alta", result[0].UserName)
	})

	t.Run("Test Get All Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.GetAllData(0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestGetDataUser(t *testing.T) {
	t.Run("Test Get User Data Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.GetDataUser(0)
		assert.Nil(t, err)
		assert.Equal(t, "alta", result.UserName)
	})

	t.Run("Test Get User Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.GetDataUser(0)
		assert.NotNil(t, err)
		assert.Equal(t, users.Core{}, result)
	})
}

func TestInsertData(t *testing.T) {
	t.Run("Test Insert Data Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		newUser := users.Core{
			UserName: "alta",
			Email:    "alta@mail.id",
			Password: "qwerty",
			Alamat:   "Jakarta",
			NoTelp:   "081234256",
		}
		result, err := userBusiness.InsertData(newUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Test Insert Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			UserName: "alta",
			Email:    "alta@mail.id",
			Password: "qwerty",
			Alamat:   "Jakarta",
			NoTelp:   "081234256",
		}
		result, err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})

	t.Run("Test Insert Data Failed When Email Empty", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		newUser := users.Core{
			UserName: "alta",
			Password: "qwerty",
			Alamat:   "Jakarta",
			NoTelp:   "081234256",
		}
		result, err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})

	t.Run("Test Insert Data Failed When Password Empty", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			UserName: "alta",
			Email:    "alta@mail.id",
			Alamat:   "Jakarta",
			NoTelp:   "081234256",
		}
		result, err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})

	t.Run("Test Insert Data Failed When Name Empty", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Email:    "alta@mail.id",
			Password: "qwerty",
			Alamat:   "Jakarta",
			NoTelp:   "081234256",
		}
		result, err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
}

func TestDeleteData(t *testing.T) {
	t.Run("Test Delete Data Succes", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.DeleteData(0)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Test Delete Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.DeleteData(0)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestUpdateData(t *testing.T) {
	t.Run("Test Update Data Succes", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		editUser := users.Core{
			UserName: "alta",
			Email:    "alta@mail.id",
			Password: "qwerty",
			Alamat:   "Jakarta",
			NoTelp:   "081234256",
		}
		result, err := userBusiness.UpdateData(0, editUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Test Update Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		editUser := users.Core{
			UserName: "alta",
			Email:    "alta@mail.id",
			Password: "qwerty",
		}
		result, err := userBusiness.UpdateData(0, editUser)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Test Login User Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		var loginEmail = "alta@mail.id"
		var loginPass = "alta123"

		result1, result2, result3, err := userBusiness.Login(loginEmail, loginPass)
		assert.Nil(t, err)
		assert.Equal(t, "abncdgru465599f8c", "abncdgru465599f8c", result1, result2, result3)
	})
	t.Run("Test Login User Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		var loginEmail = "alta@mail.id"
		var loginPass = "alta123"

		result1, result2, result3, err := userBusiness.Login(loginEmail, loginPass)
		assert.NotNil(t, err)
		assert.Equal(t, "abncdgru465599f8c", "abncdgru465599f8c", result1, result2, result3)
	})
}
