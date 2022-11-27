package diner_merger

import "container/list"

type PancakeHouseMenuIterator struct {
	items       *list.List
	position    *list.Element
	curPosition int
}

func NewPancakeHouseMenuIterator(l *list.List) *PancakeHouseMenuIterator {
	return &PancakeHouseMenuIterator{
		items:       l,
		position:    nil,
		curPosition: 0,
	}
}

func (this *PancakeHouseMenuIterator) Next() MenuItem {
	if this.position == nil {
		this.position = this.items.Front()
	}
	res := this.position.Value.(*MenuItem)
	this.position = this.position.Next()
	this.curPosition++
	return *res
}

func (this *PancakeHouseMenuIterator) HasNext() bool {
	return this.curPosition < this.items.Len()
}
