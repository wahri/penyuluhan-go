package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

type Penyuluh struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Nik       string `json:"Nik"`
	Nama      string `json:"Nama"`
	Alamat    string `json:"Alamat"`
	Email     string `json:"Email"`
	No_hp     string `json:"No_hp"`
	Foto      string `json:"Foto"`
	Status    int    `json:"Status"`
	User      User   `gorm:"polymorphic:Level;"`
	Jadwals   []Jadwal
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Penyuluh) Prepare() {
	u.ID = 0
	u.Nik = html.EscapeString(strings.TrimSpace(u.Nik))
	u.Nama = html.EscapeString(strings.TrimSpace(u.Nama))
	u.Alamat = html.EscapeString(strings.TrimSpace(u.Alamat))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.No_hp = html.EscapeString(strings.TrimSpace(u.No_hp))
	u.Foto = html.EscapeString(strings.TrimSpace(u.Foto))
	u.Status = 1
	u.User.Username = html.EscapeString(strings.TrimSpace(u.User.Username))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *Penyuluh) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nama == "" {
			return errors.New("Required Nama")
		}
		if u.No_hp == "" {
			return errors.New("Required No_hp")
		}
		if u.Alamat == "" {
			return errors.New("Required Alamat")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	default:
		if u.Nama == "" {
			return errors.New("Required Nama")
		}
		if u.No_hp == "" {
			return errors.New("Required No_hp")
		}
		if u.Alamat == "" {
			return errors.New("Required Alamat")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		if u.User.Username == "" {
			return errors.New("Required Username")
		}
		if u.User.Password == "" {
			return errors.New("Required Password")
		}
		return nil
	}
}

func (u *Penyuluh) SavePenyuluh(db *gorm.DB) (*Penyuluh, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Penyuluh{}, err
	}
	return u, nil
}

func (u *Penyuluh) FindAllPenyuluhs(db *gorm.DB) (*[]Penyuluh, error) {
	var err error
	Penyuluhs := []Penyuluh{}
	err = db.Debug().Model(&Penyuluh{}).Limit(100).Preload("User").Preload("Jadwals.Laporan").Find(&Penyuluhs).Error
	if err != nil {
		return &[]Penyuluh{}, err
	}
	return &Penyuluhs, err
}

func (u *Penyuluh) FindPenyuluhByID(db *gorm.DB, uid uint32) (*Penyuluh, error) {
	var err error
	err = db.Debug().Model(Penyuluh{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Penyuluh{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Penyuluh{}, errors.New("Penyuluh Not Found")
	}
	return u, err
}

func (u *Penyuluh) UpdatePenyuluh(db *gorm.DB, uid uint32) (*Penyuluh, error) {

	db = db.Debug().Model(&Penyuluh{}).Where("id = ?", uid).Take(&Penyuluh{}).UpdateColumns(
		map[string]interface{}{
			"nama":       u.Nama,
			"no_hp":      u.No_hp,
			"alamat":     u.Alamat,
			"email":      u.Email,
			"foto":       u.Foto,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Penyuluh{}, db.Error
	}
	// This is the display the updated Penyuluh
	err := db.Debug().Model(&Penyuluh{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Penyuluh{}, err
	}
	return u, nil
}

func (u *Penyuluh) DeletePenyuluh(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Penyuluh{}).Where("id = ?", uid).Take(&Penyuluh{}).Delete(&Penyuluh{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
