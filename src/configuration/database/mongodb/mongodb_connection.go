package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Usuario struct {
	Name string `bson:"nome"`
	Age  int8   `bson:"idade"`
}

var (
	MONGO_URL     = "MONGO_URL"
	DATABASE_NAME = "DATABASE_NAME"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongo_uri := os.Getenv(MONGO_URL)
	database_name := os.Getenv(DATABASE_NAME)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(database_name), nil
}

// func InitConnection() {

// 	// user := Usuario{
// 	// 	Name: "Luiz Eduardo",
// 	// 	Age:  31,
// 	// }

// 	// collection := client.Database("crudInit").Collection("teste")

// 	// result, err := collection.InsertOne(context.Background(), user)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// fmt.Println(result)

// 	// // consultando valores
// 	// filter := bson.D{{"nome", user.Name}} // preparando filtro para consulta
// 	// userResult := Usuario{}               // onde sera despejado o resultado com ponteiro

// 	// errFind := collection.FindOne(context.Background(), filter).Decode(&userResult) // consulta
// 	// if errFind != nil {
// 	// 	panic(errFind)
// 	// }

// 	// fmt.Println(userResult)
// }
