package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id int) (*entity.Category, error) {
	if id <= 0 {
		return nil, errors.New("Category Not Found")
	}
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	}
	return category, nil
}
