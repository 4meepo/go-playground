package learnfp

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	IO "github.com/IBM/fp-go/io"

	R "github.com/IBM/fp-go/reader"
)

func TestRandomInt(t *testing.T) {
	b := IO.Map(intToStr)
	r := randomInt(100)
	f := b(r)
	fmt.Println("nothing happens until we call f():")

	// r1 := f()
	fmt.Println(f(), f(), f())

	rd1 := R.MakeReader(foo)
	fmt.Printf("%T\n", rd1)

	rd2 := foo
	fmt.Println(rd1(context.Background()))
	fmt.Println(rd2(context.Background()))

}

func randomInt(n int) IO.IO[int] {
	return func() int {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		return r.Intn(n)
	}
}

func intToStr(n int) string {
	return fmt.Sprintf("%d", n)
}

func foo(ctx context.Context) string {
	fmt.Println("foo called")
	return "foo"
}
