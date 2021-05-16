package model

import "time"

// ReceiptModel レシート(単品の購入情報)のモデル
type ReceiptModel struct {
	ReceiptID   uint32
	CategorieID uint32
	Description string
	Price       uint32
	Datetime    time.Time
}

// MonthlyReceiptsRequest MonthlyReceiptsのリクエストパラメータ
type MonthlyReceiptsRequest struct {
	Year  uint32 `validate:"required,gte=2000,lte=3000"`
	Month uint32 `validate:"required,gte=1,lte=12"`
}

// AddReceiptRequest AddReceiptのリクエストパラメータ
type AddReceiptRequest struct {
	CategorieName string `validate:"required"`
	Description   string `validate:"required"`
	Price         uint32 `validate:"required,gte=0"`
	Datetime      string `validate:"required"`
}
