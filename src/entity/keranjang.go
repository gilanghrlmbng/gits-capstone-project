package entity

import (
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type Keranjang struct {
	Id                string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdWarga           string          `gorm:"type:varchar(50);not null" json:"id_warga" form:"id_warga"`
	IdKeluargaPenjual string          `gorm:"type:varchar(50)" json:"id_keluarga_penjual" form:"id_keluarga_penjual"`
	ItemKeranjang     []ItemKeranjang `gorm:"foreignKey:id_keranjang;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"item_keranjang,omitempty" form:"item_keranjang"`
	Harga_total       int64           `gorm:"not null" json:"harga_total" form:"harga_total"`
	CreatedAt         time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt         *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Keranjang) TableName() string {
	return "keranjang"
}

func (ord Keranjang) ValidateCreate() utils.Error {
	return utils.Error{}
}
