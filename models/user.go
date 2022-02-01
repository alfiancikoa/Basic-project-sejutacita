package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"primarykey;AUTO_INCREMENT"`
	Username  string `gorm:"type:varchar(255);not null" json:"username" form:"username"`
	Email     string `gorm:"type:varchar(100);unique;not null" json:"email" form:"email"`
	Password  string `gorm:"type:varchar(255);not null" json:"password" form:"password"`
	Role      string `gorm:"type:varchar(20)" json:"role" form:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type LoginUser struct {
	Username string `gorm:"type:varchar(255);not null" json:"username" form:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password" form:"password"`
}

type GormUserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *GormUserModel {
	return &GormUserModel{db: db}
}

// Interface User

type UserModel interface {
	Insert(User) (User, error)
	CheckDatabase(string, string) (int64, error)
	FindUserByUsername(string) (*User, error)
	FindUserBy(string, interface{}) (*User, error)
	Login(LoginUser) (*User, error)
	Edit(User, int) (*User, error)
	Delete(int) (int64, error)
	FindUsers() ([]User, error)
}

// Fungsi untuk menambahkan data user ke dalam database
func (m *GormUserModel) Insert(user User) (User, error) {
	if err := m.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Fungsi untuk mengambil dan mencari data organizer by email di database
func (m *GormUserModel) FindUserByUsername(username string) (*User, error) {
	user := User{}
	tx := m.db.Where("username=?", username).Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}

// Fungsi untuk mengambil dan mencari data organizer by request di database
func (m *GormUserModel) FindUserBy(request string, data interface{}) (*User, error) {
	user := User{}
	tx := m.db.Where(request+"=?", data).Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}

// Fungsi untuk mengambil dan mencari data organizer by request di database
func (m *GormUserModel) FindUsers() ([]User, error) {
	user := []User{}
	tx := m.db.Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return user, nil
	}
	return nil, nil
}

// Fungsi untuk Login ke dalam sistem
func (m *GormUserModel) Login(dataLogin LoginUser) (*User, error) {
	userData, err := m.FindUserByUsername(dataLogin.Username)
	if userData == nil || err != nil {
		return nil, err
	}
	check := CheckPasswordHash(dataLogin.Password, userData.Password)
	if !check {
		return nil, nil
	}
	return userData, nil
}

// Fungsi untuk mengecek ketersediaan data di database
func (m *GormUserModel) CheckDatabase(coloumn string, data string) (int64, error) {
	organizer := User{}
	tx := m.db.Where(coloumn+"=?", data).Find(&organizer)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return tx.RowsAffected, nil
}

// Fungsi untuk mengedit data user di database
func (m *GormUserModel) Edit(reqUser User, id int) (*User, error) {
	user := User{}
	tx := m.db.Find(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		return nil, nil
	}
	if err := m.db.Model(&User{}).Where("id=?", id).Updates(reqUser).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Fungsi untuk menghapus user
func (m *GormUserModel) Delete(user_id int) (int64, error) {
	tx := m.db.Delete(&User{}, user_id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return tx.RowsAffected, nil
}

// Fungsi untuk enkripsi password organizer
func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Fungsi untuk compare password organizer dengan enkripsi password organizer
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
