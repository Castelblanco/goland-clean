package models

import (
	"github.com/Castelblanco/goland-clean/src/app/db"
	"gorm.io/gorm"
)

type TaskDAL struct {
	gorm.Model

	ID          string `gorm:"type:uuid;primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Done        bool        `gorm:"default:false"`
	UserID      string      `gorm:"type:uuid"`
	User        TaskUserDAL `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (TaskDAL) TableName() string {
	return "TASKS"
}

type TaskUserDAL struct {
	gorm.Model

	ID        string `gorm:"type:uuid;primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null;uniqueIndex"`
}

func (TaskUserDAL) TableName() string {
	return "TASKS"
}

func Migrations() {
	db.DB.AutoMigrate(TaskDAL{}, TaskUserDAL{})
}
