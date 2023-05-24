package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Token     string    `gorm:"-"`
	Dosen     Dosen     `gorm:"foreignKey:UserID"`
	Mahasiswa Mahasiswa `gorm:"foreignKey:UserID"`
}

type Mahasiswa struct {
	gorm.Model
	Mahasiswa  string `json:"mahasiswa" form:"mahasiswa"`
	Email      string `json:"email" form:"email"`
	Name       string `json:"name" form:"name"`
	NIM        string `json:"nim" form:"nim"`
	Image      string `json:"image" form:"image"`
	Phone      string `json:"phone" form:"phone"`
	Jurusan    string `json:"jurusan" form:"jurusan"`
	TahunMasuk string `json:"tahun_masuk" form:"tahun_masuk"`
	IPK        string `json:"ipk" form:"ipk"`
	UserID     uint   `json:"user_id" form:"user_id"`
}
