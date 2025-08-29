package main

import (
	"asean-phonebook/model"
	"asean-phonebook/repository"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func store(pb *repository.Phonebook) *model.Person {
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

func searchCountry(pb *repository.Phonebook) []int {
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
		choice := readInt(fmt.Sprintf("Enter choice %d: ", count))
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
			return nil
		default:
			fmt.Println("Invalid choice, please try again.")
			continue
		}
		selectedCountryCodes = append(selectedCountryCodes, selectedCode)
		count++
	}

	fmt.Printf("\n%s\n", pb.PrintContactsFromCountryCodes(selectedCountryCodes))

	return nil
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

func showViewSearchMenu() {
	fmt.Println()
	fmt.Println("[1] Search by countries")
	fmt.Println("[2] Search by surname")
	fmt.Println("[3] Go back to main menu")
	fmt.Println()
}

func main() {
	pb := &repository.Phonebook{Contacts: []*model.Person{}}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMainMenu()
		fmt.Print("Enter choice: ")
		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			for {
				store(pb)
				fmt.Print("\nDo you want to enter another contact? (y/n): ")
				scanner.Scan()
				if again := strings.ToLower(scanner.Text()); again != "y" {
					break
				}
			}
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
					// searchSurname(pb)
					searchCountry(pb)

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
