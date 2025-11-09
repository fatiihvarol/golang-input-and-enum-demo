package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type GenderType int

const (
	Male GenderType = iota
	Female
	Unspecified
)

func (g GenderType) String() string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unspecified"
	}
}

type User struct {
	FirstName   string
	LastName    string
	Email       string
	Username    string
	Gender      GenderType
	DateOfBirth time.Time
	Age         int
}

func boolToEnum(isMale bool) GenderType {
	if isMale {
		return Male
	}
	return Female
}

func calculateAge(dob time.Time) int {
	today := time.Now()
	age := today.Year() - dob.Year()

	if dob.YearDay() > today.YearDay() {
		age--
	}
	return age
}

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func validateUsername(username string) bool {
	length := len(username)
	return length >= 3 && length <= 15
}

func validateEmail(email string) bool {
	if len(email) < 6 || len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}

func validateRequired(input string) bool {
	return input != ""
}

func getInput() User {
	var u User

	fmt.Println("--- User Information Input ---")

	for {
		u.FirstName = readInput("First Name: ")
		if validateRequired(u.FirstName) {
			break
		}
		fmt.Println("Error: First Name cannot be empty.")
	}

	for {
		u.LastName = readInput("Last Name: ")
		if validateRequired(u.LastName) {
			break
		}
		fmt.Println("Error: Last Name cannot be empty.")
	}

	for {
		u.Username = readInput("Username (3-15 Characters): ")
		if !validateRequired(u.Username) {
			fmt.Println("Error: Username cannot be empty.")
			continue
		}
		if validateUsername(u.Username) {
			break
		}
		fmt.Println("Error: Username must be between 3 and 15 characters long.")
	}

	for {
		u.Email = readInput("Email Address: ")
		if !validateRequired(u.Email) {
			fmt.Println("Error: Email cannot be empty.")
			continue
		}
		if validateEmail(u.Email) {
			break
		}
		fmt.Println("Error: Please enter a valid email address.")
	}

	var genderInput string
	fmt.Print("Gender (m/f): ")
	fmt.Scanln(&genderInput)
	genderInput = strings.ToLower(strings.TrimSpace(genderInput))
	isMale := (genderInput == "m")
	u.Gender = boolToEnum(isMale)

	var dobStr string
	for {
		dobStr = readInput("Date of Birth (YYYY-MM-DD): ")
		t, err := time.Parse("2006-01-02", dobStr)
		if err == nil {
			u.DateOfBirth = t
			u.Age = calculateAge(u.DateOfBirth)
			break
		}
		fmt.Println("Error: Invalid date format. Please use YYYY-MM-DD.")
	}

	return u
}

func displayInfo(u User) {
	fmt.Println("\n--- User Information Report ---")
	fmt.Println("---------------------------------")
	fmt.Printf("1. Full Name:          %s %s\n", u.FirstName, u.LastName)
	fmt.Printf("2. Username:           %s\n", u.Username)
	fmt.Printf("3. E-mail:             %s\n", u.Email)
	fmt.Printf("4. Gender (Enum):      %s\n", u.Gender)
	fmt.Printf("5. Date of Birth:      %s\n", u.DateOfBirth.Format("January 02, 2006"))
	fmt.Printf("6. Calculated Age:     %d\n", u.Age)
	fmt.Printf("7. Email Length:       %d characters\n", len(u.Email))
	fmt.Println("---------------------------------")
}

func main() {
	userData := getInput()
	displayInfo(userData)
}
