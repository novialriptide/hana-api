package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Source struct {
	ID         primitive.ObjectID `bson:"_id"`
	SourceID   string             `bson:"source_id"`
	SourceName string             `bson:"source_name"`
}
