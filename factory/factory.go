package factory

import (
	_bookBusiness "altaproject/features/books/business"
	_bookData "altaproject/features/books/data"
	_bookPresentation "altaproject/features/books/presentation"
	_shoppingCartBusiness "altaproject/features/shoppingCarts/business"
	_shoppingCartData "altaproject/features/shoppingCarts/data"
	_shoppingCartPresentation "altaproject/features/shoppingCarts/presentation"
	_userBusiness "altaproject/features/users/business"
	_userData "altaproject/features/users/data"
	_userPresentation "altaproject/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter         *_userPresentation.UserHandler
	BookPresenter         *_bookPresentation.BookHandler
	ShoppingCartPresenter *_shoppingCartPresentation.ShoppingCartHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	bookData := _bookData.NewBookRepository(dbConn)
	bookBusiness := _bookBusiness.NewBookBusiness(bookData)
	bookPresentation := _bookPresentation.NewBookHandler(bookBusiness)

	shoppingCartData := _shoppingCartData.NewShoppingCartRepository(dbConn)
	shoppingCartBusiness := _shoppingCartBusiness.NewShoppingCartBusiness(shoppingCartData)
	shoppingCartPresentation := _shoppingCartPresentation.NewShoppingCartHandler(shoppingCartBusiness)

	return Presenter{
		UserPresenter:         userPresentation,
		BookPresenter:         bookPresentation,
		ShoppingCartPresenter: shoppingCartPresentation,
	}
}
