package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
)

type edge struct {
	node1Id, node2Id int
	cost             int
}

type DisjointSet struct {
	parent []int
}

func NewDS(length int) DisjointSet {
	ds := DisjointSet{parent: make([]int, length)}
	for i := 0; i < length; i++ {
		ds.parent[i] = i
	}
	return ds
}

func (ds *DisjointSet) find(node int) int {
	if ds.parent[node] != node {
		ds.parent[node] = ds.find(ds.parent[node])
	}
	return ds.parent[node]
}
func (ds *DisjointSet) union(node1, node2 int) {
	ds.parent[node1] = ds.find(node2)

}

func minCostConnectPoints(points [][]int) int {
	edges := make([]edge, 0)
	for i, point := range points {
		for j := i + 1; j < len(points); j++ {
			distance := int(math.Abs(float64(point[0]-points[j][0])) + math.Abs(float64(point[1]-points[j][1])))
			edges = append(edges, edge{
				node1Id: i,
				node2Id: j,
				cost:    distance,
			})
		}
	}
	slices.SortFunc(edges, func(a, b edge) int {
		return cmp.Compare(a.cost, b.cost)
	})
	fmt.Println(edges)
	totalCost := 0
	ds := NewDS(len(points))
	for _, e := range edges {
		if ds.find(e.node1Id) == ds.find(e.node2Id) {
			continue
		} else {
			ds.union(ds.find(e.node1Id), ds.find(e.node2Id))
			totalCost += e.cost
		}
		fmt.Println("added ", e, " to tree. new cost:", totalCost)
		fmt.Println(ds.parent)
	}

	return totalCost
}

func main() {
	//points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	points := [][]int{{-1000000, -1000000}, {1000000, 1000000}}
	fmt.Println(minCostConnectPoints(points))
}
