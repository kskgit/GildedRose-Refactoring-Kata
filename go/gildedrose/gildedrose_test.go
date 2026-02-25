package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_通常アイテム_SellInとQualityが1減少する(t *testing.T) {
	items := []*gildedrose.Item{
		{"foo", 10, 20},
	}

	gildedrose.UpdateQuality(items)

	assertItem(t, items[0], expected{SellIn: 9, Quality: 19})
}

func Test_通常アイテム_期限切れ後はQualityが2減少する(t *testing.T) {
	items := []*gildedrose.Item{
		{"foo", 0, 20},
	}

	gildedrose.UpdateQuality(items)

	assertItem(t, items[0], expected{SellIn: -1, Quality: 18})
}

func Test_通常アイテム_Qualityは0未満にならない(t *testing.T) {
	items := []*gildedrose.Item{
		{"foo", 5, 0},
	}

	gildedrose.UpdateQuality(items)

	assertItem(t, items[0], expected{SellIn: 4, Quality: 0})
}

func Test_通常アイテム_期限切れ後もQualityは0未満にならない(t *testing.T) {
	items := []*gildedrose.Item{
		{"foo", 0, 1},
	}

	gildedrose.UpdateQuality(items)

	assertItem(t, items[0], expected{SellIn: -1, Quality: 0})
}

func Test_通常アイテム_期限切れ後にQualityが0なら0のまま(t *testing.T) {
	items := []*gildedrose.Item{
		{"foo", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	assertItem(t, items[0], expected{SellIn: -1, Quality: 0})
}

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
