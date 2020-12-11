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

// ChangeClient - смена клиента для работы с mongo
func (w *Wrap) ChangeClient(client *mongo.Client) {
	w.client = client
}

// ChangeDB - Смена рабочей базы данных
func (w *Wrap) ChangeDB(db string) {
	w.db = db
}

// Connect - Подключиться к СУБД
func (w *Wrap) Connect(ctx context.Context) error {
	return w.client.Connect(ctx)
}

// Disconnect - Разорвать соединение с СУБД
func (w *Wrap) Disconnect(ctx context.Context) error {
	return w.client.Disconnect(ctx)
}

// Fork - Метод для создания нового экземпляра Wrap
func (w *Wrap) Fork() Wrap {
	return Wrap{
		db:     w.db,
		client: w.client,
	}
}
