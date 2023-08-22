package main

import (
	"container/heap"
	"fmt"
)

type TrieNode struct {
	count   int
	nodes   []*TrieNode
	wordEnd bool
}

func addQuery(query string, trie *TrieNode) {
	trie.count++
	for len(query) > 0 {
		charVal := query[0] - 'a'
		if trie.nodes[charVal] == nil {
			count := 0
			if len(query) == 1 {
				count = 1
			}
			newNode := &TrieNode{
				count:   count,
				nodes:   make([]*TrieNode, 26),
				wordEnd: count == 1,
			}
			trie.nodes[charVal] = newNode
			query = query[1:]
			trie = newNode
		} else {
			trie = trie.nodes[charVal]
			if len(query) == 1 {
				trie.count++
			}
			query = query[1:]
		}
	}
}
func startsWith(query string, trie *TrieNode) bool {
	for len(query) > 0 {
		charVal := query[0] - 'a'
		if trie.nodes[charVal] == nil {
			return false
		} else {
			trie = trie.nodes[charVal]
			query = query[1:]
		}
	}
	return true
}
func getTop(query string, trie *TrieNode) []string {
	res := make([]string, 0)
	tempQuery := query
	for len(tempQuery) > 0 {
		charVal := tempQuery[0] - 'a'
		if trie.nodes[charVal] == nil {
			return res
		} else {
			trie = trie.nodes[charVal]
			tempQuery = tempQuery[1:]
		}
	}
	q := new(Queue)
	q.init()
	pq := make(trieHeap, 0)
	heap.Init(&pq)
	q.push(trie, query)
	for !q.isEmpty() {
		trieData := q.popFront()
		currTrie := trieData.trie
		currVal := trieData.val
		for i := 0; i < 26; i++ {
			tryTrie := currTrie.nodes[i]
			if tryTrie != nil {
				newVal := currVal + string(97+i)
				if tryTrie.wordEnd {
					tryResData := &resData{
						count: tryTrie.count,
						val:   newVal,
					}
					heap.Push(&pq, tryResData)
				}
				q.push(tryTrie, newVal)
			}
		}
	}
	top := 7
	for pq.Len() > 0 && top > 0 {
		data := heap.Pop(&pq).(*resData)
		res = append(res, data.val+" count is "+fmt.Sprint(data.count))
		top--
	}
	return res
}
