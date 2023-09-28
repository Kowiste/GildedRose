package gildedrose

type conjuredItem struct {
	*Item
}

const conjuredIncrement = -2

func (i *conjuredItem) Set(item *Item) {
	i.Item = item
}
func (i *conjuredItem) Update() {
	if i.SellIn < 0 {
		i.Quality += conjuredIncrement * 2
	} else {
		i.Quality += conjuredIncrement
	}
	i.SellIn--
	i.checkRange()
}
func (i *conjuredItem) checkRange() {
	if i.Quality > maxDefaultQuality {
		i.Quality = maxDefaultQuality
	} else if i.Quality < 0 {
		i.Quality = 0
	}
}
