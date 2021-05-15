package repository

import (
	"HouseholdAccountApp/domain/model"
	"context"
)

// BudgetRepository Budgetモデルのリポジトリ
type BudgetRepository interface {
	GetMonthlyBudgets(ctx context.Context) (*model.BudgetModel, error)
}
