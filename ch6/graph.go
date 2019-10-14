package main

import "fmt"

type Graph struct {
	Edges map[*Node][]*Node
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.Edges == nil {
		g.Edges = make(map[*Node][]*Node, 0)
	}

	g.Edges[n1] = append(g.Edges[n1], n2)
}

// searches in the Graph if someone is a Mango Seller
func (g Graph) BFS(startNode *Node) *Node {
	// if I'm already a mango seller, I don't need another!
	if startNode.MangoSeller {
		return startNode
	}

	// initialize queue
	var queue []*Node
	for _, n := range g.Edges[startNode] {
		queue = append(queue, n)
	}

	var searched []*Node

	// search all my friends
	for len(queue) != 0 {

		// pops node from queue
		var n *Node
		n, queue = queue[len(queue)-1], queue[:len(queue)-1]

		// verify if node was already searched
		var alreadySearched bool
		for _, s := range searched {
			if s == n {
				alreadySearched = true
			}
		}

		// if it wasn't searched
		if !alreadySearched {

			// if is the mango seller, return it
			if n.MangoSeller {
				return n
			} else {
				// otherwise, push all Node's friends to the list
				for _, near := range g.Edges[n] {
					queue = append([]*Node{near}, queue...)
				}

				// mark node as searched
				searched = append(searched, n)
			}
		}
	}

	return nil
}

func (g *Graph) String() {
	var s string

	for n, _ := range g.Edges {
		s += n.Name + " -> "

		for _, nn := range g.Edges[n] {
			s += nn.Name + " "
		}

		s += "\n"
	}

	fmt.Println(s)
}

type Node struct {
	Name        string
	MangoSeller bool
}

func (n *Node) String() string {
	return n.Name
}

func main() {
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

	g.String()
}
