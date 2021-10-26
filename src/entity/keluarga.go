package entity

type Keluarga struct {
	Id      string    `gorm:"type:varchar(50);primaryKey" json:"id"`
	Nama    string    `gorm:"type:varchar(50);not_null" json:"nama"`
	Warga   []Warga   `gorm:"foreignKey:id_keluarga;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"warga"`
	Tagihan []Tagihan `gorm:"foreignKey:id_keluarga;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"tagihan"`
	Produk  []Produk  `gorm:"foreignKey:id_keluarga;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"produk"`
	Alamat  string    `gorm:"type:varchar(50);not_null" json:"alamat"`
}

func (Keluarga) TableName() string {
	return "keluarga"
}
