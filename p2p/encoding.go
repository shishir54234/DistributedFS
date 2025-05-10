package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, any) error
}
type GO8Decoder struct{}

func (decoder *GO8Decoder) Decode(r io.Reader, msg any) error {
	return gob.NewDecoder(r).Decode(msg)
}
