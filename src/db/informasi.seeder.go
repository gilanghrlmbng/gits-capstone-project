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
		Judul:     "RT 1 Kebanjiran",
		Detail:    "RT 1 lagi kebanjiran, jadi pada ngungsi ke RT 3",
		Gambar:    "Banjir Lagi_2022_01-02_00_11_57",
		Kategori:  "Informasi",
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
		Judul:     "Bantuan untuk Semeru",
		Detail:    "Bantuan bahan pangan papan sandang untuk korban semeru",
		Gambar:    "Bantuan untuk Semeru_2022_01-05_15_58_40",
		Kategori:  "Kegiatan",
		Lokasi:    "Atas bumi",
		CreatedAt: time.Now(),
	}

	db.Create(&data2)

	// Data 3
	entropy3 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id3 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy3).String()
	data3 := entity.Informasi{
		Id:        Id3,
		IdRT:      listRT[0],
		Judul:     "Pos Satpam Kebakaran",
		Detail:    "Pos satpam kebakaran diduga dibakar oleh mas anan",
		Gambar:    "Kebakaran_2021_12-28_20_38_07",
		Kategori:  "Informasi",
		Lokasi:    "Atas bumi",
		CreatedAt: time.Now(),
	}

	db.Create(&data3)

	return []string{Id1, Id2}
}
