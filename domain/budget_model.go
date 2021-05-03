package domain

import "time"

//BudgetModel: 月の予実(予算・実績)のモデル
type BudgetModel struct {
	Name             string
	BudgetValue      uint32
	PerformanceValue uint32
	Date             time.Time
}

func NewBudgetModel(name string, budgetvalue uint32, performancevalue uint32, datetime time.Time) *BudgetModel {
	return &BudgetModel{Name: name, BudgetValue: budgetvalue, PerformanceValue: performancevalue, Date: datetime}
}
