package infra

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Spreadsheets Spreadsheetsを取り扱う上でのstruct
type SpreadSheets struct {
	SheetService  *sheets.Service
	SpreadsheetId string
}

// NewSpreadsheets Spreadsheetsインスタンスの生成
func NewSpreadsheets(spreadsheetId string, credentialFilePath string) SpreadSheets {

	// クレデンシャルの作成
	credential := option.WithCredentialsFile(credentialFilePath)

	// シートオブジェクトの作成
	sheetService, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	// シートの取得テスト
	_, err = sheetService.Spreadsheets.Get(spreadsheetId).Do()
	if err != nil {
		log.Fatalf("Unable to get Spreadsheets. %v", err)
	}

	return SpreadSheets{
		SheetService:  sheetService,
		SpreadsheetId: spreadsheetId,
	}
}
