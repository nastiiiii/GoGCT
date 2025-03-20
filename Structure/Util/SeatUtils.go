package Util

import (
	Models "GCT/Structure/models"
	"errors"
	"strconv"
	"strings"
)

// ParseSeatPrices Parse seat band prices from "A:100; B:200; C:300;" format
func ParseSeatPrices(data string) (map[string]float64, error) {
	prices := make(map[string]float64)
	parts := strings.Split(data, ";") // ["A:100", " B:200", " C:300"]

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		bandData := strings.Split(part, ":") // ["A", "100"]
		if len(bandData) != 2 {
			return nil, errors.New("invalid seat price format")
		}

		price, err := strconv.ParseFloat(strings.TrimSpace(bandData[1]), 64)
		if err != nil {
			return nil, err
		}

		prices[strings.TrimSpace(bandData[0])] = price
	}

	return prices, nil
}

// ParseSeatAvailability Parse seat availability from "A: 12-true, 13-false, 14-true; B: 15-true, 16-true, 17-true;"
func ParseSeatAvailability(data string) (map[string][]Models.Seats, error) {
	seatMap := make(map[string][]Models.Seats)
	sections := strings.Split(data, ";") // ["A: 12-true, 13-false, 14-true", " B: 15-true, 16-true, 17-true"]

	for _, section := range sections {
		section = strings.TrimSpace(section)
		if section == "" {
			continue
		}

		parts := strings.Split(section, ":") // ["A", " 12-true, 13-false, 14-true"]
		if len(parts) != 2 {
			return nil, errors.New("invalid seat availability format")
		}

		seatBand := strings.TrimSpace(parts[0])
		seats := strings.Split(parts[1], ",") // [" 12-true", " 13-false", " 14-true"]

		for _, seat := range seats {
			seat = strings.TrimSpace(seat)
			seatParts := strings.Split(seat, "-") // ["12", "true"]
			if len(seatParts) != 2 {
				return nil, errors.New("invalid seat entry format")
			}

			seatNumber, err := strconv.Atoi(strings.TrimSpace(seatParts[0]))
			if err != nil {
				return nil, err
			}

			availability, err := strconv.ParseBool(strings.TrimSpace(seatParts[1]))
			if err != nil {
				return nil, err
			}

			seatMap[seatBand] = append(seatMap[seatBand], Models.Seats{
				Seat:         seatNumber,
				Availability: availability,
			})
		}
	}

	return seatMap, nil
}
