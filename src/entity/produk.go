package entity

type Produk struct {
	Id         string `gorm:"type:varchar(50);primaryKey" json:"id"`
	IdKeluarga string `gorm:"type:varchar(50);not null" json:"id_keluarga"`
	Nama       string `gorm:"type:varchar(50);not null" json:"nama"`
	Detail     string `gorm:"not null" json:"detail"`
	Harga      int64  `gorm:"not null" json:"harga"`
}

func (Produk) TableName() string {
	return "produk"
}
