package main

import (
	"testing"
)

func Test_GraphAddEdge(t *testing.T) {
	g := Graph{}
	n1 := &Node{
		Name:        "jojo",
		MangoSeller: false,
	}
	n2 := &Node{
		Name:        "theodore",
		MangoSeller: true,
	}

	g.AddEdge(n1, n2)

	if g.Edges[n1][0] != n2 {
		t.Errorf("expected %v, got %v", n2, g.Edges[n1][0])
	}
}

func Test_BreadthFirstSearch(t *testing.T) {
	g := Graph{}
	jojo := &Node{Name: "jojo", MangoSeller: false}
	claire := &Node{Name: "claire", MangoSeller: false}
	jonny := &Node{Name: "jonny", MangoSeller: false}
	thom := &Node{Name: "thom", MangoSeller: false}
	alice := &Node{Name: "alice", MangoSeller: false}
	peggy := &Node{Name: "peggy", MangoSeller: false}
	bob := &Node{Name: "bob", MangoSeller: false}
	anuj := &Node{Name: "anuj", MangoSeller: false}

	g.AddEdge(jojo, claire)
	g.AddEdge(claire, thom)
	g.AddEdge(claire, jonny)
	g.AddEdge(jojo, alice)
	g.AddEdge(alice, peggy)
	g.AddEdge(jojo, bob)
	g.AddEdge(bob, anuj)
	g.AddEdge(bob, peggy)

	t.Run("claire is the mango seller", func(t *testing.T) {
		claire.MangoSeller = true

		n := g.BFS(jojo)

		if n != claire {
			t.Errorf("expected %v, got %v", claire, n)
		}

		claire.MangoSeller = false
	})

	t.Run("anuj is the mango seller", func(t *testing.T) {
		anuj.MangoSeller = true

		n := g.BFS(jojo)

		if n != anuj {
			t.Errorf("expected %v, got %v", anuj, n)
		}

		anuj.MangoSeller = false
	})

	t.Run("jojo is the mango seller", func(t *testing.T) {
		jojo.MangoSeller = true

		n := g.BFS(jojo)

		if n != jojo {
			t.Errorf("expected %v, got %v", jojo, n)
		}

		jojo.MangoSeller = false
	})
}

func Test_BFSInfiniteLoop(t *testing.T) {
	g := Graph{}
	jojo := &Node{Name: "jojo", MangoSeller: false}
	peggy := &Node{Name: "peggy", MangoSeller: false}

	g.AddEdge(jojo, peggy)
	g.AddEdge(peggy, jojo)

	n := g.BFS(jojo)
	if n != nil {
		t.Errorf("expected %v, got %v", nil, n)
	}
}
