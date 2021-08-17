package db

import (
	"Dolphin/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/* Check existing User recibe un email como parametro y chequea si ya esta en la DB */
func CheckExistingUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("dolphin")
	col := db.Collection("users")
	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
