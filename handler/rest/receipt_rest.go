package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"HouseholdAccountApp/usecase"

	"github.com/julienschmidt/httprouter"
)

// ReceiptHandler Receipt における Handler のインターフェース
type ReceiptHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
}

type receiptHandler struct {
	receiptUseCase usecase.ReceiptUseCase
}

// NewReceiptUseCase Receipt データに関する Handler を生成
func NewReceiptHandler(bu usecase.ReceiptUseCase) ReceiptHandler {
	return &receiptHandler{
		receiptUseCase: bu,
	}
}

// ReceiptIndex GET /receipts -> Receipt データ一覧を返す
func (bh receiptHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	// receiptField response 内で使用する Receipt を表す構造体
	//  -> ドメインモデルの Receipt に HTTP の関心事である JSON タグを付与したくないために Handler 層で用意
	//     簡略化のために JSON タグを付与したドメインモデルを流用するプロジェクトもしばしば見かける
	type receiptField struct {
		ReceiptID   uint32    `json:"receiptID"`
		CategorieID uint32    `json:"categorieID"`
		ProductName string    `json:"productName"`
		Price       uint32    `json:"price"`
		Methods     string    `json:"methods"`
		Datetime    time.Time `json:"datetime"`
	}

	ctx := r.Context()

	// ユースケースの呼出
	receipt, err := bh.receiptUseCase.GetReceipt(ctx)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 取得したドメインモデルを response に変換
	rf := receiptField(*receipt)

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(rf); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
