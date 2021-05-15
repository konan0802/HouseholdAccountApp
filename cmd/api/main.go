package main

import (
	"fmt"
	"log"
	"net/http"

	handler "HouseholdAccountApp/handler/rest"
	"HouseholdAccountApp/infra/persistence"
	"HouseholdAccountApp/usecase"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// 依存関係を注入（DI まではいきませんが一応注入っぽいことをしてる）
	// DI ライブラリを使えば、もっとスマートになるはず
	budgetPersistence := persistence.NewBudgetPersistence()
	budgetUseCase := usecase.NewBudgetUseCase(budgetPersistence)
	budgetHandler := handler.NewBudgetHandler(budgetUseCase)
	receiptPersistence := persistence.NewReceiptPersistence()
	receiptUseCase := usecase.NewReceiptUseCase(receiptPersistence)
	receiptHandler := handler.NewReceiptHandler(receiptUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/v1/budget", budgetHandler.Index)
	router.GET("/api/v1/receipt", receiptHandler.Index)

	// サーバ起動
	fmt.Println("========================")
	fmt.Println("Server Start >> http://localhost:3000")
	fmt.Println("========================")
	log.Fatal(http.ListenAndServe(":3000", router))
}
