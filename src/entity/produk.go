package entity

import (
	"time"

	"gorm.io/gorm"
)

type Produk struct {
	Id         string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdKeluarga string          `gorm:"type:varchar(50);not null" json:"id_keluarga" form:"id_keluarga"`
	Nama       string          `gorm:"type:varchar(50);not null" json:"nama" form:"nama"`
	Detail     string          `gorm:"not null" json:"detail" form:"detail"`
	Harga      int64           `gorm:"not null" json:"harga" form:"harga"`
	CreatedAt  time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt  time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Produk) TableName() string {
	return "produk"
}
