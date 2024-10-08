package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"` // user to whom the contact belongs
}

func GetContact(id uint) *Contact {
	contact := &Contact{}
	if err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error; err != nil {
		return nil
	}

	return contact
}

func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	if err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error; err != nil {
		log.Println(err)
		return nil
	}

	return contacts
}
