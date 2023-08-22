package main

type Node struct {
	data *TrieData
	next *Node
	last *Node
}

type Queue struct {
	head *Node
	tail *Node
}
type TrieData struct {
	trie *TrieNode
	val  string
}

func (q *Queue) init() {
	newHead := new(Node)
	newTail := new(Node)
	newHead.next = newTail
	newTail.last = newHead
	q.head = newHead
	q.tail = newTail
}

func (q *Queue) push(trie *TrieNode, val string) {
	newData := TrieData{
		trie: trie,
		val:  val,
	}
	newNode := Node{
		data: &newData,
		next: q.tail,
		last: q.tail.last,
	}
	q.tail.last.next = &newNode
	q.tail.last = &newNode
}
func (q *Queue) popFront() *TrieData {
	popNode := q.head.next
	q.head.next = popNode.next
	popNode.next.last = q.head
	return popNode.data
}

func (q *Queue) isEmpty() bool {
	return q.head.next == q.tail
}
