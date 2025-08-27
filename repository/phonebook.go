package repository

import (
	"asean-phonebook/model"
	"errors"
)

type Phonebook struct {
	contacts []*model.Person
	size     int64
}

func (pb *Phonebook) GetContactAtIndex(index int64) (*model.Person, error) {
	if index < 0 || index >= pb.size {
		return nil, errors.New("index out of range")
	}
	return pb.contacts[index], nil
}

func (pb *Phonebook) GetContact(id int64) (*model.Person, error) {
	for _, contact := range pb.contacts {
		if contact.GetID() == id {
			return contact, nil
		}
	}
	return nil, errors.New("contact not found")
}

func (pb *Phonebook) IsEmpty() bool {
	return pb.size == 0
}

func (pb *Phonebook) incrSize() {
	pb.size++
}

func (pb *Phonebook) decrSize() {
	pb.size--
}

func (pb *Phonebook) insert(p *model.Person) {
	// Insert in sorted order
}

func (pb *Phonebook) DeleteContact(id int64) error {
	for i, contact := range pb.contacts {
		if contact.GetID() == id {
			pb.contacts = append(pb.contacts[:i], pb.contacts[i+1:]...)
			pb.decrSize()
			return nil
		}
	}
	return errors.New("contact not found")
}

func (pb *Phonebook) findIndexInsertion(p *model.Person) int64 {
	// Binary search to find the correct index for insertion
	return 0
}

func (pb *Phonebook) adjustPhonebook(start, end int, direction string) {
	// Adjust the phonebook slice based on insertion or deletion
}

func (pb *Phonebook) PrintContactsFromCountryCodes(countryCodes []int16) []*model.Person {
	// Filter contacts by country codes and return the list
	return nil
}
