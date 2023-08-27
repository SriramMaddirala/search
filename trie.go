package main

import (
	"container/heap"
)

type TrieNode struct {
	count   int
	nodes   map[rune]*TrieNode
	wordEnd bool
}

func addQuery(query string, trie *TrieNode) {
	trie.count++
	for index, chaRune := range query {
		if trie.nodes[chaRune] == nil {
			count := 0
			if len(query)-1 == index {
				count = 1
			}
			newNode := &TrieNode{
				count:   count,
				nodes:   make(map[rune]*TrieNode),
				wordEnd: count == 1,
			}
			trie.nodes[chaRune] = newNode
			trie = newNode
		} else {
			trie = trie.nodes[chaRune]
			if len(query)-1 == index {
				trie.count++
			}
		}
	}
}
func startsWith(query string, trie *TrieNode) bool {
	for _, chaRune := range query {
		if trie.nodes[chaRune] == nil {
			return false
		} else {
			trie = trie.nodes[chaRune]
		}
	}
	return true
}
func getTop(query string, trie *TrieNode) []string {
	res := make([]string, 0)
	tempQuery := query
	for _, chaRune := range tempQuery {
		if trie.nodes[chaRune] == nil {
			return res
		} else {
			trie = trie.nodes[chaRune]
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
		for key, node := range currTrie.nodes {
			if node != nil {
				newVal := currVal + string(key)
				if node.wordEnd {
					tryResData := &resData{
						count: node.count,
						val:   currVal + string(key),
					}
					heap.Push(&pq, tryResData)
				}
				q.push(node, newVal)
			}
		}
	}
	top := 7
	for pq.Len() > 0 && top > 0 {
		data := heap.Pop(&pq).(*resData)
		res = append(res, data.val)
		top--
	}
	return res
}
