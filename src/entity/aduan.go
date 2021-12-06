package entity

import (
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type Aduan struct {
	Id        string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdWarga   string          `gorm:"type:varchar(50);not null" json:"id_warga" form:"id_warga"`
	Gambar    string          `gorm:"not null" json:"gambar" form:"gambar"`
	Deskripsi string          `gorm:"not null" json:"deskripsi" form:"deskripsi"`
	CreatedBy string          `gorm:"not null" json:"createdBy" form:"createdBy"`
	CreatedAt time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt time.Time       `gorm:"type:timestamptz" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Aduan) TableName() string {
	return "aduan"
}

func (a Aduan) ValidateCreate() utils.Error {

	return utils.Error{}
}
