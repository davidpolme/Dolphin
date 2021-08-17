package db

import (
	"Dolphin/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertRegister is the last stop with the DB to insert the user data */
func InsertRegister(u models.User) (string, bool, error) {

	//ctx context
	ctx, cancer := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancer()

	db := MongoCN.Database("dolphin")
	col := db.Collection("users")

	/* Funcion que encripta la contraseña */
	u.Password, _ = EncryptPasword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
