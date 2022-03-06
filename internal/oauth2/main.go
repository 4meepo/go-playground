package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

// https://github.com/golang-jwt/jwt
func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userid": "111222233",
		"name":   "yifei",
		"nam1e":  "yifei",
		"name2":  "yifei",
		"name3":  "yifei",
		"na4me":  "yifei",
		"na5me":  "yifei",
	})
	signedString, err := token.SignedString([]byte("sdfsdfsdfdsfsdfdsf"))
	if err != nil {
		panic(err)
	}
	fmt.Println(signedString)

}
