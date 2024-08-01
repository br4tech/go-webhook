package filters

import "go.mongodb.org/mongo-driver/bson"

type EqualFilter struct {
	Field string
	Value string
}

func (f *EqualFilter) ToBson() bson.M {
	return bson.M{f.Field: f.Value}
}
