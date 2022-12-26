package learn_golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {

	dogTrustBridge := Dog{Name: "trust bridge", FriendsMap: make(map[string]Dog)}
	dogTrustBridge.changeName("mother fucker trustbridge")
	assert.Equal(t, "trust bridge", dogTrustBridge.Name)

	assert.Equal(t, 0, len(dogTrustBridge.Friends))
	dogTrustBridge.addFriend("wework")
	assert.Equal(t, 1, len(dogTrustBridge.Friends))

	assert.Equal(t, 0, len(dogTrustBridge.FriendsMap))
	dogTrustBridge.addFriend2("wework")
	assert.Equal(t, 1, len(dogTrustBridge.FriendsMap))

}

func changeValue(a *Dog) {
	a.Name = "dd"
}

type Dog struct {
	Name       string
	Friends    []Dog
	FriendsMap map[string]Dog
}

func (d Dog) changeName(name string) {
	d.Name = name
}

func (d *Dog) addFriend(name string) {
	d.Friends = append(d.Friends, Dog{Name: name})
}

func (d Dog) addFriend2(name string) {
	d.FriendsMap[name] = Dog{Name: name}
}

func TestIntPointer(t *testing.T) {
	a := new(int)
	fmt.Println(a == nil)

	fn := func(ap *int) {
		*ap = 0
	}

	fn(a)

	fmt.Println(a == nil)
}
