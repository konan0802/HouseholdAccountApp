package repository

import (
	"HouseholdAccountApp/domain/model"
	"context"
)

// ReceiptRepository Receiptモデルのリポジトリ
type ReceiptRepository interface {
	AddReceipt(ctx context.Context) error
	GetMonthlyReceipts(mrr model.MonthlyReceiptsRequest) ([]*model.ReceiptModel, error)
}
