package learn_golang

import (
	"fmt"
	"testing"
)

func TestMapCompare(t *testing.T) {
	m1 := make(map[string]string)
	m2 := m1

	fmt.Println(m2)
	// fmt.Println(m1 == m2)  compile error , map 只可以和nil 用来 == 比较
}

func TestMapCopy(t *testing.T) {
	m1 := make(map[int]struct{}, 8)
	m2 := m1

	fmt.Printf("m1 pointer %p\n", &m1)
	fmt.Printf("m2 pointer %p\n", &m2)
	for i := 0; i < 5; i++ {
		m1[i] = struct{}{}
	}

	for i := 10; i < 200; i++ {
		m2[i] = struct{}{}
	}

	fmt.Printf("m1 pointer %p\n", &m1)
	fmt.Printf("m2 pointer %p\n", &m2)

}

func TestCopySlice(t *testing.T) {
	s1 := make([]int, 0, 10)
	s2 := s1
	fmt.Printf("s1 address %p\n", &s1)
	fmt.Printf("s2 address %p\n", &s2)

	strA := "hello"
	strB := strA
	fmt.Printf("strA address %p\n", &strA)
	fmt.Printf("strB address %p\n", &strB)

	strB = "world"
	fmt.Println(strA, strB)
}
