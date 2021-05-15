package persistence

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
)

type budgetPersistence struct{}

// NewBudgetPersistence Budgetデータに関するPersistenceを生成
func NewBudgetPersistence() repository.BudgetRepository {
	return &budgetPersistence{}
}

// GetBudget 対象年月の予算を取得する
func (bp *budgetPersistence) GetBudget(categorieID uint32, year uint16, month uint8) (*model.BudgetModel, error) {
	bm := model.BudgetModel{
		CategorieID:      0,
		BudgetValue:      30000,
		PerformanceValue: 0,
		Year:             2021,
		Month:            05,
	}
	return &bm, nil
}
