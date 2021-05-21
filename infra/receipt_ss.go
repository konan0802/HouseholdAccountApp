package infra

import (
	"MillionaireApp/config"
	"MillionaireApp/domain/model"
	"MillionaireApp/domain/repository"
	"log"
	"strconv"

	"google.golang.org/api/sheets/v4"
)

type ReceiptInfraSS struct {
	SS SpreadSheets
}

// NewReceiptInfraSS Receiptデータに関するInfraを生成
func NewReceiptInfraSS(sheetService SpreadSheets) repository.ReceiptRepository {
	return &ReceiptInfraSS{
		SS: sheetService,
	}
}

// GetMonthlyReceipts 対象年月の予算を取得する
func (rp *ReceiptInfraSS) GetMonthlyReceipts(mrr model.MonthlyReceiptsRequest) ([]*model.ReceiptModel, error) {
	rm1 := model.ReceiptModel{
		ReceiptID:   1,
		CategorieID: 0,
		Description: "コーヒー",
		Price:       300,
		Datetime:    "2006-01-02 15:04:05",
	}
	rm2 := model.ReceiptModel{
		ReceiptID:   2,
		CategorieID: 1,
		Description: "椅子",
		Price:       4000,
		Datetime:    "2006-01-02 15:04:05",
	}
	return []*model.ReceiptModel{&rm1, &rm2}, nil
}

// AddReceipt レシートを追加する
func (rp *ReceiptInfraSS) AddReceipt(arr model.AddReceiptRequest) (*model.ReceiptModel, error) {

	// スプレッドシートにインサートする行
	val, err := rp.SS.SheetService.Spreadsheets.Values.Get(rp.SS.SpreadsheetId, config.CheckReciptColumnsRange).Do()
	if err != nil {
		log.Fatal(err)
	}
	start := len(val.Values)
	startRow := strconv.Itoa(config.ReciptStartRow + start)
	inputeRow := config.SheetName + "!B" + startRow + ":F" + startRow

	// スプレッドシートにインサートする値
	valueRange := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values: [][]interface{}{
			[]interface{}{start, arr.Datetime, arr.CategorieName, arr.Description, arr.Price},
		},
	}

	// スプレッドシートにインサート
	_, err = rp.SS.SheetService.Spreadsheets.Values.Update(rp.SS.SpreadsheetId, inputeRow, valueRange).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatalf("Unable to write value. %v", err)
	}

	return &model.ReceiptModel{
		ReceiptID:   0,
		CategorieID: config.CategorieNametoInt[arr.CategorieName],
		Description: arr.Description,
		Price:       arr.Price,
		Datetime:    arr.Datetime,
	}, nil
}

// CheckReciptColumns レシートの挿入先チェック
func (rp *ReceiptInfraSS) CheckReciptColumns() ([][]interface{}, error) {
	resp, err := rp.SS.SheetService.Spreadsheets.Values.Get(rp.SS.SpreadsheetId, config.CheckReciptColumnsRange).Do()
	if err != nil {
		log.Fatal(err)
	}
	return resp.Values, nil
}
