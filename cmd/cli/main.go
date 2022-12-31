package main

import (
	"context"
	"fmt"

	"github.com/charlesmst/wake-my-pc-remotelly/pkg/backend"
	"github.com/charlesmst/wake-my-pc-remotelly/pkg/wakepc"
)

func main() {
	storage := backend.NewFirebaseStorage()

	r, err := storage.FindAll(context.Background())
	if err != nil {
		fmt.Printf("failed to read %v", err)
	}
	printState(r)
}
func printState(pcs []wakepc.PcState) {

	for _, pc := range pcs {
		fmt.Printf("%s\t%s\t%s\n", pc.MacAddress, pc.HostName, pc.RealState())
	}

}
