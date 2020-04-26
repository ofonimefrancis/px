package models

type User struct {
	FirstName string `json:"first_name" bson:"first_name" msgpack:"first_name"`
	LastName  string `json:"last_name" bson:"last_name" msgpack:"last_name"`
	Email     string `json:"email" bson:"email" msgpack:"email"`
	Password  string `json:"password" bson:"password" msgpack:"password"`
}
