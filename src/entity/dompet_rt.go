package entity

import (
	"net/http"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type DompetRT struct {
	Id        string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdRT      string          `gorm:"type:varchar(50);not null" json:"id_rt" form:"id_rt"`
	Nama      string          `gorm:"type:varchar(50);not null" json:"nama" form:"nama"`
	Detail    string          `gorm:"type:varchar(60); not null" json:"detail" form:"detail"`
	Jumlah    int64           `gorm:"not null" json:"jumlah" form:"jumlah"`
	CreatedAt time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt time.Time       `gorm:"type:timestamptz" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (DompetRT) TableName() string {
	return "dompet_rt"
}

func (d DompetRT) ValidateCreate() utils.Error {
	if d.Nama == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Gambar tidak boleh kosong",
		}
	}

	if d.Detail == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Detail tidak boleh kosong",
		}
	}

	return utils.Error{}
}
