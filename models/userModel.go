package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    string             `json :"first_name" validate :"required, min = 2 , max = 6"`
	Last_name     string             `json :"last_name" validate: "required min 2 max = 6"`
	Email         string             `json:"email" validate: "required"`
	Pasword       string             `json:"password" validate: "required min = 6"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	Token         string             `json:"token" validate:"required"`
	Refresh_token string             `	json:"refresh_token" validate :"required"`
	User_id       string             `json:"user_id" validate:"required"`
	Phone         string             `json:"phone" validate:"required"`
	User_type     string             `json:"user_type" validate:"required eq = ADMIN | eq = USER"`
}
