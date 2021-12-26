package db

import (
	"math/rand"
	"src/entity"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func SeedProduk(db *gorm.DB, listKeluarga []string) []string {
	// Data 1
	entropy1 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id1 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy1).String()
	data1 := entity.Produk{
		Id:         Id1,
		IdKeluarga: listKeluarga[0],
		Nama:       "Mie Goreng",
		Detail:     "Mie Goreng Mantap Pake Telor",
		Gambar:     "default_image",
		Harga:      20000,
		CreatedAt:  time.Now(),
	}

	db.Create(&data1)

	// Data 2
	entropy2 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id2 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy2).String()
	data2 := entity.Produk{
		Id:         Id2,
		IdKeluarga: listKeluarga[0],
		Nama:       "Mie Rebus",
		Detail:     "Mie Rebus Mantap Pake Telor",
		Gambar:     "default_image",
		Harga:      20000,
		CreatedAt:  time.Now(),
	}

	db.Create(&data2)

	// Data 3
	entropy3 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id3 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy3).String()
	data3 := entity.Produk{
		Id:         Id3,
		IdKeluarga: listKeluarga[0],
		Nama:       "Kentang Mustofa",
		Detail:     "Kering kentang mantap 200gr",
		Gambar:     "default_image",
		Harga:      25000,
		CreatedAt:  time.Now(),
	}

	db.Create(&data3)

	// Data 4
	data4 := entity.Produk{
		Id:         "01FQNTQ9KYRFW9YNMYGAXP1R8M",
		IdKeluarga: listKeluarga[1],
		Nama:       "Kue Donat",
		Detail:     "Donat manis aneka topping",
		Gambar:     "default_image",
		Harga:      3000,
		CreatedAt:  time.Now(),
	}

	db.Create(&data4)

	// Data 5
	data5 := entity.Produk{
		Id:         "01FQNTQ9M1WA992RJY56VTJ4YX",
		IdKeluarga: listKeluarga[1],
		Nama:       "Kue Sus",
		Detail:     "Kue sus manis",
		Gambar:     "default_image",
		Harga:      2500,
		CreatedAt:  time.Now(),
	}

	db.Create(&data5)

	return []string{Id1, Id2, Id3}
}
