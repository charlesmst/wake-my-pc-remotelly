package wakepc

import (
	"context"
	"os/exec"
)

type LinuxController struct {
}

var _ PcController = &LinuxController{}

func (c *LinuxController) FindState(ctx context.Context) (PcState, error) {
	return readState()
}

func (c *LinuxController) Shutdown(ctx context.Context) error {
	cmd := exec.Command("shutdown -n now")
	err := cmd.Run()
	return err
}

func (c *LinuxController) Wol(ctx context.Context, mac string) error {
	cmd := exec.Command("wol", mac)
	err := cmd.Run()
	return err
}
