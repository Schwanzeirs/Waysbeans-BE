package cartsdto

import (
	"waysbeans/models"
)

type CartResponse struct {
	ProductID     int            `json:"product_id"`
	Product       models.Product `json:"product"`
	TransactionID int            `json:"transaction_id"`
	Qty           int            `json:"qty"`
	SubAmount     int            `json:"sub_amount"`
}
