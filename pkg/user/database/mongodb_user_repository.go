package database

import (
	"context"
	"errors"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBUserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *MongoDBUserRepository {
	return &MongoDBUserRepository{
		collection: collection,
	}
}

func (ur *MongoDBUserRepository) Save(user *domain.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	_, err := ur.collection.InsertOne(context.Background(), user)
	return err
}

func (ur *MongoDBUserRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User

	search, errID := primitive.ObjectIDFromHex(id)

	if errID != nil {
		return nil, domain.InvalidID
	}

	err := ur.collection.FindOne(context.Background(), bson.M{"_id": search}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (ur *MongoDBUserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := ur.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (ur *MongoDBUserRepository) FindByDocumentNo(documentNo string) (*domain.User, error) {
	var user domain.User

	err := ur.collection.FindOne(context.Background(), bson.M{"documentno": documentNo}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (ur *MongoDBUserRepository) Update(user *domain.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	opts := options.Update().SetUpsert(true)
	_, err := ur.collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, bson.M{"$set": user}, opts)
	return err
}

func (ur *MongoDBUserRepository) Delete(id primitive.ObjectID) error {
	_, err := ur.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
