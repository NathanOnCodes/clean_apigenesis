package repository

import (
	"clean_architecture/api_genesis/internal/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InterfaceConversionRepository interface {
	Create(conversion *entity.Conversion) error
	FindAll() ([]entity.Conversion, error)
}

type MongoDBConversionRepository struct {
	client *mongo.Client
}


func NewMongoDBConversionRepository(mongodb *mongo.Client) *MongoDBConversionRepository {
	return &MongoDBConversionRepository{client: mongodb}
}

func (repo *MongoDBConversionRepository) Create(conversion *entity.Conversion) error {
	collection := repo.client.Database("mongodb_api_genesis").Collection("conversions")
	_, err := collection.InsertOne(context.Background(), bson.M{
		"name":   conversion.NameCoin,
		"symbol": conversion.SymbolCoin,
		"value":  conversion.Value,
	})
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoDBConversionRepository) FindAll() ([]entity.Conversion, error) {
	collection := repo.client.Database("mongodb_api_genesis").Collection("conversions")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var conversions []entity.Conversion
	for cursor.Next(context.Background()) {
		var conversion entity.Conversion
		err := cursor.Decode(&conversion)
		if err != nil {
			return nil, err
		}
		conversions = append(conversions, conversion)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return conversions, nil
}
