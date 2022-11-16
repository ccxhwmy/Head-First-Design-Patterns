package sort

type Duck struct {
	name   string
	weight int
}

type DuckSort []Duck

func (duckSort DuckSort) Less(i, j int) bool {
	return duckSort[i].weight < duckSort[j].weight
}

func (duckSort DuckSort) Len() int {
	return len(duckSort)
}

func (duckSort DuckSort) Swap(i, j int) {
	duckSort[i], duckSort[j] = duckSort[j], duckSort[i]
}
