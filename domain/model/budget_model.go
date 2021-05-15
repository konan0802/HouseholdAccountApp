package model

// BudgetModel 月の予実(予算・実績)のモデル
type BudgetModel struct {
	CategorieID      uint32
	BudgetValue      uint32
	PerformanceValue uint32
	Year             uint16
	Month            uint8
}

func NewBudgetModel(categorieID uint32, budgetvalue uint32, performancevalue uint32, year uint16, month uint8) *BudgetModel {
	return &BudgetModel{CategorieID: categorieID, BudgetValue: budgetvalue, PerformanceValue: performancevalue, Year: year, Month: month}
}
