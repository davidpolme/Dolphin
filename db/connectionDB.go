package db 

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


/* MongoCN is the connection object to the database */
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://davidpolme:<password>@dolphin.ajwlm.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")


/* ConnectDB is the function that allows the connection to the db */
func ConnectDB() *mongo.Client {
    client, err := mongo.Connect(context.TODO(),clientOptions)

    if err != nil {
        log.Fatal(err)
        return client
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
        return client
    }
    log.Println("Conexion exitosa con la BD")
    return client
}   

/* CheckConnection is the ping to the DB */
func CheckConnection() int {

    err := MongoCN.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
        return 0
    }
    return 1
}
