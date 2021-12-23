package id

import "go.mongodb.org/mongo-driver/bson/primitive"

// Converts a value into a slice of ObjectIDs
func Slice(value interface{}) []primitive.ObjectID {

	switch v := value.(type) {
	case []string:
		result := make([]primitive.ObjectID, len(v))
		for index := range v {
			result[index] = ID(v[index])
		}
		return result
	}

	return make([]primitive.ObjectID, 0)
}

func SliceOfString(value []primitive.ObjectID) []string {
	result := make([]string, len(value))

	for index := range value {
		result[index] = value[index].Hex()
	}

	return result
}
