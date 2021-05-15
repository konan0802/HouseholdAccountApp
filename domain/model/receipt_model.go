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
