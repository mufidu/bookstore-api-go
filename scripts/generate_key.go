package main

import (
	auth "bookstore-api-go/pkg/auth/user"
	"fmt"
)

func main() {
	fmt.Println(auth.GenerateRandomKey())
}
