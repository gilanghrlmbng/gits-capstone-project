package db

import (
	"math/rand"
	"src/entity"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func SeedKeluarga(db *gorm.DB, listRT []string) []string {
	// Data 1
	entropy1 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id1 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy1).String()
	data1 := entity.Keluarga{
		Id:           Id1,
		IdRT:         listRT[0],
		Nama:         "Keluarga Pak Agus",
		NamaToko:     "Keluarga Pak Agus",
		Alamat:       "Rumah No 7",
		Gambar:       "Toko Agustina_2021_12-28_17_15_59",
		KodeKeluarga: "As3ZGx",
		CreatedAt:    time.Now(),
	}

	db.Create(&data1)

	// Data 2
	entropy2 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id2 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy2).String()
	data2 := entity.Keluarga{
		Id:           Id2,
		IdRT:         listRT[0],
		Nama:         "Keluarga Pak Aleks",
		NamaToko:     "Keluarga Pak Aleks",
		Alamat:       "No 8, Tetanggaan sama Pak Agus",
		Gambar:       "Keluarga Pak Aleks_2022_01-02_09_53_34",
		KodeKeluarga: "0MK2Rr",
		CreatedAt:    time.Now(),
	}

	db.Create(&data2)

	// Data 3
	entropy3 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id3 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy3).String()
	data3 := entity.Keluarga{
		Id:           Id3,
		IdRT:         listRT[0],
		Nama:         "Keluarga Lord Rangga",
		NamaToko:     "Keluarga Lord Rangga",
		Alamat:       "Rumah No 1",
		Gambar:       "Keluarga Lord Rangga_2022_01-02_09_56_13",
		KodeKeluarga: "4Kd9D3",
		CreatedAt:    time.Now(),
	}

	db.Create(&data3)

	// Data 4
	entropy4 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id4 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy4).String()
	data4 := entity.Keluarga{
		Id:           Id4,
		IdRT:         listRT[1],
		Nama:         "Keluarga Bu Angelina",
		NamaToko:     "Keluarga Bu Angelina",
		Alamat:       "Rumah No 6",
		Gambar:       "Keluarga Pak Agus_2022_01-02_09_51_04",
		KodeKeluarga: "20KrNd",
		CreatedAt:    time.Now(),
	}

	db.Create(&data4)

	// Data 5
	entropy5 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id5 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy5).String()
	data5 := entity.Keluarga{
		Id:           Id5,
		IdRT:         listRT[1],
		Nama:         "Keluarga Lord Rangga",
		NamaToko:     "Keluarga Lord Rangga",
		Alamat:       "Rumah No 666",
		Gambar:       "default_image",
		KodeKeluarga: "1MTpd4",
		CreatedAt:    time.Now(),
	}

	db.Create(&data5)

	return []string{Id1, Id2, Id3, Id4, Id5}
}
