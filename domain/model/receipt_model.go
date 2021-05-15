package model

import "time"

// ReceiptModel レシート(単品の購入情報)のモデル
type ReceiptModel struct {
	ReceiptID   uint32
	CategorieID uint32
	ProductName string
	Price       uint32
	Methods     string
	Datetime    time.Time
}

func NewReceiptModel(receiptID uint32, categorieID uint32, price uint32, datetime time.Time, methods string) *ReceiptModel {
	return &ReceiptModel{ReceiptID: receiptID, CategorieID: categorieID, Price: price, Datetime: datetime, Methods: methods}
}
