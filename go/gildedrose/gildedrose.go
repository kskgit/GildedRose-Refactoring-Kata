package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const maxQuality = 50
const minQuality = 0

type itemBehavior interface {
	adjustQuality(sellIn, quality int) int
	adjustExpiredQuality(quality int) int
}

func updateItem(item *Item, b itemBehavior) *Item {
	quality := b.adjustQuality(item.SellIn, item.Quality)
	sellIn := item.SellIn - 1
	if sellIn < 0 {
		quality = b.adjustExpiredQuality(quality)
	}
	return &Item{Name: item.Name, SellIn: sellIn, Quality: quality}
}

func UpdateQuality(items []*Item) []*Item {
	result := make([]*Item, len(items))
	for i := 0; i < len(items); i++ {
		switch items[i].Name {
		case "Sulfuras, Hand of Ragnaros":
			result[i] = &Item{Name: items[i].Name, SellIn: items[i].SellIn, Quality: items[i].Quality}
		case "Aged Brie":
			result[i] = updateItem(items[i], agedBrie{})
		case "Backstage passes to a TAFKAL80ETC concert":
			result[i] = updateItem(items[i], backstage{})
		default:
			result[i] = updateItem(items[i], normal{})
		}
	}
	return result
}

type agedBrie struct{}

func (agedBrie) adjustQuality(sellIn, quality int) int {
	if quality >= maxQuality {
		return quality
	}
	return quality + 1
}

func (agedBrie) adjustExpiredQuality(quality int) int {
	if quality >= maxQuality {
		return quality
	}
	return quality + 1
}

type backstage struct{}

func (backstage) adjustQuality(sellIn, quality int) int {
	if quality >= maxQuality {
		return quality
	}
	quality++

	if quality >= maxQuality || sellIn >= 11 {
		return quality
	}
	quality++

	if quality >= maxQuality || sellIn >= 6 {
		return quality
	}
	quality++

	return quality
}

func (backstage) adjustExpiredQuality(quality int) int {
	return minQuality
}

type normal struct{}

func (normal) adjustQuality(sellIn, quality int) int {
	if quality <= minQuality {
		return quality
	}
	return quality - 1
}

func (normal) adjustExpiredQuality(quality int) int {
	if quality <= minQuality {
		return quality
	}
	return quality - 1
}
