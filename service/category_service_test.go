package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {

	// program mock
	// categoryRepository.Mock.On("FindById", 1).Return(nil) // gpp di comment jika pada baris 20 Get(0)

	category, err := categoryService.Get(0) // kalau get <= 0, baris 18 gpp di komen karena mengembalikan nil. Kalau 1 ke atas maka diperlukan
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   2,
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", 2).Return(category)

	result, err := categoryService.Get(2)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
