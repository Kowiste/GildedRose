package gildedrose

type brieItem struct {
	*Item
}

const brieIncrement = 1

func (i *brieItem) Set(item *Item) {
	i.Item = item
}
func (i *brieItem) Update() {
	if i.SellIn < 0 {
		i.Quality += brieIncrement * 2
	} else {
		i.Quality += brieIncrement
	}
	i.SellIn--
	i.checkRange()
}
func (i *brieItem) checkRange() {
	if i.Quality > maxDefaultQuality {
		i.Quality = maxDefaultQuality
	} else if i.Quality < 0 {
		i.Quality = 0
	}
}
