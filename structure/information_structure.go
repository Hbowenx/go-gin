package structure

import "time"

type Information struct {
	Id        int         `json:"id"`
	MasterMap string      `json:"master_map"`
	Title     string      `json:"title"`
	Text      string      `json:"text"`
	Images []string `json:"images"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type InformationAdd struct {
	MasterMap string `gorm:"column:master_map"`
	Title   string `gorm:"column:title"`
	Text    string `gorm:"column:text"`
	Images    []string `gorm:"type:json;column:images"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}