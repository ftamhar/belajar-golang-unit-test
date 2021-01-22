package repository

import (
	"belajar-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id int) *entity.Category {

	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}

	category := arguments.Get(0).(entity.Category)
	category.Id = id // set id agar sesuai dengan yang di panggil method FindById, bukan result
	return &category
}
