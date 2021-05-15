package usecase

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
)

type ReceiptUseCase interface {
	GetReceipt(year uint16, month uint8, day uint8) (*model.ReceiptModel, error)
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
func (ru receiptUseCase) GetReceipt(year uint16, month uint8, day uint8) (receipt *model.ReceiptModel, err error) {
	// Persistence（Repository）を呼出
	receipt, err = ru.receiptRepository.GetReceipt(year, month, day)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}
