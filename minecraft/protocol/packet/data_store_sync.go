package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// DataStoreSync is sent to synchronize data stores between the client and server.
type DataStoreSync struct {
	// Payload is the data store sync payload.
	Payload []byte
	// SerializationMode is the serialization mode used for the payload.
	SerializationMode uint32
}

// ID ...
func (*DataStoreSync) ID() uint32 {
	return IDDataStoreSync
}

func (pk *DataStoreSync) Marshal(io protocol.IO) {
	io.ByteSlice(&pk.Payload)
	io.Uint32(&pk.SerializationMode)
}
