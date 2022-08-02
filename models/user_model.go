package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	Id   string `json:"id,omitempty"`
	Name string
}

type User struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Roles []Role
}
