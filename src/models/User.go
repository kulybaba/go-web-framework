package models

import "time"

type User struct {
	ID        uint   `gorm:"PRIMARY_KEY"`
	FirstName string `gorm:"column:firstName", type:VARCHAR(30); NOT NULL"`
	LastName  string `gorm:"column:lastName", "type:VARCHAR(30); NOT NULL"`
	Email     string `gorm:"type:VARCHAR(40); UNIQUE; NOT NULL"`
	Password  string `gorm:"type:VARCHAR(255); NOT NULL"`
	Created   time.Time
	Updated   time.Time
}
