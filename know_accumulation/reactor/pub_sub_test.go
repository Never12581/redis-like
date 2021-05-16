package reactor

import (
	"fmt"
	"testing"
)

const (
	EVENT1 = "event1"
	EVENT2 = "event2"
)

func Test_main(m *testing.T) {
	pub := NewPublisherImpl()

	sub1 := new(SubscribeImpl)
	sub2 := new(SubscribeImpl)

	pub.Subscribe(EVENT1, sub1)
	pub.Subscribe(EVENT2, sub2)

	fmt.Println("------------")
	pub.OnEvent(EVENT1)
	fmt.Println("------------")
	pub.OnEvent(EVENT1)

	pub.UnSubscribe(EVENT1, sub1)

	fmt.Println("------------")
	pub.OnEvent(EVENT1)
	fmt.Println("------------")
	pub.OnEvent(EVENT2)

}
