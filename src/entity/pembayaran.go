package entity

import (
	"time"

	"gorm.io/gorm"
)

type Pembayaran struct {
	Id                string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdOrder           string          `gorm:"type:varchar(50);not null" json:"id_order" form:"id_order"`
	IdKeluargaPembeli string          `gorm:"type:varchar(50);not null" json:"id_keluarga_pembeli" form:"id_keluarga_pembeli"`
	IdKeluargaPenjual string          `gorm:"type:varchar(50);not null" json:"id_keluarga_penjual" form:"id_keluarga_penjual"`
	Jumlah_pembayaran int64           `gorm:"not null" json:"jumlah_pembayaran" form:"jumlah_pembayaran"`
	Jenis             string          `gorm:"type:varchar(50);not null" json:"jenis" form:"jenis"`
	CreatedAt         time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"type:timestamptz" json:"updated_at"`
	DeletedAt         *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Pembayaran) TableName() string {
	return "pembayaran"
}
