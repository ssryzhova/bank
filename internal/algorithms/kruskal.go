package algorithms

type Edge struct {
	From, To int
	Weight   int
}

type Graph struct {
	Vertices int
	Edges    []Edge
}

type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &DSU{parent}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	d.parent[d.Find(x)] = d.Find(y)
}

func sortEdges(edges []Edge) {
	for i := 0; i < len(edges); i++ {
		for j := i + 1; j < len(edges); j++ {
			if edges[i].Weight > edges[j].Weight {
				edges[i], edges[j] = edges[j], edges[i]
			}
		}
	}
}

func Kruskal(g Graph) []Edge {
	edges := g.Edges

	sortEdges(edges)

	dsu := NewDSU(g.Vertices)
	var result []Edge

	for _, e := range edges {
		if dsu.Find(e.From) != dsu.Find(e.To) {
			result = append(result, e)
			dsu.Union(e.From, e.To)
		}
	}

	return result
}
