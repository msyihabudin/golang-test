package repository

import "unittest/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
