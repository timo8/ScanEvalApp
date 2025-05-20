package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name               string    `gorm:"not null"`
	Surname            string    `gorm:"not null"`
	BirthDate          time.Time `gorm:"not null"`
	RegistrationNumber int       `gorm:"not null;unique"`
	Room               string    `gorm:"not null"`
	Score              int
	Answers            string
	ExamID             uint `gorm:"not null"`
	Pages              string
}
