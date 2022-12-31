package wakepc

import (
	"context"
	"testing"
	"time"
)

func TestPc(t *testing.T) {

	storage := NewPcStateStorageMock()
	controller := PcControllerMock{}
	pc := NewPcDaemon(&storage, &controller)

	t.Run("reposts status", func(t *testing.T) {

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
	})
	t.Run("shutdown pc", func(t *testing.T) {

		ctx, cancel := context.WithCancel(context.Background())
		go pc.Start(ctx)
		defer cancel()

		if storage.listener == nil {
			t.Fatalf("didn't start listener")
		}

		storage.listener <- PcCommand(Shutdown)

		time.Sleep(100 * time.Millisecond)
		if controller.lastState != "shutdown" {
			t.Fatalf("pc didn't shut down %v", controller.lastState)
		}

	})

}
