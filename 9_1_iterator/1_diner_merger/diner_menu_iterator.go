package diner_merger

type DinerMenuIterator struct {
	Items    []MenuItem
	position int
}

func NewDinerMenuIterator(items []MenuItem) *DinerMenuIterator {
	return &DinerMenuIterator{
		Items:    items,
		position: 0,
	}
}

func (this *DinerMenuIterator) Next() MenuItem {
	res := this.Items[this.position]
	this.position++
	return res
}

func (this *DinerMenuIterator) HasNext() bool {
	return len(this.Items) > this.position
}
