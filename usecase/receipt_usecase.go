package usecase

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
	"context"
)

type ReceiptUseCase interface {
	GetMonthlyReceipts(ctx context.Context) ([]*model.ReceiptModel, error)
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
func (ru receiptUseCase) GetMonthlyReceipts(ctx context.Context) (receipts []*model.ReceiptModel, err error) {
	// Persistence（Repository）を呼出
	receipts, err = ru.receiptRepository.GetMonthlyReceipts(ctx)
	if err != nil {
		return nil, err
	}
	return receipts, nil
}
