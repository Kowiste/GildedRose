package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 3, 20},
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Aged Brie", 2, 50},
		{"Sulfuras, Hand of Ragnaros", 1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 48},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 47},
		{"Conjured Mana Cake", 3, 6},
		{"Conjured Mana Cake", -3, 6},
		{"Aged Brie", -2, 10},
		{"Backstage passes to a TAFKAL80ETC concert", -5, 6},
		{"Conjured Mana Cake", 3, -6},
	}

	gildedrose.UpdateQuality(items)

	//Item 0: +5 Dexterity Vest
	assert.Equal(t, 19, items[0].Quality)
	assert.Equal(t, 2, items[0].SellIn)

	//Item 1: +5 Dexterity Vest
	assert.Equal(t, 19, items[1].Quality)
	assert.Equal(t, 9, items[1].SellIn)

	//Item 2: Aged Brie
	assert.Equal(t, 1, items[2].Quality)
	assert.Equal(t, 1, items[2].SellIn)

	//Item 3: Aged Brie
	assert.Equal(t, 50, items[3].Quality)
	assert.Equal(t, 1, items[3].SellIn)

	//Item 4: Sulfuras, Hand of Ragnaros
	assert.Equal(t, 80, items[4].Quality)
	assert.Equal(t, 0, items[4].SellIn)

	//Item 5: Backstage passes to a TAFKAL80ETC concert
	assert.Equal(t, 21, items[5].Quality)
	assert.Equal(t, 14, items[5].SellIn)

	//Item 6: Backstage passes to a TAFKAL80ETC concert
	assert.Equal(t, 50, items[6].Quality)
	assert.Equal(t, 9, items[6].SellIn)

	//Item 7: Backstage passes to a TAFKAL80ETC concert
	assert.Equal(t, 50, items[7].Quality)
	assert.Equal(t, 4, items[7].SellIn)

	//Item 8: Conjured Mana Cake
	assert.Equal(t, 4, items[8].Quality)
	assert.Equal(t, 2, items[8].SellIn)

	//Item 9: Conjured Mana Cake
	assert.Equal(t, 2, items[9].Quality)
	assert.Equal(t, -4, items[9].SellIn)

	//Item 10:Aged Brie
	assert.Equal(t, 12, items[10].Quality)
	assert.Equal(t, -3, items[10].SellIn)

	//Item 11: Backstage passes to a TAFKAL80ETC concert
	assert.Equal(t, 0, items[11].Quality)
	assert.Equal(t, -6, items[11].SellIn)

	//Item 12: Conjured Mana Cake
	assert.Equal(t, 0, items[12].Quality)
	assert.Equal(t, 2, items[12].SellIn)
}
