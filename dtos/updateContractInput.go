package DTOs

import (
	"database/sql"
	"time"
)

type UpdateContractInput struct {
	ContractNumber string `json:"contract_number"`
	CustomerName string `json:"customer_name"`
	StartDate *time.Time `json:"start_date"`
	EndDate *time.Time `json:"end_date"`
	Balance sql.NullInt64 `json:"balance"`
}