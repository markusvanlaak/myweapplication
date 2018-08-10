package main

import (
	"net/http"
	"fmt"
)

func main() {
	res, _ := http.Get("http://localhost:8081/person")
	fmt.Println(res.Status, res.Header)
	}
