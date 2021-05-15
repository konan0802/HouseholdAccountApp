package usecase

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
	"context"
)

type ReceiptUseCase interface {
	GetReceipt(ctx context.Context) (*model.ReceiptModel, error)
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

// GetReceipt Receiptデータを全件取得するためのユースケース
func (ru receiptUseCase) GetReceipt(ctx context.Context) (receipt *model.ReceiptModel, err error) {
	// Persistence（Repository）を呼出
	receipt, err = ru.receiptRepository.GetReceipt(ctx)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}
