package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

type Admin struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nama      string    `json:"Nama"`
	No_hp     string    `json:"No_hp"`
	Alamat    string    `json:"Alamat"`
	Email     string    `json:"Email"`
	Foto      string    `json:"Foto"`
	User      User      `gorm:"polymorphic:Level;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Admin) Prepare() {
	u.ID = 0
	u.Nama = html.EscapeString(strings.TrimSpace(u.Nama))
	u.No_hp = html.EscapeString(strings.TrimSpace(u.No_hp))
	u.Alamat = html.EscapeString(strings.TrimSpace(u.Alamat))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Foto = html.EscapeString(strings.TrimSpace(u.Foto))
	u.User.Username = html.EscapeString(strings.TrimSpace(u.User.Username))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *Admin) Validate(action string) error {
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

func (u *Admin) SaveAdmin(db *gorm.DB) (*Admin, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Admin{}, err
	}
	return u, nil
}

func (u *Admin) FindAllAdmins(db *gorm.DB) (*[]Admin, error) {
	var err error
	Admins := []Admin{}
	err = db.Debug().Model(&Admin{}).Limit(100).Preload("User").Find(&Admins).Error
	if err != nil {
		return &[]Admin{}, err
	}
	return &Admins, err
}

func (u *Admin) FindAdminByID(db *gorm.DB, uid uint32) (*Admin, error) {
	var err error
	err = db.Debug().Model(Admin{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Admin{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Admin{}, errors.New("Admin Not Found")
	}
	return u, err
}

func (u *Admin) UpdateAdmin(db *gorm.DB, uid uint32) (*Admin, error) {

	db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).UpdateColumns(
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
		return &Admin{}, db.Error
	}
	// This is the display the updated Admin
	err := db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Admin{}, err
	}
	return u, nil
}

func (u *Admin) DeleteAdmin(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).Delete(&Admin{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
