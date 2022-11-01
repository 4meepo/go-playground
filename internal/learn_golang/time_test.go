package learn_golang

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

	zeroTime := time.Time{}
	assert.True(t, zeroTime.IsZero())

}

func TestTimeLayout(t *testing.T) {
	// 1、测试parse
	// layout中缺失timezone时，默认使用UTC时间
	parsedTime, err := time.Parse("2006-01-02 15:04:05", "2022-04-10 22:11:01")
	if nil != err {
		t.Error(err)
	}
	fmt.Println(parsedTime) // 输出的时间为UTC时间

	// 2、测试parseInLocation
	parsedTime2, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-04-10 22:11:01", time.Local)
	if nil != err {
		t.Error(err)
	}
	fmt.Println(parsedTime2) // 输出的时间为Local时间,which is CST

	eightHours, _ := time.ParseDuration("8h")
	assert.Equal(t, eightHours, parsedTime.Sub(parsedTime2))

	// 通过in可以将时间转换为指定时区的时间
	fmt.Println(time.Now().In(time.UTC))
}
