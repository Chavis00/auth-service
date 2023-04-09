package models

import (
	pkg "authservice/pkg"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelAuth struct {
	ID        string    `json:"id" gorm:"primary_key;"`
	Fullname  string    `json:"fullname,omitempty" gorm:"type:varchar(255);unique;not null"`
	Email     string    `json:"email,omitempty" gorm:"type:varchar(255);unique;not null"`
	Password  string    `json:"password,omitempty" gorm:"type:varchar(255);not null"`
	Active    bool      `json:"active,omitempty" gorm:"type:bool;default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (model *ModelAuth) BeforeCreate(db *gorm.DB) error {
	model.ID = uuid.New().String()
	hashedPassword := pkg.HashPassword(model.Password)
	model.Password = hashedPassword
	model.CreatedAt = time.Now().Local()
	return nil
}
