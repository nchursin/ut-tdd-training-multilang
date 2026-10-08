package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func main() {
	fmt.Println("OMGHAI!")

	var items = []*gildedrose.Item{
		{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20},
		{Name: "Aged Brie", SellIn: 2, Quality: 0},
		{Name: "Elixir of the Mongoose", SellIn: 5, Quality: 7},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 49},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 49},
		{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6}, // <-- :O
	}

	receipts := []*gildedrose.Receipt{
		gildedrose.NewReceipt(items[0], "John Doe", time.Date(2026, 6, 15, 10, 30, 0, 0, time.Local)),
		gildedrose.NewReceipt(items[2], "Jane Smith", time.Date(2026, 6, 28, 19, 0, 0, 0, time.Local)),
		gildedrose.NewReceipt(items[1], "Bob Goodfellow", time.Date(2026, 3, 1, 14, 0, 0, 0, time.Local)),
		gildedrose.NewReceipt(items[8], "Alice Wonder", time.Date(2026, 6, 7, 9, 0, 0, 0, time.Local)),
	}

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Println("Name, SellIn, Quality")
		for i := 0; i < len(items); i++ {
			fmt.Println(items[i])
		}
		fmt.Println("")
		gildedrose.UpdateQuality(items)
	}

	gildedrose.UpdateReceipts(receipts)
	fmt.Print(gildedrose.GenerateReceiptReport(receipts))
}
