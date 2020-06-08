package DTOs

import (
	"database/sql"
	"time"
)

type CreateContractInput struct {
	ContractNumber string `json:"contract_number" binding:"required"`
	CustomerName string `json:"customer_name" binding:"required"`
	StartDate *time.Time `json:"start_date" binding:"required"`
	EndDate *time.Time `json:"end_date" binding:"required"`
	Balance sql.NullInt64 `json:"balance"`
}