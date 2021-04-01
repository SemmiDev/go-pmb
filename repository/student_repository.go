package repository

import "go-clean/entity"

type StudentRepository interface {
	Insert(student entity.Student)
	Delete(id string) string
	FindAll() (students []entity.Student)
	DeleteAll()
	GetById(id string) entity.Student
	UpdateById(id string, student entity.Student) bool
}