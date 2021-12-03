package entity

import (
	"net/http"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type Persuratan struct {
	Id        string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdRT      string          `gorm:"type:varchar(50);not null" json:"id_rt" form:"id_rt"`
	Judul     string          `gorm:"type:varchar(50);not null" json:"judul" form:"judul"`
	Penerima  string          `gorm:"type:varchar(50);not null" json:"penerima" form:"penerima"`
	Tanggal   string          `gorm:"type:varchar(50);not null" json:"tanggal" form:"tanggal"`
	Keperluan string          `gorm:"not null" json:"keperluan" form:"keperluan"`
	Link      string          `gorm:"not null" json:"link" form:"link"`
	Status    string          `gorm:"type:varchar(50); not null" json:"status" form:"status"`
	CreatedAt time.Time       `gorm:"type:timestamptz; not null" json:"created_at"`
	UpdatedAt time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Persuratan) TableName() string {
	return "persuratan"
}

func (p Persuratan) ValidateCreate() utils.Error {
	if p.Judul == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Judul surat tidak boleh kosong",
		}
	}
	if p.Penerima == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Penerima surat tidak boleh kosong",
		}
	}
	if p.Tanggal == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Tanggal dibuat surat tidak boleh kosong",
		}
	}
	if p.Keperluan == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Keperluan tidak boleh kosong",
		}
	}

	if p.Status == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Status tidak boleh kosong",
		}
	}
	return utils.Error{}
}
