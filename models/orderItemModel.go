package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	Quantity      int                `json:"quantity" validate:"required, eq=S|eq=M|eq=L"`
	Unit_price    int                `json:"unit_price" validate:"required"`
	Food_id       *string            `json:"food_id" validate:"required"`
	Order_item_id string             `json:"order_item_id"`
	Order_id      string             `json:"order_id" validate:"required"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
}
