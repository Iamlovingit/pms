package main

import (
	"context"
	"fmt"
)

type Records []Record

type Record struct {
	Name       string  `json:"name" db:"name"`
	Department string  `json:"dep" db:"dep"`
	ID         string  `json:"id" db:"di"`
	BankID     string  `json:"bankid" db:"banid"`
	January    float32 `json:"january" db:"januray"`
	February   float32 `json:"february" db:"february"`
	March      float32 `json:"march" db:"march"`
	April      float32 `json:"april" db:"april"`
	May        float32 `json:"may" db:"may"`
	June       float32 `json:"june" db:"june"`
	July       float32 `json:"july" db:"july"`
	August     float32 `json:"august" db:"august"`
	September  float32 `json:"september" db:"september"`
	October    float32 `json:"october" db:"october"`
	November   float32 `json:"november" db:"november"`
	December   float32 `json:"december" db:"december"`
	Average    float32 `json:"average" db:"average"`
	Total      float32 `json:"total" db:"total"`
	Bonus      float32 `json:"bonus" db:"bonus"`
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

//
func (a *App) GetRecords() Records {
	var rs Records
	r := Record{
		Name: "cqs",
	}
	rs = append(rs, r)
	return rs
}
