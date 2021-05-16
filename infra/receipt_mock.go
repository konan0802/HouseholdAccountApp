package infra

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
	"time"
)

type receiptInfraMock struct{}

// NewReceiptInfraMock Receiptデータに関するInfraを生成
func NewReceiptInfraMock() repository.ReceiptRepository {
	return &receiptInfraMock{}
}

// GetMonthlyReceipts 対象年月のレシートを取得する
func (rp *receiptInfraMock) GetMonthlyReceipts(mrr model.MonthlyReceiptsRequest) ([]*model.ReceiptModel, error) {
	rm1 := model.ReceiptModel{
		ReceiptID:   1,
		CategorieID: 0,
		Description: "コーヒー",
		Price:       300,
		Datetime:    time.Date(int(mrr.Year), time.May, int(mrr.Month), 5, 0, 0, 0, time.UTC),
	}
	rm2 := model.ReceiptModel{
		ReceiptID:   2,
		CategorieID: 1,
		Description: "椅子",
		Price:       4000,
		Datetime:    time.Date(int(mrr.Year), time.May, int(mrr.Month), 5, 0, 0, 0, time.UTC),
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
		Datetime:    time.Date(2021, time.May, 16, 5, 0, 0, 0, time.UTC),
	}
	return &rm, nil
}
