package wrap

import (
	"crypto/md5"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"strings"
)

// PredictableObjectID - предсказуемая генерация objectID на основе строки
func (w *Wrap) PredictableObjectID(str string) (primitive.ObjectID, error) {
	str = fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return primitive.ObjectIDFromHex(str[:24])
}

func reflectName(obj interface{}) string {
	name := reflect.TypeOf(obj).String()
	if name[0] == []byte("*")[0] {
		name = name[1:]
	}
	if name[:2] == "[]" {
		name = name[2:]
	}
	split := strings.Split(name, ".")
	return split[len(split)-1]
}

func reverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
