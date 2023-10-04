# Google Sheets Console Application

This is a command-line application that allows users to interact with Google Sheets. Users can perform various actions such as loading data from sheets, creating new sheets, adding data, updating data, renaming sheets, deleting sheets, and deleting data from specific ranges.

## Prerequisites

* Go programming language installed on your system.
* Google Cloud Console account with Google Sheets API enabled.
* Service account credentials JSON file.
* A Google Sheets spreadsheet ID.

## Setup Google Sheets Credentials

1. Create a service account in the Google Cloud Console.
2. Enable the Google Sheets API for the created project.
3. Create a service account and download the JSON credentials file.
4. Save the JSON credentials file as `credentials.json` in the project directory.

## Set Up Environment Variables

1. Create a `.env` file in the project directory.
2. Add the following environment variables to the `.env` file:

```
SPREAD_SHEET_ID=<Your Spreadsheet ID>
```

Replace `<Your Spreadsheet ID>` with the ID of the Google Sheets spreadsheet you want to access.

## Running the Application

1. Clone the repository:

```
git clone <repository-url>
cd googlesheetscli
```

2. Install Dependencies:

```
go mod tidy
```

3. Build the Application:

```
go build ./cmd/app/main.go
```

3. Run the Application:

```
./main
```

## Available Commands


###  Load Data from Sheet:
  * Select the load option and choose between loading entire sheet data or specific sheet range.
### Create New Sheet:
  * Select the create option and provide a name for the new sheet.
### Add Data to Sheet:
 * Select the add option, provide the sheet name, column letter, and pressure data to add to the sheet.
### Update Data in Sheet:
 * Select the update option, provide the sheet name, cell range, and new pressure data to update the data.
### Rename Sheet:
 * Select the update option and choose rename. Provide the current sheet name and the new name for the sheet.
### Delete Sheet:
 * Select the delete option and choose sheet. Choose the sheet you want to delete.
### Delete Data from Sheet:
 * Select the delete option and choose data. Provide the sheet name and cell range to delete data.
### Exit the Application:
 * Select the exit option to quit the program.
   
Note: Ensure you have the necessary permissions and access to the Google Sheets spreadsheet specified in the SPREAD_SHEET_ID environment variable.
