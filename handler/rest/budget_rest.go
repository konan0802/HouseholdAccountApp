package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"HouseholdAccountApp/domain/model"
	"HouseholdAccountApp/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

// BudgetHandler Budget における Handler のインターフェース
type BudgetHandler interface {
	GetMonthlybudget(http.ResponseWriter, *http.Request, httprouter.Params)
}

type budgetHandler struct {
	budgetUseCase usecase.BudgetUseCase
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// NewBudgetUseCase Budget データに関する Handler を生成
func NewBudgetHandler(bu usecase.BudgetUseCase) BudgetHandler {
	return &budgetHandler{
		budgetUseCase: bu,
	}
}

// Monthlybudget GET /budgets -> Budget データ一覧を返す
func (bh budgetHandler) GetMonthlybudget(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	type GetMonthlybudgetResponse struct {
		Budgets []model.BudgetModel `json:"budgets"`
	}

	// リクエストパラメーターの取得
	// エラーはここでは握り潰し、バリテーションに委託する
	// （エラーの場合は0が代入される）
	_ = r.ParseForm()
	yearStr := r.Form.Get("year")
	yearUint64, _ := strconv.ParseUint(yearStr, 10, 64)
	yearUint := uint32(yearUint64)
	monthStr := r.Form.Get("month")
	monthUint64, _ := strconv.ParseUint(monthStr, 10, 64)
	monthUint := uint32(monthUint64)

	// リクエストstructに埋め込み
	mbr := model.MonthlyBudgetsRequest{
		Year:  yearUint,
		Month: monthUint,
	}

	// バリデーションを実行
	validate = validator.New()
	err := validate.Struct(mbr)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error: validate.Struct", 500)
		return
	}

	// ユースケースの呼出
	budgets, err := bh.budgetUseCase.GetMonthlyBudgets(mbr)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error: GetMonthlyBudgets", 500)
		return
	}

	// 取得したドメインモデルを response に変換
	var res GetMonthlybudgetResponse
	for _, budget := range budgets {
		bm := model.BudgetModel(*budget)
		res.Budgets = append(res.Budgets, bm)
	}

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error: Encode", 500)
		return
	}
}
