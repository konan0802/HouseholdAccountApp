package repository

import "HouseholdAccountApp/domain/model"

// BudgetRepository Budgetモデルのリポジトリ
type BudgetRepository interface {
	//AddBudget(newBudget *model.BudgetModel) (*model.BudgetModel, error)
	//UpdateBudget(newBudget *model.BudgetModel) error
	GetBudget(categorieID uint32, year uint16, month uint8) (*model.BudgetModel, error)
	//DeleteBudget(name string) error
}
