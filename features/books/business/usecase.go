package business

import "altaproject/features/books"

type bookUsecase struct {
	bookData books.Data
}

func NewBookBusiness(bkData books.Data) books.Business {
	return &bookUsecase{
		bookData: bkData,
	}
}

func (uc *bookUsecase) GetAllBook(limit, offset uint) (resp []books.Core, err error) {
	resp, err = uc.bookData.SelectAllBook(limit, offset)
	return resp, err
}
