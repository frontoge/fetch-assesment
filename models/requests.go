package models

import (
	"time"
)

type ItemRequest struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ProcessReceiptRequest struct {
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Items        []ItemRequest `json:"items"`
	Total        string        `json:"total"`
}

type Item struct {
	ShortDescription string
	Price            float64
}

type MappedProcessReceiptRequest struct {
	Retailer          string
	PurchaseTimestamp time.Time
	Items             []Item
	Total             float64
}
