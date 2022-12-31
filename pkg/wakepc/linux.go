package  wakepc

import (
	"context"
	"fmt"
	"os/exec"
)

type LinuxController struct {
}

var _ PcController = &LinuxController{}

func (c *LinuxController) FindState(ctx context.Context) (PcState, error) {
	mac, err := getMacAddr()
	if err != nil {
		return PcState{}, fmt.Errorf("could not get mac address %w", err)
	}
	return PcState{MacAddress: mac[0], State: On}, nil
}

func (c *LinuxController) Shutdown(ctx context.Context) error {
	cmd := exec.Command("shutdown -f now")
	err := cmd.Run()
	return err
}

func (c *LinuxController) Wol(ctx context.Context, mac string) error {
	cmd := exec.Command("wol", mac)
	err := cmd.Run()
	return err
}
