package domain

//BudgetRepository: Budgetモデルのリポジトリ
type BudgetRepository interface {
	CreateBudget(newBudget *BudgetModel) (*BudgetModel, error)
	UpdateBudget(newBudget *BudgetModel) error
	GetBudget(name string) (*BudgetModel, error)
	DeleteBudget(name string) error
}
