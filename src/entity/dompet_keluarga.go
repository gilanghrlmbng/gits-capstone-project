package entity

import (
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type DompetKeluarga struct {
	Id         string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdKeluarga string          `gorm:"type:varchar(50);not null" json:"id_keluarga" form:"id_keluarga"`
	Jumlah     int64           `gorm:"not null" json:"jumlah" form:"jumlah"`
	CreatedAt  time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt  time.Time       `gorm:"type:timestamptz" json:"updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (DompetKeluarga) TableName() string {
	return "dompet_keluarga"
}

func (d DompetKeluarga) ValidateCreate() utils.Error {

	return utils.Error{}
}
