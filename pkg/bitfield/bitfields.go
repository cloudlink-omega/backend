package bitfield

/*
	Bitfield constants
	These constants are used to store the state of each user in the database.

	USER_BITFIELD type is used for the "users" table.
	SESSION_BITFIELD for the "sessions" table.
	ADMIN_BITFIELD for the "admins" table.
*/

type USER_BITFIELD Bitfield8

const (
	USER_IS_EMAIL_REGISTERED uint8 = 0 // If the first bit is set, the welcome email has been sent successfully.
	USER_IS_EMAIL_VERIFIED   uint8 = 1 // If the second bit is set, the email has been verified.
	USER_IS_ACCOUNT_DISABLED uint8 = 2 // If the third bit is set, the account has been disabled.
	USER_IS_ACCOUNT_BANNED   uint8 = 3 // If the fourth bit is set, the account has been banned.
	_                        uint8 = 4
	_                        uint8 = 5 // _ bit values are reserved for future use.
	_                        uint8 = 6
	USER_IS_ADMIN            uint8 = 7 // If the last bit is set, the user is an admin.
)

type SESSION_BITFIELD Bitfield8

const (
	SESSION_IS_ACTIVE uint8 = 0 // If the first bit is set, the session is active.
	SESSION_PERSIST   uint8 = 1 // If the second bit is set, the session should have no TTL and should persist.
	_                 uint8 = 2
	_                 uint8 = 3
	_                 uint8 = 4 // _ bit values are reserved for future use.
	_                 uint8 = 5
	_                 uint8 = 6
	_                 uint8 = 7
)

type ADMIN_BITFIELD Bitfield8

const (
	ADMIN_IS_ACTIVE uint8 = 0 // If the first bit is set, the admin is active.
	_               uint8 = 1
	_               uint8 = 2
	_               uint8 = 3
	_               uint8 = 4 // _ bit values are reserved for future use.
	_               uint8 = 5
	_               uint8 = 6
	_               uint8 = 7
)
