package googlesheets

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/VladdPp/googlesheetscli/util"
	"google.golang.org/api/sheets/v4"
)

// HandleLoadSheet displays a list of available sheets in the Google Sheets table
// and allows the user to select a sheet to load data from.
func HandleLoadSheet(srv *sheets.Service, spreadsheetID string) {
	fmt.Println("Select the sheet to load data from:")

	availableSheets, err := getSheetNames(srv, spreadsheetID)
	if err != nil {
		log.Fatalf("Unable to retrieve sheet names: %v", err)
	}

	for i, sheet := range availableSheets {
		fmt.Printf("%d. %s\n", i+1, sheet)
	}

	scanner := bufio.NewScanner(os.Stdin)
	selectedSheetIndex, err := getUserSheetChoice(scanner, availableSheets)
	if err != nil {
		fmt.Println(err)
		return
	}

	selectedSheet := availableSheets[selectedSheetIndex-1]
	fmt.Printf("Loading data from sheet '%s'...\n", selectedSheet)
	err = loadSheetData(srv, spreadsheetID, selectedSheet)
	if err != nil {
		log.Printf("Error loading sheet: %v", err)
	} else {
		fmt.Printf("Sheet '%s' loaded successfully.\n", selectedSheet)
	}
}

// HandleLoadRangeSheet allows the user to input the sheet name and cell range
// to load data from a specific range in the Google Sheets.
func HandleLoadRangeSheet(srv *sheets.Service, spreadsheetID string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter sheet name: ")
	scanner.Scan()
	sheetName := scanner.Text()

	fmt.Print("Enter sheet range (e.g., A1:D5): ")
	scanner.Scan()
	sheetRange := sheetName + "!" + scanner.Text()

	err := loadSheetData(srv, spreadsheetID, sheetRange)
	if err != nil {
		log.Println(err)
	}
}

// HandleAddSheetData allows the user to input sheet name, column letter, and pressure data
// to add data to the specified sheet in the Google Sheets.
func HandleAddSheetData(client *sheets.Service, spreadsheetID string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the sheet name (e.g., Sheet1): ")
	sheetName, _ := reader.ReadString('\n')
	sheetName = strings.TrimSpace(sheetName)

	fmt.Print("Enter the column letter (e.g., A): ")
	columnLetter, _ := reader.ReadString('\n')
	columnLetter = strings.TrimSpace(columnLetter)

	readRange := fmt.Sprintf("%s!%s:%s", sheetName, columnLetter, columnLetter)

	fmt.Print("Enter pressure data to add to the sheet (e.g., 100): ")
	pressureValue, _ := reader.ReadString('\n')
	pressureValue = strings.TrimSpace(pressureValue)

	err := addSheetData(client, spreadsheetID, readRange, pressureValue)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("Pressure data '%s' added to sheet \n", pressureValue)
	}
}

// HandleCreateSheet allows the user to input the name of a new sheet to create in the Google Sheets.
func HandleCreateSheet(client *sheets.Service, spreadsheetID string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the name of the new sheet to create: ")
	sheetName, _ := reader.ReadString('\n')
	sheetName = strings.TrimSpace(sheetName)
	err := createNewSheet(client, spreadsheetID, sheetName)
	if err != nil {
		log.Printf("Error creating sheet: %v", err)
	} else {
		fmt.Printf("Sheet '%s' created successfully.\n", sheetName)
	}

}

// HandleUpdateData allows the user to input sheet name, cell range, and new pressure data
// to update data in the specified range of the Google Sheets.
func HandleUpdateData(srv *sheets.Service, spreadsheetID string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter sheet name: ")
	scanner.Scan()
	sheetName := scanner.Text()

	fmt.Print("Enter sheet range (e.g., A1): ")
	scanner.Scan()
	sheetRange := sheetName + "!" + scanner.Text()

	fmt.Print("Enter pressure data to add to the sheet (e.g., 100): ")
	scanner.Scan()
	pressureValue := scanner.Text()

	err := updateSheetData(srv, spreadsheetID, sheetRange, pressureValue)
	if err != nil {
		log.Println(err)
	}
}

// HandleRenameSheet allows the user to input the current sheet name and a new name
// to rename a sheet in the Google Sheets.
func HandleRenameSheet(client *sheets.Service, spreadsheetID string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the sheet which you want to rename: ")
	sheetName, _ := reader.ReadString('\n')
	sheetName = strings.TrimSpace(sheetName)

	fmt.Print("Enter the new name: ")
	newName, _ := reader.ReadString('\n')
	newName = strings.TrimSpace(newName)
	err := renameSheet(client, spreadsheetID, sheetName, newName)
	if err != nil {
		log.Printf("Error creating sheet: %v", err)
	} else {
		fmt.Printf("Sheet '%s' renamed successfully.\n", newName)
	}
}

// HandleDeleteSheet displays a list of available sheets in the Google Sheets table
// and allows the user to select a sheet to delete.
func HandleDeleteSheet(srv *sheets.Service, spreadsheetID string) {
	fmt.Println("Select the sheet to delete:")

	availableSheets, err := getSheetNames(srv, spreadsheetID)
	if err != nil {
		log.Fatalf("Unable to retrieve sheet names: %v", err)
	}

	for i, sheet := range availableSheets {
		fmt.Printf("%d. %s\n", i+1, sheet)
	}

	scanner := bufio.NewScanner(os.Stdin)
	selectedSheetIndex, err := getUserSheetChoice(scanner, availableSheets)
	if err != nil {
		fmt.Println(err)
		return
	}

	selectedSheet := availableSheets[selectedSheetIndex-1]
	err = deleteSheet(srv, spreadsheetID, selectedSheet)
	if err != nil {
		log.Printf("Error deleting sheet: %v", err)
	} else {
		fmt.Printf("Sheet '%s' deleted successfully.\n", selectedSheet)
	}
}

// HandleDeleteData allows the user to input sheet name and cell range
// to delete data from the specified range in the Google Sheets.
func HandleDeleteData(srv *sheets.Service, spreadsheetID string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter sheet name: ")
	scanner.Scan()
	sheetName := scanner.Text()

	fmt.Print("Enter sheet range (e.g., A1:D5): ")
	scanner.Scan()
	sheetRange := sheetName + "!" + scanner.Text()

	err := deleteDataFromSheet(srv, spreadsheetID, sheetRange)
	if err != nil {
		log.Printf("Error deleting data: %v", err)
	} else {
		fmt.Printf("Data in range '%s' on sheet '%s' has been deleted.\n", sheetRange, sheetName)
	}
}

// addSheetData adds the specified pressure data to the given range in the Google Sheets.
func addSheetData(client *sheets.Service, spreadsheetID, readRange, pressureValue string) error {
	response, err := client.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from Google Sheets: %v", err)
	}

	var nextRow int
	if len(response.Values) == 0 {
		nextRow = 1
	} else {
		for i, row := range response.Values {
			if len(row) == 0 {
				nextRow = i + 1
				break
			}
		}
		if nextRow == 0 {
			nextRow = len(response.Values) + 1
		}
	}

	writeRange := fmt.Sprintf("%s%d", readRange, nextRow)
	values := [][]interface{}{{pressureValue}}
	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err = client.Spreadsheets.Values.Update(spreadsheetID, writeRange, valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		return fmt.Errorf("unable to write pressure data to sheet: %v", err)
	}

	return nil
}

// createNewSheet creates a new sheet with the given name in the Google Sheets.
func createNewSheet(client *sheets.Service, spreadsheetID string, sheetName string) error {
	createSheetRequest := sheets.Request{
		AddSheet: &sheets.AddSheetRequest{
			Properties: &sheets.SheetProperties{
				Title: sheetName,
			},
		},
	}
	batchUpdateRequest := sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{&createSheetRequest},
	}

	_, err := client.Spreadsheets.BatchUpdate(spreadsheetID, &batchUpdateRequest).Do()
	if err != nil {
		return err
	}
	return nil
}

// loadSheetData loads data from the specified range in the Google Sheets.
func loadSheetData(client *sheets.Service, spreadsheetID string, readRange string) error {
	response, err := client.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from Google Sheets: %v", err)
	}

	if len(response.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Data from Google Sheets:")
		util.PrintTable(response.Values)
	}
	return nil
}

// updateSheetData updates the specified range in the Google Sheets with new pressure data.
func updateSheetData(srv *sheets.Service, spreadsheetID, rangeName string, newValues string) error {
	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{{newValues}},
	}
	_, err := srv.Spreadsheets.Values.Update(spreadsheetID, rangeName, valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		return err
	}
	return nil
}

// renameSheet renames the specified sheet in the Google Sheets to the new name.
func renameSheet(srv *sheets.Service, spreadsheetID, sheetName, newSheetName string) error {
	requests := []*sheets.Request{
		{
			UpdateSheetProperties: &sheets.UpdateSheetPropertiesRequest{
				Properties: &sheets.SheetProperties{
					SheetId: getSheetID(srv, spreadsheetID, sheetName),
					Title:   newSheetName,
				},
				Fields: "title",
			},
		},
	}

	_, err := srv.Spreadsheets.BatchUpdate(spreadsheetID, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: requests,
	}).Do()

	return err
}

// deleteSheet deletes the specified sheet from the Google Sheets.
func deleteSheet(client *sheets.Service, spreadsheetID string, sheetName string) error {
	deleteRequest := sheets.DeleteSheetRequest{
		SheetId: getSheetID(client, spreadsheetID, sheetName),
	}

	requests := []*sheets.Request{
		{
			DeleteSheet: &deleteRequest,
		},
	}

	batchUpdateRequest := sheets.BatchUpdateSpreadsheetRequest{
		Requests: requests,
	}

	_, err := client.Spreadsheets.BatchUpdate(spreadsheetID, &batchUpdateRequest).Do()
	if err != nil {
		return err
	}
	return nil
}

// deleteDataFromSheet clears data from the specified range in the Google Sheets.
func deleteDataFromSheet(srv *sheets.Service, spreadsheetID string, deleteRange string) error {
	_, err := srv.Spreadsheets.Values.Clear(spreadsheetID, deleteRange, &sheets.ClearValuesRequest{}).Do()
	if err != nil {
		return err
	}
	return nil
}

// getSheetID retrieves the sheet ID of the specified sheet name in the Google Sheets.
func getSheetID(client *sheets.Service, spreadsheetID string, sheetName string) int64 {
	resp, err := client.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve spreadsheet: %v", err)
	}

	for _, sheet := range resp.Sheets {
		if sheet.Properties.Title == sheetName {
			return sheet.Properties.SheetId
		}
	}

	log.Fatalf("Sheet not found: %s", sheetName)
	return 0
}

// getSheetNames retrieves a list of sheet names in the Google Sheets table.
func getSheetNames(client *sheets.Service, spreadsheetID string) ([]string, error) {
	resp, err := client.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		return nil, err
	}

	var sheetNames []string
	for _, sheet := range resp.Sheets {
		sheetNames = append(sheetNames, sheet.Properties.Title)
	}

	return sheetNames, nil
}

// getUserSheetChoice prompts the user to select a sheet from the provided list.
func getUserSheetChoice(scanner *bufio.Scanner, availableSheets []string) (int, error) {
	fmt.Print("Enter the number of the sheet to load data from: ")
	scanner.Scan()
	input := scanner.Text()
	selectedSheetIndex, err := strconv.Atoi(input)
	if err != nil || selectedSheetIndex < 1 || selectedSheetIndex > len(availableSheets) {
		return 0, fmt.Errorf("invalid input. Please select a valid sheet")
	}
	return selectedSheetIndex, nil
}
