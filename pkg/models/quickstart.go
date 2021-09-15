package models

import "gorm.io/gorm"

// Quickstart represents the quickstart json content
type Quickstart struct {
	gorm.Model
	ID    uint `gorm:"primaryKey"`
	Title string
}