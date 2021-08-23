package entity

import (
	"time"
)

type Shop struct {
	Id   int   `json:"id"  bson:"id"`
	Name    string   `json:"name" bson:"name"`
	Code    string   `json:"code" bson:"code,omitempty"`
	SystemUrl    string   `json:"system_url" bson:"system_url"`

	Phone   *string    `json:"phone" bson:"phone"`
	Mobile  *string    `json:"mobile" bson:"mobile"`

	CreatedUser    string   `json:"created_user,omitempty" bson:"created_user"`
	UpdatedUser    string   `json:"updated_user,omitempty" bson:"updated_user"`
	CreatedDate  *time.Time   `json:"created_date,omitempty" bson:"created_date"`
	UpdatedDate    *time.Time   `json:"updated_date,omitempty" bson:"updated_date"`
}