package entity

import (
	"net/http"
	"net/mail"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type Warga struct {
	Id         string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdKeluarga string          `gorm:"type:varchar(50);not null" json:"id_keluarga" form:"id_keluarga"`
	Nama       string          `gorm:"type:varchar(50);not null" json:"nama" form:"nama"`
	Email      string          `gorm:"type:varchar(120);not null" json:"email" form:"email"`
	Password   string          `gorm:"type:varchar(100);not null" json:"password" form:"password"`
	CreatedAt  time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt  time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Warga) TableName() string {
	return "warga"
}

func (w Warga) ValidateCreate() utils.Error {
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
	return utils.Error{}
}
