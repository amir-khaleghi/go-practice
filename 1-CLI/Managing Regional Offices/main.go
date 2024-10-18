/* // Step 1: Initial Setup And Imports ----------------- */
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* // Step 2: Define Structs ---------------------------- */
// Office represents an office of the company
type Office struct {
	ID                int
	Name              string
	Address           string
	PhoneNumber       string
	MembershipDate    string
	NumberOfEmployees int
}

/* // Step 3: Define Global Variables ------------------- */
// File for data persistence
const dataFileName = "offices.csv"

/* ■■■■■■■■■■■■■■■■■■■■■■■■ Main ■■■■■■■■■■■■■■■■■■■■■■■■ */
func main() {
	// Step 4: Define Flags
	command := flag.String("command", "", "Command to execute (list, get, create, edit, status)")
	region := flag.String("region", "", "Region (e.g., Tehran, Isfahan, Shiraz)")
	flag.Parse()

	// Step 5: Read Existing Data from File
	offices, err := readOfficesFromFile(dataFileName)
	if err != nil {
		fmt.Println("Error reading data from file:", err)
		return
	}

	// Step 6: Execute Commands Based on Flags
	switch *command {
	case "list":
		listOffices(offices, *region)
	case "get":
		getOffice(offices)
	case "create":
		offices = createOffice(offices)
	case "edit":
		offices = editOffice(offices)
	case "status":
		statusOffices(offices, *region)
	default:
		fmt.Println("Invalid command. Available commands: list, get, create, edit, status")
	}

	// Step 7: Save Updated Data Back to File
	err = writeOfficesToFile(dataFileName, offices)
	if err != nil {
		fmt.Println("Error writing data to file:", err)
	}
}

/* ■■■■■■■■■■■■■■■■■■■■ Os Functions ■■■■■■■■■■■■■■■■■■■■ */

/* // Function To Read Offices From File ---------------- */
func readOfficesFromFile(filename string) ([]Office, error) {
	// Open file
	file, err := os.Open(filename)

	// Handle file not found
	if err != nil {
		if os.IsNotExist(err) {
			return []Office{}, nil // Return an empty slice if file doesn't exist
		}
		return nil, err
	}

	// Defer File Close to prevent resource leaks
	// When a defer statement is encountered, the function call is pushed onto a stack of deferred function calls, and it will only be executed when the surrounding function ends (either by returning or due to a runtime error).

	defer file.Close()

	// Initialize Office Slice to store new offices
	offices := []Office{}

	// Read CSV File
	reader := csv.NewReader(file)

	// loop through each line
	for {
		record, err := reader.Read()

		if err != nil {
			break
		}
		if len(record) == 0 {
			continue
		}

		// Convert id and numEmployees to int
		id, _ := strconv.Atoi(record[0])
		numEmployees, _ := strconv.Atoi(record[5])

		// Convert CSV DATA to Office Struct
		office := Office{
			ID:                id,
			Name:              record[1],
			Address:           record[2],
			PhoneNumber:       record[3],
			MembershipDate:    record[4],
			NumberOfEmployees: numEmployees,
		}

		offices = append(offices, office)
	}

	return offices, nil
}

/* // Function To Write Offices To File ----------------- */
func writeOfficesToFile(filename string, offices []Office) error {

	// create a new file if the file already exists, it will truncate its content and overwritig it
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV Writer:
	// csv.NewWriter(file): This function initializes a csv.Writer to write comma-separated values to the file. The writer will handle all data formatting required for CSV.

	writer := csv.NewWriter(file)

	// Defer CSV Writer Flush
	//This ensures that any buffered data is flushed to the file. Without flushing, data that is buffered internally may not be completely written to the file when the function exits.

	defer writer.Flush()

	//Iterate Over the Slice of Offices:

	for _, office := range offices {

		//Convert Struct to String Slice:

		record := []string{
			strconv.Itoa(office.ID),
			office.Name,
			office.Address,
			office.PhoneNumber,
			office.MembershipDate,
			strconv.Itoa(office.NumberOfEmployees),
		}
		// The CSV writer requires each row as a slice of strings ([]string).

		writer.Write(record)
	}
	return nil
}

/* ■■■■■■■■■■■■■■■■  Command Functions ■■■■■■■■■■■■■■■■ */
/*  List Offices In A Specific Region ----------------- */
func listOffices(offices []Office, region string) {
	fmt.Println("Offices in region:", region)
	for _, office := range offices {
		if strings.EqualFold(office.Address, region) {
			fmt.Printf("ID: %d, Name: %s, Address: %s, Phone: %s, Employees: %d\n",
				office.ID, office.Name, office.Address, office.PhoneNumber, office.NumberOfEmployees)
		}
	}
}

/* Get Complete Information About A Specific Office ----- */
func getOffice(offices []Office) {
	fmt.Print("Enter office ID: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	idStr := scanner.Text()
	id, _ := strconv.Atoi(idStr)

	for _, office := range offices {
		if office.ID == id {
			fmt.Printf("ID: %d, Name: %s, Address: %s, Phone: %s, Membership Date: %s, Employees: %d\n",
				office.ID, office.Name, office.Address, office.PhoneNumber, office.MembershipDate, office.NumberOfEmployees)
			return
		}
	}
	fmt.Println("Office not found")
}

/* Create A New Office ---------------------------------- */
func createOffice(offices []Office) []Office {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter office name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter office address: ")
	scanner.Scan()
	address := scanner.Text()

	fmt.Print("Enter phone number: ")
	scanner.Scan()
	phoneNumber := scanner.Text()

	fmt.Print("Enter membership date: ")
	scanner.Scan()
	membershipDate := scanner.Text()

	fmt.Print("Enter number of employees: ")
	scanner.Scan()
	numEmployeesStr := scanner.Text()
	numEmployees, _ := strconv.Atoi(numEmployeesStr)

	newID := len(offices) + 1
	office := Office{
		ID:                newID,
		Name:              name,
		Address:           address,
		PhoneNumber:       phoneNumber,
		MembershipDate:    membershipDate,
		NumberOfEmployees: numEmployees,
	}
	offices = append(offices, office)

	fmt.Println("Office created successfully!")
	return offices
}

/* Edit An Existing Office ------------------------------ */
func editOffice(offices []Office) []Office {
	fmt.Print("Enter office ID to edit: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	idStr := scanner.Text()
	id, _ := strconv.Atoi(idStr)

	for i, office := range offices {
		if office.ID == id {
			fmt.Printf("Editing Office: %s\n", office.Name)

			fmt.Print("Enter new office name (leave blank to keep current): ")
			scanner.Scan()
			name := scanner.Text()
			if name != "" {
				offices[i].Name = name
			}

			fmt.Print("Enter new office address (leave blank to keep current): ")
			scanner.Scan()
			address := scanner.Text()
			if address != "" {
				offices[i].Address = address
			}

			fmt.Print("Enter new phone number (leave blank to keep current): ")
			scanner.Scan()
			phone := scanner.Text()
			if phone != "" {
				offices[i].PhoneNumber = phone
			}

			fmt.Print("Enter new number of employees (leave blank to keep current): ")
			scanner.Scan()
			numEmployeesStr := scanner.Text()
			if numEmployeesStr != "" {
				numEmployees, _ := strconv.Atoi(numEmployeesStr)
				offices[i].NumberOfEmployees = numEmployees
			}

			fmt.Println("Office updated successfully!")
			return offices
		}
	}
	fmt.Println("Office not found")
	return offices
}

/* Display The Status Of Offices In A Specific Region --- */
func statusOffices(offices []Office, region string) {
	officeCount := 0
	totalEmployees := 0
	for _, office := range offices {
		if strings.EqualFold(office.Address, region) {
			officeCount++
			totalEmployees += office.NumberOfEmployees
		}
	}
	fmt.Printf("Number of offices in %s: %d\nTotal number of employees: %d\n", region, officeCount, totalEmployees)
}
