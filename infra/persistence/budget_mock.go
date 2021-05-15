package persistence

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
	"context"
)

type budgetPersistenceMock struct{}

// NewBudgetPersistence Budgetデータに関するPersistenceを生成
func NewBudgetPersistenceMock() repository.BudgetRepository {
	return &budgetPersistenceMock{}
}

// GetMonthlyBudgets 対象年月の予算を取得する
func (bp *budgetPersistenceMock) GetMonthlyBudgets(ctx context.Context) (*model.BudgetModel, error) {
	bm := model.BudgetModel{
		CategorieID:      0,
		BudgetValue:      30000,
		PerformanceValue: 0,
		Year:             2021,
		Month:            05,
	}
	return &bm, nil
}
