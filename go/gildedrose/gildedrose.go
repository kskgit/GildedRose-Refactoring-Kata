package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const maxQuality = 50
const minQuality = 0

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		switch items[i].Name {
		case "Sulfuras, Hand of Ragnaros":
			// do nothing
		case "Aged Brie":
			updateAgedBrie(items[i])
		case "Backstage passes to a TAFKAL80ETC concert":
			updateBackstage(items[i])
		default:
			updateNormal(items[i])
		}
	}
}

func updateAgedBrie(item *Item) {
	if item.Quality < maxQuality {
		item.Quality = item.Quality + 1
	}
	item.SellIn = item.SellIn - 1
	if item.SellIn < 0 {
		if item.Quality < maxQuality {
			item.Quality = item.Quality + 1
		}
	}
}

func updateBackstage(item *Item) {
	if item.Quality < maxQuality {
		item.Quality = item.Quality + 1
		if item.SellIn < 11 {
			if item.Quality < maxQuality {
				item.Quality = item.Quality + 1
			}
		}
		if item.SellIn < 6 {
			if item.Quality < maxQuality {
				item.Quality = item.Quality + 1
			}
		}
	}
	item.SellIn = item.SellIn - 1
	if item.SellIn < 0 {
		item.Quality = minQuality
	}
}

func updateNormal(item *Item) {
	if item.Quality > minQuality {
		item.Quality = item.Quality - 1
	}
	item.SellIn = item.SellIn - 1
	if item.SellIn < 0 {
		if item.Quality > minQuality {
			item.Quality = item.Quality - 1
		}
	}
}
