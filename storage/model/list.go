package model

type ListNoder interface {
	Prev() ListNoder
	Next() ListNoder
	Value() interface{}
	SetPrev(noder ListNoder)
	SetNext(noder ListNoder)
}

type ListNode struct {
	prev  ListNoder
	next  ListNoder
	value interface{}
}

func (l *ListNode) SetPrev(noder ListNoder) {
	l.prev = noder
}

func (l *ListNode) SetNext(noder ListNoder) {
	l.next = noder
}

func (l *ListNode) Prev() ListNoder {
	return l.prev
}

func (l *ListNode) Next() ListNoder {
	return l.next
}

func (l *ListNode) Value() interface{} {
	return l.value
}

func newListNode(value interface{}) *ListNode {
	l := new(ListNode)
	l.value = value
	return l
}

// todo 暂时不实现
type ListIter struct {
	next      ListNoder
	direction int64
}

type Lister interface {
	Size() int64
	ListFirst() ListNoder
	ListLast() ListNoder
	AddNodeHead(value interface{})
	AddNodeTail(value interface{})
	AddNodeIndex(value interface{}, index int, after bool)
	DelNode(node ListNoder)
	Iterator() *ListIter
	SearchKey(key interface{}) ListNoder
	SearchIndex(index int64) ListNoder
}

type List struct {
	head  ListNoder
	tail  ListNoder
	len   int64
	match func(n1, n2 interface{}) bool
}

func (l *List) Size() int64 {
	return l.len
}

func (l *List) ListFirst() ListNoder {
	return l.head
}

func (l *List) ListLast() ListNoder {
	return l.tail
}

func (l *List) AddNodeHead(value interface{}) {
	node := newListNode(value)
	if l.Size() == 0 {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.SetPrev(node)
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
		l.tail.SetNext(node)
		l.tail = node
	}
	l.len++
}

func (l *List) AddNodeIndex(value interface{}, index int, after bool) {
	insertNode := newListNode(value)
	currentNode := l.SearchIndex(int64(index))
	if after {
		insertNode.SetPrev(currentNode)
		insertNode.SetNext(currentNode.Next())
		if l.tail == currentNode {
			l.tail = insertNode
		}
	} else {
		insertNode.next = currentNode
		insertNode.prev = currentNode.Next()
		if l.head == currentNode {
			l.head = insertNode
		}
	}
	if insertNode.prev != nil {
		insertNode.prev.SetNext(insertNode)
	}
	if insertNode.next != nil {
		insertNode.next.SetPrev(insertNode)
	}
	l.len++
}

func (l *List) DelNode(node ListNoder) {
	if node.Prev() != nil {
		node.Prev().SetNext(node.Next())
	} else {
		l.head = node.Next()
	}
	if node.Next() != nil {
		node.Next().SetPrev(node.Next())
	} else {
		l.tail = node.Prev()
	}
	l.len--
}

func (l *List) Iterator() *ListIter {
	iter := new(ListIter)
	iter.next = l.head
	iter.direction = 0
	return iter
}

func (l *List) SearchKey(key interface{}) ListNoder {
	node := l.head
	var resultNode ListNoder
	for {
		if l.match(node.Value(), key) {
			resultNode = node
			break
		}
		node = node.Next()
	}
	return resultNode
}

/* Return the element at the specified zero-based index
 * where 0 is the head, 1 is the element next to head
 * and so on. Negative integers are used in order to count
 * from the tail, -1 is the last element, -2 the penultimate
 * and so on. If the index is out of range NULL is returned. */
func (l *List) SearchIndex(index int64) ListNoder {
	if index < 0 {
		tempIndex := -(index)
		if tempIndex > l.len {
			return nil
		}
	}

	var node ListNoder
	if index < 0 {
		index = -(index) - 1
		node = l.tail
		for index > 0 && node != nil {
			index--
			node = node.Prev()
		}
	} else {
		node = l.head
		for index > 0 && node != nil {
			index--
			node = node.Next()
		}
	}
	return node
}
