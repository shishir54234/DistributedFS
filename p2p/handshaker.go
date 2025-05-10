package p2p

import "errors"

var ErrInvalidHandhsake = errors.New("invalid handshake from peer")

type HandShakeFunc func(any) error

func NOPHandShakeFunc(any) error { return nil }
