package infra

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
	"context"
)

type budgetInfraMock struct{}

// NewBudgetInfraMock Budgetデータに関するInfraを生成
func NewBudgetInfraMock() repository.BudgetRepository {
	return &budgetInfraMock{}
}

// GetMonthlyBudgets 対象年月の予算を取得する
func (bp *budgetInfraMock) GetMonthlyBudgets(ctx context.Context) (*model.BudgetModel, error) {
	bm := model.BudgetModel{
		CategorieID:      0,
		BudgetValue:      30000,
		PerformanceValue: 0,
		Year:             2021,
		Month:            05,
	}
	return &bm, nil
}
