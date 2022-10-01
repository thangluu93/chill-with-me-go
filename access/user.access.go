package access

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main/models"
)

type User struct {
	DB             *mongo.Database
	CollectionName string
	collection     *mongo.Collection
}

func NewUser(db *mongo.Database, collectionName string) *User {
	return &User{
		DB:             db,
		CollectionName: collectionName,
		collection:     db.Collection(collectionName),
	}
}

func (a *User) GetByToken(token string) (user *models.User, err error) {
	user = new(models.User)
	err = a.collection.FindOne(context.Background(), bson.M{"token": token}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *User) GetUserById(id string) (user *models.User, err error) {
	user = new(models.User)
	err = a.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
