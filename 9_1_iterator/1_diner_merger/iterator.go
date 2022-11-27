package diner_merger

type Iterator interface {
	HasNext() bool
	Next() MenuItem
}
