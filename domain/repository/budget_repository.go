package repository

import (
	"HouseholdAccountApp/domain/model"
)

// BudgetRepository Budgetモデルのリポジトリ
type BudgetRepository interface {
	GetMonthlyBudgets(mbr model.MonthlyBudgetsRequest) ([]*model.BudgetModel, error)
}
