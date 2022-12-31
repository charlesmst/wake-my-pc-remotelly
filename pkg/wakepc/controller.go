package wakepc

import (
	"context"
	"fmt"
)

type PcController interface {
	Shutdown(ctx context.Context) error
	Wol(ctx context.Context, mac string) error
}

type PcControllerMock struct {
	lastState string
}

func (c *PcControllerMock) Shutdown(ctx context.Context) error {
	c.lastState = "shutdown"
	return nil
}

func (c *PcControllerMock) Wol(ctx context.Context, mac string) error {
	c.lastState = fmt.Sprintf("wol %s", mac)
	return nil
}
