package chat

import (
	"github.com/shobky/giggs/config/database"
	"github.com/shobky/giggs/model"
)

func InsertNewContact(contact *model.Contact) error {
	err := database.DB.Create(contact).Error
	if err != nil {
		return err
	}
	return nil
}

func GetContacgByID(id uint) (*model.Contact, bool) {
	var contact *model.Contact

	database.DB.First(&contact, "id = ?", id)

	if contact.ID == 0 {
		return nil, false
	}
	return contact, true
}

func GetAllContacts(userID uint) ([]*model.Contact, error) {
	var contacts []*model.Contact
	if result := database.DB.Where("user_id = ?", userID).Find(&contacts); result.Error != nil {
		return nil, result.Error
	}
	return contacts, nil
}
