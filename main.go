package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	b, err := ioutil.ReadFile("check")
	if err != nil {
		log.Fatal(err)
	}
	if bytes.HasPrefix(b, []byte("0")) {
		fmt.Println("FAILED. Got:", string(b))
		os.Exit(1)
	}
	fmt.Println("OK. Got:", string(b))
}
