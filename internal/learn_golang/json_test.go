package learn_golang

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

func TestJsonWithNumber(t *testing.T) {
	jsonStr := `
{

}
	`
	var payload struct {
		Id *int64 `json:"id,string"`
	}
	err := json.Unmarshal([]byte(jsonStr), &payload)
	if nil != err {
		fmt.Println(err)
	}

	fmt.Printf("%+v", payload)
	fmt.Println()
	x, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Println(string(x))

	var p struct {
		Id uint64 
	}
	p.Id = math.MaxUint64 - 1
	bs, err := json.Marshal(&p)
	if err != nil {
		fmt.Println(bs)
	}
	fmt.Println("----")
	fmt.Println(string(bs))
}
