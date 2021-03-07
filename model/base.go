package model

import (
	"time"
)

type BaseModel struct {
	ID        uint64     `gorm:"-"`
	CreatedAt *time.Time `gorm:"-"`
	UpdatedAt *time.Time `gorm:"-"`
	DeletedAt *time.Time `gorm:"-"`
}
