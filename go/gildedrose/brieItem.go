package gildedrose

type defaultItem struct {
	*Item
}

const defaultIncrement = -1

func (i *defaultItem) Set(item *Item) {
	i.Item = item
}
func (i *defaultItem) Update() {
	if i.SellIn < 0 {
		i.Quality += defaultIncrement * 2
	} else {
		i.Quality += defaultIncrement
	}
	i.SellIn--
	i.checkRange()
}
func (i *defaultItem) checkRange() {
	if i.Quality > maxDefaultQuality {
		i.Quality = maxDefaultQuality
	} else if i.Quality < 0 {
		i.Quality = 0
	}
}
