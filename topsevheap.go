package main

type trieHeap []*resData

type resData struct {
	count int
	val   string
}

func (pq trieHeap) Len() int {
	return len(pq)
}

// maxheap right now
func (pq trieHeap) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq trieHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *trieHeap) Push(x any) {
	data := x.(*resData)
	*pq = append(*pq, data)
}
func (pq *trieHeap) Pop() any {
	old := *pq
	n := len(old)
	trieData := old[n-1]
	*pq = old[0 : n-1]
	return trieData
}
