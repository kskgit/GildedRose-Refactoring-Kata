package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name == "Sulfuras, Hand of Ragnaros" {
			continue
		}

		switch items[i].Name {
		case "Aged Brie":
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
			}
		case "Backstage passes to a TAFKAL80ETC concert":
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].SellIn < 11 {
					if items[i].Quality < 50 {
						items[i].Quality = items[i].Quality + 1
					}
				}
				if items[i].SellIn < 6 {
					if items[i].Quality < 50 {
						items[i].Quality = items[i].Quality + 1
					}
				}
			}
		default:
			if items[i].Quality > 0 {
				items[i].Quality = items[i].Quality - 1
			}
		}

		items[i].SellIn = items[i].SellIn - 1

		if items[i].SellIn >= 0 {
			continue
		}

		switch items[i].Name {
		case "Aged Brie":
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
			}
		case "Backstage passes to a TAFKAL80ETC concert":
			items[i].Quality = 0
		default:
			if items[i].Quality > 0 {
				items[i].Quality = items[i].Quality - 1
			}
		}
	}

}
