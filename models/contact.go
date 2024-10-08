package models

import (
	"log"

	"github.com/jinzhu/gorm"
	u "github.com/pchchv/contacts/utils"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"` // user to whom the contact belongs
}

// Checks the mandatory parameters sent through the body of the http request.
// Returns a message and true if the requirement is met.
func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
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
