package wakepc

func NewShutdownCommand(mac string) PcCommandEvent {
	return PcCommandEvent{
		Command: Shutdown,
		Target:  mac,
	}
}

func NewRestartCommand(mac string) PcCommandEvent {
	return PcCommandEvent{
		Command: Restart,
		Target:  mac,
	}
}

func NewWolCommand(mac string) PcCommandEvent {
	return PcCommandEvent{
		Command: Wol,
		Target:  mac,
		Args:    []string{mac},
	}
}
