package learn_golang

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Println(now.String())

	l, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(l.String())

	l2, _ := time.LoadLocation("Local")
	fmt.Println(l2.String())

	utcTime := time.Now().In(time.UTC)
	fmt.Println(utcTime.String())

	l2Time := time.Now().In(l2)
	fmt.Println(l2Time.String())

}
