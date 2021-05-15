package main

import (
	"log"
	"net/http"
	"os"

	handler "HouseholdAccountApp/handler/rest"
	"HouseholdAccountApp/infra/persistence"
	"HouseholdAccountApp/usecase"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// 依存関係を注入
	budgetPersistence := persistence.NewBudgetPersistence()
	budgetUseCase := usecase.NewBudgetUseCase(budgetPersistence)
	budgetHandler := handler.NewBudgetHandler(budgetUseCase)
	receiptPersistence := persistence.NewReceiptPersistenceMock()
	receiptUseCase := usecase.NewReceiptUseCase(receiptPersistence)
	receiptHandler := handler.NewReceiptHandler(receiptUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/v1/budget", budgetHandler.Index)
	router.GET("/api/v1/receipt", receiptHandler.Index)

	// サーバ起動
	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
