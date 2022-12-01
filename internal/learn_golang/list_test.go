package learn_golang

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := list.New()

	l.PushBack(1)
	l.PushBack(2)
	var e *list.Element
	for {
		e = l.Back()
		if e == nil {
			break
		}
		l.Remove(e)
		fmt.Println(e.Value.(int))
	}
}
