package factory

import (
	_userBusiness "altaproject/features/users/business"
	_userData "altaproject/features/users/data"
	_userPresentation "altaproject/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter *_userPresentation.UserHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	return Presenter{
		UserPresenter: userPresentation,
	}
}
