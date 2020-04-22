package models

import (
	"time"
)

type GroupUser struct {
	ID        uint64     `gorm:"primary_key"`
	GroupId   uint64     `json:"-"`
	UserId    uint64     `json:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
