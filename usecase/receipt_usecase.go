package usecase

import (
	"MillionaireApp/domain/model"
	"MillionaireApp/domain/repository"
)

type ReceiptUseCase interface {
	GetMonthlyReceipts(mrr model.MonthlyReceiptsRequest) ([]*model.ReceiptModel, error)
	AddReceipt(mrr model.AddReceiptRequest) (*model.ReceiptModel, error)
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
func (ru receiptUseCase) AddReceipt(arr model.AddReceiptRequest) (*model.ReceiptModel, error) {
	// Infra（Repository）を呼出
	receipt, err := ru.receiptRepository.AddReceipt(arr)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}
