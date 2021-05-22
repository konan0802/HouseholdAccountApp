# MillionaireApp（SpreadSheet家計簿アプリ）

"MillionaireApp" is a household account book application that uses google spreadsheet.

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
    ├── app.yaml
    ├── go.mod
    ├── go.sum
    ├── credential.json
    ├── config
    │   └── config.go
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

## TODO
* Front
    * React
* Update README
    * about MillionaireApp
        * spreadsheet
        * GAE
    * How to start
    * UI image