package business

import (
	"altaproject/features/users"
	"errors"
)

type userUsecase struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUsecase{
		userData: usrData,
	}
}

func (uc *userUsecase) GetAllData(limit, offset int) (resp []users.Core, err error) {
	var param string
	resp, err = uc.userData.SelectData(param)
	return resp, err
}

func (uc *userUsecase) InsertData(input users.Core) (row int, err error) {
	if input.UserName == "" || input.Email == "" || input.Password == "" || input.Alamat == "" || input.NoTelp == "" {
		return -1, errors.New("all input data must be filled")
	}
	row, err = uc.userData.InsertData(input)
	return row, err
}

func (uc *userUsecase) UpdateData(id int, data users.Core) (row int, err error) {
	if data.UserName == "" || data.Alamat == "" || data.Email == "" || data.Password == "" || data.NoTelp == "" {
		return -1, errors.New("all input data must be filled")
	}
	row, err = uc.userData.UpdateData(id, data)
	return row, err
}

func (uc *userUsecase) GetDataUser(id int) (resp users.Core, err error) {
	resp, err = uc.userData.SelectDataUser(id)
	return resp, err
}

func (uc *userUsecase) DeleteData(id int) (row int, err error) {
	row, err = uc.userData.DeleteDataUser(id)
	return row, err
}

func (uc *userUsecase) Login(email string, password string) (token string, nama string, id int, err error) {
	token, nama, id, err = uc.userData.LoginUser(email, password)
	return token, nama, id, err
}
