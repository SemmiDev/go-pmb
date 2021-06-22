package repository

import (
	"go-clean/internal/app/entity"
	mongoConfig "go-clean/internal/config"
	exception2 "go-clean/internal/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "students"

func NewStudentRepository(database *mongo.Database) StudentRepository {
	return &db{
		Collection: database.Collection(collectionName),
	}
}

type db struct {
	Collection *mongo.Collection
}

func (r *db) Insert(student *entity.Student) (uid primitive.ObjectID) {
	ctx, cancel := mongoConfig.NewMongoContext()
	defer cancel()

	result, err := r.Collection.InsertOne(ctx, student)
	exception2.PanicIfNeeded(err)

	_id := result.InsertedID
	uid = _id.(primitive.ObjectID)

	return
}

func (r *db) GetById(id string) (student *entity.Student) {
	ctx, cancel := mongoConfig.NewMongoContext()
	defer cancel()

	err := r.Collection.FindOne(ctx, bson.M{
		"id": id,
	}).Decode(&student)

	exception2.PanicIfNeeded(err)
	return
}

func (r *db) FindAll() (students []entity.Student) {
	ctx, cancel := mongoConfig.NewMongoContext()
	defer cancel()

	cursor, err := r.Collection.Find(ctx, bson.M{})
	exception2.PanicIfNeeded(err)

	err = cursor.All(ctx, &students)
	exception2.PanicIfNeeded(err)

	return
}

func (r *db) GetByOID(oid primitive.ObjectID) (res *entity.Student, err error) {
	ctx, cancel := mongoConfig.NewMongoContext()
	defer cancel()

	err = r.Collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&res)
	exception2.PanicIfNeeded(err)
	return
}

func (r *db) GetByIdentifier(identifier string) (student *entity.Student, err error) {
	ctx, cancel := mongoConfig.NewMongoContext()
	defer cancel()

	err = r.Collection.FindOne(ctx, bson.M{
		"identifier": identifier,
	}).Decode(&student)

	exception2.PanicIfNeeded(err)
	return
}
