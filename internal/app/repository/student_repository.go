package repository

import (
	"go-clean/internal/app/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentRepository interface {
	Insert(student *entity.Student) primitive.ObjectID
	GetById(id string) *entity.Student
	FindAll() (students []entity.Student)
	GetByOID(oid primitive.ObjectID) (*entity.Student, error)
	GetByIdentifier(identifier string) (*entity.Student, error)
}