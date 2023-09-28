package gildedrose

type backStageItem struct {
	*Item
}


func (i *backStageItem) Set(item *Item) {
	i.Item = item
}
func (i *backStageItem) Update() {
	if i.SellIn < 0 {
		i.Quality = 0
	} else if i.SellIn <= 5 {
		i.Quality += 3
	} else if i.SellIn <= 10 {
		i.Quality += 2
	} else {
		i.Quality++
	}
	i.SellIn--
	i.checkRange()
}
func (i *backStageItem) checkRange() {
	if i.Quality > maxDefaultQuality {
		i.Quality = maxDefaultQuality
	} else if i.Quality < 0 {
		i.Quality = 0
	}
}
