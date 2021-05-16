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

// ReceiptHandler Receipt における Handler のインターフェース
type ReceiptHandler interface {
	GetMonthlyReceipts(http.ResponseWriter, *http.Request, httprouter.Params)
	Create(w http.ResponseWriter, r *http.Request, pr httprouter.Params)
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

// Monthlyreceipts GET /receipts -> Receipt データ一覧を返す
func (bh receiptHandler) GetMonthlyReceipts(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	type GetMonthlyReceiptsResponse struct {
		Receipts []model.ReceiptModel `json:"receipts"`
	}

	// リクエストパラメーターの取得
	// エラーはここでは握り潰し、バリテーションに委託する
	// （エラーの場合は0が代入される）
	_ = r.ParseForm()
	yearStr := r.Form.Get("year")
	yearUint64, _ := strconv.ParseUint(yearStr, 10, 64)
	yearUint := uint16(yearUint64)
	monthStr := r.Form.Get("month")
	monthUint64, _ := strconv.ParseUint(monthStr, 10, 64)
	monthUint := uint8(monthUint64)

	// リクエストstructに埋め込み
	mrr := model.MonthlyReceiptsRequest{
		Year:  yearUint,
		Month: monthUint,
	}

	// バリデーションを実行
	validate = validator.New()
	err := validate.Struct(mrr)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error:Struct", 500)
		return
	}

	// ユースケースの呼出
	receipts, err := bh.receiptUseCase.GetMonthlyReceipts(mrr)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error:GetMonthlyReceipts", 500)
		return
	}

	// 取得したドメインモデルを response に変換
	var res GetMonthlyReceiptsResponse
	for _, receipt := range receipts {
		rm := model.ReceiptModel(*receipt)
		res.Receipts = append(res.Receipts, rm)
	}

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error:Encode", 500)
		return
	}
}

// Create GET /receipts -> Receipt データ一覧を返す
func (bh receiptHandler) Create(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	ctx := r.Context()

	// ユースケースの呼出
	err := bh.receiptUseCase.AddReceipt(ctx)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error:AddReceipt", 500)
		return
	}

	successCode := map[string]string{"status": "success"}

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(successCode); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error:Encode", 500)
		return
	}
}
