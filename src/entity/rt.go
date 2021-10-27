package entity

import (
	"time"

	"gorm.io/gorm"
)

type Rt struct {
	Id           string          `gorm:"type:varchar(50);primaryKey" json:"id"`
	PengurusRT   []PengurusRT    `gorm:"foreignKey:id_rt;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"pengurusRT"`
	NamaRt       string          `gorm:"type:varchar(50);not null" json:"nama_rt"`
	NamaRw       string          `gorm:"type:varchar(50);not null" json:"nama_rw"`
	Kelurahan    string          `gorm:"type:varchar(50);not null" json:"kelurahan"`
	Kecamatan    string          `gorm:"type:varchar(50);not null" json:"kecamatan"`
	Kota         string          `gorm:"type:varchar(50);not null" json:"kota"`
	Provinsi     string          `gorm:"type:varchar(50);not null" json:"provinsi"`
	BiayaBulanan int64           `gorm:"not null" json:"biaya_bulanan"`
	CreatedAt    time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt    time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Rt) TableName() string {
	return "rt"
}
