package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID         primitive.ObjectID `bson:"_id"`
	Text       string             `json:"text"`
	Title      string             `json:"title"`
	Created_at string             `json:"created_at"`
	Updated_at string             `json:"updated_at"`
	Note_id    string             `json:"note_id"`
}
