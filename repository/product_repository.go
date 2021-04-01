package repository

import "go-clean/entity"

type ProductRepository interface {
	Insert(product entity.Product)
	FindAll() (products []entity.Product)
	DeleteAll()
}