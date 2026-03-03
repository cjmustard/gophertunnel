package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Fish is sent by the client and/or server. It contains fishing-related data.
type Fish struct {
	// GrippyNibble is a float value associated with the fish packet.
	GrippyNibble float32
	// BigMouth is a boolean value associated with the fish packet.
	BigMouth bool
	// NoTeeth is a boolean value associated with the fish packet.
	NoTeeth bool
}

// ID ...
func (*Fish) ID() uint32 {
	return IDFish
}

func (pk *Fish) Marshal(io protocol.IO) {
	io.Float32(&pk.GrippyNibble)
	io.Bool(&pk.BigMouth)
	io.Bool(&pk.NoTeeth)
}
