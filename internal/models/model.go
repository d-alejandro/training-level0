package models

import (
	"encoding/json"
	"errors"
)

type Model struct {
	OrderUID    string `json:"order_uid" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required"`
	Entry       string `json:"entry" validate:"required"`
	Delivery    struct {
		Name    string `json:"name" validate:"required"`
		Phone   string `json:"phone" validate:"required,e164"`
		Zip     string `json:"zip" validate:"required"`
		City    string `json:"city" validate:"required"`
		Address string `json:"address" validate:"required"`
		Region  string `json:"region" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
	} `json:"delivery" validate:"required"`
	Payment struct {
		Transaction  string `json:"transaction" validate:"required"`
		RequestID    string `json:"request_id" validate:"omitempty,min=1"`
		Currency     string `json:"currency" validate:"required,iso4217"`
		Provider     string `json:"provider" validate:"required"`
		Amount       int    `json:"amount" validate:"required,numeric,min=0"`
		PaymentDT    int    `json:"payment_dt" validate:"required,numeric,min=946684800"`
		Bank         string `json:"bank" validate:"required"`
		DeliveryCost int    `json:"delivery_cost" validate:"required,numeric,min=0"`
		GoodsTotal   int    `json:"goods_total" validate:"required,numeric,min=0"`
		CustomFee    int    `json:"custom_fee" validate:"required|eq=0,numeric"`
	} `json:"payment" validate:"required"`
	Items []struct {
		ChrtID      int    `json:"chrt_id" validate:"required,numeric,min=0"`
		TrackNumber string `json:"track_number" validate:"required"`
		Price       int    `json:"price" validate:"required,numeric,min=0"`
		Rid         string `json:"rid" validate:"required"`
		Name        string `json:"name" validate:"required"`
		Sale        int    `json:"sale" validate:"required,numeric,min=0"`
		Size        string `json:"size" validate:"required"`
		TotalPrice  int    `json:"total_price" validate:"required,numeric,min=0"`
		NmID        int    `json:"nm_id" validate:"required,numeric,min=0"`
		Brand       string `json:"brand" validate:"required"`
		Status      int    `json:"status" validate:"required,numeric,min=0"`
	} `json:"items" validate:"dive"`
	Locale            string `json:"locale" validate:"required,min=2,bcp47_language_tag"`
	InternalSignature string `json:"internal_signature" validate:"omitempty,min=1"`
	CustomerID        string `json:"customer_id" validate:"required"`
	DeliveryService   string `json:"delivery_service" validate:"required"`
	Shardkey          string `json:"shardkey" validate:"required"`
	SmID              int    `json:"sm_id" validate:"required,numeric,min=0"`
	DateCreated       string `json:"date_created" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	OofShard          string `json:"oof_shard" validate:"required"`
}

func (model *Model) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &model)
}
