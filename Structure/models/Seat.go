package models

// Seats is contained in performance models, not instance in database
type Seats struct {
	Seat         int
	Availability bool
}
