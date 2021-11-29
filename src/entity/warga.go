package entity

import (
	"fmt"
	"net/http"
	"net/mail"
	"src/utils"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Warga struct {
	Id           string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	Order        []Order         `gorm:"foreignKey:id_pembayaran;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"id_order,omitempty" form:"id_order"`
	IdKeluarga   string          `gorm:"type:varchar(50);not null" json:"id_keluarga" form:"id_keluarga"`
	KodeKeluarga string          `gorm:"type:varchar(100); not null" json:"kode_keluarga,omitempty" form:"kode_keluarga"`
	Nama         string          `gorm:"type:varchar(100);not null" json:"nama" form:"nama"`
	NoHandphone  string          `gorm:"type:varchar(20);not null" json:"no_hp" form:"no_hp"`
	Gender       string          `gorm:"type:varchar(20);not null" json:"gender" form:"gender"`
	Gambar       string          `gorm:"not null" json:"gambar" form:"gambar"`
	Email        string          `gorm:"type:varchar(120);not null" json:"email" form:"email"`
	Password     string          `gorm:"type:varchar(100);not null" json:"password" form:"password"`
	CreatedAt    time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt    time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Warga) TableName() string {
	return "warga"
}

func (w Warga) ValidateCreate() utils.Error {
	if strings.HasPrefix(w.NoHandphone, "62") {
		w.NoHandphone = fmt.Sprintf("0%s", strings.SplitAfter(w.NoHandphone, "62"))
	}
	if strings.HasPrefix(w.NoHandphone, "+62") {
		w.NoHandphone = fmt.Sprintf("0%s", strings.SplitAfter(w.NoHandphone, "+62"))
	}
	if w.Nama == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama tidak boleh kosong",
		}
	}
	if _, err := mail.ParseAddress(w.Email); err != nil {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Email tidak valid",
		}
	}
	if w.Password == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Password tidak boleh kosong",
		}
	}
	if w.KodeKeluarga == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Kode Keluarga tidak boleh kosong",
		}
	}
	if w.Gambar == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Gambar tidak boleh kosong",
		}
	}
	if len(w.NoHandphone) < 10 && len(w.NoHandphone) > 13 {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nomor handphone tidak valid (min 10 angka, max 13 angka)",
		}
	}
	if w.Gender != "laki-laki" && w.Gender != "perempuan" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "gender tidak valid",
		}
	}
	return utils.Error{}
}
