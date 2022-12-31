package wakepc

import (
	"context"
)

type Daemon interface {
	Start(ctx context.Context)
}
