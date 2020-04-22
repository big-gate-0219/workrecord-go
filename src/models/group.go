package models

import (
	"time"
)

type Group struct {
	ID        uint64     `json:"group_id"; gorm:"primary_key"`
	Name      string     `json:"group_name"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
