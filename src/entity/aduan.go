package entity

import (
	"net/http"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type Aduan struct {
	Id        string `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdRT      string `gorm:"type:varchar(50);not null" json:"id_rt" form:"id_rt"`
	IdWarga   string `gorm:"type:varchar(50);not null" json:"id_warga" form:"id_warga"`
	Judul     string `gorm:"not null" json:"judul" form:"judul"`
	Gambar    string `json:"gambar" form:"gambar"`
	Deskripsi string `gorm:"not null" json:"deskripsi" form:"deskripsi"`
	Status    string
	CreatedBy string          `gorm:"not null" json:"createdBy" form:"createdBy"`
	CreatedAt time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt time.Time       `gorm:"type:timestamptz" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Aduan) TableName() string {
	return "aduan"
}

func (a Aduan) ValidateCreate() utils.Error {
	if a.Judul == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Judul tidak boleh kosong",
		}
	}
	if a.Deskripsi == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Deskripsi tidak boleh kosong",
		}
	}

	return utils.Error{}
}
