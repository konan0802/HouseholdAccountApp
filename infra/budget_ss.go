package infra

import (
	"MillionaireApp/domain/model"
	"MillionaireApp/domain/repository"
	"context"
)

type budgetInfraSS struct{}

// NewBudgetInfraSS Budgetデータに関するInfraを生成
func NewBudgetInfraSS() repository.BudgetRepository {
	return &budgetInfraMock{}
}

// GetMonthlyBudgets 対象年月の予算を取得する
func (bp *budgetInfraSS) GetMonthlyBudgets(ctx context.Context) (*model.BudgetModel, error) {
	bm := model.BudgetModel{
		CategorieID:      0,
		BudgetValue:      30000,
		PerformanceValue: 0,
		Year:             2021,
		Month:            05,
	}
	return &bm, nil
}
