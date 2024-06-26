package structs

// Declare the packet format for signaling.
type SignalPacket struct {
	Opcode    string    `json:"opcode" validate:"required" label:"opcode"`                       // Required for protocol compliance
	Payload   any       `json:"payload,omitempty" validate:"required" label:"payload"`           // Required for protocol compliance
	Origin    *PeerInfo `json:"origin,omitempty" validate:"omitempty,ulid" label:"origin"`       // Server -> Client, identifies client that sent the message
	Recipient string    `json:"recipient,omitempty" validate:"omitempty,ulid" label:"recipient"` // Client -> Server, identifies client that should receive the message
	Listener  string    `json:"listener,omitempty" validate:"omitempty" label:"listener"`        // For clients to listen to server replies
}

// JSON structure for signaling INIT_OK response.
type InitOK struct {
	User      string `json:"user"`
	Id        string `json:"id"`
	Game      string `json:"game"`
	Developer string `json:"developer"`
}

type HostConfigPacket struct {
	Opcode  string `json:"opcode" validate:"required" label:"opcode"`
	Payload struct {
		LobbyID             string `json:"lobby_id" label:"lobby_id"`
		AllowHostReclaim    bool   `json:"allow_host_reclaim" validate:"boolean" label:"allow_host_reclaim"`
		AllowPeersToReclaim bool   `json:"allow_peers_to_claim_host" validate:"boolean" label:"allow_peers_to_claim_host"`
		MaximumPeers        int    `json:"max_peers" validate:"min=0,max=100" label:"max_peers"`
		Password            string `json:"password" validate:"omitempty,max=128" label:"password"`
		PublicKey           string `json:"pubkey,omitempty" validate:"omitempty" label:"pubkey"`
	} `json:"payload" validate:"required_with=LobbyID AllowHostReclaim AllowPeersToReclaim MaximumPeers" label:"payload"`
}

// Declare the packet format for the CONFIG_PEER signaling command.
type PeerConfigPacket struct {
	Opcode  string `json:"opcode" validate:"required" label:"opcode"`
	Payload struct {
		LobbyID   string `json:"lobby_id" validate:"required" label:"lobby_id"`
		Password  string `json:"password" validate:"omitempty,max=128" label:"password"`
		PublicKey string `json:"pubkey,omitempty" validate:"omitempty" label:"pubkey"`
	} `json:"payload" validate:"required_with=LobbyID" label:"payload"`
}

// Declare the packet format for the NEW_HOST signaling event.
type NewHostParams struct {
	ID        string `json:"id"`
	User      string `json:"user"`
	LobbyID   string `json:"lobby_id"`
	PublicKey string `json:"pubkey,omitempty"`
}

// Declare the packet format for the NEW_PEER signaling event.
type NewPeerParams struct {
	ID        string `json:"id"`
	User      string `json:"user"`
	PublicKey string `json:"pubkey,omitempty"`
}

type RootError struct {
	Errors []map[string]string `json:"Validation error"`
}

type PeerInfo struct {
	ID   string `json:"id"`
	User string `json:"user"`
}

type LobbyInfo struct {
	LobbyHostID       string `json:"lobby_host_id"`
	LobbyHostUsername string `json:"lobby_host_username"`
	MaximumPeers      int    `json:"max_peers"`
	CurrentPeers      int    `json:"current_peers"`
}
