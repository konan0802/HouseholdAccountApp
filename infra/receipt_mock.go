package infra

import (
	"MillionaireApp/domain/model"
	"MillionaireApp/domain/repository"
)

type receiptInfraMock struct{}

// NewReceiptInfraMock Receiptデータに関するInfraを生成
func NewReceiptInfraMock() repository.ReceiptRepository {
	return &receiptInfraMock{}
}

// GetMonthlyReceipts 対象年月のレシートを取得する
func (rp *receiptInfraMock) GetMonthlyReceipts(mrr model.MonthlyReceiptsRequest) ([]*model.ReceiptModel, error) {
	rm1 := model.ReceiptModel{
		CategorieID: 0,
		Description: "コーヒー",
		Price:       300,
		Datetime:    "2006-01-02 15:04:05",
	}
	rm2 := model.ReceiptModel{
		CategorieID: 1,
		Description: "椅子",
		Price:       4000,
		Datetime:    "2006-01-02 15:04:05",
	}
	return []*model.ReceiptModel{&rm1, &rm2}, nil
}

// AddReceipt レシートを追加する
func (rp *receiptInfraMock) AddReceipt(mrr model.AddReceiptRequest) (*model.ReceiptModel, error) {
	rm := model.ReceiptModel{
		ReceiptID:   2,
		CategorieID: 1,
		Description: "椅子",
		Price:       4000,
		Datetime:    "2006-01-02 15:04:05",
	}
	return &rm, nil
}
