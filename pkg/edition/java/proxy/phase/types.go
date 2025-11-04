package phase

// The connection types supported.
var (
	// Undetermined indicates that the connection has yet to reach the
	// point where we have a definitive answer as to what type of connection we have.
	Undetermined ConnectionType = &connType{
		initialClientPhase:  VanillaClientPhase,
		initialBackendPhase: UnknownBackendPhase,
	}
	// Vanilla indicates that a connection is a vanilla connection.
	Vanilla ConnectionType = &connType{
		initialClientPhase:  VanillaClientPhase,
		initialBackendPhase: VanillaBackendPhase,
	}
	// Undetermined17 is a 1.7 version connection type.
	Undetermined17 ConnectionType = &connType{
		initialClientPhase:  NotStartedLegacyForgeHandshakeClientPhase,
		initialBackendPhase: UnknownBackendPhase,
	}
	// LegacyForge indicates that the connection is a 1.8-1.12 Forge connection.
	LegacyForge ConnectionType = &legacyForgeConnType{
		connType: &connType{
			initialClientPhase:  NotStartedLegacyForgeHandshakeClientPhase,
			initialBackendPhase: NotStartedLegacyForgeHandshakeBackendPhase,
		},
	}
	ModernForge ConnectionType = &connType{
		initialClientPhase:  VanillaClientPhase,
		initialBackendPhase: VanillaBackendPhase,
	}
)

// ConnectionType is a connection type.
type ConnectionType interface {
	InitialClientPhase() ClientConnectionPhase
	InitialBackendPhase() BackendConnectionPhase
}

type connType struct {
	initialClientPhase  ClientConnectionPhase
	initialBackendPhase BackendConnectionPhase
}

var _ ConnectionType = (*connType)(nil)

func (c *connType) InitialClientPhase() ClientConnectionPhase {
	return c.initialClientPhase
}

func (c *connType) InitialBackendPhase() BackendConnectionPhase {
	return c.initialBackendPhase
}

type legacyForgeConnType struct{ *connType }
