package models

import "time"

type User struct {
	ID               string    `json:"ID" bson:"_id" msgpack:"_id"`
	FirstName        string    `json:"first_name" bson:"first_name" msgpack:"first_name"`
	LastName         string    `json:"last_name" bson:"last_name" msgpack:"last_name"`
	Email            string    `json:"email" bson:"email" msgpack:"email"`
	Password         string    `json:"password" bson:"password" msgpack:"password"`
	Bio              string    `json:"bio" bson:"bio" msgpack:"bio"`
	IsActive         bool      `json:"is_active" bson:"is_active" msgpack:"is_active"` //Account has been confirmed
	AvatarURL        string    `json:"avatar_url" bson:"avatar_url" msgpack:"avatar_url"`
	Website          string    `json:"website" bson:"website" msgpack:"website"`
	InstagramProfile string    `json:"instagram_url" bson:"instagram_url" msgpack:"instagram_url"`
	TwitterProfile   string    `json:"twitter_url" bson:"twitter_url" msgpack:"twitter_url"`
	Location         string    `json:"location" bson:"location" msgpack:"location"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at" msgpack:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at" msgpack:"updated_at"`
	Gender           string    `json:"gender" bson:"gender" msgpack:"gender"`
	PasswordHash     []byte
}

func (m *User) TableName() string {
	return "users"
}
