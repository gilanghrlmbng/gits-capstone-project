package db

import (
	"math/rand"
	"src/entity"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func SeedInformasi(db *gorm.DB, listRT []string) []string {
	// Data 1
	entropy1 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id1 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy1).String()
	data1 := entity.Informasi{
		Id:        Id1,
		IdRT:      listRT[0],
		Judul:     "Kerja Bakti",
		Detail:    "Mie Goreng Mantap Pake Telor",
		Gambar:    "https://dummyimage.com/500x500/eee/fff&text=MG",
		Kategori:  "Kegiatan",
		Lokasi:    "Bawah langit",
		CreatedAt: time.Now(),
	}

	db.Create(&data1)

	// Data 2
	entropy2 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id2 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy2).String()
	data2 := entity.Informasi{
		Id:        Id2,
		IdRT:      listRT[0],
		Judul:     "Kerja Rodi",
		Detail:    "Mie Rebus Mantap Pake Telor",
		Gambar:    "https://dummyimage.com/500x500/eee/fff&text=MR",
		Kategori:  "Informasi",
		Lokasi:    "Atas bumi",
		CreatedAt: time.Now(),
	}

	db.Create(&data2)

	return []string{Id1, Id2}
}
