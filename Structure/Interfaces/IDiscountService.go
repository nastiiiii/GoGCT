package Interfaces

import "GCT/Structure/models"

type IDiscountService interface {
	LoadDiscounts() ([]models.Discount, error)
	SaveDiscount(d models.Discount) error
	DeleteDiscount(id int) error
	ApplyBestDiscount(transaction *models.Transaction)
	isDiscountApplicable(discount models.Discount, transaction *models.Transaction) bool
}
