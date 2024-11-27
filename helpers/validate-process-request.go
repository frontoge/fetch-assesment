package helpers

import (
	"errors"
	"fetch/receipt-processor/models"
)

func ValidateProcessRequest(request models.ProcessReceiptRequest) error {
	if request.Retailer == "" {
		return errors.New("Missing Retailer")
	}

	if request.PurchaseDate == "" {
		return errors.New("Missing Purchase date")
	}

	if request.PurchaseTime == "" {
		return errors.New("Missing Purchase time")
	}

	if request.Total == "" {
		return errors.New("Missing Total")
	}

	if request.Items == nil {
		return errors.New("Missing Items list")
	}

	return nil
}
