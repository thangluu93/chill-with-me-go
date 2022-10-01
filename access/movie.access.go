package access

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"main/models"
)

type Movie struct {
	DB             *mongo.Database
	CollectionName string
	collection     *mongo.Collection
}

func NewMovie(db *mongo.Database, collectionName string) *User {
	return &User{
		DB:             db,
		CollectionName: collectionName,
		collection:     db.Collection(collectionName),
	}
}

func (m *Movie) GetListMovies(limit int, offset int, genre string) (movie []*models.Movie, err error) {
	ctx := context.Background()
	filterOptions := options.Find()
	if limit != 0 {
		filterOptions.SetLimit(int64(limit))
	}
	if offset != 0 {
		filterOptions.SetSkip(int64(offset))
	}
	var filter = bson.M{}
	if genre != "" {
		filter = bson.M{"genre": genre}
	}

	cursor, errCursor := m.collection.Find(ctx, filter, filterOptions)
	if errCursor != nil {
		log.Fatal(errCursor)
		return nil, errCursor
	}
	var movies []models.Movie
	if cursor.All(ctx, &movies) != nil {
		log.Fatal(errCursor)
		return nil, errCursor
	}
	return movie, nil
}
