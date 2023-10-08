package models

import "time"

const (
	Male   Gender = "male"
	Female Gender = "female"
	Others Gender = "others"
)

type (
	Gender string

	User struct {
		ID        string    `json:"id"`
		Username  string    `json:"username" validate:"required"`
		FirstName string    `json:"first_name" validate:"required"`
		LastName  string    `json:"last_name" validate:"required"`
		Email     string    `json:"email" validate:"required,email"`
		Gender    Gender    `json:"gender" validate:"required,oneof=male female others"`
		Password  string    `json:"password" validate:"required,min=6"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
