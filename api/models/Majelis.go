package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Majelis struct {
	ID           uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Nama_majelis string `gorm:"size:255;not null;unique" json:"nama_majelis"`
	Alamat       string `gorm:"size:255;not null;unique" json:"alamat"`
	Jadwals      []Jadwal
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Majelis) Prepare() {
	p.ID = 0
	p.Nama_majelis = html.EscapeString(strings.TrimSpace(p.Nama_majelis))
	p.Alamat = html.EscapeString(strings.TrimSpace(p.Alamat))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Majelis) Validate() error {

	if p.Nama_majelis == "" {
		return errors.New("Required Nama_majelis")
	}
	if p.Alamat == "" {
		return errors.New("Required Alamat")
	}
	return nil
}

func (p *Majelis) SaveMajelis(db *gorm.DB) (*Majelis, error) {
	var err error
	err = db.Debug().Model(&Majelis{}).Create(&p).Error
	if err != nil {
		return &Majelis{}, err
	}
	return p, nil
}

func (p *Majelis) FindAllMajelis(db *gorm.DB) (*[]Majelis, error) {
	var err error
	Majeliss := []Majelis{}
	err = db.Debug().Model(&Majelis{}).Limit(100).Preload("Jadwals.Laporan").Find(&Majeliss).Error
	if err != nil {
		return &[]Majelis{}, err
	}
	return &Majeliss, nil
}

func (p *Majelis) FindMajelisByID(db *gorm.DB, pid uint32) (*Majelis, error) {
	var err error
	err = db.Debug().Model(&Majelis{}).Preload("Jadwals.Laporan").Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Majelis{}, err
	}
	return p, nil
}

// func (p *Majelis) UpdateAMajelis(db *gorm.DB) (*Majelis, error) {

// 	var err error
// 	// db = db.Debug().Model(&Majelis{}).Where("id = ?", pid).Take(&Majelis{}).UpdateColumns(
// 	// 	map[string]interface{}{
// 	// 		"title":      p.Title,
// 	// 		"content":    p.Content,
// 	// 		"updated_at": time.Now(),
// 	// 	},
// 	// )
// 	// err = db.Debug().Model(&Majelis{}).Where("id = ?", pid).Take(&p).Error
// 	// if err != nil {
// 	// 	return &Majelis{}, err
// 	// }
// 	// if p.ID != 0 {
// 	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
// 	// 	if err != nil {
// 	// 		return &Majelis{}, err
// 	// 	}
// 	// }
// 	err = db.Debug().Model(&Majelis{}).Where("id = ?", p.ID).Updates(Majelis{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error
// 	if err != nil {
// 		return &Majelis{}, err
// 	}
// 	if p.ID != 0 {
// 		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
// 		if err != nil {
// 			return &Majelis{}, err
// 		}
// 	}
// 	return p, nil
// }

// func (p *Majelis) DeleteAMajelis(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

// 	db = db.Debug().Model(&Majelis{}).Where("id = ? and author_id = ?", pid, uid).Take(&Majelis{}).Delete(&Majelis{})

// 	if db.Error != nil {
// 		if gorm.IsRecordNotFoundError(db.Error) {
// 			return 0, errors.New("Majelis not found")
// 		}
// 		return 0, db.Error
// 	}
// 	return db.RowsAffected, nil
// }
