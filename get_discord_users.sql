SELECT u.*, p.id AS "discord id" FROM users as u, users_discord as p
JOIN users_discord ON u.id = p.user_id