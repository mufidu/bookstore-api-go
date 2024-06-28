package main

import (
	"bookstore-api-go/pkg/api/admin"
	"fmt"
)

func main() {
	fmt.Println(admin.GenerateRandomKey())
}
