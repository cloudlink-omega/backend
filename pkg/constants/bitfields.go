package constants

/*
	Bitfield constants
	These constants are used to define each bit used for the "state" column value in each table outlined below.
*/

// User flags
const (
	USER_IS_EMAIL_REGISTERED uint = 0 // If the first bit is set, the welcome email has been sent successfully.
	USER_IS_ACTIVE           uint = 1 // If the second bit is set, the account is activated and can be used (user has verified email).
	USER_IS_BLOCKED          uint = 2 // If the third bit is set, the account has been disabled.
	USER_IS_BANNED           uint = 3 // If the fourth bit is set, the account has been banned.
	USER_IS_EMAIL_DISABLED   uint = 4 // If the fifth bit is set, this will disable sending emails to the user (i.e. email needs to be changed manually, wrong email, etc).
	_                        uint = 5 // _ bit values are reserved for future use.
	_                        uint = 6
	USER_IS_ADMIN            uint = 7 // If the last bit is set, the user is a server admin.
)

// Session flags
const (
	SESSION_IS_ACTIVE uint = 0 // If the first bit is set, the session is active (set false to revoke the session).
	SESSION_PERSIST   uint = 1 // If the second bit is set, the session should have no TTL and should persist.
	_                 uint = 2
	_                 uint = 3
	_                 uint = 4 // _ bit values are reserved for future use.
	_                 uint = 5
	_                 uint = 6
	_                 uint = 7
)

// Admin account flags
const (
	ADMIN_IS_ACTIVE uint = 0 // If the first bit is set, the admin is active (set false to revoke access).
	_               uint = 1
	_               uint = 2
	_               uint = 3
	_               uint = 4 // _ bit values are reserved for future use.
	_               uint = 5
	_               uint = 6
	_               uint = 7
)

// Developer account flags
const (
	DEVELOPER_IS_ACTIVE   uint = 0 // If the first bit is set, the developer account is active (set false to revoke access).
	DEVELOPER_IS_VERIFIED uint = 1 // If the second bit is set, the developer account has been verified by the admin and is ready for use.
	_                     uint = 2
	_                     uint = 3
	_                     uint = 4 // _ bit values are reserved for future use.
	_                     uint = 5
	_                     uint = 6
	_                     uint = 7
)

// Game flags
const (
	GAME_IS_ACTIVE       uint = 0 // If the first bit is set, the game ID is active and will permit players to join (set false to deny access).
	GAME_IS_VERIFIED     uint = 1 // If the second bit is set, the game ID is verified and will permit protected features to be used (i.e. currency, stores, etc.).
	GAME_SUPPORTS_DISK   uint = 2 // If the third bit is set, the game supports cloud save slots.
	GAME_SUPPORTS_VOICE  uint = 3 // If the fourth bit is set, the game supports voice chat.
	GAME_IS_MATURE       uint = 4 // If the fifth bit is set, the game is considered mature.
	GAME_USES_OTHER_AUTH uint = 5 // If the sixth bit is set, the game uses other authentication methods.
	_                    uint = 6 // _ bit values are reserved for future use.
	_                    uint = 7
)

// Developer member flags
const (
	DEVMEMBER_IS_ACTIVE uint = 0 // If the first bit is set, the developer member is active (set false to revoke access).
	_                   uint = 1
	_                   uint = 2
	_                   uint = 3
	_                   uint = 4 // _ bit values are reserved for future use.
	_                   uint = 5
	_                   uint = 6
	DEVMEMBER_IS_OWNER  uint = 7 // If the last bit is set, the developer member is the owner of the developer account.
)
