# -*- coding: utf-8 -*-
from datetime import datetime, timedelta


class GildedRose(object):

    def __init__(self, items, receipts=None):
        self.items = items
        self.receipts = receipts if receipts is not None else []

    def update_quality(self):
        for item in self.items:
            if item.name != "Aged Brie" and item.name != "Backstage passes to a TAFKAL80ETC concert":
                if item.quality > 0:
                    if item.name != "Sulfuras, Hand of Ragnaros":
                        item.quality = item.quality - 1
            else:
                if item.quality < 50:
                    item.quality = item.quality + 1
                    if item.name == "Backstage passes to a TAFKAL80ETC concert":
                        if item.sell_in < 11:
                            if item.quality < 50:
                                item.quality = item.quality + 1
                        if item.sell_in < 6:
                            if item.quality < 50:
                                item.quality = item.quality + 1
            if item.name != "Sulfuras, Hand of Ragnaros":
                item.sell_in = item.sell_in - 1
            if item.sell_in < 0:
                if item.name != "Aged Brie":
                    if item.name != "Backstage passes to a TAFKAL80ETC concert":
                        if item.quality > 0:
                            if item.name != "Sulfuras, Hand of Ragnaros":
                                item.quality = item.quality - 1
                    else:
                        item.quality = item.quality - item.quality
                else:
                    if item.quality < 50:
                        item.quality = item.quality + 1

    def update_receipts(self):
        for receipt in self.receipts:
            now = datetime.now()
            if now <= receipt.return_deadline:
                receipt.status = "can_be_returned"
            else:
                if now <= receipt.return_deadline + timedelta(days=30):
                    receipt.status = "cannot_be_returned"
                else:
                    receipt.status = "archived"

    def generate_receipt_report(self):
        can_be_returned_count = 0
        cannot_be_returned_count = 0
        archived_count = 0
        for receipt in self.receipts:
            if receipt.status == "can_be_returned":
                can_be_returned_count = can_be_returned_count + 1
            if receipt.status == "cannot_be_returned":
                cannot_be_returned_count = cannot_be_returned_count + 1
            if receipt.status == "archived":
                archived_count = archived_count + 1
        report = "===== Receipt Report [%s] =====\n" % datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        report = report + "can_be_returned: %d\n" % can_be_returned_count
        report = report + "cannot_be_returned: %d\n" % cannot_be_returned_count
        report = report + "archived: %d\n" % archived_count
        report = report + "\n"
        for receipt in self.receipts:
            if receipt.status != "archived":
                report = report + str(receipt) + "\n"
        return report


class Item:
    def __init__(self, name, sell_in, quality):
        self.name = name
        self.sell_in = sell_in
        self.quality = quality

    def __repr__(self):
        return "%s, %s, %s" % (self.name, self.sell_in, self.quality)


class Receipt:
    def __init__(self, item, customer_name, purchase_datetime):
        self.item_name = item.name
        self.customer_name = customer_name
        self.purchase_datetime = purchase_datetime

        if purchase_datetime.weekday() == 5:
            deadline_days = 16
        else:
            if purchase_datetime.weekday() == 6:
                deadline_days = 15
            else:
                deadline_days = 14
        if purchase_datetime.weekday() < 5:
            if purchase_datetime.hour >= 18:
                deadline_days = deadline_days + 1

        self.return_deadline = purchase_datetime + timedelta(days=deadline_days)

        can_return = True
        if item.name != "Sulfuras, Hand of Ragnaros":
            projected_sell_in = item.sell_in - deadline_days
            if projected_sell_in <= 2:
                can_return = False
            if item.name != "Aged Brie":
                projected_quality = item.quality - deadline_days
                if projected_quality <= 2:
                    can_return = False

        if not can_return:
            self.return_deadline = purchase_datetime

        now = datetime.now()
        if now <= self.return_deadline:
            self.status = "can_be_returned"
        else:
            if now <= self.return_deadline + timedelta(days=30):
                self.status = "cannot_be_returned"
            else:
                self.status = "archived"

    def __repr__(self):
        return "%s, %s, %s, %s" % (self.item_name, self.customer_name, self.return_deadline.strftime("%Y-%m-%d %H:%M"), self.status)
