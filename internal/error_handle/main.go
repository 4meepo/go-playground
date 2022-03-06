package main

import (
	"fmt"
	"log"
)

func main() {

	err := MustError()
	if err != nil {
		fmt.Println("------")
		log.Printf("%+v", err)
	}

}
