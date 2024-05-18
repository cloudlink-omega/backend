package constants

/*
	Magic link modes
	These constants are used to define each mode used for the "mode" column value in the "magic_links" table.
*/

const (
	LINKMODE_EMAIL     uint8 = 0   // Link mode for verifying or unsubscribing an email. Used for welcome emails.
	LINKMODE_PASSWORD  uint8 = 1   // Link mode for resetting passwords.
	LINKMODE_DEVELOPER uint8 = 2   // Link mode for admin approve/deny developer account requests.
	LINKMODE_UNDEFINED uint8 = 255 // Default link mode.
)
