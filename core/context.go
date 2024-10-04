package core

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const (
	networkTimeout = 2 * time.Second
)

var (
	ErrNetworkTimeout = errors.New("network timeout")
)

func NewContextWithNetworkDeadline() (context.Context, context.CancelFunc) {
	return context.WithDeadlineCause(context.Background(), time.Now().Add(networkTimeout), ErrNetworkTimeout)
}

func NewCallOptsWithNetworkDeadline() (*bind.CallOpts, context.CancelFunc) {
	ctx, cancel := NewContextWithNetworkDeadline()
	return &bind.CallOpts{Context: ctx}, cancel
}
