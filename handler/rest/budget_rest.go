package rest

import (
	"encoding/json"
	"net/http"

	"HouseholdAccountApp/usecase"

	"github.com/julienschmidt/httprouter"
)

// BudgetHandler Budget における Handler のインターフェース
type BudgetHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
}

type budgetHandler struct {
	budgetUseCase usecase.BudgetUseCase
}

// NewBudgetUseCase Budget データに関する Handler を生成
func NewBudgetHandler(bu usecase.BudgetUseCase) BudgetHandler {
	return &budgetHandler{
		budgetUseCase: bu,
	}
}

// BudgetIndex GET /budgets -> Budget データ一覧を返す
func (bh budgetHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	// budgetField response 内で使用する Budget を表す構造体
	//  -> ドメインモデルの Budget に HTTP の関心事である JSON タグを付与したくないために Handler 層で用意
	//     簡略化のために JSON タグを付与したドメインモデルを流用するプロジェクトもしばしば見かける
	type budgetField struct {
		CategorieID      uint32 `json:"categorieID"`
		BudgetValue      uint32 `json:"budgetValue"`
		PerformanceValue uint32 `json:"performanceValue"`
		Year             uint16 `json:"year"`
		Month            uint8  `json:"month"`
	}

	ctx := r.Context()

	// ユースケースの呼出
	budget, err := bh.budgetUseCase.GetBudget(ctx)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 取得したドメインモデルを response に変換
	bf := budgetField(*budget)

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(bf); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
