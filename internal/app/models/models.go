package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID               uint   `gorm:"primaryKey"`
	Name             string `gorm:"size:100;not null"`
	Email            string `gorm:"size:100;unique;not null"`
	RegistrationDate string
	Role             string `gorm:"size:50;not null"`
	Tasks            []Task `gorm:"foreignKey:AssigneeID"`
}

type Task struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:100;not null"`
	Description string `gorm:"size:255"`
	Priority    string `gorm:"size:50"`
	Status      string `gorm:"size:50"`
	AssigneeID  uint
	ProjectID   uint
	CreatedAt   string
	CompletedAt string
	User        User    `gorm:"foreignKey:AssigneeID"`
	Project     Project `gorm:"foreignKey:ProjectID"`
}

type Project struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:100;not null"`
	Description string `gorm:"size:255"`
	StartDate   string
	EndDate     string
	ManagerID   uint
	Tasks       []Task `gorm:"foreignKey:ProjectID"`
}