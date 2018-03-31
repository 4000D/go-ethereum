package plasma

import (
	"context"
	"sync"
	"time"
)

// PlasmaAPI provides an API to access Plasma related information.
type PlasmaAPI struct {
}

// PlasmaOperatorAPI provides an API for plasma operator.
type PlasmaOperatorAPI struct {
}

type PublicPlasmaAPI struct {
	pls *Plasma

	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.
}

func NewPlasmaAPI() PlasmaAPI {
	return PlasmaAPI{}
}

func NewPlasmaOperatorAPI() PlasmaOperatorAPI {
	return PlasmaOperatorAPI{}
}

func NewPublicPlasmaAPI(pls *Plasma) *PublicPlasmaAPI {
	api := &PublicPlasmaAPI{
		pls:      pls,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the Whisper sub-protocol version.
func (api *PublicPlasmaAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}

func (api *PublicPlasmaAPI) Deposit(ctx context.Context) string {
	return ""
}
