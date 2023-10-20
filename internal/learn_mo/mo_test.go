package learnmo

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/samber/mo"
)

func TestFuture(t *testing.T) {

	f := mo.NewFuture(func(resolve func(string), reject func(error)) {
		time.Sleep(time.Second * 3)
		if time.Now().Second()%2 == 0 {
			reject(errors.New("failed."))
		} else {
			resolve("success")
		}
	}).Catch(func(err error) (string, error) {
		fmt.Println("catch error")
		return "catch", err
	}).Then(func(s string) (string, error) {
		fmt.Println("then")
		return s, nil
	})
	x, err := f.Collect()
	fmt.Println(x, err)
}
