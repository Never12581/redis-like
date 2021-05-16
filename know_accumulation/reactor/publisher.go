package reactor

type Publisher interface {
	Subscribe(event string, sub Subscriber)
	UnSubscribe(event string, sub Subscriber) bool
	OnEvent(event string)
}

type PublisherImpl struct {
	subscriberMap map[string]*SubscriberSet
}

func NewPublisherImpl() Publisher {
	s := make(map[string]*SubscriberSet)
	pub := new(PublisherImpl)
	pub.subscriberMap = s
	return pub
}

func (p *PublisherImpl) Subscribe(event string, sub Subscriber) {
	// 第一次订阅事件，创建订阅列表并加入订阅者
	subscribers := p.subscriberMap[event]
	if subscribers == nil {
		sSet := NewSubscribeSet()
		sSet.Add(sub)
		p.subscriberMap[event] = sSet
		return
	}

	// 已经订阅的，不做任何反应
	if subscribers.Has(sub) {
		return
	}

	// 将订阅者加入到订阅列表
	subscribers.Add(sub)
}

func (p *PublisherImpl) UnSubscribe(event string, sub Subscriber) bool {
	subscriberSet := p.subscriberMap[event]
	if subscriberSet == nil {
		return false
	}
	if !subscriberSet.Has(sub) {
		return true
	}
	subscriberSet.Remove(sub)
	return true
}

func (p *PublisherImpl) OnEvent(event string) {
	subscriberSet := p.subscriberMap[event]
	for _, sub := range subscriberSet.List() {
		sub.OnReceive(event)
	}
}
