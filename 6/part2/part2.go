package part2

import (
	"fmt"
	"sort"
	"strings"

	"github.com/evansalter/advent-2019/helpers"
)

type Node struct {
	Name   string
	Parent *Node
	Nodes  []*Node
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
			n := &Node{
				Name:   o.Moon,
				Parent: root,
			}
			root.Nodes = append(root.Nodes, n)
			makeTree(n, orbits)
		}
	}
	return root
}

func findNode(root *Node, name string) *Node {
	if root.Name == name {
		return root
	}
	for _, n := range root.Nodes {
		if foundNode := findNode(n, name); foundNode != nil {
			return foundNode
		}
	}
	return nil
}

func findShortestPath(paths [][]*Node) []*Node {
	if len(paths) == 0 {
		return nil
	}

	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
	return paths[0]
}

func findPathsToNode(start *Node, target string, path []*Node) [][]*Node {
	if start.Name == target {
		return [][]*Node{append(path, start)}
	}
	foundPaths := make([][]*Node, 0)
	for _, n := range start.Nodes {
		foundPath := findPathsToNode(n, target, append(path, start))
		foundPaths = append(foundPaths, foundPath...)
	}

	return foundPaths
}

func FindPathsToNode(start *Node, target string, path []*Node) [][]*Node {
	foundPaths := findPathsToNode(start, target, path)
	if start.Parent != nil {
		foundPath := FindPathsToNode(start.Parent, target, append(path, start))
		foundPaths = append(foundPaths, foundPath...)
	}
	return foundPaths
}

func Run() {
	lines := helpers.ReadInputFile(6)
	input := parseInput(lines)

	root := &Node{Name: "COM"}
	tree := makeTree(root, input)
	you := findNode(tree, "YOU")
	paths := FindPathsToNode(you, "SAN", make([]*Node, 0))
	shortest := findShortestPath(paths)
	fmt.Println(len(shortest) - 3) // Subtract 3 to get rid of YOU, SAN and origin orbit
}
