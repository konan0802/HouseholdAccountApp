package repository

import (
	"MillionaireApp/domain/model"
)

// ReceiptRepository Receiptモデルのリポジトリ
type ReceiptRepository interface {
	AddReceipt(arr model.AddReceiptRequest) (*model.ReceiptModel, error)
	GetMonthlyReceipts(mrr model.MonthlyReceiptsRequest) ([]*model.ReceiptModel, error)
}
