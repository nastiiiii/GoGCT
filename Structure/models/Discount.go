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

// Discount represents entity from the database
type Discount struct {
	DiscountId          int                `json:"discount_id"`
	DiscountName        string             `json:"discount_name"`
	DiscountValue       float64            `json:"discount_value"`
	DiscountType        DISCOUNT_TYPE_ENUM `json:"discount_type"`
	IsRecurring         bool               `json:"is_recurring"`
	IsActive            bool               `json:"is_active"`
	AppliesToSocialClub bool               `json:"applies_to_social_club"`
	CustomLogic         string             `json:"custom_logic"` //keeping to what type it belongs example: {"Bulk-Based Rules":[{"MaxBookings":20,"MinBookings":10,"Percentage":0.05}]}
	DiscountCode        string             `json:"discount_code"`
}

// BulkBasedRule bulk based discount
type BulkBasedRule struct {
	MinBookings int     `json:"min_bookings"`
	MaxBookings int     `json:"max_bookings"`
	Percentage  float64 `json:"percentage"`
}

// SpecialOfferRule special offer discount
type SpecialOfferRule struct {
	RuleName           string  `json:"rule_name"`
	BuyQuantity        int     `json:"buy_quantity"`
	GetQuantity        int     `json:"get_quantity"`
	DiscountPercentage float64 `json:"discount_percentage"`
	Active             bool    `json:"active"`
}

// DateBasedRule date based discount
type DateBasedRule struct {
	Begins time.Time `json:"begins"`
	Ends   time.Time `json:"ends"`
}

// ToCustomLogic in the Discount struct is responsible for parsing the CustomLogic field, which is stored as a JSON string in the database, into a Go map[string]interface{}.
func (d *Discount) ToCustomLogic() (map[string]interface{}, error) {
	var customLogic map[string]interface{}
	if err := json.Unmarshal([]byte(d.CustomLogic), &customLogic); err != nil {
		return nil, err
	}
	return customLogic, nil
}
