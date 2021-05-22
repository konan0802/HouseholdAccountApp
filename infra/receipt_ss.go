package infra

import (
	"MillionaireApp/config"
	"MillionaireApp/domain/model"
	"MillionaireApp/domain/repository"
	"log"
	"strconv"
	"time"

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

	// 取得範囲の作成
	getSheet := strconv.Itoa(mrr.Year) + "年" + strconv.Itoa(mrr.Month) + "月"
	getSheetRange := getSheet + config.GetReciptRange

	// レシートの取得
	resp, err := rp.SS.SheetService.Spreadsheets.Values.Get(rp.SS.SpreadsheetId, getSheetRange).Do()
	if err != nil {
		log.Fatal(err)
	}

	var rms []*model.ReceiptModel
	for _, eachReceipt := range resp.Values {
		price, err := strconv.Atoi(eachReceipt[3].(string))
		if err != nil {
			log.Fatal(err)
		}
		rm := model.ReceiptModel{
			CategorieID: config.CategorieNametoInt[eachReceipt[1].(string)],
			Description: eachReceipt[2].(string),
			Price:       price,
			Datetime:    eachReceipt[0].(string),
		}
		rms = append(rms, &rm)
	}

	return rms, nil
}

// AddReceipt レシートを追加する
func (rp *ReceiptInfraSS) AddReceipt(arr model.AddReceiptRequest) (*model.ReceiptModel, error) {

	// 追加レシートの時間をパース
	t, err := time.Parse(config.TimeLayout, arr.Datetime)
	if err != nil {
		log.Fatal(err)
	}

	// 取得対象シート
	getSheet := strconv.Itoa(t.Year()) + "年" + strconv.Itoa(int(t.Month())) + "月"

	// 取得対象範囲（インサート行の確認）
	getSheetRange := getSheet + config.CheckReciptColumnsRange

	// インサート行の確認
	val, err := rp.SS.SheetService.Spreadsheets.Values.Get(rp.SS.SpreadsheetId, getSheetRange).Do()
	if err != nil {
		log.Fatal(err)
	}

	// インサート開始行の計算
	start := len(val.Values)
	startRow := strconv.Itoa(config.ReciptStartRow + start)
	inputeRow := getSheet + "!" + config.ReceiptStartCol + startRow + ":" + config.ReceiptEndCol + startRow

	// インサートする値
	valueRange := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values: [][]interface{}{
			[]interface{}{arr.Datetime, arr.CategorieName, arr.Description, arr.Price},
		},
	}

	// スプレッドシートにインサート
	_, err = rp.SS.SheetService.Spreadsheets.Values.Update(rp.SS.SpreadsheetId, inputeRow, valueRange).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatalf("Unable to write value. %v", err)
	}

	return &model.ReceiptModel{
		CategorieID: config.CategorieNametoInt[arr.CategorieName],
		Description: arr.Description,
		Price:       arr.Price,
		Datetime:    arr.Datetime,
	}, nil
}
