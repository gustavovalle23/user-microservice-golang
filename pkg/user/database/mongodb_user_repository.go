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

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (ur *UserRepository) Save(user *domain.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	_, err := ur.collection.InsertOne(context.Background(), user)
	return err
}

func (ur *UserRepository) FindById(id primitive.ObjectID) (*domain.User, error) {
	var user domain.User

	err := ur.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) FindByEmail(email string) (*domain.User, error) {
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

func (ur *UserRepository) FindByDocumentNo(documentNo string) (*domain.User, error) {
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

func (ur *UserRepository) Update(user *domain.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	opts := options.Update().SetUpsert(true)
	_, err := ur.collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, bson.M{"$set": user}, opts)
	return err
}

func (ur *UserRepository) Delete(id primitive.ObjectID) error {
	_, err := ur.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
