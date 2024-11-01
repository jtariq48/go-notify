package models

import "time"

type APIToken struct {
	Token     string    `json:"token"`
	SecretKey string    `json:"secret_key"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
