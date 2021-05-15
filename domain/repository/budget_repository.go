package repository

import (
	"HouseholdAccountApp/domain/model"
	"context"
)

// BudgetRepository Budgetモデルのリポジトリ
type BudgetRepository interface {
	//AddBudget(newBudget *model.BudgetModel) (*model.BudgetModel, error)
	//UpdateBudget(newBudget *model.BudgetModel) error
	GetBudget(ctx context.Context) (*model.BudgetModel, error)
	//DeleteBudget(name string) error
}
