package model

import "time"

// ReceiptModel レシート(単品の購入情報)のモデル
type ReceiptModel struct {
	ReceiptID   uint32    `json:"receiptID"`
	CategorieID uint32    `json:"categorieID"`
	Description string    `json:"description"`
	Price       uint32    `json:"price"`
	Datetime    time.Time `json:"datetime"`
}

// MonthlyReceiptsRequest MonthlyReceiptsのリクエストパラメータ
type MonthlyReceiptsRequest struct {
	Year  uint16 `json:"year" validate:"required,gte=2000,lte=3000"`
	Month uint8  `json:"month" validate:"required,gte=1,lte=12"`
}

// AddReceiptRequest AddReceiptのリクエストパラメータ
type AddReceiptRequest struct {
	Year  uint16 `json:"year" validate:"required,gte=2000,lte=3000"`
	Month uint8  `json:"month" validate:"required,gte=1,lte=12"`
}
