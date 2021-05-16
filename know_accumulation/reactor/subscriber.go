package reactor

import "fmt"

type Subscriber interface {
	OnReceive(event string)
}

type SubscribeImpl struct {
}

func (s *SubscribeImpl) OnReceive(event string) {
	fmt.Printf("receive event: %v\n", event)
}
