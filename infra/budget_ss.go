package infra

import (
	"MillionaireApp/domain/model"
	"MillionaireApp/domain/repository"
)

type budgetInfraSS struct {
	SS SpreadSheets
}

// NewBudgetInfraSS Budgetデータに関するInfraを生成
func NewBudgetInfraSS(ss SpreadSheets) repository.BudgetRepository {
	return &budgetInfraSS{
		SS: ss,
	}
}

// GetMonthlyBudgets 対象年月の予算を取得する
func (bp *budgetInfraSS) GetMonthlyBudgets(mbr model.MonthlyBudgetsRequest) ([]*model.BudgetModel, error) {
	bm1 := model.BudgetModel{
		CategorieID:      0,
		BudgetValue:      30000,
		PerformanceValue: 0,
		ProgressValue:    50,
		Year:             mbr.Year,
		Month:            mbr.Month,
	}
	bm2 := model.BudgetModel{
		CategorieID:      1,
		BudgetValue:      40000,
		PerformanceValue: 2000,
		ProgressValue:    50,
		Year:             mbr.Year,
		Month:            mbr.Month,
	}
	return []*model.BudgetModel{&bm1, &bm2}, nil
}
