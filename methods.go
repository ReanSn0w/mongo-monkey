package mongo_monkey

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection - Данная функция повторяет метод Collection из структуры Wrap
func Collection(obj interface{}) *mongo.Collection {
	return defaultWrap.Collection(obj)
}

// GetSet - Данная функция повторяет метод GetSet из структуры Wrap
func GetSet(set interface{}, filter bson.D) error {
	return defaultWrap.GetSet(set, filter)
}

// Aggregate - Данная функция повторяет метод Aggregate из структуры Wrap
func AggreagateSet(set interface{}, pipeline mongo.Pipeline) error {
	return defaultWrap.AggreagateSet(set, pipeline)
}

// CreateSet - Данная функция повторяет метод CreateSet из структуры Wrap
func CreateSet(set interface{}) error {
	return defaultWrap.CreateSet(set)
}

// UpdateSet - Данная функция повторяет метод UpdateSet из структуры Wrap
func UpdateSet(set interface{}, filter bson.D, update bson.D) error {
	return defaultWrap.UpdateSet(set, filter, update)
}

// DeleteSet - Данная функция повторяет метод DeleteSet из структуры Wrap
func DeleteSet(set interface{}, filter bson.D) error {
	return defaultWrap.DeleteSet(set, filter)
}

// CountElements - Данная функция повторяет метод CountElements из структуры Wrap
func CountElements(set interface{}, filter bson.D) (int, error) {
	return defaultWrap.CountElements(set, filter)
}

// GetObj - Данная функция повторяет метод GetObj из структуры Wrap
func GetObj(id primitive.ObjectID, obj interface{}) error {
	return defaultWrap.GetObj(id, obj)
}

// CreateObj - Данная функция повторяет метод CreateObj из структуры Wrap
func CreateObj(obj interface{}) (*mongo.InsertOneResult, error) {
	return defaultWrap.CreateObj(obj)
}

// ReplaceObj - Данная функция повторяет метод ReplaceObj из структуры Wrap
func ReplaceObj(id primitive.ObjectID, obj interface{}) error {
	return defaultWrap.ReplaceObj(id, obj)
}

// UpdateObj - Данная функция повторяет метод UpdateObj из структуры Wrap
func UpdateObj(id primitive.ObjectID, obj interface{}, update bson.D) error {
	return defaultWrap.UpdateObj(id, obj, update)
}

// DeleteObj - Данная функция повторяет метод DeleteObj из структуры Wrap
func DeleteObj(id primitive.ObjectID, obj interface{}) error {
	return defaultWrap.DeleteObj(id, obj)
}

// Connect - Данная функция повторяет метод Connect из структуры Wrap
func Connect(ctx context.Context) error {
	return defaultWrap.Connect(ctx)
}

// Disconnect - Данная функция повторяет метод Disconnect из структуры Wrap
func Disconnect(ctx context.Context) error {
	return defaultWrap.Disconnect(ctx)
}