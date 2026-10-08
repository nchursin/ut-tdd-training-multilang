# -*- coding: utf-8 -*-
from __future__ import print_function

from datetime import datetime
from gilded_rose import *


def main():
    print("OMGHAI!")
    items = [
        Item(name="+5 Dexterity Vest", sell_in=10, quality=20),
        Item(name="Aged Brie", sell_in=2, quality=0),
        Item(name="Elixir of the Mongoose", sell_in=5, quality=7),
        Item(name="Sulfuras, Hand of Ragnaros", sell_in=0, quality=80),
        Item(name="Sulfuras, Hand of Ragnaros", sell_in=-1, quality=80),
        Item(name="Backstage passes to a TAFKAL80ETC concert", sell_in=15, quality=20),
        Item(name="Backstage passes to a TAFKAL80ETC concert", sell_in=10, quality=49),
        Item(name="Backstage passes to a TAFKAL80ETC concert", sell_in=5, quality=49),
        Item(name="Conjured Mana Cake", sell_in=3, quality=6),  # <-- :O
    ]
    receipts = [
        Receipt(item_name="+5 Dexterity Vest", customer_name="John Doe",
                purchase_datetime=datetime(2026, 6, 15, 10, 30)),
        Receipt(item_name="Elixir of the Mongoose", customer_name="Jane Smith",
                purchase_datetime=datetime(2026, 6, 28, 19, 0)),
        Receipt(item_name="Aged Brie", customer_name="Bob Goodfellow",
                purchase_datetime=datetime(2026, 3, 1, 14, 0)),
        Receipt(item_name="Conjured Mana Cake", customer_name="Alice Wonder",
                purchase_datetime=datetime(2026, 6, 7, 9, 0)),
    ]
    days = 2
    import sys
    if len(sys.argv) > 1:
        days = int(sys.argv[1]) + 1
    for day in range(days):
        print("-------- day %s --------" % day)
        print("name, sellIn, quality")
        for item in items:
            print(item)
        print("")
        GildedRose(items, receipts).update_quality()
    generate_receipt_report(receipts)


if __name__ == "__main__":
    main()
