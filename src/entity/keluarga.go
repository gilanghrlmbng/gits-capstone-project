package entity

import (
	"net/http"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type Keluarga struct {
	Id        string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdRT      string          `gorm:"type:varchar(50);not null" json:"id_rt" form:"id_rt"`
	Nama      string          `gorm:"type:varchar(50);not null" json:"nama" form:"nama"`
	Warga     []Warga         `gorm:"foreignKey:id_keluarga;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"warga,omitempty" form:"warga"`
	Tagihan   []Tagihan       `gorm:"foreignKey:id_keluarga;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"tagihan,omitempty" form:"tagihan"`
	Produk    []Produk        `gorm:"foreignKey:id_keluarga;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"produk,omitempty" form:"produk"`
	Alamat    string          `gorm:"type:varchar(50)" json:"alamat,omitempty" form:"alamat"`
	CreatedAt time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt time.Time       `gorm:"type:timestamptz" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Keluarga) TableName() string {
	return "keluarga"
}

func (k Keluarga) ValidateCreate() utils.Error {
	if k.Nama == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama tidak boleh kosong",
		}
	}
	return utils.Error{}
}
