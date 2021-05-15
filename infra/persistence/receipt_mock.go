package persistence

import (
	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/domain/repository"
	"time"
)

type receiptPersistence struct{}

// NewReceiptPersistence Receiptデータに関するPersistenceを生成
func NewReceiptPersistence() repository.ReceiptRepository {
	return &receiptPersistence{}
}

// GetReceipt 対象年月の予算を取得する
func (bp *receiptPersistence) GetReceipt(year uint16, month uint8, day uint8) (*model.ReceiptModel, error) {
	rm := model.ReceiptModel{
		ReceiptID:   1,
		CategorieID: 0,
		ProductName: "コーヒー",
		Price:       300,
		Methods:     "suica",
		Datetime:    time.Date(2021, time.May, 15, 5, 0, 0, 0, time.UTC),
	}
	return &rm, nil
}
