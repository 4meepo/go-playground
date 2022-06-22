package learn_golang

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNilInt(t *testing.T) {
	var data struct {
		IsMember *bool `json:"Is_Member,omitempty"`
		Count    *int64
	}

	bytes, _ := json.Marshal(&data)
	fmt.Println(string(bytes))

	if data.IsMember == nil{
		fmt.Println(" the data.IsMember is nil")
	}
}
