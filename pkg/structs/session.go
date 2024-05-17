package structs

import "github.com/cloudlink-omega/backend/pkg/bitfield"

type Session struct {
	UGI     string             // ULID
	UserID  string             // ULID
	State   bitfield.Bitfield8 // Bitfield
	Created int64              // Timestamp as UNIX time
	Expiry  int64              // Timestamp as UNIX time
	Origin  string             // Tinytext
}
