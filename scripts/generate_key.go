package main

import (
	auth "bookstore-api-go/pkg/api/admin/auth"
	"fmt"
)

func main() {
	fmt.Println(auth.GenerateRandomKey())
}
