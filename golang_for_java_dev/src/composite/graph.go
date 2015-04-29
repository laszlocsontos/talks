package main

import "fmt"

type Vertex struct {
	name string

	incomingEdges []*Edge
	outgoingEdges []*Edge
}

type Edge struct {
	tail *Vertex
	head *Vertex

	weight int
}

type Graph struct {
	verticesMap map[string]*Vertex
}

func (g *Graph) AddEdge(tailName, headName string, weight int) {
	tail, ok := g.verticesMap[tailName]

	if !ok {
		tail = &Vertex{name: tailName}
		g.verticesMap[tailName] = tail
	}

	head, ok := g.verticesMap[headName]

	if !ok {
		head = &Vertex{name: headName}
		g.verticesMap[headName] = head
	}

	edge := &Edge{tail, head, weight}

	tail.outgoingEdges = append(tail.outgoingEdges, edge)
	head.incomingEdges = append(head.incomingEdges, edge)
}

func (v *Vertex) String() string {
	return v.name
}

func main() {
	g := Graph{}
	g.verticesMap = make(map[string]*Vertex)

	g.AddEdge("a", "b", 1)
	g.AddEdge("b", "c", 1)
	g.AddEdge("c", "d", 1)
	g.AddEdge("d", "a", 1)

	for n, v := range g.verticesMap {
		for _, e := range v.outgoingEdges {
			fmt.Printf("%s -> %s\n", n, e.head)
		}
	}
}
