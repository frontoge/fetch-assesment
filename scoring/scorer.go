package scoring

import (
	"fetch/receipt-processor/models"
	"math"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
)

var savedReceipts map[string]int = make(map[string]int)

// Calculates the number of alphanumeric characters in the name
func getAlphanumericCount(name string) int {
	count := 0

	for _, c := range name {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			count += 1
		}
	}

	return count
}

func getPointsFromItems(items []models.Item) int {
	// Assign points for each pair of items
	score := (len(items) / 2) * 5

	// Assign points for each item that has a description with multiple of 3 characters
	for _, item := range items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			score += int(math.Ceil(item.Price * 0.2))
		}
	}

	return score
}

// Gets the points awarded for the receipt total
func getPointsFromTotalPrice(total float64) int {
	score := 0
	if total == math.Trunc(total) {
		score += 50
	}

	if (total / 0.25) == math.Trunc(total/0.25) {
		score += 25
	}

	return score
}

// Gets the points awarded for date and time of purchase
func getPointsFromDate(date time.Time) int {
	score := 0

	hour := date.Hour()

	// If day is odd
	if date.Day()%2 == 1 {
		score += 6
	}

	// If time is between 2pm and 4pm
	if hour >= 14 && hour < 16 {
		score += 10
	}

	return score
}

func ProcessReceipt(data models.MappedProcessReceiptRequest) string {
	// Generate a UUID for this receipt
	id := uuid.New().String()

	// Calculate the score
	score := getAlphanumericCount(data.Retailer)
	score += getPointsFromItems(data.Items)
	score += getPointsFromTotalPrice(data.Total)
	score += getPointsFromDate(data.PurchaseTimestamp)

	// Save the receipt
	savedReceipts[id] = score

	return id
}

func GetScoreById(id string) int {
	value, exists := savedReceipts[id]

	if exists {
		return value
	} else {
		return -1
	}
}
