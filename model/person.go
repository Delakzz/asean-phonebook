package model

import "fmt"

type Person struct {
	id          int64
	fName       string
	lName       string
	countryCode int16
	areaCode    int32
	occupation  string
	contactNum  string
	sex         string
}

func (p *Person) GetID() int64 {
	return p.id
}

func (p *Person) GetFullName() string {
	return p.fName + " " + p.lName
}

func (p *Person) GetPhoneNumber() string {
	return "+" + fmt.Sprintf("%d", p.countryCode) + "-" + fmt.Sprintf("%d", p.areaCode) + "-" + p.contactNum
}

func (p *Person) GetOccupation() string {
	return p.occupation
}

func (p *Person) GetSex() string {
	return p.sex
}

func (p *Person) CompareTo(other *Person) int {
	if p.lName < other.lName {
		return -1
	} else if p.lName > other.lName {
		return 1
	} else {
		if p.fName < other.fName {
			return -1
		} else if p.fName > other.fName {
			return 1
		} else {
			return 0
		}
	}
}
