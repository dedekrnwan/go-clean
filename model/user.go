package model

type (
	User struct {
		Model

		FirstName string `json:"first_name" bson:"first_name"`
		LastName  string `json:"last_name" bson:"last_name"`
		Email     string `json:"email" bson:"email"`
		Phone     string `json:"phone" bson:"phone"`
		IsActive  *bool  `json:"is_active" bson:"is_active"`
		Password  string `json:"password" bson:"password"`
	}
)
