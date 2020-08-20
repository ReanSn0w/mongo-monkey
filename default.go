package mongo_monkey

import (
	"flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"github.com/ReanSn0w/mongo-monkey/wrap"
	"os"
)

var defaultWrap wrap.Wrap

func init() {
	uri := os.Getenv("MongoURI")
	if uri != "" {
		log.Printf("MongoURI: %v\n", uri)
		err := setupDefaultWrap(uri)

		if err == nil {
			return
		}

		log.Println(err)
	}

	if ptr := flag.String("mongoURI", "mongo://localhost:27017/", "mongoURI String"); ptr != nil {
		log.Printf("mongoURL: %v\n", *ptr)
		err := setupDefaultWrap(*ptr)

		if err == nil {
			return
		}

		log.Println(err)
	}

	err := setupDefaultWrap("mongo://localhost:27017/")

	if err != nil {
		log.Println("Не удалось инициализировать mongo-monkey в автоматическом режиме.")
	}
}

// SetupDefaultWrap - настроит стандартную обертку для работы необходимым Вам клиентов и базой данных
func SetupDefaultWrap(client *mongo.Client, db string) {
	defaultWrap = wrap.CreateWrap(client, db)
}

func setupDefaultWrap(uri string) error {
	client, err := uriToClient(uri)

	if err != nil {
		return err
	}

	SetupDefaultWrap(client, "data")

	return nil
}

func uriToClient(uri string) (*mongo.Client, error) {
	return mongo.NewClient(options.Client().ApplyURI(uri))
}
