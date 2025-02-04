package f9z

import "go.mongodb.org/mongo-driver/bson/primitive"

// ConvertToObjectID 将id进行转换
func ConvertToObjectID(ids []string) ([]primitive.ObjectID, error) {
	var objectIDs []primitive.ObjectID

	for _, id := range ids {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, objectID)
	}
	return objectIDs, nil
}
