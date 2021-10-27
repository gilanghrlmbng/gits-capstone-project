package entity

import (
	"time"

	"gorm.io/gorm"
)

type Tagihan struct {
	Id         string          `gorm:"type:varchar(50);primaryKey" json:"id"`
	IdKeluarga string          `gorm:"type:varchar(50);not null" json:"id_keluarga"`
	Nama       string          `gorm:"type:varchar(50);not null" json:"nama"`
	Detail     string          `gorm:"not null" json:"detail"`
	Jumlah     int64           `gorm:"not null" json:"jumlah"`
	CreatedAt  time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt  time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Tagihan) TableName() string {
	return "tagihan"
}
