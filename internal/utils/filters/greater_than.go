package filters

import "go.mongodb.org/mongo-driver/bson"

type GreaterThan struct {
	Field string
	Value interface{}
}

func (f *GreaterThan) ToBson() bson.M {
	return bson.M{f.Field: bson.M{"$gt": f.Value}}
}
