package entity

import (
	"net/http"
	"net/mail"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type PengurusRT struct {
	Id          string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdRT        string          `gorm:"type:varchar(50);not null" json:"id_rt" form:"id_rt"`
	NoHandphone string          `gorm:"type:varchar(20);not null" json:"no_hp" form:"no_hp"`
	KodeRT      string          `gorm:"type:varchar(100); not null" json:"kode_rt,omitempty" form:"kode_rt"`
	Gender      string          `gorm:"type:varchar(20);not null" json:"gender" form:"gender"`
	Nama        string          `gorm:"type:varchar(50);not null" json:"nama" form:"nama"`
	Gambar      string          `gorm:"not null" json:"gambar" form:"gambar"`
	Email       string          `gorm:"type:varchar(120);not null" json:"email" form:"email"`
	Password    string          `gorm:"type:varchar(100);not null" json:"password" form:"password"`
	CreatedAt   time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (PengurusRT) TableName() string {
	return "pengurus_rt"
}

func (prt PengurusRT) ValidateCreate() utils.Error {
	if prt.Nama == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama tidak boleh kosong",
		}
	}
	if _, err := mail.ParseAddress(prt.Email); err != nil {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Email tidak valid",
		}
	}
	if prt.KodeRT == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Kode RT tidak boleh kosong",
		}
	}
	if prt.Password == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Password tidak boleh kosong",
		}
	}
	if prt.Gambar == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Gambar tidak boleh kosong",
		}
	}
	if len(prt.NoHandphone) < 10 && len(prt.NoHandphone) > 13 {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nomor handphone tidak valid (min 10 angka, max 13 angka)",
		}
	}
	if prt.Gender != "laki-laki" && prt.Gender != "perempuan" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "gender tidak valid",
		}
	}
	return utils.Error{}
}
