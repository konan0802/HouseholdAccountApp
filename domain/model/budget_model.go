package model

// BudgetModel 月の予実(予算・実績)のモデル
type BudgetModel struct {
	CategorieID      uint32
	BudgetValue      uint32
	PerformanceValue uint32
	ProgressValue    uint32
	Year             uint32
	Month            uint32
}

// MonthlyBudgetsRequest Monthlybudgetのリクエストパラメータ
type MonthlyBudgetsRequest struct {
	Year  uint32 `validate:"required,gte=2000,lte=3000"`
	Month uint32 `validate:"required,gte=1,lte=12"`
}
