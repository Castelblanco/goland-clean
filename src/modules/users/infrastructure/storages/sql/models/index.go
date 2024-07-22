package models

import (
	"github.com/Castelblanco/goland-clean/src/app/db"
	"gorm.io/gorm"
)

type UserDAL struct {
	gorm.Model

	ID        string        `gorm:"type:uuid;primaryKey"`
	FirstName string        `gorm:"not null"`
	LastName  string        `gorm:"not null"`
	Email     string        `gorm:"not null;uniqueIndex"`
	Tasks     []UserTaskDAL `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (UserDAL) TableName() string {
	return "USERS"
}

type UserTaskDAL struct {
	gorm.Model

	ID          string `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Done        bool   `gorm:"default:false"`
	UserID      string `gorm:"type:uuid"`
}

func (UserTaskDAL) TableName() string {
	return "TASKS"
}

func Migrations() {
	db.DB.AutoMigrate(UserDAL{}, UserTaskDAL{})
}
