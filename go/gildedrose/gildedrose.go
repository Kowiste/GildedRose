package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

type UpdateItem interface {
	Set(*Item)
	Update()
}

const maxDefaultQuality = 50

func UpdateQuality(items []*Item) {
	for _, item := range items {
		itemI := selectItem(item.Name)
		itemI.Set(item)
		itemI.Update()
	}
}

func selectItem(name string) (item UpdateItem) {
	switch name {
	case "Aged Brie":
		item = new(brieItem)
	case "Backstage passes to a TAFKAL80ETC concert":
		item = new(backStageItem)
	case "Sulfuras, Hand of Ragnaros":
		item = new(sulfurasItem)
	case "Conjured Mana Cake":
		item = new(conjuredItem)
	default:
		item = new(defaultItem)
	}
	return
}