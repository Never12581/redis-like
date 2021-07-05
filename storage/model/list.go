package model

import (
	"fmt"
	"strings"
)

type Lister interface {
	Size() int64
	ListFirst() *ListNode
	ListLast() *ListNode
	AddNodeHead(value interface{})
	AddNodeTail(value interface{})
	AddNodeIndex(value interface{}, index int, after bool)
	DelNode(node *ListNode)
	Iterator() *ListIter
	SearchKey(key interface{}) *ListNode
	SearchIndex(index int64) *ListNode
}

type ListNode struct {
	prev  *ListNode
	next  *ListNode
	value interface{}
}

func newListNode(value interface{}) *ListNode {
	l := new(ListNode)
	l.value = value
	return l
}

type ListIter struct {
	next      *ListNode
	direction int64
}

type List struct {
	head  *ListNode
	tail  *ListNode
	len   int64
	match func(n1, n2 *ListNode) bool
}

func (l *List) Size() int64 {
	return l.len
}

func (l *List) ListFirst() *ListNode {
	return l.head
}

func (l *List) ListLast() *ListNode {
	return l.tail
}

func (l *List) AddNodeHead(value interface{}) {
	node := newListNode(value)
	if l.Size() == 0 {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.len++
}

func (l *List) AddNodeTail(value interface{}) {
	node := newListNode(value)
	if l.Size() == 0 {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.len++
}

func (l *List) AddNodeIndex(value interface{}, index int, after bool) {
	insertNode := newListNode(value)
	currentNode := l.SearchIndex(int64(index))
	if after {
		insertNode.prev = currentNode
		insertNode.next = currentNode.next
		if l.tail == currentNode {
			l.tail = insertNode
		}
	} else {
		insertNode.next = currentNode
		insertNode.prev = currentNode.prev
		if l.head == currentNode {
			l.head = insertNode
		}
	}
	if insertNode.prev != nil {
		insertNode.prev.next = insertNode
	}
	if insertNode.next != nil {
		insertNode.next.prev = insertNode
	}
	l.len++
}

func (l *List) DelNode(node *ListNode) {
	panic("implement me")
}

func (l *List) Iterator() *ListIter {
	panic("implement me")
}

func (l *List) SearchKey(key interface{}) *ListNode {
	panic("implement me")
}

/* Return the element at the specified zero-based index
 * where 0 is the head, 1 is the element next to head
 * and so on. Negative integers are used in order to count
 * from the tail, -1 is the last element, -2 the penultimate
 * and so on. If the index is out of range NULL is returned. */
func (l *List) SearchIndex(index int64) *ListNode {
	if index < 0 {
		tempIndex := -(index)
		if tempIndex > l.len {
			return nil
		}
	}

	var node *ListNode
	if index < 0 {
		index = -(index) - 1
		node = l.tail
		for index > 0 && node != nil {
			index--
			node = node.prev
		}
	} else {
		node = l.head
		for index > 0 && node != nil {
			index--
			node = node.next
		}
	}
	return node
}

func (l *List) ToString() string {
	format := `len:%v`
	s := fmt.Sprintf(format, l.len)
	sb := strings.Builder{}
	sb.WriteString(s)
	node := l.head
	for node != nil {
		sb.WriteString(",")
		sb.WriteString((node.value).(string))
	}
	return sb.String()
}
