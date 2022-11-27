package learn_golang

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	type account struct {
		balance int
	}
	as := []account{
		{100},
		{200},
		{300},
	}

	for i := range as {
		as[i].balance += 100
	}
	fmt.Println(as)
}

func TestRangeMap(t *testing.T) {
	m1 := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	m2 := make(map[int]*string)
	for k, v := range m1 {
		m2[k] = &v
	}

	for k, v := range m2 {
		fmt.Println(k, *v)

	}
}

func TestMapOrder(t *testing.T) {
	m1 := make(map[int]struct{})
	m1[1] = struct{}{}
	m1[2] = struct{}{}
	m1[3] = struct{}{}
	m1[4] = struct{}{}

	for k := range m1 {
		fmt.Println(k)
	}
	fmt.Println("round 2 ")
	for k := range m1 {
		fmt.Println(k)
	}
}
