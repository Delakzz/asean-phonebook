package repository

import (
	"asean-phonebook/model"
	"errors"
)

type Phonebook struct {
	contacts []*model.Person
}

func NewPhonebook() *Phonebook {
	return &Phonebook{contacts: []*model.Person{}}
}

func (pb *Phonebook) GetContactAtIndex(index int) (*model.Person, error) {
	if index < 0 || index >= len(pb.contacts) {
		return nil, errors.New("index out of range")
	}
	return pb.contacts[index], nil
}

func (pb *Phonebook) GetContact(id int) (*model.Person, error) {
	for _, contact := range pb.contacts {
		if contact.GetID() == id {
			return contact, nil
		}
	}
	return nil, errors.New("contact not found")
}

func (pb *Phonebook) IsEmpty() bool {
	return len(pb.contacts) == 0
}

func (pb *Phonebook) DeleteContact(id int) error {
	for i, contact := range pb.contacts {
		if contact.GetID() == id {
			pb.contacts = append(pb.contacts[:i], pb.contacts[i+1:]...)
			return nil
		}
	}
	return errors.New("contact not found")
}

func (pb *Phonebook) insert(p *model.Person) {
	if len(pb.contacts) == 0 {
		pb.contacts = append(pb.contacts, p)
		return
	}
	index := pb.findIndexInsertion(p)
	pb.adjustPhonebook(index, "insert")
	pb.contacts[index] = p
}

func (pb *Phonebook) findIndexInsertion(p *model.Person) int {
	high := len(pb.contacts) - 1
	low := 0

	for low <= high {
		mid := (high + low) / 2
		if pb.contacts[mid].CompareTo(p) < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func (pb *Phonebook) adjustPhonebook(index int, direction string) {
	switch direction {
	case "insert":
		// make room by appending nil then shifting
		pb.contacts = append(pb.contacts, nil)
		copy(pb.contacts[index+1:], pb.contacts[index:])
	case "delete":
		// shift left to remove
		copy(pb.contacts[index:], pb.contacts[index+1:])
		pb.contacts = pb.contacts[:len(pb.contacts)-1]
	}
}

func (pb *Phonebook) PrintContactsFromCountryCodes(countryCodes []int) []*model.Person {
	var filtered []*model.Person
	lookup := make(map[int]struct{})
	for _, code := range countryCodes {
		lookup[code] = struct{}{}
	}
	for _, c := range pb.contacts {
		if _, ok := lookup[c.GetCountryCode()]; ok {
			filtered = append(filtered, c)
		}
	}
	return filtered
}
