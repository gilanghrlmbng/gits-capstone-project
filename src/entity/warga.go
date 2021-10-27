package entity

import (
	"time"

	"gorm.io/gorm"
)

type Warga struct {
	Id         string          `gorm:"type:varchar(50);primaryKey" json:"id"`
	IdKeluarga string          `gorm:"type:varchar(50);not null" json:"id_keluarga"`
	Nama       string          `gorm:"type:varchar(50);not null" json:"nama"`
	Alamat     string          `gorm:"type:varchar(50);not null" json:"alamat"`
	CreatedAt  time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt  time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Warga) TableName() string {
	return "warga"
}
