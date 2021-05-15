package repository

import (
	"HouseholdAccountApp/domain/model"
	"context"
)

// ReceiptRepository Receiptモデルのリポジトリ
type ReceiptRepository interface {
	//AddReceipt(ctx context.Context)
	GetMonthlyReceipts(ctx context.Context) ([]*model.ReceiptModel, error)
}
