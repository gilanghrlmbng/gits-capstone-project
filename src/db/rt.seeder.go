package db

import (
	"math/rand"
	"src/entity"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func SeedRT(db *gorm.DB) []string {
	// Data 1
	entropy1 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id1 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy1).String()
	data1 := entity.Rt{
		Id:           Id1,
		NamaRt:       "RT 01",
		NamaRw:       "RW 01",
		Kelurahan:    "Cipanas",
		Kecamatan:    "Cipanas",
		Kota:         "Bogor",
		Provinsi:     "Jawa Barat",
		BiayaBulanan: 50000,
		KodeRT:       "0MK2Rr",
		CreatedAt:    time.Now(),
	}
	db.Create(&data1)

	// Data 2
	entropy2 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id2 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy2).String()
	data2 := entity.Rt{
		Id:           Id2,
		NamaRt:       "RT 02",
		NamaRw:       "RW 05",
		Kelurahan:    "Cidingin",
		Kecamatan:    "Cibasah",
		Kota:         "Bandung",
		Provinsi:     "Jawa Barat",
		BiayaBulanan: 70000,
		KodeRT:       "93KrsT",
		CreatedAt:    time.Now(),
	}
	db.Create(&data2)

	// Data 3
	entropy3 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id3 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy3).String()
	data3 := entity.Rt{
		Id:           Id3,
		NamaRt:       "RT 12",
		NamaRw:       "RW 09",
		Kelurahan:    "Ciamis",
		Kecamatan:    "Ciamis",
		Kota:         "Ciamis",
		Provinsi:     "Jawa Barat",
		BiayaBulanan: 30000,
		KodeRT:       "02KsiT",
		CreatedAt:    time.Now(),
	}
	db.Create(&data3)
	return []string{Id1, Id2, Id3}
}
