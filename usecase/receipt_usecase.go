package usecase

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
	"context"
)

type ReceiptUseCase interface {
	GetMonthlyReceipts(mrr model.MonthlyReceiptsRequest) ([]*model.ReceiptModel, error)
	AddReceipt(ctx context.Context) error
}

type receiptUseCase struct {
	receiptRepository repository.ReceiptRepository
}

// NewReceiptUseCase : Receiptデータに関するUseCaseを生成
func NewReceiptUseCase(rr repository.ReceiptRepository) ReceiptUseCase {
	return &receiptUseCase{
		receiptRepository: rr,
	}
}

// GetMonthlyReceipts Receiptデータを全件取得するためのユースケース
func (ru receiptUseCase) GetMonthlyReceipts(mrr model.MonthlyReceiptsRequest) ([]*model.ReceiptModel, error) {
	// Infra（Repository）を呼出
	receipts, err := ru.receiptRepository.GetMonthlyReceipts(mrr)
	if err != nil {
		return nil, err
	}
	return receipts, nil
}

// AddReceipt Receiptデータを全件取得するためのユースケース
func (ru receiptUseCase) AddReceipt(ctx context.Context) error {
	// Infra（Repository）を呼出
	err := ru.receiptRepository.AddReceipt(ctx)
	if err != nil {
		return err
	}
	return nil
}
