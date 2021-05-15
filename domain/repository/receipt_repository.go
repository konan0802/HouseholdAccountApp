package repository

import (
	"HouseholdAccountApp/domain/model"
	"context"
)

// ReceiptRepository Receiptモデルのリポジトリ
type ReceiptRepository interface {
	//AddReceipt(newReceipt *model.ReceiptModel) (*model.ReceiptModel, error)
	//UpdateReceipt(newReceipt *model.ReceiptModel) error
	GetReceipt(ctx context.Context) (*model.ReceiptModel, error)
	//DeleteReceipt(name string) error
}
