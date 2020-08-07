package wrap

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func idFilter(id primitive.ObjectID) bson.D {
	return bson.D{{"_id", id}}
}

// GetObj - функция для получения объекта из mongo.
func (w *Wrap) GetObj(id primitive.ObjectID, obj interface{}) error {
	c := w.Collection(obj)
	return c.FindOne(context.TODO(), idFilter(id)).Err()
}

// CreateObj - функция создания нового объекта в mongo.
func (w *Wrap) CreateObj(obj interface{}) (*mongo.InsertOneResult, error) {
	c := w.Collection(obj)
	return c.InsertOne(context.TODO(), obj)
}

// ReplaceObj - функция для замены одного объекта в mongo новым объектом.
func (w *Wrap) ReplaceObj(id primitive.ObjectID, obj interface{}) error {
	c := w.Collection(obj)
	_, err := c.ReplaceOne(context.TODO(), idFilter(id), obj)
	return err
}

// UpdateObj - Функция для обновления полей объекта в mongo.
func (w *Wrap) UpdateObj(id primitive.ObjectID, obj interface{}, update bson.D) error {
	c := w.Collection(obj)
	_, err := c.UpdateOne(context.TODO(), idFilter(id), update)
	return err
}

// DeleteObj - функция для удаления объекта из mongo.
func (w *Wrap) DeleteObj(id primitive.ObjectID, obj interface{}) error {
	c := w.Collection(obj)
	_, err := c.DeleteOne(context.TODO(), idFilter(id))
	return err
}
