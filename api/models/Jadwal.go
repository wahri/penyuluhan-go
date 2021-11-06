package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Jadwal struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Hari_jam   time.Time `json:"hari_jam"`
	Materi     string    `json:"materi"`
	Lokasi     string    `json:"lokasi"`
	PenyuluhID uint32
	MajelisID  uint32
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Jadwal) Prepare() {
	u.ID = 0
	u.Materi = html.EscapeString(strings.TrimSpace(u.Materi))
	u.Lokasi = html.EscapeString(strings.TrimSpace(u.Lokasi))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *Jadwal) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Materi == "" {
			return errors.New("Required Materi")
		}
		if u.Lokasi == "" {
			return errors.New("Required Lokasi")
		}
		return nil
	default:
		if u.Materi == "" {
			return errors.New("Required Materi")
		}
		if u.Lokasi == "" {
			return errors.New("Required Lokasi")
		}
		return nil
	}
}

func (p *Jadwal) SaveJadwal(db *gorm.DB) (*Jadwal, error) {
	var err error
	err = db.Debug().Model(&Jadwal{}).Create(&p).Error
	if err != nil {
		return &Jadwal{}, err
	}
	return p, nil
}

func (p *Jadwal) FindAllJadwal(db *gorm.DB) (*[]Jadwal, error) {
	var err error
	Jadwals := []Jadwal{}
	err = db.Debug().Model(&Jadwal{}).Limit(100).Find(&Jadwals).Error
	if err != nil {
		return &[]Jadwal{}, err
	}
	return &Jadwals, nil
}

func (p *Jadwal) FindJadwalByID(db *gorm.DB, pid uint32) (*Jadwal, error) {
	var err error
	err = db.Debug().Model(&Jadwal{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Jadwal{}, err
	}
	return p, nil
}
