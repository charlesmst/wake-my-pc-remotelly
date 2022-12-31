package wakepc

import (
	"context"
	"testing"
	"time"
)

func TestPc(t *testing.T) {

	t.Run("reposts status", func(t *testing.T) {

		storage := NewPcStateStorageMock()
		controller := PcControllerMock{
			fakeMac: "fakemac"}
		pc := NewPcDaemon(&storage, &controller)
		ctx, cancel := context.WithCancel(context.Background())
		go pc.Start(ctx)

		time.Sleep(1 * time.Second)

		cancel()

		result, err := storage.Find(context.Background(), controller.fakeMac)
		if err != nil {
			t.Fatalf("got error findig state %v", err)
		}

		if result.State != On {
			t.Fatalf("got wrong state %v", result)
		}
	})
	t.Run("shutdown pc", func(t *testing.T) {

		storage := NewPcStateStorageMock()
		controller := PcControllerMock{}
		pc := NewPcDaemon(&storage, &controller)
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		go pc.Start(ctx)
		defer cancel()

		time.Sleep(1 * time.Millisecond)
		if storage.listener == nil {
			t.Fatalf("didn't start listener")
		}

		storage.listener <- PcCommandEvent{Command: Shutdown}

		time.Sleep(100 * time.Millisecond)
		if controller.lastState != "shutdown" {
			t.Fatalf("pc didn't shut down %v", controller.lastState)
		}

	})

	t.Run("wol other pc pc", func(t *testing.T) {
		storage := NewPcStateStorageMock()
		controller := PcControllerMock{}
		pc := NewPcDaemon(&storage, &controller)

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		go pc.Start(ctx)
		defer cancel()

		time.Sleep(1 * time.Millisecond)

		if storage.listener == nil {
			t.Fatalf("didn't start listener")
		}

		storage.listener <- PcCommandEvent{Command: Wol, Args: []string{"somepc"}}

		time.Sleep(100 * time.Millisecond)
		if controller.lastState != "wol somepc" {
			t.Fatalf("pc didn't wol %v", controller.lastState)
		}

	})

}
