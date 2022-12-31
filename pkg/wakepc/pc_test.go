package wakepc

import (
	"context"
	"testing"
	"time"
)

func TestPc(t *testing.T) {

	storage := NewPcStateStorageMock()
	pc := NewPcDaemon(&storage)

	ctx, cancel := context.WithCancel(context.Background())
	go pc.Start(ctx)
	time.Sleep(1 * time.Second)
	cancel()

	result, err := storage.Find(context.Background(), pc.mac)
	if err != nil {
		t.Fatalf("got error findig state %v", err)
	}

	if result.State != On {
		t.Fatalf("got wrong state %v", result)
	}

}
