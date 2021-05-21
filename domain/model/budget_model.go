package model

// BudgetModel 月の予実(予算・実績)のモデル
type BudgetModel struct {
	CategorieID      int
	BudgetValue      int
	PerformanceValue int
	ProgressValue    int
	Year             int
	Month            int
}

// MonthlyBudgetsRequest Monthlybudgetのリクエストパラメータ
type MonthlyBudgetsRequest struct {
	Year  int `validate:"required,gte=2000,lte=3000"`
	Month int `validate:"required,gte=1,lte=12"`
}
