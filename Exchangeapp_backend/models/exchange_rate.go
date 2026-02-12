package models

import "time"

type ExchangeRate struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	FromCurrency string    `json:"fromcurrency" binding:"required"`
	ToCurrency   string    `json:"tocurrency" binding:"required"`
	Rate         float64   `json:"rate" binding:"required"`
	Time         time.Time `json:"time"`
}
