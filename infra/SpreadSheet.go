package infra

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Spreadsheets Spreadsheetsを取り扱う上でのstruct
type SpreadSheets struct {
	SheetService *sheets.Service
}

// NewSpreadsheets Spreadsheetsインスタンスの生成
func NewSpreadsheets(spreadsheetId string, credentialFilePath string) *SpreadSheets {

	credential := option.WithCredentialsFile(credentialFilePath)

	sheetService, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	return &SpreadSheets{
		SheetService: sheetService,
	}
}
