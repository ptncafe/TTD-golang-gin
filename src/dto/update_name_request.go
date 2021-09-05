package dto

import "time"

type UpdateNameRequest struct {
	Id   int   `json:"id"  bson:"id"`
	Name    string   `json:"name" bson:"name"`
	Code    string   `json:"code" bson:"code"`
	UpdatedDate    *time.Time   `json:"updated_date,omitempty" bson:"updated_date"`
	UpdatedUser    string   `json:"updated_user,omitempty" bson:"updated_user"`

}