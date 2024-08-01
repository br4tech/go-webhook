package filters

type AndFilter struct {
	Filters []interface{}
}

// func (f *AndFilter) ToBson() bson.M {
// 	filters := make([]bson.M, len(f.Filters))
// 	for i, filter := range f.Filters {
// 		// filters[i] = filter.ToBson()
// 	}
// 	return bson.M{"$and": filters}
// }
