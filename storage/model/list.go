package model

type listNode struct {
	prev  *listNode
	next  *listNode
	value interface{}
}

type ListIter struct {
	next      *listNode
	direction int64
}

type List struct {
	head *listNode
	tail *listNode
	len  int64
}

type Lister interface {
}
