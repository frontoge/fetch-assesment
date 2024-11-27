package helpers

import (
	"fetch/receipt-processor/models"
	"log"
	"strconv"
	"time"
)

// Handle data type mapping for items list
func mapItems(items []models.ItemRequest) []models.Item {
	var mappedItems []models.Item

	for _, item := range items {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			log.Fatalf("Error parsing item price: %v", err)
		}

		mappedItems = append(mappedItems, models.Item{
			ShortDescription: item.ShortDescription,
			Price:            price,
		})
	}

	return mappedItems
}

// Handle data type mapping for process receipt request
func MapProcessRequest(request models.ProcessReceiptRequest) (models.MappedProcessReceiptRequest, error) {
	date, dateErr := time.Parse("2006-01-02 15:04", request.PurchaseDate+" "+request.PurchaseTime)
	if dateErr != nil {
		return models.MappedProcessReceiptRequest{}, dateErr
	}

	total, totalErr := strconv.ParseFloat(request.Total, 64)

	if totalErr != nil {
		log.Default().Printf("Error parsing total: %v", totalErr)
		return models.MappedProcessReceiptRequest{}, totalErr
	}

	return models.MappedProcessReceiptRequest{
		Retailer:          request.Retailer,
		PurchaseTimestamp: date,
		Items:             mapItems(request.Items),
		Total:             total,
	}, nil
}
