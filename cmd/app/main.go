package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/VladdPp/googlesheetscli/internal/googlesheets"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Path to your service account JSON file
const credentialsFile = "./credentials.json"

func main() {

	// Create a new Google Sheets client
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(credentialsFile), option.WithScopes(sheets.SpreadsheetsScope))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	// The ID of the spreadsheet you want to access.
	spreadsheetID := os.Getenv("SPREAD_SHEET_ID")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Google Sheets Console App!")

	for {
		fmt.Print("Enter an action (load, create, add, update, delete) or 'exit' to quit: ")
		actionInput, _ := reader.ReadString('\n')
		actionInput = strings.TrimSpace(actionInput)

		switch actionInput {
		case "load":
			fmt.Print("Enter load option (data, sheet): ")
			loadOption, _ := reader.ReadString('\n')
			loadOption = strings.TrimSpace(loadOption)

			switch loadOption {
			case "data":
				googlesheets.HandleLoadRangeSheet(srv, spreadsheetID)
			case "sheet":
				googlesheets.HandleLoadSheet(srv, spreadsheetID)
			default:
				fmt.Println("Invalid load option.")
			}
		case "add":
			fmt.Print("Enter add option (data): ")
			saveOption, _ := reader.ReadString('\n')
			saveOption = strings.TrimSpace(saveOption)

			if saveOption == "data" {
				fmt.Println("Adding data to Google Sheets...")
				googlesheets.HandleAddSheetData(srv, spreadsheetID)
			} else {
				fmt.Println("Invalid add option.")
			}
		case "create":
			fmt.Print("Enter create option (sheet): ")
			createOption, _ := reader.ReadString('\n')
			createOption = strings.TrimSpace(createOption)

			if createOption == "sheet" {
				googlesheets.HandleCreateSheet(srv, spreadsheetID)
			} else {
				fmt.Println("Invalid create option.")
			}
		case "update":
			fmt.Print("Enter update data option or rename sheet (data, rename): ")
			updateOption, _ := reader.ReadString('\n')
			updateOption = strings.TrimSpace(updateOption)

			switch updateOption {
			case "rename":
				googlesheets.HandleRenameSheet(srv, spreadsheetID)
			case "data":
				googlesheets.HandleUpdateData(srv, spreadsheetID)
			default:
				fmt.Println("Invalid update option.")
			}

		case "delete":
			fmt.Print("Enter delete option (sheet, data): ")
			deleteOption, _ := reader.ReadString('\n')
			deleteOption = strings.TrimSpace(deleteOption)

			switch deleteOption {
			case "sheet":
				googlesheets.HandleDeleteSheet(srv, spreadsheetID)
			case "data":
				googlesheets.HandleDeleteData(srv, spreadsheetID)
			default:
				fmt.Println("Invalid delete option.")
			}

		case "exit":
			fmt.Println("Exiting the program.")
			return

		default:
			fmt.Println("Invalid action. Please specify a valid action (load, save, update, delete, exit).")
		}
	}
}
