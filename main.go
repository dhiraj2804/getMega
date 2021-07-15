package main

import (
	"practice/getMega/pubsub"
	"time"
)

func main() {
	ps := pubsub.NewPubsub()

	ch1 := ps.CreateTopic("Math")
	ch2 := ps.CreateTopic("Geography")
	ch3 := ps.CreateTopic("Geography")

	go ps.Subscribe("1", ch1)
	go ps.Subscribe("2", ch2)
	go ps.Subscribe("3", ch3)

	time.Sleep(50 * time.Millisecond)
	ps.Pub("Math", "Binomial")
	time.Sleep(2 * time.Second)
	ps.Pub("Science", "vitamins")
	time.Sleep(2 * time.Second)
	ps.Pub("Math", "Permutation")
	time.Sleep(2 * time.Second)
	ps.Pub("Geography", "Newzealand")
	time.Sleep(2 * time.Second)
	ps.Pub("Geography", "Australia")
	time.Sleep(2 * time.Second)
	// ps.DeleteSubscription(ch1)
	ps.Pub("Math", "Quadratic equation")

	time.Sleep(50 * time.Millisecond)
}
