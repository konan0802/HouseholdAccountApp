package model

// BudgetModel 月の予実(予算・実績)のモデル
type BudgetModel struct {
	CategorieID      uint32 `json:"categorieID"`
	BudgetValue      uint32 `json:"budgetValue"`
	PerformanceValue uint32 `json:"performanceValue"`
	Year             uint16 `json:"year"`
	Month            uint8  `json:"month"`
}

// MonthlyBudgetsRequest Monthlybudgetのリクエストパラメータ
type MonthlyBudgetsRequest struct {
	Year  uint16 `json:"year" validate:"required,gte=2000,lte=3000"`
	Month uint8  `json:"month" validate:"required,gte=1,lte=12"`
}
