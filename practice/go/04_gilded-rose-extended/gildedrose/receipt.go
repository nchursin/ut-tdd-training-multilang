package gildedrose

import (
	"fmt"
	"time"
)

const (
	canBeReturned    = "can_be_returned"
	cannotBeReturned = "cannot_be_returned"
	archived         = "archived"
)

type Receipt struct {
	ItemName         string
	CustomerName     string
	PurchaseDatetime time.Time
	ReturnDeadline   time.Time
	Status           string
}

func NewReceipt(item *Item, customer string, purchase time.Time) *Receipt {
	deadlineDays := 14
	if purchase.Weekday() == time.Saturday {
		deadlineDays = 16
	} else if purchase.Weekday() == time.Sunday {
		deadlineDays = 15
	} else if purchase.Hour() >= 18 {
		deadlineDays++
	}
	deadline := purchase.AddDate(0, 0, deadlineDays)
	canReturn := true
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		if item.SellIn-deadlineDays <= 2 {
			canReturn = false
		}
		if item.Name != "Aged Brie" && item.Quality-deadlineDays <= 2 {
			canReturn = false
		}
	}
	if !canReturn {
		deadline = purchase
	}
	receipt := &Receipt{ItemName: item.Name, CustomerName: customer, PurchaseDatetime: purchase, ReturnDeadline: deadline}
	receipt.Status = receiptStatus(receipt, time.Now())
	return receipt
}

func receiptStatus(receipt *Receipt, now time.Time) string {
	if !now.After(receipt.ReturnDeadline) {
		return canBeReturned
	}
	if !now.After(receipt.ReturnDeadline.AddDate(0, 0, 30)) {
		return cannotBeReturned
	}
	return archived
}

func UpdateReceipts(receipts []*Receipt) {
	for _, receipt := range receipts {
		receipt.Status = receiptStatus(receipt, time.Now())
	}
}

func GenerateReceiptReport(receipts []*Receipt) string {
	canBeReturnedCount, cannotBeReturnedCount, archivedCount := 0, 0, 0
	for _, receipt := range receipts {
		switch receipt.Status {
		case canBeReturned:
			canBeReturnedCount++
		case cannotBeReturned:
			cannotBeReturnedCount++
		case archived:
			archivedCount++
		}
	}
	report := fmt.Sprintf("===== Receipt Report [%s] =====\n", time.Now().Format("2006-01-02 15:04:05"))
	report += fmt.Sprintf("can_be_returned: %d\n", canBeReturnedCount)
	report += fmt.Sprintf("cannot_be_returned: %d\n", cannotBeReturnedCount)
	report += fmt.Sprintf("archived: %d\n\n", archivedCount)
	for _, receipt := range receipts {
		if receipt.Status != archived {
			report += receipt.String() + "\n"
		}
	}
	return report
}

func (r *Receipt) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", r.ItemName, r.CustomerName, r.ReturnDeadline.Format("2006-01-02 15:04"), r.Status)
}
