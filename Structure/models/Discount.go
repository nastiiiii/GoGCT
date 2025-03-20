package models

import (
	"encoding/json"
	"time"
)

type DISCOUNT_TYPE_ENUM string

const (
	Date_Based    DISCOUNT_TYPE_ENUM = "Date-Based"
	Bulk_Based    DISCOUNT_TYPE_ENUM = "Bulk-Based"
	Special_Offer DISCOUNT_TYPE_ENUM = "Special-Offer"
)

type Discount struct {
	DiscountId          int                `json:"discount_id"`
	DiscountName        string             `json:"discount_name"`
	DiscountValue       float64            `json:"discount_value"`
	DiscountType        DISCOUNT_TYPE_ENUM `json:"discount_type"`
	IsRecurring         bool               `json:"is_recurring"`
	IsActive            bool               `json:"is_active"`
	AppliesToSocialClub bool               `json:"applies_to_social_club"`
	CustomLogic         string             `json:"custom_logic"`
	DiscountCode        string             `json:"discount_code"`
}

type BulkBasedRule struct {
	MinBookings int     `json:"min_bookings"`
	MaxBookings int     `json:"max_bookings"`
	Percentage  float64 `json:"percentage"`
}

type SpecialOfferRule struct {
	RuleName           string  `json:"rule_name"`
	BuyQuantity        int     `json:"buy_quantity"`
	GetQuantity        int     `json:"get_quantity"`
	DiscountPercentage float64 `json:"discount_percentage"`
	Active             bool    `json:"active"`
}

type DateBasedRule struct {
	Begins time.Time `json:"begins"`
	Ends   time.Time `json:"ends"`
}

func (d *Discount) ToCustomLogic() (map[string]interface{}, error) {
	var customLogic map[string]interface{}
	if err := json.Unmarshal([]byte(d.CustomLogic), &customLogic); err != nil {
		return nil, err
	}
	return customLogic, nil
}
