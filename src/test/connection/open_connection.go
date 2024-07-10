package connection

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ory/dockertest"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenConnection() (database *mongo.Database, close func()) {

	// criando repositorio ------------------------------------
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
	})

	if err != nil {
		log.Fatalf("Erro ao abrir conexão: %s", err)
	}

	// criando conexão ------------------------------------------
	client, err := mongo.NewClient(options.Client().ApplyURI(
		fmt.Sprintf("mongodb://127.0.0.1:%s", resource.GetPort("27017/tcp")))) //gerando porta aleatoria usando getPort

	if err != nil {
		log.Println("Erro ao criar conexão")
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Println("Erro ao tentar conectar!")
	}

	database = client.Database(os.Getenv("DATABASE_NAME"))

	close = func() {
		err := resource.Close()
		if err != nil {
			log.Println("Erro ao fechar conexão")
			return
		}

	}
	return
}
