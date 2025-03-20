package Services

import (
	"GCT/Structure/models"
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"time"
)

type DiscountService struct {
	DB *pgx.Conn
}

func (ds *DiscountService) LoadDiscounts() ([]models.Discount, error) {
	rows, err := ds.DB.Query(context.Background(), `SELECT "discountID", "discountName", "discountValue", "discountType", "isRecurring", "isActive", "appliesToSocialClub", "customLogic", "discountCode" FROM "Discounts"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var discounts []models.Discount
	for rows.Next() {
		var d models.Discount
		if err := rows.Scan(&d.DiscountId, &d.DiscountName, &d.DiscountValue, &d.DiscountType, &d.IsRecurring, &d.IsActive, &d.AppliesToSocialClub, &d.CustomLogic, &d.DiscountCode); err != nil {
			return nil, err
		}
		discounts = append(discounts, d)
	}
	return discounts, nil
}

func (ds *DiscountService) SaveDiscount(d models.Discount) error {
	_, err := ds.DB.Exec(context.Background(), `INSERT INTO "Discounts" ("discountName", "discountValue", "discountType", "isRecurring", "isActive", "appliesToSocialClub", "customLogic", "discountCode") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		d.DiscountName, d.DiscountValue, d.DiscountType, d.IsRecurring, d.IsActive, d.AppliesToSocialClub, d.CustomLogic, d.DiscountCode)
	return err
}

func (ds *DiscountService) DeleteDiscount(id int) error {
	query := `DELETE FROM "Discounts" WHERE "discountID" = $1`
	_, err := ds.DB.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DiscountService) ApplyBestDiscount(transaction *models.Transaction) error {
	discounts, err := ds.LoadDiscounts()
	if err != nil {
		return err
	}

	var bestDiscountValue float64

	for _, discount := range discounts {
		if !discount.IsActive {
			continue
		}

		// Check if the discount can be applied
		if ds.IsDiscountApplicable(discount, transaction) {
			if discount.DiscountValue > bestDiscountValue {
				bestDiscountValue = discount.DiscountValue
			}
		}
	}

	// Apply the best discount
	transaction.TotalCost -= bestDiscountValue

	return nil
}

func (ds *DiscountService) IsDiscountApplicable(discount models.Discount, transaction *models.Transaction) bool {
	// Check if discount applies only to Social Club members
	if discount.AppliesToSocialClub {
		/*if !ds.isUserSocialClubMember(transaction.AccountId) {
			return false
		}*/
	}

	// Handle Date-Based Discounts
	if discount.DiscountType == models.Date_Based {
		var dateRule models.DateBasedRule
		if err := json.Unmarshal([]byte(discount.CustomLogic), &dateRule); err == nil {
			currentTime := time.Now()
			if currentTime.After(dateRule.Begins) && currentTime.Before(dateRule.Ends) {
				return true
			}
		}
	}

	// Handle Bulk-Based Discounts
	if discount.DiscountType == models.Bulk_Based {
		var bulkRule models.BulkBasedRule
		if err := json.Unmarshal([]byte(discount.CustomLogic), &bulkRule); err == nil {
			if transaction.TotalCost >= float64(bulkRule.MinBookings) && transaction.TotalCost <= float64(bulkRule.MaxBookings) {
				return true
			}
		}
	}

	// Handle Special-Offer Discounts
	if discount.DiscountType == models.Special_Offer {
		var offerRule models.SpecialOfferRule
		if err := json.Unmarshal([]byte(discount.CustomLogic), &offerRule); err == nil {
			if offerRule.Active {
				return true
			}
		}
	}

	return false
}
