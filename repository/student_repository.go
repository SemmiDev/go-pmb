package repository

import "go-clean/entity"

type StudentRepository interface {
	Insert(product entity.Student)
	FindAll() (products []entity.Student)
	DeleteAll()
}