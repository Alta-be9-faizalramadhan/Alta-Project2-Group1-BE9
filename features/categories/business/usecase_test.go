package business

import (
	category "altaproject/features/categories"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCategoryData struct{}

func (repo mockCategoryData) SelectAllCategory() ([]category.Core, error) {
	return []category.Core{{Name: "Novel"}, {Name: "Comic"}}, nil
}

type mockCategoryDataFailed struct{}

func (repo mockCategoryDataFailed) SelectAllCategory() ([]category.Core, error) {
	return nil, fmt.Errorf("Failed to select data")
}

func TestGetAllBook(t *testing.T) {
	t.Run("Test Get All Category Success", func(t *testing.T) {
		categoryBusiness := NewCategoryBusiness(mockCategoryData{})
		result, err := categoryBusiness.GetAllCategory()
		assert.Nil(t, err)
		assert.Equal(t, "Novel", result[0].Name)
	})

	t.Run("Test Get All Category Failed", func(t *testing.T) {
		categoryBusiness := NewCategoryBusiness(mockCategoryDataFailed{})
		result, err := categoryBusiness.GetAllCategory()
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
