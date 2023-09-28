package gildedrose

type sulfurasItem struct {
	*Item
}

const maxSulfurasQuality = 80

func (i *sulfurasItem) Set(item *Item) {
	i.Item = item
}
func (i *sulfurasItem) Update() {
	i.Quality = maxSulfurasQuality
	i.SellIn = 0
}
