package repository

import (
	"go-clean/config"
	"go-clean/entity"
	"go-clean/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProductRepository(database *mongo.Database) ProductRepository {
	return &productRepositoryImpl{
		Collection: database.Collection("products"),
	}
}

type productRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repo *productRepositoryImpl) Insert(product entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repo.Collection.InsertOne(ctx, bson.M{
		"_id":      	product.Id,
		"code":      	product.Code,
		"name":      	product.Name,
		"price":      	product.Price,
		"available":    product.Avaliable,
		"stock":      	product.Stock,
	})

	exception.PanicIfNeeded(err)
}

func (repo *productRepositoryImpl) FindAll() (products []entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repo.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		products = append(products, entity.Product{
			Id: document["_id"].(string),
			Code: document["code"].(string),
			Name: document["name"].(string),
			Price: document["price"].(float64),
			Avaliable: document["available"].(bool),
			Stock: document["stock"].(int),
		})
	}
	return products
}

func (repo *productRepositoryImpl) DeleteAll() {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_,err := repo.Collection.DeleteMany(ctx, bson.M{})
	exception.PanicIfNeeded(err)
}


