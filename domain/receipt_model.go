package domain

import "time"

//ReciptModel: レシート(単品の購入情報)のモデル
type ReciptModel struct {
	ID       int64
	Name     string
	Price    uint32
	Datetime time.Time
}

func NewReciptModel(name string, price uint32, datetime time.Time) *ReciptModel {
	return &ReciptModel{Name: name, Price: price, Datetime: datetime}
}
