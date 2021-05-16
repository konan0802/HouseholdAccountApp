package infra

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
)

type budgetInfraMock struct{}

// NewBudgetInfraMock Budgetデータに関するInfraを生成
func NewBudgetInfraMock() repository.BudgetRepository {
	return &budgetInfraMock{}
}

// GetMonthlyBudgets 対象年月の予算を取得する
func (bp *budgetInfraMock) GetMonthlyBudgets(mbr model.MonthlyBudgetsRequest) ([]*model.BudgetModel, error) {
	bm1 := model.BudgetModel{
		CategorieID:      0,
		BudgetValue:      30000,
		PerformanceValue: 0,
		Year:             mbr.Year,
		Month:            mbr.Month,
	}
	bm2 := model.BudgetModel{
		CategorieID:      1,
		BudgetValue:      40000,
		PerformanceValue: 2000,
		Year:             mbr.Year,
		Month:            mbr.Month,
	}
	return []*model.BudgetModel{&bm1, &bm2}, nil
}
