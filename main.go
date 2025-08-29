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

func showMainMenu() {
	fmt.Println()
	fmt.Println("[1] Store to ASEAN phonebook")
	fmt.Println("[2] Edit entry in ASEAN phonebook")
	fmt.Println("[3] Delete entry from ASEAN phonebook")
	fmt.Println("[4] View/Search ASEAN phonebook")
	fmt.Println("[5] Exit")
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
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
