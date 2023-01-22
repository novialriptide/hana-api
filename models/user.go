package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	UserID         string             `bson:"_id"`
	CanUploadMusic bool               `bson:"can_upload_music"`
	CanDeleteMusic bool               `bson:"can_delete_music"`
}
