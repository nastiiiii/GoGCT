package models

import "time"

//@todo needs to implement better logic for the discounts
//@todo Discounts to Rules when Database up and working

type DiscountRule interface {
	ApplyDiscount() float64
	IsApplicable() bool
}

type BaseDiscountRule struct {
	IsActive  bool
	StartDate time.Time
	EndDate   time.Time
}

func (b *BaseDiscountRule) IsApplicable() bool {
	now := time.Now()
	return b.IsActive && now.After(b.StartDate) && now.Before(b.EndDate)
}

type BulkBasedRule struct {
	BaseDiscountRule
	MinBooking int
	MaxBooking int
	Percentage float64
}

// ApplyDiscount implements DiscountRule interface
func (b *BulkBasedRule) ApplyDiscount() float64 {
	if b.IsApplicable() {
		return b.Percentage // Apply discount percentage
	}
	return 0.0
}

// DateBasedRule applies discount if within a valid date range
type DateBasedRule struct {
	BaseDiscountRule
}

// ApplyDiscount implements DiscountRule interface
func (d *DateBasedRule) ApplyDiscount() float64 {
	if d.IsApplicable() {
		return 10.0 // Assume a fixed discount of 10 units
	}
	return 0.0
}

// SpecialOfferRule gives discounts for buying a specific quantity
type SpecialOfferRule struct {
	BaseDiscountRule
	BuyQuantity        int
	RuleName           string
	GetQuantity        int
	DiscountPercentage float64
}

// ApplyDiscount implements DiscountRule interface
func (s *SpecialOfferRule) ApplyDiscount() float64 {
	if s.IsApplicable() {
		return s.DiscountPercentage
	}
	return 0.0
}

type DiscountManager struct {
	DiscountRules []DiscountRule
}

// ApplyDiscount calculates the total discount based on applicable rules
func (dm *DiscountManager) ApplyDiscount() float64 {
	totalDiscount := 0.0
	for _, rule := range dm.DiscountRules {
		if rule.IsApplicable() {
			totalDiscount += rule.ApplyDiscount()
		}
	}
	return totalDiscount
}

// GetApplicableDiscounts filters and returns applicable discount rules
func (dm *DiscountManager) GetApplicableDiscounts() []DiscountRule {
	var applicableRules []DiscountRule
	for _, rule := range dm.DiscountRules {
		if rule.IsApplicable() {
			applicableRules = append(applicableRules, rule)
		}
	}
	return applicableRules
}
