package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"test_hh_1/config"
	"test_hh_1/init/domain"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GiveToken(user domain.RefreshToken) error {
	_, err := config.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	fmt.Println("успешно занесено")

	return nil
}

func (ur *UserRepo) RefreshToken(refToken, newRef string) error {
	var user domain.RefreshToken

	filter := bson.D{{"reftoken", refToken}}
	update := bson.D{
		{"$set", bson.D{
			{"reftoken", newRef},
		}},
	}

	err := config.Collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return err
	}

	_, err = config.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	fmt.Println("элемент изменен")
	return nil
}
