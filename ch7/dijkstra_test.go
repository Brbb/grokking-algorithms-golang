package main

import (
	"reflect"
	"testing"
)

func Test_Graph(t *testing.T) {
	a := &Node{Name: "a"}
	b := &Node{Name: "b"}
	c := &Node{Name: "c"}
	d := &Node{Name: "d"}

	g := Graph{}
	g.AddEdge(a, b, 1)
	g.AddEdge(a, c, 2)
	g.AddEdge(b, d, 9)
	g.AddEdge(c, b, 10)
	g.AddEdge(c, d, 5)

	t.Run("AddNode", func(t *testing.T) {
		if len(g.Nodes) != 4 {
			t.Errorf("expected %d, got %d", 4, len(g.Nodes))
		}

		nodes := []*Node{a, b, c, d}
		if !reflect.DeepEqual(g.Nodes, nodes) {
			t.Errorf("expected %v, got %v", nodes, g.Nodes)
		}
	})

	t.Run("AddEdge", func(t *testing.T) {
		if len(g.Edges) != 5 {
			t.Errorf("expected %d, got %d", 5, len(g.Edges))
		}

		edges := []*Edge{
			{
				Parent: a,
				Child:  b,
				Cost:   1,
			},
			{
				Parent: a,
				Child:  c,
				Cost:   2,
			},
			{
				Parent: b,
				Child:  d,
				Cost:   9,
			},
			{
				Parent: c,
				Child:  b,
				Cost:   10,
			},
			{
				Parent: c,
				Child:  d,
				Cost:   5,
			},
		}
		if !reflect.DeepEqual(g.Edges, edges) {
			t.Errorf("expected %v, got %v", edges, g.Edges)
		}
	})

	t.Run("NewCostTable", func(t *testing.T) {
		costTable := g.NewCostTable(a)

		expectedCostTable := make(map[*Node]int)
		expectedCostTable[a] = 0
		expectedCostTable[b] = Infinity
		expectedCostTable[c] = Infinity
		expectedCostTable[d] = Infinity

		if !reflect.DeepEqual(costTable, expectedCostTable) {
			t.Errorf("expected %v, got %v", expectedCostTable, costTable)
		}

	})
	t.Run("GedNodeEdges", func(t *testing.T) {
		edges := g.GetNodeEdges(a)

		expectedEdges := []*Edge{
			{
				Parent: a,
				Child:  b,
				Cost:   1,
			},
			{
				Parent: a,
				Child:  c,
				Cost:   2,
			},
		}

		if !reflect.DeepEqual(edges, expectedEdges) {
			t.Errorf("expected %v, got %v", expectedEdges, edges)
		}

	})
	t.Run("getClosestVisitedNode", func(t *testing.T) {
		costTable := g.NewCostTable(a)
		node := getClosestNonVisitedNode(costTable, []*Node{})

		if node != a {
			t.Errorf("expected %v, got %v", a, node)
		}

	})
	t.Run("Dijkstra", func(t *testing.T) {
		shortestPathTable := g.Dijkstra(a)

		expectedTable := `Distance from a to a = 0
Distance from a to b = 1
Distance from a to c = 2
Distance from a to d = 7
`
		if shortestPathTable != expectedTable {
			t.Errorf("expected %s, got %s", expectedTable, shortestPathTable)
		}
	})
}
