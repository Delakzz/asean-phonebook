package model

import "fmt"

type Person struct {
	ID          int
	FName       string
	LName       string
	CountryCode int
	AreaCode    int
	Occupation  string
	ContactNum  string
	Sex         string
}

func NewPerson(id int, fName, lName string, countryCode, areaCode int, contactNum, occupation, sex string) *Person {
	return &Person{
		ID:          id,
		FName:       fName,
		LName:       lName,
		CountryCode: countryCode,
		AreaCode:    areaCode,
		ContactNum:  contactNum,
		Occupation:  occupation,
		Sex:         sex,
	}
}

func (p *Person) GetID() int {
	return p.ID
}

func (p *Person) GetFullName() string {
	return p.FName + " " + p.LName
}

func (p *Person) GetCountryCode() int {
	return p.CountryCode
}

func (p *Person) GetPhoneNumber() string {
	return fmt.Sprintf("+%d-%d-%s", p.CountryCode, p.AreaCode, p.ContactNum)
}

func (p *Person) GetOccupation() string {
	return p.Occupation
}

func (p *Person) GetSex() string {
	return p.Sex
}

func (p *Person) CompareTo(other *Person) int {
	if p.LName < other.LName {
		return -1
	} else if p.LName > other.LName {
		return 1
	} else {
		if p.FName < other.FName {
			return -1
		} else if p.FName > other.FName {
			return 1
		} else {
			return 0
		}
	}
}
