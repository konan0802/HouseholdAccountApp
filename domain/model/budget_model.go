package model

// BudgetModel 月の予実(予算・実績)のモデル
type BudgetModel struct {
	CategorieID      uint32
	BudgetValue      uint32
	PerformanceValue uint32
	Year             uint16
	Month            uint8
}

func NewBudgetModel(name string, budgetvalue uint32, performancevalue uint32, year uint8, month uint8) *BudgetModel {
	return &BudgetModel{Name: name, BudgetValue: budgetvalue, PerformanceValue: performancevalue, Year: year, Month: month}
}
