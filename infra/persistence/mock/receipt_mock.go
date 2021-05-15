package persistence

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
	"context"
	"time"
)

type receiptPersistenceMock struct{}

// NewReceiptPersistenceMock Receiptデータに関するPersistenceを生成
func NewReceiptPersistenceMock() repository.ReceiptRepository {
	return &receiptPersistenceMock{}
}

// GetMonthlyReceipts 対象年月の予算を取得する
func (rp *receiptPersistenceMock) GetMonthlyReceipts(ctx context.Context) ([]*model.ReceiptModel, error) {
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