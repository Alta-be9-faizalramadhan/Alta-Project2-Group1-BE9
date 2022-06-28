package business

import (
	"altaproject/features/books"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockBookData struct{}

func (mock mockBookData) SelectAllBook(limit, offset uint) (data []books.Core, err error) {
	return []books.Core{
		{ID: 1, Title: "Kamus", Author: "Qodir", Publisher: "Gramedia", ISBN: "43215ISBN", Category: "TextBook", Price: 20000, Stock: 5, BookPage: "100", Sold: 0, Description: "Buku baru keluar", User: books.User{ID: 1, UserName: "Aldi"}},
	}, nil
}

func (mock mockBookData) InsertNewBook(input books.Core) (row int, err error) {
	return 1, nil
}

func (mock mockBookData) SelectBookById(id int) (data books.Core, err error) {
	return books.Core{
		ID: 1, Title: "Kamus", Author: "Qodir", Publisher: "Gramedia", ISBN: "43215ISBN", Category: "TextBook", Price: 20000, Stock: 5, BookPage: "100", Sold: 0, Description: "Buku baru keluar", User: books.User{ID: 1, UserName: "Aldi"},
	}, nil
}

func (mock mockBookData) UpdatedBook(id int, data books.Core) (row int, err error) {
	return 1, nil
}

func (mock mockBookData) SoftDeleteBook(id int) (row int, err error) {
	return 1, nil
}

func (mock mockBookData) SelectBookByCategory(category string) (resp []books.Core, err error) {
	return []books.Core{
		{ID: 1, Title: "Kamus", Author: "Qodir", Publisher: "Gramedia", ISBN: "43215ISBN", Category: "TextBook", Price: 20000, Stock: 5, BookPage: "100", Sold: 0, Description: "Buku baru keluar", User: books.User{ID: 1, UserName: "Aldi"}},
	}, nil
}

// func (mock mockBookData) SelectBookByUserId(id int) (resp []books.Core, err error) {
// 	return []books.Core{
// 		{ID: 1, Title: "Kamus", Author: "Qodir", Publisher: "Gramedia", ISBN: "43215ISBN", Category: "TextBook", Price: 20000, Stock: 5, BookPage: "100", Sold: 0, Description: "Buku baru keluar", User: books.User{ID: 1, UserName: "Aldi"}},
// 	}, nil
// }

type mockBookDataFailed struct{}

func (mock mockBookDataFailed) SelectAllBook(limit, offset uint) (data []books.Core, err error) {
	return nil, fmt.Errorf("Failed to select data")
}

func (mock mockBookDataFailed) InsertNewBook(input books.Core) (row int, err error) {
	return 0, fmt.Errorf("Failed to select data")
}

func (mock mockBookDataFailed) SelectBookById(id int) (data books.Core, err error) {
	return books.Core{}, fmt.Errorf("Failed to select data")
}

func (mock mockBookDataFailed) UpdatedBook(id int, data books.Core) (row int, err error) {
	return 0, fmt.Errorf("Failed to select data")
}

func (mock mockBookDataFailed) SoftDeleteBook(id int) (row int, err error) {
	return 0, fmt.Errorf("Failed to select data")
}

func (mock mockBookDataFailed) SelectBookByCategory(category string) (resp []books.Core, err error) {
	return nil, fmt.Errorf("Failed to select data")
}

// func (mock mockBookDataFailed) SelectBookByUserId(id int) (resp []books.Core, err error) {
// 	return nil, fmt.Errorf("Failed to select data")
// }

func TestGetAllBook(t *testing.T) {
	t.Run("Test Get All Book Success", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookData{})
		result, err := userBusiness.GetAllBook(0, 0)
		assert.Nil(t, err)
		assert.Equal(t, "Kamus", result[0].Title)
	})

	t.Run("Test Get All Book Failed", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookDataFailed{})
		result, err := userBusiness.GetAllBook(0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestCreateNewBook(t *testing.T) {
	t.Run("Test Create New Book Success", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookData{})
		newBook := books.Core{
			Title: "Kamus", Author: "Qodir", Publisher: "Gramedia", ISBN: "43215ISBN", Category: "TextBook", Price: 20000, Stock: 5, BookPage: "100", Sold: 0, Description: "Buku baru keluar", User: books.User{ID: 1, UserName: "Aldi"},
		}
		result, err := userBusiness.CreateNewBook(newBook)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Create New Book Failed", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookDataFailed{})
		newBook := books.Core{
			Title: "Kamus", Author: "Qodir", Publisher: "Gramedia", ISBN: "43215ISBN", Category: "TextBook", Price: 20000, Stock: 5, BookPage: "100", Sold: 0, Description: "Buku baru keluar", User: books.User{ID: 1, UserName: "Aldi"},
		}
		result, err := userBusiness.CreateNewBook(newBook)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
	t.Run("Test Create New Book Failed When Title Empty", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookDataFailed{})
		newBook := books.Core{
			Author: "Qodir", Publisher: "Gramedia", ISBN: "43215ISBN", Category: "TextBook", Price: 20000, Stock: 5, BookPage: "100", Sold: 0, Description: "Buku baru keluar", User: books.User{ID: 1, UserName: "Aldi"},
		}
		result, err := userBusiness.CreateNewBook(newBook)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
	t.Run("Test Create New Book Failed When Author Empty", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookDataFailed{})
		newBook := books.Core{
			Title: "Kamus", Publisher: "Gramedia", ISBN: "43215ISBN", Category: "TextBook", Price: 20000, Stock: 5, BookPage: "100", Sold: 0, Description: "Buku baru keluar", User: books.User{ID: 1, UserName: "Aldi"},
		}
		result, err := userBusiness.CreateNewBook(newBook)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})

}

func TestGetBookById(t *testing.T) {
	t.Run("Test Get Book By Id Success", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookData{})
		result, err := userBusiness.GetBookById(1)
		assert.Nil(t, err)
		assert.Equal(t, "Kamus", result.Title)
	})
	t.Run("Test Get Book By Id Failed", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookDataFailed{})
		result, err := userBusiness.GetBookById(0)
		assert.NotNil(t, err)
		assert.Equal(t, "", result.Title)
	})
}

func TestUpdatedBook(t *testing.T) {
	t.Run("Test Updated Book Success", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookData{})
		updBook := books.Core{
			Title: "Kamus", Author: "Qodir", Publisher: "Gramedia", ISBN: "43215ISBN", Category: "TextBook", Price: 20000, Stock: 5, BookPage: "100", Sold: 0, Description: "Buku baru keluar", User: books.User{ID: 1, UserName: "Aldi"},
		}
		result, err := userBusiness.UpdatedBook(0, updBook)
		assert.Nil(t, err)
		assert.NotNil(t, 1, result)
	})
	t.Run("Test Get Updated Book Failed", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookDataFailed{})
		updBook := books.Core{
			Title: "Kamus", Author: "Qodir", Publisher: "Gramedia", ISBN: "43215ISBN", Category: "TextBook", Price: 20000, Stock: 5, BookPage: "100", Sold: 0, Description: "Buku baru keluar", User: books.User{ID: 1, UserName: "Aldi"},
		}
		result, err := userBusiness.UpdatedBook(0, updBook)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestSoftDeleteBook(t *testing.T) {
	t.Run("Test Soft Delete Book Success", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookData{})
		var id = 0
		result, err := userBusiness.SoftDeleteBook(id)
		assert.Nil(t, err)
		assert.NotNil(t, 1, result)
	})
	t.Run("Test Soft Delete Book Failed", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookDataFailed{})
		var id = 0
		result, err := userBusiness.SoftDeleteBook(id)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestSelectBookByCategory(t *testing.T) {
	t.Run("Test Select Book By Category Success", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookData{})
		var ctg = books.Core{Category: "TextBook"}
		result, err := userBusiness.SelectBookByCategory(ctg.Category)
		assert.Nil(t, err)
		assert.NotNil(t, "TextBook", result[0].Category)
	})
	t.Run("Test Select Book By Category Failed", func(t *testing.T) {
		userBusiness := NewBookBusiness(mockBookDataFailed{})
		var ctg = books.Core{Category: "TextBook"}
		result, err := userBusiness.SelectBookByCategory(ctg.Category)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
