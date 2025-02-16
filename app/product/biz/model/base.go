package model

import "time"

type Base struct {
	ID       int `gorm:"primaryKey"`
	CreateAt time.Time
	UpdateAt time.Time
}
