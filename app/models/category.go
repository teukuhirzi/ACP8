package models

import "time"

type Category struct {
	ID        string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name      string `gorm:"size:100;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
