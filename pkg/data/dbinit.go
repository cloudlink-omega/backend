package data

import (
	"log"

	"github.com/huandu/go-sqlbuilder"
)

func (mgr *Manager) InitDB() {

	// Bypass if in authless mode
	if mgr.AuthlessMode {
		return
	}

	log.Print("[DB] Initializing...")
	mgr.createUsersTable()
	mgr.createDevelopersTable()
	mgr.createGamesTable()
	mgr.createAdminsTable()
	mgr.createSessionsTable()
	mgr.createSavesTable()
	mgr.createGamesAuthorizedOriginsTable()
	mgr.createDeveloperMembersTable()
	mgr.createIPWhitelistTable()
	mgr.createIPBlocklistTable()
	mgr.createMagicLinksTable()
	log.Print("[DB] Ready!")
}

func (mgr *Manager) buildTable(tablename string, sb *sqlbuilder.CreateTableBuilder) {
	query, args := sb.Build()
	if _, err := mgr.DB.Query(query, args...); err != nil {
		log.Printf(`[DB] Failed to prepare table "%s": %s`, tablename, err)
	}
}

func (mgr *Manager) createGamesTable() {
	sb := sqlbuilder.NewCreateTableBuilder()
	sb.CreateTable("games").IfNotExists().
		Define(
			`id`,
			`CHAR(26) PRIMARY KEY UNIQUE NOT NULL`, // ULID string
		).
		Define(
			`developerid`,
			`CHAR(26) NOT NULL REFERENCES developers(id) ON DELETE CASCADE`, // ULID string
		).
		Define(
			`name`,
			`TINYTEXT NOT NULL DEFAULT ''`, // 255 maximum length
		).
		Define(
			`state`,
			`TINYINT unsigned NOT NULL DEFAULT 0`,
		).
		Define(
			`created`,
			`BIGINT NOT NULL DEFAULT CURRENT_TIMESTAMP`, // UNIX Timestamp
		)
	mgr.buildTable("games", sb)
}

func (mgr *Manager) createDevelopersTable() {
	sb := sqlbuilder.NewCreateTableBuilder()
	sb.CreateTable("developers").IfNotExists().
		Define(
			`id`,
			`CHAR(26) PRIMARY KEY UNIQUE NOT NULL`, // ULID string
		).
		Define(
			`name`,
			`TINYTEXT NOT NULL DEFAULT ''`, // 255 maximum length
		).
		Define(
			`state`,
			`TINYINT unsigned NOT NULL DEFAULT 0`,
		).
		Define(
			`created`,
			`BIGINT NOT NULL DEFAULT UNIX_TIMESTAMP()`, // UNIX Timestamp
		)
	mgr.buildTable("developers", sb)
}

func (mgr *Manager) createUsersTable() {
	sb := sqlbuilder.NewCreateTableBuilder()
	sb.CreateTable("users").IfNotExists().
		Define(
			`id`,
			`CHAR(26) PRIMARY KEY UNIQUE NOT NULL`, // ULID string
		).
		Define(
			`username`,
			`TINYTEXT UNIQUE NOT NULL DEFAULT ''`, // 255 maximum length
		).
		Define(
			`password`,
			`TINYTEXT NOT NULL DEFAULT ''`, // Scrypt hash
		).
		Define(
			`email`,
			`VARCHAR(320) UNIQUE NOT NULL DEFAULT ''`, // Longest (insane) email address is 320 characters long. But why would you do that to yourself?
		).
		Define(
			`created`,
			`BIGINT NOT NULL DEFAULT UNIX_TIMESTAMP()`, // UNIX Timestamp
		)
	mgr.buildTable("users", sb)
}

func (mgr *Manager) createAdminsTable() {
	sb := sqlbuilder.NewCreateTableBuilder()
	sb.CreateTable("admins").IfNotExists().
		Define(
			`userid`,
			`CHAR(26) NOT NULL REFERENCES users(id) ON DELETE CASCADE`, // ULID string
		).
		Define(
			`state`,
			`TINYINT unsigned NOT NULL DEFAULT 0`,
		).
		Define(
			`created`,
			`BIGINT NOT NULL DEFAULT UNIX_TIMESTAMP()`, // UNIX Timestamp
		)
	mgr.buildTable("admins", sb)
}

func (mgr *Manager) createSessionsTable() {
	sb := sqlbuilder.NewCreateTableBuilder()
	sb.CreateTable("sessions").IfNotExists().
		Define(
			`id`,
			`CHAR(26) PRIMARY KEY UNIQUE NOT NULL`, // ULID string
		).
		Define(
			`userid`,
			`CHAR(26) NOT NULL REFERENCES users(id) ON DELETE CASCADE`, // ULID string
		).
		Define(
			`state`,
			`TINYINT unsigned NOT NULL DEFAULT 0`,
		).
		Define(
			`created`,
			`BIGINT NOT NULL DEFAULT UNIX_TIMESTAMP()`, // UNIX Timestamp
		).
		Define(
			`expires`,
			`BIGINT NOT NULL DEFAULT (UNIX_TIMESTAMP() + 86400)`, // UNIX Timestamp + 24 hours
		).
		Define(
			`origin`,
			`TINYTEXT NOT NULL DEFAULT ''`, // 255 maximum length, IP address
		)
	mgr.buildTable("sessions", sb)
}

func (mgr *Manager) createSavesTable() {
	sb := sqlbuilder.NewCreateTableBuilder()
	sb.CreateTable("saves").IfNotExists().
		Define(
			`userid`,
			`CHAR(26) NOT NULL REFERENCES users(id) ON DELETE CASCADE`, // ULID string
		).
		Define(
			`gameid`,
			`CHAR(26) NOT NULL REFERENCES games(id)`, // ULID string
		).
		Define(
			`slotid`,
			`TINYINT unsigned NOT NULL DEFAULT 0`, // 10 save slots, using 0-9 index.
		).
		Define(
			`contents`,
			`VARCHAR(10000) NOT NULL DEFAULT ''`, // Any desired format, within 10,000 characters
		)
	mgr.buildTable("saves", sb)
}

func (mgr *Manager) createGamesAuthorizedOriginsTable() {
	sb := sqlbuilder.NewCreateTableBuilder()
	sb.CreateTable("games_authorized_origins").IfNotExists().
		Define(
			`gameid`,
			`CHAR(26) NOT NULL REFERENCES games(id)`, // ULID string
		).
		Define(
			`origin`,
			`TINYTEXT NOT NULL DEFAULT ''`, // IP address
		).
		Define(
			`state`,
			`TINYINT unsigned NOT NULL DEFAULT 0`,
		)
	mgr.buildTable("games_authorized_origins", sb)
}

func (mgr *Manager) createDeveloperMembersTable() {
	sb := sqlbuilder.NewCreateTableBuilder()
	sb.CreateTable("developer_members").IfNotExists().
		Define(
			`developerid`,
			`CHAR(26) NOT NULL REFERENCES developers(id) ON DELETE CASCADE`, // ULID string
		).
		Define(
			`userid`,
			`CHAR(26) NOT NULL REFERENCES users(id) ON DELETE CASCADE`, // ULID string
		).
		Define(
			`description`,
			`TINYTEXT NOT NULL`, // ULID string
		)
	mgr.buildTable("developer_members", sb)
}

func (mgr *Manager) createIPBlocklistTable() {
	sb := sqlbuilder.NewCreateTableBuilder()
	sb.CreateTable("ip_blocklist").IfNotExists().
		Define(
			`address`,
			`TINYTEXT NOT NULL`, // IP address
		)
	mgr.buildTable("ip_blocklist", sb)
}

func (mgr *Manager) createIPWhitelistTable() {
	sb := sqlbuilder.NewCreateTableBuilder()
	sb.CreateTable("ip_whitelist").IfNotExists().
		Define(
			`address`,
			`TINYTEXT NOT NULL`, // IP address
		)
	mgr.buildTable("ip_whitelist", sb)
}

func (mgr *Manager) createMagicLinksTable() {
	sb := sqlbuilder.NewCreateTableBuilder()
	sb.CreateTable("magic_links").IfNotExists().
		Define(
			`id`,
			`CHAR(26) PRIMARY KEY UNIQUE NOT NULL`, // ULID string, used for the magic link ID
		).
		Define(
			`mode`,
			`TINYINT unsigned NOT NULL DEFAULT 255`, // 0-255, magic link mode. See MagicLinkMode constants
		).
		Define(
			`userid`,
			`CHAR(26) NOT NULL REFERENCES users(id) ON DELETE CASCADE`, // ULID string, used for identifying which user the magic link belongs to
		).
		Define(
			`created`,
			`BIGINT NOT NULL DEFAULT UNIX_TIMESTAMP()`, // UNIX Timestamp
		).
		Define(
			`expires`,
			`BIGINT DEFAULT NULL`, // UNIX Timestamp (i.e. security codes) or null (i.e. verification links)
		)
	mgr.buildTable("magic_links", sb)
}
