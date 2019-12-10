package part2

import (
	"fmt"
	"sort"
	"strings"

	"github.com/evansalter/advent-2019/helpers"
)

type Node struct {
	Name  string
	Nodes []*Node
}

type Orbit struct {
	Planet string
	Moon   string
}

func parseInput(input []string) []*Orbit {
	orbits := make([]*Orbit, len(input))
	for i, l := range input {
		parts := strings.Split(l, ")")
		orbits[i] = &Orbit{
			Planet: parts[0],
			Moon:   parts[1],
		}
	}

	sort.Slice(orbits, func(i, j int) bool {
		return orbits[i].Planet < orbits[j].Planet
	})

	return orbits
}

func makeTree(root *Node, orbits []*Orbit) *Node {
	for _, o := range orbits {
		if o.Planet == root.Name {
			n := &Node{Name: o.Moon}
			root.Nodes = append(root.Nodes, n)
			makeTree(n, orbits)
		}
	}
	return root
}

func countNodes(depth int, root *Node) int {
	sum := depth
	for _, n := range root.Nodes {
		sum += countNodes(depth+1, n)
	}
	return sum
}

func Run() {
	lines := helpers.ReadInputFile(6)
	input := parseInput(lines)

	root := &Node{Name: "COM"}
	tree := makeTree(root, input)
	numOrbits := countNodes(0, tree)
	fmt.Println(numOrbits)
}
