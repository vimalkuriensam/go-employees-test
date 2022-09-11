package models

import (
	"time"
)

type Employee struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Age       int       `bson:"age" json:"age"`
	Email     string    `bson:"email" json:"email"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type EmployeeDeleteData struct {
	ID          string   `json:"id"`
	DeleteCount int      `json:"deleteCount"`
	DeleteData  Employee `json:"deleteData"`
}
