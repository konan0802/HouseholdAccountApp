package infra

import (
	"MillionaireApp/config"
	"MillionaireApp/domain/model"
	"MillionaireApp/domain/repository"
	"log"
	"strconv"
)

type budgetInfraSS struct {
	SS SpreadSheets
}

// NewBudgetInfraSS Budgetデータに関するInfraを生成
func NewBudgetInfraSS(sheetService SpreadSheets) repository.BudgetRepository {
	return &budgetInfraSS{
		SS: sheetService,
	}
}

// GetMonthlyBudgets 対象年月の予算を取得する
func (bp *budgetInfraSS) GetMonthlyBudgets(mbr model.MonthlyBudgetsRequest) ([]*model.BudgetModel, error) {

	var bms []*model.BudgetModel
	for cName, cCol := range config.BudgetCategorytoCol {

		getRange := strconv.Itoa(mbr.Year) + "年" + strconv.Itoa(mbr.Month) + "月" + "!" + cCol + config.BugetStartCol + ":" + cCol + config.BugetEndCol
		resp, err := bp.SS.SheetService.Spreadsheets.Values.Get(bp.SS.SpreadsheetId, getRange).Do()
		if err != nil {
			log.Fatal(err)
		}

		budgetValue, err := strconv.Atoi(resp.Values[0][0].(string))
		if err != nil {
			log.Fatal(err)
		}
		performanceValue, err := strconv.Atoi(resp.Values[1][0].(string))
		if err != nil {
			log.Fatal(err)
		}
		progressValue, err := strconv.Atoi(resp.Values[2][0].(string))
		if err != nil {
			log.Fatal(err)
		}

		bm := model.BudgetModel{
			CategorieID:      config.CategorieNametoInt[cName],
			BudgetValue:      budgetValue,
			PerformanceValue: performanceValue,
			ProgressValue:    progressValue,
			Year:             mbr.Year,
			Month:            mbr.Month,
		}
		bms = append(bms, &bm)
	}

	return bms, nil
}
