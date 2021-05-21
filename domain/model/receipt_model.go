package model

// ReceiptModel レシート(単品の購入情報)のモデル
type ReceiptModel struct {
	ReceiptID   int
	CategorieID int
	Description string
	Price       int
	Datetime    string
}

// MonthlyReceiptsRequest MonthlyReceiptsのリクエストパラメータ
type MonthlyReceiptsRequest struct {
	Year  int `validate:"required,gte=2000,lte=3000"`
	Month int `validate:"required,gte=1,lte=12"`
}

// AddReceiptRequest AddReceiptのリクエストパラメータ
type AddReceiptRequest struct {
	CategorieName string `validate:"required"`
	Description   string `validate:"required"`
	Price         int    `validate:"required,gte=0"`
	Datetime      string `validate:"required"`
}
