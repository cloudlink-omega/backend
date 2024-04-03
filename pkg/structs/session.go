package structs

type Session struct {
	UGI     string // ULID
	UserID  string // ULID
	State   int    // Bitfield, spec TBD
	Created int64  // Timestamp as UNIX time
	Expiry  int64  // Timestamp as UNIX time
	Origin  string // Tinytext
}
