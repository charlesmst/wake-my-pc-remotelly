package wakepc

import "context"

type PcController interface {
	Shutdown(ctx context.Context) error
}

type PcControllerMock struct {
	lastState string
}

func (c *PcControllerMock) Shutdown(ctx context.Context) error {
	c.lastState = "shutdown"
	return nil
}
