package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

//
// 通常アイテム
//

func Test_通常アイテム_SellInとQualityが1減少する(t *testing.T) {
	items := []*gildedrose.Item{
		{"foo", 10, 20},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 9, Quality: 19})
}

func Test_通常アイテム_期限切れ後はQualityが2減少する(t *testing.T) {
	items := []*gildedrose.Item{
		{"foo", 0, 20},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: -1, Quality: 18})
}

func Test_通常アイテム_Qualityは0未満にならない(t *testing.T) {
	items := []*gildedrose.Item{
		{"foo", 5, 0},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 4, Quality: 0})
}

func Test_通常アイテム_期限切れ後もQualityは0未満にならない(t *testing.T) {
	items := []*gildedrose.Item{
		{"foo", 0, 1},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: -1, Quality: 0})
}

func Test_通常アイテム_期限切れ後にQualityが0なら0のまま(t *testing.T) {
	items := []*gildedrose.Item{
		{"foo", 0, 0},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: -1, Quality: 0})
}

//
// Sulfuras
//

func Test_Sulfuras_SellInもQualityも変化しない(t *testing.T) {
	items := []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 0, Quality: 80})
}

func Test_Sulfuras_SellInが正の場合も変化しない(t *testing.T) {
	items := []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 10, 80},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 10, Quality: 80})
}

func Test_Sulfuras_SellInが負の場合も変化しない(t *testing.T) {
	items := []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", -1, 80},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: -1, Quality: 80})
}

//
// AgedBrie
//

func Test_AgedBrie_日が経つほどQualityが上がる(t *testing.T) {
	items := []*gildedrose.Item{
		{"Aged Brie", 10, 20},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 9, Quality: 21})
}

func Test_AgedBrie_1日前はQualityが1上がる_境界値(t *testing.T) {
	items := []*gildedrose.Item{
		{"Aged Brie", 1, 20},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 0, Quality: 21})
}

func Test_AgedBrie_期限切れ後はQualityが2上がる_境界値(t *testing.T) {
	items := []*gildedrose.Item{
		{"Aged Brie", 0, 20},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: -1, Quality: 22})
}

func Test_AgedBrie_Qualityは50を超えない(t *testing.T) {
	items := []*gildedrose.Item{
		{"Aged Brie", 10, 50},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 9, Quality: 50})
}

func Test_AgedBrie_期限切れ後もQualityは50を超えない(t *testing.T) {
	items := []*gildedrose.Item{
		{"Aged Brie", 0, 49},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: -1, Quality: 50})
}

//
// Backstage passes
//

const backstage = "Backstage passes to a TAFKAL80ETC concert"

func Test_Backstage_11日前はQualityが1上がる(t *testing.T) {
	items := []*gildedrose.Item{
		{backstage, 11, 20},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 10, Quality: 21})
}

func Test_Backstage_10日前はQualityが2上がる_境界値(t *testing.T) {
	items := []*gildedrose.Item{
		{backstage, 10, 20},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 9, Quality: 22})
}

func Test_Backstage_6日前はQualityが2上がる(t *testing.T) {
	items := []*gildedrose.Item{
		{backstage, 6, 20},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 5, Quality: 22})
}

func Test_Backstage_5日前はQualityが3上がる_境界値(t *testing.T) {
	items := []*gildedrose.Item{
		{backstage, 5, 20},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 4, Quality: 23})
}

func Test_Backstage_1日前はQualityが3上がる(t *testing.T) {
	items := []*gildedrose.Item{
		{backstage, 1, 20},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 0, Quality: 23})
}

func Test_Backstage_コンサート終了後はQualityが0になる(t *testing.T) {
	items := []*gildedrose.Item{
		{backstage, 0, 20},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: -1, Quality: 0})
}

func Test_Backstage_Qualityは50を超えない(t *testing.T) {
	items := []*gildedrose.Item{
		{backstage, 5, 49},
	}

	updated := gildedrose.UpdateQuality(items)

	assertItem(t, updated[0], expected{SellIn: 4, Quality: 50})
}

//
// Conjured
//

// func Test_Conjured_Qualityが2減少する(t *testing.T) {
// 	items := []*gildedrose.Item{
// 		{"Conjured Mana Cake", 10, 20},
// 	}

// 	gildedrose.UpdateQuality(items)

// 	assertItem(t, items[0], expected{SellIn: 9, Quality: 18})
// }

// func Test_Conjured_期限切れ後はQualityが4減少する(t *testing.T) {
// 	items := []*gildedrose.Item{
// 		{"Conjured Mana Cake", 0, 20},
// 	}

// 	gildedrose.UpdateQuality(items)

// 	assertItem(t, items[0], expected{SellIn: -1, Quality: 16})
// }

// func Test_Conjured_Qualityは0未満にならない(t *testing.T) {
// 	items := []*gildedrose.Item{
// 		{"Conjured Mana Cake", 5, 1},
// 	}

// 	gildedrose.UpdateQuality(items)

// 	assertItem(t, items[0], expected{SellIn: 4, Quality: 0})
// }

// func Test_Conjured_期限切れ後もQualityは0未満にならない(t *testing.T) {
// 	items := []*gildedrose.Item{
// 		{"Conjured Mana Cake", 0, 3},
// 	}

// 	gildedrose.UpdateQuality(items)

// 	assertItem(t, items[0], expected{SellIn: -1, Quality: 0})
// }

// func Test_Conjured_期限切れ後にQualityが0なら0のまま(t *testing.T) {
// 	items := []*gildedrose.Item{
// 		{"Conjured Mana Cake", 0, 0},
// 	}

// 	gildedrose.UpdateQuality(items)

// 	assertItem(t, items[0], expected{SellIn: -1, Quality: 0})
// }

type expected struct {
	SellIn  int
	Quality int
}

func assertItem(t *testing.T, item *gildedrose.Item, e expected) {
	t.Helper()
	if item.SellIn != e.SellIn {
		t.Errorf("SellIn: expected %d but got %d", e.SellIn, item.SellIn)
	}
	if item.Quality != e.Quality {
		t.Errorf("Quality: expected %d but got %d", e.Quality, item.Quality)
	}
}
