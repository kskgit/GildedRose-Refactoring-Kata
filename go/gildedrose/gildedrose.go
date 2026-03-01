package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const maxQuality = 50
const minQuality = 0

type itemBehavior interface {
	adjustQuality(item *Item)
	adjustExpiredQuality(item *Item)
}

func updateItem(item *Item, b itemBehavior) {
	b.adjustQuality(item)
	item.SellIn--
	if item.SellIn < 0 {
		b.adjustExpiredQuality(item)
	}
}

func UpdateQuality(items []*Item) {
	// TODO イミュータブルにする
	for i := 0; i < len(items); i++ {
		switch items[i].Name {
		case "Sulfuras, Hand of Ragnaros":
			// do nothing
		case "Aged Brie":
			updateItem(items[i], agedBrie{})
		case "Backstage passes to a TAFKAL80ETC concert":
			updateItem(items[i], backstage{})
		default:
			updateItem(items[i], normal{})
		}
	}
}

type agedBrie struct{}

func (agedBrie) adjustQuality(item *Item) {
	if item.Quality < maxQuality {
		item.Quality++
	}
}

func (agedBrie) adjustExpiredQuality(item *Item) {
	if item.Quality < maxQuality {
		item.Quality++
	}
}

type backstage struct{}

func (backstage) adjustQuality(item *Item) {
	if item.Quality < maxQuality {
		item.Quality++
		if item.SellIn < 11 {
			if item.Quality < maxQuality {
				item.Quality++
			}
		}
		if item.SellIn < 6 {
			if item.Quality < maxQuality {
				item.Quality++
			}
		}
	}
}

func (backstage) adjustExpiredQuality(item *Item) {
	item.Quality = minQuality
}

type normal struct{}

func (normal) adjustQuality(item *Item) {
	if item.Quality > minQuality {
		item.Quality--
	}
}

func (normal) adjustExpiredQuality(item *Item) {
	if item.Quality > minQuality {
		item.Quality--
	}
}
