package models

import common_models "cfotech/common/models"

type Order struct {
	OrderID string                    `json:"order_id"`
	Status  common_models.OrderStatus `json:"status"`
}
