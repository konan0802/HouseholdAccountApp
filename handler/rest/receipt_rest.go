package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"MillionaireApp/domain/model"
	"MillionaireApp/usecase"

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
func (rh receiptHandler) GetMonthlyReceipts(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	type GetMonthlyReceiptsResponse struct {
		Receipts []model.ReceiptModel `json:"receipts"`
	}

	// リクエストパラメーターの取得
	// エラーはここでは握り潰し、バリテーションに委託する
	// （エラーの場合は0が代入される）
	_ = r.ParseForm()
	yearStr := r.Form.Get("year")
	yearUint, err := strconv.Atoi(yearStr)
	if err != nil {
		log.Fatal(err)
	}
	monthStr := r.Form.Get("month")
	monthUint, err := strconv.Atoi(monthStr)
	if err != nil {
		log.Fatal(err)
	}

	// リクエストstructに埋め込み
	mrr := model.MonthlyReceiptsRequest{
		Year:  yearUint,
		Month: monthUint,
	}

	// バリデーションを実行
	validate = validator.New()
	err = validate.Struct(mrr)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error:Struct", 500)
		return
	}

	// ユースケースの呼出
	receipts, err := rh.receiptUseCase.GetMonthlyReceipts(mrr)
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
func (rh receiptHandler) Create(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	type CreateResponse struct {
		Receipt model.ReceiptModel `json:"receipt"`
	}

	// リクエストパラメーターの取得
	// エラーはここでは握り潰し、バリテーションに委託する
	// （エラーの場合は0が代入される）
	_ = r.ParseForm()
	categorieName := r.Form.Get("categorie_name")
	description := r.Form.Get("description")
	priceStr := r.Form.Get("price")
	priceUint32, err := strconv.Atoi(priceStr)
	if err != nil {
		log.Fatal(err)
	}
	datetime := r.Form.Get("datetime")

	// リクエストstructに埋め込み
	arr := model.AddReceiptRequest{
		CategorieName: categorieName,
		Description:   description,
		Price:         priceUint32,
		Datetime:      datetime,
	}

	// バリデーションを実行
	validate = validator.New()
	err = validate.Struct(arr)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error: validate.Struct", 500)
		return
	}

	// ユースケースの呼出
	rm, err := rh.receiptUseCase.AddReceipt(arr)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error:AddReceipt", 500)
		return
	}

	var res CreateResponse
	res.Receipt = *rm

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error:Encode", 500)
		return
	}
}
