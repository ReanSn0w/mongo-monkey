package mongo_monkey

import (
	"errors"
	"flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"github.com/ReanSn0w/mongo-monkey/wrap"
	"os"
)

var defaultWrap wrap.Wrap

func init() {
	if err := setupMongoURI(); err == nil {
		setupDatabase()
		return
	}else{
		log.Println(err)
	}
}

func setupMongoURI() error {
	uri := os.Getenv("MongoURI")

	if uri != "" {
		log.Printf("MongoURI: %v\n", uri)
		err := setupDefaultWrap(uri)

		if err == nil {
			return nil
		}
	}

	if ptr := flag.String("mongoURI", "mongo://localhost:27017/", "mongoURI String"); ptr != nil {
		log.Printf("mongoURL: %v\n", *ptr)
		err := setupDefaultWrap(*ptr)

		if err == nil {
			return nil
		}
	}

	return errors.New("Не удалось инициализировать mongo-monkey в автоматическом режиме.")
}

func setupDatabase() {
	database := os.Getenv("Database")
	if database != "" {
		defaultWrap.ChangeDB(database)
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

// ChangeDB - повторяет метод ChangeDB из пакета Wrap
func ChangeDB(db string) {
	defaultWrap.ChangeDB(db)
}