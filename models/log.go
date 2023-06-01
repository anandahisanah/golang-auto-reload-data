package models

import "time"

type Log struct {
	Id            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	StatusWaterId uint       `gorm:"not null;foreignKey:StatusWaterId" json:"status_water_id"`
	StatusWater   Status    `gorm:"foreignKey:StatusWaterId" json:"-"`
	StatusWindId  uint       `gorm:"not null;foreignKey:StatusWindId" json:"status_wind_id"`
	StatusWind    Status    `gorm:"foreignKey:StatusWindId" json:"-"`
	Water         int        `gorm:"not null" json:"water"`
	Wind          int        `gorm:"not null" json:"wind"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
}
