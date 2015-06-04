package quan

import (
	"time"
)

type Orders struct {
	Id              int64     `json:"id"`
	Code            string    `json:"code"`
	Uid             int64     `json:"uid"`
	Keywords        string    `json:"keywords"`
	Name            string    `json:"name"`
	Phone           string    `json:"phone"`
	Postal          string    `json:"postal"`
	Province        string    `json:"province"`
	City            string    `json:"city"`
	District        string    `json:"district"`
	Addr            string    `json:"addr"`
	PaymentStatus   int       `json:"payment_status"`
	LogisticsStatus int       `json:"logistics_status"`
	Freight         float64   `json:"freight"`
	FreightEx       float64   `json:"freight_ex"`
	Total           float64   `json:"total"`
	TotalEx         float64   `json:"total_ex"`
	TotalPay        float64   `json:"total_pay"`
	Discount        string    `json:"discount"`
	Address         string    `json:"address"`
	PaymentId       string    `json:"payment_id"`
	Intro           string    `json:"intro"`
	DeletedAt       time.Time `xorm:"deleted" json:"deleted_at"`
	CreatedAt       time.Time `xorm:"created" json:"created_at"`
	UpdatedAt       time.Time `xorm:"updated" json:"updated_at"`
}

type OrdersService struct {
	OrdersGetInfoById func(id int64) (*Orders, error)
}

var OrdersClient *OrdersService
