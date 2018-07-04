package graph

type Graph map[string]map[string]int

func (g Graph) Add(key1 string, key2 string, value int) {
	if _, ok := g[key1]; !ok {
		g[key1] = make(map[string]int)
	}
	if _, ok := g[key2]; !ok {
		g[key2] = make(map[string]int)
	}
	g[key1][key2] = value
}
