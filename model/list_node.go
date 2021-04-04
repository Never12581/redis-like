package model

type ListNode struct {
	prev  *ListNode
	next  *ListNode
	value interface{}
}

func NewListNode(val interface{}) *ListNode {
	return &ListNode{value: val}
}

type List struct {
	head   *ListNode
	tail   *ListNode
	length int64
}

func NewList() *List {
	return &List{}
}

func (l *List) ListLength() int64 {
	return l.length
}

func (l *List) ListFirst() *ListNode {
	return l.head
}

func (l *List) ListEmpty() {
	l = NewList()
}

func (l *List) ListAddNodeHead(value interface{}) {
	node := NewListNode(value)
	if l.length == 0 {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.length++
}

func (l *List) ListAddNodeTail(value interface{}) {
	node := NewListNode(value)
	if l.length == 0 {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.length++
}

func (l *List) ListInsertNode(oldNode *ListNode, value interface{}, after bool) {
	node := NewListNode(value)
	if after {
		node.prev = oldNode
		node.next = oldNode.next
		if l.tail == oldNode {
			l.tail = node
		}
	} else {
		node.next = oldNode
		node.prev = oldNode.prev
		if l.head == oldNode {
			l.head = node
		}
	}
	l.length++
}

func (l *List) ListDelNode(node *ListNode) {
	if node == l.head {
		l.head = node.next
	} else {
		node.prev.next = node.next
	}
	if node == l.tail {
		l.tail = node.prev
	} else {
		node.next.prev = node.prev
	}
	l.length--
}
