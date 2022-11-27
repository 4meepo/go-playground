package learn_golang

import "testing"

func TestMethodReceiver(t *testing.T) {
	u := User{Name: "Tom"}
	(&u).Hello()
	(&u).Hello()
}

type User struct {
	Name string
}

func (u User) Hello() {
	println("Hello", u.Name)
	u.Name = "Jack"
}
