package usecase

import (
	"MillionaireApp/domain/model"
	"MillionaireApp/domain/repository"
)

type BudgetUseCase interface {
	GetMonthlyBudgets(mbr model.MonthlyBudgetsRequest) ([]*model.BudgetModel, error)
}

type budgetUseCase struct {
	budgetRepository repository.BudgetRepository
}

// NewBudgetUseCase : Budgetデータに関するUseCaseを生成
func NewBudgetUseCase(rr repository.BudgetRepository) BudgetUseCase {
	return &budgetUseCase{
		budgetRepository: rr,
	}
}

// GetMonthlyBudgets Budgetデータを全件取得するためのユースケース
func (ru budgetUseCase) GetMonthlyBudgets(mbr model.MonthlyBudgetsRequest) ([]*model.BudgetModel, error) {
	// Infra（Repository）を呼出
	budgets, err := ru.budgetRepository.GetMonthlyBudgets(mbr)
	if err != nil {
		return nil, err
	}
	return budgets, nil
}
