package wrap

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

// Wrap - обертка над драйвером монго для упрощения работы с БД и коллекциями
type Wrap struct {
	client *mongo.Client
	db     string
}

// CreateWrap - создает экземпляр обертки над драйвером mongoDB
func CreateWrap(client *mongo.Client, db string) Wrap {
	return Wrap{
		client: client,
		db:     db,
	}
}

func (w *Wrap) ChangeClient(client *mongo.Client) {
	w.client = client
}

func (w *Wrap) ChangeDB(db string) {
	w.db = db
}

func (w *Wrap) Connect(ctx context.Context) error {
	return w.client.Connect(ctx)
}

func (w *Wrap) Disconnect(ctx context.Context) error {
	return w.client.Disconnect(ctx)
}
