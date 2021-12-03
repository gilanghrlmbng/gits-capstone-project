package entity

import (
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type DompetRT struct {
	Id        string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdRT      string          `gorm:"type:varchar(50);not null" json:"id_rt" form:"id_rt"`
	Jumlah    int64           `gorm:"not null" json:"jumlah" form:"jumlah"`
	CreatedAt time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt time.Time       `gorm:"type:timestamptz" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (DompetRT) TableName() string {
	return "dompet_rt"
}

func (d DompetRT) ValidateCreate() utils.Error {

	return utils.Error{}
}
