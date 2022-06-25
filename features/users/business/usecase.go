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

func (uc *userUsecase) GetDataUser(idnya int) (resp users.Core, err error) {
	resp, err = uc.userData.SelectDataUser(idnya)
	return resp, err
}
