package main

import (
	"fmt"
	"strings"
)

type martian struct{}

type laser struct {
	pulses int
}
type talker interface {
	talk() string
}

func newLaser(p int) laser {

	return laser{pulses: p}

}

func (l laser) talk() string {
	return strings.Repeat("pew", l.pulses)
}

func (m martian) talk() string {
	return "nack, nack!!"
}

func shout(t talker) string {

	return strings.ToUpper(t.talk())

}

func main() {

	m := martian{}
	l := newLaser(4)

	fmt.Println(m.talk())
	fmt.Println(l.talk())

	var t talker

	t = m

	fmt.Println(t.talk())

	fmt.Println(shout(t))

	t = l

	fmt.Println(t.talk())

	fmt.Println(shout(t))

}
