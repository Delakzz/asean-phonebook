package main

import (
	"asean-phonebook/model"
	"asean-phonebook/repository"
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func storeEntry(pb *repository.Phonebook) *model.Person {
	scanner := bufio.NewScanner(os.Stdin)

	readString := func(prompt string) string {
		fmt.Print(prompt)
		scanner.Scan()
		return scanner.Text()
	}

	readInt := func(prompt string) int {
		for {
			fmt.Print(prompt)
			scanner.Scan()
			input := scanner.Text()
			value, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid integer.")
				continue
			}
			return value
		}
	}

	studentNumber := readInt("\nEnter student number: ")
	surname := readString("Enter surname: ")
	firstName := readString("Enter first name: ")
	occupation := readString("Enter occupation: ")
	gender := readString("Enter gender (M for male, F for female): ")
	countryCode := readInt("Enter country code: ")
	areaCode := readInt("Enter area code: ")
	phoneNumber := readString("Enter phone number: ")

	person := model.NewPerson(studentNumber, firstName, surname, countryCode, areaCode, phoneNumber, occupation, gender)
	pb.Insert(person)

	fmt.Println("\nContact stored successfully!")
	return person
}

func deleteEntry(pb *repository.Phonebook) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter student number to delete: ")
		scanner.Scan()
		input := scanner.Text()
		studentNumber, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		fmt.Print("Are you sure you want to delete it (y/n)? ")
		scanner.Scan()
		confirm := strings.ToLower(scanner.Text())
		if confirm != "y" {
			fmt.Println("Deletion did not proceed")
			return
		}
		pb.DeleteContact(studentNumber)
		fmt.Println("Deletion successful")
		return
	}
}

func showPersonInformation(pb *repository.Phonebook, id int) error {
	person, err := pb.GetContact(id)
	if err != nil {
		return errors.New("contact not found")
	}

	fmt.Printf("\nHere is the existing information about %d:\n", person.GetID())
	fmt.Printf("%s\n", person.GetPersonDetails())
	return nil
}

func editEntry(pb *repository.Phonebook, person *model.Person) {
	scanner := bufio.NewScanner(os.Stdin)

	readString := func(prompt string) string {
		fmt.Print(prompt)
		scanner.Scan()
		return scanner.Text()
	}

	readInt := func(prompt string) int {
		for {
			fmt.Print(prompt)
			scanner.Scan()
			input := scanner.Text()
			value, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid integer.")
				continue
			}
			return value
		}
	}

	updatedContact := person

	for {
		showPersonInformation(pb, person.ID)
		showEditMenu()
		choice := readString("Enter choice: ")
		switch choice {
		case "1":
			newID := readInt("Enter new student number: ")
			updatedContact.ID = newID
		case "2":
			newFName := readString("Enter new first name: ")
			updatedContact.FName = newFName
		case "3":
			newLName := readString("Enter new surname: ")
			updatedContact.LName = newLName
		case "4":
			newOccupation := readString("Enter new occupation: ")
			updatedContact.Occupation = newOccupation
		case "5":
			newCountryCode := readInt("Enter new country code: ")
			updatedContact.CountryCode = newCountryCode
		case "6":
			newAreaCode := readInt("Enter new area code: ")
			updatedContact.AreaCode = newAreaCode
		case "7":
			newPhoneNumber := readString("Enter new phone number: ")
			updatedContact.ContactNum = newPhoneNumber
		case "8":
			newGender := readString("Enter new gender: ")
			updatedContact.Sex = newGender
		case "9":
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
		pb.UpdateContact(person.ID, updatedContact)
	}
}

func searchCountry(pb *repository.Phonebook) {
	scanner := bufio.NewScanner(os.Stdin)

	readInt := func(prompt string) int {
		for {
			fmt.Print(prompt)
			scanner.Scan()
			input := scanner.Text()
			value, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid integer.")
				continue
			}
			return value
		}
	}

	countryCodes := map[string]int{
		"Malaysia":    60,
		"Indonesia":   62,
		"Philippines": 63,
		"Vietnam":     84,
		"Thailand":    66,
		"Myanmar":     95,
		"Cambodia":    855,
		"Laos":        856,
		"Singapore":   65,
		"Brunei":      673,
		"Timor Leste": 670,
	}

	fmt.Println("From which country:")
	fmt.Println("[1] Myanmar [2] Cambodia [3] Thailand [4] Vietnam [5] Malaysia")
	fmt.Println("[6] Philippines [7] Indonesia [8] Timor Leste [9] Laos")
	fmt.Println("[10] Brunei [11] Singapore [12] All [0] No more")

	selectedCountryCodes := []int{}
	count := 1
	for {
		choice := readInt(fmt.Sprintf("\nEnter choice %d: ", count))
		if choice == 0 {
			break
		}
		var selectedCode int
		switch choice {
		case 1:
			selectedCode = countryCodes["Myanmar"]
		case 2:
			selectedCode = countryCodes["Cambodia"]
		case 3:
			selectedCode = countryCodes["Thailand"]
		case 4:
			selectedCode = countryCodes["Vietnam"]
		case 5:
			selectedCode = countryCodes["Malaysia"]
		case 6:
			selectedCode = countryCodes["Philippines"]
		case 7:
			selectedCode = countryCodes["Indonesia"]
		case 8:
			selectedCode = countryCodes["Timor Leste"]
		case 9:
			selectedCode = countryCodes["Laos"]
		case 10:
			selectedCode = countryCodes["Brunei"]
		case 11:
			selectedCode = countryCodes["Singapore"]
		case 12:
			for _, code := range countryCodes {
				fmt.Printf("\n%s\n", pb.PrintContactsFromCountryCodes([]int{code}))
			}
			return
		default:
			fmt.Println("Invalid choice, please try again.")
			continue
		}
		if slices.Contains(selectedCountryCodes, selectedCode) {
			fmt.Println("Country already selected, please choose another.")
			continue
		}
		selectedCountryCodes = append(selectedCountryCodes, selectedCode)
		count++
	}

	// for displaying selected countries
	var matches []string
	for name, code := range countryCodes {
		if slices.Contains(selectedCountryCodes, code) {
			matches = append(matches, name)
		}
	}

	var selectedCountries string
	switch len(matches) {
	case 0:
		selectedCountries = ""
	case 1:
		selectedCountries = matches[0]
	case 2:
		selectedCountries = matches[0] + " and " + matches[1]
	default:
		selectedCountries = strings.Join(matches[:len(matches)-1], ", ") +
			", and " + matches[len(matches)-1]
	}

	fmt.Println()
	fmt.Printf("Here are the students from the %s:", selectedCountries)
	result := pb.PrintContactsFromCountryCodes(selectedCountryCodes)
	fmt.Println()
	fmt.Printf("\n%s\n", strings.Join(result, "\n\n"))
}

func searchSurname(pb *repository.Phonebook) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter surname: ")
	scanner.Scan()
	surname := scanner.Text()
	results := pb.GetSurnames(surname)
	fmt.Println()
	if len(results) == 0 {
		fmt.Printf("No contacts found with the surname '%s'.\n", surname)
	} else {
		fmt.Printf("Contacts with the surname '%s':\n", surname)
		fmt.Printf("\n%s\n", strings.Join(results, "\n\n"))
	}
}

func showMainMenu() {
	fmt.Println()
	fmt.Println("[1] Store to ASEAN phonebook")
	fmt.Println("[2] Edit entry in ASEAN phonebook")
	fmt.Println("[3] Delete entry from ASEAN phonebook")
	fmt.Println("[4] View/Search ASEAN phonebook")
	fmt.Println("[5] Exit")
	fmt.Println()
}

func showEditMenu() {
	fmt.Println()
	fmt.Println("Which of the following information do you wish to change?")
	fmt.Println("[1] Student number [2] First name [3] Surname [4] Occupation")
	fmt.Println("[5] Country code [6] Area code [7] Phone number [8] Gender")
	fmt.Println("[9] None - Go back to main menu")
	fmt.Println()
}

func showViewSearchMenu() {
	fmt.Println()
	fmt.Println("[1] Search by countries")
	fmt.Println("[2] Search by surname")
	fmt.Println("[3] Go back to main menu")
	fmt.Println()
}

func main() {
	pb, err := repository.NewPhonebook("db/phonebook.json")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMainMenu()
		fmt.Print("Enter choice: ")
		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			for {
				storeEntry(pb)
				fmt.Print("\nDo you want to enter another contact? (y/n): ")
				scanner.Scan()
				if again := strings.ToLower(scanner.Text()); again != "y" {
					break
				}
			}
		case "2":
			if pb.IsEmpty() {
				fmt.Println("\nPhonebook is empty. Please add contacts first.")
				break
			}
			var student *model.Person
			for {
				fmt.Print("\nEnter student number to edit: ")
				scanner.Scan()
				input := scanner.Text()
				studentNumber, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Invalid input. Please enter a valid integer.")
					continue
				}

				student, err = pb.GetContact(studentNumber)
				if err != nil {
					fmt.Println("Contact not found.")
					continue
				}
				break
			}
			editEntry(pb, student)
		case "3":
			if pb.IsEmpty() {
				fmt.Println("\nPhonebook is empty. Please add contacts first.")
				continue
			}
			deleteEntry(pb)
		case "4":
		subMenu:
			for {
				showViewSearchMenu()
				fmt.Print("Enter choice: ")
				scanner.Scan()
				subChoice := scanner.Text()

				switch subChoice {
				case "1":
					if pb.IsEmpty() {
						fmt.Println("\nPhonebook is empty. Please add contacts first.")
						continue
					}
					searchCountry(pb)

				case "2":
					if pb.IsEmpty() {
						fmt.Println("\nPhonebook is empty. Please add contacts first.")
						continue
					}
					searchSurname(pb)

				case "3":
					fmt.Println("Returning to main menu...")
					break subMenu

				default:
					fmt.Println("Invalid choice, please try again.")
				}
			}
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
