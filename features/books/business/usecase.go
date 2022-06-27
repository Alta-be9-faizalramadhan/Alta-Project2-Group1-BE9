package business

import (
	"altaproject/features/books"
	"errors"
)

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

func (uc *bookUsecase) CreateNewBook(input books.Core) (row int, err error) {
	if input.Title == "" || input.Author == "" || input.Publisher == "" || input.User.ID == 0 || input.Price == 0 || input.ISBN == "" {
		return -1, errors.New("all input data must be filled")
	}
	row, err = uc.bookData.InsertNewBook(input)
	return row, err
}

func (uc *bookUsecase) GetBookById(id int) (resp books.Core, err error) {
	resp, err = uc.bookData.SelectBookById(id)
	return resp, err
}

func (uc *bookUsecase) UpdatedBook(id int, data books.Core) (row int, err error) {
	row, err = uc.bookData.UpdatedBook(id, data)
	return row, err
}

func (uc *bookUsecase) SoftDeleteBook(id int) (row int, err error) {
	row, err = uc.bookData.SoftDeleteBook(id)
	return row, err
}
