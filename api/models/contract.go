package models

import (
	"database/sql"
	"time"
)

type Contract struct {
	ID uint `json:"id" gorm:"primary_key"`
	ContractNumber string `json:"contract_number"`
	CustomerName string `json:"customer_name"`
	StartDate *time.Time `json:"start_date"`
	EndDate *time.Time `json:"end_date"`
	Balance sql.NullInt64 `json:"balance"`
}