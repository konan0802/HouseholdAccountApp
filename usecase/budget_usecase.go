package usecase

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
)

type BudgetUseCase interface {
	GetBudget(categorieID uint32, year uint16, month uint8) (*model.BudgetModel, error)
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
func (ru budgetUseCase) GetBudget(categorieID uint32, year uint16, month uint8) (budget *model.BudgetModel, err error) {
	// Persistence（Repository）を呼出
	budget, err = ru.budgetRepository.GetBudget(categorieID, year, month)
	if err != nil {
		return nil, err
	}
	return budget, nil
}
