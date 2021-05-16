package reactor

import "sync"

type SubscriberSet struct {
	m map[Subscriber]bool
	sync.RWMutex
}

func NewSubscribeSet() *SubscriberSet {
	mm := make(map[Subscriber]bool, 0)
	set := new(SubscriberSet)
	set.m = mm
	return set
}

func (s *SubscriberSet) Add(item Subscriber) {
	//写锁
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *SubscriberSet) Remove(item Subscriber) {
	//写锁
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *SubscriberSet) Has(item Subscriber) bool {
	//允许读
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *SubscriberSet) List() []Subscriber {
	//允许读
	s.RLock()
	defer s.RUnlock()
	var outList []Subscriber
	for value := range s.m {
		outList = append(outList, value)
	}
	return outList
}

func (s *SubscriberSet) Len() int {
	return len(s.List())
}

func (s *SubscriberSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[Subscriber]bool{}
}

func (s *SubscriberSet) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}
