package entity

import (
	"net/http"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type Informasi struct {
	Id        string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdRT      string          `gorm:"type:varchar(50);not null" json:"id_rt" form:"id_rt"`
	Gambar    string          `gorm:"type:varchar(60);not null" json:"gambar" form:"gambar"`
	Detail    string          `gorm:"type:varchar(60); not null" json:"detail" form:"detail"`
	Kategori  string          `gorm:"type:varchar(60); not null" json:"kategori" form:"kategori"`
	Lokasi    string          `gorm:"type:varchar(60); not null" json:"lokasi" form:"lokasi"`
	CreatedBy string          `gorm:"type:varchar(50);not null" json:"created_by" form:"created_by"`
	CreatedAt time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt time.Time       `gorm:"type:timestamptz" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Informasi) TableName() string {
	return "informasi"
}

func (i Informasi) ValidateCreate() utils.Error {
	if i.Gambar == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Gambar tidak boleh kosong",
		}
	}

	if i.Detail == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Detail tidak boleh kosong",
		}
	}

	if i.CreatedBy == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Author tidak boleh kosong",
		}
	}

	if i.Lokasi == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Lokasi tidak boleh kosong",
		}
	}

	if i.Kategori == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Kategori tidak boleh kosong",
		}
	}
	return utils.Error{}
}
