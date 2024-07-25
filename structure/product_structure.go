package structure

import "time"

type Product struct {
	Id        int         `json:"id"`
	Desc 	  string      `json:"desc"`
	Images      interface{} `json:"images"`
	Sort      int	   `json:"sort"`
	LabelID   int     `json:"label_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ProductAdd struct {
	Desc 	  string      `gorm:"column:desc"`
	Images    []string `gorm:"type:json;column:images"`
	Sort      int	   `gorm:"column:sort"`
	LabelId   int     `gorm:"column:label_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}