package wakepc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type PcController interface {
	FindState(ctx context.Context) (PcState, error)
	Shutdown(ctx context.Context) error
	Restart(ctx context.Context) error
	Wol(ctx context.Context, mac string) error
}

type PcControllerMock struct {
	lastState string
	fakeMac   string
}

func (c *PcControllerMock) FindState(ctx context.Context) (PcState, error) {
	return PcState{MacAddress: c.fakeMac, State: On}, nil
}

func (c *PcControllerMock) Shutdown(ctx context.Context) error {
	c.lastState = "shutdown"
	return nil
}

func (c *PcControllerMock) Restart(ctx context.Context) error {
	c.lastState = "restart"
	return nil
}

func (c *PcControllerMock) Wol(ctx context.Context, mac string) error {
	c.lastState = fmt.Sprintf("wol %s", mac)
	return nil
}

type MacController struct {
}

var _ PcController = &MacController{}

func (c *MacController) FindState(ctx context.Context) (PcState, error) {
	return readState()
}

func (c *MacController) Shutdown(ctx context.Context) error {
	cmd := exec.Command("shutdown", "-h", "now")
	err := cmd.Run()
	return err
}

func (c *MacController) Restart(ctx context.Context) error {
	cmd := exec.Command("shutdown", "-r", "now")
	err := cmd.Run()
	return err
}

func (c *MacController) Wol(ctx context.Context, mac string) error {
	cmd := exec.Command("wakeonlan", mac)
	err := cmd.Run()
	return err
}

type LinuxController struct {
}

var _ PcController = &LinuxController{}

func (c *LinuxController) FindState(ctx context.Context) (PcState, error) {
	return readState()
}

func (c *LinuxController) Shutdown(ctx context.Context) error {
	cmd := exec.Command("shutdown", "-n", "now")
	err := cmd.Run()
	return err
}

func (c *LinuxController) Restart(ctx context.Context) error {
	cmd := exec.Command("reboot")
	err := cmd.Run()
	return err
}

func (c *LinuxController) Wol(ctx context.Context, mac string) error {
	cmd := exec.Command("wol", mac)
	err := cmd.Run()
	return err
}

type WindowsController struct {
}

var _ PcController = &WindowsController{}

func (c *WindowsController) FindState(ctx context.Context) (PcState, error) {
	return readState()
}

func (c *WindowsController) Shutdown(ctx context.Context) error {
	return exec.Command("cmd", "/C", "shutdown", "/s").Run()
}

func (c *WindowsController) Restart(ctx context.Context) error {
	return exec.Command("cmd", "/C", "shutdown", "/r").Run()
}

func (c *WindowsController) Wol(ctx context.Context, mac string) error {
	return fmt.Errorf("unsupported wol command")
}

func readState() (PcState, error) {

	mac, err := getMacAddr()
	if err != nil {
		return PcState{}, fmt.Errorf("could not get mac address %w", err)
	}
	host, _ := os.Hostname()
	return PcState{
		MacAddress: mac[0],
		HostName:   host,
		State:      On,
		ReadTime:   time.Now().Unix(),
	}, nil
}

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

func ResolveController() PcController {
	switch runtime.GOOS {
	case "linux":
		return &LinuxController{}

	case "darwin":
		return &MacController{}

	case "windows":
		return &WindowsController{}
	default:
		log.Fatalf("unsupported OS")
		return nil
	}
}
