package structure

import "time"

type Label struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Sort int `json:"sort"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}



type LabelAdd struct {
	Name string `gorm:"column:name"`
	Sort   int `gorm:"column:sort"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}


