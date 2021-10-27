package entity

import (
	"net/http"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type Rt struct {
	Id           string          `gorm:"type:varchar(50);primaryKey" json:"id"`
	PengurusRT   []PengurusRT    `gorm:"foreignKey:id_rt;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"pengurus_rt"`
	Keluarga     []Keluarga      `gorm:"foreignKey:id_rt;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"keluarga"`
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

func (rt Rt) ValidateCreate() utils.Error {
	if rt.NamaRt == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama RT tidak boleh kosong",
		}
	}
	if rt.NamaRw == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama RW tidak boleh kosong",
		}
	}
	if rt.Kelurahan == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama Kelurahan tidak boleh kosong",
		}
	}
	if rt.Kecamatan == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama Kecamatan tidak boleh kosong",
		}
	}
	if rt.Kota == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama Kota tidak boleh kosong",
		}
	}
	if rt.Provinsi == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama Provinsi tidak boleh kosong",
		}
	}
	if rt.BiayaBulanan == 0 {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Biaya Bulanan tidak boleh kosong",
		}
	}
	return utils.Error{}
}
