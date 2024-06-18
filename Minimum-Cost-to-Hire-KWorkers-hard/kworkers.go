package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"math"
	"slices"
)

type Worker struct {
	quality int
	rank    float64
}

type WorkerHeap []*Worker

func (w WorkerHeap) Len() int {
	return len(w)
}

func (w WorkerHeap) Less(i, j int) bool {
	return w[i].quality > w[j].quality // FIXME: swap?
}

func (w WorkerHeap) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w *WorkerHeap) Push(x any) {
	*w = append(*w, x.(*Worker))
}

func (w *WorkerHeap) Pop() any {
	old := *w
	n := len(old)
	last := old[n-1]
	*w = old[0 : n-1]
	return last
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	workers := make([]Worker, len(quality))
	rank := make([]float64, len(quality))
	for i := 0; i < len(quality); i++ {
		rank[i] = float64(wage[i]) / float64(quality[i])
		workers[i] = Worker{quality: quality[i], rank: rank[i]}
	}

	slices.SortFunc(workers, func(a, b Worker) int {
		return cmp.Compare(a.rank, b.rank)
	})
	//fmt.Println(workers)

	h := make(WorkerHeap, 0)
	heap.Init(&h)

	minCost := math.MaxFloat64
	cumulativeQuality := 0
	for _, w := range workers {
		heap.Push(&h, &w)
		cumulativeQuality += w.quality
		//fmt.Println(fmt.Sprintf("checkpre-> cumulq:%d, minCost:%.2f", cumulativeQuality, minCost))
		if h.Len() == k {
			minCost = min(minCost, float64(cumulativeQuality)*w.rank)
		} else if h.Len() > k {
			last := heap.Pop(&h).(*Worker)
			cumulativeQuality -= last.quality
			minCost = min(minCost, float64(cumulativeQuality)*w.rank)
			//fmt.Println("qu:", last.quality, "last:", last)
		}
		//fmt.Println(fmt.Sprintf("check-> cumulq:%d, minCost:%.2f", cumulativeQuality, minCost))

	}
	//fmt.Println(minCost)
	//heap.Push(&h, &Worker{quality: 1, rank: 5.0})

	//for h.Len() > 0 {
	//	item := heap.Pop(&h).(*Worker)
	//	fmt.Printf("%.2f:%d-", item.rank, item.quality)
	//}
	return minCost
}

func main() {
	quality := []int{3, 1, 10, 10, 1}
	wage := []int{4, 8, 2, 2, 7}
	k := 3
	fmt.Println(mincostToHireWorkers(quality, wage, k))

}
