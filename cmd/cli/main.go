package main

import (
	"context"
	"fmt"

	"github.com/charlesmst/wake-my-pc-remotelly/pkg/backend"
)

func main() {
	storage := backend.NewFirebaseStorage()

	r, err := storage.FindAll(context.Background())
	if err != nil{
		fmt.Printf("failed to read %v", err)
	}
	fmt.Printf("%v", r)


}
