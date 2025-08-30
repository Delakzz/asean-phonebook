package repository

import (
	"asean-phonebook/model"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"sync"
)

type Phonebook struct {
	filePath string
	mutex    sync.RWMutex
	Contacts []*model.Person
}

func NewPhonebook(filePath string) (*Phonebook, error) {
	pb := &Phonebook{filePath: filePath}
	if err := pb.load(); err != nil {
		return nil, err
	}
	return pb, nil
}

func (pb *Phonebook) load() error {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()

	if _, err := os.Stat(pb.filePath); os.IsNotExist(err) {
		pb.Contacts = []*model.Person{}
		return nil
	}

	data, err := os.ReadFile(pb.filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &pb.Contacts)
}

func (pb *Phonebook) save() error {
	data, err := json.MarshalIndent(pb.Contacts, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(pb.filePath, data, 0644)
}

func (pb *Phonebook) GetContactAtIndex(index int) (*model.Person, error) {
	pb.mutex.RLock()
	defer pb.mutex.RUnlock()
	if index < 0 || index >= len(pb.Contacts) {
		return nil, errors.New("index out of range")
	}
	return pb.Contacts[index], nil
}

func (pb *Phonebook) GetContact(id int) (*model.Person, error) {
	pb.mutex.RLock()
	defer pb.mutex.RUnlock()
	for _, contact := range pb.Contacts {
		if contact.GetID() == id {
			return contact, nil
		}
	}
	return nil, errors.New("contact not found")
}

func (pb *Phonebook) IsEmpty() bool {
	pb.mutex.RLock()
	defer pb.mutex.RUnlock()
	return len(pb.Contacts) == 0
}

func (pb *Phonebook) UpdateContact(id int, updated *model.Person) error {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()
	for i, contact := range pb.Contacts {
		if contact.GetID() == id {
			pb.Contacts[i] = updated
			return pb.save()
		}
	}
	return errors.New("contact not found")
}

func (pb *Phonebook) DeleteContact(id int) error {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()
	for i, contact := range pb.Contacts {
		if contact.GetID() == id {
			pb.adjustPhonebook(i, "delete")
			return pb.save()
		}
	}
	return errors.New("contact not found")
}

func (pb *Phonebook) Insert(p *model.Person) error {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()
	if len(pb.Contacts) == 0 {
		pb.Contacts = append(pb.Contacts, p)
		return pb.save()
	}
	index := pb.findIndexInsertion(p)
	pb.adjustPhonebook(index, "insert")
	pb.Contacts[index] = p
	return pb.save()
}

func (pb *Phonebook) findIndexInsertion(p *model.Person) int {
	high := len(pb.Contacts) - 1
	low := 0

	for low <= high {
		mid := (high + low) / 2
		if pb.Contacts[mid].CompareTo(p) < 0 {
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
		// to right
		pb.Contacts = append(pb.Contacts, nil)
		copy(pb.Contacts[index+1:], pb.Contacts[index:])
	case "delete":
		// to left
		copy(pb.Contacts[index:], pb.Contacts[index+1:])
		pb.Contacts = pb.Contacts[:len(pb.Contacts)-1]
	}
}

func (pb *Phonebook) PrintContactsFromCountryCodes(selectedCountryCodes []int) []string {
	pb.mutex.RLock()
	defer pb.mutex.RUnlock()
	var filtered []string
	lookup := make(map[int]struct{})
	for _, code := range selectedCountryCodes {
		lookup[code] = struct{}{}
	}
	for _, c := range pb.Contacts {
		if _, ok := lookup[c.GetCountryCode()]; ok {
			filtered = append(filtered, c.GetPersonDetails())
		}
	}
	return filtered
}

func (pb *Phonebook) GetSurnames(selectedSurname string) []string {
	pb.mutex.RLock()
	defer pb.mutex.RUnlock()
	var surnames []string
	for _, c := range pb.Contacts {
		if strings.EqualFold(c.GetLastName(), selectedSurname) {
			surnames = append(surnames, c.GetPersonDetails())
		}
	}
	return surnames
}
