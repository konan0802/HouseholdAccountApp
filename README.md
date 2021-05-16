# HouseholdAccountApp（SpreadSheet家計簿アプリ）

"HouseholdAccountApp" is a household account book application that uses google spreadsheet.

## Requirements
- Google App Engine (Server)
    - Golang go1.15
    - React
- Google SpreadSheet (DB)

## Layered Architecture + DDD
    [handler]
    　　↓
    [usecase]
    　　↓
    [domain]
    　　↑
    [infra]

## Directory structure
    ├── main.go
    ├── handler
    │   └── rest
    │       ├── budget_rest.go
    │       └── receipt_rest.go
    ├── usecase
    │   ├── budget_usecase.go
    │   └── receipt_usecase.go
    ├── domain
    │   ├── model
    │   │   ├── budget_model.go
    │   │   └── receipt_model.go
    │   └── repository
    │       ├── budget_repository.go
    │       └── receipt_repository.go
    └── infra
        └── infra
            ├── budget_ss.go
            ├── receipt_ss.go
            ├── budget_mock.go
            └── receipt_mock.go

## API Specification
#### GET:receipt/monthlyreceipts

#### POST:receipt/create

#### GET:budget/monthlybudget
https://householdaccountapp.an.r.appspot.com//receipt/create?categorie_name=%E9%A3%9F%E8%B2%BB&description=%E8%B1%86%E8%85%90&price=150&datetime=2021/5/16

budget page