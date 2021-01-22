package repository

import "belajar-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id int) *entity.Category
}
