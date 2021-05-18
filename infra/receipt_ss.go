package infra

import (
	"MillionaireApp/domain/model"
	"MillionaireApp/domain/repository"
	"context"
	"time"
)

type receiptInfraSS struct{}

// NewReceiptInfraSS Receiptデータに関するInfraを生成
func NewReceiptInfraSS() repository.ReceiptRepository {
	return &receiptInfraMock{}
}

// GetMonthlyReceipts 対象年月の予算を取得する
func (rp *receiptInfraSS) GetMonthlyReceipts(ctx context.Context) ([]*model.ReceiptModel, error) {
	rm1 := model.ReceiptModel{
		ReceiptID:   1,
		CategorieID: 0,
		Description: "コーヒー",
		Price:       300,
		Datetime:    time.Date(2021, time.May, 15, 5, 0, 0, 0, time.UTC),
	}
	rm2 := model.ReceiptModel{
		ReceiptID:   2,
		CategorieID: 1,
		Description: "椅子",
		Price:       4000,
		Datetime:    time.Date(2021, time.May, 16, 5, 0, 0, 0, time.UTC),
	}
	return []*model.ReceiptModel{&rm1, &rm2}, nil
}

// AddReceipt レシートを追加する
func (rp *receiptInfraSS) AddReceipt(ctx context.Context) error {
	return nil
}
