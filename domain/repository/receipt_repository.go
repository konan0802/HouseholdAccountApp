package repository

import "HouseholdAccountApp/domain/model"

// ReceiptRepository Receiptモデルのリポジトリ
type ReceiptRepository interface {
	//AddReceipt(newReceipt *model.ReceiptModel) (*model.ReceiptModel, error)
	//UpdateReceipt(newReceipt *model.ReceiptModel) error
	GetReceipt(year uint16, month uint8, day uint8) (*model.ReceiptModel, error)
	//DeleteReceipt(name string) error
}
