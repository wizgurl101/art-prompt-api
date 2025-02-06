package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Token    string             `json:"token,omitempty"`
}

type contextKey string

const UserContextKey contextKey = "user"
