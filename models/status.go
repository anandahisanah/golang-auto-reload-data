package models

type Status struct {
	Id         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Code       string `gorm:"not null" json:"code"`
	Name       string `gorm:"not null" json:"name"`
	RangeStart int    `gorm:"not null" json:"range_start"`
	RangeEnd   int    `gorm:"not null" json:"range_end"`
}
