package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email    string `gorm:"unique"`
    Password string
    Role     string // "doctor" or "receptionist"
}

type Patient struct {
    gorm.Model
    Name      string
    Age       int
    Gender    string
    Diagnosis string
}