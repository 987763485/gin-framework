/**
 * @Author: wuchunle<wuchunle@gsaxns.com>
 * @Version: 1.0.0
 * @Description:
 * @File:  user
 * @Time: 2020/9/17 10:05 上午
 */

package models

import (
	"github.com/987763485/gin-framework/application/lib"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type User struct {
	ID          string     `gorm:"type:char(19)" json:"id"`
	LoginName   string     `gorm:"type:varchar(60);not null;unique" json:"login_name"`
	Password    string     `gorm:"type:varchar(60);not null" json:"-"`
	Status      uint8      `gorm:"type:tinyint(1);not null;unsigned;default 1" json:"status"`
	Name        string     `gorm:"type:varchar(60);not null" json:"name"`
	OfficePhone string     `gorm:"not null" json:"office_phone"`
	MobilePhone string     `gorm:"not null" json:"mobile_phone"`
	Department  string     `gorm:"not null" json:"department"`
	Position    string     `gorm:"not null" json:"position"`
	Email       string     `gorm:"type:varchar(50);not null" json:"email"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	var ID uint64
	var err error
	if ID, err = lib.GenerateGUID(); err != nil {
		return err
	}
	if err = scope.SetColumn("ID", strconv.Itoa(int(ID))); err != nil {
		return err
	}
	return err
}
