package structs

import "github.com/cloudlink-omega/backend/pkg/bitfield"

type UserQuery struct {
	ID       string
	Username string
	Gamertag string
	Email    string
	Created  string
	Password string
	State    bitfield.Bitfield8
}

type BasicUserQuery struct {
	Username string
	Email    string
}
