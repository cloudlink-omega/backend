package structs

// Managing lobbies
type LobbyConfigStore struct {
	ID                   string
	MaximumPeers         int
	AllowHostReclaim     bool
	AllowPeersToReclaim  bool
	CurrentOwnerID       uint64 // For client manager tracking only
	CurrentOwnerULID     string // For signaling
	CurrentOwnerUsername string // For lobby manager
	Password             string // Scrypt hash or empty
	IsPublic             bool
	Locked               bool
}
