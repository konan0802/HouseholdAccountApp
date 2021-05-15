package model

// BudgetModel 月の予実(予算・実績)のモデル
type BudgetModel struct {
	CategorieID      uint32
	BudgetValue      uint32
	PerformanceValue uint32
	Year             uint16
	Month            uint8
}
