package wrap

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection - Метод для получения коллекции на основе типа документа ( Используется во всех методах для работы с докуметами и коллекциями в mongoWrap )
func (w *Wrap) Collection(obj interface{}) *mongo.Collection {
	return w.client.Database(w.db).Collection(reflectName(obj))
}

// GetSet - Получение нескольких объектов из mongo.
func (w *Wrap) GetSet(set interface{}, filter bson.D) error {
	c := w.Collection(set)

	cur, err := c.Find(context.TODO(), filter)

	if err != nil {
		return err
	}

	return cur.All(context.TODO(), set)
}

// Aggregate - Агрегация объектов mongo.
func (w *Wrap) AggreagateSet(set interface{}, pipeline mongo.Pipeline) error {
	c := w.Collection(set)
	cur, err := c.Aggregate(context.TODO(), pipeline)

	if err != nil {
		return err
	}

	return cur.All(context.TODO(), set)
}

// CreateSet - создаст несколько в коллекции mongo.
func (w *Wrap) CreateSet(set interface{}) error {
	c := w.Collection(set)
	objs := set.([]interface{})
	_, err := c.InsertMany(context.TODO(), objs)
	return err
}

// UpdateSet - обновит несколько документов в mongo.
func (w *Wrap) UpdateSet(set interface{}, filter bson.D, update bson.D) error {
	c := w.Collection(set)
	_, err := c.UpdateMany(context.TODO(), filter, update)
	return err
}

// DeleteSet - удалит несколько документов из mongo.
func (w *Wrap) DeleteSet(set interface{}, filter bson.D) error {
	c := w.Collection(set)
	_, err := c.DeleteMany(context.TODO(), filter)
	return err
}

// CountElements - рассчет колличества элементов в коллекции, удовлетворяющих фильтру
func (w *Wrap) CountElements(set interface{}, filter bson.D) (int, error) {
	c := w.Collection(set)
	count, err := c.CountDocuments(context.TODO(), filter)
	return int(count), err
}