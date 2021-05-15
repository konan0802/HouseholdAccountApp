package usecase

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
	"context"
)

type BudgetUseCase interface {
	GetBudget(ctx context.Context) (*model.BudgetModel, error)
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

// GetBudget Budgetデータを全件取得するためのユースケース
func (ru budgetUseCase) GetBudget(ctx context.Context) (budget *model.BudgetModel, err error) {
	// Persistence（Repository）を呼出
	budget, err = ru.budgetRepository.GetBudget(ctx)
	if err != nil {
		return nil, err
	}
	return budget, nil
}
