package config

import (
	"time"
)

var Now time.Time

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	Now = time.Now().In(loc)
}

// CategorieIDtoName CategorieIDとカテゴリ名の対応
var CategorieIDtoName = map[int]string{
	0: "食費",
	1: "生活用品",
	2: "娯楽",
	3: "水道代",
	4: "家具家電",
}

// CategorieNametoInt CategorieIDとカテゴリ名の対応
var CategorieNametoInt = map[string]int{
	"食費":   0,
	"生活用品": 1,
	"娯楽":   2,
	"水道代":  3,
	"家具家電": 4,
}

// NowYM 対象月のシート名
//var NowYM = strconv.Itoa(Now.Year()) + "年" + Now.Month().String() + "月"
var NowYM = "2022年1月"

// CheckReciptColumnsRange レシートの挿入先チェック
var CheckReciptColumnsRange = "!B7:B1500"

// レシートの開始・終了カラム
var ReceiptStartCol = "B"
var ReceiptEndCol = "E"

// ReciptStartRow レシートの記載が始まる行
var ReciptStartRow = 7

// GetReciptRange 対象レシートの取得
var GetReciptRange = "!B7:E1000"

// BudgetCategorytoCol 予算の絡む番号
var BudgetCategorytoCol = map[string]string{
	"食費":   "B",
	"生活用品": "C",
	"娯楽":   "D",
	"水道代":  "E",
	"家具家電": "F",
}

// BugetStartCol 予実が始まる行
var BugetStartCol = "2"

// BugetEndCol 予実が終わる行
var BugetEndCol = "4"
