package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Laporan struct {
	ID             uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Judul_laporan  string `json:"Judul_laporan"`
	Isi_kegiatan   string `json:"Isi_kegiatan"`
	Jumlah_peserta int    `json:"Jumlah_peserta"`
	JadwalID       uint32
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Laporan) Prepare() {
	u.ID = 0
	u.Judul_laporan = html.EscapeString(strings.TrimSpace(u.Judul_laporan))
	u.Isi_kegiatan = html.EscapeString(strings.TrimSpace(u.Isi_kegiatan))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *Laporan) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Judul_laporan == "" {
			return errors.New("Required Judul_laporan")
		}
		if u.Isi_kegiatan == "" {
			return errors.New("Required Isi_kegiatan")
		}
		return nil
	default:
		if u.Judul_laporan == "" {
			return errors.New("Required Judul_laporan")
		}
		if u.Isi_kegiatan == "" {
			return errors.New("Required Isi_kegiatan")
		}
		return nil
	}
}

func (p *Laporan) SaveLaporan(db *gorm.DB) (*Laporan, error) {
	var err error
	err = db.Debug().Model(&Laporan{}).Create(&p).Error
	if err != nil {
		return &Laporan{}, err
	}
	return p, nil
}

func (p *Laporan) FindAllLaporan(db *gorm.DB) (*[]Laporan, error) {
	var err error
	Laporans := []Laporan{}
	err = db.Debug().Model(&Laporan{}).Limit(100).Find(&Laporans).Error
	if err != nil {
		return &[]Laporan{}, err
	}
	return &Laporans, nil
}

func (p *Laporan) FindLaporanByID(db *gorm.DB, pid uint32) (*Laporan, error) {
	var err error
	err = db.Debug().Model(&Laporan{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Laporan{}, err
	}
	return p, nil
}
